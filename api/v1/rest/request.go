package rest

import (
	"encoding/json"
	"net/http"

	"github.com/clysec/dagosy-worker/common"
	"github.com/clysec/greq"
)

// GET Reqest
// @Summary Make a GET Request
// @Description Make a GET Request
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /request/get [get]
func MakeRequest(w http.ResponseWriter, r *http.Request) {
	reqData := BaseRequest{}
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := greq.NewRequest(greq.Method(reqData.Method), reqData.URL)

	if len(reqData.Headers) > 0 {
		req = req.WithHeaders(reqData.Headers)
	}

	if len(reqData.Params) > 0 {
		req = req.WithQueryParams(reqData.Params)
	}

	switch reqData.BodyType {
	case BodyString:
		req = req.WithStringBody(reqData.BodyObject.(string))
	case BodyBytes:
		req = req.WithByteBody(reqData.BodyObject.([]byte))
	case BodyJson:
		req = req.WithJSONBody(reqData.BodyObject, nil)
	case BodyFormUrlencoded:
		req = req.WithUrlencodedFormBody(reqData.BodyObject.(map[string]interface{}), nil)
	case BodyFormMultipart:
		mpBody, err := greq.MultipartFieldsFromMap(reqData.BodyObject.(map[string]interface{}))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		req = req.WithMultipartFormBody(mpBody)
	}

	resp, err := req.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.StreamResponse(w, resp.StatusCode, resp.Headers["Content-Type"][0], resp.Response.Body)
}
