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
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/rest/request [get]
// @Router /api/v1/rest/request [post]
// @Router /api/v1/rest/request [put]
// @Router /api/v1/rest/request [delete]
// @Router /api/v1/rest/request [patch]
// @Router /api/v1/rest/request [options]
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

	if reqData.UseCertAuth {
		certAuth := greq.NewClientCertificateAuth().FromX509Bytes([]byte(reqData.Certificate), []byte(reqData.PrivateKey))
		if !reqData.ValidateCa {
			certAuth = certAuth.WithInsecureSkipVerify(true)
		}

		req = req.WithAuth(certAuth)
	}

	resp, err := req.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.StreamResponse(w, resp.StatusCode, resp.Headers, resp.Response.Body)
}
