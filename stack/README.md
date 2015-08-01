A stack is an instance of a YAML elora template

# Stack

Manage stacks from the elora cli.


## Create a stack

Create a stack to be managed from elora.

`elora stack create name path`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the stack to be identified in elora cli commands |
| path        | path to the YAML elora template |


## Inspect a stack

Retrive stack dynamic information.

`elora stack inspect name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the stack to be retrieved |


## Export a stack

Retrieve the stack YAML elora template.

`elora stack export name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the stack to be retrieved |


## List Stacks

Retrive all stacks information.

`elora stack list`

## Update a stack

Update a stack managed by elora.

`elora stack update name path`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the host to be updated |
| path        | path to the new YAML elora template |


## Delete a stack

Remove a stack from elora. Containers created for this stack are also deleted.

`elora stack remove name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the stack to be deleted |
