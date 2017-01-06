Change Log
==========

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/) and this
project adheres to [Semantic Versioning](http://semver.org/).


Unreleased
==========

### Changed

- Exported Position (a Region consists of two position).
- Exported all fields of Ctx.
- Exported all fields of Err.
- Proper handling of Region with zero columns. It is now treated as a full line.

### Removed

- Method `Error.Inner()` as it is now exported as field `Inner`.
- Methods of `Ctx`: `WithPath`, `WithContext` and `WithHint` as they're now exported as fields.

0.1.0
=====

### Added

- Initial public release.


[Unreleased]: https://github.com/nochso/ctxerr/compare/0.1.0...HEAD