#!/usr/bin/env bash
RUN_NAME="twitter"

mkdir -p output/bin output/conf output/static/html output/static/css

cp script/* output/
cp conf/*.yaml output/conf/
cp static/html/* output/static/html/
cp static/css/* output/static/css/
chmod +x output/*.sh

go build -o output/bin/${RUN_NAME}