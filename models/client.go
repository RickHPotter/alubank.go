package models

type Client struct {
	name       string
	taxId      string
	occupation string
}

/*
DEV INTERFACE
*/

/*
USER INTERFACE
*/

// DATA

func (client *Client) GetName() string {
	return client.name
}

func (client *Client) GetId() string {
	return client.taxId
}

func (client *Client) GetOccupation() string {
	return client.occupation
}

func (client *Client) Print() string {
	return client.GetName() + client.GetId() + client.GetOccupation()
}

func NewClient(name string, taxId string, occupation string) *Client {
	client := Client{name, taxId, occupation}
	return &client
}
