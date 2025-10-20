# Contributing to Compass Protocol

Thank you for your interest in contributing to Compass Protocol! This document provides guidelines and instructions for contributing to this project.

## Code of Conduct

Please be respectful and constructive in all interactions with the project and its community.

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git
- (Optional) golangci-lint for linting

### Setting Up Your Development Environment

1. Fork the repository on GitHub
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/compass-protocol.git
   cd compass-protocol
   ```

3. Add the upstream repository:
   ```bash
   git remote add upstream https://github.com/brewinski/compass-protocol.git
   ```

4. Install dependencies:
   ```bash
   make deps
   ```

## Development Workflow

### Making Changes

1. Create a new branch for your changes:
   ```bash
   git checkout -b feature/my-new-feature
   ```

2. Make your changes and test them:
   ```bash
   make test
   make build
   ```

3. Format your code:
   ```bash
   make fmt
   ```

4. Run linting (if golangci-lint is installed):
   ```bash
   make lint
   ```

### Committing Changes

- Write clear, concise commit messages
- Follow the conventional commits format when possible:
  - `feat:` for new features
  - `fix:` for bug fixes
  - `docs:` for documentation changes
  - `test:` for test changes
  - `chore:` for maintenance tasks

### Submitting a Pull Request

1. Push your changes to your fork:
   ```bash
   git push origin feature/my-new-feature
   ```

2. Open a pull request on GitHub
3. Describe your changes and link any related issues
4. Wait for review and address any feedback

## Testing

- All code changes should include appropriate tests
- Run tests with `make test`
- Check test coverage with `make test-coverage`
- Ensure all tests pass before submitting a PR

## Building

- Build the project with `make build`
- The binary will be created in the project root as `compass-protocol`

## Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting (run `make fmt`)
- Run `golangci-lint` to catch common issues (run `make lint`)
- Keep functions focused and well-documented
- Write clear comments for exported functions

## Project Structure

```
.
├── main.go              # Main server implementation
├── main_test.go         # Tests
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── Makefile             # Build and development tasks
├── .gitignore           # Git ignore patterns
├── .golangci.yml        # Linter configuration
├── .goreleaser.yml      # Release configuration
└── .github/
    └── workflows/       # GitHub Actions workflows
```

## Release Process

Releases are automated through GitHub Actions:

1. Maintainers create a new tag:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

2. GitHub Actions automatically:
   - Builds binaries for all platforms
   - Creates a GitHub release
   - Uploads release artifacts

## Questions or Issues?

- Open an issue on GitHub for bugs or feature requests
- Check existing issues before creating a new one
- Provide as much context as possible when reporting issues

## License

By contributing to this project, you agree that your contributions will be licensed under the MIT License.
