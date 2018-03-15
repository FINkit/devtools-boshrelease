# bosh release for the FinKit build stack

[![Build Status](https://travis-ci.org/FINkit/devtools-boshrelease.svg?branch=master)](https://travis-ci.org/FINkit/devtools-boshrelease)

## 1. Overview

This is a bosh release of Gerrit/Jenkins/Sonarqube & Nexus all linked together configured with a base set of plugins that can be used to build FINkit services.

## 2. Release

### Create

```
bosh -e MY_ENV \
  create-release
```
### Upload

```
bosh -e MY_ENV \
  upload-release
```
