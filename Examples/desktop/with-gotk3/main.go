package main

import (
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK
	gtk.Init(nil)

	// Create a new window
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	
	win.SetTitle("Golang Desktop App")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a button
	button, err := gtk.ButtonNewWithLabel("Click me!")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	button.Connect("clicked", func() {
		dialog := gtk.MessageDialogNew(win, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "Hello, World!")
		dialog.Run()
		dialog.Destroy()
	})

	// Add the button to the window
	win.Add(button)

	// Show all the elements
	win.ShowAll()

	// Run the GTK main loop
	gtk.Main()
}
