//go:build windows

package chromium

import (
	"encoding/base64"
	"errors"
	"os"

	"PassGet/log"
	"PassGet/modules/utils/browser/crypto"
	"PassGet/modules/utils/browser/fileutil"
	"PassGet/modules/utils/browser/types"
	"github.com/tidwall/gjson"
)

var errDecodeMasterKeyFailed = errors.New("decode master key failed")

func (c *Chromium) GetMasterKey() ([]byte, error) {
	b, err := fileutil.ReadFile(types.ChromiumKey.TempFilename())
	if err != nil {
		return nil, err
	}
	defer os.Remove(types.ChromiumKey.TempFilename())

	encryptedKey := gjson.Get(b, "os_crypt.encrypted_key")
	if !encryptedKey.Exists() {
		return nil, nil
	}

	key, err := base64.StdEncoding.DecodeString(encryptedKey.String())
	if err != nil {
		return nil, errDecodeMasterKeyFailed
	}
	c.masterKey, err = crypto.DecryptWithDPAPI(key[5:])
	if err != nil {
		log.Errorf("decrypt master key failed, err %v", err)
		return nil, err
	}
	log.Debugf("get master key success, browser %s", c.name)
	return c.masterKey, nil
}