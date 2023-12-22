package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)
	reader.Comma = '|'

	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).Headless(true).MustLaunch()
	page := rod.New().ControlURL(u).MustConnect().MustPage("")

	for i, row := range records {
		PageQuiz(page, row, i)
	}

}

func PageQuiz(page *rod.Page, row []string, i int) {
	dataQuiz1, err := stringsToInts(row[9:49])
	if err != nil {
		log.Fatal(err, i)
	}
	dataQuiz2, err := stringsToInts(row[49:79])
	if err != nil {
		log.Fatal(err, i)
	}
	dataResponden := row[:9]

	page.MustNavigate("https://docs.google.com/forms/d/e/1FAIpQLSe3TPcju4fsYew5gLidWZefV0PqEmV5MVT_euWaKzmFyi46HQ/viewform")

	Responden(page, dataResponden)

	Quiz(page, dataQuiz1)
	Submit(page)
	Quiz(page, dataQuiz2)
	Submit(page)

	fmt.Printf("Responden %s Telah mengisi\n", dataResponden[0])
	rand.NewSource(time.Now().UnixNano())
	randomNumber := time.Duration(rand.Intn(15) + 30)

	time.Sleep(time.Second * randomNumber)
}

func Quiz(page *rod.Page, answer []int) {
	page.MustWaitStable()
	quiz := page.MustElements(`div.Od2TWd.hYsg7c`)

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

func Responden(page *rod.Page, data []string) {

	jk, _ := strconv.Atoi(data[1])
	jurusan, _ := strconv.Atoi(data[2])
	angkatan, _ := strconv.Atoi(data[3])
	online, _ := strconv.Atoi(data[4])
	kegiatan, _ := strconv.Atoi(data[5])
	uangsaku, _ := strconv.Atoi(data[6])
	freq, _ := strconv.Atoi(data[7])

	page.MustWaitStable()

	textbox := page.MustElements(`input.whsOnd.zHQkBf`)
	textbox[0].MustInput(data[0])

	radioBut := page.MustElements(`.Od2TWd.hYsg7c`)
	// Jenis Kelamin
	radioBut[jk].MustClick()
	sleepytime()
	radioBut[jurusan].MustClick()

	radioBut[angkatan].MustClick()
	radioBut[online].MustClick()
	sleepytime()
	radioBut[kegiatan].MustClick()
	radioBut[uangsaku].MustClick()
	sleepytime()
	radioBut[freq].MustClick()
	textbox[1].MustInput(data[8])
	sleepytime()
	page.MustElementX(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div`).MustClick()
}

func sleepytime() {
	time.Sleep(time.Millisecond * 300)
}

func Submit(page *rod.Page) {
	page.MustElementX(`//*[@id="mG61Hd"]/div[2]/div/div[3]/div[1]/div[1]/div[2]`).MustClick()
}

func stringsToInts(strSlice []string) ([]int, error) {
	intSlice := make([]int, len(strSlice))

	for i, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("error converting string '%s' to int: %v", str, err)
		}
		intSlice[i] = num
	}

	return intSlice, nil
}
