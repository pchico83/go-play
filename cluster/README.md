# Cluster

Manage clusters from the d2k8 cli.


## Create a cluster

Create a cluster to be managed by d2k8.

`d2k8 cluster create name -skip-consul -skip-weave -weave-pwd pwd -skip-cleanup`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the cluster to be identified in d2k8 cli commands |
| skip-consul | 'true' to skip 'consult' installation for DNS in this cluster |
| skip-weave  | 'true' to skip 'weave' installation for SDN in this cluster |
| weave-pwd   | weave router password for traffic encryption. "" for non-encrypted traffic |
| skip-cleanup | 'true' to skip 'cleanup' service installation (cleans unused volumes and images) |


## Inspect a cluster

Retrieve cluster information.

`d2k8 cluster inspect name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the cluster to be retrieved |


## List clusters

Retrieve all clusters information.

`d2k8 cluster list`


## Delete a cluster

Remove a cluster from d2k8. Nodes created in this cluster are also deleted.

`d2k8 cluster remove name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the cluster to be deleted |
