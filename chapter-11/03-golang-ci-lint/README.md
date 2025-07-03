Finally, if you’d rather take the buffet approach to tool selection, there’s `golangci-lint`. It is designed to make it as efficient as possible to configure and run over 50 code-quality tools, including go vet, staticcheck, and revive.

While you can use `go install golangci-lint`, it is recommended that you download a binary version instead. Follow the installation instructions on the website. 
Once it is installed, you run golangci-lint:
```
$ golangci-lint run
```