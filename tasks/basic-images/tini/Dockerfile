FROM alpine:latest

LABEL author ragaws1674@gmail.com
LABEL version v1
LABEL using simplebash

# Install necessary packages
RUN apk update && apk add --no-cache \ 
openssl \
    curl \
    gnupg \
    wget \ 
terraform \
    busybox-extras \
    python3 \
    tini

ENTRYPOINT [ "/sbin/tini","--"]

CMD ["/bin/sh","-c","sleep infinity"]
