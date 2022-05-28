#!/bin/bash

# Checking arguments
if [ $# -ne 2 ]
then
  echo "!!! INVALID INPUT !!! Usage: $(basename "$0") {package name} {collector name}"
  exit 1
fi

# Checking working dir
if [ "${PWD##*/}" != "collectors" ]
then
  echo "!!! INVALID WORKING DIR !!! cd to xray-agent-linux/collectors"
  exit 1
fi

# Copy template
newPackageName="$1"
newCollectorName="$2"
cp -r ./_template ./"$newPackageName"

# Sed names
for filename in ./"$newPackageName"/*; do
  sed "s/template/$newPackageName/" -i ./"$filename";
  sed "s/Template/$newCollectorName/" -i ./"$filename";
done