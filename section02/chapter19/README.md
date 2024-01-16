get test cover

```bash
    go test ./utils/ -coverprofile=coverage.out && go tool cover -html=coverage.out
```