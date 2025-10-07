# GitHub Actions CI/CD Pipeline

This repository implements a comprehensive CI/CD pipeline following trunk-based development practices with semantic versioning and automated release management.

## Pipeline Overview

### 🔄 Continuous Integration

The CI pipeline triggers on every push to any branch and includes:

- **Compile & Test**: Go compilation, unit tests, and code coverage
- **Code Quality**: Static analysis with `go vet`, `staticcheck`, and format checking
- **Security Scanning**: Vulnerability scanning with Gosec and GitLeaks
- **Container Build**: Multi-arch Docker images with Trivy security scanning
- **Integration Testing**: Full stack testing with Docker Compose

### 🚀 Continuous Deployment

The CD pipeline triggers on main branch and includes:

- **Semantic Release**: Automated versioning based on conventional commits
- **Release Notes**: Auto-generated changelogs
- **Container Registry**: Automated image publishing to GitHub Container Registry
- **GitHub Releases**: Binary artifacts and release notes

## Branch Protection

### Protected Branches

- **main**: Production branch with strict protection rules
- **develop**: Integration branch (if used) with standard protection

### Protection Rules

- ✅ Require pull request reviews (1 approver minimum)
- ✅ Require status checks to pass before merging
- ✅ Require conversation resolution
- ✅ Dismiss stale reviews when new commits are pushed
- ✅ Require review from code owners
- ❌ Allow force pushes (disabled)
- ❌ Allow deletions (disabled)

### Required Status Checks

All the following checks must pass before merging to protected branches:

- `provisioning-service-ci` - Go compilation, tests, and quality checks
- `docker-build` - Container build and security scan
- `integration-tests` - Full stack integration testing
- `security-scan` - Security vulnerability scanning
- `quality-gate` - Overall quality gate for protected branch merges

## Semantic Versioning

This project uses [Semantic Versioning](https://semver.org/) with automated release management via [semantic-release](https://semantic-release.gitbook.io/).

### Commit Message Format

We follow [Conventional Commits](https://www.conventionalcommits.org/) specification:

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Commit Types

| Type | Description | Version Impact |
|------|-------------|----------------|
| `feat` | New feature | Minor version bump |
| `fix` | Bug fix | Patch version bump |
| `perf` | Performance improvement | Patch version bump |
| `refactor` | Code refactoring | Patch version bump |
| `build` | Build system changes | Patch version bump |
| `docs` | Documentation changes | No version bump |
| `style` | Code style changes | No version bump |
| `test` | Test changes | No version bump |
| `ci` | CI/CD changes | No version bump |
| `chore` | Maintenance tasks | No version bump |

### Breaking Changes

For breaking changes, add `BREAKING CHANGE:` in the commit footer or use `!` after the type:

```
feat!: remove deprecated API endpoint

BREAKING CHANGE: The /api/v1/old-endpoint has been removed. Use /api/v2/new-endpoint instead.
```

This will trigger a **major** version bump.

### Examples

```bash
# Feature (minor version bump)
git commit -m "feat(auth): add JWT token refresh endpoint"

# Bug fix (patch version bump)  
git commit -m "fix(config): handle missing environment variables gracefully"

# Breaking change (major version bump)
git commit -m "feat!: migrate to new authentication API"

# Documentation (no version bump)
git commit -m "docs: update API documentation for device registration"
```

## Workflow Files

### `.github/workflows/ci-cd.yml`

Main CI/CD pipeline with the following jobs:

1. **detect-changes**: Detects which parts of the codebase changed
2. **provisioning-service-ci**: Go service compilation and testing
3. **docker-build**: Container build with security scanning
4. **integration-tests**: Full stack testing with Docker Compose
5. **security-scan**: Security vulnerability analysis
6. **quality-gate**: Quality gate for protected branches
7. **release**: Semantic release (main branch only)
8. **notify**: Build status notifications

### `.github/workflows/branch-protection.yml`

Automated branch protection rule configuration. Run manually or on workflow file changes.

## Container Registry

Docker images are automatically built and published to GitHub Container Registry:

- **Registry**: `ghcr.io`
- **Image**: `ghcr.io/{owner}/{repo}/provisioning-service`
- **Tags**: 
  - `latest` (main branch)
  - `{branch-name}` (feature branches)
  - `{branch-name}-{sha}` (commit-specific)
  - `pr-{number}` (pull requests)

### Multi-Architecture Support

Images are built for multiple architectures:
- `linux/amd64`
- `linux/arm64`

## Security

### Vulnerability Scanning

- **Code**: Gosec static analysis
- **Secrets**: GitLeaks secret scanning  
- **Containers**: Trivy vulnerability scanning
- **Dependencies**: GitHub Security Advisories

### Security Reports

Scan results are uploaded to GitHub Security tab for tracking and remediation.

## Local Development

### Prerequisites

```bash
# Install required tools
go install honnef.co/go/tools/cmd/staticcheck@latest
```

### Running CI Checks Locally

```bash
# Navigate to service directory
cd AuthNZ/provisioning-service

# Format check
gofmt -s -l .

# Linting
go vet ./...
staticcheck ./...

# Tests
go test -v -race -coverprofile=coverage.out ./...

# Build
go build -v ./cmd/main.go
```

### Integration Testing

```bash
# Start the stack
cd AuthNZ
docker-compose up -d

# Setup services  
./scripts/setup-keycloak.sh
./scripts/setup-hono.sh

# Run tests
./scripts/test-mqtt.sh
./scripts/test-http.sh

# Cleanup
docker-compose down -v
```

## Monitoring and Alerts

### Build Status

Monitor build status through:
- GitHub Actions tab
- Pull request status checks
- Commit status indicators

### Quality Metrics

Track code quality through:
- Test coverage reports (Codecov)
- Security scan results (GitHub Security)
- Container vulnerability reports (Trivy)

## Troubleshooting

### Common Issues

1. **Tests failing locally but passing in CI**
   - Ensure consistent Go version (1.21)
   - Check for race conditions in tests
   - Verify all dependencies are available

2. **Docker build failing**
   - Check Dockerfile syntax
   - Verify all COPY paths exist
   - Ensure base image is accessible

3. **Integration tests failing**
   - Verify Docker Compose services start successfully
   - Check service health endpoints
   - Review service logs for errors

4. **Semantic release not creating versions**
   - Verify commit message format follows conventional commits
   - Check that commits include version-triggering types
   - Ensure GITHUB_TOKEN has necessary permissions

### Debug Commands

```bash
# Check workflow syntax
act --list

# Test Docker build locally
docker build -t test AuthNZ/provisioning-service/

# Validate semantic-release config
npx semantic-release --dry-run

# Check commit message format
npx commitlint --from HEAD~1 --to HEAD --verbose
```

## Contributing

1. Create feature branch: `git checkout -b {task-id}-{description}`
2. Make changes following coding standards
3. Write/update tests for new functionality
4. Use conventional commit messages
5. Create pull request to main/develop
6. Ensure all CI checks pass
7. Get code review approval
8. Merge after all requirements met

## Release Process

Releases are fully automated:

1. **Commit** to main branch with conventional commit message
2. **CI Pipeline** validates all checks pass
3. **Semantic Release** analyzes commits and determines version
4. **GitHub Release** created with auto-generated notes
5. **Container Images** tagged and published
6. **Changelog** updated in repository

No manual intervention required for releases!