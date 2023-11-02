package validator

import (
	"context"
	"time"

	"github.com/foomo/contentfulvalidation/constants"
	"github.com/foomo/keel/log"
	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
)

type AttributeProviderFunc func() constants.Attributes
type AttributeUpdateFunc func(ctx context.Context) constants.Attributes

type AttributeProvider struct {
	l          *zap.Logger
	ctx        context.Context
	attributes constants.Attributes
	updateFunc AttributeUpdateFunc
}

func NewAttributeProvider(ctx context.Context, l *zap.Logger, uf AttributeUpdateFunc) *AttributeProvider {
	return &AttributeProvider{
		l:          l,
		ctx:        ctx,
		updateFunc: uf,
	}
}

func (ap *AttributeProvider) Init() error {
	ap.attributes = ap.updateFunc(ap.ctx)

	// TODO: make configurable
	s := gocron.NewScheduler(time.Local)
	_, err := s.Every(1).Day().At("03:00").Do(func() {
		ap.attributes = ap.updateFunc(ap.ctx)
	})
	if err != nil {
		log.Must(ap.l, err, "failed to initialize attribute provider scheduler")
	}
	s.StartAsync()
	return nil
}

func (ap *AttributeProvider) GetAttributes() constants.Attributes {
	return ap.attributes
}
