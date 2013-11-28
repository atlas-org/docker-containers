#!/bin/sh

set -e
set -x

HWAF_ROOT=$1;    shift
HWAF_VERSION=$1; shift

### ---------------------------------------------------------------------------
echo "::: install hwaf-${HWAF_VERSION}... ($HWAF_ROOT)"
mkdir -p $HWAF_ROOT
curl -L http://cern.ch/hwaf/downloads/tar/hwaf-$HWAF_VERSION-linux-amd64.tar.gz | \
    tar -C $HWAF_ROOT -zxf -

echo "::: install hwaf-${HWAF_VERSION}... ($HWAF_ROOT) [ok]"

## EOF

