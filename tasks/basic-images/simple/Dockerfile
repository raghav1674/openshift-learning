FROM registry.redhat.io/ubi8:8.2-347

LABEL author ragaws1674@gmail.com
LABEL version v1
LABEL using simplebash

# Install necessary packages

COPY ./terraform.repo /etc/yum.repos.d/terraform.repo

RUN yum update -y && yum install -y \ 
    net-tools \
    terraform \
    python3

CMD ["sleep","infinity"]
