FROM registry.redhat.io/ubi8:8.2-347

LABEL author ragaws1674@gmail.com
LABEL version v1
LABEL using supervisord

# Install necessary packages

COPY ./terraform.repo /etc/yum.repos.d/terraform.repo

RUN yum update -y && yum install -y \ 
net-tools \
    terraform \
    python3

RUN python3 -m pip install supervisor

COPY ./supervisord.conf /etc/supervisord.conf

RUN chmod -R 0640 /etc/supervisord.conf

# run supervisord
CMD ["/usr/local/bin/supervisord","-c","/etc/supervisord.conf"]
