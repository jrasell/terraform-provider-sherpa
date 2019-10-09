package sherpa

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/jrasell/sherpa/pkg/api"
	clientCfg "github.com/jrasell/sherpa/pkg/config/client"
)

type ProviderConfig struct {
	client *api.Client
}

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"addr": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHERPA_ADDR", "http://localhost:8000"),
			},
			"ca_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHERPA_CA_CERT", ""),
			},
			"cert_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHERPA_CLIENT_CERT", ""),
			},
			"key_file": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SHERPA_CLIENT_KEY", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"sherpa_policy": policyResource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	defaultConfig := clientCfg.Config{
		Addr:        d.Get("addr").(string),
		CertPath:    d.Get("cert_file").(string),
		CertKeyPath: d.Get("key_file").(string),
		CAPath:      d.Get("ca_file").(string),
	}
	mergedConfig := api.DefaultConfig(&defaultConfig)

	client, err := api.NewClient(mergedConfig)
	if err != nil {
		return nil, err
	}
	return ProviderConfig{client: client}, nil
}
