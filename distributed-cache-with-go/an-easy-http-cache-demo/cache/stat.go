package cache

type Stat struct {
	Count     int64
	KeySize   int64
	ValueSize int64
}

func (s *Stat) add(key string, value []byte) {
	s.Count++
	s.KeySize += int64(len(key))
	s.ValueSize += int64(len(value))
}

func (s *Stat) del(key string, value []byte) {
	s.Count--
	s.KeySize -= int64(len(key))
	s.ValueSize -= int64(len(value))
}
