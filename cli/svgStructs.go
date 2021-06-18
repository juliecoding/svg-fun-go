package cli

type CoreAttributes struct {
	Id 		string	`json:"id"`
	Lang 	string	`json:"lang"`
}

type FilterPrimitiveAttributes struct {
	X 		string 	`json:"x"`		// x, y, height and width can all be percentage or number
	Y 		string	`json:"y"`
	Height 	string	`json:"height"`
	Width 	string	`json:"width"`
	Result 	string	`json:"result"`	// Defines the assigned name for this filter primitive
}

type GlobalAttributes struct {
	CoreAttributes
	FilterPrimitiveAttributes
	PresentationAttributes
	Class 	string	`json:"class"`
	Style 	string	`json:"style"`
}

type PresentationAttributes struct {
	AlignmentBaseline			string	`json:"alignmentBaseline"`
	BaselineShift 				string	`json:"baselineShift"`
	Clip 						string	`json:"clip"`
	ClipPath 					string	`json:"clipPath"`
	ClipRule 					string	`json:"clipRule"`
	Color 						string	`json:"color"`
	ColorInterpolation 			string	`json:"colorInterpolation"`
	ColorInterpolationFilters	string	`json:"colorInterpolationFilters"`
	ColorProfile				string	`json:"colorProfile"`
	ColorRendering				string	`json:"colorRendering"`
	Cursor						string	`json:"cursor"`
	Direction					string	`json:"direction"`
	Display						string	`json:"display"`
	DominantBaseline			string	`json:"dominantBaseline"`
	EnableBackground			string	`json:"enableBackground"`
	Fill						string	`json:"fill"`
	FillOpacity					string	`json:"fillOpacity"`
	FillRule					string	`json:"fillRule"`
	Filter						string	`json:"filter"`
	FloodColor					string	`json:"floodColor"`
	FloodOpacity				string	`json:"floodOpacity"`
	FontFamily					string	`json:"fontFamily"`
	FontSize					string	`json:"fontSize"`
	FontSizeAdjust				string	`json:"fontSizeAdjust"`
	FontStretch					string	`json:"fontStretch"`
	FontStyle					string	`json:"fontStyle"`
	FontVariant					string	`json:"fontVariant"`
	FontWeight					string	`json:"fontWeight"`
	GlyphOrientationHorizontal	string	`json:"glyphOrientationHorizontal"`
	GlyphOrientationVertical 	string	`json:"glyphOrientationVertical"`
	ImageRendering 				string	`json:"imageRendering"`
	Kerning						string	`json:"kerning"`
	LetterSpacing				string	`json:"letterSpacing"`
	LightingColor				string	`json:"lightingColor"`
	MarkerEnd					string	`json:"markerEnd"`
	MarkerMid					string	`json:"markerMid"`
	MarkerStart					string	`json:"markerStart"`
	Mask						string	`json:"mask"`
	Opacity						string	`json:"opacity"`
	Overflow					string	`json:"overflow"`
	PointerEvents				string	`json:"pointerEvents"`
	ShapeRendering				string	`json:"shapeRendering"`
	StopColor					string	`json:"stopColor"`
	StopOpacity					string	`json:"stopOpacity"`
	Stroke						string	`json:"stroke"`
	StrokeDasharray				string	`json:"strokeDasharray"`
	StrokeDashoffset			string	`json:"strokeDashoffset"`
	StrokeLinecap				string	`json:"strokeLinecap"`
	StrokeLinejoin				string	`json:"strokeLinejoin"`
	StrokeMiterlimit			string	`json:"strokeMiterlimit"`
	StrokeOpacity				string	`json:"strokeOpacity"`
	StrokeWidth					string	`json:"strokeWidth"`
	TextAnchor					string	`json:"textAnchor"`
	TextDecoration				string	`json:"textDecoration"`
	TextRendering				string	`json:"textRendering"`
	Transform					string	`json:"transform"`
	TransformOrigin				string	`json:"transformOrigin"`
	UnicodeBidi					string	`json:"unicodeBidi"`
	VectorEffect				string	`json:"vectorEffect"`
	Visibility					string	`json:"visibility"`
	WordSpacing					string	`json:"wordSpacing"`
	WritingMode					string	`json:"writingMode"`
}

// TRY TO GET RID OF THE DUPLICATE FIELDS WITH FILTER EFFECT?
type FilterElement struct {
	GlobalAttributes
	FilterEffects	[]filterEffect
}

type filterEffect struct {
	GlobalAttributes
	Which string	`json:"which" validate:"oneof=feBlend feColorMatrix feComponentTransfer feComposite feConvolveMatrix feDiffuseLighting feDisplacementMap feDistantLight feFlood feGaussianBlur feImage feMerge feMorphology feOffset fePointLight feSpecularLighting feSpotLight feTile feTurbulence"`
}

type FeBlend struct {
	In		string	// All In and In2's for SVGs can be one of:
	In2		string	// SourceGraphic | SourceAlpha | BackgroundImage | BackgroundAlpha | FillPaint | StrokePaint | <filter-primitive-reference>
	Mode	string	
}

