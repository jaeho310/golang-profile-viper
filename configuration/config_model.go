package configuration

var RuntimeConf = RuntimeConfig{}

type RuntimeConfig struct {
	Datasource Datasource `yaml:"datasource"`
	Server     Server     `yaml:"server"`
}

type Datasource struct {
	DbType   string `yaml:"dbType"`
	Url      string `yaml:"url"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
}

type Server struct {
	Port int `yaml:"port"`
}
