package md2

import (
	"encoding/hex"
	"testing"
)

func TestRFC1319Vectors(t *testing.T) {
	tests := []struct{ input, want string }{
		{"", "8350e5a3e24c153df2275c9f80692773"},
		{"a", "32ec01ec4a6dac72c0ab96fb34c0b5d1"},
		{"abc", "da853b0d3f88d99b30283a69e6ded6bb"},
		{"message digest", "ab4f496bfb2a530b219ff33031fe06b0"},
		{"abcdefghijklmnopqrstuvwxyz", "4e8ddff3650292ab5a4108c3aa47940b"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", "da33def2a42df13975352846c30338cd"},
		{"12345678901234567890123456789012345678901234567890123456789012345678901234567890", "d5976f79d83d3a0dc9806c3c66f3efd8"},
	}
	for _, tt := range tests {
		got := Sum([]byte(tt.input))
		if hex.EncodeToString(got[:]) != tt.want {
			t.Errorf("Sum(%q) = %x, want %s", tt.input, got, tt.want)
		}
	}
}

func TestWriteInChunks(t *testing.T) {
	input := []byte("the quick brown fox jumps over the lazy dog; " + "0123456789abcdef")
	want := Sum(input)
	h := New()
	for i := 0; i < len(input); i += 3 {
		end := i + 3
		if end > len(input) { end = len(input) }
		n, err := h.Write(input[i:end])
		if err != nil || n != end-i { t.Fatalf("Write returned (%d, %v)", n, err) }
	}
	got := h.Sum(nil)
	if string(got) != string(want[:]) { t.Fatalf("chunked Sum = %x, want %x", got, want) }
}

func TestSumDoesNotChangeState(t *testing.T) {
	h := New()
	_, _ = h.Write([]byte("prefix"))
	first := h.Sum(nil)
	second := h.Sum(nil)
	if string(first) != string(second) { t.Fatal("repeated Sum changed the digest") }
	_, _ = h.Write([]byte("suffix"))
	want := Sum([]byte("prefixsuffix"))
	if got := h.Sum(nil); string(got) != string(want[:]) { t.Fatalf("after Sum = %x, want %x", got, want) }
}

func TestReset(t *testing.T) {
	h := New()
	_, _ = h.Write([]byte("discarded"))
	h.Reset()
	got := h.Sum(nil)
	want := Sum(nil)
	if string(got) != string(want[:]) { t.Fatalf("after Reset = %x, want %x", got, want) }
}

func TestAPIConstants(t *testing.T) {
	h := New()
	if h.Size() != Size || h.BlockSize() != BlockSize { t.Fatal("unexpected hash dimensions") }
}
