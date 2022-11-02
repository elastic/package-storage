# Kibana

The Kibana integration collects events from your [Kibana](https://www.elastic.co/guide/en/kibana/current/introduction.html) instance.

## Configuration parameters

If the Kibana instance is using a basepath in its URL, you must set the `basepath` setting for this integration with the same value.

## Compatibility

The `kibana` package works with Kibana 8.5.0 and later.

## Usage for Stack Monitoring

The `kibana` package can be used to collect metrics shown in our Stack Monitoring
UI in Kibana. To enable this usage, set `xpack.enabled: true` on the package config.

## Logs

### Audit

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Event timestamp. | date |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| ecs.version | ECS version this event conforms to. `ecs.version` is a required field and must exist in all events. When querying across multiple indices -- which may conform to slightly different ECS versions -- this field lets integrations adjust to the schema version of the events. | keyword |
| event.action | The action captured by the event. This describes the information in the event. It is more specific than `event.category`. Examples are `group-add`, `process-started`, `file-created`. The value is normally defined by the implementer. | keyword |
| event.category | This is one of four ECS Categorization Fields, and indicates the second level in the ECS category hierarchy. `event.category` represents the "big buckets" of ECS categories. For example, filtering on `event.category:process` yields all events relating to process activity. This field is closely related to `event.type`, which is used as a subcategory. This field is an array. This will allow proper categorization of some events that fall in multiple categories. | keyword |
| event.dataset | Name of the dataset. If an event source publishes more than one type of log or events (e.g. access log, error log), the dataset is used to specify which one the event comes from. It's recommended but not required to start the dataset name with the module name, followed by a dot, then the dataset name. | keyword |
| event.ingested | Timestamp when an event arrived in the central data store. This is different from `@timestamp`, which is when the event originally occurred.  It's also different from `event.created`, which is meant to capture the first time an agent saw the event. In normal conditions, assuming no tampering, the timestamps should chronologically look like this: `@timestamp` \< `event.created` \< `event.ingested`. | date |
| event.kind | This is one of four ECS Categorization Fields, and indicates the highest level in the ECS category hierarchy. `event.kind` gives high-level information about what type of information the event contains, without being specific to the contents of the event. For example, values of this field distinguish alert events from metric events. The value of this field can be used to inform how these kinds of events should be handled. They may warrant different retention, different access control, it may also help understand whether the data coming in at a regular interval or not. | keyword |
| event.outcome | This is one of four ECS Categorization Fields, and indicates the lowest level in the ECS category hierarchy. `event.outcome` simply denotes whether the event represents a success or a failure from the perspective of the entity that produced the event. Note that when a single transaction is described in multiple events, each event may populate different values of `event.outcome`, according to their perspective. Also note that in the case of a compound event (a single event that contains multiple logical events), this field should be populated with the value that best captures the overall success or failure from the perspective of the event producer. Further note that not all events will have an associated outcome. For example, this field is generally not populated for metric events, events with `event.type:info`, or any events for which an outcome does not make logical sense. | keyword |
| http.request.method | HTTP request method. The value should retain its casing from the original event. For example, `GET`, `get`, and `GeT` are all considered valid values for this field. | keyword |
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
| log.level | Original log level of the log event. If the source of the event provides a log level or textual severity, this is the one that goes in `log.level`. If your source doesn't specify one, you may put your event transport's severity here (e.g. Syslog severity). Some examples are `warn`, `err`, `i`, `informational`. | keyword |
| log.logger | The name of the logger inside an application. This is usually the name of the class which initialized the logger, or can be a custom name. | keyword |
| message | For log events the message field contains the log message, optimized for viewing in a log viewer. For structured logs without an original message field, other fields can be concatenated to form a human-readable summary of the event. If multiple messages exist, they can be combined into one message. | match_only_text |
| process.pid | Process id. | long |
| service.node.roles | Roles of a service node. This allows for distinction between different running roles of the same service. In the case of Kibana, the `service.node.role` could be `ui` or `background_tasks` or both. In the case of Elasticsearch, the `service.node.role` could be `master` or `data` or both. Other services could use this to distinguish between a `web` and `worker` role running as part of the service. | keyword |
| trace.id | Unique identifier of the trace. A trace groups multiple events like transactions that belong together. For example, a user request handled by multiple inter-connected services. | keyword |
| transaction.id | Unique identifier of the transaction within the scope of its trace. A transaction is the highest level of work measured within a service, such as a request to a server. | keyword |
| url.domain | Domain of the url, such as "www.elastic.co". In some cases a URL may refer to an IP and/or port directly, without a domain name. In this case, the IP address would go to the `domain` field. If the URL contains a literal IPv6 address enclosed by `[` and `]` (IETF RFC 2732), the `[` and `]` characters should also be captured in the `domain` field. | keyword |
| url.path | Path of the request, such as "/search". | wildcard |
| url.port | Port of the request, such as 443. | long |
| url.query | The query field describes the query string of the request, such as "q=elasticsearch". The `?` is excluded from the query string. If a URL contains no `?`, there is no query field. If there is a `?` but no query, the query field exists with an empty string. The `exists` query can be used to differentiate between the two cases. | keyword |
| url.scheme | Scheme of the request, such as "https". Note: The `:` is not part of the scheme. | keyword |
| user.name | Short name or login of the user. | keyword |
| user.name.text | Multi-field of `user.name`. | match_only_text |
| user.roles | Array of user roles at the time of the event. | keyword |


### Log

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Event timestamp. | date |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| ecs.version | ECS version this event conforms to. `ecs.version` is a required field and must exist in all events. When querying across multiple indices -- which may conform to slightly different ECS versions -- this field lets integrations adjust to the schema version of the events. | keyword |
| event.dataset | Name of the dataset. If an event source publishes more than one type of log or events (e.g. access log, error log), the dataset is used to specify which one the event comes from. It's recommended but not required to start the dataset name with the module name, followed by a dot, then the dataset name. | keyword |
| event.ingested | Timestamp when an event arrived in the central data store. This is different from `@timestamp`, which is when the event originally occurred.  It's also different from `event.created`, which is meant to capture the first time an agent saw the event. In normal conditions, assuming no tampering, the timestamps should chronologically look like this: `@timestamp` \< `event.created` \< `event.ingested`. | date |
| event.kind | This is one of four ECS Categorization Fields, and indicates the highest level in the ECS category hierarchy. `event.kind` gives high-level information about what type of information the event contains, without being specific to the contents of the event. For example, values of this field distinguish alert events from metric events. The value of this field can be used to inform how these kinds of events should be handled. They may warrant different retention, different access control, it may also help understand whether the data coming in at a regular interval or not. | keyword |
| event.outcome | This is one of four ECS Categorization Fields, and indicates the lowest level in the ECS category hierarchy. `event.outcome` simply denotes whether the event represents a success or a failure from the perspective of the entity that produced the event. Note that when a single transaction is described in multiple events, each event may populate different values of `event.outcome`, according to their perspective. Also note that in the case of a compound event (a single event that contains multiple logical events), this field should be populated with the value that best captures the overall success or failure from the perspective of the event producer. Further note that not all events will have an associated outcome. For example, this field is generally not populated for metric events, events with `event.type:info`, or any events for which an outcome does not make logical sense. | keyword |
| http.request.headers.accept |  | keyword |
| http.request.headers.authorization |  | keyword |
| http.request.headers.content-length |  | keyword |
| http.request.headers.content-type |  | keyword |
| http.request.headers.user-agent |  | keyword |
| http.request.headers.x-elastic-client-meta |  | keyword |
| http.request.headers.x-elastic-product-origin |  | keyword |
| http.request.headers.x-opaque-id |  | keyword |
| http.request.id | A unique identifier for each HTTP request to correlate logs between clients and servers in transactions. The id may be contained in a non-standard HTTP header, such as `X-Request-ID` or `X-Correlation-ID`. | keyword |
| http.request.method | HTTP request method. The value should retain its casing from the original event. For example, `GET`, `get`, and `GeT` are all considered valid values for this field. | keyword |
| http.response.body.bytes | Size in bytes of the response body. | long |
| http.response.headers.content-length |  | keyword |
| http.response.headers.content-type |  | keyword |
| http.response.headers.x-elastic-product |  | keyword |
| http.response.headers.x-opaque-id |  | keyword |
| http.response.status_code | HTTP response status code. | long |
| log.level | Original log level of the log event. If the source of the event provides a log level or textual severity, this is the one that goes in `log.level`. If your source doesn't specify one, you may put your event transport's severity here (e.g. Syslog severity). Some examples are `warn`, `err`, `i`, `informational`. | keyword |
| log.logger | The name of the logger inside an application. This is usually the name of the class which initialized the logger, or can be a custom name. | keyword |
| message | For log events the message field contains the log message, optimized for viewing in a log viewer. For structured logs without an original message field, other fields can be concatenated to form a human-readable summary of the event. If multiple messages exist, they can be combined into one message. | match_only_text |
| process.pid | Process id. | long |
| service.node.roles | Roles of a service node. This allows for distinction between different running roles of the same service. In the case of Kibana, the `service.node.role` could be `ui` or `background_tasks` or both. In the case of Elasticsearch, the `service.node.role` could be `master` or `data` or both. Other services could use this to distinguish between a `web` and `worker` role running as part of the service. | keyword |
| session_id |  | keyword |
| tags | List of keywords used to tag each event. | keyword |
| trace.id | Unique identifier of the trace. A trace groups multiple events like transactions that belong together. For example, a user request handled by multiple inter-connected services. | keyword |
| transaction.id | Unique identifier of the transaction within the scope of its trace. A transaction is the highest level of work measured within a service, such as a request to a server. | keyword |
| url.path | Path of the request, such as "/search". | wildcard |
| url.query | The query field describes the query string of the request, such as "q=elasticsearch". The `?` is excluded from the query string. If a URL contains no `?`, there is no query field. If there is a `?` but no query, the query field exists with an empty string. The `exists` query can be used to differentiate between the two cases. | keyword |


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
    "agent": {
        "name": "docker-fleet-agent",
        "id": "44d99b67-3ac6-44a7-aa72-63367a8c2f8b",
        "type": "metricbeat",
        "ephemeral_id": "ab3cdd2a-3336-4682-a038-6844197893f4",
        "version": "8.5.0"
    },
    "process": {
        "pid": 7
    },
    "@timestamp": "2022-08-06T22:34:12.983Z",
    "ecs": {
        "version": "8.0.0"
    },
    "data_stream": {
        "namespace": "default",
        "type": "metrics",
        "dataset": "kibana.stats"
    },
    "service": {
        "address": "https://kibana:5601/api/stats?extended=true",
        "id": "79307ef1-725a-4f29-992a-446bcbedf380",
        "type": "kibana",
        "version": "8.5.0"
    },
    "elastic_agent": {
        "id": "44d99b67-3ac6-44a7-aa72-63367a8c2f8b",
        "version": "8.5.0",
        "snapshot": true
    },
    "host": {
        "hostname": "docker-fleet-agent",
        "os": {
            "kernel": "5.10.47-linuxkit",
            "codename": "focal",
            "name": "Ubuntu",
            "type": "linux",
            "family": "debian",
            "version": "20.04.4 LTS (Focal Fossa)",
            "platform": "ubuntu"
        },
        "containerized": true,
        "ip": [
            "172.21.0.7"
        ],
        "name": "docker-fleet-agent",
        "mac": [
            "02:42:ac:15:00:07"
        ],
        "architecture": "x86_64"
    },
    "metricset": {
        "period": 10000,
        "name": "stats"
    },
    "event": {
        "duration": 22471757,
        "agent_id_status": "verified",
        "ingested": "2022-08-06T22:34:13Z",
        "module": "kibana",
        "dataset": "kibana.stats"
    },
    "kibana": {
        "elasticsearch": {
            "cluster": {
                "id": "wMZ6Mw1nR1ydMG25AiiOLw"
            }
        },
        "stats": {
            "request": {
                "total": 4,
                "disconnects": 0
            },
            "process": {
                "memory": {
                    "resident_set_size": {
                        "bytes": 510763008
                    },
                    "heap": {
                        "total": {
                            "bytes": 354033664
                        },
                        "used": {
                            "bytes": 280320136
                        },
                        "size_limit": {
                            "bytes": 4345298944
                        }
                    }
                },
                "event_loop_delay": {
                    "ms": 10.395972266666668
                },
                "uptime": {
                    "ms": 64365
                }
            },
            "os": {
                "distroRelease": "Ubuntu-20.04",
                "distro": "Ubuntu",
                "memory": {
                    "used_in_bytes": 4305055744,
                    "total_in_bytes": 35739144192,
                    "free_in_bytes": 31434088448
                },
                "load": {
                    "5m": 0.66,
                    "15m": 0.25,
                    "1m": 1.66
                },
                "platformRelease": "linux-5.10.47-linuxkit",
                "platform": "linux"
            },
            "name": "kibana",
            "host": {
                "name": "0.0.0.0"
            },
            "index": ".kibana",
            "response_time": {
                "avg": {
                    "ms": 8
                },
                "max": {
                    "ms": 11
                }
            },
            "concurrent_connections": 10,
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

### Cluster actions

Cluster actions metrics documentation

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Date/time when the event originated. This is the date/time extracted from the event, typically representing when the event was generated by the source. If the event source has no original timestamp, this value is typically populated by the first time the event was received by the pipeline. Required field for all events. | date |
| cluster_uuid |  | alias |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| ecs.version | ECS version this event conforms to. `ecs.version` is a required field and must exist in all events. When querying across multiple indices -- which may conform to slightly different ECS versions -- this field lets integrations adjust to the schema version of the events. | keyword |
| error.message | Error message. | match_only_text |
| event.dataset | Name of the dataset. If an event source publishes more than one type of log or events (e.g. access log, error log), the dataset is used to specify which one the event comes from. It's recommended but not required to start the dataset name with the module name, followed by a dot, then the dataset name. | keyword |
| event.duration | Duration of the event in nanoseconds. If event.start and event.end are known this value should be the difference between the end and start time. | long |
| event.module | Name of the module this data is coming from. If your monitoring agent supports the concept of modules or plugins to process events of a given source (e.g. Apache logs), `event.module` should contain the name of this module. | keyword |
| host.name | Name of the host. It can contain what `hostname` returns on Unix systems, the fully qualified domain name, or a name specified by the user. The sender decides which value to use. | keyword |
| kibana.cluster_actions.overdue.count |  | long |
| kibana.cluster_actions.overdue.delay.p50 |  | float |
| kibana.cluster_actions.overdue.delay.p99 |  | float |
| kibana.elasticsearch.cluster.id |  | keyword |
| kibana_stats.kibana.uuid |  | alias |
| kibana_stats.kibana.version |  | alias |
| kibana_stats.timestamp |  | alias |
| process.pid | Process id. | long |
| service.address | Address where data about this service was collected from. This should be a URI, network address (ipv4:port or [ipv6]:port) or a resource path (sockets). | keyword |
| service.id | Unique identifier of the running service. If the service is comprised of many nodes, the `service.id` should be the same for all nodes. This id should uniquely identify the service. This makes it possible to correlate logs and metrics for one specific service, no matter which particular node emitted the event. Note that if you need to see the events from one specific host of the service, you should filter on that `host.name` or `host.id` instead. | keyword |
| service.type | The type of the service data is collected from. The type can be used to group and correlate logs and metrics from one service type. Example: If logs or metrics are collected from Elasticsearch, `service.type` would be `elasticsearch`. | keyword |
| service.version | Version of the service the data was collected from. This allows to look at a data set only for a specific version of a service. | keyword |
| timestamp |  | alias |


An example event for `cluster_actions` looks as following:

```json
{
    "agent": {
        "name": "docker-fleet-agent",
        "id": "83c9f2b5-5134-4df2-88d8-ae48906024fc",
        "type": "metricbeat",
        "ephemeral_id": "f0c34fc3-ac35-4a80-80ed-a0de44ff6be0",
        "version": "8.5.0"
    },
    "service.id": "543c4fcf-bf38-4483-8cc4-df01fcb095e1",
    "elastic_agent": {
        "id": "83c9f2b5-5134-4df2-88d8-ae48906024fc",
        "version": "8.5.0",
        "snapshot": true
    },
    "@timestamp": "2022-08-06T21:38:59.780Z",
    "service.version": "8.5.0",
    "ecs": {
        "version": "8.0.0"
    },
    "service": {
        "address": "https://kibana:5601/api/monitoring_collection/cluster_actions",
        "type": "kibana"
    },
    "service.address": "0.0.0.0:5601",
    "data_stream": {
        "namespace": "default",
        "type": "metrics",
        "dataset": "kibana.cluster_actions"
    },
    "host": {
        "hostname": "docker-fleet-agent",
        "os": {
            "kernel": "5.10.47-linuxkit",
            "codename": "focal",
            "name": "Ubuntu",
            "type": "linux",
            "family": "debian",
            "version": "20.04.4 LTS (Focal Fossa)",
            "platform": "ubuntu"
        },
        "containerized": true,
        "ip": [
            "172.20.0.7"
        ],
        "name": "docker-fleet-agent",
        "mac": [
            "02:42:ac:14:00:07"
        ],
        "architecture": "x86_64"
    },
    "metricset": {
        "period": 10000,
        "name": "cluster_actions"
    },
    "event": {
        "duration": 13732239,
        "agent_id_status": "verified",
        "ingested": "2022-08-06T21:39:00Z",
        "module": "kibana",
        "dataset": "kibana.cluster_actions"
    },
    "kibana": {
        "elasticsearch.cluster.id": "Og-OqdQZQ62JHTfGBMc0CA",
        "cluster_actions": {
            "overdue": {
                "delay": {
                    "p99": 0,
                    "p50": 0
                },
                "count": 0
            }
        }
    }
}
```

### Cluster rules

Cluster rules metrics

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Date/time when the event originated. This is the date/time extracted from the event, typically representing when the event was generated by the source. If the event source has no original timestamp, this value is typically populated by the first time the event was received by the pipeline. Required field for all events. | date |
| cluster_uuid |  | alias |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| ecs.version | ECS version this event conforms to. `ecs.version` is a required field and must exist in all events. When querying across multiple indices -- which may conform to slightly different ECS versions -- this field lets integrations adjust to the schema version of the events. | keyword |
| error.message | Error message. | match_only_text |
| event.dataset | Name of the dataset. If an event source publishes more than one type of log or events (e.g. access log, error log), the dataset is used to specify which one the event comes from. It's recommended but not required to start the dataset name with the module name, followed by a dot, then the dataset name. | keyword |
| event.duration | Duration of the event in nanoseconds. If event.start and event.end are known this value should be the difference between the end and start time. | long |
| event.module | Name of the module this data is coming from. If your monitoring agent supports the concept of modules or plugins to process events of a given source (e.g. Apache logs), `event.module` should contain the name of this module. | keyword |
| host.name | Name of the host. It can contain what `hostname` returns on Unix systems, the fully qualified domain name, or a name specified by the user. The sender decides which value to use. | keyword |
| kibana.cluster_rules.overdue.count |  | long |
| kibana.cluster_rules.overdue.delay.p50 |  | float |
| kibana.cluster_rules.overdue.delay.p99 |  | float |
| kibana.elasticsearch.cluster.id |  | keyword |
| kibana_stats.kibana.uuid |  | alias |
| kibana_stats.kibana.version |  | alias |
| kibana_stats.timestamp |  | alias |
| process.pid | Process id. | long |
| service.address | Address where data about this service was collected from. This should be a URI, network address (ipv4:port or [ipv6]:port) or a resource path (sockets). | keyword |
| service.id | Unique identifier of the running service. If the service is comprised of many nodes, the `service.id` should be the same for all nodes. This id should uniquely identify the service. This makes it possible to correlate logs and metrics for one specific service, no matter which particular node emitted the event. Note that if you need to see the events from one specific host of the service, you should filter on that `host.name` or `host.id` instead. | keyword |
| service.type | The type of the service data is collected from. The type can be used to group and correlate logs and metrics from one service type. Example: If logs or metrics are collected from Elasticsearch, `service.type` would be `elasticsearch`. | keyword |
| service.version | Version of the service the data was collected from. This allows to look at a data set only for a specific version of a service. | keyword |
| timestamp |  | alias |


An example event for `cluster_rules` looks as following:

```json
{
    "agent": {
        "name": "docker-fleet-agent",
        "id": "83c9f2b5-5134-4df2-88d8-ae48906024fc",
        "ephemeral_id": "f0c34fc3-ac35-4a80-80ed-a0de44ff6be0",
        "type": "metricbeat",
        "version": "8.5.0"
    },
    "service.id": "543c4fcf-bf38-4483-8cc4-df01fcb095e1",
    "elastic_agent": {
        "id": "83c9f2b5-5134-4df2-88d8-ae48906024fc",
        "version": "8.5.0",
        "snapshot": true
    },
    "@timestamp": "2022-08-06T21:41:29.650Z",
    "service.version": "8.5.0",
    "ecs": {
        "version": "8.0.0"
    },
    "service": {
        "address": "https://kibana:5601/api/monitoring_collection/cluster_rules",
        "type": "kibana"
    },
    "data_stream": {
        "namespace": "default",
        "type": "metrics",
        "dataset": "kibana.cluster_rules"
    },
    "service.address": "0.0.0.0:5601",
    "host": {
        "hostname": "docker-fleet-agent",
        "os": {
            "kernel": "5.10.47-linuxkit",
            "codename": "focal",
            "name": "Ubuntu",
            "type": "linux",
            "family": "debian",
            "version": "20.04.4 LTS (Focal Fossa)",
            "platform": "ubuntu"
        },
        "containerized": true,
        "ip": [
            "172.20.0.7"
        ],
        "name": "docker-fleet-agent",
        "mac": [
            "02:42:ac:14:00:07"
        ],
        "architecture": "x86_64"
    },
    "metricset": {
        "period": 10000,
        "name": "cluster_rules"
    },
    "event": {
        "duration": 8419517,
        "agent_id_status": "verified",
        "ingested": "2022-08-06T21:41:30Z",
        "module": "kibana",
        "dataset": "kibana.cluster_rules"
    },
    "kibana": {
        "elasticsearch.cluster.id": "Og-OqdQZQ62JHTfGBMc0CA",
        "cluster_rules": {
            "overdue": {
                "delay": {
                    "p99": 0,
                    "p50": 0
                },
                "count": 0
            }
        }
    }
}
```

### Node actions

Node actions metrics

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Date/time when the event originated. This is the date/time extracted from the event, typically representing when the event was generated by the source. If the event source has no original timestamp, this value is typically populated by the first time the event was received by the pipeline. Required field for all events. | date |
| cluster_uuid |  | alias |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| ecs.version | ECS version this event conforms to. `ecs.version` is a required field and must exist in all events. When querying across multiple indices -- which may conform to slightly different ECS versions -- this field lets integrations adjust to the schema version of the events. | keyword |
| error.message | Error message. | match_only_text |
| event.dataset | Name of the dataset. If an event source publishes more than one type of log or events (e.g. access log, error log), the dataset is used to specify which one the event comes from. It's recommended but not required to start the dataset name with the module name, followed by a dot, then the dataset name. | keyword |
| event.duration | Duration of the event in nanoseconds. If event.start and event.end are known this value should be the difference between the end and start time. | long |
| event.module | Name of the module this data is coming from. If your monitoring agent supports the concept of modules or plugins to process events of a given source (e.g. Apache logs), `event.module` should contain the name of this module. | keyword |
| host.name | Name of the host. It can contain what `hostname` returns on Unix systems, the fully qualified domain name, or a name specified by the user. The sender decides which value to use. | keyword |
| kibana.elasticsearch.cluster.id |  | keyword |
| kibana.node_actions.executions |  | long |
| kibana.node_actions.failures |  | long |
| kibana.node_actions.timeouts |  | long |
| kibana_stats.kibana.uuid |  | alias |
| kibana_stats.kibana.version |  | alias |
| kibana_stats.timestamp |  | alias |
| process.pid | Process id. | long |
| service.address | Address where data about this service was collected from. This should be a URI, network address (ipv4:port or [ipv6]:port) or a resource path (sockets). | keyword |
| service.id | Unique identifier of the running service. If the service is comprised of many nodes, the `service.id` should be the same for all nodes. This id should uniquely identify the service. This makes it possible to correlate logs and metrics for one specific service, no matter which particular node emitted the event. Note that if you need to see the events from one specific host of the service, you should filter on that `host.name` or `host.id` instead. | keyword |
| service.type | The type of the service data is collected from. The type can be used to group and correlate logs and metrics from one service type. Example: If logs or metrics are collected from Elasticsearch, `service.type` would be `elasticsearch`. | keyword |
| service.version | Version of the service the data was collected from. This allows to look at a data set only for a specific version of a service. | keyword |
| timestamp |  | alias |


An example event for `node_actions` looks as following:

```json
{
    "agent": {
        "name": "docker-fleet-agent",
        "id": "83c9f2b5-5134-4df2-88d8-ae48906024fc",
        "type": "metricbeat",
        "ephemeral_id": "f0c34fc3-ac35-4a80-80ed-a0de44ff6be0",
        "version": "8.5.0"
    },
    "service.id": "543c4fcf-bf38-4483-8cc4-df01fcb095e1",
    "elastic_agent": {
        "id": "83c9f2b5-5134-4df2-88d8-ae48906024fc",
        "version": "8.5.0",
        "snapshot": true
    },
    "@timestamp": "2022-08-06T21:42:19.560Z",
    "ecs": {
        "version": "8.0.0"
    },
    "service.version": "8.5.0",
    "service.address": "0.0.0.0:5601",
    "service": {
        "address": "https://kibana:5601/api/monitoring_collection/node_actions",
        "type": "kibana"
    },
    "data_stream": {
        "namespace": "default",
        "type": "metrics",
        "dataset": "kibana.node_actions"
    },
    "host": {
        "hostname": "docker-fleet-agent",
        "os": {
            "kernel": "5.10.47-linuxkit",
            "codename": "focal",
            "name": "Ubuntu",
            "family": "debian",
            "type": "linux",
            "version": "20.04.4 LTS (Focal Fossa)",
            "platform": "ubuntu"
        },
        "containerized": true,
        "ip": [
            "172.20.0.7"
        ],
        "name": "docker-fleet-agent",
        "mac": [
            "02:42:ac:14:00:07"
        ],
        "architecture": "x86_64"
    },
    "metricset": {
        "period": 10000,
        "name": "node_actions"
    },
    "event": {
        "duration": 6658572,
        "agent_id_status": "verified",
        "ingested": "2022-08-06T21:42:20Z",
        "module": "kibana",
        "dataset": "kibana.node_actions"
    },
    "kibana": {
        "elasticsearch.cluster.id": "Og-OqdQZQ62JHTfGBMc0CA",
        "node_actions": {
            "failures": 0,
            "executions": 0,
            "timeouts": 0
        }
    }
}
```

### Node rules

Node rules metrics

**Exported fields**

| Field | Description | Type |
|---|---|---|
| @timestamp | Date/time when the event originated. This is the date/time extracted from the event, typically representing when the event was generated by the source. If the event source has no original timestamp, this value is typically populated by the first time the event was received by the pipeline. Required field for all events. | date |
| cluster_uuid |  | alias |
| data_stream.dataset | Data stream dataset. | constant_keyword |
| data_stream.namespace | Data stream namespace. | constant_keyword |
| data_stream.type | Data stream type. | constant_keyword |
| ecs.version | ECS version this event conforms to. `ecs.version` is a required field and must exist in all events. When querying across multiple indices -- which may conform to slightly different ECS versions -- this field lets integrations adjust to the schema version of the events. | keyword |
| error.message | Error message. | match_only_text |
| event.dataset | Name of the dataset. If an event source publishes more than one type of log or events (e.g. access log, error log), the dataset is used to specify which one the event comes from. It's recommended but not required to start the dataset name with the module name, followed by a dot, then the dataset name. | keyword |
| event.duration | Duration of the event in nanoseconds. If event.start and event.end are known this value should be the difference between the end and start time. | long |
| event.module | Name of the module this data is coming from. If your monitoring agent supports the concept of modules or plugins to process events of a given source (e.g. Apache logs), `event.module` should contain the name of this module. | keyword |
| host.name | Name of the host. It can contain what `hostname` returns on Unix systems, the fully qualified domain name, or a name specified by the user. The sender decides which value to use. | keyword |
| kibana.elasticsearch.cluster.id |  | keyword |
| kibana.node_rules.executions |  | long |
| kibana.node_rules.failures |  | long |
| kibana.node_rules.timeouts |  | long |
| kibana_stats.kibana.uuid |  | alias |
| kibana_stats.kibana.version |  | alias |
| kibana_stats.timestamp |  | alias |
| process.pid | Process id. | long |
| service.address | Address where data about this service was collected from. This should be a URI, network address (ipv4:port or [ipv6]:port) or a resource path (sockets). | keyword |
| service.id | Unique identifier of the running service. If the service is comprised of many nodes, the `service.id` should be the same for all nodes. This id should uniquely identify the service. This makes it possible to correlate logs and metrics for one specific service, no matter which particular node emitted the event. Note that if you need to see the events from one specific host of the service, you should filter on that `host.name` or `host.id` instead. | keyword |
| service.type | The type of the service data is collected from. The type can be used to group and correlate logs and metrics from one service type. Example: If logs or metrics are collected from Elasticsearch, `service.type` would be `elasticsearch`. | keyword |
| service.version | Version of the service the data was collected from. This allows to look at a data set only for a specific version of a service. | keyword |
| timestamp |  | alias |


An example event for `node_rules` looks as following:

```json
{
    "agent": {
        "name": "docker-fleet-agent",
        "id": "83c9f2b5-5134-4df2-88d8-ae48906024fc",
        "ephemeral_id": "f0c34fc3-ac35-4a80-80ed-a0de44ff6be0",
        "type": "metricbeat",
        "version": "8.5.0"
    },
    "service.id": "543c4fcf-bf38-4483-8cc4-df01fcb095e1",
    "elastic_agent": {
        "id": "83c9f2b5-5134-4df2-88d8-ae48906024fc",
        "version": "8.5.0",
        "snapshot": true
    },
    "@timestamp": "2022-08-06T21:42:59.474Z",
    "service.version": "8.5.0",
    "ecs": {
        "version": "8.0.0"
    },
    "service.address": "0.0.0.0:5601",
    "data_stream": {
        "namespace": "default",
        "type": "metrics",
        "dataset": "kibana.node_rules"
    },
    "service": {
        "address": "https://kibana:5601/api/monitoring_collection/node_rules",
        "type": "kibana"
    },
    "host": {
        "hostname": "docker-fleet-agent",
        "os": {
            "kernel": "5.10.47-linuxkit",
            "codename": "focal",
            "name": "Ubuntu",
            "type": "linux",
            "family": "debian",
            "version": "20.04.4 LTS (Focal Fossa)",
            "platform": "ubuntu"
        },
        "containerized": true,
        "ip": [
            "172.20.0.7"
        ],
        "name": "docker-fleet-agent",
        "mac": [
            "02:42:ac:14:00:07"
        ],
        "architecture": "x86_64"
    },
    "metricset": {
        "period": 10000,
        "name": "node_rules"
    },
    "kibana": {
        "elasticsearch.cluster.id": "Og-OqdQZQ62JHTfGBMc0CA",
        "node_rules": {
            "failures": 0,
            "executions": 0,
            "timeouts": 0
        }
    },
    "event": {
        "duration": 9031470,
        "agent_id_status": "verified",
        "ingested": "2022-08-06T21:43:00Z",
        "module": "kibana",
        "dataset": "kibana.node_rules"
    }
}
```