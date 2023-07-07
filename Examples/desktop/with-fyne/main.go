package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var fObject fyne.CanvasObject
var canvas fyne.Canvas

type Resourse struct{}

func (r *Resourse) Name() string {
	return "go app"
}

func (r *Resourse) Content() []byte {
	return []byte{}
}

func main() {
	// Create a new Fyne application
	myApp := app.New()

	// Create a new window
	myWindow := myApp.NewWindow("Golang Desktop App")

	// Create a button widget
	f := widget.NewButton("Click me!", func() {
		fmt.Println("bosildi")
		widget.ShowPopUp(fObject, canvas)
	})

	f2 := widget.NewButton("don't Click me!", func() {
		fmt.Println("bosildi")
		widget.ShowPopUp(fObject, canvas)
	})

	// Create a container to hold the button
	myContent := container.NewCenter(f, f2)

	// Set the content of the window
	myWindow.SetContent(myContent)

	// Show the window
	myWindow.ShowAndRun()
}
