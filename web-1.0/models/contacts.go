package models

type Contact struct {
	ID    int
	First string
	Last  string
	Phone string
	Email string
}

type Contacts []Contact

var ids = 0

func (cs Contacts) Init() Contacts {
	cs = append(cs, Contact{0, "marcos", "Capistrano", "(35)99899-2632", "mcapistrano50@gmail.com"})
	cs = append(cs, Contact{1, "Marcos", "Capistrano", "(35)99899-2632", "mcapistrano50@gmail.com"})
	cs = append(cs, Contact{2, "Marcos", "Capistrano", "(35)99899-2632", "mcapistrano50@gmail.com"})
	cs = append(cs, Contact{3, "Marcos", "Capistrano", "(35)99899-2632", "mcapistrano50@gmail.com"})
	cs = append(cs, Contact{4, "Marcos", "Capistrano", "(35)99899-2632", "mcapistrano50@gmail.com"})
	ids = 5

	return cs
}

func (cs *Contacts) AddContact(contact Contact) map[string]string {
	var errors map[string]string = make(map[string]string)

	if contact.Email == "" {
		errors["email"] = "Email deve ser preenchido"
	}

	if contact.Phone == "" {
		errors["phone"] = "Phone deve ser preenchido"
	}

	if contact.First == "" {
		errors["first"] = "First deve ser preenchido"
	}

	if contact.Last == "" {
		errors["last"] = "Last deve ser preenchido"
	}

	if contact.Email != "" &&
		contact.Phone != "" &&
		contact.First != "" &&
		contact.Last != "" {
		*cs = append(*cs, contact)
	}

	return errors
}

func (cs Contacts) GetAll() Contacts {
	return cs
}

func (cs Contacts) GetByFirst(first string) (bool, Contact) {
	for _, c := range cs {
		if c.First == first {
			return true, c
		}
	}

	return false, Contact{}
}

func (cs Contacts) GetByID(id int) (bool, Contact) {
	for _, c := range cs {
		if c.ID == id {
			return true, c
		}
	}

	return false, Contact{}
}

func (cs Contacts) DeleteByID(id int) Contacts {
	for i, c := range cs {
		if c.ID == id {
			cs[i] = cs[len(cs)-1]
			return cs[:len(cs)-1]
		}
	}

	return Contacts{}
}
