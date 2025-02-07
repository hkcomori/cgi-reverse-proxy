package application

import (
	"iter"
	"maps"
)

type QueryMap[K comparable, V any] iter.Seq2[K, V]

func NewQueryMap[K comparable, V any](m map[K]V) QueryMap[K, V] {
	return QueryMap[K, V](maps.All(m))
}

func NewQueryMapFromIter[K comparable, V any](iterator iter.Seq2[K, V]) QueryMap[K, V] {
	return QueryMap[K, V](iterator)
}

func (f QueryMap[K, V]) toIter() iter.Seq2[K, V] {
	return iter.Seq2[K, V](f)
}

func (f QueryMap[K, V]) Filter(cond func(k K, v V) bool) QueryMap[K, V] {
	return func(yield func(k K, v V) bool) {
		for k, v := range f {
			if !cond(k, v) {
				continue
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

func (f QueryMap[K, V]) Replace(selector func(k K, v V) (K, V)) QueryMap[K, V] {
	return func(yield func(k K, v V) bool) {
		for k, v := range f {
			if !(yield(selector(k, v))) {
				return
			}
		}
	}
}

func (f QueryMap[K, V]) All() map[K]V {
	return maps.Collect(f.toIter())
}
