package main

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

const refreshInterval = 500 * time.Millisecond

var (
	timeBox *tview.TextView
	app     *tview.Application
)

func currentTimeString() string {
	t := time.Now()
	return fmt.Sprintf(t.Format("Current time is 15:04:05"))
}

// Concurrently update the time
func updateTime() {
	for {
		time.Sleep(refreshInterval)
		app.QueueUpdateDraw(func() {
			timeBox.SetText(currentTimeString())
		})
	}
}

// Function called when the application is exited
func doneFunc(buttonIndex int, buttonLabel string) {
	if buttonLabel == "Quit" {
		app.Stop()
	}
}

func main() {
	// Set up application view
	app = tview.NewApplication()
	infoBox := tview.NewFlex()
	timeBox = tview.NewTextView().SetText(currentTimeString())
	infoBox.AddItem(timeBox, 0, 2, true)

	go updateTime()
	if err := app.SetRoot(infoBox, true).Run(); err != nil {
		panic(err)
	}
}
