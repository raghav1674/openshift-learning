FROM alpine:latest

LABEL author ragaws1674@gmail.com
LABEL version v1
LABEL using supervisord

# Install necessary packages
RUN apk update && apk add --no-cache supervisor \ 
openssl \
    curl \
    gnupg \
    wget \ 
terraform \
    busybox-extras \
    python3

RUN adduser -D sysadmin -u 1000

# Create config files for supervisord
RUN mkdir -p /data/supervisor/log && mkdir -p /data/supervisor/run && mkdir -p /data/supervisor/conf
COPY ./supervisord.conf /data/supervisor/conf/supervisord.conf

RUN chown -R 1000:1000 /data/supervisor

# run supervisord
CMD ["/usr/bin/supervisord","-c","/data/supervisor/conf/supervisord.conf"]
