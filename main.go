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

	page.MustElementX(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div`).MustClick()

	jawab := []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
	Quiz(page, jawab)
	Submit(page)
	Quiz(page, jawab)
}

func Quiz(page *rod.Page, answer []int) {
	page.MustWaitStable()
	quiz := page.MustElements(`div.Od2TWd.hYsg7c`)
	fmt.Println(len(quiz) / 3)
	for i, d := range answer {
		if d == 3 {
			d = 1
		} else if d == 1 {
			d = 3
		}
		i += 1
		idxQuiz := (3 * i) - d

		quiz[idxQuiz].MustClick()
	}

}

func Responden(page *rod.Page) {
	page.MustWaitStable()
	textbox := page.MustElements(`input.whsOnd.zHQkBf`)
	textbox[0].MustInput("Evan Roy")

	radioBut := page.MustElements(`.Od2TWd.hYsg7c`)
	radioBut[0].MustClick()
	sleepytime()
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

func Submit(page *rod.Page) {
	page.MustElementX(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div[2]`).MustClick()
	//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div[2]

}
