#!/usr/bin/env bash
RUN_NAME="twitter"

mkdir -p output/bin output/conf

cp script/* output/
cp conf/*.yaml output/conf/
chmod +x output/*.sh

go build -o output/bin/${RUN_NAME}