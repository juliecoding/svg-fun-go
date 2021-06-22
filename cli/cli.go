package cli

import (
	"bufio"
	"flag"
	"fmt"
	// "image"
	"io"

	// Decoding any particular image format requires the prior registration of a decoder function. Registration is typically automatic as a side effect of initializing that format's package, so here we're using _ to import the package purely for its initialization side effects.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"strings"

	"github.com/ajstarks/svgo"
)

func Run(args []string) int {
	// Conventionally, for os.Exit, code zero indicates success, non-zero an error.
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

type appEnv struct {
	filters	filters
	in		string
	out		string
}

func (app *appEnv) fromArgs(args []string) error {
	flag.StringVar(&app.out, "out", "./out/img.svg", "Path to SVG output")
	flag.StringVar(&app.in, "in", "", "Path to access input image")
	flag.Var(&app.filters, "filters", "Name of filters from the filters.json file to apply to the input file. Multiple filter names can be provided as a single string with spaces between them, enclosed in double quotes")
	// Parse parses the command-line flags from os.Args[1:].
	// Must be called after all flags are defined and before flags are accessed by the program.
	// Also, internally it ignores errors because CommandLine, which it uses, is set for ExitOnError
	flag.Parse()
	app.validateInput()
	return nil
}

func (app *appEnv) validateInput() {
	app.checkForMissingFields()
	// Validate input/output paths by attempting to interact with the files
	// Allows the user to correct paths if necessary
	inputFile := app.openInputFile()
	outputFileWriter := app.createOutputFile()
	// JAK
	w, h := app.getImageDimensions(inputFile)
	createSvg(app.in, outputFileWriter, app.filters, w, h)
}

func (app *appEnv) checkForMissingFields() {
	if (app.out == "") {
		app.out = getUserInput("An output filepath is required. Please enter one (without quotes): ")
	}
	if (app.in == "") {
		app.in = getUserInput("An input filepath is required. Please enter one (without quotes): ")
	}
}

func (app *appEnv) createOutputFile() io.Writer {
	// If the file already exists, Create will truncate it
	outputFileWriter, err := os.Create(app.out)
	if err != nil {
		explanation := "An error occurred while trying to create file at "
		outputError(explanation, err)
		app.in = getUserInput("Please enter a valid filepath for your output: ")
		app.createOutputFile()
	}
	return outputFileWriter
}

func (app *appEnv) openInputFile() *os.File {
	inputFile, err := os.Open(app.in)
	// Deferring file close and then entering a bad input loop creates issues
	defer inputFile.Close()
	if err != nil {
		explanation := "An error occurred trying to open your input file: "
		outputError(explanation, err)
		app.in = getUserInput("Please enter a valid input filepath: ")
	}
	return inputFile
}

func createSvg(in string, outputFileWriter io.Writer, filters filters, w int, h int) {
	// New is the SVG constructor, specifying the io.Writer where the generated SVG is written.
	canvas := svg.New(outputFileWriter)
	canvas.Start(w, h)
	canvas.Def()
	addFilters(*canvas, filters)
	canvas.DefEnd()
	canvas.Image(0, 0, w, h, in, `filter="url(#__filters)"`)
	canvas.End()
}

// JAK
func (app *appEnv) getImageDimensions(inputFile *os.File) (int, int) {
	return 1024, 1073
	// img, _, err := image.DecodeConfig(inputFile)
	// width := img.Width
	// height := img.Height
	// return width, height
}

// Because this is a custom flag Var, we have to implement String and Set methods
type filters []string

func (fl *filters) String() string {
	return strings.Join(*fl, " ")
}

func (fl *filters) Set(s string) error {
	sp := strings.Split(s, " ")
	*fl = append(*fl, sp...)
	return nil
}

func getUserInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(prompt)
	scanner.Scan()
	return scanner.Text()
}

func outputError(explanation string, err error, ) {
	fmt.Println(explanation)
	fmt.Println(err)
}

func addFilters(canvas svg.SVG, filters filters) svg.SVG {
	canvas.Filter("__filters")

	for _, filter := range filters {
		switch {
		case filter == "dawn":
			result := "floodOut"
			floodId := "__df"
			floodSpec := svg.Filterspec{ Result: result }
			blendSpec := svg.Filterspec{ In2: result, In: "SourceGraphic"}
			canvas.FeFlood(floodSpec, "#EFB2D1", 0.0, fmt.Sprintf(`id="%s"`, floodId))
			canvas.FeBlend(blendSpec, "multiply")
			canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 0, 1, 5, 1)


		case filter == "dusk":

		}
	}
	canvas.Fend()
	return canvas
}
