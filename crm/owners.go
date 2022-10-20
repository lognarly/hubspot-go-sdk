package crm

import (
	"fmt"
)

type Owners interface {
	List(query *OwnerListQuery) (*OwnerList, error)
	Read(ownerId string, query *OwnerReadQuery) (*Owner, error)
}

type owners struct {
	client *Client
}

type OwnerListQuery struct {
	Email    string `url:"email,omitempty"`
	After    string `ur:"after,omitempty"`
	Limit    string `url:"limit,omitempty"`
	Archived bool   `url:"archived,omitempty"`
}

type OwnerList struct {
	Results []Owner `json:"results"`
	Pagination
}

type Owner struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	UserId    int64  `json:"userId,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Archived  bool   `json:"archived"`
	Teams     []Team `json:"teams,omitempty"`
}

type Team struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Primary bool   `json:"primary"`
}

type OwnerReadQuery struct {
	IdProperty string `url:"idProperty,omitempty"`
	Archived   bool   `url:"archived,omitempty"`
}

func (z *owners) List(query *OwnerListQuery) (*OwnerList, error) {
	u := fmt.Sprintf("/crm/v3/owners/")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.owners.List(): newHttpRequest(): %v", err)
	}

	ol := &OwnerList{}

	err = z.client.do(req, ol)
	if err != nil {
		return nil, fmt.Errorf("client.owners.List(): do(): %v", err)
	}
	
	return ol, nil
}

func (z *owners) Read(ownerId string, query *OwnerReadQuery) (*Owner, error) {
	u := fmt.Sprintf("/crm/v3/owners/%s", ownerId)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.owners.Read(): newHttpRequest(): %v", err)
	}

	owner := &Owner{}

	err = z.client.do(req, owner)
	if err != nil {
		return nil, fmt.Errorf("client.owners.Read(): do(): %v", err)
	}
	
	return owner, nil
}