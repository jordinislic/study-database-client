package ContactClient

import (
	"encoding/json"
	"fmt"
	"github.com/jordinislic/utilities/utilities/ClientJordi"
)

type Contact struct {
	Id      int
	Name    string
	Surname string
	Number  string
}

type ContactClient struct {
	cl ClientJordi.ClientJordi
}

func New(url string) ContactClient {
	return ContactClient{
		cl: ClientJordi.New(url),
	}
}

func (b ContactClient) List() {
	resp := b.cl.Get("/contacts")
	jsonData := b.cl.GetBodyResp(resp)

	ContactResp := []Contact{}
	err := json.Unmarshal(jsonData, &ContactResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(ContactResp)
}

func (b ContactClient) Get(name string, surname string) {
	var variables []string
	variables = append(variables, name)
	variables = append(variables, surname)
	resp := b.cl.GetWithVariables("/contacts", variables)
	jsonData := b.cl.GetBodyResp(resp)

	ContactResp := Contact{}
	err := json.Unmarshal(jsonData, &ContactResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(ContactResp)
}

func (b ContactClient) GetName(name string) {
	var variables []string
	variables = append(variables, name)
	resp := b.cl.GetWithVariables("/contacts", variables)
	jsonData := b.cl.GetBodyResp(resp)

	ContactResp := []Contact{}
	err := json.Unmarshal(jsonData, &ContactResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(ContactResp)
}

func (b ContactClient) GetSurname(surname string) {
	var variables []string
	variables = append(variables, surname)
	resp := b.cl.GetWithVariables("/contacts", variables)
	jsonData := b.cl.GetBodyResp(resp)

	ContactResp := []Contact{}
	err := json.Unmarshal(jsonData, &ContactResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(ContactResp)
}

func (b ContactClient) Add(body string) {
	resp := b.cl.Add("/contacts", body)
	jsonData := b.cl.GetBodyResp(resp)

	ContactResp := Contact{}
	err := json.Unmarshal(jsonData, &ContactResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(ContactResp)
}

func (b ContactClient) Delete(name string, surname string) {
	variables := []string{}
	variables = append(variables, name)
	variables = append(variables, surname)
	resp := b.cl.Delete("/contacts", variables)
	jsonData := b.cl.GetBodyResp(resp)

	ContactResp := []Contact{}
	err := json.Unmarshal(jsonData, &ContactResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(ContactResp)
}
