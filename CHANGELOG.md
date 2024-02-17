# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](https://semver.org/).


## [v3.1.0](https://github.com/virtualtam/ccache_exporter/releases/tag/v3.1.0) - 2024-02-17

### Added

- Add metrics for remote storage (e.g. Redis)
- Add test cases for Debian 12 and Ubuntu 24.04
- Update configuration parser to support new output format for `max_size` values
- Add Go vulnerability checker to the CI workflow

### Changed

- Require Go 1.22
- Refactor and document testdata generation utilities
- Update example Docker Compose stack and dashboard


## [v3.0.0](https://github.com/virtualtam/ccache_exporter/releases/tag/v3.0.0) - 2022-08-11

Major refactoring with breaking changes.

### Added

- Add support for ccache 3.7 and above, that comes with a new format for
  machine-readable statistics
- Add utilities to generate reference testdata using ccache on long-term-support
  Debian and Ubuntu releases

### Changed

- Switch to `main` as the default Git branch
- Refactor ccache wrapper and shell abstractions to handle both formats:
    - pre-3.7 mixed configuration/statistics output (--show-stats)
    - 3.7 and above configuration (--show-config) and statistics (--print-stats)
      output
- Relocate Prometheus collector
- Bump Go to 1.19 for CI, Promu and Docker builds
- Switch to Github Actions
- Rewrite README and CHANGELOG in Markdown
- Refactor tests and test helpers to make it easier to add golden test cases
- Update Prometheus Makefile library


## [v2.0.0](https://github.com/virtualtam/ccache_exporter/releases/tag/v2.0.0) - 2019-10-26

Major refactoring.

### Added

- Allow to specify the path to the ccache binary
- Setup promu from cross-platform exporter builds
- Add license information to all Go source headers
- Document all exported types and methods

### Changed

- Explicitly handle or silence errors
- Relocate parser code to this repository
- Refactor and simplify parser tests by switching to table-driven testing
- Refactor the codebase to adopt a flat package structure, with more idiomatic
  type and method names
- Trim build paths

## [v1.1.0](https://github.com/virtualtam/ccache_exporter/releases/tag/v1.1.0) - 2018-10-20

### Changed

- Add support for ccache 3.5
- Setup out-of-tree builds

## [v1.0.0](https://github.com/virtualtam/ccache_exporter/releases/tag/v1.0.0) - 2018-10-09

Initial release.

### Added

- Setup Go project with module support
- Gather metrics  from local ccache statistics
- Export Go, scraping and ccache metrics under ``/metrics``
- Setup a sandbox Docker Compose monitoring environment with Node Exporter, Prometheus and Grafana
