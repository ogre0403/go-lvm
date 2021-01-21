package lvm_test

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/ogre0403/go-lvm"
)

func Test_OpenVG(t *testing.T) {

	vg, err := lvm.VgOpen("vg-aaa", "r")

	if assert.Error(t, err) {
		assert.Equal(t, errors.New(fmt.Sprintf("Volume group \"vg-aaa\" not found")), err)
		return
	}

	defer vg.Close()
}

func Test_CreateLV(t *testing.T) {

	vgo, err := lvm.VgOpen("vg-0", "w")
	assert.NoError(t, err)
	defer vgo.Close()

	// Output some data of the VG
	//fmt.Printf("pvlist: %v\n", vgo.ListPVs())
	//fmt.Printf("free size: %4.2f %s\n", lvm.BytesToHumanReadable(vgo.GetFreeSize(), lvm.GiB), lvm.GiB)
	//fmt.Printf("size: %4.2f %s\n", lvm.BytesToHumanReadable(vgo.GetSize(), lvm.GiB), lvm.GiB)

	// Create a LV
	lv, err := vgo.CreateLvLinear("go-lvm-example-test-lv", lvm.HumanReadableToBytes(152, lvm.MiB))
	assert.NoError(t, err)
	assert.Equal(t, float64(152), lvm.BytesToHumanReadable(lv.GetSize(), lvm.MiB))

}

func Test_GetExistingLV(t *testing.T) {
	vgo, err := lvm.VgOpen("vg-0", "w")
	assert.NoError(t, err)
	defer vgo.Close()

	lv, err := vgo.LvFromName("go-lvm-example-test-lv")

	assert.NoError(t, err)

	assert.Equal(t, float64(152), lvm.BytesToHumanReadable(lv.GetSize(), lvm.MiB))

	// Remove LV
	err = lv.Remove()
	assert.NoError(t, err)
}

func TestVgObject_CreateLvThin(t *testing.T) {
	vgo, err := lvm.VgOpen("vg-0", "w")
	assert.NoError(t, err)
	defer vgo.Close()

	lv, err := vgo.CreateLvThin("pool0", "thin", lvm.HumanReadableToBytes(800, lvm.MiB))
	assert.NoError(t, err)
	// Verify first bit of attribute is V
	//https://www.mankier.com/8/lvs
	assert.Equal(t, "V", string(lv.GetAttr()[0]))

	err = lv.Remove()
	assert.NoError(t, err)
}
