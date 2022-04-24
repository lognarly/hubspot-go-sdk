package main

import {
	"fmt"
}

type Companies interface {
	List(query *CompanyListQuery) (*CompanyList, error)
	Create(options *CompanyCreateOptions) (*Company, error)
	Read(query *CompanyReadQuery, companyId string) (*Company, error)
	Update(options *CompanyCreateOrUpdateOptions, companyId string) (*Company, error)
}

type companies struct {
	client *Client
}

type CompanyPropertiesQuery string

type CompanyPropertiesWithHistoryQuery string

type CompanyListQuery struct {
	ListOptions
	Properties            CompanyPropertiesQuery            `url:"properties,omitempty"`
	PropertiesWithHistory CompanyPropertiesWithHistoryQuery `url:"propertiesWithHistory,omitempty"`
	Associations          CompanyAssociationQuery           `url:"associations,omitempty"`
	Archived              bool                              `url:"archived,omitempty"`
}

type CompanyList struct {
	Companies  []Company `json:"results"`
	Pagination

}

type Company struct {
	ID         string            `json:"id"`
	Properties CompanyProperties `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

type CompanyProperties struct {
	City             string `json:"city"`
	CreateDate       string `json:"createdate"`
	Domain           string `json:"domain"`
	HSObjectID       string `hs_object_id`
	LastModifiedDate string `json:"hs_lastmodifieddate"`
	Industry         string `json:"industry"`
	Name             string `json:"name"`
	Phone            string `json:"phone"`
	State            string `json:"state"`
}

type CompanyCreateOptions struct {
	Properties CompanyCreateOrUpdateProperties `json:"properties"`
}

type CompanyCreateOrUpdateProperties struct {
	City             string `json:"city"`
	Domain           string `json:"domain"`
	Industry         string `json:"industry"`
	Name             string `json:"name"`
	Phone            string `json:"phone"`
	State            string `json:"state"`
}

/*type CompanyReadQuery struct {
	Properties            CompanyPropertiesQuery            `url:"properties,omitempty"`
	PropertiesWithHistory CompanyPropertiesWithHistoryQuery `url:"propertiesWithHistory,omitempty"`
	Associations          CompanyAssociationQuery           `url:"associations,omitempty"`
	Archived              bool                              `url:"archived,omitempty"`
	IdProperty            string                            `url:"idProperty,omitempty"`
}

type CompanyUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}*/



func (c *companies) List(query *CompanyListQuery) (*CompanyList, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies")
	req, err := c.client.newHttpRequest("GET", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.List(): newHttpRequest(): %v", err)
	}

	cl := &CompanyList{}

	err = l.client.do(req, cl)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.List(): do(): %v", err)
	}
	
	return cl, nil
}

func (c *companies) Create(options *CompanyCreateOptions) (*Company, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies")
	req, err := c.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.Create(): newHttpRequest(): %v", err)
	}

	company := &Company{}

	err = c.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.Create(): do(): %v", err)
	}
	
	return company, nil
}

func (c *companies) Read(query *CompanyReadQuery, companyId string) (*Company, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", companyId)
	req, err := c.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.Read(): newHttpRequest(): %v", err)
	}

	company := &Company{}

	err = c.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.Read(): do(): %+v", err)
	}

	return company, nil
}

func (c *companies) Update(options *CompanyCreateOrUpdateOptions, companyId string) (*Company, error) {
	u := fmt.Sprintf("crm/v3/objects/companies/%s", companyId)
	req, err := c.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.Update(): newHttpRequest(): %v", err)
	}

	company := &Company{}

	err = c.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.Update(): do(): %+v", err)
	}

	return company, nil
}