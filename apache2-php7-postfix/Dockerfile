FROM r0mdau/stretch-apache2-php7
LABEL name="r0mdau/stretch-apache2-php7-postfix"

ENV DEBIAN_FRONTEND noninteractive

RUN set -x && \
    apt-get update && \
    apt-get install --no-install-recommends --no-install-suggests -y postfix && \
    rm -rf /var/lib/apt/lists/* && \
    useradd -s /bin/bash email && \
    echo "default_transport = file_route" >> /etc/postfix/main.cf && \
    echo "file_route   unix  -       n       n       -       -       pipe" >> /etc/postfix/master.cf && \
    echo "  flags=F user=email argv=/usr/local/bin/mailToFile.php" >> /etc/postfix/master.cf

COPY mailToFile.php /usr/local/bin/
RUN chmod 755 /usr/local/bin/mailToFile.php

COPY launch-foreground /usr/local/bin/

ENTRYPOINT ["launch-foreground"]