<<<<<<< HEAD
package tp_test
=======
package erpc_test
>>>>>>> 9a46c7d6f1738b0bb85ebe4a21db0e6dcc501751

import (
	"testing"

<<<<<<< HEAD
	tp "github.com/henrylee2cn/teleport"
=======
	"github.com/henrylee2cn/erpc/v6"
	"github.com/stretchr/testify/assert"
>>>>>>> 9a46c7d6f1738b0bb85ebe4a21db0e6dcc501751
)

func TestHTTPServiceMethodMapper(t *testing.T) {
	var cases = []struct{ src, dst string }{
		{"AaBb", "/aa_bb"},
		{"ABcXYz", "/abc_xyz"},
		{"Aa__Bb", "/aa_bb"},
		{"Aa________Bb", "/aa_bb"},
		{"aa__bb", "/aa_bb"},
		{"ABC__XYZ", "/abc_xyz"},
		{"Aa_Bb", "/aa/bb"},
		{"aa_bb", "/aa/bb"},
		{"ABC_XYZ", "/abc/xyz"},
	}
	for _, c := range cases {
<<<<<<< HEAD
		got := tp.HTTPServiceMethodMapper("", c.src)
=======
		got := erpc.HTTPServiceMethodMapper("", c.src)
>>>>>>> 9a46c7d6f1738b0bb85ebe4a21db0e6dcc501751
		if got != c.dst {
			t.Fatalf("%s: got: %s, expect: %s", c.src, got, c.dst)
		}
	}
}

func TestRPCServiceMethodMapper(t *testing.T) {
	var cases = []struct{ src, dst string }{
		{"AaBb", "AaBb"},
		{"ABcXYz", "ABcXYz"},
		{"Aa__Bb", "Aa_Bb"},
		{"Aa________Bb", "Aa_Bb"},
		{"aa__bb", "aa_bb"},
		{"ABC__XYZ", "ABC_XYZ"},
		{"Aa_Bb", "Aa.Bb"},
		{"aa_bb", "aa.bb"},
		{"ABC_XYZ", "ABC.XYZ"},
	}
	for _, c := range cases {
<<<<<<< HEAD
		got := tp.RPCServiceMethodMapper("", c.src)
=======
		got := erpc.RPCServiceMethodMapper("", c.src)
>>>>>>> 9a46c7d6f1738b0bb85ebe4a21db0e6dcc501751
		if got != c.dst {
			t.Fatalf("%s: got: %s, expect: %s", c.src, got, c.dst)
		}
	}
}
<<<<<<< HEAD
=======

func TestFakeAddr(t *testing.T) {
	addr := erpc.NewFakeAddr("", "", "")
	assert.Equal(t, "0.0.0.0:0", addr.String())
	assert.Equal(t, "tcp", addr.Network())

	addr = erpc.NewFakeAddr("tcp", "", "1234")
	assert.Equal(t, "0.0.0.0:1234", addr.String())

	addr, err := erpc.NewFakeAddr2("", "")
	assert.NoError(t, err)
	assert.Equal(t, "0.0.0.0:0", addr.String())
	assert.Equal(t, "tcp", addr.Network())
	assert.Equal(t, "0.0.0.0", addr.Host())
	assert.Equal(t, "0", addr.Port())

	addr, err = erpc.NewFakeAddr2("tcp6", ":1234")
	assert.NoError(t, err)
	assert.Equal(t, "0.0.0.0:1234", addr.String())
	assert.Equal(t, "tcp6", addr.Network())
	assert.Equal(t, "0.0.0.0", addr.Host())
	assert.Equal(t, "1234", addr.Port())

	addr, err = erpc.NewFakeAddr2("tcp6", "192.0.0.10:1234")
	assert.NoError(t, err)
	assert.Equal(t, "192.0.0.10:1234", addr.String())
	assert.Equal(t, "tcp6", addr.Network())
	assert.Equal(t, "192.0.0.10", addr.Host())
	assert.Equal(t, "1234", addr.Port())
}
>>>>>>> 9a46c7d6f1738b0bb85ebe4a21db0e6dcc501751
