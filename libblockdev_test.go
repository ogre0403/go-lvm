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
}
