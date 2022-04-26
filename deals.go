package main

import (
	"fmt"
)

type Deals interface {
	List(query *DealListQuery) (*DealList, error)
	Create(options *DealCreateOptions) (*Deal, error)
	Read(dealId string) (*Deal, error)
	Update(dealId string, options *DealUpdateOptions) (*Deal, error)
	Archive(dealId string) (error)
	BatchArchive(dealIds []string) (error)
	BatchCreate(options *DealBatchCreateOptions) (*DealBatchCreateOutput, error)
}

type deals struct {
	client *Client
}

type DealListQuery struct {
	ListQuery
}

type DealList struct {
	Deals      []Deal `json:"results"`
	Pagination
}

type Deal struct {
	ID         string         `json:"id"`
	Properties DealProperties `json:"properties"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type DealProperties struct {
	Amount           string `json:"amount"`
	Closedate        string `json:"closedate"`
	CreateDate       string `json:"createdate"`
	DealName         string `json:"dealname"`
	DealStage        string `json:"dealstage"`
	LastModifiedDate string `json:"hs_lastmodifieddate"`
	HubSpotOwnerId   string `json:"hubspot_owner_id"`
	Pipeline         string `json:"pipeline"`
}

type DealCreateOptions struct {
	Properties DealCreateOrUpdateProperties `json:"properties"`
}

type DealCreateOrUpdateProperties struct {
	Amount           string `json:"amount"`
	Closedate        string `json:"closedate"`
	DealName         string `json:"dealname"`
	DealStage        string `json:"dealstage"`
	HubSpotOwnerId   string `json:"hubspot_owner_id"`
	Pipeline         string `json:"pipeline"`
}

type DealUpdateOptions struct {
	Properties DealCreateOrUpdateProperties `json:"properties"`
}

type DealBatchCreateOptions struct {
	Inputs []DealCreateOptions `json:"inputs"`
}

type DealBatchCreateOutput struct {
	Status  string `json:"status"`
	Results []Deal `json:"results"`
}

func (z *deals) List(query *DealListQuery) (*DealList, error) {
	u := fmt.Sprintf("crm/v3/objects/deals")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.deal.List(): newHttpRequest(): %v", err)
	}

	dl := &DealList{}

	err = z.client.do(req, dl)
	if err != nil {
		return nil, fmt.Errorf("client.deal.List(): do()(): %v", err)
	}

	return dl, nil
}

func (z *deals) Create(options *DealCreateOptions) (*Deal, error) {
	u := fmt.Sprintf("crm/v3/objects/deals")
	
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.deal.Create(): newHttpRequest(): %v", err)
	}

	deal := &Deal{}

	
	err = z.client.do(req, deal)
	if err != nil {
		return nil, fmt.Errorf("client.deal.Create(): do()(): %v", err)
	}

	return deal, nil
}

func (z *deals) Read(dealId string) (*Deal, error) {
	u := fmt.Sprintf("crm/v3/objects/deals/%s", dealId)
	
	req, err := z.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.deal.Read(): newHttpRequest(): %v", err)
	}

	deal := &Deal{}

	
	err = z.client.do(req, deal)
	if err != nil {
		return nil, fmt.Errorf("client.deal.Read(): do()(): %v", err)
	}

	return deal, nil
}

func (z *deals) Update(dealId string, options *DealUpdateOptions) (*Deal, error) {
	u := fmt.Sprintf("crm/v3/objects/deals/%s", dealId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Update(): newHttpRequest(): %v", err)
	}

	deal := &Deal{}

	err = z.client.do(req, deal)
	if err != nil {
		return nil, fmt.Errorf("client.deals.Update(): do(): %+v", err)
	}

	return deal, nil
}

func (z *deals) Archive(dealId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/deals/%s", dealId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.deals.Archive(): newHttpRequest(): %v", err)
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

func (z *deals) BatchCreate(options *DealBatchCreateOptions) (*DealBatchCreateOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.deals.BatchCreate(): newHttpRequest(): %v", err)
	}

	deals := &DealBatchCreateOutput{}

	err = z.client.do(req, deals)
	if err != nil {
		return nil, fmt.Errorf("client.deals.BatchCreate(): do(): %+v", err)
	}

	return deals, nil
}