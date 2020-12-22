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
    BIN_DIR = "${env.HOME}/bin"
    KUBECTL = "${env.BIN_DIR}/kubectl"
    CLOUDSDK_ROOT_DIR = "${env.BIN_DIR}"
    GCLOUD_VERSION = "320.0.0"
    JOB_GIT_CREDENTIALS = "f6c7695a-671e-4f4f-a331-acdce44ff9ba"
    CREDENTIALS_FILE = 'credentials.json'
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
    choice(choices: ['none', 'snapshot', 'staging', 'prod', 'experimental', '7-9'], description: 'Environment to Rollout.', name: 'environment')
  }
  stages {
    stage('Rollout') {
      when {
        expression {
          return params.environment != 'none'
        }
      }
      environment{
        PACKAGE_REGISTRY_DEPLOYMENT_NAME = "package-registry-${params.environment}-vanilla"
      }
      steps {
        withPackageRegistryEnv(secret: 'secret/observability-team/ci/package-registry-deployment'){
          installGcloud()
          installKubectl()
          withGCPCredentials(secret: "secret/gce/${GOOGLE_PROJECT}/service-account/package-registry-rollout"){
            sh(label: "Rollout ${PACKAGE_REGISTRY_DEPLOYMENT_NAME} deployment", script: '''
              ${KUBECTL} -n package-registry rollout restart deployment ${PACKAGE_REGISTRY_DEPLOYMENT_NAME}
            ''')
          }
        }
      }
    }
  }
}

def getVaultSecretRetry(Map args){
  def secret = args.containsKey('secret') ? args.secret : error('Secret not valid')
  return getVaultSecret(secret: secret)
}

def withPackageRegistryEnv(Map args, Closure body){
  def jsonValue = getVaultSecretRetry(args)?.data
  withEnvMask(vars: [
    [var: "GOOGLE_PROJECT", password: jsonValue.google_project],
    [var: "REGION", password: jsonValue.region],
    [var: "CLUSTER_CREDENTIALS_NAME", password: jsonValue.cluster_credentials_name],
    [var: "KUBECONFIG", password: "${HOME}/.kubeconfig"],
  ]){
    withEnv([
      "PATH+GCLOUD=${CLOUDSDK_ROOT_DIR}/google-cloud-sdk/bin",
    ]){
      body()
    }
  }
}

def withGCPCredentials(Map args, Closure body){
  def jsonValue = getVaultSecretRetry(args)?.data
  writeFile(file: "${CREDENTIALS_FILE}", text: jsonValue.credentials)
  sh(label: 'Activate GCP credentials', script: '''
    gcloud auth activate-service-account --key-file ${CREDENTIALS_FILE}
    gcloud --project=${GOOGLE_PROJECT} container clusters get-credentials ${CLUSTER_CREDENTIALS_NAME} --region ${REGION}
  ''')
  body()
  sh(label: 'delete credentials', script: 'rm -fr ${CREDENTIALS_FILE} ${KUBECONFIG}')
}

def installKubectl(){
  sh(label: 'Install Kubectl', script: '''
  curl -Lo ${KUBECTL} https://storage.googleapis.com/kubernetes-release/release/v1.19.0/bin/linux/amd64/kubectl
  chmod +x ${KUBECTL}
  ''')
}

def installGcloud(){
  sh(label: 'Install gcloud', script: '''#!/bin/bash
    set -eo pipefail
    ARCH=$(uname|tr '[:upper:]' '[:lower:]')

    mkdir -p "${CLOUDSDK_ROOT_DIR}"
    cd "${CLOUDSDK_ROOT_DIR}"
    curl -sSLo google-cloud-sdk.tar.gz \
      https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-${GCLOUD_VERSION}-${ARCH}-x86_64.tar.gz
    tar zxf google-cloud-sdk.tar.gz google-cloud-sdk
    "${CLOUDSDK_ROOT_DIR}/google-cloud-sdk/install.sh" -q
  ''')
}
