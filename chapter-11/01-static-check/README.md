If you had to pick one third-party scanner, use staticcheck. It is supported by many companies that are active in the Go community, includes more than 150 code-quality checks, and tries to produce few to no false positives.

It is installed via:

```
go install
honnef.co/go/tools/cmd/staticcheck@latest
```

Invoke it with staticcheck `./...` to examine your module.
