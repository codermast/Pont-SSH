package constant

import (
	"github.com/vrischmann/userdir"
	"path/filepath"
)

var DatabaseFilePath = filepath.Join(userdir.GetConfigHome(), "PontSSH", "pontssh.db")

var ConfigFilePath = filepath.Join(userdir.GetConfigHome(), "PontSSH", "config.yaml")
