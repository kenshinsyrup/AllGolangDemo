package main

import (
	"fmt"
	"runtime"
)

// fooCallers() wants to know who called it
func fooCallers(callserSkip int) {
	fpcs := make([]uintptr, 1)
	// Skip 2 levels to get the caller
	n := runtime.Callers(callserSkip, fpcs)
	if n == 0 {
		fmt.Println("MSG: NO CALLER")
	}

	// FuncForPC deprecated
	// caller := runtime.FuncForPC(fpcs[0] - 1)
	// if caller == nil {
	// 	fmt.Println("MSG CALLER WAS NIL")
	// }
	// // Print the file name and line number
	// fmt.Println(caller.FileLine(fpcs[0] - 1))

	// 	// Print the name of the function
	// 	fmt.Println(caller.Name())

	frames := runtime.CallersFrames(fpcs)
	if frames == nil {
		fmt.Println("MSG CALLER WAS NIL")
	}

	for {
		f, more := frames.Next()
		fmt.Println("function name: ", f.Function)
		fmt.Println("file name: ", f.File)
		fmt.Println("line number: ", f.Line)
		fmt.Println("********")
		if !more {
			break
		}
	}
}

func fooCaller(callerSkip int) {
	pc, file, line, ok := runtime.Caller(callerSkip)
	if !ok {
		fmt.Println("Failed")
		return
	}
	fmt.Println("program counter: ", pc)
	fmt.Println("file name: ", file)
	fmt.Println("line number: ", line)
	fmt.Println("&&&&&")
}

func wraper(f func(int), skip int) {
	f(skip)
}

// 总结
// 根据go文档的说法，runtime.Caller与runtime.Callers的skip参数由于历史原因，其指代的内容是不同的

func main() {
	// callers skip 1 指向的栈是runtime.Callers函数被调用的地方
	fooCallers(1)
	// callers skip 2 指向的栈是直接含了runtime.Callers函数调用的函数，这里是fooCallers，被调用的地方
	fooCallers(2)
	// 验证：将fooCaller包含在一个wraper函数中，skip 2还是可以正确的找到fooCallers被调用的地方，即58行
	wraper(fooCallers, 2)

	// caller skip 0 指向的栈是runtime.Caller函数被调用的地方
	fooCaller(0)
	// caller skip 1 指向的栈是直接包含了runtime.Caller函数调用的函数，这里是fooCaller，被调用的地方
	fooCaller(1)
	// 验证：将fooCaller包含在一个wraper函数中，skip 1还是可以正确的找到fooCaller被调用的地方，即58行
	wraper(fooCaller, 1)

	// 所以，假设wrapper是一个error报告函数，那么我们希望得到wrapper的调用点，而wrapper内部的f函数才是真正的记录调用栈信息时
	// 可以将skip按需要向上增加即可
	// 得到wrapper被调用处的信息，即82行
	wraper(fooCaller, 2)
}
