#!/bin/sh

set -e

### ---------------------------------------------------------------------------
./install-hwaf.sh $HWAF_ROOT $HWAF_VERSION

### ---------------------------------------------------------------------------
echo "::: build lcg stack..."

mkdir -p /tmp
pushd /tmp

/bin/rm -rf lcg-builders
git clone -b lcg-65-branch git://github.com/atlas-org/lcg-builders
pushd lcg-builders
hwaf init
hwaf setup -variant=$HWAF_VARIANT
hwaf configure --prefix=$SITEROOT
hwaf
popd # lcg-builders

popd # /tmp
echo "::: build lcg stack... [ok]"

### ----
echo "::: cleaning up filesystem..."
/bin/rm -rf /tmp/lcg-builders
echo "::: cleaning up filesystem... [ok]"

## EOF ##
