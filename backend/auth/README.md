# authentication

[![Keep a Changelog v1.0.0 badge][changelog-badge]][changelog] [![Version 1.0.0 Badge][version-badge]][changelog]

This microservice is used to authenticate users and renew sessions.

## Actions
* `login` creates a new session based on the `username` and `password` provided by the requester.
* `refresh-token` renovate the current session if the `refresh_token` has not expired yet and return a new valid session.

## Dependencies
* Golang ([see version][golang-version] recommended)
* Docker ([see version][docker-version] recommended)

## Contribute

Please do contribute! Issues and pull requests are welcome.
Thank you for your help improving software one changelog at a time!


[golang-version]: .golang-version
[docker-version]: .docker-version
[version-badge]: https://img.shields.io/badge/version-1.0.0-blue.svg
[changelog]: ./CHANGELOG.md
[changelog-badge]: https://img.shields.io/badge/changelog-Keep%20a%20Changelog%20v1.0.0-%23E05735
