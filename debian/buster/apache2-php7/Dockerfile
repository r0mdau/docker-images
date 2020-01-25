FROM debian:buster-slim
LABEL name="r0mdau/buster-apache2-php7"

RUN set -x && \
    apt-get update && \
    apt-get install --no-install-recommends --no-install-suggests -y \
        ca-certificates \
        apache2 \
        libapache2-mod-php7.3 \
        php7.3 \
        php-apcu \
        php7.3-bcmath \
        php7.3-cli \
        php7.3-common \
        php7.3-curl \
        php7.3-json \
        php7.3-ldap \
        php7.3-mbstring \
        php7.3-mysql \
        php7.3-opcache \
        php7.3-readline \
        php7.3-xml && \
    echo 'ServerName localhost' >> /etc/apache2/apache2.conf && \
    echo 'date.timezone = "Europe/Paris"' >> /etc/php/7.3/apache2/php.ini && \
    a2enmod headers && \
    a2enmod rewrite && \
    a2enmod ssl && \
    rm -rf /var/lib/apt/lists/*

RUN ln -sf /dev/stdout /var/log/apache2/access.log \
    && ln -sf /dev/stderr /var/log/apache2/error.log

COPY apache2-foreground /usr/local/bin/

EXPOSE 80 443

ENTRYPOINT ["apache2-foreground"]