package sum

import "golang.org/x/net/context"

type v1Args struct {
	Num1 uint `key:"num1" description:"First number"`
	Num2 uint `key:"num2" description:"Second number"`
}

type v1Res struct {
	Sum      uint `json:"sum" description:"Sum of numbers"`
	Overflow bool `json:"overflow" description:"True if sum out of range"`
}

func (*Handler) V1(ctx context.Context, opts *v1Args) (v1Res, error) {
	res := v1Res{
		Sum:      opts.Num1 + opts.Num2,
		Overflow: ^uint(0)-opts.Num1 < opts.Num2,
	}

	return res, nil
}
