package fs

import (
	"bytes"
	"os"
	"testing"
)

func TestFileSystem(t *testing.T) {
	saveEnv := os.Getenv(useDirEnv)
	t.Cleanup(func() {
		_ = os.Setenv(useDirEnv, saveEnv)
	})

	testCases := []struct {
		name            string
		useDir          bool
		canAccessHidden bool
	}{
		{
			name:            "uses embed.FS",
			useDir:          false,
			canAccessHidden: false,
		},
		{
			name:            "uses os.DirFS",
			useDir:          true,
			canAccessHidden: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.useDir {
				if err := os.Setenv(useDirEnv, "1"); err != nil {
					t.Fatalf("cannot set %q env: %v", useDirEnv, err)
					defer os.Unsetenv(useDirEnv)
				}

				fsys := FileSystem()

				buf := &bytes.Buffer{}
				f1, err := fsys.Open("test.txt")
				if err != nil {
					t.Errorf("cannot open test.txt: %v", err)
				}
				defer f1.Close()

				if _, err := buf.ReadFrom(f1); err != nil || buf.String() != "THIS IS A TEST!" {
					t.Errorf("could not read text or text does not match: %v", err)
				}

				f2, err := fsys.Open(".hidden.txt")
				if err != nil {
					if tc.canAccessHidden {
						t.Fatalf("expected to open .hidden.txt, but it failed: %v", err)
					}
					if !os.IsNotExist(err) {
						t.Fatalf("expected os.ErrNotExist, got: %v", err)
					}
				}
				defer f2.Close()
			}
		})
	}
}
