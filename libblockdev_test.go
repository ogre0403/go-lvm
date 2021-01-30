package lvm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	Initialize()

}

func TestBD_LVM_LvCreate(t *testing.T) {
	e := BD_LVM_LvCreate("vg-0", "kkk", HumanReadableToBytes(12, MiB))
	assert.NoError(t, e)

	t.Cleanup(func() {
		e := BD_LVM_LvRemove("vg-0", "kkk")
		assert.NoError(t, e)
	})
}

func TestBD_LVM_ThLVCreate(t *testing.T) {
	e := BD_LVM_ThLVCreate("vg-0", "pool0", "thth", HumanReadableToBytes(20, MiB))
	assert.NoError(t, e)

	v, e := BD_LVM_LvInfo("vg-0", "thth")
	assert.NoError(t, e)
	if e == nil {
		assert.Equal(t, v.IsThinVolume(), true)
	}

	t.Cleanup(func() {
		e := BD_LVM_LvRemove("vg-0", "thth")
		assert.NoError(t, e)
	})
}
