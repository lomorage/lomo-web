#!/bin/bash
set -e

if [ "$#" -ne 1 ]; then
    echo "pack.sh [release-version]"
    exit 2
fi

PACKAGE_NAME="lomo-web"
VERSION=$1
ARCH=$(dpkg --print-architecture)
BUILD_NAME=$PACKAGE_NAME"_"$ARCH"_"$VERSION

if [ -d $BUILD_NAME ]; then
    rm -rf $BUILD_NAME
fi

mkdir $BUILD_NAME
mkdir $BUILD_NAME/DEBIAN

cat << EOF > $BUILD_NAME/DEBIAN/control
Package: $PACKAGE_NAME
Version: $VERSION
Section: net
Priority: optional
Architecture: $ARCH
Depends:
Maintainer: Jeromy Fu<fuji246@gmail.com>
Description: Lomorage Web App 
EOF

cat << EOF > $BUILD_NAME/DEBIAN/preinst
if [ -f "/lib/systemd/system/lomow.service" ]
then
  systemctl stop lomow || true
fi
EOF
chmod +x $BUILD_NAME/DEBIAN/preinst

cat << EOF > $BUILD_NAME/DEBIAN/postinst
#!/bin/bash
CUR_USER=${SUDO_USER:-$(logname)}
sudo sed -i "s/User=pi/User=$CUR_USER/g" /lib/systemd/system/lomow.service
chmod +x /opt/lomorage/bin/lomo-web
systemctl enable lomow
systemctl daemon-reload || true
systemctl restart lomow || true
EOF
chmod +x $BUILD_NAME/DEBIAN/postinst

mkdir -p $BUILD_NAME/lib/systemd/system
cp lomow.service $BUILD_NAME/lib/systemd/system/

mkdir -p $BUILD_NAME/opt/lomorage/bin
cp lomo-web $BUILD_NAME/opt/lomorage/bin

dpkg -b $BUILD_NAME
