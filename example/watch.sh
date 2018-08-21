#!/bin/sh

find ./*go ../components ./templates | entr ./restart.sh "$@"