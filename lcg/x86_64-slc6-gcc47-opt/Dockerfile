## a Dockerfile to install the whole LCG stack
FROM binet/slc
MAINTAINER binet@cern.ch

ENV HOME     /root
RUN export   HOME

ENV SITEROOT /opt/atlas-sw
RUN export SITEROOT

ENV HWAF_VARIANT x86_64-slc6-gcc47-opt
ENV HWAF_VERSION 20131126
ENV HWAF_ROOT    $SITEROOT/hwaf/hwaf-$HWAF_VERSION/linux-amd64
ENV PATH         $HWAF_ROOT/bin:$PATH

RUN export HWAF_VARIANT
RUN export HWAF_VERSION
RUN export HWAF_ROOT
RUN export PATH

## system deps
RUN yum install -y binutils-devel bzip2-devel bzip2 file git gcc gcc-c++ libXpm-devel libXft-devel libXext-devel m4 make ncurses-devel patch tar texinfo

ADD ../../install-hwaf.sh    /root/install-hwaf.sh
ADD build-lcg.sh             /root/build-lcg.sh
RUN /root/build-lcg.sh
RUN /bin/rm /root/build-lcg.sh

## EOF

