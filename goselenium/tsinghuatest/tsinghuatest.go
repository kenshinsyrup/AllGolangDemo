package tsinghuatest

import (
	"fmt"

	"sourcegraph.com/sourcegraph/go-selenium"

	"time"

	"errors"

	"github.com/golang/glog"
)

func TsinghuaPatrol(userEmail, userPassword string) error {
	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	if webDriver, err = selenium.NewRemote(caps, "http://192.168.99.100:4444/wd/hub"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return err
	}
	defer webDriver.Quit()

	// get tsinghua webpage
	err = webDriver.Get("http://goglobal.tsinghua.edu.cn/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return err
	}
	// pageSource, err := webDriver.PageSource()
	// if err != nil {
	// 	return err
	// }
	// fmt.Println("page source:", pageSource)
	if title, err := webDriver.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		fmt.Printf("Failed to get page title: %s", err)
		return err
	}

	handles, err := webDriver.WindowHandles()
	if err != nil {
		fmt.Println("handle err: ", err)
		return err
	}
	fmt.Println("handle: ", handles)

	status, err := webDriver.FindElement(selenium.ByXPATH, "//div[@id='navbar']/div/a")
	if err != nil {
		fmt.Println("find element status err: ", err)
		return err
	}
	statusContent, err := status.Text()
	if err != nil {
		glog.Infoln("check status content err: ", err)
		return err
	}
	glog.Infoln("status text: ", statusContent)
	if statusContent != "Login/Register" {
		glog.Infoln("should logout first")
		// TODO: logout
	}
	err = status.Click()
	if err != nil {
		glog.Infoln("click status err: ", err)
		return err
	}

	// login
	email, err := webDriver.FindElement(selenium.ByXPATH, "//input[@type='email']")
	if err != nil {
		fmt.Println("find element email err: ", err)
		return err
	}
	err = email.SendKeys(userEmail)
	if err != nil {
		fmt.Println("send keys err: ", err)
		return err
	}
	pwd, err := webDriver.FindElement(selenium.ByXPATH, "//input[@type='password']")
	if err != nil {
		fmt.Println("find element pwd err: ", err)
		return err
	}
	err = pwd.SendKeys(userPassword)
	if err != nil {
		fmt.Println("send pwd err: ", err)
		return err
	}
	form, err := webDriver.FindElement(selenium.ByXPATH, "//div[@id='react-root']/div/form")
	if err != nil {
		fmt.Println("get email&pwd form err: ", err)
		return err
	}
	err = form.Submit()
	if err != nil {
		glog.Infoln("submit form err: ", err)
		return err
	}

	// wait for login，等待一段时间，dom树更新的比较慢，最初采用的方法是time.Sleep 2秒，但是这样有两个不好的地方：1、如果3sredirect成功，
	// 结果我们报错了，所以最好指定由外部超时时间；2、如果1s就完成了redirect，那么等待2s太浪费了。最后该用了waitForRedirect函数的形式来完成
	// time.Sleep(2 * time.Second)
	err = waitForRedirect(time.Second*5, func() (bool, error) {
		// check if login successful
		userElem, err := webDriver.FindElement(selenium.ByXPATH, "//div[@id='navbar']/div/div/a")
		if err != nil {
			glog.Infoln("get user element err: ", err, ", continue")
			return false, nil
		}
		name, err := userElem.Text()
		if err != nil {
			glog.Infoln("get user name err: ", err)
			return false, err
		}
		glog.Infoln("user name: ", name)
		if name == "testxiaohao" {
			return true, nil
		}
		return false, fmt.Errorf("user name should be: testxiaohao, actual: %s", name)
	})
	if err != nil {
		glog.Infoln("waitForRedirect login to home err: ", err)
		return err
	}

	// Home
	homeElem, err := webDriver.FindElement(selenium.ByXPATH, "//div[@class='header home']/nav/div/div[2]/ul/li[1]/a")
	if err != nil {
		glog.Infoln("get home element err: ", err)
		return err
	}
	homeContent, err := homeElem.Text()
	if err != nil {
		glog.Infoln("get home content err: ", err)
		return err
	}
	glog.Infoln("home content: ", homeContent)

	// Overview
	//  International Education at Tsinghua
	// dropdown的子级内容是invisible的，如果直接获取该元素（不会出错）之后直接点击会报错：Element not visible。
	// 解决办法是曲线救国：首先获取dropdown的那个元素，也就是实际要获取的元素的父级元素，然后调用其moveto方法，该方法是移动鼠标，而其参数是相对位置，因此0，0参数就是移动到其自身，相当于鼠标的hover动作，从而激活了该dropdown
	/*overviewElem, err := webDriver.FindElement(selenium.ByXPATH, "//div[@class='header home']/nav/div/div[2]/ul/li[2]")
	if err != nil {
		glog.Infoln("get overview element err: ", err)
		return err
	}
	// hover to active dropdown menu
	err = overviewElem.MoveTo(0, 0)
	if err != nil {
		glog.Infoln("move to element err: ", err)
		return err
	}*/
	// move to overview元素，来激活下拉菜单
	err = moveToElem(webDriver, selenium.ByXPATH, "//div[@class='header home']/nav/div/div[2]/ul/li[2]")
	if err != nil {
		glog.Infoln("move to elem err: ", err)
		return err
	}

	internationalElem, err := webDriver.FindElement(selenium.ByXPATH, "//div[@class='header home']/nav/div/div[2]/ul/li[2]/ul/li[1]/a")
	if err != nil {
		glog.Infoln("get internationalElem err: ", err)
		return err
	}
	content, err := internationalElem.Text()
	if err != nil {
		glog.Infoln("get international content  err: ", err)
		return err
	}
	glog.Infoln("international content: ", content)
	glog.Infoln("*****************")
	err = internationalElem.Click()
	if err != nil {
		glog.Infoln("click international err: ", err)
		return err
	}
	// time.Sleep(2 * time.Second)
	// currentURL, err := webDriver.CurrentURL()
	// if err != nil {
	// 	glog.Infoln("international currentURL err: ", err)
	// 	return err
	// }
	// glog.Infoln("international currentURL: ", currentURL)
	err = waitForRedirect(5*time.Second, func() (bool, error) {
		currentURL, err := webDriver.CurrentURL()
		if err != nil {
			glog.Infoln("international currentURL err: ", err)
			return false, err
		}
		glog.Infoln("international currentURL: ", currentURL)
		if currentURL != "http://goglobal.tsinghua.edu.cn/overview/global_vision" {
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		glog.Infoln("wait for redirect to intertional err: ", err)
		return err
	}

	// move to overview元素，来激活下拉菜单，注意由于我们上面已经跳转到了internation页面，因此xpath变了
	// study and research abroad
	err = moveToElem(webDriver, selenium.ByXPATH, "//div[@class='header overview']/nav/div/div[2]/ul/li[2]")
	if err != nil {
		glog.Infoln("move to elem err: ", err)
		return err
	}
	studyElem, err := webDriver.FindElement(selenium.ByXPATH, "//div[@class='header overview']/nav/div/div[2]/ul/li[2]/ul/li[2]/a")
	if err != nil {
		glog.Infoln("get studyElem err: ", err)
		return err
	}
	content, err = studyElem.Text()
	if err != nil {
		glog.Infoln("get study content err: ", err)
		return err
	}
	glog.Infoln("study content: ", content)
	err = studyElem.Click()
	if err != nil {
		glog.Infoln("click study err: ", err)
		return err
	}
	time.Sleep(2 * time.Second)

	currentURL, err := webDriver.CurrentURL()
	if err != nil {
		glog.Infoln("study url err: ", err)
		return err
	}
	glog.Infoln("study url: ", currentURL)

	// user center
	userElem, err := webDriver.FindElement(selenium.ByXPATH, "//div[@id='navbar']/div/div")
	if err != nil {
		glog.Infoln("user elem err: ", err)
		return err
	}
	err = userElem.Click()
	if err != nil {
		glog.Infoln("user click err: ", err)
		return err
	}
	usercenterElem, err := webDriver.FindElement(selenium.ByXPATH, "//div[@id='navbar']/div/div/ul/li[1]")
	if err != nil {
		glog.Infoln("user center elem err: ", err)
		return err
	}
	err = usercenterElem.Click()
	if err != nil {
		glog.Infoln("user center click err: ", err)
		return err
	}
	time.Sleep(2 * time.Second)
	currentURL, err = webDriver.CurrentURL()
	if err != nil {
		glog.Infoln("current URL err: ", err)
		return err
	}
	glog.Infoln("currentURL: ", currentURL)

	// home页面的explore中点击undergraduate

	return nil
}

func moveToElem(webDriver selenium.WebDriver, by, path string) error {
	elem, err := webDriver.FindElement(by, path)
	if err != nil {
		glog.Infoln("find element err: ", err)
		return err
	}
	return elem.MoveTo(0, 0)
}

func waitForRedirect(timeout time.Duration, check func() (bool, error)) error {
	now := time.Now()
	for {
		interval := time.Now().Sub(now)
		if interval > timeout {
			return errors.New("time out")
		}

		ok, err := check()
		if err != nil {
			glog.Infoln("check err: ", err)
			return err
		}
		if ok {
			glog.Infoln("operation time: ", interval)
			break
		}
	}
	return nil
}
