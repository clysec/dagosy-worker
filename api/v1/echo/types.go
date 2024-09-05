package echo

type EchoResponse struct {
	RequestType   string              `json:"requestType"`
	RequestURL    string              `json:"requestURL"`
	RequestSource string              `json:"requestSource"`
	RequestBody   string              `json:"requestBody"`
	Headers       map[string][]string `json:"headers"`
	QueryParams   map[string][]string `json:"queryParams"`
	FormData      map[string][]string `json:"formData"`
	FormFiles     map[string][]string `json:"formFiles"`
}
