#!/bin/bash
# 
# Author: Mathew Robinson <chasinglogic@gmail.com>
# 
# This script builds praelatus with the frontend and deploys it to Github
# Requires go, npm, node, and curl be installed.
#


function parse_git_branch {
    ref=$(git symbolic-ref HEAD 2> /dev/null) || return
    echo "${ref#refs/heads/} "
}

BRANCH=$(parse_git_branch)

# if [ $BRANCH != "master" ] && [ $BRANCH != "develop" ]; then
#     echo "you aren't on master or develop, refusing to package a release"
#     exit 1
# fi

if [ "$GOOS" == "" ]; then
    echo "\$GOOS not set defaulting to linux"
    export GOOS="linux"
fi

function check_if_success() {
    if [ $? -ne 0 ]; then
        echo "error running last command"
        exit $?
    fi
}

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

if ! [ -x "$(command -v glide)" ]; then
    echo "glide not detected attempting to install..."
    go get github.com/Masterminds/glide
    if ! [ -x "$(command -v glide)" ]; then
        echo "installed glide but \$GOBIN isn't in \$PATH"
        exit 1
    fi
fi

if [ -d "build" ]; then
    echo "cleaning build directory..."
    rm -rf build
fi

# create the final build and build/client directories
mkdir -p build/client

# install deps for backend
echo "installing dependencies for backend"
glide install &>/dev/null

# compile the backend
# TODO add cross-compilation
echo "compiling the backend"
export GOARCH="amd64"
go build -o build/praelatus &>/dev/null
check_if_success

# get the frontend
echo "downloading the frontend"
git clone https://github.com/praelatus/frontend build/frontend &>/dev/null

# change to frontend git repo
cd build/frontend

# # install frontend deps
# echo "installing dependencies for frontend"
# yarn install &>/dev/null

# echo "compiling the frontend"
# webpack -p &>/dev/null
# mv build/debug/static ../client/
# cp index.html ../client/index.html

echo "cleaning up"
cd $STARTING_DIR
rm -rf build/frontend

echo "building release tar"
cd build

PACKAGE_NAME="praelatus-$TAG_NAME-$GOOS-$GOARCH.tar.gz"

if ! [ -z "$PRERELEASE" ]; then
    echo "is a PRELEASE"
    PACKAGE_NAME="praelatus-$TAG_NAME-prelease-$GOOS-$GOARCH.tar.gz"
fi

echo $PACKAGE_NAME
if [ -f $PACKAGE_NAME ]; then
    echo "old package detected removing..."
    rm $PACKAGE_NAME
fi

tar czf ../$PACKAGE_NAME *

# # create the tag
# echo "tagging release..."
# if [ $TAG_NAME == "nightly" ]; then
#     # we just move nightly up to the current commit instead of retagging
#     git tag -af $TAG_NAME -m $RELEASE_NAME
# else
#     git tag -a $TAG_NAME -m $RELEASE_NAME
# fi

# # push the tag
# echo "Pushing tags..."
# git push --follow-tags

