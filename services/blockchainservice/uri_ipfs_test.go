package blockchainservice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"gitlab.123xe.vn/TrustKeysV2/Common/MappingUIDPubkey/utils/log"
)

func TestURIIpfs(t *testing.T) {
	resp, err := http.Get("https://ipfs.io/ipfs/Qmewft6TqiPq1yj1WB6c2zh5arsfyBMMhkp5QcpLKBRy8M/metadata.json")
	// resp, err := http.Get("https://ipfs.io/ipfs/bafkreif5crcyihj34zwahyjwrdoyyf34lq5cre4j5ql4gxwnqoxaxn756e")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	dataUri := make(map[string]interface{})
	err = json.Unmarshal(body, &dataUri)
	if err != nil {
		panic(err)
	}

	log.Println(dataUri)
}
