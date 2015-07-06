# go-travis

go-travis is a Go client library for accessing the [Travis CI API][].

**Documentation:** [![GoDoc](https://godoc.org/github.com/AbletonAppDev/go-travis/travis?status.svg)](https://godoc.org/github.com/AbletonAppDev/go-travis/travis)
**Build Status:** [![Build Status](https://travis-ci.org/AbletonAppDev/go-travis.svg?branch=master)](https://travis-ci.org/AbletonAppDev/go-travis)

go-travis requires Go version 1.1 or greater.

## Usage ##

```go
import travis "github.com/AbletonAppDev/go-travis"
```

Construct a new Travis CI client, then use the various services on the client to
access different parts of the Travis CI API.  For example, to list all
builds for the authenticated user:

```go
client := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "AQFvXR7r88s2Db5-dMYo3g")
builds, _, err := client.Builds.List(nil)
```

*Nota*: The ``NewClient`` constructor function takes the Travis CI API target url as first argument. The library exposes two constants for you to select the endpoint you wish to use:
  + ``TRAVIS_API_DEFAULT_URL``: the default api.travis-ci.org endpoint for the travis default plan.
  + ``TRAVIS_API_PRO_URL``: the api.travis-ci.com endpoint for the travis pro plan.

Some API methods have optional parameters that can be passed.  For example,
to list builds for the "MyGithubUser/chucknorris" repository:

```go
client := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "AQFvXR7r88s2Db5-dMYo3g")
repos, _, err := client.Repositories.ListByRepository("MyGithubUser/chucknorris", nil)
```

Moreover, some API methods exposes an opt (generally of a Options suffixed type, like ListBuildsOptions) parameters that can be passed in order to add filters to the request. For example, to list builds of pull request only for repository "MyGithubUser/chucknorris":

```go
client := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "AQFvXR7r88s2Db5-dMYo3g")
opt := &ListBuildsOptions{Slug: "MyGithubUser/chucknorris", EventType: "pull_request"}
repos, _, err := client.Repositories.ListByRepository(opt)
```

### Authentication ###

The go-travis client supports both authenticated and unauthenticated interactions with the Travis CI API.

```go
// Unauthenticated client
unauthClient := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "")  // Unauthenticated client

// Client authenticated with a Travis CI API access token
authClient := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "AQFvXR7r88s2Db5-dMYo3g")

// Authenticate your client using a Github personal token
client := travis.NewClient(travis.TRAVIS_API_DEFAULT_URL, "")
access_token, _, err := client.Authentication.UsingGithubToken("mygithubtoken")  // Your client is now authenticated
```


## Roadmap ##

This library is being initially developed for internal applications at
[Ableton](http://ableton.com), so API methods will likely be implemented in the order that they are
needed by that application. Eventually, we would like to cover the entire
Travis API, so contributions are of course [always welcome][contributing].

[contributing]: CONTRIBUTING.md

## Maintainers

* Theo Crevon <theo.crevon@ableton.com>

## Disclaimer

This library design is heavily inspired from the amazing Google's [go-github](https://github.com/google/go-github) library. Some pieces of code have been directly extracted from there too. Therefore any obvious similarities would not be adventitious.

## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
