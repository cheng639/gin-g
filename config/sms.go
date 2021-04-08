package config

type SMS struct {
	AccessId  string `mapstructure:"access-id" json:"accessId" yaml:"access-id"`
	AccessKey string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`
	SignName  string `mapstructure:"sign-name" json:"signName" yaml:"sign-name"`
	TemplateCode  string `mapstructure:"template-code" json:"stemplateCode" yaml:"template-code"`
	RegionId  string `mapstructure:"region-id" json:"regionId" yaml:"region-id"`
}