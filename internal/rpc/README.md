# rpc

This package is a `JSON-RPC 2.0` client derived from
[github.com/ethereum/go-ethereum/rpc](https://github.com/ethereum/go-ethereum/tree/master/rpc)
and stripped down to only the HTTP client.

Only HTTP (`http`, `https`) transport is supported.

This package maintains the original [GNU Lesser General Public License v3.0
(LGPL-3.0)](./LICENSE). Authors of the original package can be found in the root
of the that repository
([AUTHORS](https://github.com/ethereum/go-ethereum/blob/master/AUTHORS)).

This package was vendored instead of imported due to the size of the parent
package (~190MB), considering only one half (the client) of a sub-package would
be used.

Additionally, the presence of a cryptocurrency mining framework as a dependency
may give the false impression that the parent package
(`github.com/wranders/terraform-provider-xcpng`) is introducing
malware / unwanted applications into the end user's infrastructure.

This package is internal as it is only intended to be used by the parent package
(`github.com/wranders/terraform-provider-xcpng`). Use of this package outside of
this context is unsupported.

## Usage

```go
package main

import (
    "crpyto/tls"
    "fmt"
    "net/http"

    "github.com/wranders/terraform-provider-xcpng/internal/rpc"
)

const endpoint = "https://endpoint"

func main() {
    httpClient := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            }
        }
    }
    client, err := rpc.DialHTTPWithClient(endpoint, httpClient)
    if err != nil {
        panic(err)
    }

    var result string
    if err := client.Call(&result, "Resource.method"); err != nil {
        panic(err)
    }
    fmt.Println(result)
}
```
