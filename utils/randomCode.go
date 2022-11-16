package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomCode() string {
	rand.Seed(time.Now().UnixNano())
	var code string
	for i := 0; i < 5; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}
	return code
}
