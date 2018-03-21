package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	loadFont("./impact.svg")
	ret := textToPath("Simple text", 20)
	fmt.Println(ret)
}

var unitsPerEm int
var ascent int
var descent int

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
	Descent    string `xml:"descent,attr"`
	UnitsPerEm string `xml:"units-per-em,attr"`
	FontFamily string `xml:"font-family,attr"`
	//  font-stretch string `xml:"font-stretch,attr"`
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
	bytes, err := ioutil.ReadFile("./impact.svg")
	if err != nil {
		panic(err)
	}
	xml.Unmarshal(bytes, &svg)
	unitsPerEm, err = strconv.Atoi(svg.Defs.Font.FontFace.UnitsPerEm)
	if err != nil {
		panic(err)
	}
	ascent, err = strconv.Atoi(svg.Defs.Font.FontFace.Ascent)
	if err != nil {
		panic(err)
	}
	descent, err = strconv.Atoi(svg.Defs.Font.FontFace.Descent)
	if err != nil {
		panic(err)
	}
	fmt.Println("font familly: ", svg.Defs.Font.FontFace.FontFamily)

	for _, g := range svg.Defs.Font.Glyphs {
		nums := utf8ToUnicode(g.Unicode)
		if len(nums) == 0 {
			fmt.Println("g.Unicode fati to utf8ToUnicode: ", g.Unicode)
			continue
		}
		var key string
		for _, num := range nums {
			key = key + num
		}

		mapG := glyph{
			d:         g.D,
			horizAdvX: g.HorizAdvX,
		}
		if g.HorizAdvX == "" {
			mapG.horizAdvX = svg.Defs.Font.HorizAdvX
		}
		glyphMap[key] = mapG
	}
}

func textToPath(text string, asize float32) string {
	lines := strings.Split(text, "\n")
	result := ""
	horizAdvY := 0

	for _, line := range lines {
		size := asize / float32(unitsPerEm)
		result = result + fmt.Sprintf("<g transform=\"scale(%f) translate(0, %d)\">", size, horizAdvY)
		horizAdvX := 0
		// for _, unicodeNum := range unicodeNums {
		for _, letter := range line {
			unicodeNums := utf8ToUnicode(string(letter))
			var key string
			for _, num := range unicodeNums {
				key = key + num
			}
			result = result + fmt.Sprintf("<path transform=\"translate(%d,%d) rotate(180) scale(-1, 1)\" d=\"%s\" />", horizAdvX, horizAdvY, glyphMap[key].d)
			x, err := strconv.Atoi(glyphMap[key].horizAdvX)
			if err != nil {
				panic(err)
			}
			horizAdvX += x
		}
		result = result + "</g>"
		horizAdvY += ascent + descent
	}
	return result
}
