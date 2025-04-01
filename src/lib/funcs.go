package lib

import (
	log "github.com/NikosGour/logging/src"
)

func Add(x int, y int) {
	log.Debug("%d + %d = %d", x, y, x+y)
}
