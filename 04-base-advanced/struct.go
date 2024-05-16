package main

import "fmt"

type Client struct {
	ID   int64
	Name string
	Age  int
	IMG  *Avatar // при указателе: main.Client{ID:0, Name:"", Age:0, IMG:(*main.Avatar)(nil)}
}
type Avatar struct {
	URL  string
	Size int64
}

func updateAvatar(client *Client) {
	client.IMG.URL = "https://www.google.com"
}

func updateClient(client Client) {
	client.Name = "Артем"
}

func main() {
	i := new(int)
	_ = i

	client := Client{}
	//client.IMG = new(Avatar) //main.Client{ID:0, Name:"", Age:0, IMG:(*main.Avatar)(0x1400000c018)}
	//или можно так:
	client.IMG = &Avatar{} //main.Client{ID:0, Name:"", Age:0, IMG:(*main.Avatar)(0x140000aa000)}

	//fmt.Printf("%#v\n", client)
	//main.Client{ID:0, Name:"", Age:0, IMG:main.Avatar{URL:"", Size:0}}

	updateAvatar(&client) //добавил указатель
	fmt.Printf("%#v\n", client)
	//main.Client{ID:0, Name:"", Age:0, IMG:main.Avatar{URL:"https://www.google.com", Size:0}}
}
