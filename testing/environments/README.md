Before using the Package Registry, remember to `mage build` the project to prepare the volume with packages
(`public` directory).

Refresh docker images:

```bash
$ docker-compose -f snapshot.yml pull
```

Run docker containers (Elasticsearch, Kibana, Package Registry):

```bash
$ docker-compose -f snapshot.yml -f local.yml up --force-recreate --build
```
