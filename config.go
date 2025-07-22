package oss_aliyun

import (
	"fmt"
	"github.com/hdget/common/intf"
	"github.com/pkg/errors"
	"net/url"
)

type aliyunOssConfig struct {
	Region       string `mapstructure:"region"`
	Domain       string `mapstructure:"domain"`
	Bucket       string `mapstructure:"bucket"`
	AccessKey    string `mapstructure:"access_key"`
	AccessSecret string `mapstructure:"access_secret"`
}

const (
	configSection = "sdk.oss"
)

var (
	errInvalidConfig = errors.New("invalid oss provider config")
)

func newConfig(configProvider intf.ConfigProvider) (*aliyunOssConfig, error) {
	if configProvider == nil {
		return nil, errInvalidConfig
	}

	var c *aliyunOssConfig
	err := configProvider.Unmarshal(&c, configSection)
	if err != nil {
		return nil, err
	}

	if err := validateConfig(c); err != nil {
		return nil, err
	}

	return c, nil
}

func validateConfig(config *aliyunOssConfig) error {
	if _, err := url.Parse(config.Domain); err != nil {
		return fmt.Errorf("invalid oss domain")
	}

	if config.AccessKey == "" {
		return fmt.Errorf("oss access key is empty")
	}

	if config.AccessSecret == "" {
		return fmt.Errorf("oss access secret is empty")
	}

	if config.Bucket == "" {
		return fmt.Errorf("oss bucket is empty")
	}

	return nil
}
