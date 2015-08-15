# Node

Manage hosts from the d2k8 cli.


## Create a node

Create a node to be managed by d2k8 and installs docker on it.

`d2k8 node create name -ip ip -cert docker_cert_path -tags staging,database`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be identified in elora cli commands |
| url         | url to connect to the node |
| cert        | path to docker node certificates |
| tags        | list of tags associated to this node |


## Configure a docker client to connect to a node

Retrieve Docker environment variables to allow docker client configuration by executing `$(d2k8 node env name)`.

`d2k8 node env name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be retrieved |


## Inspect a node

Retrieve node information.

`d2k8 node inspect name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be retrieved |


## List nodes

Retrive all nodes information.

`d2k8 node list`


## Update a node

Update a node managed by d2k8.

`d2k8 node update name -ip ip -cert docker_cert_path -tags staging,database`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be updated |
| url         | new url of the node (optional) |
| cert        | new path to docker node certificates (optional) |
| tags        | new list of tags associated to this node (optional) |


## Delete a node

Remove a node from d2k8. Containers created in this node are also deleted.

`d2k8 node remove name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be deleted |


# Integration with Docker-Machine

[Docker-Machine](https://github.com/docker/machine) is a beuatiful Docker project to create virtual machines where docker 
is automatically configured.
At d2k8 we don't want to re-implement all this functionality, but allow for easier integration of Docker-Machine users instead.
This is achieved by allowing `d2k8 node create` flags to be passed via environment variables.
In particular, the value of `-url` is taken from `DOCKER_HOST` and the value of `-cert` is taken from `DOCKER_CERT_PATH`
if they are not specified in the `node create` command.

For example, to create a virtualbox node and register it into d2k8 you can do like this:

```
docker-machine create --driver virtualbox test
$(docker-machine env test)
d2k8 node create test
```

And the node `test` will be ready to deploy stacks from d2k8.
A full list of Docker-Machine supported providers and its configuration options is available [here](https://docs.docker.com/machine/#drivers).
