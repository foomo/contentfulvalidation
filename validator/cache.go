package validator

import (
	"sync"

	"go.uber.org/zap"
)

type Cache struct {
	sync.RWMutex
	l    *zap.Logger
	pool ValidationResults
}

func NewCache(l *zap.Logger) (*Cache, error) {
	logger := l.With(zap.String("routine", "contentfulvalidation-cache"))
	c := &Cache{
		RWMutex: sync.RWMutex{},
		l:       logger,
		pool:    map[ModelType]map[ModelID]*ValidationResult{},
	}
	return c, nil
}

func (c *Cache) Get(modelType ModelType, modelID ModelID) (*ValidationResult, bool) {
	c.RLock()
	defer c.RUnlock()
	// check if the modelType is populated at all
	_, typeMapExists := c.pool[modelType]
	if !typeMapExists {
		return nil, false
	}
	// check if the modelID has a validation result and return it
	validationResult, ok := c.pool[modelType][modelID]
	return validationResult, ok
}

func (c *Cache) GetForType(modelType ModelType) map[ModelID]*ValidationResult {
	c.RLock()
	defer c.RUnlock()
	results, typeMapExists := c.pool[modelType]
	if !typeMapExists {
		return nil
	}
	return results
}

func (c *Cache) SetForType(modelType ModelType, results map[ModelID]*ValidationResult) {
	c.Lock()
	defer c.Unlock()
	c.pool[modelType] = results
}

func (c *Cache) GetPool() ValidationResults {
	return c.pool
}
