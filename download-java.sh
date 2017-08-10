#!/bin/sh

JRE_VERSION=8u144
JRE_BUILD=01

JRE_TAR=jre-${JRE_VERSION}-linux-x64.tar.gz
if [ ! -e ${JRE_TAR} ]; then
    JRE_URL=https://download.oracle.com/otn-pub/java/jdk/${JRE_VERSION}-b${JRE_BUILD}/jre-${JRE_VERSION}-linux-x64.tar.gz
    wget -c -O ${JRE_TAR} --no-check-certificate --no-cookies --header "Cookie: oraclelicense=accept-securebackup-cookie" "${JRE_URL}"
    #bosh add blob $ROOTDIR/.downloads/jre.tar.gz java
fi
