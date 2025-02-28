package pkg

import (
	"fmt"
	"github.com/hdget/common/intf"
	"github.com/pkg/errors"
	"net/url"
)

type aliyunOssConfig struct {
	domain       string `mapstructure:"domain"`
	endpoint     string `mapstructure:"endpoint"`
	bucket       string `mapstructure:"bucket"`
	accessKey    string `mapstructure:"access_key"`
	accessSecret string `mapstructure:"access_secret"`
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
	if _, err := url.Parse(config.domain); err != nil {
		return fmt.Errorf("invalid oss domain")
	}

	if config.endpoint == "" {
		return fmt.Errorf("oss endpoint is empty")
	}

	if config.accessKey == "" {
		return fmt.Errorf("oss access key is empty")
	}

	if config.accessSecret == "" {
		return fmt.Errorf("oss access secret is empty")
	}

	if config.bucket == "" {
		return fmt.Errorf("oss bucket is empty")
	}

	return nil
}
