# Kibana

## Configuration parameters

If the Kibana instance is using a basepath in its URL, you must set the `basepath` setting for this integration with the same value.

## Compatibility

The `kibana` package works with Kibana 6.7.0 and later.

## Usage for Stack Monitoring

The `kibana` package can be used to collect metrics shown in our Stack Monitoring
UI in Kibana. To enable this usage, set `xpack.enabled: true` on the package config.

## Logs

### Audit

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Date/time when the event originated. This is the date/time extracted from the event, typically representing when the event was generated by the source. If the event source has no original timestamp, this value is typically populated by the first time the event was received by the pipeline. Required field for all events. | date |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| kibana.add_to_spaces | The set of space ids that a saved object was shared to. | keyword |
| kibana.authentication_provider | The authentication provider associated with a login event. | keyword |
| kibana.authentication_realm | The Elasticsearch authentication realm name which fulfilled a login event. | keyword |
| kibana.authentication_type | The authentication provider type associated with a login event. | keyword |
| kibana.delete_from_spaces | The set of space ids that a saved object was removed from. | keyword |
| kibana.lookup_realm | The Elasticsearch lookup realm which fulfilled a login event. | keyword |
| kibana.saved_object.id | The id of the saved object associated with this event. | keyword |
| kibana.saved_object.type | The type of the saved object associated with this event. | keyword |
| kibana.session_id | The ID of the user session associated with this event. Each login attempt results in a unique session id. | keyword |
| kibana.space_id | The id of the space associated with this event. | keyword |


### Log

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Event timestamp. | date |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| http.request.method | HTTP request method. Prior to ECS 1.6.0 the following guidance was provided: "The field value must be normalized to lowercase for querying." As of ECS 1.6.0, the guidance is deprecated because the original case of the method may be useful in anomaly detection.  Original case will be mandated in ECS 2.0.0 | keyword |
| http.request.referrer | Referrer for this HTTP request. | keyword |
| http.response.status_code | HTTP response status code. | long |
| kibana.add_to_spaces | The set of space ids that a saved object was shared to. | keyword |
| kibana.authentication_provider | The authentication provider associated with a login event. | keyword |
| kibana.authentication_realm | The Elasticsearch authentication realm name which fulfilled a login event. | keyword |
| kibana.authentication_type | The authentication provider type associated with a login event. | keyword |
| kibana.delete_from_spaces | The set of space ids that a saved object was removed from. | keyword |
| kibana.log.meta |  | object |
| kibana.log.state | Current state of Kibana. | keyword |
| kibana.log.tags | Kibana logging tags. | keyword |
| kibana.lookup_realm | The Elasticsearch lookup realm which fulfilled a login event. | keyword |
| kibana.saved_object.id | The id of the saved object associated with this event. | keyword |
| kibana.saved_object.type | The type of the saved object associated with this event. | keyword |
| kibana.session_id | The ID of the user session associated with this event. Each login attempt results in a unique session id. | keyword |
| kibana.space_id | The id of the space associated with this event. | keyword |
| source.address | Some event source addresses are defined ambiguously. The event will sometimes list an IP, a domain or a unix socket.  You should always store the raw address in the `.address` field. Then it should be duplicated to `.ip` or `.domain`, depending on which one it is. | keyword |
| url.original | Unmodified original url as seen in the event source. Note that in network monitoring, the observed URL may be a full URL, whereas in access logs, the URL is often just represented as a path. This field is meant to represent the URL as it was observed, complete or not. | wildcard |
| user_agent.original | Unparsed user_agent string. | keyword |


## Metrics

### Stats

