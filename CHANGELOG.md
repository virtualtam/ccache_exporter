# Change Log

All notable changes to this project will be documented in this file.

The format is based on `Keep a Changelog`_ and this project adheres to
`Semantic Versioning`_.

.. _Keep A Changelog: http://keepachangelog.com/
.. _Semantic Versioning: https://semver.org/

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
