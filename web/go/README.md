# Go web project

A minimal Go web server built on [Echo v5](https://echo.labstack.com), created
by the human Start Project wizard.

## Run

```sh
make run
```

Then open http://localhost:8080. Set `PORT` to use a different port.

## Test

```sh
make test
```

## Checks

All tools are version-pinned in `go.mod` (via `tool` directives) and download
automatically on first use — no separate install step.

```sh
make lint     # go vet + staticcheck
make sec      # gosec + govulncheck
make secrets  # gitleaks secret scan
make check    # test + lint + sec + secrets
```

CI (`.github/workflows/ci.yml`) runs the same checks on every push and pull
request against `main`.
