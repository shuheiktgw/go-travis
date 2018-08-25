go-travis
====

[![GitHub release](http://img.shields.io/github/release/shuheiktgw/go-travis.svg?style=flat-square)](release)
[![Build Status](https://travis-ci.org/shuheiktgw/travis.svg?branch=master)](https://travis-ci.org/shuheiktgw/ghbr)
[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

go-travis is a Go client library to interact with the [Travis CI API](https://docs.travis-ci.com/user/developer/).

go-travis requires Go version 1.10 or greater.

## Motivation

[shuheiktgw/go-travis](https://github.com/shuheiktgw/go-travis) is originally forked from [Ableton/go-travis](https://github.com/Ableton/go-travis). Unfortunately, the original library does not seem to be maintained any more, and it does not support the latest Travis CI API, V3.

shuheiktgw/go-travis updates the original library and supports the latest Travis CI API.

## Dive

```go
import (
    "log"
    travis "github.com/shuheiktgw/go-travis"
)

client := travis.NewDefaultClient("")
builds, _, _, resp, err := client.Builds.ListFromRepository("shuheiktgw/go-travis", nil)
if err != nil {
    log.Fatal(err)
}

// Now do something with the builds
```

## Installation

```bash
$ go get github.com/shuheiktgw/go-travis
```

## Usage

Interaction with the Travis CI API is done through a `Client` instance.

```go
import travis "github.com/shuheiktgw/go-travis"

client := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "asuperdupertoken")
```

Constructing it with the ``NewClient`` helper requires two arguments:
* The Travis CI API URL you wish to communicate with. Different Travis CI plans are accessed through different URLs. go-travis exposes constants for these URLs:
  * ``TRAVIS_API_DEFAULT_URL``: default *api.travis-ci.org* endpoint for the free Travis "Open Source" plan.
  * ``TRAVIS_API_PRO_URL``: the *api.travis-ci.com* endpoint for the paid Travis pro plans.
* A Travis CI token with which to authenticate. If you wish to run requests unauthenticated, pass an empty string. It is possible at any time to authenticate the Client instance with a Travis token or a Github token. For more information see [Authentication]().


### Services oriented design

The ``Client`` instance's ``Service`` attributes provide access to Travis CI API resources.

```go
opt := &travis.BuildListOptions{EventType: "pull request"}
builds, response, err := client.Builds.ListFromRepository("mygithubuser/mygithubrepo", opt)
if err != nil {
        log.Fatal(err)
}
```

**Non exhaustive list of implemented services**:
+ Authentication
+ Branches
+ Builds
+ Commits
+ Jobs
+ Logs
+ Repositories
+ Requests
+ Users

(*For an up to date exhaustive list, please check out the documentation*)


**Nota**: Service methods will often take an *Option* (sub-)type instance as input. These types, like ``BuildListOptions`` allow narrowing and filtering your requests.


### Authentication

The Client instance supports both authenticated and unauthenticated interaction with the Travis CI API. **Note** that both Pro and Enterprise plans will require almost all API calls to be authenticated.


#### Unuathenticated

It is possible to use the client unauthenticated. However some resources won't be accesible.

```go
unauthClient := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "")
builds, _, _, resp, err := unauthClient.Builds.ListFromRepository("mygithubuser/myopensourceproject", nil)
// Do something with your builds

_, err := unauthClient.Jobs.Cancel(12345)
if err != nil {
        // This operation is unavailable in unauthenticated mode and will
        // throw an error.
}
```

#### Authenticated

The Client instance supports authentication with both Travis token and Github token.

```go
authClient := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "mytravistoken")
builds, _, _, resp, err := authClient.Builds.ListFromRepository("mygithubuser/myopensourceproject",
nil)
// Do something with your builds

_, err := unauthClient.Jobs.Cancel(12345)
// Your job is succesfully canceled
```

However, authentication with a Github token will require and extra step (and request).

```go
authWithGithubClient := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "")
// authWithGithubClient.IsAuthenticated() will return false

err := authWithGithubClient.Authentication.UsingGithubToken("mygithubtoken")
if err != nil {
        log.Fatal(err)
}
// authWithGithubClient.IsAuthenticated()  will return true

builds, _, _, resp, err := authClient.Builds.ListFromRepository("mygithubuser/myopensourceproject",
nil)
// Do something with your builds
```


### Pagination

The services support resource pagination through the `ListOption` type. Every services `Option` type implements the `ListOption` type.

```go
client := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "mysuperdupertoken")
opt := &travis.BuildListOptions{}

for {
        travisBuilds, _, _, _, err := tc.Builds.ListFromRepository(target, opt)
        if err != nil {
                log.Fatal(err)
        }

        // Do something with the builds

        opt.GetNextPage(travisBuilds)
        if opt.AfterNumber <= 1 {  // Travis CI resources are one-indexed (not zero-indexed)
                break
        }
}
```


## Disclaimer

This library is based on [Ableton/go-travis](https://github.com/Ableton/go-travis) and most of the credits of this library is attributed to them.
