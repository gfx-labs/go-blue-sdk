package gosol

type Mapping[K comparable, V any] interface {
	Get(key K) (value V, err error)
	Set(key K, value V) error
}

type mapMapping[K comparable, V any] struct {
	m map[K]V
}

func (m *mapMapping[K, V]) Get(key K) (value V, err error) {
	value, _ = m.m[key]
	return value, nil
}

func (m *mapMapping[K, V]) Set(key K, value V) error {
	m.m[key] = value
	return nil
}

func NewMapMapping[K comparable, V any]() Mapping[K, V] {
	return &mapMapping[K, V]{
		m: make(map[K]V),
	}
}
