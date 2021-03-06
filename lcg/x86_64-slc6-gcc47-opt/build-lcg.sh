#!/bin/sh

TOPDIR=$1;       shift
HWAF_VERSION=$1; shift
HWAF_VARIANT=$1; shift
SITEROOT=$1;     shift
LCG_BRANCH=$1;   shift
WORCH_CONF=$1;   shift

HWAF_ROOT=$SITEROOT/hwaf/hwaf-$HWAF_VERSION/linux-amd64

PATH=$HWAF_ROOT/bin:$PATH

export SITEROOT
export HWAF_VARIANT
export HWAF_VERSION
export HWAF_ROOT
export PATH

export MODULEPATH=$SITEROOT/sw/modules:$MODULEPATH

echo "=========================================================================="
echo "SITEROOT:     [$SITEROOT]"
echo "HWAF_VARIANT: [$HWAF_VARIANT]"
echo "HWAF_VERSION: [$HWAF_VERSION]"
echo "HWAF_ROOT:    [$HWAF_ROOT]"
echo "PATH:         [$PATH]"
echo "=========================================================================="

set -e
set -x

### ---------------------------------------------------------------------------
echo "::: install base RPMs..."
yum install -y \
    autoconf automake \
    binutils-devel bzip2-devel bzip2 \
    environment-modules \
    file \
    gcc gcc-c++ git glibc-devel glibc-static \
    libtool libXpm-devel libXft-devel libXext-devel \
    m4 make \
    ncurses-devel \
    patch \
    readline readline-devel \
    tar texinfo

## setup env-modules
echo '
# enable environment-modules
if [ -f "/etc/profile.d/modules.sh" ] ; then
  source /etc/profile.d/modules.sh
fi
' >> /root/.bash_profile

### ---------------------------------------------------------------------------
echo "::: install hwaf-${HWAF_VERSION}... ($HWAF_ROOT)"
mkdir -p $HWAF_ROOT
curl -L http://cern.ch/hwaf/downloads/tar/hwaf-$HWAF_VERSION-linux-amd64.tar.gz | \
    tar -C $HWAF_ROOT -zxf -

echo "::: install hwaf-${HWAF_VERSION}... ($HWAF_ROOT) [ok]"

### ---------------------------------------------------------------------------
echo "::: build lcg stack..."

/bin/rm -rf $TOPDIR/scratch
/bin/mkdir -p $TOPDIR/scratch
pushd $TOPDIR/scratch

/bin/rm -rf lcg-builders
git clone -b $LCG_BRANCH git://github.com/atlas-org/lcg-builders
pushd lcg-builders
hwaf init
hwaf setup -variant=$HWAF_VARIANT
hwaf configure --prefix=$SITEROOT --hwaf-worch-conf=$WORCH_CONF
hwaf
popd # lcg-builders

popd # $TOPDIR/scratch
echo "::: build lcg stack... [ok]"

### ----
echo "::: cleaning up filesystem..."
/bin/rm -rf $TOPDIR/scratch/lcg-builders
echo "::: cleaning up filesystem... [ok]"

## EOF ##
