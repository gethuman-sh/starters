# human starters

Starter project templates for the [human](https://github.com/gethuman-sh/human)
Start Project wizard.

## Layout contract

One template per `<type>/<language>/` subdirectory:

```
web/go/     Web Project, Go
```

The `human` desktop app downloads this repository as a tarball from the `main`
branch and extracts exactly one template subdirectory into the user's project
directory. Because of that:

- Keep every template self-contained inside its own subdirectory.
- Keep templates small — a handful of files that run out of the box.
- Do not rename the repository or the `main` branch without updating the
  `internal/starter` constants in the `human` CLI.

Adding a template: create the `<type>/<language>/` directory here, then
register it in `internal/starter/templates.go` in the `human` repository.
