package dns

import (
	"os"
	"reflect"
	"testing"

	"github.com/appscode/go-dns/aws"
	"github.com/stretchr/testify/assert"
)

var (
	apiKey    string
	apiSecret string
)

func init() {
	apiSecret = os.Getenv("EXOSCALE_API_SECRET")
	apiKey = os.Getenv("EXOSCALE_API_KEY")
}

func restoreExoscaleEnv() {
	os.Setenv("EXOSCALE_API_KEY", apiKey)
	os.Setenv("EXOSCALE_API_SECRET", apiSecret)
}

func TestKnownDNSProviderSuccess(t *testing.T) {
	os.Setenv("EXOSCALE_API_KEY", "abc")
	os.Setenv("EXOSCALE_API_SECRET", "123")
	provider, err := NewDNSProvider("exoscale")
	assert.NoError(t, err)
	assert.NotNil(t, provider)
	if reflect.TypeOf(provider) != reflect.TypeOf(&aws.DNSProvider{}) {
		t.Errorf("Not loaded correct DNS proviver: %v is not *exoscale.DNSProvider", reflect.TypeOf(provider))
	}
	restoreExoscaleEnv()
}

func TestKnownDNSProviderError(t *testing.T) {
	os.Setenv("EXOSCALE_API_KEY", "")
	os.Setenv("EXOSCALE_API_SECRET", "")
	_, err := NewDNSProvider("exoscale")
	assert.Error(t, err)
	restoreExoscaleEnv()
}

func TestUnknownDNSProvider(t *testing.T) {
	_, err := NewDNSProvider("foobar")
	assert.Error(t, err)
}
