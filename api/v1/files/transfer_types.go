package files

type SourceDestinationRequest struct {
	SourceRemote      RemoteConfiguration `json:"sourceRemote"`
	DestinationRemote RemoteConfiguration `json:"destinationRemote"`
	SourcePath        string              `json:"sourcePath"`
	DestinationPath   string              `json:"destinationPath"`
}

type RmdirRequest RemotePathRequest
type RmdirsRequest struct {
	RemotePathRequest
	LeaveRoot bool `json:"leaveRoot"`
}

type CopyFileRequest SourceDestinationRequest
type MoveFileRequest SourceDestinationRequest
type MoveBackupDirRequest SourceDestinationRequest

type CopyURLRequest struct {
	RemotePathRequest
	URL                   string `json:"url"`
	AutoFilename          bool   `json:"autoFilename"`
	DstFilenameFromHeader bool   `json:"dstFilenameFromHeader"`
	NoClobber             bool   `json:"noClobber"`
}

type CopyURLResponse struct {
	WriteFileResponseItem
}

type SyncRequest struct {
	SourceDestinationRequest
	CopyEmptyDirs bool `json:"copyEmptyDirs"`
}

type SyncCopyDirRequest SyncRequest

type SyncMoveDirRequest struct {
	SyncRequest
	DeleteEmptySrcDirs bool `json:"deleteEmptySrcDirs"`
}

type CheckEqualRequest SourceDestinationRequest

type CheckEqualResponse struct {
	Equal bool `json:"equal"`
}
