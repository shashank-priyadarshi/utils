# Changelog

All notable changes to this project will be documented in this file.

## [unreleased]

### 🚀 Features

- Add implementation of Data interface for JWT
- Implement http server using echo and refactor network
- Add echo server startup with or without TLS
- Add initial implementation of a worker pool
- Add scaling based on wait time, job statuses and job timeout & cancellation
- Add logger to Pool class and functional options
- Add idx config
- Add impl for logger using logrus, zap & zerolog
- *(repo)* Update module name to vanity go.ssnk.in/utils

### 🐛 Bug Fixes

- Update prepare-hook to install commitlint conventional config and add commitlint.config.js
- Segregate invalid parameter and type errors
- Update mongodb impl with proper invalid parameter and type errors
- Update maxExecutionTime & waitTime for each work, and fix status update bugs for timed out jobs

### ⚙️ Miscellaneous Tasks

- Update TODO to include CI
- Add goreleaser config
- Remove logs from commit and add to gitignore
- Update TODO
- Add progress for redundancy and free up head to new dev
- Upgrade go version and deps to latest
- Vendor deps
- *(repo)* Revendor deps, add updated changelog & doc
- *(repo)* Remove external deps & use net/http & golang.org/x/net/websocket

### 🚜 Refactor

- Structure integration tests
- Update network ports and add empty files for sessions, datalink, tcp, udp, sctp
- Remove log DI from database package and add ut for sql
- Remove log DI from data package and rename constructors
- Remove redundant usages of type keyword in database package
- Update database constructor to accept referenced config model
- Cleanup models and constants, fiddling with ports for solutions
- Cleanup ports except HTTP REST using echo implementation
- Move integrationtests to tests, add other tests' structure
- *(logger)* Refactor log config, ports & constructor func names
- *(logger)* Update commitlint scope enum and package import
- *(test)* Rename tests dir & setup unified suite with more test types
- *(setup)* Add commit check and releaser scripts, tools & config
- *(secretxform)* Remove redundant impl for JWT & SAML from data pkg
- *(database)* Update mock import paths to updated path from test
- *(setup,repo)* Refactor test & mocks structure and update tool configs
- *(repo)* Restart new head with minimal packages for brevity

### 🎨 Styling

- Fix formatting and import order

### 🧪 Testing

- *(database)* Update module dependencies and add mongo db and sql test file
- Add unit tests for mongodb
- Add integration tests for worker pool package
- *(mocks)* Move mocks to test dir for easy management
- Add algo test type with algo module to test suite

### 🚜 Refactor

- Setup errcheck, fieldaligment and staticcheck, patch commitlint issues, patch slog.go
- Update commitlint config
- Update gitignore & add scripts for go dl & envs in web editors
- Add workflows
- Cleanup & retain sonarqube and codecov workflows
- Fix workflow yaml, config & env var names
- Fix workflow yaml, config & env var names
- Update workflow trigger for pull request on main branch
- Update workflow trigger for push to main & dev
- Update default branch for sonarqube & increased verbosity
- Remove run key from dependency review workflow
- Add workflow run on pull request to branch main
- Update license list in dependency review job
- Use active action-semantic-pull-request action
- Downgrade amann/action-semantic-pull-request from main v5
- Continue-on-error true for semantic-pull-request to skip not found err
- Resolve sonar project key & org detection with quote error
- Resolve sonarqube external issues not found using golangci_lint output
- Resolve ci yaml formatting error & segregate configs
- Standardize yml ext & golangci-lint action with bin manual install
- Resolve golangci-lint action path error & cache actions
- Remove usage of go123 with golangcilint version 122
- Bump golangci-lint version to latest for go123 support
- Add action to install go123
- Update golangci regex yml to go convention
- Update golangci regex yml to go convention
- Remove golangci yaml exclude files regex
- Continue on error if golangci lint action fails
- Add external issues json for sonarqube
- Add checkstyle golangci-lint output integration with sonarqube
- Update golangci-run command for actions test
- Update golangci-run command for actions test

### TODO

- Updated TODO

### WIP

- Added interfaces for databases and implemented logger using slog
- Added interfaces for Data and Network
- Added worker and pubsub
- Added implementation of Create, Update & Query for MongoDB
- Minor bug fix in rdbms.go
- Minor bug fix in rdbms.go
- Added module and package level constants & usages, and implemented Create for Redis
- Implemented CRUD for Redis and refactored methods for MongoDB operations for standardized errors and no naked returns
- Completed implemented implementation of MongoDB delete operation
- Updated TODO to reflect latest agenda and progress
- Refactored constructor names for meaningfulness and migrated from snake case to small case import aliases in database
- Corrected function argument aliases from i or params to args and added expectation in sql package in database
- Completed implementation of CRUD for sql in database
- Updated TODO
- Setup miniredis, mockery & testify, generated database port mocks, and happy path unit tests for redis operations in database
- Added commit-msg and pre-commit git hooks alongside prepare hook script

<!-- generated by git-cliff -->
