ctxerr
======

[![MIT license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/nochso/ctxerr?status.svg)](https://godoc.org/github.com/nochso/ctxerr)
[![GitHub release](https://img.shields.io/github/release/nochso/ctxerr.svg)](https://github.com/nochso/ctxerr/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/nochso/ctxerr)](https://goreportcard.com/report/github.com/nochso/ctxerr)
[![Build Status](https://travis-ci.org/nochso/ctxerr.svg?branch=master)](https://travis-ci.org/nochso/ctxerr)
[![Coverage Status](https://coveralls.io/repos/github/nochso/ctxerr/badge.svg?branch=master)](https://coveralls.io/github/nochso/ctxerr?branch=master)

**ctxerr** is a [Go][] library and CLI utility for pretty-printing linter/parser errors.

Instead of just describing what's wrong, it lets you point at the wrong input:

```
narf.txt:1:8: string is missing closing quote
 1 | "foo bar
   |         ^ missing closing quote
```

- Create pretty error messages for linters, parsers, etc. by pointing at the
  relevant position in the source code.
- Parse existing error messages into `Ctx` structs.
- CLI utility `ctx` scans STDIN for linter errors and enhances them with the
  reported source line.

---

- [Installation](#installation)
- [Documentation](#documentation)
- [Changes](#changes)
- [License](#license)


Installation
------------


### Library

```bash
go get -u github.com/nochso/ctxerr
```

### CLI utility `ctx`

```bash
go install github.com/nochso/ctxerr/cmd/ctx
```


Documentation
-------------

### Library

The [GoDoc](https://godoc.org/github.com/nochso/ctxerr) pages contain plenty
of examples.


### CLI utility `ctx`

```shell
$ ctx -h
ctx 1.0.0-beta
Pretty prints parser errors from stdin.

Possible input:
  path/file.ext:1:5: some error on line 1, column 5
  file.ext:123: column is optional and so is the message:
  file.ext:1

Usage:
  ctx < log.txt
  gometalinter . | ctx

Flags:
  -context NUM
        print NUM lines of context surrounding an error. negative for all, positive for limited and 0 for none (default)
  -no-color
        disable any color output
  -pessimistic
        print only matching errors, ignore everything else
```


Changes
-------

All notable changes to this project will be documented in the [changelog].

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this
project adheres to [Semantic Versioning](http://semver.org/).


License
-------

This project is released under the [MIT license](LICENSE).


[changelog]: CHANGELOG.md
[releases]: https://github.com/nochso/ctxerr/releases
[Go]: https://golang.org