#!/bin/sh

GO=go

echo "Changes detected - killing server..."
pkill -f gotea-example
echo "Recompiling and starting..."
# Instead of run, we use build as we need to specify an output name
# because we need to be able to shut the server down by name
$GO build -o=gotea-example *.go && ./gotea-example "$@" &