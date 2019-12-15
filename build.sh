#!/bin/bash

nowDate=$(date +"%Y-%m-%d")
nowTime=$(date +"%H-%M-%S")
commitHash=$(git rev-parse --short HEAD)
versionString="$nowDate.$nowTime.0.$commitHash"
echo $versionString

versionOld=$(grep "const LomoWebVersion" main.go)
echo "old verion: $versionOld"
sed -i.bak -E "s/[[:digit:]]{4}-[[:digit:]]{2}-[[:digit:]]{2}\.[[:digit:]]{2}-[[:digit:]]{2}-[[:digit:]]{2}\.0\.[a-zA-Z0-9]{7}/$versionString/g" main.go
versionNew=$(grep "const LomoWebVersion" main.go)
echo "new verion: $versionNew"

rice embed-go

if [ "$(uname)" == "Darwin" ]; then
    CGO_CFLAGS=-mmacosx-version-min=10.10 CGO_LDFLAGS=-mmacosx-version-min=10.10 go build -o lomo-web
    zip -r lomoWebOSX.zip lomo-web
    shasum -a256 lomoWebOSX.zip
elif [ "$(uname)" == "Linux" ]; then
    go build -o lomo-web
    sudo ./pack.sh $versionString
elif [ "$(expr substr $(uname -s) 1 5)" == "MINGW" ]; then
    go build -o lomo-web
    zip -r lomoWebWin.zip lomo-web
    certUtil -hashfile lomoWebWin.zip SHA256
else
    go build -o lomo-web
fi


