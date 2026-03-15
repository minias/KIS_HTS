// internal/domain/market/tick.go
package market

// Tick represents real-time market data
// Domain Layer: 외부 API, DB, UI 의존성 없음
type Tick struct {
	Symbol string  // 종목코드
	Price  float64 // 현재가
	Volume int64   // 거래량
}
