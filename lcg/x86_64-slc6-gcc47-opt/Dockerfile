## a Dockerfile to install the whole LCG stack
FROM binet/slc6-64
MAINTAINER binet@cern.ch

ENV HOME     /root
ENV SITEROOT /opt/atlas-sw

ENV HWAF_VARIANT x86_64-slc6-gcc47-opt
ENV HWAF_VERSION 20131125
ENV HWAF_ROOT    $SITEROOT/hwaf/hwaf-$HWAF_VERSION/linux-amd64


ADD build-lcg.sh /root/build-lcg.sh
RUN /root/build-lcg.sh
RUN /bin/rm /root/build-lcg.sh

## EOF

