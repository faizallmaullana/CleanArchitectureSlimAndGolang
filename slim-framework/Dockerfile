FROM php:8.3-fpm-alpine

RUN apk add --no-cache --virtual .build-deps $PHPIZE_DEPS openssl-dev \
	&& pecl install mongodb \
	&& docker-php-ext-enable mongodb \
	&& apk del .build-deps

WORKDIR /var/www/html

CMD ["php-fpm", "-F"]