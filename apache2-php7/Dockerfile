FROM debian:stretch
LABEL name="r0mdau/stretch-apache2-php7"

RUN set -x && \
    apt-get update && \
    apt-get install --no-install-recommends --no-install-suggests -y \
        ca-certificates \
        apache2 \
        libapache2-mod-php7.0 \
        libapache2-mod-shib2 \
        php7.0 \
        php7.0-apcu \
        php7.0-bcmath \
        php7.0-cli \
        php7.0-common \
        php7.0-curl \
        php7.0-dev \
        php7.0-gd \
        php7.0-json \
        php7.0-ldap \
        php7.0-mbstring \
        php7.0-mcrypt \
        php7.0-mysql \
        php7.0-opcache \
        php7.0-readline \
        php7.0-sqlite3 \
        php7.0-xdebug \
        php7.0-xml && \
    echo "TLS_REQCERT allow" >> /etc/ldap/ldap.conf && \
    echo "ServerName localhost" >> /etc/apache2/apache2.conf && \
    echo 'date.timezone = "Europe/Paris"' >> /etc/php/7.0/apache2/php.ini && \
    a2enmod headers && \
    a2enmod rewrite && \
    a2enmod ssl && \
    rm -rf /var/lib/apt/lists/*

RUN ln -sf /dev/stdout /var/log/apache2/access.log \
    && ln -sf /dev/stderr /var/log/apache2/error.log

COPY apache2-foreground /usr/local/bin/

EXPOSE 80 443

ENTRYPOINT ["apache2-foreground"]
