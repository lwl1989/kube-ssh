#!/bin/bash

WORKSPACE=/home/www/go
MODULE=$1
ENV=$2
PIPELINE_ID=$3
BUILD_NUMBER=$4

#创建基本目录
projectDir=${WORKSPACE}/${MODULE}/${ENV}
logsDir=${projectDir}/logs
filesDir=${projectDir}/files
mkdir -p ${logsDir}
mkdir -p ${filesDir}
mkdir -p ${projectDir}/etc

#复制配置文件
cp "${projectDir}/tmp/etc/api_prod.yaml" "${projectDir}/etc/api.yaml"
cp "${projectDir}/tmp/etc/api_remote_local.yaml" "${projectDir}/etc/api_remote_local.yaml"
# 复制编译文件
cp "${projectDir}/tmp/${MODULE}" "${projectDir}/${PIPELINE_ID}-${BUILD_NUMBER}"
rm -rf mv "${projectDir}/dist"
cp -rf "${projectDir}/tmp/dist" "${projectDir}/dist"
# 重建软连接
if [ -e "${projectDir}/${MODULE}" ]; then
	unlink "${projectDir}/${MODULE}"
fi
ln -sf "${PIPELINE_ID}-${BUILD_NUMBER}" "${projectDir}/${MODULE}"

total_version=$(ls -l "${WORKSPACE}/${MODULE}/${ENV}" | grep "$PIPELINE_ID" |wc -l)
if [ "$total_version" -gt 10 ]; then
  rm `ls -ldt  "${WORKSPACE}/${MODULE}/${ENV}"/* | grep ${PIPELINE_ID} | tail -n 5| awk '{print $9}'`
fi