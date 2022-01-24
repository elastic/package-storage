# Package Storage
This is a storage repository for the packages served through the package registry service.  See the [`package-registry` repository](https://github.com/elastic/package-registry) for basic registry API usage and examples.

The `package-storage` repository contains 3 branches with the packages for the different environments:

* snapshot
* staging
* production

Here's how these branches relate to repositories and other aspects of packages.

|                         | Snapshot                    | Staging                    | Production                |
|-------------------      |-------------------------    |------------------------    |-----------------------    |
| URL                     |  [epr-snapshot.elastic.co](https://epr-snapshot.elastic.co/search) | [epr-staging.elastic.co](https://epr-staging.elastic.co/search) | [epr.elastic.co](https://epr.elastic.co/search) |
| How to add a package      | Commit to elastic/integrations* | [`elastic-package promote`](https://github.com/elastic/elastic-package#elastic-package-promote) | [`elastic-package promote`](https://github.com/elastic/elastic-package#elastic-package-promote) |
| Allow version overwrite?| yes**                        | if needed                  | no                        |
| Allow version removal?  | yes                         | special exceptions only    | only version increments   |
| Stack Ver vs Storage Ver| [main branch](https://github.com/elastic/kibana/blob/main/x-pack/plugins/fleet/server/services/epm/registry/registry_url.ts#L24) | all `-SNAPSHOT` Kibana versions | all shipped or BC versions*** |
| Registry Version        | Fixed dev or latest stable release | Stable release      | Stable release            |
| Branch                  | snapshot                    | staging                    | production                |
| Packages                | snapshot+staging+prod       | staging+production         | production                |
| Release                 | Manual                      | Manual                     | Manual                    |
| Docker image            | snapshot                    | staging                    | production                |

`*` https://github.com/elastic/integrations is the development home of most packages, though not all. Promotion process to package storage repository is discussed below.

** removing packages can cause short-term dev problems when a package is in use and then is deleted, some manual maintenance is expected and this does not reflect user experience in production.

*** for example, during the development of 8.0 as main, all 7.x -SNAPSHOT stack versions (including cloud deploys) will use the `staging` repo to facilitate testing prior to package release, while the formal build candidates and shipped stack versions of Kibana are coded to use the `production` registry.  To update the package storage branch used in a self managed Kibana deploy you can set the below, for example, in the kibana/config/kibana.yml file:
xpack.fleet.registryUrl: "htttp://epr-snapshot.elastic.co/"

# Update Package Registry for a distribution

Each distribution (snapshot, staging, production) uses a specific version of the registry. The registry to be used is referenced in the Dockerfile. To update the package registry version, the referenced must be updated.

In the Dockerfile there is a reference looking similar to:

```
ARG PACKAGE_REGISTRY=v0.6.0
FROM docker.elastic.co/package-registry/package-registry:${PACKAGE_REGISTRY}

LABEL package-registry=${PACKAGE_REGISTRY}
```

The above example uses `v0.6.0` as the registry version. Each tagged version in the registry can be used or commits made to package-registry main. A valid registry version would also be `02ca71731cdc092db213cd4c9069db43b74ac01b` as this is a commit hash from main.

In addition to updating the Dockerfile, also the package-registry version on the Golang side must be updated as it is used for some testing. For this, run the following commands:

```
go get github.com/elastic/package-registry@v0.7.0
mage vendor
```

This will update the go module files and pull in the most recent vendor files.

With all the changes above, a pull request can be opened. As soon as the pull request is merged, a new Docker image is automatically built by CI and can be deployed.

As each branch/distribution references its own version, each branch has to be updated with the above. It is encourage to keep the branches in sync related to registry versions.

# Pull request titles

To easily differentiate PRs against snapshot, staging and production, each PR to one of these branches should be prefixed with `[{branch-name}]`, for example `[snapshot]`. This also makes sure if the same package is promoted to all 3 branches, they don't have the exact same names.

# Package promotion

A package will usually go through different stages from snapshot to staging to production.  It is recommended to use the [`elastic-package` tool](https://github.com/elastic/elastic-package) to achieve package promotion. It is the package developer's burden to fulfill tests on snapshot stage and to coordinate any desired tests on staging prior to promoting to production storage.  Note that production is in use by all 7.10+ released versions of Elastic and testing should reflect this to avoid problems.

So, briefly, to promote a package from one distribution to another, it must first be added to the new distribution and as soon as it is added, removed from the old distribution. As an example, a package `foo-1.2.3` is in `snapshot`. Now the content is copied over to `staging` and as soon as the package is merged, the package `foo-1.2.3` should be removed from `snapshot`. For example see this PR [promoting the `system-0.10.0` package from `snapshot` to `staging`](https://github.com/elastic/package-storage/pull/824) and this corresponding PR [removing the `system-0.10.0` package from `snapshot`](https://github.com/elastic/package-storage/pull/825).

If the same package version exists in both branches, the first one is taken. In the above case this is `staging`.

As the staging distribution consists of `production + staging` packages and `snapshot` of `production + staging + snapshot` packages, a package that was promoted to the next stage, is always also available in the previous stage.

The above implies, that as soon as the packages for `staging` as an example are built and the Docker image is available, also `snapshot` Docker image must be rebuilt as it depends on the `staging` one.

# Release distribution

Currently each distribution is released manually but should be release automatically in the future. To release a distribution with a new package, first the building of the Docker image must be completed which is automatic. For `snapshot` this pushes a new image to `docker.elastic.co/package-registry/distribution:snapshot`. At the same time, an image is built for each commit hash, for example `docker.elastic.co/package-registry/distribution:48f3935a72b0c5aacc6fec8ef36d559b089a238b`. The distribution specific images are constantly overwritten, the commit hash images stays as is in case an environment is need that does not change.

As soon as the Docker image is built and added to the Docker registry, the rollout command for the k8s cluster can be run. For `staging` this looks as following:

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
* production-7.9: Packages for epr-7-9.elastic.co. These packages are served for Kibana 7.9 versions. It is expected that all future changes to packages go into the production branch.
* experimental: Packages for epr-experimental.elastic.co. These packages are served for Kibana 7.8 and will disappear in the future. No updates should happen to this branch.
* main: Contains docs and comment scripts for the distribution branches.

# Tags

At the time of the Elastic Stack release, a tag off the `production` branch is created with an Elastic Stack release version. This will allow environments like Kibana tests to use a specific "known" tag for their testing. An additional benefit is that in the future if users potentially run the registry on premise, these tags can be used to have the on premise registry aligned with the Elastic Stack version instead of constantly pulling the most recent version.

These tags trigger the build of the following docker images (7.9.0 as an example):

```
docker.elastic.co/package-registry/distribution:7.9.0
```

The tag does not contain a `v` prefix like most other version tags do. The reason is that this tag is used for the docker image name where the v prefix is not used.
