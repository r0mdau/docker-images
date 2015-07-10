## Build

docker build -t r0mdau/nginx:latest .

## Host

Application host : *.dev (http://something.dev:8080)

## Run

docker run --name nginx-static -p 8080:80 -v /home/romain/project/public:/srv/http/:ro -d r0mdau/nginx:latest
