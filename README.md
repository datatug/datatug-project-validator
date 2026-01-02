# datatug-project-validator
Validate DataTug projects locally or in CI pipelines

## Run as GitHub action

In your repository create file: `/.github/workflows/datatug.yml`:
```yaml
name: DataTug Project Validator

on:
  pull_request:
  push:

jobs:
  validate:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: datatug/datatug-project-validator@main
```
