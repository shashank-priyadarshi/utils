## Changelog
* ac76b23 - Updated go modules dependencies - Reformatted TODO based on markdown linter - Return type of database ports updated to *model.Response and corresponding adapters were updated - Added implementation of Server start and shutdown for Echo
* f83c124 Initial commit
* 629cb13 TODO: Updated TODO
* 600e411 WIP: Added commit-msg and pre-commit git hooks alongside prepare hook script
* bf9fd76 WIP: Added implementation of Create, Update & Query for MongoDB
* 1f662c6 WIP: Added interfaces for Data and Network
* 5f4fffc WIP: Added interfaces for databases and implemented logger using slog
* 77c7e8a WIP: Added module and package level constants & usages, and implemented Create for Redis
* 6bc67b3 WIP: Added worker and pubsub
* c5a0420 WIP: Completed implementation of CRUD for sql in database
* e5d4234 WIP: Completed implemented implementation of MongoDB delete operation
* 002e172 WIP: Corrected function argument aliases from i or params to args and added expectation in sql package in database
* 6f2fbe6 WIP: Implemented CRUD for Redis and refactored methods for MongoDB operations for standardized errors and no naked returns
* 02cf32e WIP: Minor bug fix in rdbms.go
* c5fa5a3 WIP: Minor bug fix in rdbms.go
* 19374e8 WIP: Refactored constructor names for meaningfulness and migrated from snake case to small case import aliases in database
* 2cde675 WIP: Setup miniredis, mockery & testify, generated database port mocks, and happy path unit tests for redis operations in database
* fe59ae4 WIP: Updated TODO
* 22b746f WIP: Updated TODO to reflect latest agenda and progress
* 8cdd1df chore: add goreleaser config
* bf487da chore: update TODO to include CI
* c9787d1 fix: update prepare-hook to install commitlint conventional config and add commitlint.config.js
* a0bfb2a test(database): update module dependencies and add mongo db and sql test file
