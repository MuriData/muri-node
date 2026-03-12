package chain

import (
	"errors"
	"testing"
)

func TestIsNonceTooLow(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nonce too low", errors.New("nonce too low"), true},
		{"uppercase", errors.New("Nonce Too Low"), true},
		{"wrapped", errors.New("execution reverted: nonce too low for account"), true},
		{"unrelated", errors.New("insufficient funds"), false},
		{"empty", errors.New(""), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNonceTooLow(tt.err)
			if got != tt.want {
				t.Errorf("isNonceTooLow(%q) = %v, want %v", tt.err, got, tt.want)
			}
		})
	}
}

func TestDecodeRevertData(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{"empty", nil, "unknown"},
		{"proof invalid selector", []byte{0x7f, 0xcd, 0xd1, 0xf4}, "ProofInvalid()"},
		{"public input selector", []byte{0xa5, 0x4f, 0x8e, 0x27}, "PublicInputNotInField()"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := decodeRevertData(tt.data)
			if got != tt.want {
				t.Errorf("decodeRevertData = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestDecodeRevertData_ErrorString(t *testing.T) {
	// Build ABI-encoded Error(string) with message "test error"
	// Selector: 0x08c379a0
	// Offset:   32 (0x20)
	// Length:   10 (0x0a)
	// Data:     "test error"
	data := make([]byte, 100)
	data[0] = 0x08
	data[1] = 0xc3
	data[2] = 0x79
	data[3] = 0xa0
	// offset = 32
	data[35] = 0x20
	// length = 10
	data[67] = 0x0a
	// "test error"
	copy(data[68:], "test error")

	got := decodeRevertData(data)
	if got != "test error" {
		t.Errorf("decodeRevertData = %q, want %q", got, "test error")
	}
}

func TestExtractHexFromError(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int // expected length of returned bytes, 0 means nil
	}{
		{"with hex", "execution reverted: 0x08c379a0abcd", 6},
		{"no hex", "some error without hex", 0},
		{"short hex", "error 0xab", 0}, // less than 4-byte selector
		{"minimum selector", "0x08c379a0", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractHexFromError(tt.s)
			if tt.want == 0 {
				if got != nil {
					t.Errorf("expected nil, got %x", got)
				}
			} else if len(got) != tt.want {
				t.Errorf("expected %d bytes, got %d", tt.want, len(got))
			}
		})
	}
}

func TestIsHexChar(t *testing.T) {
	for _, c := range "0123456789abcdefABCDEF" {
		if !isHexChar(byte(c)) {
			t.Errorf("isHexChar(%c) = false, want true", c)
		}
	}
	for _, c := range "ghijklGHIJKL!@#$% " {
		if isHexChar(byte(c)) {
			t.Errorf("isHexChar(%c) = true, want false", c)
		}
	}
}
