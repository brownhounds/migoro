# Changelog

## [0.1.3] - 2024-07-13

### Added

- Version Linting.
- Release script.
- Conditional run of `Build Docker Image` workflow based on conclusion of `Go Build Release`.
- `changelog-lint` pre commit hook.
- `changelog-lint` and `golangci` linting stages in github workflow.
- `NOTES.md` with development requirements.

### Changed

- Tag name to contain `v0.x.x` in alignment with golang module versioning.

## [0.1.2] - 2024-07-13

### Changed

- Remove `xgo` from github workflow for cross compilation.
- Package namespace from `migoro` to `github.com/brownhounds/migoro`

## [0.1.1] - 2024-07-10

### Added

- Changelog.
- Release Notes.
