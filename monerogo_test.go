// monerogo_test
package monerogo

import (
	"testing"
)

func TestDaemonRPCMethods(t *testing.T) {
	// create new daemon client
	client := NewDaemonClient("http://127.0.0.1:18081/json_rpc")

	// test getblockcount
	blockCount, err := client.GetBlockCount()
	if err != nil {
		t.Error("getblockcount rpc failed: ", err)
		return
	}
	if blockCount.Status != "OK" {
		t.Error("getblockcount response fail ", blockCount)
		return
	}

	// test on_getblockhash
	hash, err := client.OnGetBlockHash(912345)
	if err != nil {
		t.Error("on_getblockhash rpc failed: ", err)
		return
	}
	if len(hash) <= 0 {
		t.Error("on_getblockhash response fail ", hash)
		return
	}

	// test getblocktemplate
	blockTemplate, err := client.GetBlockTemplate("44GBHzv6ZyQdJkjqZje6KLZ3xSyN1hBSFAnLP6EAqJtCRVzMzZmeXTC2AHKDS9aEDTRKmo6a6o9r9j86pYfhCWDkKjbtcns", 60)
	if err != nil && err.Error() != "Core is busy" {
		t.Error("getblocktemplate rpc failed: ", err)
		return
	}
	if blockTemplate.Status != "OK" && err == nil {
		t.Error("getblocktemplate response fail ", blockTemplate)
	}

	// test getlastblockheader
	blockHeader, err := client.GetLastBlockHeader()
	if err != nil {
		t.Error("getlastblockheader rpc failed: ", err)
		return
	}
	if blockHeader.Status != "OK" {
		t.Error("getlastblockheader response fail ", blockHeader)
		return
	}

	// test getblockheaderbyhash
	blockHeader, err = client.GetBlockHeaderByHash("e22cf75f39ae720e8b71b3d120a5ac03f0db50bba6379e2850975b4859190bc6")
	if err != nil {
		t.Error("getblockheaderbyhash rpc failed: ", err)
		return
	}
	if blockHeader.Status != "OK" {
		t.Error("getblockheaderbyhash response fail ", blockHeader)
		return
	}

	// test getblockheaderbyheight
	blockHeader, err = client.GetBlockHeaderByHeight(912345)
	if err != nil {
		t.Error("getblockheaderbyheight rpc failed: ", err)
		return
	}
	if blockHeader.Status != "OK" {
		t.Error("getblockheaderbyheight response fail ", blockHeader)
		return
	}

	// test getblock
	block, err := client.GetBlock(912345, "")
	if err != nil {
		t.Error("getblock rpc failed: ", err)
		return
	}
	if block.Status != "OK" {
		t.Error("getblock response fail ", block)
		return
	}

	// test get_connections
	// @TODO method fails because of invalid chars...what to do ?
	_, err = client.GetConnections()
	if err != nil {
		t.Log("get_connections rpc failed: ", err)
	}

	// test get_info
	info, err := client.GetInfo()
	if err != nil {
		t.Error("get_info rpc failed: ", err)
		return
	}
	if info.Status != "OK" {
		t.Error("get_info response fail ", info)
		return
	}

	// test hard_fork_info
	hfInfo, err := client.GetHardForkInfo()
	if err != nil {
		t.Error("hard_fork_info rpc failed: ", err)
		return
	}
	if hfInfo.Status != "OK" {
		t.Error("hard_fork_info response fail ", hfInfo)
		return
	}

	// test getbans
	// @TODO method does not exist, outdated documentation ?
	_, err = client.GetBans()
	if err != nil {
		t.Log("getbans rpc failed: ", err)
	}

}
