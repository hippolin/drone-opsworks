# drone-aws-opsworks

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-aws-opsworks/status.svg)](http://beta.drone.io/drone-plugins/drone-aws-opsworks)
[![](https://badge.imagelayers.io/plugins/drone-aws-opsworks:latest.svg)](https://imagelayers.io/?images=plugins/drone-aws-opsworks:latest 'Get your own badge on imagelayers.io')

Drone plugin for deploying to AWS OpsWorks

## Usage

```sh
./drone-aws-opsworks <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
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
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```
make deps build docker
```

### Example

```sh
docker run -i plugins/drone-aws-opsworks <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
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
    }
}
EOF
```
