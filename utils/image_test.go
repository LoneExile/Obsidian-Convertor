package utils

import "testing"

func TestCopyImage(t *testing.T) {
	type args struct {
		srcPath string
		dstPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{srcPath: "../example/SecondBrain/Assets/image/test.jpg", dstPath: "../test/test.jpg"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CopyImage(tt.args.srcPath, tt.args.dstPath); (err != nil) != tt.wantErr {
				t.Errorf("CopyImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConvertImage(t *testing.T) {
	type args struct {
		srcPath string
		dstPath string
		format  ImageFormat
		quality ImageQuality
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{srcPath: "../example/SecondBrain/Assets/image/test.jpg", dstPath: "../test/test.webp", format: FormatWEBP, quality: QualityMedium}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertImage(tt.args.srcPath, tt.args.dstPath, tt.args.format, tt.args.quality); (err != nil) != tt.wantErr {
				t.Errorf("ConvertImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
