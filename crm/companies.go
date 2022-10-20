package crm

import (
	"fmt"
)

type Companies interface {
	ListAssociations(query *CompanyAssociationsQuery, companyId string, toObjectType string) (*CompanyAssociations, error)
	Associate(companyId string, toObjectType string, toObjectId string, associationType string) (*Company, error)
	Disassociate(companyId string, toObjectType string, toObjectId string, associationType string) (error)
	List(query *CompanyListQuery) (*CompanyList, error)
	Create(options *CompanyCreateOrUpdateOptions) (*Company, error)
	Read(query *CompanyReadQuery, companyId string) (*Company, error)
	Update(options *CompanyCreateOrUpdateOptions, companyId string) (*Company, error)
	Archive(companyId string) (error)
	BatchArchive(companyIds []string) (error)
	BatchCreate(options *CompanyBatchCreateOptions) (*CompanyBatchOutput, error)
	BatchRead(options *CompanyBatchReadOptions) (*CompanyBatchOutput, error)
	BatchUpdate(options *CompanyBatchUpdateOptions) (*CompanyBatchOutput, error)
	Search(options *CompanySearchOptions) (*CompanySearchResults, error)
	Merge(options *CompanyMergeOptions) (*Company, error)
}

type companies struct {
	client *Client
}

type CompanyAssociationsQuery struct {
	ListAssociationsQuery
}

type CompanyAssociations struct {
	Results []CompanyAssociation `json:"results"`
	Pagination
}

type CompanyAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CompanyListQuery struct {
	ListQuery
}

type CompanyList struct {
	Companies  []Company `json:"results"`
	Pagination
}

type Company struct {
	Id         string            `json:"id"`
	Properties CompanyProperties `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

type CompanyCreateOrUpdateOptions struct {
	Properties CompanyProperties `json:"properties"`
}

type CompanyReadQuery struct {
	ReadQuery
}

type CompanyUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}

type CompanyBatchOutput struct {
	Status      string    `json:"status"`
	Results     []Company `json:"results"`
	RequestedAt string    `json:"requestedAt"`
	StartedAt   string    `json:"startedAt"`
	CompletedAt string    `json:"completedAt"`
}

type CompanyBatchReadOptions struct {
	BatchReadOptions
}

type CompanyBatchCreateOptions struct {
	Inputs []CompanyCreateOrUpdateOptions `json:"inputs"`
}

type CompanyBatchUpdateOptions struct {
	Inputs []CompanyBatchUpdateProperties `json:"inputs"`
}

type CompanyBatchUpdateProperties struct {
	Id         string            `json:"id"`
	Properties CompanyProperties `json:"properties"`
}

type CompanySearchOptions struct {
	SearchOptions
}

type CompanySearchResults struct {
	Total      int64     `json:"total"`
	Results    []Company `json:"results"`
	Pagination
}

type CompanyMergeOptions struct {
	MergeOptions
}

func (z *companies) ListAssociations(query *CompanyAssociationsQuery, companyId string, toObjectType string) (*CompanyAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/%s/associations/%s", companyId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.companies.ListAssociations(): newHttpRequest(): %v", err)
	}

	ca := &CompanyAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, fmt.Errorf("client.companies.ListAssociations(): do(): %v", err)
	}
	
	return ca, nil
}

func (z *companies) Associate(companyId string, toObjectType string, toObjectId string, associationType string) (*Company, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/%s/associations/%s/%s/%s", companyId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Associate(): newHttpRequest(): %v", err)
	}

	company := &Company{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Associate(): do(): %v", err)
	}
	
	return company, nil
}

func (z *companies) Disassociate(companyId string, toObjectType string, toObjectId string, associationType string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/%s/associations/%s/%s/%s", companyId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.companies.Disassociate(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *companies) List(query *CompanyListQuery) (*CompanyList, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.companies.List(): newHttpRequest(): %v", err)
	}

	cl := &CompanyList{}

	err = z.client.do(req, cl)
	if err != nil {
		return nil, fmt.Errorf("client.companies.List(): do(): %v", err)
	}
	
	return cl, nil
}

func (z *companies) Create(options *CompanyCreateOrUpdateOptions) (*Company, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Create(): newHttpRequest(): %v", err)
	}

	company := &Company{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Create(): do(): %v", err)
	}
	
	return company, nil
}

func (z *companies) Read(query *CompanyReadQuery, companyId string) (*Company, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", companyId)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Read(): newHttpRequest(): %v", err)
	}

	company := &Company{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Read(): do(): %+v", err)
	}

	return company, nil
}

func (z *companies) Update(options *CompanyCreateOrUpdateOptions, companyId string) (*Company, error) {
	u := fmt.Sprintf("crm/v3/objects/companies/%s", companyId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Update(): newHttpRequest(): %v", err)
	}

	company := &Company{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Update(): do(): %+v", err)
	}

	return company, nil
}

func (z *companies) Archive(companyId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/companies/%s", companyId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.companies.Archive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *companies) BatchArchive(companyIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, companyId := range companyIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: companyId})
	}

	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.companies.BatchArchive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *companies) BatchCreate(options *CompanyBatchCreateOptions) (*CompanyBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.companies.BatchCreate(): newHttpRequest(): %v", err)
	}

	companies := &CompanyBatchOutput{}

	err = z.client.do(req, companies)
	if err != nil {
		return nil, fmt.Errorf("client.companies.BatchCreate(): do(): %+v", err)
	}

	return companies, nil
}

func (z *companies) BatchRead(options *CompanyBatchReadOptions) (*CompanyBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.companies.BatchUpdate(): newHttpRequest(): %v", err)
	}

	companies := &CompanyBatchOutput{}

	err = z.client.do(req, companies)
	if err != nil {
		return nil, fmt.Errorf("client.companies.BatchUpdate(): do(): %+v", err)
	}

	return companies, nil
}

func (z *companies) BatchUpdate(options *CompanyBatchUpdateOptions) (*CompanyBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.companies.BatchUpdate(): newHttpRequest(): %v", err)
	}

	companies := &CompanyBatchOutput{}

	err = z.client.do(req, companies)
	if err != nil {
		return nil, fmt.Errorf("client.companies.BatchUpdate(): do(): %+v", err)
	}

	return companies, nil
}

func (z *companies) Search(options *CompanySearchOptions) (*CompanySearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Search(): newHttpRequest(): %v", err)
	}

	companies := &CompanySearchResults{}

	err = z.client.do(req, companies)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Search(): do(): %+v", err)
	}

	return companies, nil
}

func (z *companies) Merge(options *CompanyMergeOptions) (*Company, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/merge")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Merge(): newHttpRequest(): %v", err)
	}

	company := &Company{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.companies.Merge(): do(): %+v", err)
	}

	return company, nil
}