// internal/infrastructure/kis/parser.go
package kis

import (
	"strings"

	"KIS_HTS/internal/domain/market"
)

// ParseTick parse websocket tick message
func ParseTick(msg string) (market.Tick, bool) {

	parts := strings.Split(msg, "|")

	if len(parts) < 5 {
		return market.Tick{}, false
	}

	return market.Tick{
		Symbol: parts[2],
	}, true
}
