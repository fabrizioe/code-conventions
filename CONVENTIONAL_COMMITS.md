# Conventional Commits Enforcement

This project enforces [Conventional Commits](https://www.conventionalcommits.org/) format for all commits to ensure consistent commit history and enable automated semantic versioning.

## 🔒 Enforcement Levels

### All Branches
- **Basic Format Validation**: Ensures commit messages follow `<type>: <description>` format
- **Type Validation**: Only allows approved commit types
- **Automated Feedback**: Provides helpful error messages with examples

### Protected Branches (`main`, `develop`)
- **Strict Format Validation**: Enforces additional rules for protected branches
- **Character Limits**: Description must be 1-72 characters
- **Case Sensitivity**: Type and scope must be lowercase
- **No Trailing Periods**: Descriptions cannot end with periods

## 📝 Commit Message Format

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

### Required Components

1. **Type**: Must be one of the allowed types (see below)
2. **Description**: Short summary of the change (1-72 characters)

### Optional Components

1. **Scope**: Context of the change (e.g., component, module, or feature name)
2. **Body**: Longer description with details
3. **Footer**: Additional metadata (breaking changes, issue references)

## 🏷️ Allowed Commit Types

| Type | Description | Version Impact | Examples |
|------|-------------|----------------|----------|
| `feat` | New feature | Minor | `feat(auth): add JWT authentication` |
| `fix` | Bug fix | Patch | `fix(api): handle null response gracefully` |
| `docs` | Documentation | None | `docs: update API documentation` |
| `style` | Code style/formatting | None | `style: fix indentation in config.go` |
| `refactor` | Code refactoring | Patch | `refactor(db): optimize query performance` |
| `perf` | Performance improvement | Patch | `perf(cache): implement Redis caching` |
| `test` | Tests | None | `test(auth): add unit tests for login` |
| `build` | Build system changes | Patch | `build: update Go version to 1.21` |
| `ci` | CI/CD changes | None | `ci: add security scanning to pipeline` |
| `chore` | Maintenance tasks | None | `chore: update dependencies` |
| `revert` | Revert previous commit | Varies | `revert: revert commit abc123` |

## 💥 Breaking Changes

For breaking changes, add `!` after the type or include `BREAKING CHANGE:` in the footer:

```bash
# Method 1: Exclamation mark
feat!: remove deprecated API endpoint

# Method 2: Footer (recommended for detailed explanation)
feat(api): update authentication system

BREAKING CHANGE: The authentication API has been updated to use OAuth 2.0.
The old token-based authentication is no longer supported.
Clients must migrate to the new OAuth flow.
```

## ✅ Valid Examples

### Basic Commits
```bash
feat: add user registration endpoint
fix: resolve database connection timeout
docs: update README with setup instructions
style: format code according to style guide
refactor: extract validation logic to separate module
perf: optimize database queries
test: add integration tests for user service
build: update Docker base image
ci: add automated testing workflow
chore: update npm dependencies
```

### With Scope
```bash
feat(auth): implement JWT token refresh
fix(ui): correct button alignment on mobile
docs(api): add authentication examples
refactor(db): simplify connection pooling
test(integration): add end-to-end user flow tests
```

### Breaking Changes
```bash
feat!: migrate to new database schema
fix!: correct API response format

feat(api): redesign user management endpoints

BREAKING CHANGE: User API endpoints now return different field names.
See migration guide for details.
```

### Multi-line with Body
```bash
feat(monitoring): add Prometheus metrics endpoint

This commit adds a new /metrics endpoint that exposes
application metrics in Prometheus format, including:
- HTTP request duration
- Database connection pool stats
- Custom business metrics

Closes #123
```

## ❌ Invalid Examples

```bash
# Missing type
Add user authentication

# Invalid type
feature: add new endpoint

# No description
feat:

# Description too long (over 72 characters)
feat: implement a very long and detailed description that exceeds the maximum allowed character limit for commit message descriptions

# Trailing period
feat: add user service.

# Wrong case
FEAT: add user service
Feat: add user service

# Invalid scope format
feat(User_Service): add authentication
```

## 🔧 Tools and Configuration

### Commitlint Configuration
The project uses `@commitlint/config-conventional` with custom rules in `commitlint.config.js`:

- **Type validation**: Only allows specified commit types
- **Case validation**: Enforces lowercase types and scopes
- **Length validation**: Header max 72 characters, body lines max 100 characters
- **Format validation**: Ensures proper conventional commit structure

### GitHub Actions Validation
The `.github/workflows/commit-lint.yml` workflow runs on every push and PR:

1. **Install commitlint**: Uses Node.js and npm to install validation tools
2. **Validate commits**: Checks all commits in push/PR against conventional format
3. **Provide feedback**: Shows helpful error messages with examples on failure
4. **Block invalid commits**: Prevents merging PRs with invalid commit messages

### Local Development Setup

Install commitlint locally for immediate feedback:

```bash
# Install commitlint globally
npm install -g @commitlint/cli @commitlint/config-conventional

# Validate the last commit
npx commitlint --from HEAD~1 --to HEAD --verbose

# Install git hook (optional)
npx husky add .husky/commit-msg 'npx --no -- commitlint --edit ${1}'
```

## 🚫 Bypass and Exceptions

### Emergency Hotfixes
For critical production issues, you may need to bypass validation temporarily:

```bash
# Use conventional format even for hotfixes
fix!: critical security vulnerability in auth system

BREAKING CHANGE: All users must re-authenticate due to security fix.
```

### Merge Commits
Merge commits from GitHub are automatically formatted and don't require manual formatting.

### Revert Commits
Use the `revert` type for reverting previous commits:

```bash
revert: revert "feat(auth): add JWT authentication"

This reverts commit abc123def456.
```

## 🔍 Troubleshooting

### Common Issues

1. **"Type must be lowercase"**
   ```bash
   # ❌ Wrong
   FEAT: add new feature
   
   # ✅ Correct
   feat: add new feature
   ```

2. **"Subject cannot be empty"**
   ```bash
   # ❌ Wrong
   feat:
   
   # ✅ Correct
   feat: add user authentication
   ```

3. **"Header too long"**
   ```bash
   # ❌ Wrong (over 72 characters)
   feat: implement comprehensive user authentication system with JWT tokens and refresh capabilities
   
   # ✅ Correct
   feat: implement JWT authentication system
   ```

4. **"Subject cannot end with period"**
   ```bash
   # ❌ Wrong
   feat: add user service.
   
   # ✅ Correct
   feat: add user service
   ```

### Fixing Invalid Commits

#### Amend Last Commit
```bash
# Fix the most recent commit message
git commit --amend -m "feat: corrected commit message"
```

#### Rebase Multiple Commits
```bash
# Fix multiple commits (last 3 in this example)
git rebase -i HEAD~3

# In the editor, change 'pick' to 'reword' for commits to fix
# Save and close, then edit each commit message
```

#### Squash and Rewrite
```bash
# Squash multiple commits into one with proper message
git reset --soft HEAD~3
git commit -m "feat(auth): implement complete authentication system"
```

## 📚 Additional Resources

- [Conventional Commits Specification](https://www.conventionalcommits.org/)
- [Semantic Versioning](https://semver.org/)
- [Commitlint Documentation](https://commitlint.js.org/)
- [Angular Commit Guidelines](https://github.com/angular/angular/blob/main/CONTRIBUTING.md#commit)

## 🎯 Quick Reference

### Most Common Types
```bash
feat: new features
fix: bug fixes
docs: documentation
test: testing
refactor: code improvements
chore: maintenance
```

### Quick Templates
```bash
# Feature
feat(scope): add [what]

# Bug fix
fix(scope): resolve [issue]

# Documentation
docs: update [what] documentation

# Breaking change
feat!: [change description]
BREAKING CHANGE: [detailed explanation]
```