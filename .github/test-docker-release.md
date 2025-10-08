# Test File for Docker Release Tagging

This file was created to test the workflow's ability to tag Docker images with semantic-release version tags.

When a new release is created, the `docker-release` job should:

1. Extract the release tag (e.g., v2.0.2 -> 2.0.2)
2. Build the Docker image with that tag
3. Push to ghcr.io with the semantic version tag

This test ensures the workflow properly handles release events and tags images accordingly.
