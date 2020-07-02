# Here the version of the registry is specified this storage branch uses.
# It should always be a specific version to make sure builds are reproducible.
ARG PACKAGE_REGISTRY=41c150c8020efc53ab16e3bba774c62a419b51ea
FROM docker.elastic.co/package-registry/package-registry:${PACKAGE_REGISTRY}

LABEL package-registry=${PACKAGE_REGISTRY}

# Adds specific config and packages
COPY deployment/package-registry.yml /registry/config.yml
COPY packages /packages/production

# Cleanup
RUN rm -r packages

# Sanity check on the packages. If packages are not valid, container does not even build.
RUN ./package-registry -dry-run