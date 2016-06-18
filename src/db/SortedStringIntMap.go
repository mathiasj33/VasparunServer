package db

type SortedStringIntMap struct {
	keys []string
	values []int
}

func NewSortedStringIntMap() *SortedStringIntMap {
	m := SortedStringIntMap{make([]string, 0), make([]int, 0)}
	return &m
}

func (s *SortedStringIntMap) Put(key string, value int) {
	s.keys = append(s.keys, key)
	s.values = append(s.values, value)
}

func (s *SortedStringIntMap) Get(key string) int {
	i := indexOfString(key, s.keys)
	return s.values[i]
}

func (s *SortedStringIntMap) GetFromIndex(i int) (string, int) {
	return s.keys[i], s.values[i]
}

func (s *SortedStringIntMap) Length() int {
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

func indexOfInt(e int, slice []int) int {
	for i,e := range slice {
		if e == slice[i] {
			return i
		}
	}
	return -1
}