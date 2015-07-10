FROM nginx:latest
MAINTAINER  romain.dauby@gmail.com
RUN mkdir /srv/http && chown www-data:www-data /srv/http
RUN echo "server {listen   80; root /srv/http; index index.html; server_name *.dev;} " > /etc/nginx/conf.d/application.conf
