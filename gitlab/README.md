# Introduction

Dockerfile to build a [GitLab CI](https://about.gitlab.com/gitlab-ci/) container image.

# Installation

Pull the latest version of the image from the docker index. This is the recommended method of installation as it is easier to update image in the future. These builds are performed by the **Docker Trusted Build** service.

```bash
docker pull r0mdau/gitlab:latest
```

Alternately you can build the image yourself.

```bash
git clone https://github.com/r0mdau/docker-images.git
cd docker-images/gitlab/
docker build -t r0mdau/gitlab .
```
# Quick Start

```bash
docker run --detach --publish 8080:80 --publish 2222:22 r0mdau/gitlab
```

Now you can access to gitlab through your web browser : http://localhost:8080

```bash
user : root
password : 5iveL!fe
```
