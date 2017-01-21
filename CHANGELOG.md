Change Log
==========

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this
project adheres to [Semantic Versioning](http://semver.org/).


Unreleased
==========

### Added

- Global integer `DefaultContext` allows you to override the default amount of
  context lines (default 0).
- Function `NewFromPath(path string, region ctxerr.Region) (ctxerr.Ctx, error)`
  returns a new Ctx based on a path to an existing file.
- Method `Ctx.Error()` implementing the error interface, replacing `Ctx.String()`.
- Function `Parse(line string) (*Ctx, error)` parses typical linter output into
  Ctx structs.
  Returns error `ErrNoMatch` when unlikely to be a linter error.
  - Currently supports output of `gometalinter` and `npm run lint`.
- Command `ctx` enhances stdin with pretty errors and optional context.

### Changed

- Exported Position (a Region consists of two position).
- Exported all fields of Ctx and Region.
- Proper handling of Region with zero columns. It is now treated as a full line.

### Removed

- Methods of `Ctx`: `WithPath`, `WithContext` and `WithHint` as they're now exported as fields.
- Struct `Error` and method `Ctx.ToError()`. Use the exported `Err` field of `Ctx` instead.

## Fixed

- Tabs are now properly handled when padding and pointing to regions.
- No more panic when pointing outside of the line (to the right).


0.1.0
=====

### Added

- Initial public release.


[Unreleased]: https://github.com/nochso/ctxerr/compare/0.1.0...HEAD