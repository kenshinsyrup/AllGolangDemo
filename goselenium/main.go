package main

import (
	"allgolangdemo/goselenium/tsinghuatest"

	"github.com/golang/glog"
)

func main() {
	// ExampleFindElement()
	err := tsinghuatest.TsinghuaPatrol("useremail@applysquare.com", "password")
	glog.Infoln("tsinghuaPatrol err: ", err)
}
