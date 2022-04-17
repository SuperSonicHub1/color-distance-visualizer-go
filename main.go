package main

import (
	"flag"
	"fmt"
	"os"
)

var input string
var output string
var unchangedPixels bool
var vcodec string

func init() {
	flag.StringVar(&input, "input", "", "The video that we want to analyze.")
	flag.StringVar(&output, "output", "", "Where we want to save the output.")
	flag.BoolVar(&unchangedPixels, "show-unchanged-pixels", false, "For pixels that haven't changed, display the pixel from the orginal frame instead of black.")
	flag.StringVar(&vcodec, "vcodec", "", "What codec you want the output to be saved in. FFV1 is the recommended lossless codec. Using lossy codecs like MPEG-4 will result in significant loss of detail. (default: codec of input)")
}

func main() {
	flag.Parse()
	if input == "" || output == "" {
		fmt.Fprintln(os.Stderr, "Input/output parameters are empty.")
		flag.Usage()
		os.Exit(1)
	}

	err := visualize(
		input,
		output,
		unchangedPixels,
		vcodec,
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
