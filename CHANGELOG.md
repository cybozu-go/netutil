# Change Log

All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

## [1.4.3] - 2022-08-26

### Changed
- Change LICENSE from MIT to Apache2.0 (#18)
- Update dependencies (#19)
    - Update Golang used for testing from 1.15 to 1.18

## [1.4.2] - 2021-07-30

### Changed
- Update dependencies (#17)

## [1.4.1] - 2021-03-11

### Changed
- Make `DetectMTU` linux specific (#14)
- Build and test on MacOS (#14)

## [1.4.0] - 2021-03-05

### Added
- MTU detection utility (#11)

## [1.3.0] - 2020-12-15

### Added
- Add `IPAdd` and `IPDiff` for IPv4/v6 address calculations
- Add `EqualIP` matcher for Ginkgo tests

### Deprecated
- `IP4ToInt` and `IntToIP4`: use `IPAdd` and `IPDiff` instead.

## [1.2.0] - 2018-09-14
### Added
- Opt in to [Go modules](https://github.com/golang/go/wiki/Modules).

## [1.1.0] - 2016-09-11
### Added
- `CipherSuiteString` returns string for tls.TLS_* constants.
- `TLSVersionString` returns string for tls.Version* constants.

## [1.0.1] - 2016-08-28
### Added
- `IsNetworkUnreachable`, `IsConnectionRefused`, `IsNoRouteToHost` functions to identify network errors.

[Unreleased]: https://github.com/cybozu-go/netutil/compare/v1.4.3...HEAD
[1.4.3]: https://github.com/cybozu-go/netutil/compare/v1.4.2...v1.4.3
[1.4.2]: https://github.com/cybozu-go/netutil/compare/v1.4.1...v1.4.2
[1.4.1]: https://github.com/cybozu-go/netutil/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/cybozu-go/netutil/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/cybozu-go/netutil/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/cybozu-go/netutil/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/cybozu-go/netutil/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/cybozu-go/netutil/compare/v1.0.0...v1.0.1
