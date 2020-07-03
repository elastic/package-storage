# Here the version of the registry is specified this storage branch uses.
# It should always be a specific version to make sure builds are reproducible.

FROM docker.elastic.co/package-registry/distribution:production AS production
FROM docker.elastic.co/package-registry/distribution:staging AS staging

FROM docker.elastic.co/package-registry/package-registry:0579a6edb887c957c0fa64fc8ae82ca3f205a63b
LABEL package-registry=0579a6edb887c957c0fa64fc8ae82ca3f205a63b

COPY --from=production /packages/production /packages/production
COPY --from=staging /packages/production /packages/staging

# Adds specific config and packages
COPY deployment/package-registry.yml /package-registry/config.yml
COPY packages /packages/staging

WORKDIR /package-registry

# Sanity check on the packages. If packages are not valid, container does not even build.
RUN ./package-registry -dry-run