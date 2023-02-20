package main

import (
	"log"

	"honnef.co/go/js/dom"
)

func main() {
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
		addNewUser(document)
		hideUserInput(document)		
	})
	refreshBtn.AddEventListener("click", true, func(e dom.Event) {
		refreshBtn.SetTextContent("Refresh not implemented yet")
	})
}

func hideUserInput(document dom.Document) {
	userInput := document.GetElementByID("userInput")
	userInput.Class().Add("d-none")
	name := document.GetElementByID("nameInput").(*dom.HTMLInputElement)
	country := document.GetElementByID("countryInput").(*dom.HTMLInputElement)
	funFact := document.GetElementByID("funFactInput").(*dom.HTMLTextAreaElement)
	name.Value = ""
	country.Value = ""
	funFact.Value = ""
}

func showUserInput(document dom.Document) {
	userInput := document.GetElementByID("userInput")
	userInput.Class().Remove("d-none")
}

func addNewUser(document dom.Document) {
	name := document.GetElementByID("nameInput").(*dom.HTMLInputElement)
	country := document.GetElementByID("countryInput").(*dom.HTMLInputElement)
	funFact := document.GetElementByID("funFactInput").(*dom.HTMLTextAreaElement)
	table := document.GetElementByID("attendeeTable").GetElementsByTagName("tbody")[0]
	st := table.(*dom.HTMLTableSectionElement)
	row := st.InsertRow(len(st.Rows()))
	row.InsertCell(0).SetTextContent("xx")
	row.InsertCell(1).SetTextContent(name.Value)
	row.InsertCell(2).SetTextContent(country.Value)
	row.InsertCell(3).SetTextContent(funFact.Value)
}