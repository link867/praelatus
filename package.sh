#!/bin/bash
# 
# Author: Mathew Robinson <chasinglogic@gmail.com>
# 
# This script builds praelatus with the frontend and deploys it to Github
# Requires go, npm, node, and curl be installed.
#

if [ "$GOOS" == "" ]; then
    echo "\$GOOS not set defaulting to linux"
    export GOOS="linux"
fi

function print_help() {
    echo "Usage: 
    ./package.sh tag_name name_of_release prelease_bool:optional

Examples:
    This would deploy to tag v0.0.1 naming the release MVP and specify it is a 
    prerelease

    ./package.sh v0.0.1 MVP true

    This would deploy to tag v1.0.0 naming the release Aces High and specify it
    as not a prerelease

    ./package.sh v1.0.0 \"Aces High\" false

    Alternatively prelease_bool can be omitted (defaults: false)

    ./package.sh v1.0.0 \"Aces High\""
}

if [ "$1" == "--help" ] || [ "-h" == "$1" ]; then
    print_help
    exit 0
fi

if [ "$#" -ne 3 ] && [ "$#" -ne 2 ]; then
    echo "wrong number of arguments $#"
    print_help
    exit 1
fi

TAG_NAME=$1
RELEASE_NAME=$2
PRELEASE=$3
STARTING_DIR=$(pwd)

echo "Tag Name: $TAG_NAME"
echo "Release Name: $RELEASE_NAME"
echo "Prelease: $PRELEASE"
echo ""

echo "Checking for dependencies..."
if ! [ -x "$(command -v go)" ]; then
    echo "You need to install the go tool. https://golang.org/download"
    exit 1
fi

if ! [ -x "$(command -v npm)" ]; then
    echo "You need to install npm. https://nodejs.org/en/download/"
    exit 1
fi

if ! [ -x "$(command -v node)" ]; then
    echo "You need to install node. https://nodejs.org/en/download/"
    exit 1
fi

if ! [ -x "$(command -v curl)" ]; then
    echo "You need to install curl"
    exit 1
fi

if ! [ -x "$(command -v yarn)" ]; then
    echo "yarn not detected attempting to install..."
    sudo npm install -g yarn
fi

if ! [ -x "$(command -v webpack)" ]; then
    echo "webpack not detected attempting to install..."
    sudo npm install -g webpack
fi

mkdir build
if [ "$?" -ne  0 ]; then
    echo "cleaning build directory..."
    rm -rf build
    mkdir build
fi

# create the final client directory
mkdir build/client

# install deps for backend
echo "installing dependencies for backend"
glide install &>/dev/null

# compile the backend
# TODO add cross-compilation
echo "compiling the backend"
export GOARCH="amd64"
go build -o build/praelatus

# get the frontend
echo "downloading the frontend"
git clone https://github.com/praelatus/frontend build/frontend &>/dev/null

# change to frontend git repo
cd build/frontend

# install frontend deps
echo "installing dependencies for frontend"
yarn install &>/dev/null

echo "compiling the frontend"
webpack -p &>/dev/null
mv build/debug/static ../client/
cp index.html ../client/index.html

echo "cleaning up"
rm -rf build/frontend

# go back to backend repo
cd $STARTING_DIR
echo $PWD

echo "building release tar"
tar czf praelatus-$TAG_NAME-$GOOS-$GOARCH.tar.gz build/*

# create the tag
# echo "Tagging release..."
# git tag -a $TAG_NAME -m $RELEASE_NAME

# # push the tag
# echo "Pushing tags..."
# git push --follow-tags

