package main

import (
	"log"
	"strings"

	_ "embed"

	"github.com/liamchampton/tech-days/frontend/data"
	"honnef.co/go/js/dom"
)

//go:embed config
var backendURL string

//go:embed countries
var countriesInput string

func main() {
	ds, err := data.NewDataService(strings.TrimSpace(backendURL))
	if err != nil {
		log.Fatal(err)
	}
	countries := strings.Split(countriesInput, "\n")
	document := dom.GetWindow().Document()
	addRowBtn := document.GetElementByID("addRowBtn")
	refreshBtn := document.GetElementByID("refreshBtn")
	if addRowBtn == nil || refreshBtn == nil {
		log.Fatal("action buttons are nil ")
	}
	addRowBtn.AddEventListener("click", true, func(e dom.Event) {
		showUserInput(document)
	})
	populateCountries(document, countries)
	cancelBtn := document.GetElementByID("cancelBtn")
	cancelBtn.AddEventListener("click", true, func(e dom.Event) {
		hideUserInput(document)
	})
	submitBtn := document.GetElementByID("submitBtn")
	submitBtn.AddEventListener("click", true, func(e dom.Event) {
		readUserInput(ds, document)
		hideUserInput(document)
		refreshEntries(ds, document)
	})
	refreshBtn.AddEventListener("click", true, func(e dom.Event) {
		refreshEntries(ds, document)
	})
}

func populateCountries(document dom.Document, countries []string) {
	selEl := document.GetElementByID("countryDropdown").(*dom.HTMLSelectElement)
	for _, c := range countries {
		o := document.CreateElement("option")
		o.SetTextContent(c)
		selEl.AppendChild(o)
	}
}

func readUserInput(ds *data.DataService, document dom.Document) {
	name := document.GetElementByID("nameInput").(*dom.HTMLInputElement)
	countrySelection := document.GetElementByID("countryDropdown").(*dom.HTMLSelectElement)
	countryElement := countrySelection.Options()[countrySelection.SelectedIndex]
	if err := ds.PostEntry(data.DataEntry{
		Name:    name.Value,
		Country: countryElement.Text,
	}); err != nil {
		log.Fatal(err)
	}
}

func refreshEntries(ds *data.DataService, document dom.Document) {
	ds.GetEntries(func(d []data.DataEntry) {
		table := document.GetElementByID("attendeeTable").GetElementsByTagName("tbody")[0]
		ts := table.(*dom.HTMLTableSectionElement)
		for i := 0; i < len(ts.Rows()); i++ {
			ts.DeleteRow(i)
		}
		for _, e := range d {
			populateUser(ts, e)
		}
	})
}

func hideUserInput(document dom.Document) {
	userInput := document.GetElementByID("userInput")
	userInput.Class().Add("d-none")
	name := document.GetElementByID("nameInput").(*dom.HTMLInputElement)
	countrySelection := document.GetElementByID("countryDropdown").(*dom.HTMLSelectElement)
	name.Value = ""
	countrySelection.SelectedIndex = 0
}

func showUserInput(document dom.Document) {
	userInput := document.GetElementByID("userInput")
	userInput.Class().Remove("d-none")
}

func populateUser(tableSection *dom.HTMLTableSectionElement, entry data.DataEntry) {
	row := tableSection.InsertRow(0)
	row.InsertCell(0).SetTextContent(entry.Name)
	row.InsertCell(1).SetTextContent(entry.Country)
}
