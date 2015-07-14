# Introduction

Dockerfile to build an iceweasel accessible by vnc in a docker conainter image

# Installation

Pull the latest version of the image from the docker index. This is the recommended method of installation as it is easier to update image in the future. These builds are performed by the **Docker Trusted Build** service.

```bash
docker pull r0mdau/iceweasel:latest
```

Alternately you can build the image yourself.

```bash
git clone https://github.com/r0mdau/docker-images.git
cd docker-images/iceweasel-vnc/
docker build -t r0mdau/iceweasel .
```
# Quick Start

```bash
docker run --name iceweasel -p 5900:5900 r0mdau/iceweasel x11vnc -forever -usepw -create
```

Now you can access to iceweasel through a vnc viewer on : localhost:5900

```bash
password : password
```
