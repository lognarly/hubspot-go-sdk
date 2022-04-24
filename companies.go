package main

import {
	"fmt"
}

type Companies interface {
	List(query *CompanyListQuery) (*CompanyList, error)
	Create(options *CompanyCreateOptions) (*Company, error)
	Read(query *CompanyReadQuery, companyId string) (*Company, error)
	Update(options *CompanyCreateOrUpdateOptions, companyId string) (*Company, error)
	Delete(companyId string) (error)
	BatchArchive(options *ComapnyBatchArchiveOptions) (error)
	BatchCreate(options *CompanyBatchCreateOptions) (*CompanyBatchCreateOrUpdateResults, error)
	BatchUpdate(options *CompanyBatchUpdateOptions) (*CompanyBatchCreateOrUpdateResults, error)
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

type CompanyBatchArchiveOptions struct {
	Inputs []ArchiveCompany `json:"inputs"`
}

type ArchiveCompany struct {
	CompanyId string `json:"id"`
}

type LineItemBatchCreateResults struct {
	Status  string     `json:"status"`
	Results []Comany `json:"results"`
}

type CompanyBatchUpdateOptions struct {
	Inputs []CompanyBatchUpdateProperties `json:"inputs"`
}

type CompanyBatchUpdateProperties struct {
	ID string `json:"id"`
	Properties CompanyCreateOrUpdateProperties `json:"properties"`
}

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

func (c *companies) Delete(companyId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/companies/%s", companyId)
	req, err := c.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.Delete(): newHttpRequest(): %v", err)
	}

	return c.client.do(req, nil)
}

func (c *companies) BatchArchive(options *ComapnyBatchArchiveOptions) (error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/batch/archive")
	req, err := c.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.Comanies.BatchArchive(): newHttpRequest(): %v", err)
	}

	return c.client.do(req, nil)
}

func (c *companies) BatchCreate(options *CompanyBatchCreateOptions) (*CompanyBatchCreateOrUpdateResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/batch/create")
	req, err := c.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.BatchCreate(): newHttpRequest(): %v", err)
	}

	companies := &CompanyBatchCreateOrUpdateResults{}

	err = c.client.do(req, companies)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.BatchCreate(): do(): %+v", err)
	}

	return companies, nil
}

func (c *companies) BatchUpdate(options *CompanyBatchUpdateOptions) (*CompanyBatchCreateOrUpdateResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/batch/update")
	req, err := c.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.BatchUpdate(): newHttpRequest(): %v", err)
	}

	companies := &CompanyBatchCreateOrUpdateResults{}

	err = c.client.do(req, companies)
	if err != nil {
		return nil, fmt.Errorf("client.Companies.BatchUpdate(): do(): %+v", err)
	}

	return companies, nil
}