package go_stream

type Predicate func(interface{}) bool
type Function func(interface{}) interface{}
type Comparator func(interface{}, interface{}) int
type Consumer func(interface{})

type Stream interface {
	Of(...interface{}) Stream

	Filter(predicate Predicate) Stream

	Map(function Function) Stream

	FlatMap(function Function) Stream

	Distinct() Stream

	Sorted(comparator Comparator) Stream

	Peek(consumer Consumer) Stream

	Limit(limit int64) Stream

	Skip(skip int64) Stream

	Concat(streams ...Stream) Stream

	ForEach(consumer Consumer)

	Min(comparator Comparator) interface{}

	Max(comparator Comparator) interface{}

	Count() int64

	AnyMatch(predicate Predicate) bool

	NonMatch(predicate Predicate) bool

	AllMatch(predicate Predicate) bool

	FindFirst() interface{}

	FindAny() interface{}

	ToSlice() []interface{}
}
