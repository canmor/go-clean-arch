# go-clean-arch

An analyzer for Clean Arch dependency rules for Go.

## Guides

### Check with `go vet`

1. Install binary

    ```shell
    go install github.com/canmor/go-clean-arch/cmd/go-clean-arch@latest
    ```

2. Run with `go vet` to check your repo

    ```shell
    go vet -vettool=$GOPATH/bin/go-clean-arch ./...
    ```
