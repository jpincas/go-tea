#!/bin/sh

find ../*go ./*go ../components ./templates | entr ./restart.sh "$@"