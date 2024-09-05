package files

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/clysec/dagosy-worker/common"
	_ "github.com/rclone/rclone/backend/all"
	"github.com/rclone/rclone/fs/object"
	"github.com/rclone/rclone/fs/operations"
)

// List Files
// @Summary List Files
// @Description List Files
// @Tags files
// @Accept json
// @Produce json
// @Param remote body ListFilesRequest true "Remote Configuration"
// @Success 200 {object} ListFilesResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/list [post]
func ListFiles(w http.ResponseWriter, r *http.Request) {
	var request ListFilesRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request.Remote.Parameters["fast-list"] = true

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	outData := ListFilesResponse{}

	err = operations.ListJSON(context.Background(), tfs, request.Path, &request.Options, func(item *operations.ListJSONItem) error {
		outData.Files = append(outData.Files, *item)
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outData.Total = int64(len(outData.Files))
	common.JsonResponse(w, http.StatusOK, outData)
}

// Filtered List Files
// @Summary Filtered List Files
// @Description List files in a given directory with a filter
// @Tags files
// @Accept json
// @Produce json
// @Param remote body FilteredListFilesRequest true "Remote Configuration"
// @Success 200 {object} FilteredListFilesResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/list [post]
func FilteredListFiles(w http.ResponseWriter, r *http.Request) {
	var request FilteredListFilesRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request.Remote.Parameters["fast-list"] = true

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	outData := ListFilesResponse{}

	filterFunc := func(item *operations.ListJSONItem) bool {
		return true
	}

	switch request.FilterType {
	case FilterTypePrefix:
		filterFunc = func(item *operations.ListJSONItem) bool {
			return strings.HasPrefix(item.Name, request.Filter)
		}
	case FilterTypeSuffix:
		filterFunc = func(item *operations.ListJSONItem) bool {
			return strings.HasSuffix(item.Name, request.Filter)
		}
	case FilterTypeRegex:
		filterFunc = func(item *operations.ListJSONItem) bool {
			return regexp.MustCompile(request.Filter).MatchString(item.Name)
		}
	case FilterTypeWildcard:
		filterFunc = func(item *operations.ListJSONItem) bool {
			return strings.Contains(item.Name, request.Filter)
		}
	}

	err = operations.ListJSON(context.Background(), tfs, request.Path, &request.Options, func(item *operations.ListJSONItem) error {
		if filterFunc(item) {
			outData.Files = append(outData.Files, *item)
		}

		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	outData.Total = int64(len(outData.Files))
	common.JsonResponse(w, http.StatusOK, outData)
}

// Read File
// @Summary Read File
// @Description Read File
// @Tags files
// @Accept json
// @Produce json
// @Param remote body ReadFileRequest true "Remote Configuration"
// @Success 200 {object} ReadFileResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/read [post]
func ReadFile(w http.ResponseWriter, r *http.Request) {
	var request ListFilesRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request.Remote.Parameters["fast-list"] = true

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tfo, err := tfs.NewObject(r.Context(), request.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if tfo.Size() == 0 || tfo.Size() == -1 {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	stat, err := operations.StatJSON(r.Context(), tfs, request.Path, &operations.ListJSONOpt{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content, err := operations.ReadFile(r.Context(), tfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ByteResponse(w, http.StatusOK, stat.MimeType, content)
}

// Write File
// @Summary Write File
// @Description Write File. You can attach an arbitrary number of files to the request. All have to be placed in the "file" field.
// @Tags files
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Param remote formData string true "Remote Configuration"
// @Param path formData string true "Path"
// @Param overwrite formData string false "Overwrite"
// @Success 200 {object} WriteFileResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/write [put]
func WriteFile(w http.ResponseWriter, r *http.Request) {
	request := WriteFileRequest{}

	remoteData := r.FormValue("remote")
	if remoteData == "" {
		http.Error(w, "Remote configuration is required", http.StatusBadRequest)
		return
	}

	overwrite := strings.Trim(r.FormValue("overwrite"), "\r\n")
	if slices.Contains([]string{"true", "1", "yes", "y"}, overwrite) {
		request.Overwrite = true
	}

	json.Unmarshal([]byte(remoteData), &request.Remote)

	pathData := strings.Trim(r.FormValue("path"), "\r\n")
	if !strings.HasSuffix(pathData, "/") && pathData != "" && pathData != "/" {
		pathData = pathData + "/"
	}

	request.Path = pathData

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.MultipartForm.File["file"] == nil {
		http.Error(w, "File field 'file' is required", http.StatusBadRequest)
		return
	}

	resp := WriteFileResponse{}

	for _, file := range r.MultipartForm.File["file"] {
		fPath := request.Path + file.Filename

		statx, err := operations.StatJSON(r.Context(), tfs, fPath, &operations.ListJSONOpt{})
		if err == nil && statx != nil && !request.Overwrite {
			http.Error(w, "File already exists", http.StatusInternalServerError)
			return
		}

		fo, err := file.Open()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer fo.Close()

		memobj := object.NewMemoryObject(fPath, time.Now(), nil)

		ob, err := tfs.Put(context.Background(), fo, memobj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Files = append(resp.Files, WriteFileResponseItem{
			Name:     fPath,
			Size:     ob.Size(),
			MimeType: file.Header.Get("Content-Type"),
			ModTime:  ob.ModTime(r.Context()),
		})
	}

	common.JsonResponse(w, http.StatusOK, resp)
}

// Delete File
// @Summary Delete File
// @Description Delete File
// @Tags files
// @Accept json
// @Produce text/plain
// @Param remote body DeleteFileRequest true "Remote Configuration"
// @Success 200 {string} string "File deleted successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/delete [post]
func DeleteFile(w http.ResponseWriter, r *http.Request) {
	var request DeleteFileRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tfo, err := tfs.NewObject(r.Context(), request.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = operations.DeleteFile(r.Context(), tfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ByteResponse(w, http.StatusOK, "text/plain", []byte("File deleted successfully"))
}
