package main

import (
        "fmt"
        "io"
	//"encoding/hex"
	"encoding/json"
        "net/http"
	"reflect"
	"unsafe"

        //"github.com/davecgh/go-spew/spew"
        //"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	//types1 "github.com/tendermint/tendermint/proto/tendermint/types"
)

func main() {
        GetRequest("http://124.156.136.62:28757/blocks/6476190")
	//a := "73B0F27A94E8413FE3A6C6451B73B2FE39C1B77A905082ECF8C7EA20C58C1110"
	//fmt.Printf("a: %v", stringtoslicebyte(a))
}

type block struct {
	Block_id blockid
	Block block2
}

type blockid struct {
	Hash string
}

type block2 struct {
	Header header
	Data data
}

type header struct {
	Chain_id string
}

type data struct {
	Txs []string
}

func GetRequest(url string) ([]byte, error) {
	fmt.Printf("url: %v\n", url)
        res, err := http.Get(url) // nolint:gosec
        if err != nil {
                return nil, err
        }
        defer res.Body.Close()

        body, err := io.ReadAll(res.Body)
        if err != nil {
                return nil, err
        }

        var getBlock block
        err = json.Unmarshal(body, &getBlock)
        if err != nil {
                fmt.Printf("json.Unmarrshal err: %v\n", err)
        }
        //spew.Printf("getBlock: %v\n", getBlock)
        fmt.Printf("block Hash: %x\n", getBlock.Block_id.Hash)
	fmt.Printf("chainid: %v\n", getBlock.Block.Header.Chain_id)
        txs := getBlock.Block.Data.Txs
	fmt.Printf("txs len: %v\n", len(txs))
	fmt.Printf("txs[0]: %v\n", txs[0])

        return body, nil
}

func BytesToString(b []byte) string {
    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
    sh := reflect.StringHeader{bh.Data, bh.Len}
    return *(*string)(unsafe.Pointer(&sh))
}

func stringtoslicebyte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
