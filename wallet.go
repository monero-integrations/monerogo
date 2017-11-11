// Copyright 2017 Marin Basic <marin@marin-basic.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package monerogo

type WalletClient struct {
	endpoint string
}

// Creates new wallet client
func NewWalletClient(endpoint string) *WalletClient {
	return &WalletClient{endpoint: endpoint}
}
