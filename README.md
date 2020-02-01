# Scripts To Rule Them All Completions

Script to rule them all from Github contains a number of scripts like `deploy`,
`console` and these scripts all accept a environment variable.

This projects enables completions for these scripts.

Currently it supports completions for:

- `script/deploy`

And the following environments:

- `development`
- `staging`
- `production`
- `approval`
- `sandbox`

These completions are made available using the [posener/complete](https://github.com/posener/complete)

## Install

```bash
go get -u github.com/niels-s/scripts-to-rule-them-all-complete
COMP_INSTALL=1 scripts-to-rule-them-all-complete
```

## Uninstall

```bash
COMP_UNINSTALL=1 scripts-to-rule-them-all-complete
```
