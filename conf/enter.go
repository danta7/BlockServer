package conf

type Config struct {
	System System `yaml:"system"` // ip、端口、环境
	Log    Log    `yaml:"log"`
	DB     DB     `yaml:"db"`  // 读库
	DB1    DB     `yaml:"db1"` // 写库
	Jwt    Jwt    `yaml:"jwt"` // jwt
	Redis  Redis  `yaml:"redis"`
	Site   Site   `yaml:"site"`
	Email  Email  `yaml:"email"`
	QQ     QQ     `yaml:"qq"`
	QiNiu  QiNiu  `yaml:"qiNiu"`
	Ai     Ai     `yaml:"ai"`
	Upload Upload `yaml:"upload"`
}
