## a Dockerfile to install the whole LCG stack
FROM binet/slc
MAINTAINER binet@cern.ch

ENV HOME     /root
RUN export   HOME

ENV SITEROOT /opt/atlas-sw
RUN export SITEROOT

ENV HWAF_VARIANT x86_64-slc6-gcc47-opt
ENV HWAF_VERSION 20131204
ENV HWAF_ROOT    $SITEROOT/hwaf/hwaf-$HWAF_VERSION/linux-amd64
ENV PATH         $HWAF_ROOT/bin:$PATH

RUN export HWAF_VARIANT
RUN export HWAF_VERSION
RUN export HWAF_ROOT
RUN export PATH

VOLUME ["/build"]
ADD build-lcg.sh /build/build-lcg.sh

RUN /build/build-lcg.sh /build $HWAF_VERSION $HWAF_VARIANT $SITEROOT build-lcg-base.cfg

ENV MODULEPATH $SITEROOT/sw/modules/$HWAF_VARIANT:$MODULEPATH
RUN export MODULEPATH
RUN . /etc/profile.d/modules.sh
#RUN module load lcg

## EOF

