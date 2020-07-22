# EXPERIMENTAL: This is only for experimental use

# Package Storage
This is a storage repository for the packages served through the [package registry](https://github.com/elastic/package-registry) service.

There are three package registries in operation and they correspond to the three branches above as follows:

| Package registry | Registry URL | Source branches for packages served by registry |
| ---- | ---- | ---- |
| Snapshot | epr-snapshot.elastic.co | `snapshot`, `staging`, and `production` |
| Staging | epr-staging.elastic.co | `staging` and `production` |
| Production | epr.elastic.co | `production` |

# A day in the life of a package

Ultimately, every package dreams of of making it out in the wide world. But rather than be suddenly cast in the harsh limelight, it must gradually make its way. Here's the journey a package takes from birth to prime time.

## A package is born (or updated)

1. Every new package or a new version of a package starts out life in the `snapshot` branch of this repository. It makes it way into the branch via a PR. The PR must pass CI and be approved by a teammate before it can be merged.

2. Once the PR is merged into the `snapshot` branch, it is eligible to be served by the Snapshot registry. However, this doesn't happen automatically! One must [kick off a build of the `snapshot` branch](https://beats-ci.elastic.co/job/Beats/job/package-storage/job/snapshot/build?delay=0sec), which will publish the Docker image containing packages for the Snapshot registry.

  QUESTION: Once the Docker image is published, does the Snapshot registry immediately / with some delay automatically deploy the new image? Or is there a manual step here needed as well?

3. Confirm that the package is available in the Snapshot registry by visiting https://epr-snapshot.elastic.co/search?package=YOUR_PACKAGE (obviously, replace `YOUR_PACKAGE` with, your package name). Make sure the version of the package being returned is the one you've just been working with.

## A package warms up it's act

1. Before a package can put itself out in the wild world, it must "warm up" it's act in front of friendlier audiences. For this it must progress to the Staging registry.

1. To progress to the Staging registry, create a PR to the `staging` branch of this repository. The PR will contain the package's contents as they are in the `snapshot` branch. The PR must pass CI and be approved by a teammate before it can be merged.

2. Once the PR is merged into the `staging` branch, it is eligible to be served by the Staging (and Snapshot) registries. However, this doesn't happen automatically! One must [kick off a build of the `staging` branch](https://beats-ci.elastic.co/job/Beats/job/package-storage/job/staging/build?delay=0sec), which will publish the Docker image containing packages for the Staging registry.

3. Confirm that the package is available in the Staging registry by visiting https://epr-staging.elastic.co/search?package=YOUR_PACKAGE (obviously, replace `YOUR_PACKAGE` with, your package name). Make sure the version of the package being returned is the one you've just been working with.

3. Once you've confirmed that the package is being served correctly from the Staging registry, it's time to remove it from the `snapshot` branch. Create a PR to the `snapshot` branch of this repository, removing the package's contents. The PR must pass CI and be approved by a teammate before it can be merged.

4. Once the PR is merged into the `snapshot` branch, we must tell the Snapshot registry about this change. To do this, [kick off a build of the `snapshot` branch](https://beats-ci.elastic.co/job/Beats/job/package-storage/job/snapshot/build?delay=0sec), which will publish the Docker image containing the updated list of packages for the Snapshot registry.

## A package goes out into the wild world

1. The golden moment has arrived! cCreate a PR to the `production` branch of this repository. The PR will contain the package's contents as they are in the `staging` branch. The PR must pass CI and be approved by a teammate before it can be merged.

2. Once the PR is merged into the `production` branch, it is eligible to be served by the Production (and Staging and Snapshot) registries. However, this doesn't happen automatically! One must [kick off a build of the `production` branch](https://beats-ci.elastic.co/job/Beats/job/package-storage/job/production/build?delay=0sec), which will publish the Docker image containing packages for the Production registry.

3. Confirm that the package is available in the Staging registry by visiting https://epr.elastic.co/search?package=YOUR_PACKAGE (obviously, replace `YOUR_PACKAGE` with, your package name). Make sure the version of the package being returned is the one you've just been working with.

3. Once you've confirmed that the package is being served correctly from the Production registry, it's time to remove it from the `staging` branch. Create a PR to the `staging` branch of this repository, removing the package's contents. The PR must pass CI and be approved by a teammate before it can be merged.

4. Once the PR is merged into the `staging` branch, we must tell the Staging and Snapshot registries about this change. To do this, [kick off a build of the `staging` branch](https://beats-ci.elastic.co/job/Beats/job/package-storage/job/staging/build?delay=0sec), which will publish the Docker image containing the updated list of packages for the Staging registry. Then, [kick off a build of the `snapshot` branch](https://beats-ci.elastic.co/job/Beats/job/package-storage/job/snapshot/build?delay=0sec), which will publish the Docker image containing the updated list of packages for the Snapshot registry.

5. Hooray, your package is a certified star!
