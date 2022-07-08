package configs

type Keys struct {
	Secretkey string `yaml: "secretkey"`
	Salt      string `yaml: "salt"`
}