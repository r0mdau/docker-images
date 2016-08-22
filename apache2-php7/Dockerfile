FROM debian:stretch

RUN set -x \
    && apt-get update \
    && apt-get install -y \
        apache2 \
        php7.0 \
        php7.0-bcmath \
        php7.0-curl \
        php7.0-gd \
        php7.0-ldap \
        php7.0-mcrypt \
        php7.0-mysql \
        php7.0-sqlite3 \
        php-pear \
    --no-install-recommends \
    && apt-get install -y php7.0-dev \
    && rm -rf /var/lib/apt/lists/* \
    && pecl channel-update pecl.php.net \
    && printf "\n" | pecl install apcu \
    && echo "extension=apcu.so" >> /etc/php/7.0/apache2/php.ini \
    && echo "ServerName localhost" >> /etc/apache2/apache2.conf \
    && a2enmod headers && a2enmod rewrite

RUN ln -sf /dev/stdout /var/log/apache2/access.log \
    && ln -sf /dev/stderr /var/log/apache2/error.log

CMD ["apache2ctl",  "-D", "FOREGROUND"]
