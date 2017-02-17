#!/usr/bin/env python3
#
# Author: Mathew Robinson <chasinglogic@gmail.com>
# 
# This script builds praelatus with the frontend and deploys it to Github
# Requires the requests lib `pip3 install requests` and python3
#

import os
import shutil
import requests

from sys import argv, exit
from subprocess import run, PIPE

def print_help():
    print("""
Usage: 
    ./build.py tag_name name_of_release prelease_bool:optional

Examples:
    This would deploy to tag v0.0.1 naming the release MVP and specify it is a 
    prerelease

    ./build.py v0.0.1 MVP true

    This would deploy to tag v1.0.0 naming the release Aces High and specify it
    as not a prerelease

    ./build.py v1.0.0 "Aces High" false

    Alternatively prelease_bool can be omitted (defaults: false)

    ./build.py v1.0.0 "Aces High"
""")

if len(argv) != 4:
    print("wrong number of arguments")
    print_help()
    exit(1)

if "help" in argv[1] or "-h" in argv[1]:
    print_help()
    exit(0)

def getInd(lst, idx, default=None):
    try:
        return lst[idx]
    except:
        return default

jsn = {
    'tag_name': getInd(argv, 1),
    'release_name': getInd(argv, 2),
    'prerelease': bool(getInd(argv, 3, False)),
}

print(jsn)

starting_dir = os.getcwd()

try:
    os.mkdir("./build")
except FileExistsError:
    shutil.rmtree("./build")
    os.mkdir("./build")

def cmd(args):
    p = run(args, stdout=PIPE, stdin=PIPE, stderr=PIPE)
    if p.returncode != 0:
        print("Failing due to subprocess:", p.stderr)
        exit(p.returncode)

# install deps for backend
cmd(["glide", "install"])

# compile the backend
cmd(["go", "build", "-o", "build/praelatus"])

# get the frontend
cmd(["git", "clone", "https://github.com/praelatus/frontend", "build/frontend"])

# change to frontend git repo
os.chdir("build/frontend")

# install frontend deps
cmd(["npm", "install"])


# go back to backend repo
os.chdir(starting_dir)

# create the tag
print("Tagging release...")
cmd(["git", "tag", "-a", jsn['tag_name'], "-m", jsn["release_name"]])

# push the tag
print("Pushing tags...")
cmd(["git", "push", "--follow-tags"]) 

