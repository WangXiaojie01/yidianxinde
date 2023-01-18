#!/bin/bash
PROJ_PATH=/cygdrive/e/PerformanceAnalysis/engine_center
HTML_PATH=/cygdrive/d/nginx-1.14.0/engine_center/build

SVN_USERNAME=sdk_mac
SVN_PASSWORD=lfdalld

cd ${PROJ_PATH}

if [[ `echo $?` != 0 ]]; then
  echo goto package_web error eccurs
  exit -1;
fi

svn update --username=${SVN_USERNAME} --password=${SVN_PASSWORD} --no-auth-cache
if [[ `echo $?` != 0 ]]; then
  echo svn error eccurs
  exit -1;
fi

yarn build

rm -rf ${HTML_PATH}/*

cp -a ${PROJ_PATH}/build/* ${HTML_PATH}
