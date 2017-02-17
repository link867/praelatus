#!/bin/bash
# 
# Author: Mathew Robinson <chasinglogic@gmail.com>
# 
# This script builds praelatus with the frontend and deploys it to Github
# Requires go, npm, node, and curl be installed.
#

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

if [[ "$1" == *"help"* || "-h" == "$1" ]]; then
    print_help
    exit 0
fi

if [[ "$#" -ne 4 && "$#" -ne 3 ]]; then
    echo "wrong number of arguments"
    print_help
    exit 1
fi

TAG_NAME = $1
RELEASE_NAME = $2
PRELEASE = $3
STARTING_DIR = $(pwd)

echo $TAG_NAME
echo $RELEASE_NAME
echo $PRELEASE
echo $STARTING_DIR

exit 1

mkdir build
if [ "$?" -ne  0 ]; then
    rm -rf build
    mkdir build
fi

# install deps for backend
glide install

# compile the backend
go build -o build/praelatus

# get the frontend
git clone https://github.com/praelatus/frontend build/frontend

# change to frontend git repo
cd build/frontend

# install frontend deps
npm install

# go back to backend repo
cd $STARTING_DIR

# create the tag
echo "Tagging release..."
git tag -a $TAG_NAME -m $RELEASE_NAME

# push the tag
echo "Pushing tags..."
git push --follow-tags

