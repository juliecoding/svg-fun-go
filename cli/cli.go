package cli

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"log"

	// Decoding any particular image format requires the prior registration of a decoder function.
	// Registration is typically automatic as a side effect of initializing that format's package,
	// so here we're using _ to import the package purely for its initialization side effects.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/ioutil"
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
	flag.Var(&app.filters, "filters", "Name of filters from the filters.json file to apply to the input file. Possible values include:\ndawn, dusk, night, and black_and_white.")
	// Parse parses the command-line flags from os.Args[1:].
	// Must be called after all flags are defined and before flags are accessed by the program.
	// Also, internally it ignores errors because CommandLine, which it uses, is set for ExitOnError
	flag.Parse()
	app.validateInput()
	app.op()
	return nil
}

func (app *appEnv) validateInput() {
	if (app.out == "") {
		app.out = getUserInput("An output filepath is required. Please enter one (without quotes): ")
		app.validateInput()
	}
	if (app.in == "") {
		app.in = getUserInput("An input filepath is required. Please enter one (without quotes): ")
		app.validateInput()
	}
	if (len(app.filters) == 0) {
		fmt.Println("Just FYI -- it looks like you didn't provide any filters")
	}
}

func (app *appEnv) op() {
	outWriter, errC := os.Create(app.out)
	for errC != nil {
		outputError(fmt.Sprintf("There was an issue creating the file %s", app.out), errC)
		app.out = getUserInput("Please enter a valid filepath for your output (without quotes): ")
		outWriter, errC = os.Create(app.out)
	}
	inputFile, errO := os.Open(app.in)
	for errO != nil {
		inputFile.Close()
		outputError(fmt.Sprintf("There was an issue opening the input file at %s", app.in), errO)
		app.in = getUserInput("Please enter a valid input filepath (without quotes): ")
		inputFile, errO = os.Open(app.in)
	}
	// Middle term = Format name used during format registration
	img, _, errD := image.DecodeConfig(inputFile)
	for (errD != nil) {
		outputError(fmt.Sprintf("There was an issue decoding the input file at %s", app.in), errD)
		app.in = getUserInput("Please enter a valid input filepath (without quotes): ")
		img, _, errD = image.DecodeConfig(inputFile)
	}
	inputFile.Close()
	createSvg(app.in, outWriter, app.filters, img.Width, img.Height)
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

func createSvg(in string, outWriter io.Writer, filters filters, w int, h int) {
	// New is the SVG constructor, specifying the io.Writer where the generated SVG is written.
	canvas := svg.New(outWriter)
	canvas.Start(w, h)
	canvas.Def()
	addFilters(*canvas, filters)
	canvas.DefEnd()
	// JAK, allow 0 filters
	canvas.Image(0, 0, w, h, in, `filter="url(#__filters)"`)
	canvas.End()
}

func addFilters(canvas svg.SVG, filters filters) svg.SVG {
	animationDelay := 0.0
	canvas.Filter("__filters")
	for _, filter := range filters {
		switch {
		case filter == "dawn":
			addDawnFilter(canvas, animationDelay)
		case filter == "dusk":
			addDuskFilter(canvas, animationDelay)
		case filter == "night":
			addNightFilter(canvas, animationDelay)
		case filter == "black_and_white":
			addBlackAndWhiteFilter(canvas)
		}
	}
	canvas.FeMerge(filters)
	canvas.Fend()
	return canvas
}

func addDawnFilter(canvas svg.SVG, animationDelay float64) svg.SVG {
	result := "floodOut"
	floodId := "__df"
	floodSpec := svg.Filterspec{ Result: result }
	blendSpec := svg.Filterspec{ In2: result, In: "SourceGraphic"}
	canvas.FeFlood(floodSpec, "#EFB2D1", 0, fmt.Sprintf(`id="%s"`, floodId))
	canvas.FeBlend(blendSpec, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 1, 0, animationDelay + 7, 1)
	return canvas
}

func addDuskFilter(canvas svg.SVG, animationDelay float64) svg.SVG  {
	result1 := "duskFlood"
	floodId1 := "__df1"
	floodSpec1 := svg.Filterspec{ Result: result1 }
	blendSpec1 := svg.Filterspec{ In2: result1, In: "SourceGraphic" }
	canvas.FeFlood(floodSpec1, "#e3b249", 1, fmt.Sprintf(`id="%s"`, floodId1))
	canvas.FeBlend(blendSpec1, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId1), "flood-opacity", 0, 1, animationDelay + 7, 1)
	return canvas
}

func addNightFilter(canvas svg.SVG, animationDelay float64) svg.SVG {
	result1 := "nightFlood"
	floodId1 := "__nf1"
	floodSpec1 := svg.Filterspec{ Result: result1 }
	blendSpec1 := svg.Filterspec{ In2: result1, In: "SourceGraphic" }
	canvas.FeFlood(floodSpec1, "#0c0b0f", 1, fmt.Sprintf(`id="%s"`, floodId1))
	canvas.FeBlend(blendSpec1, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId1), "flood-opacity", 0, 1, animationDelay + 7, 1)
	return canvas
}

func addBlackAndWhiteFilter(canvas svg.SVG) svg.SVG {
	var fcm FeColorMatrix
	fp := "./filters/bw.json"
	def, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error trying to open %s", fp))
		log.Fatal(err)
	}
	err = json.Unmarshal(def, &fcm)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error unmarshalling JSON in %s", fp))
		log.Fatal(err)
	}
	filterSpec := svg.Filterspec{ In: fcm.In }
	canvas.FeColorMatrixSaturate(filterSpec, fcm.Values)
	return canvas
}
