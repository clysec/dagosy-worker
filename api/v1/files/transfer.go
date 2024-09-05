package files

import (
	"encoding/json"
	"net/http"

	"github.com/clysec/dagosy-worker/common"
	"github.com/rclone/rclone/fs/operations"
	"github.com/rclone/rclone/fs/sync"
)

// Copy File
// @Summary Copy File
// @Description Copy File
// @Tags files
// @Accept json
// @Produce json
// @Param copyFileRequest body CopyFileRequest true "Copy File Request"
// @Success 200 {string} string "File copied successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/copy [post]
func CopyFile(w http.ResponseWriter, r *http.Request) {
	var request CopyFileRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcFs, err := request.SourceRemote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFs, err := request.DestinationRemote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = operations.CopyFile(r.Context(), srcFs, dstFs, request.DestinationPath, request.SourcePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ByteResponse(w, http.StatusOK, "text/plain", []byte("File copied successfully"))
}

// Move File
// @Summary Move File
// @Description Move File
// @Tags files
// @Accept json
// @Produce json
// @Param moveFileRequest body MoveFileRequest true "Move File Request"
// @Success 200 {string} string "File moved successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/moveFile [post]
func MoveFile(w http.ResponseWriter, r *http.Request) {
	var request MoveFileRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcFs, err := request.SourceRemote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFs, err := request.DestinationRemote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = operations.MoveFile(r.Context(), srcFs, dstFs, request.DestinationPath, request.SourcePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ByteResponse(w, http.StatusOK, "text/plain", []byte("File moved successfully"))
}

// Move Backup Dir
// @Summary Move Backup Dir
// @Description Move a file to a backup directory
// @Tags files
// @Accept json
// @Produce json
// @Param moveBackupDirRequest body MoveBackupDirRequest true "Move Backup Dir Request"
// @Success 200 {string} string "Backup directory moved successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/moveBackupDir [post]
func MoveBackupDir(w http.ResponseWriter, r *http.Request) {
	var request MoveBackupDirRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcFs, err := request.SourceRemote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFs, err := request.DestinationRemote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFsDir, err := dstFs.NewObject(r.Context(), request.DestinationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	operations.MoveBackupDir(r.Context(), srcFs, dstFsDir)
}

// Copy URL
// @Summary Copy URL to filesystem
// @Description Copy URL to filesystem
// @Tags files
// @Accept json
// @Produce json
// @Param copyURLRequest body CopyURLRequest true "Copy URL Request"
// @Success 200 {string} string "File copied successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/copyUrl [post]
func CopyURL(w http.ResponseWriter, r *http.Request) {
	var request CopyURLRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dst, err := operations.CopyURL(r.Context(), tfs, request.Path, request.URL, request.AutoFilename, request.DstFilenameFromHeader, request.NoClobber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respitem := CopyURLResponse{
		WriteFileResponseItem: WriteFileResponseItem{
			Name:    dst.Fs().Name(),
			Size:    dst.Size(),
			ModTime: dst.ModTime(r.Context()),
		},
	}

	common.JsonResponse(w, http.StatusOK, respitem)
}

// Remove Dir
// @Summary Remove Empty Directory
// @Description Removes an empty directory
// @Tags files
// @Accept json
// @Produce json
// @Param rmdirsRequest body RmdirRequest true "Remove Directory Request"
// @Success 200 {string} string "Removed successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/removeDir [post]
func Rmdir(w http.ResponseWriter, r *http.Request) {
	var request RmdirsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = operations.Rmdir(r.Context(), tfs, request.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ByteResponse(w, http.StatusOK, "text/plain", []byte("Removed successfully"))
}

// Remove Dirs (force)
// @Summary Remove Directories
// @Description Recursively remove directories
// @Tags files
// @Accept json
// @Produce json
// @Param rmdirsRequest body RmdirsRequest true "Remove Directories Request"
// @Success 200 {string} string "Removed successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/removeDirs [post]
func Rmdirs(w http.ResponseWriter, r *http.Request) {
	var request RmdirsRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tfs, err := request.Remote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = operations.Rmdirs(r.Context(), tfs, request.Path, request.LeaveRoot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ByteResponse(w, http.StatusOK, "text/plain", []byte("Removed successfully"))
}

// Sync MoveDir
// @Summary Sync MoveDir
// @Description Sync MoveDir
// @Tags files
// @Accept json
// @Produce json
// @Param syncMoveDirRequest body SyncMoveDirRequest true "Sync MoveDir Request"
// @Success 200 {string} string "Synced successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/syncMoveDir [post]
func SyncMoveDir(w http.ResponseWriter, r *http.Request) {
	var request SyncMoveDirRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcFs, err := request.SourceRemote.GetFilesystemAtPath(r, w, request.SourcePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFs, err := request.DestinationRemote.GetFilesystemAtPath(r, w, request.DestinationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = sync.MoveDir(r.Context(), srcFs, dstFs, request.DeleteEmptySrcDirs, request.CopyEmptyDirs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ByteResponse(w, http.StatusOK, "text/plain", []byte("Synced successfully"))
}

// Sync CopyDir
// @Summary Sync CopyDir
// @Description Sync CopyDir
// @Tags files
// @Accept json
// @Produce json
// @Param syncCopyDirRequest body SyncCopyDirRequest true "Sync CopyDir Request"
// @Success 200 {string} string "Synced successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/syncCopyDir [post]
func SyncCopyDir(w http.ResponseWriter, r *http.Request) {
	var request SyncCopyDirRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcFs, err := request.SourceRemote.GetFilesystemAtPath(r, w, request.SourcePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFs, err := request.DestinationRemote.GetFilesystemAtPath(r, w, request.DestinationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = sync.CopyDir(r.Context(), srcFs, dstFs, request.CopyEmptyDirs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ByteResponse(w, http.StatusOK, "text/plain", []byte("Synced successfully"))
}

// Sync Sync
// @Summary Sync
// @Description Sync
// @Tags files
// @Accept json
// @Produce json
// @Param syncRequest body SyncRequest true "Sync Request"
// @Success 200 {string} string "Synced successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/sync [post]
func Sync(w http.ResponseWriter, r *http.Request) {
	var request SyncRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcFs, err := request.SourceRemote.GetFilesystemAtPath(r, w, request.SourcePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFs, err := request.DestinationRemote.GetFilesystemAtPath(r, w, request.DestinationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = sync.Sync(r.Context(), srcFs, dstFs, true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	common.ByteResponse(w, http.StatusOK, "text/plain", []byte("Synced successfully"))
}

// Check Equal
// @Summary Check Equal
// @Description Check if two files are equal
// @Tags files
// @Accept json
// @Produce json
// @Param checkEqualRequest body CheckEqualRequest true "Check Equal Request"
// @Success 200 {string} string "Checked successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/files/checkEqual [post]
func CheckEqual(w http.ResponseWriter, r *http.Request) {
	var request CheckEqualRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcFs, err := request.SourceRemote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFs, err := request.DestinationRemote.GetFilesystem(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	srcFsFile, err := srcFs.NewObject(r.Context(), request.SourcePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dstFsFile, err := dstFs.NewObject(r.Context(), request.DestinationPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	equal, _, err := operations.CheckHashes(r.Context(), srcFsFile, dstFsFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := CheckEqualResponse{
		Equal: equal,
	}

	common.JsonResponse(w, http.StatusOK, resp)
}
