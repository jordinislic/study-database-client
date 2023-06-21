package ContactClient

import (
	"fmt"
	"github.com/jordinislic/utilities/utilities/ClientJordi"
)

type ContactClient struct {
	cl ClientJordi.ClientJordi
}

func New(url string) ContactClient {
	return ContactClient{
		cl: ClientJordi.New(url),
	}
}

func (b ContactClient) List() {
	err := b.cl.Get("/contacts")
	if err != nil {
		fmt.Println(err)
	}
}

func (b ContactClient) Get(name string, surname string) {
	var variables []string
	variables = append(variables, name)
	variables = append(variables, surname)
	err := b.cl.GetWithVariables("/contacts", variables)
	if err != nil {
		fmt.Println(err)
	}
}

func (b ContactClient) GetName(name string) {
	var variables []string
	variables = append(variables, name)
	err := b.cl.GetWithVariables("/contacts", variables)
	if err != nil {
		fmt.Println(err)
	}
}

func (b ContactClient) GetSurname(surname string) {
	var variables []string
	variables = append(variables, surname)
	err := b.cl.GetWithVariables("/contacts", variables)
	if err != nil {
		fmt.Println(err)
	}
}

func (b ContactClient) Add(body string) {
	err := b.cl.Add("/contacts", body)
	if err != nil {
		fmt.Println(err)
	}
}

func (b ContactClient) Delete(name string, surname string) {
	err := b.cl.Delete("/contacts", name, surname)
	if err != nil {
		fmt.Println(err)
	}
}
