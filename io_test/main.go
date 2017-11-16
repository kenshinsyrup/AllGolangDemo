package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// open a file fot test
	f, e := os.Open("test.txt")
	if e != nil {
		log.Println("open file err: ", e)
	}
	// get file info
	fInfo, e := f.Stat()
	if e != nil {
		log.Println("state file err: ", e)
	}
	log.Println("size now: ", fInfo.Size())

	// 测试1: copy file 到 buffer, 大小和内容都如预期
	var buf bytes.Buffer
	_, e = io.Copy(&buf, f)
	if e != nil {
		log.Println("copy to buffer err: ", e)
	}
	log.Println("buf content:", buf.String())
	log.Println("buf size: ", buf.Len())
	fmt.Println("测试1 end.")

	// 测试2: 将原buffer的内容拷贝到新buffer，如预期，的原buffer被io.Copy方法读取完，size变为0.
	// 验证了io.Copy并不是复制src到dst，而是读出src的所有内容给到dst。这里的原因是io.Copy自身的注释定义：
	// If src implements the WriterTo interface, the copy is implemented by calling src.WriteTo(dst).
	var buff bytes.Buffer
	io.Copy(&buff, &buf)
	log.Println("now buf size: ", buf.Len())
	log.Println("buff size: ", buff.Len())
	fmt.Println("测试2 end.")

	// 测试3: 验证测试2
	var bufff bytes.Buffer
	(&buff).WriteTo(&bufff)
	log.Println("buff size: ", buff.Len())
	log.Println("bufff size: ", bufff.Len())
	fmt.Println("测试3 end.")

	// 测试4
	// 现在看文件的测试，基于测试2，我们知道，io.Copy使用的是src的WriteTo方法
	// 因此如果我们直接再次使用io.Copy操作file，势必会发现无法读取任何内容，尽管file对size并没有变化也就是说file还是有内容的
	fInfo, e = f.Stat()
	if e != nil {
		log.Println("state file err: ", e)
	}
	log.Println("size now: ", fInfo.Size())
	var fBuf bytes.Buffer
	n, e := io.Copy(&fBuf, f)
	if e != nil {
		log.Println("copy file to buffer err: ", e)
	}
	log.Println("copy file to buffer size: ", n)
	log.Println("buffer size: ", fBuf.Len())
	log.Println("buff content: ", fBuf.String())
	fmt.Println("测试4 end.")

	// 测试5
	// 这就是file和普通的io.Reader不同的地方，其WriteTo，也就是被io.Copy之后，只是把文件内部对读取指针进行了移动
	// 我们恢复文件对读取指针对位置
	_, e = f.Seek(0, 0)
	if e != nil {
		log.Println("file seek err: ", e)
	}
	fInfo, e = f.Stat()
	if e != nil {
		log.Println("state file err: ", e)
	}
	log.Println("size now: ", fInfo.Size())

	var fBuff bytes.Buffer
	n, e = io.Copy(&fBuff, f)
	if e != nil {
		log.Println("copy file to buffer err: ", e)
	}
	log.Println("copy file to buffer size: ", n)
	log.Println("buffer size: ", fBuff.Len())
	log.Println("buff content: ", fBuff.String())
	fmt.Println("测试5 end.")

	log.Println("Done!")
}
