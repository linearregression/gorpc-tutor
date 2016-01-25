package world

import "golang.org/x/net/context"

func (*Handler) V1(ctx context.Context, opts *struct{}) (string, error) {
	return "Hello world", nil
}
