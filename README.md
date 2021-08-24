# Eiffel Goer

[![Sandbox badge](https://img.shields.io/badge/Stage-Sandbox-yellow)](https://github.com/eiffel-community/community/blob/master/PROJECT_LIFECYCLE.md#stage-sandbox)

Eiffel Goer is a go implementation of the event repository API.

## Description

Eiffel Goer implements the event repository API and is intended as an open source alternative to the Eiffel Event Repository.

## Features

- Simple implementation of the Eiffel ER API.
- Event searching

## Installation

### Docker

    docker run -e CONNECTION_STRING=yourdb -e DATABASE_NAME=dbname -e API_PORT=8080 registry.nordix.org/eiffel/goer

### Running a development server locally for testing. Will restart on code changes.

    make start

### Building a local executable

    make build

### Running tests

    make test

## Contribute

- Issue Tracker: https://github.com/eiffel-community/eiffel-goer/issues
- Source Code: https://github.com/eiffel-community/eiffel-goer

## Support

If you are having issues, please let us know.

There is a mailing list at: eiffel-goer-maintainers@google-groups.com
or just write an Issue.
