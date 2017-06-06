package main

import (
	"fmt"
	"time"

	"sourcegraph.com/sourcegraph/go-selenium"
)



func ExampleFindElement() {
	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	if webDriver, err = selenium.NewRemote(caps, "http://192.168.99.100:4444/wd/hub"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer webDriver.Quit()

	err = webDriver.Get("https://www.applysquare.com/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	if title, err := webDriver.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		fmt.Printf("Failed to get page title: %s", err)
		return
	}

	// var elem selenium.WebElement
	// elem, err = webDriver.FindElement(selenium.ByCSSSelector, ".author")
	// if err != nil {
	// 	fmt.Printf("Failed to find element: %s\n", err)
	// 	return
	// }

	// if text, err := elem.Text(); err == nil {
	// 	fmt.Printf("Author: %s\n", text)
	// } else {
	// 	fmt.Printf("Failed to get text of element: %s\n", err)
	// 	return
	// }

	handles, err := webDriver.WindowHandles()
	if err != nil {
		fmt.Println("handle err: ", err)
		return
	}
	fmt.Println("handle: ", handles)

	status, err := webDriver.FindElement(selenium.ById, "status")
	if err != nil {
		fmt.Println("find element status err: ", err)
		return
	}
	statusChildren, err := status.FindElements(selenium.ByClassName, "js-modal")
	if err != nil {
		fmt.Println("find element login err: ", err)
		return
	}
	login := statusChildren[1]
	err = login.Click()
	if err != nil {
		fmt.Println("click login err: ", err)
		return
	}

	time.Sleep(time.Second * 2)
	handles, err = webDriver.WindowHandles()
	if err != nil {
		fmt.Println("handle err: ", err)
		return
	}
	fmt.Println("handle: ", handles)

	// alert, err := webDriver.AlertText()
	// if err != nil {
	// 	fmt.Println("alert err: ", err)
	// 	return
	// }
	// fmt.Println("alert: ", alert)

	// , err := webDriver.ActiveElement()
	// if err != nil {
	// 	fmt.Println("current active element err: ", err)
	// 	return
	// }

	fmt.Println("*****")
	email, err := webDriver.FindElement(selenium.ByXPATH, "//input[@name='email']")
	if err != nil {
		fmt.Println("find element email err: ", err)
		return
	}
	err = email.SendKeys("kenshinsyrup@gmail.com")
	if err != nil {
		fmt.Println("send keys err: ", err)
		return
	}
	pwd, err := webDriver.FindElement(selenium.ByXPATH, "//input[@name='password']")
	if err != nil {
		fmt.Println("find element pwd err: ", err)
		return
	}
	err = pwd.SendKeys("wangzhao")
	if err != nil {
		fmt.Println("send pwd err: ", err)
		return
	}

	form, err := webDriver.FindElement(selenium.ByXPATH, "//div[@class='modal-content']/form")
	if err != nil {
		fmt.Println("submit find err: ", err)
		return
	}
	// modalFooter.FindElement(selenium.ByClassName, "btn btn-primary")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// err = modalFooter.Click()
	// if err != nil {
	// 	fmt.Println("click err: ", err)
	// 	return
	// }
	err = form.Submit()
	if err != nil {
		fmt.Println("submit err: ", err)
		return
	}

	time.Sleep(time.Second * 3)
	fmt.Println("&&&&&")

	span, err := webDriver.FindElement(selenium.ByXPATH, "//ul[@class='nav navbar-nav navbar-status']/li/a/span")
	if err != nil {
		fmt.Println("span err: ", err)
		return
	}
	text, err := span.Text()
	if err != nil {
		fmt.Println("text err: ", err)
		return
	}
	fmt.Println("span text: ", text)

	// output:
	// Page title: GitHub - sourcegraph/go-selenium: Selenium WebDriver client for Go
	// Author: sourcegraph
}
