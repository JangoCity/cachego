# Cachego

[![Build Status](https://img.shields.io/travis/fabiorphp/cachego/master.svg?style=flat-square)](https://travis-ci.org/fabiorphp/cachego)
[![Coverage Status](https://img.shields.io/coveralls/fabiorphp/cachego/master.svg?style=flat-square)](https://coveralls.io/github/fabiorphp/cachego?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/fabiorphp/cachego)

Simple interface around cache drivers

## Installation

Cachego requires Go 1.8 or later.

```
go get github.com/fabiorphp/cachego
```

If you want to get an specific version, please use the example below:

```
go get gopkg.in/fabiorphp/cachego.v0
```

## Usage Examples

### Memcached

```go
package main

import (
    "github.com/fabiorphp/cachego"
    "github.com/bradfitz/gomemcache/memcache"
)

var cache cachego.Cache

func init() {
    cache = cachego.NewMemcached(memcached.New("localhost:11211"))
}
```

### Redis

```go
package main

import (
    "github.com/fabiorphp/cachego"
    "gopkg.in/redis.v4"
)

var cache cachego.Cache

func init() {
    cache = cachego.NewRedis(
        redis.NewClient(&redis.Options{
            Addr: ":6379",
        }),
    )
}
```

### File

```go
package main

import (
    "github.com/fabiorphp/cachego"
)

var cache cachego.Cache

func init() {
    cache = cachego.NewFile(
        "/cache-dir/",
    )
}
```

### Map

```go
package main

import (
    "github.com/fabiorphp/cachego"
)

var cache cachego.Cache

func init() {
    cache = NewMap()
}
```

### MongoDB

```go
package main

import (
    "github.com/fabiorphp/cachego"
    "gopkg.in/mgo.v2"
)

var cache cachego.Cache

func init() {
    session, _ := mgo.Dial(address)

    cache = cachego.NewMongo(
        session.DB("cache").C("cache"),
    )
}
```

### Chain

```go
package main

import (
    "github.com/fabiorphp/cachego"
)

var cache cachego.Cache

func init() {
    memcached := cachego.NewMemcached(
        memcached.New("localhost:11211"),
    )

    redis := cachego.NewRedis(
        redis.NewClient(&redis.Options{
            Addr: ":6379",
        }),
    )

    file := cachego.NewFile(
        "/cache-dir/"
    )

    cache = cachego.NewChain(
        cachego.NewMap(),
        memcached,
        redis,
        file,
    )
}
```

### Usage

```go
package main

import (
    "github.com/fabiorphp/cachego"
    "github.com/bradfitz/gomemcache/memcache"
)

func main() {
    cache.Save("foo", "bar")
    cache.Save("john", "doe")

    value, err := cache.Fetch("foo")

    multiple := cache.FetchMulti([]string{"foo", "john"})

    if cache.Contains("foo") {
        cache.Delete("foo")
    }

    cache.Flush()
}
```

## Documentation

Read the full documentation at [https://godoc.org/github.com/fabiorphp/cachego](https://godoc.org/github.com/fabiorphp/cachego).

## Development

### Requirements

- Install [docker](https://docs.docker.com/install/) and [docker-compose](https://docs.docker.com/compose/install/)
- Install [go dep](https://github.com/golang/dep)

### Run tests
```sh
// tests
$ make test

// test with coverage
$ make test-coverage

// clean-up
$ make clean

// configure (download dependencies and run docker containers)
$ make configure
```

## License

This project is released under the MIT licence. See [LICENSE](https://github.com/fabiorphp/cachego/blob/master/LICENSE) for more details.
