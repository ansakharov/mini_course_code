package src

import (
	"bytes"
	"io"
	"testing"
)

func TestCountingReader_Read(t *testing.T) {
	type fields struct {
		reader    io.Reader
		bytesRead int64
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          int
		wantErr       bool
		wantBytesRead int64
		wantContent   string
	}{
		{
			name:          "Basic read test",
			fields:        fields{reader: bytes.NewReader([]byte("Sample"))},
			args:          args{p: make([]byte, 6)},
			want:          6,
			wantErr:       false,
			wantBytesRead: 6,
			wantContent:   "sample",
		},
		{
			name:          "Read with smaller buffer",
			fields:        fields{reader: bytes.NewReader([]byte("Sample"))},
			args:          args{p: make([]byte, 4)},
			want:          4,
			wantErr:       false,
			wantBytesRead: 4,
			wantContent:   "samp",
		},
		{
			name:          "Read empty input",
			fields:        fields{reader: bytes.NewReader([]byte(""))},
			args:          args{p: make([]byte, 6)},
			want:          0,
			wantErr:       true,
			wantBytesRead: 0,
			wantContent:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cr := &CountingToLowerReaderImpl{
				Reader:         tt.fields.reader,
				TotalBytesRead: tt.fields.bytesRead,
			}
			got, err := cr.Read(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
			if cr.TotalBytesRead != tt.wantBytesRead {
				t.Errorf("Bytes read = %d, want %d", cr.TotalBytesRead, tt.wantBytesRead)
			}
			if string(tt.args.p[:got]) != tt.wantContent {
				t.Errorf("Content read = %s, want %s", string(tt.args.p[:got]), tt.wantContent)
			}
		})
	}
}

func TestReadAll(t *testing.T) {
	type args struct {
		cr      *CountingToLowerReaderImpl
		bufSize int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ReadAll with buffer size equal to content length",
			args: args{
				cr:      NewCountingReader(bytes.NewReader([]byte("Sample Text"))),
				bufSize: 10,
			},
			want:    "sample text",
			wantErr: false,
		},
		{
			name: "ReadAll with smaller buffer size",
			args: args{
				cr:      NewCountingReader(bytes.NewReader([]byte("Sample Text"))),
				bufSize: 4,
			},
			want:    "sample text",
			wantErr: false,
		},
		{
			name: "ReadAll with empty content",
			args: args{
				cr:      NewCountingReader(bytes.NewReader([]byte(""))),
				bufSize: 4,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.cr.ReadAll(tt.args.bufSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
