package main

import (
	"baliance.com/gooxml/document"
	"github.com/golang/glog"
)

func main() {
	doc, err := document.Open("simple.docx")
	if err != nil {
		glog.Infoln("open err: ", err)
		return
	}

	paragraphs := doc.Paragraphs()
	for _, p := range paragraphs {
		rs := p.Runs()
		for _, r := range rs {
			// t := r.Text()
			r.ReplaceText("Text", "hahahahh", -1)
		}
	}
	doc.SaveToFile("new-simple.docx")
}
