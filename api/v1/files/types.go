package files

import (
	"net/http"

	"github.com/rclone/rclone/fs"
	fscfg "github.com/rclone/rclone/fs/config"
)

var ValidRemotes = []string{
	"alias",
	"azureblob",
	"azurefiles",
	"b2",
	"box",
	"cache",
	"chunker",
	"combine",
	"compress",
	"crypt",
	"drive",
	"dropbox",
	"fichier",
	"filefabric",
	"filescom",
	"ftp",
	"gofile",
	"googlecloudstorage",
	"googlephotos",
	"hasher",
	"hdfs",
	"hidrive",
	"http",
	"imagekit",
	"internetarchive",
	"jottacloud",
	"koofr",
	"linkbox",
	"local",
	"mailru",
	"mega",
	"memory",
	"netstorage",
	"onedrive",
	"opendrive",
	"oracleobjectstorage",
	"pcloud",
	"pikpak",
	"pixeldrain",
	"premiumizeme",
	"protondrive",
	"putio",
	"qingstor",
	"quatrix",
	"s3",
	"seafile",
	"sftp",
	"sharefile",
	"sia",
	"smb",
	"storj",
	"sugarsync",
	"swift",
	"ulozto",
	"union",
	"uptobox",
	"webdav",
	"yandex",
	"zoho",
}

type RemoteConfiguration struct {
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	Parameters map[string]interface{} `json:"parameters"`
	Options    fscfg.UpdateRemoteOpt  `json:"options"`
}

func (rc *RemoteConfiguration) CreateRemote(r *http.Request, w http.ResponseWriter) error {
	_, err := fscfg.CreateRemote(r.Context(), rc.Name, rc.Type, rc.Parameters, rc.Options)
	return err
}

func (rc *RemoteConfiguration) GetFilesystem(r *http.Request, w http.ResponseWriter) (fs.Fs, error) {
	err := rc.CreateRemote(r, w)
	if err != nil {
		return nil, err
	}

	return fs.NewFs(r.Context(), rc.Name+":")
}

func (rc *RemoteConfiguration) GetFilesystemAtPath(r *http.Request, w http.ResponseWriter, path string) (fs.Fs, error) {
	err := rc.CreateRemote(r, w)
	if err != nil {
		return nil, err
	}

	return fs.NewFs(r.Context(), rc.Name+":"+path)
}
