package newsfeed

//Interfaces that allow us to interact with the newsfeed
//as function parameters
//They allow us to simply mock the instances instead of using them as pointers.

type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

//Definition of the item object
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

//Definition of the newsfeed (as a repository containing items)

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

//Functions for GET and POST requests

func (r *Repo) GetAll() []Item {
	return r.Items
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}
