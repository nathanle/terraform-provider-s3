Pending plan: Build provider with support for things like bucket policies.


# terraform-provider-s3

## Description
A more generic S3 provider for Terraform.  The objective is to have an independent S3 provider that is compatible with AWS, RGW, GCE, etc.

## Requirements
* [Terraform](https://github.com/hashicorp/terraform)
* [Minio Go Library](https://github.com/minio/minio-go)

## Installation
The installation process assumes a properly installed and configured Go environment.
```
go install github.com/nathanle/terraform-provider-s3@latest
```
The provider will be available in ```$GOPATH/bin```

## Usage

### Provider Configuration
The provider requires some variables to be configured in order to gain access to the S3 compatible server:

* **s3_server**: S3 Server
* **s3_region**: S3 Server region (default: us-east-1)
* **s3_access_key**: S3 Server Access Key
* **s3_secret_key**: S3 Server Secret Key
* **s3_api_signature**: S3 Server API Signature (type: string, options: v2 or v4, default: v4)
* Supported cloud storage providers:
   * AWS Signature Version 4
      * Amazon S3
      * Minio
   * AWS Signature Version 2
      * Google Cloud Storage (Compatibility Mode)
      * Openstack Swift + Swift3 middleware
      * Ceph Object Gateway
      * Riak CS
* **s3_ssl**: Connect using SSL
* **s3_debug**: Enable Debug messages

#### `s3`
```
provider "s3" {
    s3_server        = "http://localhost:9000"
    s3_region        = "custom-region-1"
    s3_access_key    = "Access Key here"
    s3_secret_key    = "Secret Key here"
    s3_api_signature = "v2"
    s3_ssl           = false
    s3_debug         = true
}
```

### Resource Configuration (s3_bucket)
```s3_bucket``` resources represent a bucket in the S3 server.  It requires a bucket name to operate:

* **bucket**: Name of the bucket to use

```
resource "s3_bucket" "resource_name" {
	bucket = "my_bucket_name"
}
```


### Resource Configuration (s3_file)
```s3_file``` resources represent a file to be uploaded to the S3 server or downloaded from it.  It currently takes the following arguments:
* **bucket**: Bucket in the S3 server
* **name**: S3 Object name
* **file_path**: Local file path where to read or save the object
* **content_type**: The content type of the object.  Defaults to: ```application/octet-stream```
* **debug**: Print debug messages
```
resource "s3_file" "resource_name" {
    bucket       = "my_bucket_name"
    name         = "my_object_name"
    file_path    = /tmp/my_file.bin
    content_type = "application/octet-stream"
    debug        = true
}
```
