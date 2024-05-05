package world

import (
	"ejercicio-poo/challenge/yu-gi-oh/common"
)

func GetRandomWorld(worlds []World) World {
	return common.GetRandomObject(worlds)
}
