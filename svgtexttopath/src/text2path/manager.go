package text2path

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Manager controls a glyphMap stroring unicode:glyph pairs
type Manager struct {
	unitsPerEm float32
	glyphMap   map[string]glyph
}

// NewManager returns a new manager
func NewManager(fontFile string) *Manager {
	m := &Manager{
		glyphMap: make(map[string]glyph),
	}
	m.loadFont(fontFile)
	return m
}

type glyphStruct struct {
	// attr
	D         string `xml:"d,attr"`
	Unicode   string `xml:"unicode,attr"`
	HorizAdvX string `xml:"horiz-adv-x,attr"`
	// GlyphName string `xml:"glyph-name,attr"`
}

type fontFaceStruct struct {
	// attr
	UnitsPerEm string `xml:"units-per-em,attr"`
	// FontFamily  string `xml:"font-family,attr"`
	// FontStretch string `xml:"font-stretch,attr"`
	// Ascent string `xml:"ascent,attr"`
	// Descent     string `xml:"descent,attr"`
	//  bbox string `xml:"bbox,attr"`  ="102" ="U+0020-FB02" ="2048" ="1327"
	//  cap-height string `xml:"cap-height,attr"`
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
	// ID        string `xml:"id,attr"`
	// field
	FontFace fontFaceStruct `xml:"font-face"`
	Glyphs   []glyphStruct  `xml:"glyph"`
}

type svgStruct struct {
	Font fontStruct `xml:"defs>font"`
}

type glyph struct {
	d         string
	horizAdvX float32
}

func (m *Manager) loadFont(filePath string) {
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
	m.unitsPerEm = float32(v)

	for _, g := range svg.Font.Glyphs {
		// <glyph> has no unicode symbol
		if g.Unicode == "" {
			continue
		}

		horizAdvX := g.HorizAdvX
		if g.HorizAdvX == "" {
			horizAdvX = svg.Font.HorizAdvX
		}
		v, err := strconv.ParseFloat(horizAdvX, 32)
		if err != nil {
			panic(err)
		}
		m.glyphMap[g.Unicode] = glyph{
			d:         g.D,
			horizAdvX: float32(v),
		}
	}
}

// TextToPath returns svg path
func (m *Manager) TextToPath(text string, svgWidth, svgHeight, asize float32) string {
	lines := strings.Split(text, "\n")
	var horizAdvY float32
	var linePaths []string
	size := asize / m.unitsPerEm
	fmt.Println("unitsPerEm: ", m.unitsPerEm)
	fmt.Println("size: ", size)

	var biggestX float32
	var biggestY float32
	for _, line := range lines {
		var horizAdvX float32
		var paths []string
		horizAdvY += m.unitsPerEm
		if biggestY < horizAdvY {
			biggestY = horizAdvY
		}
		for _, c := range line {
			char := string(c)
			g := m.glyphMap[char]
			paths = append(paths, fmt.Sprintf(`        <path transform="translate(%f) rotate(180) scale(-1, 1)" d="%s" />`, horizAdvX, g.d))
			horizAdvX += g.horizAdvX
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
