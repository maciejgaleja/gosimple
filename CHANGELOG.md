# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.4] - 2023-07-09
### Changed
- Get type for NoSQL.

## [0.0.3] - 2023-07-08
### Added
- Interface for NoSQL.
- Implementation for NoSQL using AWS DynamoDb.
- Implementation for KeyValue using NoSQL.

## [0.0.2] - 2023-05-31
### Added
- Function for hashing various inputs (strings, buffers, files).
- Type for file extension.

### Changed
- Type name for hash.
- Get function for Key Value store. Now it works as json.Unmarshal.

## [0.0.1] - 2023-05-30

### Added
- Storage interface and filesystem implementation.
- Key-value store and JSON file implementation.
- [draft] Authenticator interface and key-value store implementation.
- Hash implementation.
- Server for filesystem storage.
- This CHANGELOG file to hopefully serve as an evolving example of a
  standardized open source project CHANGELOG.

[unreleased]: https://github.com/maciejgaleja/gosimple/compare/v0.0.4...HEAD
[0.0.4]: https://github.com/maciejgaleja/gosimple/compare/v0.0.4...v0.0.4
[0.0.3]: https://github.com/maciejgaleja/gosimple/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/maciejgaleja/gosimple/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/maciejgaleja/gosimple/releases/tag/v0.0.1