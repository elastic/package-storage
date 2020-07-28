# EXPERIMENTAL: This is only for experimental use

# Package Storage
This is a storage repository for the packages served through the [package registry](https://github.com/elastic/package-registry) service.

It contains 3 branches with the packages for the different environments:

* snapshot
* staging
* production

Here's how these branches relate to repositories and other aspects of packages.

|                       | Snapshot                    | Staging                    | Production                |
|-------------------    |-------------------------    |------------------------    |-----------------------    |
| URL                   | epr-snapshot.elastic.co     | epr-staging.elastic.co     | epr.elastic.co            |
| Add package           | Commit                      | Commit                     | PR                        |
| Version overwrite     | yes                         | if needed                  | no                        |
| Stack Version         | *-SNAPSHOT                  | *-SNAPSHOT                 | Released (candidates)     |
| Registry Version      | fixed version                | Stable release             | Stable release            |
| Branch                | snapshot                    | staging                    | production                |
| Packages              | snapshot+staging+prod       | staging+production         | production                |
| Release               | Manual                      | Manual                  | Manual                 |
| Docker image          | snapshot                    | staging                    | production                |

# Update Package Registry for a distribution

Each distribution (snapshot, staging, production) uses a specific version of the registry. The registry to be used is referenced in the Dockerfile. To update the package registry version, the referenced must be updated.

In the Dockerfile there is a reference looking similar to:

```
ARG PACKAGE_REGISTRY=v0.6.0
FROM docker.elastic.co/package-registry/package-registry:${PACKAGE_REGISTRY}

LABEL package-registry=${PACKAGE_REGISTRY}
```

The above example uses `v0.6.0` as the registry version. Each tagged version in the registry can be used or commits made to package-registry master. A valid registry version would also be `02ca71731cdc092db213cd4c9069db43b74ac01b` as this is a commit hash from master.

In addition to updating the Dockerfile, also the package-registry version on the Golang side must be updated as it is used for some testing. For this, run the following commands:

```
go get github.com/elastic/package-registry@v0.7.0
mage vendor
```

This will update the go module files and pull in the most recent vendor files.

With all the changes above, a pull request can be opened. As soon as the pull request is merged, a new Dockerimage is automatically built by CI and can be deployed.

As each branch/distribution references its own version, each branch has to be updated with the above. It is encourage to keep the branches in sync related to registry versions.

# Package promotion

A package can go through different stages from snapshot to staging to production. To promote a package from one distribution to an other, it must first be added to the new distribution and as soon as it is added, removed from the old distribution. As an example, a package `foo-1.2.3` is in `snapshot`. Now the content is copied over to `staging` and as soon as the package is merged, the package `foo-1.2.3` should be removed from `snapshot`. If the same package version exists in both branches, the first one is taken. In the above case this is `staging`.

As the staging distribution consists of `production + staging` packages and `snapshot` of `production + staging + snapshot` packages, a package that was promoted to the next stage, is always also available in the previous stage.

The above implies, that as soon as the packages for `staging` as an example are built and the Dockerimage is available, also `snapshot` Dockerimage must be rebuilt as it depends on the `staging` one.

# Release distribution

Currently each distribution is released manually but should be release automatically in the future. To release a distribution with a new package, first the building of the Dockerimage must be completed which is automatic. For `snapshot` this pushes a new image to `docker.elastic.co/package-registry/distribution:snapshot`. At the same time, an image is built for each commit hash, for example `docker.elastic.co/package-registry/distribution:48f3935a72b0c5aacc6fec8ef36d559b089a238b`. The distribution specific images are constantly overwritten, the commit hash images stays as is in case an environment is need that does not change.

As soon as the Dockerimage is built and added to the Docker registry, the rollout command for the k8s cluster can be run. For `staging` this looks as following:

```
kubectl rollout restart deployment package-registry-staging-vanilla -n package-registry
```

Each environment has its own deployment. Note: The above command only works if you have access to the specific k8s clusters.

As soon as the rollout is triggered, it takes up to a minute until all containers are updated and from there, on the CDN side it takes the cache time to update the assets. New packages will be available immediately as they were not served by the CDN previously.

# Branches

The current `package-storage` repository has a few branches. This is a quick summary of the use cases of each branch:

* production: Packages for epr.elastic.co
* staging: Packages for epr-staging.elastic.co
* snapshot: Packages for epr-snapshot.elastic.co
* experiemental: Packages for epr-experimental.elastic.co. These packages are served by Kibana 7.8 and will disappear in the future. No updates should happen to this branch.
* master: Contains docs and comment scripts for the distribution branches.