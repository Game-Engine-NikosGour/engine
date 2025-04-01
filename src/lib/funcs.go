package lib

import (
	log "github.com/NikosGour/logging/src"
)

func add(x int, y int) {
	log.Debug("%d + %d = %d", x, y, x+y)
}
