package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// HTTPDetectorResult holds the schema definition for the HTTPDetectorResult entity.
type HTTPDetectorResult struct {
	ent.Schema
}

// HTTPDetectorSubResult http detector result
type HTTPDetectorSubResult struct {
	Result                     DetectorResult `json:"result,omitempty"`
	Addrs                      []string       `json:"addrs,omitempty"`
	Addr                       string         `json:"addr,omitempty"`
	Protocol                   string         `json:"protocol,omitempty"`
	TLSVersion                 string         `json:"tlsVersion,omitempty"`
	TLSCipherSuite             string         `json:"tlsCipherSuite,omitempty"`
	CertificateDNSNames        []string       `json:"certificateDNSNames,omitempty"`
	CertificateExpirationDates []string       `json:"certificateExpirationDates,omitempty"`
	DNSLookup                  int            `json:"dnsLookup,omitempty"`
	TCPConnection              int            `json:"tcpConnection,omitempty"`
	TLSHandshake               int            `json:"tlsHandshake,omitempty"`
	ServerProcessing           int            `json:"serverProcessing,omitempty"`
	ContentTransfer            int            `json:"contentTransfer,omitempty"`
	Duration                   int            `json:"duration,omitempty"`
	Message                    string         `json:"message,omitempty"`
}

type HTTPDetectorSubResults []*HTTPDetectorSubResult

// Fields of the HTTPDetectorResult.
func (HTTPDetectorResult) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").
			Comment("检测URL"),
		field.JSON("results", HTTPDetectorSubResults{}).
			Comment("检测结果列表"),
	}
}

// Mixin of the HTTPDetectorResult
func (HTTPDetectorResult) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DetectorResultMixin{},
	}
}

// Edges of the HTTPDetectorResult.
func (HTTPDetectorResult) Edges() []ent.Edge {
	return nil
}
