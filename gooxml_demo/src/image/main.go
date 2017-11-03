// Copyright 2017 Baliance. All rights reserved.
package main

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"log"

	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"

	wml "baliance.com/gooxml/schema/schemas.openxmlformats.org/wordprocessingml"
)

var lorem = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin lobortis, lectus dictum feugiat tempus, sem neque finibus enim, sed eleifend sem nunc ac diam. Vestibulum tempus sagittis elementum`

func main() {
	doc := document.New()

	// img, err := document.ImageFromFile("gophercolor.png")
	// if err != nil {
	// 	log.Fatalf("unable to create image: %s", err)
	// }

	tmplImageName := "gophercolor.png"
	content, err := ioutil.ReadFile(tmplImageName)
	fmt.Println("readfile err: ", err)
	imgBuf := bytes.NewBuffer(content)
	fmt.Println("********************")
	fmt.Println("err: ", err)
	if err != nil {
		fmt.Println("fail to read image ", tmplImageName)
	}
	fmt.Println("size: ", imgBuf.Len())

	img := document.Image{
		Path: tmplImageName + "1",
	}
	imgDec, ifmt, err := image.Decode(imgBuf)
	if err != nil {
		fmt.Println("decode error: ", err)
	}
	img.Format = ifmt
	img.Size = imgDec.Bounds().Size()

	iref, err := doc.AddImage(img)
	if err != nil {
		log.Fatalf("unable to add image to document: %s", err)
	}

	para := doc.AddParagraph()
	anchored, err := para.AddRun().AddDrawingAnchored(iref)
	if err != nil {
		log.Fatalf("unable to add anchored image: %s", err)
	}
	anchored.SetName("Gopher")
	anchored.SetSize(2*measurement.Inch, 2*measurement.Inch)
	anchored.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVTopMargin)
	anchored.SetHAlignment(wml.WdST_AlignHCenter)
	anchored.SetYOffset(3 * measurement.Inch)
	anchored.SetTextWrapSquare(wml.WdST_WrapTextBothSides)

	run := para.AddRun()
	for i := 0; i < 16; i++ {
		run.AddText(lorem)
	}
	doc.SaveToFile("image.docx")
}
