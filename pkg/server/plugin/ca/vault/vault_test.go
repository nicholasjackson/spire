package vault

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var mockController *gomock.Controller
var mockClient *MockVaultClient

func setupTests(t *testing.T) func() {
	mockController = gomock.NewController(t)
	// setup the mock factory
	createVaultClient = func() (VaultClient, error) {
		return NewMockVaultClient(mockController), nil
	}

	// cleanup
	return func() {
		mockController.Finish()
	}
}

func TestNewWithDefaultCreatesClientWithDefaultConfig(t *testing.T) {
	ca := NewWithDefault()

	assert.NotNil(t, ca, "Should have returned a ca.ServerCa instance")

	vcf := ca.(*VaultPlugin)
	assert.Equal(t, vcf.config.BackdateSecs, 0)
	assert.Equal(t, vcf.config.VaultAddr, "http://localhost:8200")
	assert.Fail(t, "Implemented the rest of this test")
}

func TestConfigureCreatesClientWithValidOptions(t *testing.T) {
	cleanup := setupTests(t)
	defer cleanup()

	assert.Fail(t, "Not implemented")
}
