ctxerr
======

[![MIT license](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/nochso/ctxerr?status.svg)](https://godoc.org/github.com/nochso/ctxerr)
[![GitHub release](https://img.shields.io/github/release/nochso/ctxerr.svg)](https://github.com/nochso/ctxerr/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/nochso/ctxerr)](https://goreportcard.com/report/github.com/nochso/ctxerr)
[![Build Status](https://travis-ci.org/nochso/ctxerr.svg?branch=master)](https://travis-ci.org/nochso/ctxerr)
[![Coverage Status](https://coveralls.io/repos/github/nochso/ctxerr/badge.svg?branch=master)](https://coveralls.io/github/nochso/ctxerr?branch=master)

**ctxerr** is a [Go][] library for printing pretty parser errors.

Instead of just describing what's wrong, it lets you point at the wrong input:

```
narf.txt:1:8:
 1 | "foo bar
   |         ^ missing closing quote
```


Installation
------------

```bash
go get -u github.com/nochso/ctxerr
```

Documentation
-------------

The [GoDoc](https://godoc.org/github.com/nochso/ctxerr) pages contain plenty
of examples.


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