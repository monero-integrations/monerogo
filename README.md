# monerogo

# Description
Monero JSON RPC Client

# Installation
    go get github.com/MarinX/monerogo

# HowTo

### Daemon Client
To use *DaemonClient* you need to start your daemon `./monerod` which will listen on `localhost:18081`
```go
package main

import (
	"fmt"

	"github.com/MarinX/monerogo"
)

func main() {
	// creates new daemon client
	// the endpoint is http://127.0.0.1:18081/json_rpc
	client := monerogo.NewDaemonClient("http://127.0.0.1:18081/json_rpc")

	blockCount, err := client.GetBlockCount()
	if err != nil {
		fmt.Println(err)
		return
	}

	// always check the status of the response from RPC
	if blockCount.Status != "OK" {
		fmt.Println("RPC method is not OK, got", blockCount.Status)
		return
	}

	fmt.Println("Count:", blockCount.Count)
	
	// you can check more examples in monerogo_test.go file
}
```

### Wallet Client
//TODO

# Roadmap
- [x] Create Daemon RPC client
- [ ] Create Wallet RPC client

# License
This library is under the MIT License

# Author
Marin Basic 

## Donate XMR
`45gSodJY5hnAL441jiYg2C72LrPoekuPNgLgWKAbW66Bdt9fyC4RcdH8A3qCAogkGPiiwvQuQAmMdDqBbar6oLyeLHfVBzR`
