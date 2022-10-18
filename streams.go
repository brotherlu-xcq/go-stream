package go_stream

type PredicateWithIndex func(int, interface{}) bool

type sequentialStream struct {
	datas         []interface{}
	newData       []interface{}
	executorChain []interface{}
}

type mapData struct {
	key   interface{}
	value interface{}
}

func Streams(datas ...interface{}) *sequentialStream {
	return &sequentialStream{
		datas: datas,
	}
}

func StreamsSlice(slice []interface{}) *sequentialStream {
	return &sequentialStream{
		datas: slice,
	}
}

func StreamsMap(mapdatas map[interface{}]interface{}) *sequentialStream {
	datas := make([]interface{}, len(mapdatas))
	for k, v := range mapdatas {
		datas = append(datas, mapData{
			key:   k,
			value: v,
		})
	}
	return &sequentialStream{
		datas: datas,
	}
}

func (s *sequentialStream) Of(datas ...interface{}) Stream {
	s.datas = append(s.datas, datas)
	return s
}

func (s *sequentialStream) Filter(predicate Predicate) Stream {
	s.executorChain = append(s.executorChain, predicate)
	return s
}

func (s *sequentialStream) Map(function Function) Stream {
	s.executorChain = append(s.executorChain, function)
	return s
}

func (s *sequentialStream) FlatMap(function Function) Stream {
	s.executorChain = append(s.executorChain, function)
	return s
}

func (s *sequentialStream) Distinct() Stream {

}

func (s *sequentialStream) Sorted(comparator Comparator) Stream {
	s.executorChain = append(s.executorChain, comparator)
	return s
}

func (s *sequentialStream) Peek(consumer Consumer) Stream {
	s.executorChain = append(s.executorChain, consumer)
	return s
}

func (s *sequentialStream) Limit(limit int64) Stream {
	s.executorChain = append(s.executorChain)
	return s
}

func (s *sequentialStream) Skip(skip int64) Stream {
	s.executorChain = append(s.executorChain)
	return s
}

func (s *sequentialStream) Concat(streams ...Stream) Stream {
	for _, stream := range streams {
		s.datas = append(s.datas, stream.ToSlice())
	}
	return s
}

func (s *sequentialStream) ForEach(consumer Consumer) {}

func (s *sequentialStream) Min(comparator Comparator) interface{} {
	s.executorChain = append(s.executorChain, comparator)
	s.evaluateSequential()
	return s.newData[0]
}

func (s *sequentialStream) Max(comparator Comparator) interface{} {
	s.executorChain = append(s.executorChain, comparator)
	s.evaluateSequential()
	return s.newData[0]
}

func (s *sequentialStream) Count() int64 {
	s.evaluateSequential()
	return int64(len(s.newData))
}

func (s *sequentialStream) AnyMatch(predicate Predicate) bool {
	s.evaluateSequential()
	return len(s.newData) >= 1
}

func (s *sequentialStream) NonMatch(predicate Predicate) bool {
	s.evaluateSequential()
	return len(s.newData) == 0
}

func (s *sequentialStream) AllMatch(predicate Predicate) bool {
	s.evaluateSequential()
	return len(s.newData) == len(s.datas)
}

func (s *sequentialStream) FindFirst() interface{} {
	return s.newData[0]
}

func (s *sequentialStream) FindAny() interface{} {
	return s.newData[len(s.datas)-1]
}

func (s *sequentialStream) ToSlice() []interface{} {
	return s.newData
}

func (s *sequentialStream) evaluateSequential() {

}
