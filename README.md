go-travis
====

[![GitHub release](http://img.shields.io/github/release/shuheiktgw/go-travis.svg?style=flat-square)](https://github.com/shuheiktgw/go-travis/releases/latest)
[![Actions Status](https://github.com/shuheiktgw/go-travis/workflows/Test/badge.svg)](https://github.com/shuheiktgw/go-travis/actions)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)
[![GoDoc](https://godoc.org/github.com/shuheiktgw/go-travis?status.svg)](https://godoc.org/github.com/shuheiktgw/go-travis)

go-travis is a Go client library to interact with the [Travis CI API V3](https://developer.travis-ci.com/).

## Installation

```bash
$ go get github.com/shuheiktgw/go-travis
```

## Usage

Interaction with the Travis CI API is done through a `Client` instance.

```go
import "github.com/shuheiktgw/go-travis"

client := travis.NewClient(travis.ApiOrgUrl, "TravisApiToken")

// List all the builds which belongs to the current user
builds, res, err := client.Builds.List(context.Background(), nil)
```

### URL
Currently, there are two possible options for Travis CI API URL.

- `https://api.travis-ci.org/`
- `https://api.travis-ci.com/`

You should know which URL your project belongs to, and hand it in to `NewClient` method as an argument. We provide two constants, `ApiOrgUrl` for `https://api.travis-ci.org/` and `ApiComUrl` for `https://api.travis-ci.com/`, so please choose one of them.

Travis CI is migrating projects in `https://api.travis-ci.org/` to `https://api.travis-ci.com/`, and please visit [their documentation page](https://docs.travis-ci.com/user/open-source-on-travis-ci-com#existing-private-repositories-on-travis-cicom) for more information on the migration.  


### Authentication

```go
client := travis.NewClient(travis.ApiOrgUrl, "TravisApiToken")

// Jobs.Cancel will success
_, _, err := client.Jobs.Cancel(context.Background(), 12345)
```

You can issue Travis API token and hand it in to `NewClient` method directly. You can issue your token by visiting your Travis CI [Profile page](https://travis-ci.com/profile) or using Travis CI [command line tool](https://github.com/travis-ci/travis.rb#readme). 

For more information on how to issue Travis CI API token, please visit [their documentation](https://docs.travis-ci.com/user/triggering-builds/).

### Unauthenticated client

It is possible to interact with the API without authentication. However, most resources are not accessible.

```go
client := travis.NewClient(travis.ApiOrgUrl, "")

// Builds.ListByRepoSlug is available without authentication
builds, resp, err := client.Builds.ListByRepoSlug(context.Background(), "shuheiktgw/go-travis", nil)

// Jobs.Cancel is unavailable without authentication
_, _, err = client.Jobs.Cancel(context.Background(), 12345)
```

## Standard Representation / Minimal Representation 

Travis CI API V3  provides two types of resource representations, standard and minimal. The API returns the resource you request in a standard representation and the resources related to the original resource in a minimal representation.

If you want **the related resources** in a standard representation, you need to [eager load](https://developer.travis-ci.com/eager-loading#eager%20loading) them by specifying `include` option.

For example, to eager load `repository` and `commits` when fetching a build, one can specify `Include` in `BuildOption`:


```go
opt := BuildOption{Include: []string{"build.repository", "build.commits"}}
build, _, err := client.Builds.Find(context.Background(), 123, &opt)
```  

## Contribution
Contributions are of course always welcome!

1. Fork shuheiktgw/go-travis (https://github.com/shuheiktgw/go-travis/fork)
2. Run `make install` to install dependencies
3. Create a feature branch
4. Commit your changes
5. Run test using `make test`
6. Create a Pull Request

See [`CONTRIBUTING.md`](https://github.com/shuheiktgw/go-travis/blob/master/CONTRIBUTING.md) for details.

## Acknowledgements

This library is originally forked from [Ableton/go-travis](https://github.com/Ableton/go-travis) and most of the credits of this library is attributed to them.
