package tcp

type TCPFrame interface {
	init([]byte) error
	get() error
	set([]byte) error
}

type FrameHeader struct {
	KeyLength   int64
	ValueLength int64
}

func (fh *FrameHeader) init(key, value []byte) {
	fh.KeyLength = int64(len(key))
	fh.ValueLength = int64(len(value))
}

type Frame struct {
	Header *FrameHeader
	Key    []byte
	Value  []byte
}

func (f *Frame) init(key, value []byte) error {
	header := FrameHeader{}
	header.init(key, value)
	f.Header = &header
	f.Key = key
	f.Value = value
	return nil
}

func (f *Frame) get() []byte {
	return f.Key
}

func (f *Frame) set(value []byte) error {
	f.Header.ValueLength = int64(len(value))
	f.Value = value
	return nil
}
