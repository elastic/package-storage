#!/usr/bin/env groovy
// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
@Library('apm@current') _

pipeline {
  agent { label 'linux && immutable' }
  environment {
    REPO = "package-storage"
    NOTIFY_TO = credentials('notify-to')
    PIPELINE_LOG_LEVEL = 'INFO'
    LANG = "C.UTF-8"
    LC_ALL = "C.UTF-8"
    HOME = "${env.WORKSPACE}"
    PATH = "${env.HOME}/bin:${env.WORKSPACE}/src/.ci/scripts:${env.PATH}"
    JOB_GIT_CREDENTIALS = "f6c7695a-671e-4f4f-a331-acdce44ff9ba"
    TAG = "${params.commit}"
    BRANCH_NAME = "${params.environment}"
  }
  options {
    buildDiscarder(logRotator(numToKeepStr: '20', artifactNumToKeepStr: '20', daysToKeepStr: '30'))
    timestamps()
    ansiColor('xterm')
    disableResume()
    durabilityHint('PERFORMANCE_OPTIMIZED')
    timeout(time: 2, unit: 'HOURS')
    disableConcurrentBuilds()
  }
  parameters {
    choice(choices: ['snapshot', 'staging', 'production'], description: 'Branch and environment to Release the distribution.', name: 'environment')
    string(name: 'commit', defaultValue: "", description: "Commit SHA1/Branch tag of the Docker image.")
  }
  stages {
    stage('Select Docker Image'){
      when {
        expression { return '' == env.TAG }
      }
      steps {
        gitCheckout(basedir: 'src', branch: "${env.BRANCH_NAME}",
          repo: "git@github.com:elastic/${REPO}.git",
          credentialsId: "${env.JOB_GIT_CREDENTIALS}"
        )
        dir('src'){
          sh(label: 'Get latest commits', script: '''
            echo "[" > latest-changes.json
            git log -n 5 --pretty=format:'{"sha":"%H","msg":"%s","author":"%an"},' >> latest-changes.json
            echo "]" >> latest-changes.json
          ''')
          script {
            def json = readJSON(file: 'latest-changes.json')
            def options = json?.collect{item -> "${item.msg}-${item.sha}"}
            def choose = input(message: "Choose the version to release to ${env.BRANCH_NAME}",
                        parameters: [
                          choice(choices: options, 
                            description: 'Select the commit to release',
                            name: 'tag')
                        ]
                      )
            setEnvVar('TAG', choose[-40..-1])
            echo 
          }
        }
      }
    }
    stage('Release') {
      steps {
        deleteDir()
        gitCheckout(basedir: 'infra', branch: "master",
          repo: "git@github.com:elastic/infra.git",
          credentialsId: "${env.JOB_GIT_CREDENTIALS}")
        dir("infra"){
          sshagent(["${env.JOB_GIT_CREDENTIALS}"]) {
            sh(label: 'Create local branch', script: """
              git checkout -b ${REPO}-release
            """)
            script {
              def ymlFile = "terraform/providers/gcp/env/elastic-apps-web/helm/values/package-registry/package-registry-${env.BRANCH_NAME}.yaml"
              def yml = readYaml(file: ymlFile)
              yml.image.tag = "${env.TAG}"
              writeYaml(file: ymlFile, data: yml, overwrite: true)
              setEnvVar('DOCKER_IMAGE', yml.image.repository)
            }
            sh(label: 'Git config', script: """
              git config --global user.name observability-robots
              git config --global user.email observability-robots@users.noreply.github.com
            """)
            sh(label: 'Commit local changes', script: "git commit -a -m '[package-registry]: Release ${env.TAG} to ${env.BRANCH_NAME}'")
            sh(label: 'Push changes', script: "git push origin package-registry-${env.REPO}-release --force")
            sh(label: 'Check Docker image', script: "docker pull ${env.DOCKER_IMAGE}:${env.TAG}")
            githubCreatePullRequest(
              title: "[DO NO MERGE][package-registry]: Release ${env.TAG} to ${env.BRANCH_NAME}",
              base: "elastic:master",
              labels: 'area:web',
              assign: 'elasticmachine',
              reviewer: 'observability-robots',
              description: """
This is an automatic generated PR to release `docker.elastic.co/package-registry/distribution:${env.TAG}` to ${env.BRANCH_NAME}.
  * [ ] Check the TAG ${env.TAG} and the enviroment ${env.BRANCH_NAME} is correct.
  * [ ] Review that the changes are correct.
  * [ ] Launch the infra test `jenkins test this please`
  * [ ] Merge the changes.
  """       )
            
          }
        }
      }
    }
  }
}
