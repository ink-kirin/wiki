ARG PHP_VERSION
FROM php:${PHP_VERSION}-fpm-alpine
LABEL maintainer="simba"

ARG ALPINE_REPOSITORIES
ARG PHP_REDIS_VERSION
ARG PHP_SWOOLE_VERSION

RUN ln -snf /use/share/zoneinfo/$TZ /etc/localtime \
    && echo $TZ > /etc/timezone \
    && if [ "${ALPINE_REPOSITORIES}" != "" ]; then \
        sed -i "s/dl-cdn.alpinelinux.org/${ALPINE_REPOSITORIES}/g" /etc/apk/repositories; \
    fi \
    && apk add --no-cache autoconf g++ libtool make curl-dev libxml2-dev linux-headers \
    freetype-dev libjpeg-turbo-dev libpng-dev bzip2-dev libzip-dev \
    && docker-php-ext-configure gd --with-freetype-dir=/usr/include/ --with-jpeg-dir=/usr/include/ \
    && mkdir -p /usr/src/php/ext \
    && wget https://github.com/phpredis/phpredis/archive/${PHP_REDIS_VERSION}.tar.gz \
    && tar -zxvf ${PHP_REDIS_VERSION}.tar.gz \
    && mv phpredis-${PHP_REDIS_VERSION} /usr/src/php/ext/redis \
    && wget https://github.com/swoole/swoole-src/archive/v${PHP_SWOOLE_VERSION}.tar.gz \
    && tar -zxvf v${PHP_SWOOLE_VERSION}.tar.gz \
    && mv swoole-src-${PHP_SWOOLE_VERSION} /usr/src/php/ext/swoole \
    && docker-php-ext-install -j$(nproc) pdo_mysql zip mysqli mbstring bcmath \
    opcache sockets gd bz2 curl redis swoole exif

WORKDIR /var/www/html
