#!/bin/bash

if [[ "$1" == "" ]]; then
  exit 1
fi

sed -E -e "s/\"git-bus-factor [^\"]+\"/\"git-bus-factor $1\"/" -i.bak main.go || exit 1
rm -f main.go.bak || exit 1

