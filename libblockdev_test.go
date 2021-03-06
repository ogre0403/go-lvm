package lvm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	Initialize()

}

func TestBD_LVM_PvCreate(t *testing.T) {
	e := BD_LVM_PvCreate("/dev/loop3")
	assert.NoError(t, e)
}

func TestBD_LVM_VgCreate(t *testing.T) {

	pvlist := []string{"/dev/loop3"}
	e := BD_LVM_VgCreate("vg-1", pvlist)
	assert.NoError(t, e)
}

func TestBD_LVM_Lv_CreateRemove(t *testing.T) {
	e := BD_LVM_LvCreate("vg-1", "kkk", HumanReadableToBytes(12, MiB))
	assert.NoError(t, e)

	t.Cleanup(func() {
		e := BD_LVM_LvRemove("vg-1", "kkk")
		assert.NoError(t, e)
	})
}

func TestBD_LVM_ThPoolCreate(t *testing.T) {
	e := BD_LVM_ThPoolCreate("vg-1", "pool1", HumanReadableToBytes(200, MiB))
	assert.NoError(t, e)
}

func TestBD_LVM_ThLv_CreateRemove(t *testing.T) {
	e := BD_LVM_ThLvCreate("vg-1", "pool1", "thth", HumanReadableToBytes(300, MiB))
	assert.NoError(t, e)

	v, e := BD_LVM_LvInfo("vg-1", "thth")
	assert.NoError(t, e)
	if e == nil {
		assert.Equal(t, v.IsThinVolume(), true)
	}

	t.Cleanup(func() {
		e := BD_LVM_LvRemove("vg-1", "thth")
		assert.NoError(t, e)
	})
}

func TestBD_LVM_VgRemove(t *testing.T) {
	e := BD_LVM_VgRemove("vg-1")
	assert.NoError(t, e)
}

func TestBD_LVM_PvRemove(t *testing.T) {
	e := BD_LVM_PvRemove("/dev/loop3")
	assert.NoError(t, e)
}
