# drone-opsworks

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-opsworks/status.svg)](http://beta.drone.io/drone-plugins/drone-opsworks)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-opsworks/coverage.svg)](https://aircover.co/drone-plugins/drone-opsworks)
[![](https://badge.imagelayers.io/plugins/drone-opsworks:latest.svg)](https://imagelayers.io/?images=plugins/drone-opsworks:latest 'Get your own badge on imagelayers.io')

Drone plugin to deploy or update a project on AWS OpsWorks

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-opsworks <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "access_key": "970d28f4dd477bc184fbd10b376de753",
        "secret_key": "9c5785d3ece6a9cdefa42eb99b58986f9095ff1c",
        "region": "us-east-1",
        "stack_id": "my-stack",
        "app_id": "my-app",
        "command": "deploy",
        "arguments": {
            "arg_name1": [
                "value1",
                "value2"
            ],
            "arg_name2": [
                "value1",
                "value2"
            ]
        },
        "instances": [
            "instance1",
            "instance2",
            "instance3"
        ]
    }
}
EOF
```

## Docker

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-opsworks <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "access_key": "970d28f4dd477bc184fbd10b376de753",
        "secret_key": "9c5785d3ece6a9cdefa42eb99b58986f9095ff1c",
        "region": "us-east-1",
        "stack_id": "my-stack",
        "app_id": "my-app",
        "command": "deploy",
        "arguments": {
            "arg_name1": [
                "value1",
                "value2"
            ],
            "arg_name2": [
                "value1",
                "value2"
            ]
        },
        "instances": [
            "instance1",
            "instance2",
            "instance3"
        ]
    }
}
EOF
```
