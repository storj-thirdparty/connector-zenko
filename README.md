# storj-zenko connector (uplink v1.0.5)

[![Go Report Card](https://goreportcard.com/badge/github.com/storj-thirdparty/connector-zenko)](https://goreportcard.com/report/github.com/storj-thirdparty/connector-zenko)
![Cloud Build](https://storage.googleapis.com/storj-utropic-services-badges/builds/connector-framework/branches/master.svg)


## Overview

Command line application (on Windows/Linux/Mac) for taking data backup from Zenko to Storj. Application connects to Zenko server and the souce code for interaction to Storj for cloud storage which is written in Golang.

Zenko is infrastructure software to control Data in Multi-Cloud IT Environments without cloud lock-in and has features such as enabling unified data management from anywhere through a secure cloud portal, providing a single S3 endpoint through which data can be stored, retrieved and searched across any location.

### Features of storj-zenko:
* Connects S3 compatible cloud storages (e.g. Amazon AWS, Azure Blob, Google Cloud Storage, Wasabi) to the Zenko instance for backing up their data to StorJ V3 network.
* Upload any type of data from Zenko to Storj (single or multiple at once) whether it is a folder, document, data file, image, video, etc.
```
Usage:
  storj-zenko [command] <flags>

Available Commands:
  help        Help about any command
  store       Command to upload data to a Storj V3 network.
  version     Prints the version of the tool
```  
  
```store``` - Connect to the specified Zenko instance (default: ```zenko_property.json```). Backups of the Zenko storage are generated using tooling provided by Zenko then uploaded to the Storj network. Connect to a Storj v3 network using the access specified in the Storj configuration file (default: ```storj_config.json```).


Sample configuration files are provided in the ```./config``` folder.

## Requirements and Install
To build from scratch, [install the latest Go](https://golang.org/doc/install#install).

> Note: Ensure go modules are enabled (GO111MODULE=on)

#### Option #1: clone this repo (most common)
To clone the repo
```
git clone https://github.com/storj-thirdparty/storj-zenko.git
```
Then, build the project using the following:
```
cd storj-zenko
go build
```
#### Option #2: go get into your gopath
To download the project inside your GOPATH use the following command:
```
go get github.com/storj-thirdparty/storj-zenko
```
## Run (short version)
Once you have built the project run the following commands as per your requirement:

##### Get help
```
$ ./storj-zenko --help
```
##### Check version
```
$ ./storj-zenko --version
```
##### Create backup from Zenko and upload to Storj
```
$ ./storj-zenko store
```
## Documentation

For more information on runtime flags, configuration, testing, and diagrams, check out the [Detail](//github.com/storj-thirdparty/connector-zenko/wiki) or jump to:


* [Config Files](//github.com/storj-thirdparty/connector-zenko/wiki/#config-files)
* [Run (long version)](//github.com/storj-thirdparty/connector-zenko/wiki/#run)
* [Testing](//github.com/storj-thirdparty/connector-zenko/wiki/#testing)
* [Flow Diagram](//github.com/storj-thirdparty/connector-zenko/wiki/#flow-diagram)
