## Flow Diagram

![](https://github.com/storj-thirdparty/connector-zenko/blob/master/README.assets/arch.drawio.png)


## Config Files
There are two config files that contain Storj network and Zenko connection information. The tool is designed so you can specify a config file as part of your tooling/workflow.

##### ```zenko_property.json```
Inside the ```./config``` directory there is a ```zenko_property.json``` file, with following information about your Zenko instance:

* zenkoEndpoint - S3 End point of Zenko Instance
* accessKeyID - S3 Access Key ID created in Zenko Instance
* secretAccessKey - S3 Secret Access Key created in Zenko Instance

##### ```storj_config.json```
Inside the ```./config``` directory a ```storj_config.json``` file, with Storj network configuration information in JSON format:

* apiKey - API Key created in Storj Satellite GUI
* satelliteURL - Storj Satellite URL
* encryptionPassphrase - Storj Encryption Passphrase.
* bucket - Name of the bucket to upload data into.
* uploadPath - Path on Storj Bucket to store data (optional) or "" or "/" (mandatory)
* serializedAccess - Serialized access shared while uploading data used to access bucket without API Key
* allowDownload - Set true to create serialized access with restricted download
* allowUpload - Set true to create serialized access with restricted upload
* allowList - Set true to create serialized access with restricted list access
* allowDelete - Set true to create serialized access with restricted delete
* notBefore - Set time that is always before notAfter
* notAfter - Set time that is always after notBefore

## Run
Zenko files are iterated through and streamed in 32KB chunks to the Storj network.

The following flags can be used with the ```store``` command:

* ```accesskey``` - Connects to the Storj network using a serialized access key instead of an API key, satellite url and encryption passphrase .
* ```share``` - Generates a restricted shareable serialized access with the restrictions specified in the Storj configuration file.

Once you have built the project you can run the following:

##### Get help
```
$ ./connector-zenko --help
```
##### Check version
```
$ ./connector-zenko version
```
##### Create backup from Zenko and upload them to Storj
```
$ ./connector-zenko store --zenko <path_to_zenko_config_file> --storj <path_to_storj_config_file>
```
##### Create backup files from Zenko and upload them to Storj bucket using Access Key
```
$ ./connector-zenko store --accesskey
```
##### Create backup files from Zenko and upload them to Storj and generate a Shareable Access Key based on restrictions in ```storj_config.json```.
```
$ ./connector-zenko store --share
```
		
## Testing
* The project has been tested on the following operating systems:
```
	* Windows
		* Version: 10 Pro
		* Processor: Intel(R) Core(TM) i3-5005U CPU @ 2.00GHz 2.00GHz

	* macOS Catalina
		* Version: 10.15.4
		* Processor: 2.5 GHz Dual-Core Intel Core i5

	* ubuntu
		* Version: 16.04 LTS
		* Processor: AMD A6-7310 APU with AMD Radeon R4 Graphics × 4
```		
