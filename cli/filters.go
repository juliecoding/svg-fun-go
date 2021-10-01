package cli

import (
	"fmt"

	"github.com/ajstarks/svgo"
)

type filterInterface interface { // cap???
	apply()					// cap???
	addToFilters()
}

type dawnFilter struct { // JAK, is it appropriate to make this a type?? 
						// Alternatives???
	name string
}

type duskFilter struct {
	name string
}

type nightFilter struct {
	name string
}

// Nothing in this method uses properties of the dawn filter itself. Hmmm.
func (dw *dawnFilter) apply(canvas svg.SVG, animationDelay float64) svg.SVG {
	result := "floodOut"
	floodId := "__df"
	floodSpec := svg.Filterspec{ Result: result }
	blendSpec := svg.Filterspec{ In2: result, In: "SourceGraphic"}
	canvas.FeFlood(floodSpec, "#EFB2D1", 0, fmt.Sprintf(`id="%s"`, floodId))
	canvas.FeBlend(blendSpec, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 1, 0, animationDelay + 7, 1)
	return canvas
}

func (dk *duskFilter) apply(canvas svg.SVG, animationDelay float64) svg.SVG {
	result := "floodOut"
	floodId := "__df"
	floodSpec := svg.Filterspec{ Result: result }
	blendSpec := svg.Filterspec{ In2: result, In: "SourceGraphic"}
	canvas.FeFlood(floodSpec, "#EFB2D1", 0, fmt.Sprintf(`id="%s"`, floodId))
	canvas.FeBlend(blendSpec, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId), "flood-opacity", 1, 0, animationDelay + 7, 1)
	return canvas
}

func (ni *nightFilter) apply(canvas svg.SVG, animationDelay float64) svg.SVG {
	result1 := "nightFlood"
	floodId1 := "__nf1"
	floodSpec1 := svg.Filterspec{ Result: result1 }
	blendSpec1 := svg.Filterspec{ In2: result1, In: "SourceGraphic" }
	canvas.FeFlood(floodSpec1, "#0c0b0f", 1, fmt.Sprintf(`id="%s"`, floodId1))
	canvas.FeBlend(blendSpec1, "multiply")
	canvas.Animate(fmt.Sprintf("#%s", floodId1), "flood-opacity", 0, 1, animationDelay + 7, 1)
	return canvas
}

