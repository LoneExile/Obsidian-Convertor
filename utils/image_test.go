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
		{name: "test2", args: args{srcPath: "../example/SecondBrain/Assets/image/test_no_exists.jpg", dstPath: "../test/test.jpg"}, wantErr: true},
		{name: "test3", args: args{srcPath: "../example/SecondBrain/Assets/image/test.jpg", dstPath: ""}, wantErr: true},
		// no way that `io.Copy` returns an error?
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
		{name: "test2", args: args{srcPath: "../example/SecondBrain/Assets/image/test.jpg", dstPath: "../test/test.webp", format: FormatJPG, quality: QualityLow}, wantErr: false},
		{name: "test2", args: args{srcPath: "../example/SecondBrain/Assets/image/test.jpg", dstPath: "../test/test.webp", format: FormatPNG, quality: QualityHigh}, wantErr: false},
		{name: "test3", args: args{srcPath: "../example/SecondBrain/Assets/image/test.jpg", dstPath: "../test/test.webp", format: FormatAVIF, quality: QualityMedium}, wantErr: false},
		{name: "test4", args: args{srcPath: "../example/SecondBrain/Assets/image/test_no_exists.jpg", dstPath: "../test/test.webp", format: FormatSame, quality: QualityMedium}, wantErr: true},
		{name: "test6", args: args{srcPath: "../example/SecondBrain/Assets/image/test.jpg", dstPath: "", format: FormatSame, quality: QualityMedium}, wantErr: true},
		{name: "test7", args: args{srcPath: "../example/SecondBrain/Assets/image/test.jpg", dstPath: "../test/", format: FormatAVIF, quality: QualityMedium}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ConvertImage(tt.args.srcPath, tt.args.dstPath, tt.args.format, tt.args.quality); (err != nil) != tt.wantErr {
				t.Errorf("ConvertImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
