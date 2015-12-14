# deis-integration

This repository contains Deis v2 integration tests. Each integration test is a Yaml file inside a specified directory. Each file has the following structure:

```yaml
version: 1
name: test-name
description: this is a test that does something interesting
commands:
  # these are a series of shell commands
  # the expectations below must hold after each command
  - shell: command 1
  - shell: some other command
```
