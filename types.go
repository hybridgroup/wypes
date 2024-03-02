package wypes

type Int8 int8

func (Int8) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

func (Int8) Lift(s Store) Int8 {
	return Int8(s.Stack.Pop())
}

func (v Int8) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type Int16 int16

func (Int16) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

func (Int16) Lift(s Store) Int16 {
	return Int16(s.Stack.Pop())
}

func (v Int16) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type Int32 int32

func (Int32) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI32}
}

func (Int32) Lift(s Store) Int32 {
	return Int32(s.Stack.Pop())
}

func (v Int32) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type Int64 int64

func (Int64) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

func (Int64) Lift(s Store) Int64 {
	return Int64(s.Stack.Pop())
}

func (v Int64) Lower(s Store) {
	s.Stack.Push(Raw(v))
}

type Int int

func (Int) ValueTypes() []ValueType {
	return []ValueType{ValueTypeI64}
}

func (Int) Lift(s Store) Int {
	return Int(s.Stack.Pop())
}

func (v Int) Lower(s Store) {
	s.Stack.Push(Raw(v))
}
