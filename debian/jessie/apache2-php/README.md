# Introduction

Dockerfile to build apache2 with PHP using Apt

## Installation
Pull the latest version of the image from the docker index. This is the recommended method of installation as it is easier to update image in the future. These builds are performed by the **Docker Trusted Build** service.

```bash
docker pull r0mdau/apache2-php:latest
```

## Build

```bash
git clone https://github.com/r0mdau/docker-images.git
cd docker-images/apache2-php/
docker build -t r0mdau/apache2-php .
```

## Run

```bash
docker run --name apache2-php -p 80:80 -v /pathToProject:/var/www/html:rw -v /pathToVhost:/etc/apache2/sites-enabled:ro -d r0mdau/apache2-php:latest
```

Path to vhost is not mandatory.

## To customize

Dockerfile example :

```
FROM r0mdau/apache2-php:latest
RUN apt-get update
RUN apt-get install -y YOURPACKAGE
RUN rm -rf /var/lib/apt/lists/*
```