Stats data stream uses the stats endpoint of Kibana, which is available in 6.4 by default.

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Date/time when the event originated. This is the date/time extracted from the event, typically representing when the event was generated by the source. If the event source has no original timestamp, this value is typically populated by the first time the event was received by the pipeline. Required field for all events. | date |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| kibana.stats.concurrent_connections | Number of client connections made to the server. Note that browsers can send multiple simultaneous connections to request multiple server assets at once, and they can re-use established connections. | long |
| kibana.stats.host.name | Kibana instance hostname | keyword |
| kibana.stats.index | Name of Kibana's internal index | keyword |
| kibana.stats.kibana.status |  | keyword |
| kibana.stats.name | Kibana instance name | keyword |
| kibana.stats.os.distro |  | keyword |
| kibana.stats.os.distroRelease |  | keyword |
| kibana.stats.os.load.15m |  | half_float |
| kibana.stats.os.load.1m |  | half_float |
| kibana.stats.os.load.5m |  | half_float |
| kibana.stats.os.memory.free_in_bytes |  | long |
| kibana.stats.os.memory.total_in_bytes |  | long |
| kibana.stats.os.memory.used_in_bytes |  | long |
| kibana.stats.os.platform |  | keyword |
| kibana.stats.os.platformRelease |  | keyword |
| kibana.stats.process.event_loop_delay.ms | Event loop delay in milliseconds | scaled_float |
| kibana.stats.process.memory.heap.size_limit.bytes | Max. old space size allocated to Node.js process, in bytes | long |
| kibana.stats.process.memory.heap.total.bytes | Total heap allocated to process in bytes | long |
| kibana.stats.process.memory.heap.uptime.ms | Uptime of process in milliseconds | long |
| kibana.stats.process.memory.heap.used.bytes | Heap used by process in bytes | long |
| kibana.stats.process.memory.resident_set_size.bytes |  | long |
| kibana.stats.process.uptime.ms |  | long |
| kibana.stats.request.disconnects | Number of requests that were disconnected | long |
| kibana.stats.request.total | Total number of requests | long |
| kibana.stats.response_time.avg.ms | Average response time in milliseconds | long |
| kibana.stats.response_time.max.ms | Maximum response time in milliseconds | long |
| kibana.stats.snapshot | Whether the Kibana build is a snapshot build | boolean |
| kibana.stats.status | Kibana instance's health status | keyword |
| kibana.stats.usage.index |  | keyword |
| service.id | Unique identifier of the running service. If the service is comprised of many nodes, the `service.id` should be the same for all nodes. This id should uniquely identify the service. This makes it possible to correlate logs and metrics for one specific service, no matter which particular node emitted the event. Note that if you need to see the events from one specific host of the service, you should filter on that `host.name` or `host.id` instead. | keyword |
| service.version | Version of the service the data was collected from. This allows to look at a data set only for a specific version of a service. | keyword |

An example event for `stats` looks as following:

```json
{
    "@timestamp": "2021-08-11T09:37:48.038Z",
    "agent": {
        "hostname": "docker-fleet-agent",
        "name": "docker-fleet-agent",
        "id": "09cdd3e1-f67a-4aca-bd69-ab2a5127490c",
        "ephemeral_id": "c73a88e9-ff0d-4bc0-8454-a4eace232146",
        "type": "metricbeat",
        "version": "7.15.0"
    },
    "process": {
        "pid": 1218
    },
    "elastic_agent": {
        "id": "09cdd3e1-f67a-4aca-bd69-ab2a5127490c",
        "version": "7.15.0",
        "snapshot": true
    },
    "ecs": {
        "version": "1.10.0"
    },
    "elasticsearch": {
        "cluster": {
            "id": "hEwxs-BJRuWNwJOV__gppg"
        }
    },
    "service": {
        "address": "http://kibana:5601/api/stats?extended=true",
        "name": "kibana",
        "id": "e7e31ce0-d42c-4829-8465-baf52f0b8334",
        "type": "kibana",
        "version": "7.15.0"
    },
    "data_stream": {
        "namespace": "default",
        "type": "metrics",
        "dataset": "kibana.stats"
    },
    "metricset": {
        "period": 10000,
        "name": "stats"
    },
    "event": {
        "duration": 16850171,
        "agent_id_status": "verified",
        "ingested": "2021-08-11T09:37:51.538117032Z",
        "module": "kibana",
        "dataset": "kibana.stats"
    },
    "kibana": {
        "stats": {
            "request": {
                "total": 4,
                "disconnects": 0
            },
            "process": {
                "memory": {
                    "heap": {
                        "total": {
                            "bytes": 296554496
                        },
                        "used": {
                            "bytes": 228129512
                        },
                        "size_limit": {
                            "bytes": 4345298944
                        }
                    }
                },
                "event_loop_delay": {
                    "ms": 0.56603500014171
                },
                "uptime": {
                    "ms": 1088184
                }
            },
            "host": {
                "name": "0.0.0.0"
            },
            "name": "kibana",
            "index": "kibana",
            "response_time": {
                "avg": {
                    "ms": 11
                },
                "max": {
                    "ms": 16
                }
            },
            "concurrent_connections": 8,
            "snapshot": true,
            "status": "green"
        }
    }
}
```

