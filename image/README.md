# Image

Manage private images from the d2k8 cli.


## Create an image

Create a private image to be managed by d2k8.

`d2k8 image create name -user user -pwd pwd -url url`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the image to be identified in d2k8 cli commands |
| user        | username to access the private image |
| pwd         | password to access the private image |
| url         | url to access the private image |


## Inspect a private image

Retrive image information.

`d2k8 image inspect name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the image to be retrieved |


## List private images

Retrieve all private images information.

`d2k8 image list`


## Update a private image

Update a private image managed by d2k8.

`d2k8 image update name --user user -pwd pwd -url url`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the node to be updated |
| user        | new username to access the private image (optional) |
| pwd         | new password to access the private image (optional) |
| url         | new url to access the private image (optional) |


## Delete a private image

Remove a private image from d2k8.

`d2k8 image remove name`

| Argument    | Meaning       |
| ----------- | ------------- |
| name        | name of the image to be deleted |
