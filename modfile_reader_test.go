package gobump

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestReadWriter(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		expectedPath string
		target       string
		wantErr      bool
	}{
		{
			name:         "testdata/go.mod.txt",
			path:         "testdata/go.mod.txt",
			expectedPath: "testdata/go.mod.txt.expected",
			target:       "1.15",
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadWriter(tt.path, tt.target); (err != nil) != tt.wantErr {
				t.Errorf("ReadWriter() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := ioutil.ReadFile(tt.path)
			if err != nil {
				t.Errorf("failed to read file %s: %v", tt.path, err)
			}

			expected, err := ioutil.ReadFile(tt.expectedPath)
			if err != nil {
				t.Errorf("failed to read file %s: %v", tt.expectedPath, err)
			}

			if !bytes.Equal(got, expected) {
				t.Errorf("got file %s and expected %s aren't equal", tt.path, tt.expectedPath)
			}
		})
	}
}
