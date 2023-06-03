package primary

type FeatureFunc[In any, Out any] func(In) (Out, error)
