A stack is an instance of a YAML d2k8 template

# Stack

Manage stacks from the d2k8 cli.


## Create a stack

Create a stack to be managed from d2k8.

`d2k8 stack create name path`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the stack to be identified in d2k8 cli commands |
| path        | path to the YAML d2k8 template |


## Inspect a stack

Retrive stack dynamic information.

`d2k8 stack inspect name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the stack to be retrieved |


## Export a stack

Retrieve the stack YAML d2k8 template.

`d2k8 stack export name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the stack to be retrieved |


## List Stacks

Retrive all stacks information.

`d2k8 stack list`

## Update a stack

Update a stack managed by d2k8.

`d2k8 stack update name path`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the host to be updated |
| path        | path to the new YAML d2k8 template |


## Delete a stack

Remove a stack from d2k8. Containers created for this stack are also deleted.

`d2k8 stack remove name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the stack to be deleted |
