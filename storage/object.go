package storage

func (s *ObjectSystem) ListObjects(Key string) string {
	return s.bucket
}

func (s *ObjectSystem) RetrieveObject(Key string) string {
	return s.bucket
}
