package gosol

type Mapping[K comparable, V any] interface {
	Get(key K) (value V, ok bool)
	Set(key K, value V)
}

type mapMapping[K comparable, V any] struct {
	m map[K]V
}

func (m *mapMapping[K, V]) Get(key K) (value V, ok bool) {
	value, ok = m.m[key]
	return
}

func (m *mapMapping[K, V]) Set(key K, value V) {
	m.m[key] = value
}

func NewMapMapping[K comparable, V any]() Mapping[K, V] {
	return &mapMapping[K, V]{
		m: make(map[K]V),
	}
}
