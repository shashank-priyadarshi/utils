# TODO

- CI

    - Add git hooks for local checks: DONE
    - Add go-releaser: DONE
    - Add change log using git cliff: IN PROGRESS
    - Add commitlint
    - Add golangci-lint
    - Add air for integration tests with testcontainers
    - Add unit, fuzz and mutation testing setup
    - Setup godoc for package level docs

- CD

- Documentation

    - Add generic pool docs
    - Add initial docs for all package with godoc compliance
    - Update workflow of worker
    - Update design of pubsub
    - Add missing package references & documentation to README

- DDD

    - Update packages for DDD
        - Database
        - Network
        - Pubsub
        - Worker
        - Security
    - Remove unnecessary abstraction using interfaces

- Pool

    - Add generic basic pool impl

- Data

    - Add interfaces: DONE
    - Replace placeholder code of Data interface implementation with actual implementation for JWT: DONE

- Database

    - Replace placeholder code of Database CRUD operations with actual implementation for MongoDB, Redis and SQL: DONE
    - Replace placeholder code of Operations interface implementation with actual implementation for MongoDB, Redis and
      SQL: DONE
    - Write Unit Tests for MongoDB, Redis and SQL: HOLD
    - Replace placeholder code of Transactions interface implementation with actual implementation for MongoDB, Redis
      and SQL
    - Integration tests

- Logger

    - File logging
    - Log rotation
    - Log forwarding
    - Tinted logs
    - Integration tests

- Network

    - Add correct segregation of Protocols, REST and RPCs: HOLD
    - Add interfaces for REST: DONE
    - Replace placeholder code of Network & REST interface implementation with actual implementation for Echo: DONE
    - Integration tests

- Security

    - Add correct segregation of Basic and Bearer authentication, JWT, SAML, Cookie and Session based authentication,
      API keys and OAuth: IN PROGRESS
    - Add interfaces

- PubSub

    - Add interfaces

- Worker

    - Add interfaces: DONE
    - Add implementation: DONE
