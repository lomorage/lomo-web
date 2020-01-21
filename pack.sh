#!/bin/bash
set -e

if [ "$#" -ne 1 ]; then
    echo "pack.sh [release-version]"
    exit 2
fi

PACKAGE_NAME="lomo-web"
VERSION=$1
BUILD_NAME=$PACKAGE_NAME"_"$VERSION

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
Architecture: all
Depends:
Maintainer: Jeromy Fu<fuji246@gmail.com>
Description: Lomorage Web App 
EOF

cat << EOF > $BUILD_NAME/DEBIAN/preinst
if [ -f "/lib/systemd/system/lomow.service" ]
then
  systemctl stop lomow
fi
EOF
chmod +x $BUILD_NAME/DEBIAN/preinst

cat << EOF > $BUILD_NAME/DEBIAN/postinst
#!/bin/bash
chmod +x /opt/lomorage/bin/lomo-web
systemctl enable lomow
systemctl daemon-reload 
systemctl restart lomow
EOF
chmod +x $BUILD_NAME/DEBIAN/postinst

mkdir -p $BUILD_NAME/lib/systemd/system
cp lomow.service $BUILD_NAME/lib/systemd/system/

mkdir -p $BUILD_NAME/opt/lomorage/bin
cp lomo-web $BUILD_NAME/opt/lomorage/bin

chown root:root -R $BUILD_NAME
dpkg -b $BUILD_NAME
