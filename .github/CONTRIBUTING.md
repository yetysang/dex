# Contributing to Dex

Thank you for your interest in contributing to Dex! This document outlines the process for contributing to this project.

## Code of Conduct

This project adheres to a [Code of Conduct](.github/CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Developer Certificate of Origin

All contributions must be signed off in accordance with the [Developer Certificate of Origin (DCO)](DCO). This means adding a `Signed-off-by` line to your commit messages:

```
Signed-off-by: Your Name <your.email@example.com>
```

You can do this automatically by using `git commit -s`.

## Getting Started

### Prerequisites

- Go 1.21 or later
- Docker (for integration tests)
- Make

### Setting Up Your Development Environment

1. Fork the repository on GitHub.
2. Clone your fork locally:
   ```bash
   git clone https://github.com/<your-username>/dex.git
   cd dex
   ```
3. Add the upstream remote:
   ```bash
   git remote add upstream https://github.com/dexidp/dex.git
   ```
4. If you use `direnv`, allow the `.envrc` file:
   ```bash
   direnv allow
   ```

## Making Changes

1. Create a new branch for your changes:
   ```bash
   git checkout -b feat/my-new-feature
   ```
2. Make your changes, ensuring you follow the coding standards below.
3. Write or update tests for your changes.
4. Run the test suite to ensure nothing is broken:
   ```bash
   make test
   ```
5. Commit your changes with a descriptive message and DCO sign-off:
   ```bash
   git commit -s -m "feat: add my new feature"
   ```

## Coding Standards

- Follow standard Go conventions and idioms.
- Run `gofmt` or `goimports` before committing.
- Ensure `go vet` passes without errors.
- Keep functions focused and well-documented.
- Add comments for exported types and functions.

## Submitting a Pull Request

1. Push your branch to your fork:
   ```bash
   git push origin feat/my-new-feature
   ```
2. Open a Pull Request against the `main` branch of this repository.
3. Fill out the PR template completely.
4. Ensure all CI checks pass.
5. Address any review feedback promptly.

## Reporting Bugs

Please use the [bug report template](.github/ISSUE_TEMPLATE/bug_report.yaml) when filing issues.

## Questions?

Feel free to open a GitHub Discussion or reach out via the issue tracker.
