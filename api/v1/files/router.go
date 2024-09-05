package files

import "github.com/gorilla/mux"

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/list", ListFiles).Methods("POST")
	router.HandleFunc("/read", ReadFile).Methods("POST")
	router.HandleFunc("/write", WriteFile).Methods("POST")
	router.HandleFunc("/delete", DeleteFile).Methods("POST")

	router.HandleFunc("/copy", CopyFile).Methods("POST")
	router.HandleFunc("/moveFile", MoveFile).Methods("POST")
	router.HandleFunc("/moveBackupDir", MoveBackupDir).Methods("POST")

	router.HandleFunc("/copyURL", CopyURL).Methods("POST")
	router.HandleFunc("/removeDir", Rmdir).Methods("POST")
	router.HandleFunc("/removeDirs", Rmdirs).Methods("POST")

	router.HandleFunc("/sync", Sync).Methods("POST")
	router.HandleFunc("/syncCopyDir", SyncCopyDir).Methods("POST")
	router.HandleFunc("/syncMoveDir", SyncMoveDir).Methods("POST")
	router.HandleFunc("/checkEqual", CheckEqual).Methods("POST")
}
