package node

import "strings"

// extractIPFSRef extracts the full IPFS reference (CID or CID/path) from a URI,
// stripping the ipfs:// scheme and any query parameters.
// Supports formats: "ipfs://CID", "ipfs://CID/path/to/file", "ipfs://CID?type=raw", or raw CID.
// The returned value can be passed directly to IPFS Cat or BlockGet.
func extractIPFSRef(uri string) string {
	uri = strings.TrimSpace(uri)
	if strings.HasPrefix(uri, "ipfs://") {
		uri = strings.TrimPrefix(uri, "ipfs://")
	}
	// Strip query parameters (e.g. ?type=raw)
	if idx := strings.Index(uri, "?"); idx >= 0 {
		uri = uri[:idx]
	}
	uri = strings.TrimRight(uri, "/")
	if uri == "" {
		return ""
	}
	return uri
}

// isRawBlockURI returns true if the URI has ?type=raw, indicating the CID
// points to a raw IPFS block (e.g. a directory DAG node) that must be fetched
// via block/get instead of cat.
func isRawBlockURI(uri string) bool {
	uri = strings.TrimSpace(uri)
	idx := strings.Index(uri, "?")
	if idx < 0 {
		return false
	}
	query := uri[idx+1:]
	for _, param := range strings.Split(query, "&") {
		if param == "type=raw" {
			return true
		}
	}
	return false
}

// extractRootCID extracts just the root CID from a URI, stripping any subpath
// and query parameters. Use this for Pin/Unpin operations that apply to the whole DAG.
func extractRootCID(uri string) string {
	ref := extractIPFSRef(uri)
	if idx := strings.Index(ref, "/"); idx > 0 {
		return ref[:idx]
	}
	return ref
}
