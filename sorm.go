package sorm

type sormer interface {
	Build(...interface{})
}

func NewSormer()sormer{
	return NewBuilder()
}