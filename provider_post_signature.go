package oss_aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"time"
)

type AliyunOssSignature struct {
	AccessKeyId string
	Host        string
	ExpireIn    int64
	Signature   string
	Directory   string
	Policy      string
}

// GetPostSignature 生成oss直传post签名
func (p *aliyunOssProvider) GetPostSignature(dir string) (*AliyunOssSignature, error) {
	policyBase64, policySigned, err := p.generatePolicy(dir, defaultSignatureExpiresIn)
	if err != nil {
		return nil, err
	}

	return &AliyunOssSignature{
		AccessKeyId: p.config.AccessKey,
		Host:        p.config.Domain,
		ExpireIn:    defaultSignatureExpiresIn,
		Signature:   policySigned,
		Directory:   dir,
		Policy:      policyBase64,
	}, nil
}

// generatePolicy 生成访问策略
func (p *aliyunOssProvider) generatePolicy(dir string, expiresIn int64) (string, string, error) {
	// 定义策略
	policy := map[string]any{
		// 多少秒后签名过期
		"expiration": time.Now().Add(time.Duration(expiresIn) * time.Second).Format("2006-01-02T15:04:05Z"),
		"conditions": [][]any{
			{"starts-with", "$key", dir},                    // 限制上传目录， 上传的文件名必须以dir开头
			{"content-length-range", 1, defaultMaxFileSize}, // 文件大小限制
		},
	}

	policyJSON, err := json.Marshal(policy)
	if err != nil {
		return "", "", err
	}

	policyBase64 := base64.StdEncoding.EncodeToString(policyJSON)

	// 生成HMAC-SHA1签名
	h := hmac.New(sha1.New, []byte(p.config.AccessSecret))
	h.Write([]byte(policyBase64))
	policySigned := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return policyBase64, policySigned, nil
}
