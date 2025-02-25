package types

type Recipe struct {
	ID           string
	Name         string
	Ingredients  []Ingredient
	Instructions string
}

type Ingredient struct {
	ID     string
	Name   string
	Amount int
	Unit   string
}

type RequestShape struct {
	Version        string `json:"version"`
	RouteKey       string `json:"routeKey"`
	RawPath        string `json:"rawPath"`
	RawQueryString string `json:"rawQueryString"`
	Headers        struct {
		ContentLength       string `json:"content-length"`
		XAmznTLSVersion     string `json:"x-amzn-tls-version"`
		XForwardedProto     string `json:"x-forwarded-proto"`
		XForwardedPort      string `json:"x-forwarded-port"`
		XForwardedFor       string `json:"x-forwarded-for"`
		Accept              string `json:"accept"`
		XAmznTLSCipherSuite string `json:"x-amzn-tls-cipher-suite"`
		XAmznTraceID        string `json:"x-amzn-trace-id"`
		Host                string `json:"host"`
		RequestStartTime    string `json:"request-start-time"`
		ContentType         string `json:"content-type"`
		AcceptEncoding      string `json:"accept-encoding"`
		UserAgent           string `json:"user-agent"`
	} `json:"headers"`
	RequestContext struct {
		AccountID    string `json:"accountId"`
		APIID        string `json:"apiId"`
		DomainName   string `json:"domainName"`
		DomainPrefix string `json:"domainPrefix"`
		HTTP         struct {
			Method    string `json:"method"`
			Path      string `json:"path"`
			Protocol  string `json:"protocol"`
			SourceIP  string `json:"sourceIp"`
			UserAgent string `json:"userAgent"`
		} `json:"http"`
		RequestID string `json:"requestId"`
		RouteKey  string `json:"routeKey"`
		Stage     string `json:"stage"`
		Time      string `json:"time"`
		TimeEpoch int64  `json:"timeEpoch"`
	} `json:"requestContext"`
	Body            string `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded"`
}
