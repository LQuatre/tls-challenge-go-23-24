package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.LoafOfBread("deliciousbread"))
	fmt.Print(piscine.LoafOfBread("This is a loaf of bread"))
	fmt.Print(piscine.LoafOfBread("loaf"))
	fmt.Print(piscine.LoafOfBread("Loaf of bread"))
	fmt.Print(piscine.LoafOfBread("amazing loaf of bread"))
	fmt.Print("\"" + piscine.LoafOfBread("This is a loaf of brea"+"\""))
}
