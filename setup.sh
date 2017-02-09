#! /bin/bash

if [ -z "$1" -o -z "$2" ]; then
  echo "USAGE: ./setup.sh <MACKEREL_API_KEY> <MACKEREL_SERVICE_NAME>" >&2
  exit -1
fi

export MACKEREL_API_KEY=$1
export MACKEREL_SERVICE_NAME=$2
