package oss_aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"hash"
	"io"
	"time"
)

type aliyunOssPolicy struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type AliyunOssSignature struct {
	AccessKeyId string
	Host        string
	Expire      int64
	Signature   string
	Directory   string
	Policy      string
}

const (
	defaultExpireTime = 600
)

// GenSignature 生成oss直传token
func (p *aliyunOssProvider) GenSignature(dir string) (*AliyunOssSignature, error) {
	expiresIn := time.Now().Unix() + defaultExpireTime
	policyData, err := p.getPolicyData(dir, expiresIn)
	if err != nil {
		return nil, err
	}

	// create post policy json
	stdPolicyData := base64.StdEncoding.EncodeToString(policyData)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(p.config.AccessSecret))
	_, err = io.WriteString(h, stdPolicyData)
	if err != nil {
		return nil, err
	}

	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return &AliyunOssSignature{
		AccessKeyId: p.config.AccessKey,
		Host:        p.config.Domain,
		Expire:      expiresIn,
		Signature:   signedStr,
		Directory:   dir,
		Policy:      stdPolicyData,
	}, nil
}

func (a *aliyunOssProvider) getPolicyData(dir string, expiresIn int64) ([]byte, error) {
	strExpireTime := time.Unix(expiresIn, 0).UTC().Format("2006-01-02T15:04:05Z")

	// 指定此次上传的文件名必须以user-dir开头
	condition := []string{"starts-with", "$key", dir}
	config := aliyunOssPolicy{
		Expiration: strExpireTime,
		Conditions: [][]string{
			condition,
		},
	}

	// calculate signature
	data, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	return data, nil
}
