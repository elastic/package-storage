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