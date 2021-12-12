package medium

/*
	284. Peeking Iterator
*/

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type Iterator struct {
	sl []int
	i  int
}

func (it *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return it.i < len(it.sl)
}

func (it *Iterator) next() int {
	// Returns the next element in the iteration.
	n := it.sl[it.i]
	it.i++

	return n
}

type PeekingIterator struct {
	cached   bool
	cachedEl int
	itr      *Iterator
}

func peekingConstructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		itr: iter,
	}
}

func (pi *PeekingIterator) hasNext() bool {
	if pi.cached {
		return pi.cached
	}

	return pi.itr.hasNext()
}

func (pi *PeekingIterator) next() int {
	if pi.cached {
		next := pi.cachedEl

		pi.cached = false
		return next
	}

	return pi.itr.next()
}

func (pi *PeekingIterator) peek() int {
	if pi.cached {
		return pi.cachedEl
	} else if pi.itr.hasNext() {
		next := pi.itr.next()

		pi.cached = true
		pi.cachedEl = next

		return next
	}
	return -1
}
