#!/bin/bash

## Starts the fake app, reloading on change

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
dir="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"
cd "$dir"

ulimit -n 2048
air
