# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

## v0.5.0
### Removed
- Storage server. It belongs in a separate project.
### Added
- Support for nested objects in filesystem storage.
- Support for S3 storage backend.
### Fixed
- Returned value of hash functions.

## v0.0.4
### Changed
- Get type for NoSQL.

## v0.0.3 
### Added
- Interface for NoSQL.
- Implementation for NoSQL using AWS DynamoDb.
- Implementation for KeyValue using NoSQL.

## v0.0.2
### Added
- Function for hashing various inputs (strings, buffers, files).
- Type for file extension.

### Changed
- Type name for hash.
- Get function for Key Value store. Now it works as json.Unmarshal.

## v0.0.1
### Added
- Storage interface and filesystem implementation.
- Key-value store and JSON file implementation.
- [draft] Authenticator interface and key-value store implementation.
- Hash implementation.
- Server for filesystem storage.
- This CHANGELOG file to hopefully serve as an evolving example of a
  standardized open source project CHANGELOG.

