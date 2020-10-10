package gobump

import (
	"io"
	"testing"
)

func TestWalk(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name             string
		isDir            bool
		err              error
		target           string
		walkFuncError    error
		skips            []SkipFile
		wantErr          bool
		wantWalkFuncCall bool
	}{
		{
			name:             "happy case",
			isDir:            false,
			target:           "1.14",
			wantErr:          false,
			wantWalkFuncCall: true,
		},
		{
			name:   "skip - is dir",
			isDir:  true,
			target: "1.14",
		},
		{
			name:   "skip - not interesting",
			target: "1.14",
			skips: []SkipFile{
				func(path string) bool {
					return true
				},
			},
		},
		{
			name:    "pass error",
			err:     io.ErrUnexpectedEOF,
			target:  "1.14",
			wantErr: true,
		},
		{
			name:             "walk func error",
			target:           "1.14",
			walkFuncError:    io.ErrUnexpectedEOF,
			wantErr:          true,
			wantWalkFuncCall: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {

			isDir := dirDetectorFunc(func() bool {
				return tt.isDir
			})

			called := 0
			walkFunc := func(path, target string) error {
				if target != tt.target {
					t.Errorf("expected target %s, got %s", tt.target, target)
				}

				called++
				return tt.walkFuncError
			}

			tt.skips = append(tt.skips, func(path string) bool {
				if path != "/tmp" {
					t.Error("expected path /tmp , got", path)
				}
				return false
			})

			if err := walk("/tmp", isDir, tt.err, tt.target, walkFunc, tt.skips...); (err != nil) != tt.wantErr {
				t.Errorf("walk() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantWalkFuncCall != (called == 1) {
				t.Error("want walk func called once, got", called)
			}
		})
	}
}

func TestValidateTarget(t *testing.T) {

	tests := []struct {
		name    string
		target  string
		wantErr bool
	}{
		{
			name:   "1.14",
			target: "1.14",
		},
		{
			name:   "1.15",
			target: "1.15",
		},
		{
			name:   "2.0",
			target: "2.0",
		},
		{
			name:    "2",
			target:  "2",
			wantErr: true,
		},
		{
			name:    ".2",
			target:  ".2",
			wantErr: true,
		},
		{
			name:    "a.14",
			target:  "a.14",
			wantErr: true,
		},
		{
			name:    "1.a",
			target:  "1.a",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateTarget(tt.target); (err != nil) != tt.wantErr {
				t.Errorf("validateTarget() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
