package oss_aliyun

import (
	"reflect"
	"testing"
)

func Test_aliyunOssProvider_GetPresignedURL(t *testing.T) {
	type args struct {
		dir         string
		filename    string
		contentType string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   map[string]string
		wantErr bool
	}{
		{
			name: "Test_aliyunOssProvider_GetPresignedURL",
			args: args{
				dir:         "xxx",
				filename:    "2.txt",
				contentType: "text/plain",
			},
			want:    "",
			want1:   nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &aliyunOssProvider{
				config: &aliyunOssConfig{
					Region:       "",
					Bucket:       "",
					AccessKey:    "",
					AccessSecret: "",
				},
			}
			got, got1, err := p.GetPresignedURL(tt.args.dir, tt.args.filename, tt.args.contentType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPresignedURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetPresignedURL() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPresignedURL() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
