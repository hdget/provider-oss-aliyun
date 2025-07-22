package oss_aliyun

import (
	"bytes"
	"context"
	"crypto/rand"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/hdget/common/intf"
	"github.com/hdget/common/types"
	"log"
	"math"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type aliyunOssProvider struct {
	config *aliyunOssConfig
}

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

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

func (p *aliyunOssProvider) Upload(dir, filename string, data []byte) (string, error) {
	objectKey := p.getObjectKey(dir, filename)

	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr(p.config.Bucket), // 存储空间名称
		Key:          oss.Ptr(objectKey),       // 存储对象路径
		Body:         bytes.NewReader(data),
		StorageClass: oss.StorageClassStandard, // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPublicRead,  // 指定对象的访问权限
		Metadata: map[string]string{
			"yourMetadataKey1": "yourMetadataValue1", // 设置对象的元数据
		},
	}

	// 执行上传对象的请求
	_, err := oss.NewClient(oss.LoadDefaultConfig()).PutObject(context.TODO(), putRequest)
	if err != nil {
		log.Fatalf("failed to put object from file %v", err)
	}

	return objectKey, nil
}

func (p *aliyunOssProvider) getObjectKey(dir, filename string) string {
	strDate := time.Now().Format("20060102")
	year, month, day := strDate[:4], strDate[4:6], strDate[6:8]
	return path.Join(dir, year, month, day, generateSafeFileName(filename))
}

func generateSafeFileName(filename string) string {
	safeFileName := filepath.Base(filename)                   // 移除路径分隔符
	safeFileName = strings.ReplaceAll(safeFileName, " ", "_") // 替换空格等特殊字符

	ext := filepath.Ext(safeFileName)
	name := safeFileName[0 : len(safeFileName)-len(ext)]

	return fmt.Sprintf("%s_%s%s", name, randStr(6), ext) // 防止相同文件名被覆盖
}

func randStr(size int) string { // 高效随机字符串
	chars := []rune(alphabet)
	mask := getMask(len(chars))
	// estimate how many random bytes we will need for the ID, we might actually need more but this is tradeoff
	// between average case and worst case
	ceilArg := 1.6 * float64(mask*size) / float64(len(alphabet))
	step := int(math.Ceil(ceilArg))

	id := make([]rune, size)
	bytes := make([]byte, step)
	for j := 0; ; {
		_, _ = rand.Read(bytes)
		for i := 0; i < step; i++ {
			currByte := bytes[i] & byte(mask)
			if currByte < byte(len(chars)) {
				id[j] = chars[currByte]
				j++
				if j == size {
					return string(id[:size])
				}
			}
		}
	}
}

func getMask(alphabetSize int) int {
	for i := 1; i <= 8; i++ {
		mask := (2 << uint(i)) - 1
		if mask >= alphabetSize-1 {
			return mask
		}
	}
	return 0
}
