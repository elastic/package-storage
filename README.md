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
