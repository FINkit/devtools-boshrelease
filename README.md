# bosh release for the FIKkit build stack

[![Build Status](https://travis-ci.org/garethjevans/buildstack-boshrelease.svg?branch=master)](https://travis-ci.org/garethjevans/buildstack-boshrelease)

# Overview

TODO

# Installation

To deploy the buildstack the following network is required:

```
networks:
- name: buildstack
```

Deploy using:

```
bosh -d buildstack deploy \
    --vars-store ../buildstack-deployment-vars.yml \
    -o operations/use-cloudsql-database.yml \
    buildstack.yml
```

# Ops Files

There are a number of provided bosh2 ops files that can be used to customise the deployment:

Operations File | Use
------------ | -------------
use-cloudsql-database.yml | Configure the deployment to use a Google Cloud SQL database, this will also install a 'cloud_sql_proxy' on each node.
use-alternative-network.yml | Uses an alternative network name to the default 'buildstack' option.
