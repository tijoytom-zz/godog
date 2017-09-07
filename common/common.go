package common

type Iterator interface {

    Next() bool
	
	Value() interface{}

	Key() interface{}

	Begin()

	First() bool
}

type ReverseIterator interface {
	Prev() bool
	
	End()

	Last() bool

	Iterator
}