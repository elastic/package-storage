# Here the version of the registry is specified this storage branch uses.
# It should always be a specific version to make sure builds are reproducible.

FROM docker.elastic.co/package-registry/distribution:production AS production
FROM docker.elastic.co/package-registry/distribution:staging AS staging

FROM docker.elastic.co/package-registry/package-registry:v0.12.0
LABEL package-registry=v0.12.0

COPY --from=production /packages/production /packages/production
COPY --from=staging /packages/staging /packages/staging

# Adds specific config and packages
COPY deployment/package-registry.yml /package-registry/config.yml
COPY packages /packages/snapshot

WORKDIR /package-registry

# Sanity check on the packages. If packages are not valid, container does not even build.
RUN ./package-registry -dry-run
