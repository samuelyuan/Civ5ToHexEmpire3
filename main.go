package main

import (
	"flag"
	"fmt"

	"github.com/samuelyuan/Civ5ToHexEmpire3/fileio"
)

func main() {
	inputPtr := flag.String("input", "", "Input filename")
	outputPtr := flag.String("output", "output.he3", "Output filename")
	flag.Parse()

	fmt.Println("Input filename: ", *inputPtr)
	fmt.Println("Output filename: ", *outputPtr)
	fileio.ConvertCiv5MapToHE3Map(*inputPtr, *outputPtr)
}
