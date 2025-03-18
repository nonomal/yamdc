package handler

import (
	"context"
	"fmt"
	"yamdc/enum"
	"yamdc/model"
	"yamdc/translator"
)

type translaterHandler struct {
}

func (p *translaterHandler) Name() string {
	return HTranslater
}

func (p *translaterHandler) translateSingle(ctx context.Context, name string, in string, lang string, out *string) error {
	if len(in) == 0 {
		return nil
	}
	if len(lang) == 0 || lang == enum.MetaLangZHTW || lang == enum.MetaLangZH {
		return nil
	}
	res, err := translator.Translate(ctx, in, "auto", "zh")
	if err != nil {
		return fmt.Errorf("translate failed, name:%s, err:%w", name, err)
	}
	*out = res
	return nil
}

func (p *translaterHandler) Handle(ctx context.Context, fc *model.FileContext) error {
	if !translator.IsTranslatorEnabled() {
		return nil
	}
	var errs []error
	errs = append(errs, p.translateSingle(ctx, "title", fc.Meta.Title, fc.Meta.TitleLang, &fc.Meta.TitleTranslated))
	errs = append(errs, p.translateSingle(ctx, "plot", fc.Meta.Plot, fc.Meta.PlotLang, &fc.Meta.PlotTranslated))

	for _, err := range errs {
		if err != nil {
			return fmt.Errorf("translate part failed, err:%w", err)
		}
	}
	return nil
}

func init() {
	Register(HTranslater, HandlerToCreator(&translaterHandler{}))
}
