package ipinfo

import (
	"net/netip"

	"github.com/qdm12/gluetun/internal/models"
)

type Response struct {
	IP       netip.Addr `json:"ip,omitempty"`
	Region   string     `json:"region,omitempty"`
	Country  string     `json:"country,omitempty"`
	City     string     `json:"city,omitempty"`
	Hostname string     `json:"hostname,omitempty"`
	Loc      string     `json:"loc,omitempty"`
	Org      string     `json:"org,omitempty"`
	Postal   string     `json:"postal,omitempty"`
	Timezone string     `json:"timezone,omitempty"`
}

func (r *Response) ToPublicIPModel() (model models.PublicIP) {
	return models.PublicIP{
		IP:           r.IP,
		Region:       r.Region,
		Country:      r.Country,
		City:         r.City,
		Hostname:     r.Hostname,
		Location:     r.Loc,
		Organization: r.Org,
		PostalCode:   r.Postal,
		Timezone:     r.Timezone,
	}
}
