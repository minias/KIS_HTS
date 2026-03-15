// internal/infrastructure/kis/dispatcher.go
package kis

import (
	"KIS_HTS/internal/application"
)

type Dispatcher struct {
	Service *application.MarketService
}

func (d *Dispatcher) Handle(msg []byte) {

	tick, ok := ParseTick(string(msg))
	if !ok {
		return
	}

	d.Service.UpdateTick(tick)
}
