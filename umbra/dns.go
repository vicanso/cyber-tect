package umbra

import (
	"errors"
	"time"

	"github.com/miekg/dns"
)

type (
	// DNS dns checker
	DNS struct {
		Server   string
		Hostname string
	}
)

// Check check the dns resolve
func (d *DNS) Check() (healthy bool, extra map[string]interface{}, err error) {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(d.Hostname), dns.TypeA)
	m.RecursionAvailable = true
	started := time.Now()
	r, _, err := c.Exchange(m, d.Server)
	if err != nil {
		return
	}
	if r.Rcode != dns.RcodeSuccess {
		err = errors.New("get rcode fail")
		return
	}
	extra = make(map[string]interface{})
	answers := []string{}
	for _, a := range r.Answer {
		answers = append(answers, a.String())
	}
	extra["answers"] = answers
	extra["stats"] = map[string]string{
		"total": time.Since(started).String(),
	}
	healthy = true
	return
}

// GetDescription get the description of dns checker
func (d *DNS) GetDescription() (description map[string]interface{}) {
	description = make(map[string]interface{})
	description["type"] = TypeDNS
	description["hostname"] = d.Hostname
	description["server"] = d.Server
	return
}
