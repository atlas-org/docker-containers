## Dockerfile to create a minimal SLC6 development environment
FROM binet/slc-base
MAINTAINER Sebastien Binet <binet@cern.ch>

## install a few minimal dev RPMs
RUN yum install -y file tar which
RUN yum clean metadata

ENV PYTHONSTARTUP /root/.pythonrc.py
RUN export PYTHONSTARTUP


## add files last (as this invalids caches)
ADD dot-pythonrc.py  /.pythonrc.py
ADD dot-bashrc       /.bashrc
ADD dot-bash_profile /.bash_profile

## EOF

