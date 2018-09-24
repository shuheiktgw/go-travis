go-travis
====

[![GitHub release](http://img.shields.io/github/release/shuheiktgw/go-travis.svg?style=flat-square)](https://github.com/shuheiktgw/go-travis/releases/latest)
[![Build Status](https://travis-ci.org/shuheiktgw/go-travis.svg?branch=master)](https://travis-ci.org/shuheiktgw/go-travis)
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
import travis "github.com/shuheiktgw/go-travis"

client := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "myTravisToken")

// list all builds of current user
builds, res, err := client.Builds.Find(context.Background(), nil)
```

Constructing it with the ``NewClient`` helper requires two arguments:
* The Travis CI API URL you wish to communicate with. Different Travis CI plans are accessed through different URLs. go-travis exposes constants for these URLs:
  * ``TRAVIS_API_DEFAULT_URL``: default *api.travis-ci.org* endpoint for the free Travis "Open Source" plan.
  * ``TRAVIS_API_PRO_URL``: the *api.travis-ci.com* endpoint for the paid Travis pro plans.
* A Travis CI token with which to authenticate. If you wish to run requests unauthenticated, pass an empty string. It is possible at any time to authenticate the Client instance with a Travis token or a Github personal access token. For more information see [Authentication](#Authentication).

### Authentication

The Client instance supports both authenticated and unauthenticated interaction with the Travis CI API. **Note** that both Pro and Enterprise plans will require almost all API calls to be authenticated.

#### Authenticated

The Client instance supports authentication with both Travis token and Github personal access token.

##### Authentication with Travis token

```go
// client.IsAuthenticated() returns true
client := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "myTravisToken")

// Job.Cancel will success
_, err := client.Job.Cancel(context.Background(), 12345)
```

##### Authentication with GitHub personal access token
Authentication with a Github personal access token will require some extra steps. [This GitHub help page](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) guides you thorough how to create one. 

```go
// client.IsAuthenticated() returns false
client := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "")

// client.IsAuthenticated() returns true
err := client.Authentication.UsingGithubToken("myGitHubToken")

// Job.Cancel will success
_, err := client.Job.Cancel(context.Background(), 12345)
```

#### Unauthenticated

It is possible to use the client unauthenticated. However some resources won't be accessible.

```go
client := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "")

// Builds.FindByRepoSlug is available without authentication
builds, resp, err := client.Builds.FindByRepoSlug(context.Background(), "shuheiktgw/go-travis", nil)

// Job.Cancel is unavailable without authentication
_, err := client.Job.Cancel(context.Background(), 12345)
```


## Acknowledgements

This library is originally forked from [Ableton/go-travis](https://github.com/Ableton/go-travis) and most of the credits of this library is attributed to them.
