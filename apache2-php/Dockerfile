FROM debian:jessie

RUN apt-get update \
    && apt-get install --no-install-recommends --no-install-suggests -y \
        apache2 \
        php5 \
        php5-dev \
        php5-mysql \
        php5-curl \
        php5-ldap \
        php5-gd \
        php5-mcrypt \
        php5-memcache \
        php-apc \
    && echo "ServerName localhost" >> /etc/apache2/apache2.conf \
    && a2enmod headers && a2enmod rewrite \
    && rm -rf /var/lib/apt/lists/*

RUN ln -sf /dev/stdout /var/log/apache2/access.log \
    && ln -sf /dev/stderr /var/log/apache2/error.log

CMD ["apache2ctl",  "-D", "FOREGROUND"]
