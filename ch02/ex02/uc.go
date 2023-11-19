package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/adieumonks/gpl-exercise/ch02/ex02/unitconv"
)

func main() {
	for _, arg := range getArgs() {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "uc: %v\n", err)
			os.Exit(1)
		}
		showTemperature(t)
		showLength(t)
		showWeight(t)
	}
}

func getArgs() []string {
	args := []string{}
	if len(os.Args) == 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			args = append(args, strings.Split(input.Text(), " ")...)
		}
	} else {
		return os.Args[1:]
	}
	return args
}

func showTemperature(t float64) {
	f := unitconv.Fahrenheit(t)
	c := unitconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FToC(f), c, unitconv.CToF(c))
}

func showLength(t float64) {
	f := unitconv.Feet(t)
	m := unitconv.Meter(t)
	fmt.Printf("%s = %s, %s = %s\n", f, unitconv.FeetToMeter(f), m, unitconv.MeterToFeet(m))
}

func showWeight(t float64) {
	p := unitconv.Pound(t)
	k := unitconv.Kilogram(t)
	fmt.Printf("%s = %s, %s = %s\n", p, unitconv.PoundToKilogram(p), k, unitconv.KilogramToPound(k))
}
