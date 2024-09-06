package files

import (
	"time"

	"github.com/rclone/rclone/fs/operations"
)

type RemotePathRequest struct {
	Remote RemoteConfiguration `json:"remote"`
	Path   string              `json:"path"`
}

type ReadFileRequest RemotePathRequest
type DeleteFileRequest RemotePathRequest

type ListFilesRequest struct {
	RemotePathRequest
	Recurse bool                   `json:"recurse"`
	Options operations.ListJSONOpt `json:"options"`
}

type FilterType string

const (
	FilterTypePrefix   FilterType = "prefix"
	FilterTypeSuffix   FilterType = "suffix"
	FilterTypeRegex    FilterType = "regex"
	FilterTypeWildcard FilterType = "wildcard"
)

type FilteredListFilesRequest struct {
	ListFilesRequest
	FilterType FilterType `json:"filterType"`
	Filter     string     `json:"filter"`
}

type ListFilesResponse struct {
	Files []operations.ListJSONItem `json:"files"`
	Total int64                     `json:"total"`
}

type ReadFileResponseItem struct {
	Name string `json:"name"`
	Size int64  `json:"size"`
}

type ReadFileResponse struct {
	Files []ReadFileResponseItem `json:"files"`
}

type WriteFileRequest struct {
	RemotePathRequest
	Overwrite bool   `json:"overwrite"`
	File      []byte `json:"file"`
}

type WriteFileResponse struct {
	Files []WriteFileResponseItem `json:"files"`
}

type WriteFileResponseItem struct {
	Name     string    `json:"name"`
	Size     int64     `json:"size"`
	MimeType string    `json:"mimeType"`
	ModTime  time.Time `json:"modTime"`
}

type BulkRenameFilesRequest struct {
	RemotePathRequest
	NameMap map[string]string `json:"nameMap"`
}

type BulkRenameFilesResponse struct {
	RenamedFiles map[string]string `json:"renamedFiles"`
	Errors       map[string]string `json:"errors"`
}
