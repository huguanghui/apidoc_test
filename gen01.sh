#!/usr/bin/env bash

set -x

# 重新生成
rm -rf www/apidoc
apidoc -i routers -o www/apidoc
