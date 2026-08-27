package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/jasontconnell/jpwd/generate"
)

func main() {
	length := flag.Int("length", 16, "length of the password to generate")
	ucase := flag.Bool("uppercase", false, "include uppercase")
	lcase := flag.Bool("lowercase", false, "include lowercase")
	numbers := flag.Bool("numbers", false, "include numbers")
	symbols := flag.String("symbols", "", "include symbols")
	flag.Parse()

	if *length == 0 || (!*ucase && !*lcase && !*numbers && len(*symbols) == 0) {
		flag.PrintDefaults()
		log.Fatal("please provide a length and it must include one of uppercase, lowercase, numbers, or symbols")
	}

	pg := generate.NewPasswordGenerator()

	pwd := pg.GeneratePassword(16, *lcase, *ucase, *numbers, *symbols)
	fmt.Println(pwd)
}
