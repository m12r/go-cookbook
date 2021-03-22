package fs

import "testing"

func TestEmbedString(t *testing.T) {
	if version != "v1.2.3" {
		t.Errorf("expected \"v1.2.3\", got %q", version)
	}
}
