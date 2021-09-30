# Here the version of the registry is specified this storage branch uses.
# It should always be a specific version to make sure builds are reproducible.
ARG PACKAGE_REGISTRY=v1.3.0
FROM docker.elastic.co/package-registry/package-registry:${PACKAGE_REGISTRY}

LABEL package-registry=${PACKAGE_REGISTRY}

# Adds specific config and packages
COPY deployment/package-registry.yml /package-registry/config.yml
COPY packages /packages/production

# Sanity check on the packages. If packages are not valid, container does not even build.
RUN ./package-registry -dry-run

# Override CMD to disable package validation (already done).
CMD ["--address=0.0.0.0:8080", "-disable-package-validation"]
