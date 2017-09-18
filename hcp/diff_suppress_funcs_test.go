package hcp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha512sum(t *testing.T) {
	assert.Equal(t, "b109f3bbbc244eb82441917ed06d618b9008dd09b3befd1b5e07394c706a8bb980b1d7785e5976ec049b46df5f1326af5a2ea6d103fd07c95385ffab0cacbc86", sha512sum("password"))
	assert.Equal(t, "7dd29a9c9643fd524e1b4360964b89ce59914e68d1fd1ab04dd61fbaaabc58e579dcffb5b7454ab01e586c8ae98e538b5d6e0ff3ae7dd442de7333486dc9df1a", sha512sum("newpassword"))
}

func TestSuppressPasswordDiffs(t *testing.T) {
	assert.True(t, suppressPasswordDiffs("password", "b109f3bbbc244eb82441917ed06d618b9008dd09b3befd1b5e07394c706a8bb980b1d7785e5976ec049b46df5f1326af5a2ea6d103fd07c95385ffab0cacbc86", "password", nil))
	assert.False(t, suppressPasswordDiffs("password", "b109f3bbbc244eb82441917ed06d618b9008dd09b3befd1b5e07394c706a8bb980b1d7785e5976ec049b46df5f1326af5a2ea6d103fd07c95385ffab0cacbc86", "newpassword", nil))
}

func TestNormalizeHardQuota(t *testing.T) {
	assert.Equal(t, "100.00 GB", normalizeHardQuota("100.00 GB"))
	assert.Equal(t, "100.00 GB", normalizeHardQuota("100.0 GB"))
	assert.Equal(t, "100.00 GB", normalizeHardQuota("100 GB"))

	assert.Equal(t, "10.00 GB", normalizeHardQuota("10.00 GB"))
	assert.Equal(t, "10.00 GB", normalizeHardQuota("10.0 GB"))
	assert.Equal(t, "10.00 GB", normalizeHardQuota("10 GB"))

	assert.Equal(t, "1.00 GB", normalizeHardQuota("1.00 GB"))
	assert.Equal(t, "1.00 GB", normalizeHardQuota("1.0 GB"))
	assert.Equal(t, "1.00 GB", normalizeHardQuota("1 GB"))

	assert.Equal(t, "100.55 GB", normalizeHardQuota("100.55 GB"))
	assert.Equal(t, "100.50 GB", normalizeHardQuota("100.5 GB"))

}
