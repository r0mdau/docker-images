FROM debian:jessie
MAINTAINER romain.dauby@orange.com

RUN apt-get update
RUN apt-get install -y apache2
RUN echo "ServerName localhost" >> /etc/apache2/apache2.conf
RUN a2enmod headers && a2enmod rewrite
RUN apt-get install -y \
php5 \
php5-dev \
php5-mysql \
php5-curl \
php5-ldap \
php5-gd \
php5-mcrypt \
php5-memcache \
php-apc
RUN /etc/init.d/apache2 restart
CMD ["apache2ctl",  "-D", "FOREGROUND"]
