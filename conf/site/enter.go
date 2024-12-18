package site

type SiteInfo struct {
	Title string `json:"title" yaml:"title"` // 网站的标题
	Logo  string `json:"logo" yaml:"logo"`   // 网站的logo
	Beian string `json:"beian" yaml:"beian"` // 网站的编号
	Mode  int8   `json:"mode" yaml:"mode"`   // 网站的运行模式 1 社区模式 2 博客模式

}

type Project struct {
	Title   string `json:"title" yaml:"title"`
	Icon    string `json:"icon" yaml:"icon"`
	WebPath string `json:"webPath" yaml:"webPath"`
}

type Seo struct {
	Keywords    string `json:"keywords" yaml:"keywords"`
	Description string `json:"description" yaml:"description"`
}

type About struct {
	SiteDate string `json:"siteDate" yaml:"siteDate"` // 建站时间 年月日
	QQ       string `json:"QQ" yaml:"qq"`             // qq 二维码
	Wechat   string `json:"wechat" yaml:"wechat"`     // 微信二维码
	Gitee    string `json:"gitee" yaml:"gitee"`
	Bilibili string `json:"bilibili" yaml:"bilibili"`
	Github   string `json:"github" yaml:"github"`
}

type Login struct {
	QQLogin          bool `json:"QQLogin" yaml:"qqLogin"`
	UsernamePwdLogin bool `json:"UsernamePwdLogin" yaml:"usernamePwdLogin"`
	EmailLogin       bool `json:"EmailLogin" yaml:"emailLogin"`
	Captcha          bool `json:"Captcha" yaml:"captcha"`
}

type ComponentInfo struct {
	Title  string `json:"title" yaml:"title"`
	Enable bool   `json:"enable" yaml:"enable"`
}

type IndexRight struct {
	List []ComponentInfo `json:"list" yaml:"list"`
}

type Article struct {
	NoExamine bool `yaml:"noExamine" json:"noExamine"` // 免审核
}
