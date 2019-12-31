FROM nginx:1.15

ADD ./docker/nginx/vhost.conf /etc/nginx/conf.d/default.conf

RUN apt-get update 

RUN apt-get install nano

WORKDIR /var/www