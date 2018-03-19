package main

import (
	"fmt"
	"log"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

/*
Let's use this~
PDF什么的都是异端，写好HTML专程PDF多好，要什么格式写什么格式
*/

func main() {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return
	}

	htmlStr := `<!DOCTYPE html>
	<html>
	<head>
	<meta charset="UTF-8">
	</head>
		<body>
			<h1 style="color:red;">This is an html from pdf to test color
			</h1>
			<table style="width:100%">
				<tr>
					<th>第一列</th>
					<th>Second Column</th> 
					<th>第三列</th>
				</tr>
				<tr>
					<td>1</td>
					<td>Two</td> 
					<td>哈哈哈哈哈哈哈哈</td>
				</tr>
				<tr>
					<td>how are you</td>
					<td>good, thank you, you?</td> 
					<td>挺好的</td>
				</tr>
			</table>
		</body>
	</html>`

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlStr)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	err = pdfg.WriteFile("./table.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
