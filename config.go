package testbed

import (
	"encoding/json"
	"net"
	"time"
)

// Config is the configuration for the testbed server
// Note: The custom unmarshallers & marshallers below should be
// updated  if fields are added or removed from config
type Config struct {
	NodeID                   string
	GRPCAddress              string
	TLSServerCertificate     string
	TLSServerKey             string
	TLSClientCertificate     string
	TLSClientKey             string
	TLSInsecureSkipVerify    bool
	ContainerAddr            string
	Namespace                string
	Subnet                   *net.IPNet
	DataDir                  string
	StateDir                 string
	Bridge                   string
	UpstreamDNSAddr          string
	ProxyHTTPPort            int
	ProxyHTTPSPort           int
	ProxyTLSEmail            string
	ProxyHealthcheckInterval time.Duration
	GatewayAddress           string
	EventsAddress            string
	EventsHTTPAddress        string
}

// MarshalJSON is a custom JSON marshaller
func (c *Config) MarshalJSON() ([]byte, error) {
	type Alias Config
	return json.Marshal(&struct {
		*Alias
		Subnet                   string
		ProxyHealthcheckInterval string
	}{
		Alias:                    (*Alias)(c),
		Subnet:                   c.Subnet.String(),
		ProxyHealthcheckInterval: c.ProxyHealthcheckInterval.String(),
	})
}

// UnmarshalJSON is a custom JSON unmarshaller
func (c *Config) UnmarshalJSON(data []byte) error {
	type Alias Config
	tmp := &struct {
		*Alias
		Subnet                   string
		ProxyHealthcheckInterval string
	}{Alias: (*Alias)(c)}

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	_, subnet, err := net.ParseCIDR(tmp.Subnet)
	if err != nil {
		return err
	}
	c.Subnet = subnet
	d, err := time.ParseDuration(tmp.ProxyHealthcheckInterval)
	if err != nil {
		return err
	}
	c.ProxyHealthcheckInterval = d

	return nil
}
