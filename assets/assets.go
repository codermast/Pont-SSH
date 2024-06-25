package assets

import _ "embed"

//go:embed Pontssh.sql
var InitSql string

//go:embed config.yaml
var Config []byte
