package hcp

import (
	"github.com/digipost/hcp"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldValidateNamespaceNames(t *testing.T) {
	assert.True(t, valid("a-b-c", validateNamespaceName))
	assert.True(t, valid("a", validateNamespaceName))
	assert.False(t, valid("way-too-long-way-too-long-way-too-long-way-too-long-way-too-long-way-too-long-way-too-long", validateNamespaceName))
	assert.False(t, valid("-cannot-start-with-hyphen", validateNamespaceName))
	assert.False(t, valid("cannot-end-with-hyphen-", validateNamespaceName))
	assert.False(t, valid("xn--illeval-start", validateNamespaceName))
	assert.False(t, valid("name containing whitespace", validateNamespaceName))
}

func TestShouldValidateNamespaceHardQuotas(t *testing.T) {
	assert.True(t, valid("1 GB", validateHardQuota))
	assert.True(t, valid("10 GB", validateHardQuota))
	assert.True(t, valid("100 GB", validateHardQuota))
	assert.True(t, valid("1000 GB", validateHardQuota))
	assert.True(t, valid("1 TB", validateHardQuota))
	assert.True(t, valid("10 TB", validateHardQuota))
	assert.True(t, valid("100 TB", validateHardQuota))
	assert.True(t, valid("1.0 GB", validateHardQuota))
	assert.True(t, valid("10.0 GB", validateHardQuota))
	assert.True(t, valid("10.10 TB", validateHardQuota))
	assert.False(t, valid("10.100 TB", validateHardQuota))
	assert.False(t, valid("10 KB", validateHardQuota))
}

func TestShoulValidateNamespaceSoftQuotas(t *testing.T) {
	assert.True(t, valid(0, validateSoftQuota))
	assert.True(t, valid(75, validateSoftQuota))
	assert.True(t, valid(100, validateSoftQuota))
	assert.False(t, valid(-1, validateSoftQuota))
	assert.False(t, valid(101, validateSoftQuota))
}

func TestShouldValidateHashSchemes(t *testing.T) {
	assert.True(t, valid(hcp.SHA_256, validateHashScheme))
	assert.False(t, valid("SHA256", validateHashScheme))
}

func TestShouldValidateOptimizedFor(t *testing.T) {
	assert.True(t, valid(hcp.CLOUD, validateOptimizedFor))
	assert.False(t, valid("NFS", validateOptimizedFor))
}

func valid(value interface{}, validationFunction schema.SchemaValidateFunc) bool {
	warnings, errors := validationFunction(value, "")
	return len(warnings) == 0 && len(errors) == 0
}
