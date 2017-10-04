package syro

import (
	"testing"
)

func TestECC(t *testing.T) {
	bs := []byte("test1234\x00")
	ecc := ECC(bs)

	// calculated using:
	// https://github.com/korginc/volcasample/blob/9fced19d985cba37d489a0621ba6f1098658437b/syro/korg_syro_func.c#L87
	want := uint32(13578240)

	if ecc != want {
		t.Errorf("ECC was incorrect, got: %d, want: %d.", ecc, want)
	}
}
