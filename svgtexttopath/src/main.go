package main

import (
	"allgolangdemo/svgtexttopath/src/data"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	fontFile := "./data/Microsoft-Yahei.svg" //good
	// fontFile := "./data/times.svg" //times font dont support Chinese
	// fontFile := "./data/impact.svg" // this font file is from a rubby repo, but it may have problem about rotate. dont use it.
	loadFont(fontFile)

	ret := textToPath(data.InTheEnd, 0, 0, 20)
	// the file here is for debug, vscode integrated terminal has a stupid bug, will lose space when string is too long and is wrapped to next line
	ioutil.WriteFile(fmt.Sprintf("./ret/%v.html", time.Now()), []byte(ret), 0777)
}

var unitsPerEm float32

type glyphStruct struct {
	// attr
	D         string `xml:"d,attr"`
	Unicode   string `xml:"unicode,attr"`
	GlyphName string `xml:"glyph-name,attr"`
	HorizAdvX string `xml:"horiz-adv-x,attr"`
}

type fontFaceStruct struct {
	// attr
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

type fontStruct struct {
	// attr
	HorizAdvX string `xml:"horiz-adv-x,attr"`
	ID        string `xml:"id,attr"`
	// field
	FontFace fontFaceStruct `xml:"font-face"`
	Glyphs   []glyphStruct  `xml:"glyph"`
}

type svgStruct struct {
	Font fontStruct `xml:"defs>font"`
}

type glyph struct {
	d         string
	horizAdvX string
}

var glyphMap = make(map[string]glyph)

func loadFont(filePath string) {
	var svg svgStruct
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(bytes, &svg)
	if err != nil {
		panic(err)
	}

	v, err := strconv.ParseFloat(svg.Font.FontFace.UnitsPerEm, 32)
	if err != nil {
		panic(err)
	}
	unitsPerEm = float32(v)

	for _, g := range svg.Font.Glyphs {
		// <glyph> has no unicode symbol
		if g.Unicode == "" {
			continue
		}

		mapG := glyph{
			d:         g.D,
			horizAdvX: g.HorizAdvX,
		}
		if g.HorizAdvX == "" {
			mapG.horizAdvX = svg.Font.HorizAdvX
		}
		glyphMap[g.Unicode] = mapG
	}
}

func textToPath(text string, svgWidth, svgHeight, asize float32) string {
	lines := strings.Split(text, "\n")
	var horizAdvY float32
	var linePaths []string
	size := asize / unitsPerEm
	fmt.Println("unitsPerEm: ", unitsPerEm)
	fmt.Println("size: ", size)

	var biggestX float32
	var biggestY float32
	for _, line := range lines {
		var horizAdvX float32
		var paths []string
		horizAdvY += unitsPerEm
		if biggestY < horizAdvY {
			biggestY = horizAdvY
		}
		for _, c := range line {
			char := string(c)
			g := glyphMap[char]
			paths = append(paths, fmt.Sprintf(`        <path transform="translate(%f) rotate(180) scale(-1, 1)" d="%s" />`, horizAdvX, g.d))
			v, err := strconv.ParseFloat(g.horizAdvX, 32)
			if err != nil {
				panic(err)
			}
			horizAdvX += float32(v)
			if biggestX < horizAdvX {
				biggestX = horizAdvX
			}
		}
		linePaths = append(linePaths, fmt.Sprintf(`    <g transform="scale(%f) translate(0,%f)">
%s
</g>`, size, horizAdvY, strings.Join(paths, "\n")))
	}
	if biggestX*size > svgWidth {
		svgWidth = biggestX * size
	}
	if biggestY*size > svgHeight {
		svgHeight = biggestY * size
	}

	svgTag := fmt.Sprintf(`<svg height="%fpx" width="%fpx">
%s
</svg>`, svgHeight, svgWidth, strings.Join(linePaths, "\n"))

	return svgTag
}
