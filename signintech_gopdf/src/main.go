package main

import (
	"fmt"
	"log"

	"github.com/signintech/gopdf"
)

/*
直接操作PDF太难了，虽然可以画图，这个库目测不错
但是没有模版的情况下，无法使用手写坐标画图的方式来达到画出正确的表格并调整内容格式的要求
*/

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("Microsoft Yahei", "./ttf/microsoft_yahei.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("Microsoft Yahei", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	/*
		err := pdf.AddTTFFont("HDZB_5", "../ttf/wts11.ttf")
		if err != nil {
			log.Print(err.Error())
			return
		}
		err = pdf.SetFont("HDZB_5", "", 14)
		if err != nil {
			log.Print(err.Error())
			return
		}
	*/
	// pdf.SetGrayFill(0.5)
	pdf.Cell(nil, "您好")
	pdf.Br(18)

	// try to draw a form, corordinate system origional point is upper left
	currentX := pdf.GetX()
	currentY := pdf.GetY()
	fmt.Println("currentX: ", currentX)
	fmt.Println("currentY: ", currentY)

	// upper line, width 200
	pdf.Line(currentX, currentY, currentX+200, currentY)
	// left line, height 100
	pdf.Line(currentX, currentY, currentX, currentY-100)
	// right line
	pdf.Line(currentX+200, currentY, currentX+200, currentY-100)
	// lower line
	pdf.Line(currentX, currentY-100, currentX+200, currentY-100)

	pdf.WritePdf("hello_yahei.pdf")
}
