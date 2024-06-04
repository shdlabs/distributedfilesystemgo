package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSome(t *testing.T) {
	// t.Run("empty msg with DefautlDecoder", func(t *testing.T) {
	// 	assert := assert.New(t)
	// 	d := DefaultDecoder{}
	// 	r := bufio.NewReader(os.Stdout)
	// 	msg := new(RPC)
	//
	// 	t.Logf("MSG before: %v", msg)
	// 	assert.NoError(d.Decode(r, msg))
	// 	t.Logf("MSG after : %v", msg)
	// })

	t.Run("check later", func(t *testing.T) {
		assert := assert.New(t)

		assert.Equal(1, 1, "some thing that should fail")
	})
}
