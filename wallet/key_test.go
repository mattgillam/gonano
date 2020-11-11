package wallet

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeriveKey(t *testing.T) {
	seed, _ := hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000001")
	key, err := deriveKey(seed, 1)
	require.Nil(t, err)
	assert.Equal(t, "1495f2d49159cc2eaaaa97ebb42346418e1268aff16d7fca90e6bad6d0965520", hex.EncodeToString(key))
}

func TestDeriveAddress(t *testing.T) {
	key, _ := hex.DecodeString("781186FB9EF17DB6E3D1056550D9FAE5D5BBADA6A6BC370E4CBB938B1DC71DA3")
	pubkey, err := derivePubkey(key)
	require.Nil(t, err)
	assert.Equal(t, "3068bb1ca04525bb0e416c485fe6a67fd52540227d267cc8b6e8da958a7fa039", hex.EncodeToString(pubkey))
	address, err := deriveAddress(pubkey)
	require.Nil(t, err)
	assert.Equal(t, "nano_1e5aqegc1jb7qe964u4adzmcezyo6o146zb8hm6dft8tkp79za3sxwjym5rx", address)
}

func TestAddressToPubkey(t *testing.T) {
	pubkey, err := addressToPubkey("nano_1e5aqegc1jb7qe964u4adzmcezyo6o146zb8hm6dft8tkp79za3sxwjym5rx")
	require.Nil(t, err)
	assert.Equal(t, "3068bb1ca04525bb0e416c485fe6a67fd52540227d267cc8b6e8da958a7fa039", hex.EncodeToString(pubkey))
}
