package main

import (
	"log"

	_ "embed"

	"github.com/liamchampton/tech-days/frontend/data"
	"honnef.co/go/js/dom"
)

//go:embed config
var backendURL string

func main() {
	ds, err := data.NewDataService(backendURL)
	if err != nil {
		log.Fatal(err)
	}
	document := dom.GetWindow().Document()
	addRowBtn := document.GetElementByID("addRowBtn")
	refreshBtn := document.GetElementByID("refreshBtn")
	if addRowBtn == nil || refreshBtn == nil {
		log.Println("action buttons are nil ")
		return
	}
	addRowBtn.AddEventListener("click", true, func(e dom.Event) {
		showUserInput(document)
	})
	cancelBtn := document.GetElementByID("cancelBtn")
	cancelBtn.AddEventListener("click", true, func(e dom.Event) {
		hideUserInput(document)
	})
	submitBtn := document.GetElementByID("submitBtn")
	submitBtn.AddEventListener("click", true, func(e dom.Event) {
		name := document.GetElementByID("nameInput").(*dom.HTMLInputElement)
		location := document.GetElementByID("locationInput").(*dom.HTMLInputElement)
		funFact := document.GetElementByID("funFactInput").(*dom.HTMLTextAreaElement)
		if err := ds.PostEntry(data.DataEntry{
			Name:     name.Value,
			Location: location.Value,
			Fact:     funFact.Value,
		}); err != nil {
			log.Fatal(err)
		}

		hideUserInput(document)
	})
	refreshBtn.AddEventListener("click", true, func(e dom.Event) {
		ds.GetEntries(func(d []data.DataEntry) {
			table := document.GetElementByID("attendeeTable").GetElementsByTagName("tbody")[0]
			ts := table.(*dom.HTMLTableSectionElement)
			for i:= 0; i < len(ts.Rows()); i++ {
				ts.DeleteRow(i)
			}
			for _, e := range d {
				populateUser(ts, e)
			}
		})
	})
}

func hideUserInput(document dom.Document) {
	userInput := document.GetElementByID("userInput")
	userInput.Class().Add("d-none")
	name := document.GetElementByID("nameInput").(*dom.HTMLInputElement)
	location := document.GetElementByID("locationInput").(*dom.HTMLInputElement)
	funFact := document.GetElementByID("funFactInput").(*dom.HTMLTextAreaElement)
	name.Value = ""
	location.Value = ""
	funFact.Value = ""
}

func showUserInput(document dom.Document) {
	userInput := document.GetElementByID("userInput")
	userInput.Class().Remove("d-none")
}

func populateUser(tableSection *dom.HTMLTableSectionElement, entry data.DataEntry) {
	row := tableSection.InsertRow(len(tableSection.Rows()))
	row.InsertCell(0).SetTextContent(entry.ID)
	row.InsertCell(1).SetTextContent(entry.Name)
	row.InsertCell(2).SetTextContent(entry.Location)
	row.InsertCell(3).SetTextContent(entry.Fact)
}
