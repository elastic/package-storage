# Here the version of the registry is specified this storage branch uses.
# It should always be a specific version to make sure builds are reproducible.
ARG PACKAGE_REGISTRY=v1.5.1

FROM docker.elastic.co/package-registry/distribution:production AS production

FROM docker.elastic.co/package-registry/package-registry:${PACKAGE_REGISTRY}
LABEL package-registry=${PACKAGE_REGISTRY}

COPY --from=production /packages/production /packages/production

# Adds specific config and packages
COPY deployment/package-registry.yml /package-registry/config.yml
COPY packages /packages/staging

WORKDIR /package-registry

# Sanity check on the packages. If packages are not valid, container does not even build.
RUN ./package-registry -dry-run

# Disable package validation (already done).
ENV EPR_DISABLE_PACKAGE_VALIDATION=true

# Make sure it is accessible out of the container.
ENV EPR_ADDRESS=0.0.0.0:8080
