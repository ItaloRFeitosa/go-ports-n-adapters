package ginadapter

type RequestMapper[In any] interface {
	ToInput() In
}

type ResponseMapper[Out any] interface {
	FromOutput(output Out) any
	StatusCode() int
}
