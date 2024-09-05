package echo

import (
	"io"
	"net/http"
	"strings"

	"github.com/clysec/dagosy-worker/common"
)

// Universal Echo
// @Summary Echo the request
// @Description Echo the request
// @Tags echo
// @Accept json
// @Produce json
// @Success 200 {object} EchoResponse
// @Router /api/v1/echo [get]
// @Router /api/v1/echo [put]
// @Router /api/v1/echo [patch]
// @Router /api/v1/echo [post]
// @Router /api/v1/echo [delete]
func UniversalEcho(w http.ResponseWriter, r *http.Request) {
	response := EchoResponse{
		RequestType:   r.Method,
		RequestURL:    r.RequestURI,
		RequestSource: r.RemoteAddr,
		Headers:       r.Header,
		QueryParams:   r.URL.Query(),
		FormData:      make(map[string][]string, 0),
		RequestBody:   "",
	}

	if strings.HasPrefix(r.Header.Get("content-type"), "application/x-www-form-urlencoded") {
		r.ParseForm()
		response.FormData = r.Form
	} else if strings.HasPrefix(r.Header.Get("content-type"), "multipart/form-data") {
		r.ParseMultipartForm(10 << 20)

		response.FormData = r.MultipartForm.Value
		response.FormFiles = make(map[string][]string)
		for key, value := range r.MultipartForm.File {
			response.FormFiles[key] = []string{}
			for _, file := range value {
				response.FormFiles[key] = append(response.FormFiles[key], file.Filename)
			}
		}
	} else {
		body, _ := io.ReadAll(r.Body)
		response.RequestBody = string(body)
	}

	common.JsonResponse(w, http.StatusOK, response)
}
