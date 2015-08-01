# Cluster

Manage clusters from the elora cli.


## Create a cluster

Create a cluster to be managed by elora.

`elora cluster create name -skip-consul -skip-weave -weave-pwd pwd -skip-cleanup`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the cluster to be identified in elora cli commands |
| skip-consul | 'true' to skip 'consult' installation for DNS in this cluster |
| skip-weave  | 'true' to skip 'weave' installation for SDN in this cluster |
| weave-pwd   | weave router password for traffic encryption. "" for non-encrypted traffic |
| skip-cleanup | 'true' to skip 'cleanup' service installation (cleans unused volumes and images) |


## Inspect a cluster

Retrieve cluster information.

`elora cluster inspect name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the cluster to be retrieved |


## List clusters

Retrieve all clusters information.

`elora cluster list`


## Delete a cluster

Remove a cluster from elora. Nodes created in this cluster are also deleted.

`elora cluster remove name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the cluster to be deleted |
