package syro

import (
	"testing"
)

func TestCRC16(t *testing.T) {
	bs := []byte("test1234\x00")
	crc := CRC16(bs)

	// calculated using:
	// https://github.com/korginc/volcasample/blob/9fced19d985cba37d489a0621ba6f1098658437b/syro/korg_syro_func.c#L70
	want := uint16(50948)

	if crc != want {
		t.Errorf("CRC16 was incorrect, got: %d, want: %d.", crc, want)
	}
}
