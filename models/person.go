package models

type Company struct {
	Name      string
	WebSite   string
}

type Person struct {
	FirstName string
	LastName  string
	NickName  string
	Companies   []Company //บุคคล 1 คนอาจมีหลายบริษัท
	Title     string
	Mobile    string
	Email     string
}