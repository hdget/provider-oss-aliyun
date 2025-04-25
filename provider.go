package oss_aliyun

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/hdget/common/intf"
	"github.com/hdget/common/types"
	"path"
	"time"
)

type aliyunOssProvider struct {
	config *aliyunOssConfig
}

const (
	defaultMiddleDirFormat = "060102" // year,month,day
)

func New(configProvider intf.ConfigProvider, logger intf.LoggerProvider) (intf.OssProvider, error) {
	config, err := newConfig(configProvider)
	if err != nil {
		return nil, err
	}

	return &aliyunOssProvider{config: config}, nil
}

func (p *aliyunOssProvider) GetCapability() types.Capability {
	return Capability
}

func (p *aliyunOssProvider) Upload(rootDir, filename string, data []byte) (string, error) {
	// 获取存储空间
	client, err := oss.New(p.config.Endpoint, p.config.AccessKey, p.config.AccessSecret)
	if err != nil {
		return "", err
	}

	buk, err := client.Bucket(p.config.Bucket)
	if err != nil {
		return "", err
	}

	absPath := path.Join(rootDir, p.getMiddleDir(), filename)

	// 上传Byte数组
	err = buk.PutObject(absPath, bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	return absPath, nil
}

func (p *aliyunOssProvider) getMiddleDir() string {
	s := time.Now().Format(defaultMiddleDirFormat)
	return path.Join(s[:2], s[2:4], s[4:6])
}
