package cache

// A ByteView holds an immutable view of bytes
type ByteView struct {
	b []byte
}

// Len returns the length of view
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}