package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var inTheEnd = `In The End -- Linkin Park(林肯公园)
It starts with one thing
I don't know why
It doesn't even matter how hard you try
Keep that in mind
I designed this rhyme
To explain in due time
All I know
Time is a valuable thing
Watch it fly by as the pendulum swings
Watch it count down to the end of the day
The clock ticks life away

It's so unreal
Didn't look out below
Watch the time go right out the window
Trying to hold on, but you didn't even know
Wasted it all just to watch you go
I kept everything inside
And even though I tried, it all fell apart
What it meant to me
Will eventually be a memory of a time when

I tried so hard
And got so far
But in the end
It doesn't even matter
I had to fall
To lose it all
But in the end
It doesn't even matter

One thing, I don't know why
It doesn't even matter how hard you try
Keep that in mind
I designed this rhyme
To remind myself of a time when
I tried so hard
In spite of the way you were mocking me
Acting like I was part of your property
Remembering all the times you fought with me
I'm surprised it got so
Things aren't the way they were before
You wouldn't even recognize me anymore
Not that you knew me back then
But it all comes back to me in the end
You kept everything inside
And even though I tried, it all fell apart
What it meant to me will eventually be a memory of a time when

I tried so hard
And got so far
But in the end
It doesn't even matter
I had to fall
To lose it all
But in the end
It doesn't even matter

I've put my trust in you
Pushed as far as I can go
For all this
There's only one thing you should know
I've put my trust in you
Pushed as far as I can go
For all this
There's only one thing you should know

I tried so hard
And got so far
But in the end
It doesn't even matter
I had to fall
To lose it all
But in the end
It doesn't even matter`

func main() {
	loadFont("./impact.svg")
	ret := textToPath(inTheEnd, 20)
	// the file here is for debug, vscode integrated terminal has a stupid bug, will lose space when string is too long and is wrapped to next line
	ioutil.WriteFile("testfile1.xml", []byte(ret), 0777)
}

var unitsPerEm float32
var ascent float32
var descent float32
var widthCorrection float32

func utf8ToUnicode(str string) []string {
	rs := []rune(str)
	var unicodeNums []string
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			unicodeNums = append(unicodeNums, fmt.Sprintf("%d", rint))
		} else {
			unicodeNums = append(unicodeNums, strconv.FormatInt(int64(rint), 16))
		}
	}
	return unicodeNums
}

type GlyphStruct struct {
	D         string `xml:"d,attr"`
	Unicode   string `xml:"unicode,attr"`
	GlyphName string `xml:"glyph-name,attr"`
	HorizAdvX string `xml:"horiz-adv-x,attr"`
}

type FontFaceStruct struct {
	Ascent string `xml:"ascent,attr"`
	//  bbox string `xml:"bbox,attr"`  ="102" ="U+0020-FB02" ="2048" ="1327"
	//  cap-height string `xml:"cap-height,attr"`
	Descent     string `xml:"descent,attr"`
	UnitsPerEm  string `xml:"units-per-em,attr"`
	FontFamily  string `xml:"font-family,attr"`
	FontStretch string `xml:"font-stretch,attr"`
	//  font-weight string `xml:"font-weight,attr"`
	//  panose-1 string `xml:"panose-1,attr"`
	//  underline-position string `xml:"underline-position=,attr"`
	//  underline-thickness string `xml:"underline-thickness,attr"`
	//  unicodeRange string `xml:"unicode-range,attr"`
	//  x-height string `xml:"x-height,attr"`
}

type FontStruct struct {
	// attr
	HorizAdvX string `xml:"horiz-adv-x,attr"`
	ID        string `xml:"id,attr"`
	// field
	FontFace FontFaceStruct `xml:"font-face"`
	Glyphs   []GlyphStruct  `xml:"glyph"`
}

type DefsStruct struct {
	// fields
	Font FontStruct `xml:"font"`
}

