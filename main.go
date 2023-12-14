package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

func main() {
	wsURL := launcher.New().Leakless(false).Headless(false).MustLaunch()
	browser := rod.New().ControlURL(wsURL).MustConnect()
	page := stealth.MustPage(browser)

	page.MustNavigate("https://docs.google.com/forms/d/e/1FAIpQLSe3TPcju4fsYew5gLidWZefV0PqEmV5MVT_euWaKzmFyi46HQ/viewform")

	Responden(page)

	// #mG61Hd > div.RH5hzf.RLS9Fe > div > div.ThHDze > div.DE3NNc.CekdCb > div.lRwqcd > div > span > span
	page.MustElementX(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div`).MustClick()
}

func Quiz(page *rod.Page, answer []int) {
	quiz := page.MustElements(`.Od2TWd.hYsg7c`)
	_ = quiz
}

func Responden(page *rod.Page) {
	textbox := page.MustElements(`input.whsOnd.zHQkBf`)
	textbox[0].MustInput("Evan Roy")

	radioBut := page.MustElements(`.Od2TWd.hYsg7c`)
	radioBut[0].MustClick()
	sleepytime()
	fmt.Println(len(radioBut))
	radioBut[2].MustClick()

	radioBut[9].MustClick()
	radioBut[12].MustClick()
	radioBut[14].MustClick()
	sleepytime()
	radioBut[17].MustClick()
	radioBut[20].MustClick()
	sleepytime()
	textbox[1].MustInput("shopee")
}

func sleepytime() {
	time.Sleep(time.Millisecond * 300)
}
