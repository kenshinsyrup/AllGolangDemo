package main

import (
	"allgolangdemo/svgtexttopath/src/data"
	"allgolangdemo/svgtexttopath/src/text2path"

	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	fontFile := "./data/Microsoft-Yahei.svg" //good
	// fontFile := "./data/times.svg" //times font dont support Chinese
	// fontFile := "./data/impact.svg" // this font file is from a rubby repo, but it may have problem about rotate. dont use it.
	text2pathM := text2path.NewManager(fontFile)
	ret := text2pathM.TextToPath(data.InTheEnd, 0, 0, 20)
	// the file here is for debug, vscode integrated terminal has a stupid bug:
	// fmt.Println(ret) will lose the space when string is too long and is wrapped to next line
	ioutil.WriteFile(fmt.Sprintf("./ret/%v.html", time.Now()), []byte(ret), 0777)
}
