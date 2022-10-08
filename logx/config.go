package logx

type Config struct {
	Prefix string `yaml:"prefix" json:"prefix"`
	Level  string `yaml:"level" json:"level,default=info,options=debug|info|warn|error|alert"`
	Type   string `yaml:"type" json:"type,default=console,options=console|file"`
	Path   string `yaml:"path" json:"path,default=logs"`
}
