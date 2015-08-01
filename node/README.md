# Node

Manage hosts from the elora cli.


## Create a node

Create a node to be managed by elora and installs docker on it.

`elora node create name -ip ip -cert docker_cert_path -tags staging,database`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be identified in elora cli commands |
| url         | url to connect to the node |
| cert        | path to docker node certificates |
| tags        | list of tags associated to this node |


## Configure a docker client to connect to a node

Retrieve Docker environment variables to allow docker client configuration by executing `$(elora node env name)`.

`elora node env name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be retrieved |


## Inspect a node

Retrieve node information.

`elora node inspect name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be retrieved |


## List nodes

Retrive all nodes information.

`elora node list`


## Update a node

Update a node managed by elora.

`elora node update name -ip ip -cert docker_cert_path -tags staging,database`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be updated |
| url         | new url of the node (optional) |
| cert        | new path to docker node certificates (optional) |
| tags        | new list of tags associated to this node (optional) |


## Delete a node

Remove a node from elora. Containers created in this node are also deleted.

`elora node remove name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be deleted |


# Integration with Docker-Machine

[Docker-Machine](https://github.com/docker/machine) is a beuatiful Docker project to create virtual machines where docker 
is automatically configured.
At elora we don't want to re-implement all this functionality, but allow for easier integration of Docker-Machine users instead.
This is achieved by allowing `elora node create` flags to be passed via environment variables.
In particular, the value of `-url` is taken from `DOCKER_HOST` and the value of `-cert` is taken from `DOCKER_CERT_PATH`
if they are not specified in the `node create` command.

For example, to create a virtualbox node and register it into elora you can do like this:

```
docker-machine create --driver virtualbox test
$(docker-machine env test)
elora node create test
```

And the node `test` will be ready to deploy stacks from elora.
A full list of Docker-Machine supported providers and its configuration options is available [here](https://docs.docker.com/machine/#drivers).
