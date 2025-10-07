# GitHub Actions CI/CD Pipeline

This repository implements a comprehensive CI/CD pipeline following trunk-based development practices with semantic versioning and automated release management. Semantic-release automatically tags new releases on the main branch after pull requests are merged. The develop branch is used for feature integration and does not trigger releases.

## Pipeline Overview

### 🔄 Continuous Integration

The CI pipeline triggers on every push to any branch and includes:

- **Compile & Test**: Go compilation, unit tests, and code coverage
- **Code Quality**: Static analysis with `go vet`, `staticcheck`, and format checking
- **Security Scanning**: Vulnerability scanning with Gosec and GitLeaks
- **Container Build**: Multi-arch Docker images with Trivy security scanning
- **Configurable Services**: Easily adaptable for different services via environment variables

### 🚀 Continuous Deployment

The CD pipeline triggers on main branch and includes:

- **Semantic Release**: Automated versioning and tagging based on conventional commits (runs only on main branch after PR merge)
- **Release Notes**: Auto-generated changelogs
- **Container Registry**: Automated image publishing to GitHub Container Registry (requires PAT with write:packages scope)
- **GitHub Releases**: Binary artifacts and release notes

## Branch Protection

### Protected Branches

- **main**: Production branch for official releases. Semantic-release tags and publishes releases only from main after PRs are merged.
- **develop**: Integration branch for new features. No releases or tags are created from develop.

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

- `service-ci` - Go compilation, tests, and quality checks
- `docker-build` - Container build and security scan
- `security-scan` - Security vulnerability scanning
- `quality-gate` - Overall quality gate for protected branch merges

## Configuration

### Environment Variables

The workflow is configurable via environment variables defined in the `env` section:

| Variable | Description | Default Value |
|----------|-------------|---------------|
| `SERVICE_NAME` | Name of the service for reports and coverage | `my-service` |
| `DOCKER_IMAGE_NAME` | Docker image name for container builds | `my-service` |
| `WORKING_DIRECTORY` | Working directory containing the service code | `./src` |
| `GO_VERSION` | Go version for compilation and testing | `1.21` |
| `REGISTRY` | Container registry for image publishing | `ghcr.io` |

### Customization

To adapt this workflow for your service:

1. Update the environment variables in `.github/workflows/ci-cd.yml`
2. Ensure your service code is in the specified `WORKING_DIRECTORY`
3. Update the Docker image name to match your service
4. Customize the service name for reporting

## Semantic Versioning

This project uses [Semantic Versioning](https://semver.org/) with automated release management via [semantic-release](https://semantic-release.gitbook.io/).

### Commit Message Format

We follow [Conventional Commits](https://www.conventionalcommits.org/) specification:

```text
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

```bash
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
2. **service-ci**: Go service compilation and testing
3. **docker-build**: Container build with security scanning
4. **security-scan**: Security vulnerability analysis
5. **quality-gate**: Quality gate for protected branches
6. **release**: Semantic release (main branch only)
7. **notify**: Build status notifications

### `.github/workflows/branch-protection.yml`

Automated branch protection rule configuration. Run manually or on workflow file changes.

## Container Registry

Docker images are automatically built and published to GitHub Container Registry:

- **Registry**: `ghcr.io`
- **Image**: `ghcr.io/{owner}/{repo}/{DOCKER_IMAGE_NAME}`
- **Tags**:
   - `latest` (main branch)
   - `{branch-name}` (feature branches)
   - `{branch-name}-{sha}` (commit-specific)
   - `pr-{number}` (pull requests)

### Authentication for Docker Push

To push images to GitHub Container Registry, you must use a Personal Access Token (PAT) with `write:packages` and `read:packages` scopes. Add this token as a secret (e.g., `GHCR_PAT`) in your repository or organization settings. The workflow uses this secret for Docker login:

```yaml
- name: Log in to Container Registry
   uses: docker/login-action@v3
   with:
      registry: ghcr.io
      username: ${{ github.actor }}
      password: ${{ secrets.GHCR_PAT }}
```

### Image Naming

The Docker image name is configurable via the `DOCKER_IMAGE_NAME` environment variable in the workflow. Update this variable to match your service name.

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
   - Ensure consistent Go version (check `GO_VERSION` env var)
   - Check for race conditions in tests
   - Verify all dependencies are available

2. **Docker build failing**
   - Check Dockerfile syntax in your `WORKING_DIRECTORY`
   - Verify all COPY paths exist relative to the working directory
   - Ensure base image is accessible

3. **Service detection not working**
   - Verify your code is in the specified `WORKING_DIRECTORY`
   - Check that the path filter in `detect-changes` matches your structure
   - Ensure the working directory contains Go modules (`go.mod`)

4. **Semantic release not creating versions**
   - Verify commit message format follows conventional commits
   - Check that commits include version-triggering types
   - Ensure GITHUB_TOKEN has necessary permissions

### Debug Commands

```bash
# Check workflow syntax
act --list

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

Releases are fully automated and only occur on the main branch:

1. **Merge Pull Request** into main branch (from develop or feature branch)
2. **CI Pipeline** validates all checks pass
3. **Semantic Release** analyzes commits and determines version
4. **GitHub Release** created with auto-generated notes and semver tag
5. **Container Images** tagged and published
6. **Changelog** updated in repository

No manual intervention required for releases! Develop branch is used for feature integration and does not trigger releases or tags.
