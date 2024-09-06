package rest

type RequestMethod string
type BodyType string
type ResponseType string

const (
	GET    RequestMethod = "GET"
	POST   RequestMethod = "POST"
	PUT    RequestMethod = "PUT"
	PATCH  RequestMethod = "PATCH"
	DELETE RequestMethod = "DELETE"

	BodyString         BodyType = "string"
	BodyJson           BodyType = "json"
	BodyFormMultipart  BodyType = "form-multipart"
	BodyFormUrlencoded BodyType = "form-urlencoded"
	BodyBytes          BodyType = "bytes"

	ResponseString ResponseType = "string"
	ResponseJson   ResponseType = "json"
	ResponseBytes  ResponseType = "bytes"
)

type BaseRequest struct {
	Method       RequestMethod `json:"method"`
	ResponseType ResponseType  `json:"responseType"`

	URL     string                 `json:"url"`
	Headers map[string]interface{} `json:"headers"`
	Params  map[string]string      `json:"params"`

	BodyType   BodyType    `json:"bodyType"`
	BodyObject interface{} `json:"bodyObject"`

	UseCertAuth bool   `json:"useCertAuth"`
	Certificate string `json:"certificate"`
	PrivateKey  string `json:"privateKey"`

	ValidateCa     bool   `json:"validateCa"`
	ValidateCaCert string `json:"validateCaCert"`
}
