package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/marcosCapistrano/web-1.0/models"
)

type NewContactTemplateData struct {
	Contact models.Contact
	Errors  map[string]string
}
type ContactsTemplateData struct {
	Search   string
	Contacts []models.Contact
}

func main() {
	var allContacts models.Contacts
	allContacts.Init()

	newContactHandler := func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("./templates/contacts/new.html")
		if err != nil {
			panic(err)
		}

		if r.Method == "GET" {
			err = template.Execute(w, NewContactTemplateData{Contact: models.Contact{}, Errors: make(map[string]string)})
			if err != nil {
				fmt.Println("could not execute new contact template: ", err.Error())
				io.WriteString(w, err.Error())
			}
		} else if r.Method == "POST" {
			r.ParseForm()

			form := r.Form
			newContact := models.Contact{
				Email: form.Get("email"),
				Phone: form.Get("phone"),
				First: form.Get("first"),
				Last:  form.Get("last"),
			}

			fmt.Println(form)

			errors := allContacts.AddContact(newContact)

			if len(errors) == 0 {
				http.Redirect(w, r, "/contacts", http.StatusSeeOther)
			} else {
				err = template.Execute(w, NewContactTemplateData{Contact: newContact, Errors: errors})
			}

			if err != nil {
				fmt.Println("could not execute new contact template: ", err.Error())
				io.WriteString(w, err.Error())
			}
		}
	}

	contactsHandler := func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("./templates/contacts/contacts.html")
		if err != nil {
			panic(err)
		}

		var search = ""
		if r.URL.Query().Has("q") {
			search = r.URL.Query().Get("q")
		}

		err = template.Execute(w, ContactsTemplateData{Search: search, Contacts: allContacts})
		if err != nil {
			fmt.Println("could not execute contacts template: ", err.Error())
			io.WriteString(w, err.Error())
		}

	}

	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
	}

	http.HandleFunc("/contacts/new", newContactHandler)
	http.HandleFunc("/contacts", contactsHandler)
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("test")
}
