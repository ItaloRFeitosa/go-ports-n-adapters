package internal

type ControllerFunc[In any, Out any] func(In) (Out, error)
