package handler

import (
	"av-capture/model"
	"context"
	"fmt"
)

type IHandler interface {
	Handle(ctx context.Context, fc *model.FileContext) error
}

var mp = make(map[string]CreatorFunc)

type CreatorFunc func(args interface{}) (IHandler, error)

func Register(name string, fn CreatorFunc) {
	mp[name] = fn
}

func CreateHandler(name string, args interface{}) (IHandler, error) {
	cr, ok := mp[name]
	if !ok {
		return nil, fmt.Errorf("handler:%s not found", name)
	}
	return cr(args)
}

func HandlerToCreator(h IHandler) CreatorFunc {
	return func(args interface{}) (IHandler, error) {
		return h, nil
	}
}
