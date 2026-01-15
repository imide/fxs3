[![PkgGoDev](https://pkg.go.dev/badge/github.com/imide/fxs3)](https://pkg.go.dev/github.com/imide/fxs3)

# Fx S3 Module

This module is for personal use only. Support is not guaranteed.
Bug fixes and contributions are welcome, however, unless sufficient interest is shown for this module, it will remain S3 only and for personal use.

> [Fx](https://uber-go.github.io/fx/) module for [S3](https://docs.aws.amazon.com/sdk-for-go/v2/).

<!-- TOC -->
* [Overview](#overview)
* [Installation](#installation)
* [Configuration](#configuration)
<!-- TOC -->

## Overview

This module provides to your Fx application a [s3.Client](https://docs.aws.amazon.com/sdk-for-go/v2/developer-guide/go_s3_code_examples.html),
that you can `inject` anywhere to interact with [S3](https://docs.aws.amazon.com/sdk-for-go/v2/).

## Installation

Install the module:

```shell
go get github.com/imide/fxs3
```

Then activate them in your application bootstrapper:

```go
// internal/bootstrap.go
package internal

import (
	"github.com/imide/fxs3"
	"github.com/ankorstore/yokai/fxcore"
)

var Bootstrapper = fxcore.NewBootstrapper().WithOptions(
	// load modules
	fxs3.FxS3Module,
	// ...
)
```

## Configuration

Configuration reference:

```yaml
# ./configs/config.yaml
app:
  name: app
  env: dev
  version: 0.1.0
  debug: true
modules:
  s3:
    endpoint: ${S3_BASE_ENDPOINT}
    aws_access_key_id: ${AWS_ACCESS_KEY_ID}
    aws_secret_access_key: ${AWS_SECRET_ACCESS_KEY}
```
