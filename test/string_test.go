package test

import (
	"fmt"
	"mall/utils"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	join := strings.Join([]string{"sss", "dddd"}, "  ")
	fmt.Println(join)
}

func TestCode(t *testing.T) {
	fmt.Println(utils.RandomCode())
	fmt.Println(utils.RandomCode())
	fmt.Println(utils.RandomCode())
	fmt.Println(utils.RandomCode())
	fmt.Println(utils.RandomCode())
	fmt.Println(utils.RandomCode())
	fmt.Println(utils.RandomCode())
	/**
	56224
	67846
	57324
	80225
	20510
	77163
	46553
	*/
}
