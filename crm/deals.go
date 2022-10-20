package crm

import (
	"fmt"
)

type Deals interface {
	ListAssociations(query *DealAssociationsQuery, dealId string, toObjectType string) (*DealAssociations, error)
	Associate(dealId string, toObjectType string, toObjectId string, associationType string) (*Deal, error)
	Disassociate(dealId string, toObjectType string, toObjectId string, associationType string) (error)
	List(query *DealListQuery) (*DealList, error)
	Create(options *DealCreateOrUpdateOptions) (*Deal, error)
	Read(query *DealReadQuery, dealId string) (*Deal, error)
	Update(dealId string, options *DealCreateOrUpdateOptions) (*Deal, error)
	Archive(dealId string) (error)
	BatchArchive(dealIds []string) (error)
	BatchCreate(options *DealBatchCreateOptions) (*DealBatchOutput, error)
	BatchRead(options *DealBatchReadOptions) (*DealBatchOutput, error)
	BatchUpdate(options *DealBatchUpdateOptions) (*DealBatchOutput, error)
	Search(options *DealSearchOptions) (*DealSearchResults, error)
	Merge(options *DealMergeOptions) (*Deal, error)
}

type deals struct {
	client *Client
}

type DealAssociationsQuery struct {
	ListAssociationsQuery
}

type DealAssociations struct {
	Results    []DealAssociation `json:"results"`
	Pagination
}

type DealAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type DealListQuery struct {
	ListQuery
}

type DealList struct {
	Deals      []Deal `json:"results"`
	Pagination
}

type Deal struct {
	Id         string         `json:"id"`
	Properties DealProperties `json:"properties"` //This can be found in deal_properties.go as to not clutter this file
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type DealCreateOrUpdateOptions struct {
	Properties DealProperties `json:"properties"`
}

type DealReadQuery struct {
	ReadQuery
}

type DealBatchOutput struct {
	Status      string `json:"status"`
	Results     []Deal `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type DealBatchReadOptions struct {
	BatchReadOptions
}

type DealBatchCreateOptions struct {
	Inputs []DealCreateOrUpdateOptions `json:"inputs"`
}

type DealBatchUpdateOptions struct {
	Inputs []DealBatchUpdateProperties `json:"inputs"`
}

type DealBatchUpdateProperties struct {
	Id         string         `json:"id"`
	Properties DealProperties `json:"properties"` //This can be found in deal_properties.go as to not clutter this file
}

type DealSearchOptions struct {
	SearchOptions
}

type DealSearchResults struct {
	Total      int64 `json:"total"`
	Results    []Deal `json:"results"`
	Pagination
}

type DealMergeOptions struct {
	MergeOptions
}

func (z *deals) ListAssociations(query *DealAssociationsQuery, dealId string, toObjectType string) (*DealAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/%s/associations/%s", dealId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.deals.ListAssociations(): newHttpRequest(): %v", err)
	}

	da := &DealAssociations{}

	err = z.client.do(req, da)
	if err != nil {
		return nil, fmt.Errorf("client.deals.ListAssociations(): do(): %v", err)
	}
	
	return da, nil
}

func (z *deals) Associate(dealId string, toObjectType string, toObjectId string, associationType string) (*Deal, error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/%s/associations/%s/%s/%s", dealId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Associate(): newHttpRequest(): %v", err)
	}

	deal := &Deal{}

	err = z.client.do(req, deal)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Associate(): do(): %v", err)
	}
	
	return deal, nil
}

func (z *deals) Disassociate(dealId string, toObjectType string, toObjectId string, associationType string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/%s/associations/%s/%s/%s", dealId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.deals.Disassociate(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *deals) List(query *DealListQuery) (*DealList, error) {
	u := fmt.Sprintf("crm/v3/objects/deals")

	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.deals.List(): newHttpRequest(): %v", err)
	}

	dl := &DealList{}
	
	err = z.client.do(req, dl)
	if err != nil {
		return nil, fmt.Errorf("client.deals.List(): do()(): %v", err)
	}

	return dl, nil
}

func (z *deals) Create(options *DealCreateOrUpdateOptions) (*Deal, error) {
	u := fmt.Sprintf("crm/v3/objects/deals")
	
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Create(): newHttpRequest(): %v", err)
	}

	deal := &Deal{}
	
	err = z.client.do(req, deal)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Create(): do()(): %v", err)
	}

	return deal, nil
}

func (z *deals) Read(query *DealReadQuery, dealId string) (*Deal, error) {
	u := fmt.Sprintf("crm/v3/objects/deals/%s", dealId)
	
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Read(): newHttpRequest(): %v", err)
	}

	deal := &Deal{}

	
	err = z.client.do(req, deal)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Read(): do()(): %v", err)
	}

	return deal, nil
}

func (z *deals) Update(dealId string, options *DealCreateOrUpdateOptions) (*Deal, error) {
	u := fmt.Sprintf("crm/v3/objects/deals/%s", dealId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.dealss.Update(): newHttpRequest(): %v", err)
	}

	deal := &Deal{}

	err = z.client.do(req, deal)
	if err != nil {
		return nil, fmt.Errorf("client.dealss.Update(): do(): %+v", err)
	}

	return deal, nil
}

func (z *deals) Archive(dealId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/deals/%s", dealId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.dealss.Archive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *deals) BatchArchive(dealIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, dealId := range dealIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: dealId})
	}

	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.contact.BatchArchive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *deals) BatchCreate(options *DealBatchCreateOptions) (*DealBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.dealss.BatchCreate(): newHttpRequest(): %v", err)
	}

	deals := &DealBatchOutput{}

	err = z.client.do(req, deals)
	if err != nil {
		return nil, fmt.Errorf("client.dealss.BatchCreate(): do(): %+v", err)
	}

	return deals, nil
}

func (z *deals) BatchRead(options *DealBatchReadOptions) (*DealBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.deals.BatchUpdate(): newHttpRequest(): %v", err)
	}

	deals := &DealBatchOutput{}

	err = z.client.do(req, deals)
	if err != nil {
		return nil, fmt.Errorf("client.deals.BatchUpdate(): do(): %+v", err)
	}

	return deals, nil
}

func (z *deals) BatchUpdate(options *DealBatchUpdateOptions) (*DealBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.deals.BatchUpdate(): newHttpRequest(): %v", err)
	}

	deals := &DealBatchOutput{}

	err = z.client.do(req, deals)
	if err != nil {
		return nil, fmt.Errorf("client.deals.BatchUpdate(): do(): %+v", err)
	}

	return deals, nil
}

func (z *deals) Search(options *DealSearchOptions) (*DealSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Search(): newHttpRequest(): %v", err)
	}

	deals := &DealSearchResults{}

	err = z.client.do(req, deals)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Search(): do(): %+v", err)
	}

	return deals, nil
}

func (z *deals) Merge(options *DealMergeOptions) (*Deal, error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/merge")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Merge(): newHttpRequest(): %v", err)
	}

	deal := &Deal{}

	err = z.client.do(req, deal)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Merge(): do(): %+v", err)
	}

	return deal, nil
}