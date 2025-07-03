Another good linting option is revive. It is based on golint, a tool that used to be maintained by the Go team. 

Install revive with the command: 

```
go install github.com/mgechev/revive@latest
```

With a configuration file, you can turn on many more rules. For example, to enable a check for shadowing of universe block identifiers, create a file named `built_in.toml` with the following contents:
```
[rule.redefines-builtin-id]
```