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

# Create config files for supervisord
COPY ./supervisord.conf /etc/supervisord.conf

RUN chmod -R 0640 /etc/supervisord.conf

# run supervisord
CMD ["/usr/bin/supervisord"]
