package oss_aliyun

import (
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/hdget/common/intf"
	"github.com/hdget/common/types"
)

type aliyunOssProvider struct {
	config             *aliyunOssConfig
	allowContentTypes  []string
	maxFileSize        int64
	signatureExpiresIn int64
}

const (
	defaultSignatureExpiresIn = 180                      // 上传签名默认失效时间, 3分钟
	defaultMaxFileSize        = int64(100 * 1024 * 1024) // 上传文件的最大尺寸, 100M
)

var (
	// ImageContentTypes 图像类
	ImageContentTypes = []string{
		"image/jpeg",
		"image/png",
		"image/gif",
		"image/bmp",
		"image/webp",
		"image/svg+xml",
		"image/webp",
		"image/tiff",
		"image/vnd.microsoft.icon",
	}

	// VideoContentTypes 视频类
	VideoContentTypes = []string{
		"video/mp4",
		"video/mpeg",
		"video/ogg",
		"video/webm",
		"video/quicktime",
		"video/x-msvideo",
		"video/x-ms-wmv",
	}

	// DocumentContentTypes 文档类
	DocumentContentTypes = []string{
		"text/plain",
		"application/vnd.ms-excel",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		"application/pdf",
		"application/msword",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		"application/vnd.ms-powerpoint",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation",
		"text/html",
		"application/json",
	}

	// ZipContentTypes 压缩类
	ZipContentTypes = []string{
		"application/zip",
		"application/gzip",
		"application/x-tar",
		"application/x-rar-compressed",
	}
)

func New(configProvider intf.ConfigProvider) (intf.OssProvider, error) {
	config, err := newConfig(configProvider)
	if err != nil {
		return nil, err
	}

	return &aliyunOssProvider{
		config:             config,
		allowContentTypes:  make([]string, 0),
		maxFileSize:        defaultMaxFileSize,
		signatureExpiresIn: defaultSignatureExpiresIn,
	}, nil
}

func (p *aliyunOssProvider) GetCapability() types.Capability {
	return Capability
}

func (p *aliyunOssProvider) SetContentTypes(contentTypes []string) *aliyunOssProvider {
	if len(contentTypes) > 0 {
		p.allowContentTypes = contentTypes
	}
	return p
}

func (p *aliyunOssProvider) SetMaxFileSize(size int64) *aliyunOssProvider {
	if size > 0 {
		p.maxFileSize = size
	}
	return p
}

func (p *aliyunOssProvider) SetSignatureExpires(expiresIn int64) *aliyunOssProvider {
	if expiresIn > 0 {
		p.signatureExpiresIn = expiresIn
	}
	return p
}

func (p *aliyunOssProvider) newOSSClient() *oss.Client {
	// 构建凭证提供者
	credProvider := credentials.NewStaticCredentialsProvider(p.config.AccessKey, p.config.AccessSecret)

	// 创建OSS配置
	ossCfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credProvider).
		WithRegion(p.config.Region) // region: cn-shanghai, 不需要带oss

	return oss.NewClient(ossCfg)
}
