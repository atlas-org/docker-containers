#!/bin/sh

set -e

### ---------------------------------------------------------------------------
echo "::: install hwaf-${HWAF_VERSION}... ($HWAF_ROOT)"
mkdir -p $HWAF_ROOT
curl -L http://cern.ch/hwaf/downloads/tar/hwaf-$HWAF_VERSION-linux-amd64.tar.gz | \
    tar -C $HWAF_ROOT -zxf -

export PATH=$HWAF_ROOT/bin:$PATH
echo "::: install hwaf-${HWAF_VERSION}... ($HWAF_ROOT) [ok]"


### ---------------------------------------------------------------------------
echo "::: build lcg stack..."

mkdir -p /build
pushd /build

git clone git://github.com/atlas-org/lcg-builders
pushd lcg-builders
hwaf init
hwaf setup -variant=$HWAF_VARIANT
hwaf configure --prefix=$SITEROOT/sw/lcg/external
hwaf
popd # lcg-builders

popd # /build
echo "::: build lcg stack... [ok]"
## EOF ##
