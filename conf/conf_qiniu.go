package conf

// 对象存储
type QiNiu struct {
	Enable    bool   `json:"enable" yaml:"enable"`
	AccessKey string `json:"accessKey" yaml:"accessKey"`
	SecretKey string `json:"secretKey" yaml:"secretKey"`
	Bucket    string `json:"bucket" yaml:"bucket"` // 存储桶
	Uri       string `json:"uri" yaml:"uri"`
	Region    string `json:"region" yaml:"region"`
	Prefix    string `json:"prefix" yaml:"prefix"` // 上传的目录
	Size      int    `json:"size" yaml:"size"`     // 大小限制 单位mb
	Expiry    int    `json:"expiry" yaml:"expiry"` // token过期时间 s
}
