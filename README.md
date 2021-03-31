# Overview

It's an API Skeleton project based on Echo framework.

## Directories

1. **/config** has structures which contains service config.

2. **/models** includes structures describing data models.

3. **/requests** has structures describing the parameters of incoming requests, and validator.

4. **/responses** includes structures describing the parameters of outgoing response.

5. **/server/handlers** contains request handlers.

6. **/server/routes** has a file for configuring routes.

## Code quality

For control code quality we are use [golangci-lint](https://github.com/golangci/golangci-lint). Golangci-lint is a
linters aggregator.

Why we use linters? Linters help us:

1. Finding critical bugs
2. Finding bugs before they go live
3. Finding performance errors
4. To speed up the code review, because reviewers do not spend time searching for syntax errors and searching for
   violations of generally accepted code style
5. The quality of the code is guaranteed at a fairly high level.

### How to use

Linter tool wrapped to docker-compose and first of all need to build container with linters

- `make lint-build`

Next you need to run linter to check bugs ant errors

- `make lint-check` - it will log to console what bugs and errors linters found

Finally, you need to fix all problems manually or using autofixing (if it's supported by the linter)

- `make lint-fix`

## Libraries

Jwt - https://github.com/dgrijalva/jwt-go

Swagger - https://github.com/swaggo/echo-swagger

gocloak - https://github.com/Nerzal/gocloak

gocloak-echo - https://github.com/Nerzal/gocloak-echo

## License

MIT
