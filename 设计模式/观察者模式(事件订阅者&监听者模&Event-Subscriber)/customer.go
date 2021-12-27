package main

import "log"

type customer struct {
	id string
}

func (c *customer) update(itemName string) {
	log.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}
