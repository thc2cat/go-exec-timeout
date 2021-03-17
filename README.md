# go-exec-timeout

You have a cli command that you would like to have a timeout when executing ?

go-exec-timeout will shadow the command and execute it with timeout.

```shell
# go build  // Build binary
# mv mycmd mycmd.org // rename original command with .org suffix
# mv go-exec-timeout mycmd // move go-exec-timeout to original place and name
# /path/to/mycmd mycmdarg1 mycmdarg... // execute as normal
```

Default timeout is 90s but ca be changed via TIMEOUT environment variable.

Log of timeouts will be sent to daemon.info syslog facility
