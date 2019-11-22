#!/bin/bash

nowDate=$(date +"%Y_%m_%d")
nowTime=$(date +"%H_%M_%S")
commitHash=$(git rev-parse --short HEAD)
versionString="$nowDate.$nowTime.0.$commitHash"
echo $versionString

versionOld=$(grep "const LomoWebVersion" main.go)
echo "old verion: $versionOld"
sed -i ".bak" -E "s/[[:digit:]]{4}_[[:digit:]]{2}_[[:digit:]]{2}\.[[:digit:]]{2}_[[:digit:]]{2}_[[:digit:]]{2}\.0\.[a-zA-Z0-9]{7}/$versionString/g" main.go
versionNew=$(grep "const LomoWebVersion" main.go)
echo "new verion: $versionNew"

rice embed-go

if [ "$(uname)" == "Darwin" ]; then
    CGO_CFLAGS=-mmacosx-version-min=10.10 CGO_LDFLAGS=-mmacosx-version-min=10.10 go build -o lomo-web
    rice append --exec lomo-web
    zip -r lomoWebOSX.zip lomo-web
    shasum -a256 lomoWebOSX.zip
else
    go build -o lomo-web
fi


