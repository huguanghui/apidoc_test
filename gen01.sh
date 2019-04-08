#!/usr/bin/env bash

set -x

rm -rf www/apidoc
apidoc -i routers -o www/apidoc