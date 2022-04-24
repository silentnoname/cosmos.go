package key

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateMnemonic(t *testing.T) {
	_, err := CreateMnemonic()
	assert.NoError(t, err)
}

func Test_DrivePrivKey(t *testing.T) {
	mnemonic, err := CreateMnemonic()
	assert.NoError(t, err)

	// Only Secp256k1 is supported
	_, err = DerivePrivKeyBz(mnemonic, CreateHDPath(1, 118, 1))
	assert.NoError(t, err)
}

func Test_GenerateMnemonicandAddress(t *testing.T) {
	_, _, err := GenerateMnemonicandAddress(118, "cosmos")
	assert.NoError(t, err)
}
