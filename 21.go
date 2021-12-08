package main

import "fmt"

type computer interface {
	insert()
}

type mac struct {}
type windows struct{}
type client struct{}

func (c *client) insertConnector(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insert()
}

func (m *mac) insert() {
    fmt.Println("Lightning connector is plugged into mac machine.")
}

func (w *windows) insert() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	fmt.Println("USB connector is plugged into windows machine.")
}

func main() {
	client := &client{}
    mac := &mac{}

	client.insertConnector(mac)

	windows := &windows{}
	
	client.insertConnector(windows)
}