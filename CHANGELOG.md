## [2.1.1](https://github.com/fabrizioe/code-conventions/compare/v2.1.0...v2.1.1) (2025-10-08)


### Bug Fixes

* **workflow:** run docker-release after semantic-release on main push\n\nChanged docker-release job to run after the release job completes on main branch push events, instead of requiring a separate release event trigger. This allows it to build and tag Docker images immediately after semantic-release creates a new version tag. ([7243dd6](https://github.com/fabrizioe/code-conventions/commit/7243dd68ee4a0e08d9d8cb1139159751b685e665))

# [2.1.0](https://github.com/fabrizioe/code-conventions/compare/v2.0.1...v2.1.0) (2025-10-08)


### Features

* add test file for docker release tagging verification\n\nThis file documents the expected behavior of the docker-release workflow job that tags images with semantic-release version tags. ([2885f89](https://github.com/fabrizioe/code-conventions/commit/2885f89c2ed48d9fcb8b1f5eec5533733932c518))

## [2.0.1](https://github.com/fabrizioe/code-conventions/compare/v2.0.0...v2.0.1) (2025-10-08)


### Bug Fixes

* **workflow:** tag Docker images with semantic-release release tag\n\nDocker images are now tagged and pushed using the semantic-release version (e.g., 1.0.0) on release events. This ensures proper versioning and traceability for published images. ([f7e1f5d](https://github.com/fabrizioe/code-conventions/commit/f7e1f5d9eec61237e30f6a03d0510b1d8d835951))

# [2.0.0](https://github.com/fabrizioe/code-conventions/compare/v1.0.0...v2.0.0) (2025-10-08)


* feat!: remove deprecated API endpoint ([a32b533](https://github.com/fabrizioe/code-conventions/commit/a32b533182a2b807f5c52ce2edde8a2c1ffef47b))


### BREAKING CHANGES

* The /api/v1/old-endpoint has been removed. Use /api/v2/new-endpoint instead.

# 1.0.0 (2025-10-08)


### Bug Fixes

* release job runs for any event on main branch ([84c7c82](https://github.com/fabrizioe/code-conventions/commit/84c7c8292f1be8128f75963620a28971a9154ef6))
* release job runs for any event on main branch ([5a8c8f8](https://github.com/fabrizioe/code-conventions/commit/5a8c8f8c78a1e748697a72d2f3a4009daf9aa8b5))
* run all jobs after PR merge to main ([b8526e5](https://github.com/fabrizioe/code-conventions/commit/b8526e57adccfbea303a6c0d69db9a7e602f3701))
* run all jobs after PR merge to main ([68613f5](https://github.com/fabrizioe/code-conventions/commit/68613f5fef3856a809ac3e11155276b2f2f22132))
* sem ver test ([23a32b7](https://github.com/fabrizioe/code-conventions/commit/23a32b7df57d978b64fa6a78b5ade7e90f78df30))
* test semver ([#5](https://github.com/fabrizioe/code-conventions/issues/5)) ([70d4767](https://github.com/fabrizioe/code-conventions/commit/70d476732524207a34c3e9d91ddea0718605d63d))


### Features

* add comprehensive code conventions and GitHub workflows ([5090a7e](https://github.com/fabrizioe/code-conventions/commit/5090a7e14a8fb8d23f405cfab84eb783a339ceff))
* Add HelloWorld service with complete CI/CD integration ([#1](https://github.com/fabrizioe/code-conventions/issues/1)) ([0925c70](https://github.com/fabrizioe/code-conventions/commit/0925c705820277340da44d84a55e366ef341ee17))
