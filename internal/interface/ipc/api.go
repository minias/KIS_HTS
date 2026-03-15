package ipc

import "KIS_HTS/internal/application"

// MarketAPI exposed to frontend
type MarketAPI struct {
	Service *application.MarketService
}

// GetPrice example method
func (m *MarketAPI) GetPrice(symbol string) (float64, error) {
	return m.Service.GetLastPrice(symbol)
}
