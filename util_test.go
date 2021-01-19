package lvm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BytesToHumanReadable(t *testing.T) {
	r := BytesToHumanReadable(1*1024*1024, MiB)

	assert.Equal(t, float64(1), r)

}

func Test_HumanReadableToBytes(t *testing.T) {
	r := HumanReadableToBytes(1, MiB)
	assert.Equal(t, uint64(1048576), r)
}

func Test_UnitTranslate(t *testing.T) {
	r := UnitTranslate(1*1024*1024, B, MiB)
	assert.Equal(t, float64(1), r)
}
