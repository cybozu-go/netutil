# Change Log

All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

### Changed

- Update dependencies in [#39](https://github.com/cybozu-go/netutil/pull/39)
    - Upgrade direct dependencies in go.mod
    - Update Ubuntu used for testing from 22.04 to 24.04

## [1.4.9] - 2024-11-12

### Changed
- Update dependencies in [#36](https://github.com/cybozu-go/netutil/pull/36)
  - Upgrade direct dependencies in go.mod

## [1.4.8] - 2024-07-12

### Changed
- Bump golang.org/x/net from 0.22.0 to 0.23.0 [#32](https://github.com/cybozu-go/netutil/pull/32)
- Update dependencies in [#33](https://github.com/cybozu-go/netutil/pull/33)
  - Upgrade direct dependencies in go.mod
  - Update GitHub actions
- Update release procedure to use gh command [#34](https://github.com/cybozu-go/netutil/pull/34)

## [1.4.7] - 2024-03-08

### Changed
- Update dependencies in [#30](https://github.com/cybozu-go/netutil/pull/30)
  - Upgrade direct dependencies in go.mod
  - Update Golang used for testing from 1.21 to 1.22

## [1.4.6] - 2023-11-14

### Changed
- Update dependencies in [#28](https://github.com/cybozu-go/netutil/pull/28)
  - Upgrade direct dependencies in go.mod
  - Update Golang used for testing from 1.20 to 1.21

## [1.4.5] - 2023-07-14

### Changed
- Update dependencies in [#25](https://github.com/cybozu-go/netutil/pull/25)
  - Upgrade direct dependencies in go.mod
  - Update Golang used for testing from 1.19 to 1.20

## [1.4.4] - 2023-01-20

### Changed
- Update dependencies in [#22](https://github.com/cybozu-go/netutil/pull/22)
    - Upgrade direct dependencies in go.mod
    - Update Golang used for testing from 1.18 to 1.19
    - Update GitHub Actions
- Fix for deprecated "io/ioutil" in [#22](https://github.com/cybozu-go/netutil/pull/22)

## [1.4.3] - 2022-08-26

### Changed
- Change LICENSE from MIT to Apache2.0 ([#18](https://github.com/cybozu-go/netutil/pull/18))
- Update dependencies ([#19](https://github.com/cybozu-go/netutil/pull/19))
    - Update Golang used for testing from 1.15 to 1.18

## [1.4.2] - 2021-07-30

### Changed
- Update dependencies ([#17](https://github.com/cybozu-go/netutil/pull/17))

## [1.4.1] - 2021-03-11

### Changed
- Make `DetectMTU` linux specific ([#14](https://github.com/cybozu-go/netutil/pull/14))
- Build and test on MacOS ([#14](https://github.com/cybozu-go/netutil/pull/14))

## [1.4.0] - 2021-03-05

### Added
- MTU detection utility ([#11](https://github.com/cybozu-go/netutil/pull/11))

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

[Unreleased]: https://github.com/cybozu-go/netutil/compare/v1.4.9...HEAD
[1.4.9]: https://github.com/cybozu-go/netutil/compare/v1.4.8...v1.4.9
[1.4.8]: https://github.com/cybozu-go/netutil/compare/v1.4.7...v1.4.8
[1.4.7]: https://github.com/cybozu-go/netutil/compare/v1.4.6...v1.4.7
[1.4.6]: https://github.com/cybozu-go/netutil/compare/v1.4.5...v1.4.6
[1.4.5]: https://github.com/cybozu-go/netutil/compare/v1.4.4...v1.4.5
[1.4.4]: https://github.com/cybozu-go/netutil/compare/v1.4.3...v1.4.4
[1.4.3]: https://github.com/cybozu-go/netutil/compare/v1.4.2...v1.4.3
[1.4.2]: https://github.com/cybozu-go/netutil/compare/v1.4.1...v1.4.2
[1.4.1]: https://github.com/cybozu-go/netutil/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/cybozu-go/netutil/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/cybozu-go/netutil/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/cybozu-go/netutil/compare/v1.1.0...v1.2.0
[1.1.0]: https://github.com/cybozu-go/netutil/compare/v1.0.1...v1.1.0
[1.0.1]: https://github.com/cybozu-go/netutil/compare/v1.0.0...v1.0.1