type SvgStruct struct {
	// svg attr
	// Height string `xml:"height,attr"`
	// Width  string `xml:"width,attr"`
	// Xmlns  string `xml:"xmlns,attr"`

	// svg fields
	// Metadata string     `xml:"metadata"`
	Defs DefsStruct `xml:"defs"`
}

type XMLStruct struct {
	Svg SvgStruct `xml:"svg"`
}

type glyph struct {
	d         string
	horizAdvX string
}

var glyphMap = make(map[string]glyph)

func loadFont(filePath string) {
	var svg SvgStruct
	bytes, err := ioutil.ReadFile("./Microsoft-Yahei.svg")
	if err != nil {
		panic(err)
	}
	xml.Unmarshal(bytes, &svg)

	v, err := strconv.ParseFloat(svg.Defs.Font.FontFace.UnitsPerEm, 32)
	if err != nil {
		panic(err)
	}
	unitsPerEm = float32(v)

	v, err = strconv.ParseFloat(svg.Defs.Font.FontFace.Ascent, 32)
	if err != nil {
		panic(err)
	}
	ascent = float32(v)

	v, err = strconv.ParseFloat(svg.Defs.Font.FontFace.Descent, 32)
	if err != nil {
		panic(err)
	}
	descent = float32(v)

	fmt.Println("font familly: ", svg.Defs.Font.FontFace.FontFamily)
	if svg.Defs.Font.FontFace.FontStretch == "condensed" {
		widthCorrection = -50
	}

	for _, g := range svg.Defs.Font.Glyphs {
		if g.Unicode == "" {
			continue
		}

		mapG := glyph{
			d:         g.D,
			horizAdvX: g.HorizAdvX,
		}
		if g.HorizAdvX == "" {
			mapG.horizAdvX = svg.Defs.Font.HorizAdvX
		}
		glyphMap[g.Unicode] = mapG
	}
}

func lineToPath(line string, fillColor string, x, y, asize float32) string {
	var advance float32
	var content []string
	var previousGlyph *glyph
	size := asize / float32(unitsPerEm)

	for _, char := range line {
		if previousGlyph != nil {
			value, err := strconv.ParseFloat(previousGlyph.horizAdvX, 32)
			if err != nil {
				panic(err)
			}
			advance += float32(value) + widthCorrection
		}
		key := string(char)
		fmt.Println("key: ", key)
		g := glyphMap[key]
		previousGlyph = &g
		content = append(content, fmt.Sprintf(`<path data-glyph="%s" transform="translate(%f) scale(1, -1)" d="%s" />`, key, advance, g.d))
	}

	result := fmt.Sprintf(`<g data-tspan="%s" fill="%s" transform="translate(%f, %f) scale(%f)">
	%s
	</g>`, line, fillColor, x, y, size, strings.Join(content, "\n"))

	return result
}

func textToPath(text string, asize float32) string {
	lines := strings.Split(text, "\n")
	var horizAdvY float32
	var linePaths []string
	size := asize / unitsPerEm
	fmt.Println("unitsPerEm: ", unitsPerEm)
	fmt.Println("size: ", size)
	fmt.Println("ascent: ", ascent, " descent: ", descent)
	fmt.Println("y: ", ascent+descent)

	for _, line := range lines {
		fmt.Println(line)
		var horizAdvX float32
		var paths []string
		// horizAdvY += ascent + descent
		// horizAdvY += ascent + descent + unitsPerEm
		horizAdvY += unitsPerEm

		for _, c := range line {
			char := string(c)
			g := glyphMap[char]
			paths = append(paths, fmt.Sprintf("<path transform=\"translate(%f) rotate(180) scale(-1, 1)\" d=\"%s\" />", horizAdvX, g.d))
			v, err := strconv.ParseFloat(g.horizAdvX, 32)
			if err != nil {
				continue
				// panic(err)
			}
			horizAdvX += float32(v)
		}
		linePaths = append(linePaths, fmt.Sprintf("<g transform=\"scale(%f) translate(0,%f)\">%s</g>", size, horizAdvY, strings.Join(paths, "\n")))
	}

	return strings.Join(linePaths, "\n")
}
