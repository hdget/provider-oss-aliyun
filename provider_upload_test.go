package oss_aliyun

import "testing"

func Test_aliyunOssProvider_Upload(t *testing.T) {
	type args struct {
		dir      string
		filename string
		data     []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test_aliyunOssProvider_Upload",
			args: args{
				dir:      "xxx",
				filename: "example.txt",
				data:     []byte("example content"),
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &aliyunOssProvider{
				config: &aliyunOssConfig{
					Region:       "",
					Domain:       "",
					Bucket:       "",
					AccessKey:    "",
					AccessSecret: "",
				},
			}
			got, err := p.Upload(tt.args.dir, tt.args.filename, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Upload() got = %v, want %v", got, tt.want)
			}
		})
	}
}
