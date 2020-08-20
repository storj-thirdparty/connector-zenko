# connector-zenko (uplink v1.0.5)

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/b76faa16299a4236ad85baa9f39b7294)](https://app.codacy.com/gh/storj-thirdparty/connector-zenko?utm_source=github.com&utm_medium=referral&utm_content=storj-thirdparty/connector-zenko&utm_campaign=Badge_Grade_Dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/storj-thirdparty/connector-zenko)](https://goreportcard.com/report/github.com/storj-thirdparty/connector-zenko)
![Cloud Build](https://storage.googleapis.com/storj-utropic-services-badges/builds/connector-zenko/branches/master.svg)

## Overview

Command line application (on Windows/Linux/Mac) for taking data backup from Zenko to Storj. Application connects to Zenko server and the souce code for interaction to Storj for cloud storage which is written in Golang.

Zenko is infrastructure software to control Data in Multi-Cloud IT Environments without cloud lock-in and has features such as enabling unified data management from anywhere through a secure cloud portal, providing a single S3 endpoint through which data can be stored, retrieved and searched across any location.

```
Usage:
  storj-zenko [command] <flags>

Available Commands:
  help        Help about any command
  store       Command to upload data to a Storj V3 network.
  version     Prints the version of the tool

```

`store` - Connect to the specified Zenko instance (default: `zenko_property.json`). 
Backups of the Zenko storage are generated using tooling provided by Zenko then uploaded to the Storj network. Connect to a Storj v3 network using the access specified in the Storj configuration file (default: `storj_config.json`).

The following flags  can be used with the `store` command:

* `accesskey` - Connects to the Storj network using a serialized access key instead of an API key, satellite url and encryption passphrase.
* `share` - Generates a restricted shareable serialized access with the restrictions specified in the Storj configuration file.

Sample configuration files are provided in the `./config` folder.



## Requirements and Install

To build from scratch, [install the latest Go](https://golang.org/doc/install#install).

> Note: Ensure go modules are enabled (GO111MODULE=on)



### Option #1: clone this repo (most common)

To clone the repo

```
git clone https://github.com/storj-thirdparty/connector-zenko.git
```

Then, build the project using the following:

```
cd connector-zenko
go build
```



### Option #2:  ``go get`` into your gopath

To download the project inside your GOPATH use the following command:

```
go get github.com/storj-thirdparty/connector-zenko
```


## Run (short version)

Once you have built the project run the following commands as per your requirement:

### Get help

```
$ ../connector-zenko --help
```

### Check version

```
$ ./connector-zenko --version
```

### Create backup from Zenko and upload to Storj

```
$ ./connector-zenko store
```


## Flow Diagram

![Flow Diagram](/_images/arch.drawio.png ':include :type=iframe width=100% height=1000px')