type FeColorMatrix struct {
	In		string	`json:"in"`
	Type	string	`json:"type" validate:"oneof=matrix saturate hueRotate luminanceToAlpha"`
	Values	string
}

type FeComponentTransfer struct {
	In		string	`json:"in"`
}

type FeComposite struct {
	In			string	`json:"in"`
	In2			string	`json:"in2"`
	Operator	string	`json:"operator" validate:"oneof=over in out atop xor lighter arithmetic"`
	K1			string	`json:"k1" validate:"numeric"`
	K2			string	`json:"k2" validate:"numeric"`
	K3			string	`json:"k3" validate:"numeric"`
	K4			string	`json:"k4" validate:"numeric"`
}

type FeConvolveMatrix struct {
	In					string	`json:"in"`
	Order				string	`json:"order" validate:"numeric"`
	KernelMatrix		string	`json:"kernelMatrix"`
	Divisor				string	`json:"divisor" validate:"numeric"`
	Bias				string	`json:"bias" validate:"numeric"`
	TargetX				string	`json:"targetX" validate:"numeric"` // MUST BE INTEGER!
	TargetY				string	`json:"targetY" validate:"numeric"` // MUST BE INTEGER!
	EdgeMode			string	`json:"edgeMode" validate:"oneof=duplicate wrap none"`	// DEFAULT duplicate
	PreserveAlpha		bool	`json:"preserveAlpha"`			// DEFAULT false
}

type FeDiffuseLighting struct {
	In					string	`json:"in"`
	SurfaceScale		string	`json:"surfaceScale" validate:"numeric"`		// DEFAULT is 1
	DiffuseConstant		string	`json:"diffuseConstant" validate:"numeric"`		// DEFAULT is 1
}

type FeDisplacementMap struct {
	In					string `json:"in"`
	In2					string `json:"in2"`
	Scale				string `json:"scale" validate:"numeric"`
	XChannelSelector	string `json:"xChannelSelector" validate:"oneof=R G B A"`	// DEFAULT "A"
	YChannelSelector	string `json:"yChannelSelector" validate:"oneof=R G B A"`	// DEFAULT "A"
}

type FeDistantLight struct {
	Azimuth		string	`json:"azimuth" validate:"numeric"`
	Elevation	string	`json:"elevation" validate:"numeric"`
}

type FeFlood struct { 
	FloodColor		string `json:"floodColor"`
	FloodOpacity	string `json:"floodOpacity"`
}

type FeGaussianBlur struct { 
	In				string	`json:"in"`
	StdDeviation	string	`json:"stdDeviation"`	
	EdgeMode		string	`json:"edgeMode" validate:"oneof=duplicate wrap none"` // DEFAULTS to duplicate
}

type FeImage struct { 
	PreserveAspectRatio	string	`json:"preserveAspectRatio"`
	Href				string	`json:"href"`
}

type FeMerge struct { 			// Has child FeMergeNode
	// No specific properties
}

type FeMorphology struct { 
	In			string	`json:"in"`
	Operator	string	`json:"operator" validate:"oneof=over in out atop xor lighter arithmetic"`	// DEFAULTS to over
	Radius		string	`json:"radius"`	// Default value 0
}

type FeOffset struct { 
	In	string	`json:"in"`
	Dx	string	`json:"dx" validate:"numeric"`
	Dy	string	`json:"dy" validate:"numeric"`
}

type FePointLight struct { 
	X	string	`json:"x"`
	Y	string	`json:"y"`
	Z	string	`json:"z"`
}

type FeSpecularLighting struct { 
	In					string	`json:"in"`
	SurfaceScale		string	`json:"surfaceScale" validate:"numeric"`
	SpecularConstant	string	`json:"specularConstant" validate:"numeric"` // DEFAULTS TO 1
	SpecularExponent	string	`json:"specularExponent" validate:"numeric"` // DEFAULTS TO 1
}

type FeSpotLight struct { 
	X					string	`json:"x"`
	Y					string	`json:"y"`
	Z					string	`json:"z"`
	PointsAtX			string	`json:"pointsAtX" validate:"numeric"`
	PointsAtY			string	`json:"pointsAtY" validate:"numeric"`
	PointsAtZ			string	`json:"pointsAtZ" validate:"numeric"`
	SpecularExponent	string	`json:"specularExponent" validate:"numeric"`
	LimitingConeAngle	string	`json:"limitingConeAngle" validate:"numeric"`
}

type FeTile struct {
	In	string	`json:"in"`
}

type FeTurbulence struct { 
	BaseFrequency	string	`json:"baseFrequency"`
	NumOctaves		string	`json:"numOctaves" validate:"numeric"`	// MUST BE INTEGER!
	Seed			string	`json:"seed" validate:"numeric"`
	StitchTiles		string	`json:"stitchTiles" validate:"oneof=noStitch stitch"` // DEFAULTS TO noStitch
	Type			string	`json:"type" validate:"oneof=fractalNoise turbulence"` // DEFAULTS TO TURBULENCE
}
