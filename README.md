<h1 align="center">
  <br>
  S3 Tools
  <br>
  <br>
</h1>

<h4 align="center">S3 common utilities</h4>

<p align="center">
  <a href="https://travis-ci.org/pedrolopesme/s3tools"> <img src="https://api.travis-ci.org/pedrolopesme/s3tools.svg?branch=master" /></a>
  <a href="https://goreportcard.com/report/github.com/pedrolopesme/s3tools"> <img src="https://goreportcard.com/badge/github.com/pedrolopesme/s3tools" /></a>
  <a href="https://codeclimate.com/github/pedrolopesme/s3tools/maintainability"> <img src="https://api.codeclimate.com/v1/badges/802610de39eefd49d4e4/maintainability" /></a>
</p>
<br>

### Makefile

This project provides a Makefile with all common operations need to develop, test and build call-it.

* build: generates binaries
* test: runs all tests
* clean: removes binaries
* fmt: runs gofmt for all go files

### Running tests

Tests were write using [Testify](https://github.com/stretchr/testify). In order to run them, just type:

```shell
$ make test
```

### Credits

These are the main external packages that make up Call It:

| packages | description |
|---|---|
| **[AWS SDK Go](https://github.com/aws/aws-sdk-go)** | **AWS SDK for the Go programming language** |
| **[Cobra](https://github.com/spf13/cobra)** | **A Commander for modern Go CLI interactions** |
| **[Fsnotify](https://github.com/fsnotify/fsnotify)** | **Cross-platform file system notifications for Go.** |
| **[HCL](https://github.com/hashicorp/hcl)** | **HCL is the HashiCorp configuration language.** |
| **[go-difflib](https://github.com/pmezard/go-difflib)** | **Partial port of Python difflib package to Go.** |
| **[go-homedir](https://github.com/mitchellh/go-homedir)** | **Go library for detecting and expanding the user's home directory without cgo.** |
| **[go-jmespath](https://github.com/jmespath/go-jmespath)** | **Golang implementation of JMESPath.** |
| **[go-toml](https://github.com/pelletier/go-toml)** | **Go library for the TOML language.** |
| **[mapstructure](https://github.com/jmespath/mapstructure)** | **Go library for decoding generic map values into native Go structures.** |
| **[Properties](https://github.com/magiconair/properties)** | **Java properties scanner for Go** |
| **[Testify](https://github.com/stretchr/testify)** | **A toolkit with common assertions and mocks that plays nicely with the standard library** |


### License

[MIT](LICENSE.md)
~