### Status

This status endpoint is available in 6.0 by default and can be enabled in Kibana >= 5.4 with the config option `status.v6ApiFormat: true`.

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Date/time when the event originated. This is the date/time extracted from the event, typically representing when the event was generated by the source. If the event source has no original timestamp, this value is typically populated by the first time the event was received by the pipeline. Required field for all events. | date |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| kibana.status.metrics.concurrent_connections | Current concurrent connections. | long |
| kibana.status.metrics.requests.disconnects | Total number of disconnected connections. | long |
| kibana.status.metrics.requests.total | Total number of connections. | long |
| kibana.status.name | Kibana instance name. | keyword |
| kibana.status.status.overall.state | Kibana overall state. | keyword |
| service.id | Unique identifier of the running service. If the service is comprised of many nodes, the `service.id` should be the same for all nodes. This id should uniquely identify the service. This makes it possible to correlate logs and metrics for one specific service, no matter which particular node emitted the event. Note that if you need to see the events from one specific host of the service, you should filter on that `host.name` or `host.id` instead. | keyword |
| service.version | Version of the service the data was collected from. This allows to look at a data set only for a specific version of a service. | keyword |

An example event for `status` looks as following:

```json
{
    "agent": {
        "hostname": "docker-fleet-agent",
        "name": "docker-fleet-agent",
        "id": "09cdd3e1-f67a-4aca-bd69-ab2a5127490c",
        "type": "metricbeat",
        "ephemeral_id": "09e64d5e-02f5-4ab0-859d-080e0aa1a4bb",
        "version": "7.15.0"
    },
    "elastic_agent": {
        "id": "09cdd3e1-f67a-4aca-bd69-ab2a5127490c",
        "version": "7.15.0",
        "snapshot": true
    },
    "@timestamp": "2021-08-11T09:39:06.207Z",
    "ecs": {
        "version": "1.10.0"
    },
    "service": {
        "address": "http://kibana:5601/api/status",
        "name": "kibana",
        "id": "e7e31ce0-d42c-4829-8465-baf52f0b8334",
        "type": "kibana",
        "version": "7.15.0"
    },
    "data_stream": {
        "namespace": "default",
        "type": "metrics",
        "dataset": "kibana.status"
    },
    "metricset": {
        "period": 10000,
        "name": "status"
    },
    "event": {
        "duration": 8391247,
        "agent_id_status": "verified",
        "ingested": "2021-08-11T09:39:09.730373425Z",
        "module": "kibana",
        "dataset": "kibana.status"
    },
    "kibana": {
        "status": {
            "name": "kibana",
            "metrics": {
                "requests": {
                    "total": 5,
                    "disconnects": 0
                },
                "concurrent_connections": 5
            },
            "status": {
                "overall": {
                    "state": "green"
                }
            }
        }
    }
}
```