// Copyright 2017 Marin Basic <marin@marin-basic.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package monerogo

import (
	"bytes"

	"net/http"

	"github.com/gorilla/rpc/v2/json2"
)

// Calls JSON RPC V2 with encoding request and decoding response
func call(url string, method string, req interface{}, res interface{}) error {
	clientPayload, err := json2.EncodeClientRequest(method, req)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(clientPayload))
	if err != nil {
		return err
	}

	return json2.DecodeClientResponse(resp.Body, res)
}
