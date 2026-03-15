// internal/application/market_service.go
package application

import (
	"sync"

	"KIS_HTS/internal/domain/market"
)

// MarketService manages market data state
type MarketService struct {
	mu    sync.RWMutex
	store map[string]market.Tick
}

// NewMarketService creates service instance
func NewMarketService() *MarketService {
	return &MarketService{
		store: make(map[string]market.Tick),
	}
}

// UpdateTick update last price
func (m *MarketService) UpdateTick(tick market.Tick) {

	m.mu.Lock()
	defer m.mu.Unlock()

	m.store[tick.Symbol] = tick
}

// GetLastPrice return latest price
func (m *MarketService) GetLastPrice(symbol string) (float64, error) {

	m.mu.RLock()
	defer m.mu.RUnlock()

	if v, ok := m.store[symbol]; ok {
		return v.Price, nil
	}

	return 0, nil
}
