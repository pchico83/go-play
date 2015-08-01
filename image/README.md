# Image

Manage private images from the elora cli.


## Create an image

Create a private image to be managed by elora.

`elora image create name -user user -pwd pwd -url url`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the image to be identified in elora cli commands |
| user        | username to access the private image |
| pwd         | password to access the private image |
| url         | url to access the private image |


## Inspect a private image

Retrive image information.

`elora image inspect name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the image to be retrieved |


## List private images

Retrieve all private images information.

`elora image list`


## Update a private image

Update a private image managed by elora.

`elora image update name --user user -pwd pwd -url url`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be updated |
| user        | new username to access the private image (optional) |
| pwd         | new password to access the private image (optional) |
| url         | new url to access the private image (optional) |


## Delete a private image

Remove a private image from elora.

`elora image remove name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the image to be deleted |
