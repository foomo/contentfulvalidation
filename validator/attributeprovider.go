package validator

import (
	"context"
	"time"

	"github.com/bestbytes/catalogue/vo"
	"github.com/foomo/keel/log"
	"github.com/go-co-op/gocron"
	"go.uber.org/zap"
)

type AttributeProviderFunc func() vo.Attributes
type AttributeUpdateFunc func(ctx context.Context) vo.Attributes

type AttributeProvider struct {
	l          *zap.Logger
	ctx        context.Context
	attributes vo.Attributes
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
	log.Must(ap.l, err, "failed to ...")
	s.StartAsync()
	return nil
}

func (ap *AttributeProvider) GetAttributes() vo.Attributes {
	return ap.attributes
}
