package vault

import (
	"context"

	"github.com/hashicorp/vault/api"
	spi "github.com/spiffe/spire/proto/common/plugin"
	"github.com/spiffe/spire/proto/server/ca"
)

type configuration struct {
	VaultAddr    string `hcl:"vault_address" json:"vault_address"`
	VaultToken   string `hcl:"vault_token" json:"vault_token"`
	PKIMountPath string `hcl:"pki_mount_path" json:"pki_mount_path"`
	TrustDomain  string `hcl:"trust_domain" json:"trust_domain"`
	BackdateSecs int    `hcl:"backdate_seconds" json:"backdate_seconds"`
	KeySize      int    `hcl:"key_size" json:"key_size"`
	//	CertSubject  certSubjectConfig `hcl:"cert_subject" json:"cert_subject"`
	DefaultTTL int `hcl:"default_ttl" json:"default_ttl"`
}

type vaultClientFactory func() (VaultClient, error)

var createVaultClient vaultClientFactory = func() (VaultClient, error) {
	return api.NewClient(api.DefaultConfig())
}

type VaultPlugin struct {
	config configuration
	client *api.Client
}

func (m *VaultPlugin) Configure(ctx context.Context, req *spi.ConfigureRequest) (*spi.ConfigureResponse, error) {
	return nil, nil
}

func (m *VaultPlugin) GetPluginInfo(ctx context.Context, req *spi.GetPluginInfoRequest) (*spi.GetPluginInfoResponse, error) {
	return &spi.GetPluginInfoResponse{}, nil
}

func (m *VaultPlugin) SignCsr(ctx context.Context, request *ca.SignCsrRequest) (*ca.SignCsrResponse, error) {
	return nil, nil
}

func (m *VaultPlugin) GenerateCsr(ctx context.Context, req *ca.GenerateCsrRequest) (*ca.GenerateCsrResponse, error) {
	return nil, nil
}

func (m *VaultPlugin) FetchCertificate(ctx context.Context, request *ca.FetchCertificateRequest) (*ca.FetchCertificateResponse, error) {
	return nil, nil
}

func (m *VaultPlugin) LoadCertificate(ctx context.Context, request *ca.LoadCertificateRequest) (response *ca.LoadCertificateResponse, err error) {
	return nil, nil
}

func NewWithDefault() ca.ServerCa {
	c := configuration{
		VaultAddr:    "http://localhost:8200",
		VaultToken:   "root",
		PKIMountPath: "pki",
		TrustDomain:  "mydomain.com",
		BackdateSecs: 0,
		KeySize:      4096,
		DefaultTTL:   24,
	}

	return &VaultPlugin{config: c}
}
