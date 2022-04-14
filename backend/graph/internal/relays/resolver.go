package relays

import (
	"log"

	"github.com/hsblhsn/queues"
)

type ResolverFunc[K int, V any] func(K) (V, error)

type Resolver[K int, V any] struct {
	resolve ResolverFunc[K, V]
	arr     []K
}

func NewResolver[K int, V any](arr []K, resolve ResolverFunc[K, V]) *Resolver[K, V] {
	return &Resolver[K, V]{
		arr:     arr,
		resolve: resolve,
	}
}

func (r *Resolver[K, V]) Resolve(before, after *string, first, last *int) (*Connection[V], error) {
	pager, err := Paginate(len(r.arr), ConnectionInput{
		After:  after,
		Before: before,
		First:  first,
		Last:   last,
	})
	if err != nil {
		return nil, err
	}
	var (
		fetchables = r.arr[pager.Start:pager.End]
		edges      = make([]*Edge[V], len(fetchables))
	)
	resolve := func(q *queues.Q, list []*Edge[V], index int, id K) {
		defer q.Done()
		node, err := r.resolve(id)
		if err != nil {
			log.Println("relay: could not resolve", err)
			// not returning here.
			// because return will stop the func.
			// and the edge[index] will be nil.
		}
		list[index] = &Edge[V]{
			Cursor: NewCursor(pager.Start + index),
			Node:   node,
		}
	}
	q := queues.New(MaxConcurrency)
	for index, v := range fetchables {
		q.Add(1)
		go resolve(q, edges, index, v)
	}
	q.Wait()
	return &Connection[V]{
		PageInfo: &PageInfo{
			HasNextPage:     pager.End < len(r.arr),
			HasPreviousPage: pager.Start > 0,
			PageCursor:      NewCursor(max(pager.Start-1, 0)),
			StartCursor:     NewCursor(pager.Start),
			EndCursor:       NewCursor(pager.End),
		},
		TotalCount: len(r.arr),
		Edges:      edges,
	}, nil
}
