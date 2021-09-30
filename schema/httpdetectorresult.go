package schema

import (
	"encoding/json"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// HTTPDetectorResult holds the schema definition for the HTTPDetectorResult entity.
type HTTPDetectorResult struct {
	ent.Schema
}

// HTTPDetectorSubResult http detector result
type HTTPDetectorSubResult struct {
	Result                     DetectorResult `json:"result"`
	ResultDesc                 string         `json:"resultDesc"`
	Addrs                      []string       `json:"addrs"`
	Addr                       string         `json:"addr"`
	Protocol                   string         `json:"protocol"`
	TLSVersion                 string         `json:"tlsVersion"`
	TLSCipherSuite             string         `json:"tlsCipherSuite"`
	CertificateDNSNames        []string       `json:"certificateDNSNames"`
	CertificateExpirationDates []string       `json:"certificateExpirationDates"`
	DNSLookup                  int            `json:"dnsLookup"`
	TCPConnection              int            `json:"tcpConnection"`
	TLSHandshake               int            `json:"tlsHandshake"`
	ServerProcessing           int            `json:"serverProcessing"`
	ContentTransfer            int            `json:"contentTransfer"`
	Duration                   int            `json:"duration"`
	Message                    string         `json:"message"`
}
type MarshalHTTPDetectorSubResult HTTPDetectorSubResult

func (sr *HTTPDetectorSubResult) MarshalJSON() ([]byte, error) {
	tmp := (*MarshalHTTPDetectorSubResult)(sr)
	tmp.ResultDesc = tmp.Result.String()
	return json.Marshal(tmp)
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
