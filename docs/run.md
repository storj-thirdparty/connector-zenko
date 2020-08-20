# Run

> Zenko files are iterated through and streamed in 32KB chunks to the Storj network.

The following flags can be used with the `store` command:

* `accesskey` - Connects to the Storj network using a serialized access key instead of an API key, satellite url and encryption passphrase.
* `share` - Generates a restricted shareable serialized access with the restrictions specified in the Storj configuration file.

Once you have built the project you can run the following:

## Get help

```
$ ./connector-zenko --help
```

## Check version

```
$ ./connector-zenko version
```

## Create backup from Zenko and upload them to Storj

```
$ ./connector-zenko store --zenko <path_to_zenko_config_file> --storj <path_to_storj_config_file>
```

## Create backup files from Zenko and upload them to Storj bucket using Access Key

```
$ ./connector-zenko store --accesskey
```

## Create backup files from Zenko and upload them to Storj and generate a Shareable Access Key based on restrictions in `storj_config.json`

```
$ ./connector-zenko store --share
```
