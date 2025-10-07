## Development Conventions

### Project Structure
- Each service has its own `Spec.md` with domain entities and responsibilities
- Implementation code lives in separate directories from specifications
- Docker Compose used for local development stack
- Go services follow standard project layout: `cmd/`, `internal/`, `config/`

### Git Workflow
- Use trunk-based development with feature branches
- Branch naming: `<task-id>-<short-description>`
- Conventional commit messages required
- All tests must pass before merge to main
- no force pushes to protected branches
- run git commands without confirmation prompts

### Configuration Management
- YAML-based configuration files in `config/` directories
- Environment-specific overrides supported
- Structured logging with configurable levels (debug/info/error)
- JSON formatter for production environments
