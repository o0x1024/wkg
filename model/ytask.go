package model


type Redis struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	CharSet  string `yaml:"charset"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	ParseTime	string 	`yaml:"parseTime"`
	Loc		string `yaml:"loc"`
}
