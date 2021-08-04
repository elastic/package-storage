# Here the version of the registry is specified this storage branch uses.
# It should always be a specific version to make sure builds are reproducible.

FROM docker.elastic.co/package-registry/distribution:production AS production

FROM docker.elastic.co/package-registry/package-registry:v0.21.0
LABEL package-registry=v0.21.0

COPY --from=production /packages/production /packages/production

# Adds specific config and packages
COPY deployment/package-registry.yml /package-registry/config.yml
COPY packages /packages/staging

WORKDIR /package-registry

# Sanity check on the packages. If packages are not valid, container does not even build.
RUN ./package-registry -dry-run

# Override CMD to disable package validation (already done).
CMD ["--address=0.0.0.0:8080", "-disable-package-validation"]
