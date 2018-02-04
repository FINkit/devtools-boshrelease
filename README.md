# bosh release for the FINkit build stack

[![Build Status](https://travis-ci.org/FINkit/buildstack-boshrelease.svg?branch=master)](https://travis-ci.org/FINkit/buildstack-boshrelease)

# Overview

This is a bosh release of Gerrit/Jenkins/Sonarqube & Nexus all linked together configured with a base set of plugins that can be used to build FINkit services.

# Installation

To deploy the buildstack the following network is required:

```
networks:
- name: buildstack
```

Deploy using:

```
bosh -d buildstack deploy \
  --vars-store buildstack-deployment-vars.yml \
  buildstack.yml
```

This will create the variable store `buildstack-deployment-vars.yml` and all generated keys & passwords will be stored in here.

# Ops Files

There are a number of provided bosh2 ops files that can be used to customise the deployment:

Operations File | Use
------------ | -------------
use-alternative-network.yml | Uses an alternative network name to the default `buildstack` option.
