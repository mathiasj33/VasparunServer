package db

type SortedStringFloatMap struct {
	keys []string
	values []float32
}

func NewSortedStringFloatMap() *SortedStringFloatMap {
	m := SortedStringFloatMap{make([]string, 0), make([]float32, 0)}
	return &m
}

func (s *SortedStringFloatMap) Put(key string, value float32) {
	s.keys = append(s.keys, key)
	s.values = append(s.values, value)
}

func (s *SortedStringFloatMap) Get(key string) float32 {
	i := indexOfString(key, s.keys)
	return s.values[i]
}

func (s *SortedStringFloatMap) GetFromIndex(i int) (string, float32) {
	return s.keys[i], s.values[i]
}

func (s *SortedStringFloatMap) Length() int {
	return len(s.keys)
}

func indexOfString(e string, slice []string) int {
	for i,e := range slice {
		if e == slice[i] {
			return i
		}
	}
	return -1
}

func indexOfFloat(e float32, slice []float32) int {
	for i,e := range slice {
		if e == slice[i] {
			return i
		}
	}
	return -1
}