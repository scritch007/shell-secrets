This repo aims at creating a small library to help securing secrets in a given shell.
This is inspired by the ssh-agent mechanism but simplified.

User needs to run the setup command which will generate an env variable that holds a symmetric key.

The user can than add/get/remove an entry.

linux
=====

```
eval $(shell-scret setup)
shell-secret add toto '{"key": "value"}'
shell-secret get toto
shell-secret delete toto
```