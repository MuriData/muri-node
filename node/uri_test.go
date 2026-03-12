package node

import "testing"

func TestExtractIPFSRef(t *testing.T) {
	tests := []struct {
		name string
		uri  string
		want string
	}{
		{"full ipfs URI", "ipfs://QmHash123", "QmHash123"},
		{"with subpath", "ipfs://QmHash123/path/to/file", "QmHash123/path/to/file"},
		{"with query param", "ipfs://QmHash123?type=raw", "QmHash123"},
		{"with subpath and query", "ipfs://QmHash123/sub/path?type=raw", "QmHash123/sub/path"},
		{"raw CID", "QmHash123", "QmHash123"},
		{"trailing slash", "ipfs://QmHash123/", "QmHash123"},
		{"whitespace", "  ipfs://QmHash123  ", "QmHash123"},
		{"empty", "", ""},
		{"only scheme", "ipfs://", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractIPFSRef(tt.uri)
			if got != tt.want {
				t.Errorf("extractIPFSRef(%q) = %q, want %q", tt.uri, got, tt.want)
			}
		})
	}
}

func TestIsRawBlockURI(t *testing.T) {
	tests := []struct {
		name string
		uri  string
		want bool
	}{
		{"raw type", "ipfs://QmHash?type=raw", true},
		{"no query", "ipfs://QmHash", false},
		{"other param", "ipfs://QmHash?format=dag", false},
		{"raw among others", "ipfs://QmHash?format=dag&type=raw", true},
		{"empty", "", false},
		{"partial match", "ipfs://QmHash?type=rawdata", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isRawBlockURI(tt.uri)
			if got != tt.want {
				t.Errorf("isRawBlockURI(%q) = %v, want %v", tt.uri, got, tt.want)
			}
		})
	}
}

func TestExtractRootCID(t *testing.T) {
	tests := []struct {
		name string
		uri  string
		want string
	}{
		{"simple CID", "ipfs://QmHash123", "QmHash123"},
		{"with subpath", "ipfs://QmHash123/path/to/file", "QmHash123"},
		{"with query", "ipfs://QmHash123?type=raw", "QmHash123"},
		{"with both", "ipfs://QmHash123/sub?type=raw", "QmHash123"},
		{"raw CID no scheme", "QmHash123", "QmHash123"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractRootCID(tt.uri)
			if got != tt.want {
				t.Errorf("extractRootCID(%q) = %q, want %q", tt.uri, got, tt.want)
			}
		})
	}
}
