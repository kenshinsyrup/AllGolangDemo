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

注意需要先安装wkhtmltopdf
在Mac端使用homebrew的话:
brew install Caskroom/cask/wkhtmltopdf·
*/

var timetableHTML = `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<style>

body {
 text-align:center;
 font-size: 20pt;
}

#block_container
{
    text-align:center;
}
#bloc1, #bloc2
{
    display:inline;
}

.title {
width: 100%;
text-align: center;
}

table, td, th {
    border: 1px solid black;
}
table.timetable {
    empty-cells: show;
    width:100%;
    border:1px solid black;
    border-collapse: collapse;
    }
.blank_row
{
    height: 10px !important; /* overwrites any other rules */
    background-color: #FFFFFF;
}

</style>
</head>
<body>

<div class="title">
海阳一中2017级教师课表
</div>
<p />

<div class="title" id="block_container" >
<div id = "bloc1">姓名：　孙智广</div>
<div id = "bloc1">年级：　2017级</div>
<div id = "bloc1">学科：　体育</div>
</div>
<p />

<table class="timetable">
<tr>
<th>节次</th>
<th>周一</th>
<th>周二</th>
<th>周三</th>
<th>周四</th>
<th>周五</th>
<th>周日</th>
</tr>

<tr>
<td>早自习</td>
<td>1.0</td>
<td>2.0</td>
<td>3.0</td>
<td>4.0</td>
<td>5.0</td>
<td>7.0</td>
</tr>

<tr class="blank_row">
    <td colspan="7"></td>
</tr>


<tr>
<td>上午第一节</td>
<td>1.1</td>
<td>2.1</td>
<td>3.1</td>
<td>4.1</td>
<td>5.1</td>
<td>7.1</td>
</tr>

<tr>
<td>上午第二节</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>

<tr>
<td>上午第三节</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>体育，1班，操场</td>
</tr>

<tr>
<td>上午第四节</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>

<tr class="blank_row">
    <td colspan="7"></td>
</tr>

<tr>
<td>下午第一节</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>

<tr>
<td>下午第二节</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>

<tr>
<td>下午第三节</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>体育，1班，操场</td>
<td>&nbsp;</td>
</tr>

<tr>
<td>下午第四节</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>没啥事</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>


<tr class="blank_row">
    <td colspan="7"></td>
</tr>

<tr>
<td>晚上第一节</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>

<tr>
<td>晚上第二节</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
</tr>

<tr>
<td>晚上第三节</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>&nbsp;</td>
<td>体育，1班，操场</td>
<td>&nbsp;</td>
</tr>



</table>

</body>
</html>

`

func main() {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return
	}

	/*
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
	*/
	pdfg.PageSize.Set(wkhtml.PageSizeA4)
	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(timetableHTML)))
	fmt.Println(pdfg.ArgString())

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// file bytes
	// fmt.Println(len(pdfg.Bytes()))

	//Your Pdf Name
	err = pdfg.WriteFile("./timetable.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
