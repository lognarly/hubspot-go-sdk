package main

import (
	"fmt"
)

type LineItems interface {
	List(query *LineItemListQuery) (*LineItemList, error)
	Create(options *LineItemCreateOptions) (*LineItem, error)
	Read(lineItemId string) (*LineItem, error)
	Update(lineItemId string, options *LineItemUpdateOptions) (*LineItem, error)
	Archive(lineItemId string) (error)
	BatchArchive(lineItemIds []string) (error)
	BatchCreate(options *LineItemBatchCreateOptions) (*LineItemBatchCreateResults, error)
}

type lineItems struct {
	client *Client
}

type LineItemQuery string

type LineItemListQuery struct {
	ListQuery
}

type LineItemList struct {
	LineItems      []LineItem `json:"results"`
	Pagination
}

type LineItem struct {
	ID         string             `json:"id"`
	Properties LineItemProperties `json:"properties"`
	CreatedAt  string             `json:"createdAt"`
	UpdatedAt  string             `json:"updatedAt"`
	Archived   bool               `json:"archived"`
}

type LineItemProperties struct {
	CreateDate       string `json:"createdate"`
	Description      string `json:"description"`
	LastModifiedDate string `json:"hs_lastmodifieddate"`
	HSObjectID       string `json:"hs_object_id"`
	Name             string `json:"name"`
	Price            string `json:"price"`
	SKU              string `json:"hs_sku"`
}

type LineItemCreateOptions struct {
	Properties LineItemCreateOrUpdateProperties `json:"properties"`
}

type LineItemCreateOrUpdateProperties struct {
	Description      string `json:"description"`
	Name             string `json:"name"`
	Price            string `json:"price"`
	SKU              string `json:"hs_sku"`
}

type LineItemUpdateOptions struct {
	Properties LineItemCreateOrUpdateProperties `json:"properties"`
}

type LineItemBatchCreateOptions struct {
	Inputs []LineItemCreateOptions `json:"inputs"`
}

type LineItemBatchCreateResults struct {
	Status  string    `json:"status"`
	Results []LineItem `json:"results"`
}

type LineItemBatchReadOptions struct {
	BatchReadOptions
}

type LineItemBatchReadResults struct {
	Status  string     `json:"status"`
	Results []LineItem `json:"results"`
	RequestedAt string  `json:"requestedAt"`
	StartedAt   string  `json:"startedAt"`
	CompletedAt string  `json:"completedAt"`
}

func (l *lineItems) List(query *LineItemListQuery) (*LineItemList, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems")
	req, err := l.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.List(): newHttpRequest(): %v", err)
	}

	pl := &LineItemList{}

	err = l.client.do(req, pl)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.List(): do(): %v", err)
	}
	
	return pl, nil
}

func (l *lineItems) Create(options *LineItemCreateOptions) (*LineItem, error) {
	u := fmt.Sprintf("/crm/v3/objects/lineitems")
	req, err := l.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.Create(): newHttpRequest(): %+v", err)
	}

	lineItem := &LineItem{}

	err = l.client.do(req, lineItem)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.Create(): do(): %+v", err)
	}

	return lineItem, nil
}

func (l *lineItems) Read(lineItemId string) (*LineItem, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", lineItemId)
	req, err := l.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.Read(): newHttpRequest(): %v", err)
	}

	lineItem := &LineItem{}

	err = l.client.do(req, lineItem)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.Read(): do(): %+v", err)
	}

	return lineItem, nil
}

func (l *lineItems) Update(lineItemId string, options *LineItemUpdateOptions) (*LineItem, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", lineItemId)
	req, err := l.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.Update(): newHttpRequest(): %v", err)
	}

	lineItem := &LineItem{}

	err = l.client.do(req, lineItem)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.Update(): do(): %+v", err)
	}

	return lineItem, nil
}

func (l *lineItems) Archive(lineItemId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", lineItemId)
	req, err := l.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.lineItem.Delete(): newHttpRequest(): %v", err)
	}

	return l.client.do(req, nil)
}

func (l *lineItems) BatchArchive(lineItemIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/lineitems/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, lineItemId := range lineItemIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: lineItemId})
	}

	req, err := l.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.lineItem.BatchArchive(): newHttpRequest(): %v", err)
	}

	return l.client.do(req, nil)
}

func (l *lineItems) BatchCreate(options *LineItemBatchCreateOptions) (*LineItemBatchCreateResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/lineitems/batch/create")
	req, err := l.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.BatchCreate(): newHttpRequest(): %v", err)
	}

	lineItems := &LineItemBatchCreateResults{}

	err = l.client.do(req, lineItems)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.BatchCreate(): do(): %+v", err)
	}

	return lineItems, nil
}

func (z *lineItems) BatchRead(options *LineItemBatchReadOptions) (*LineItemBatchReadResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.BatchRead(): newHttpRequest(): %v", err)
	}

	lbrr := &LineItemBatchReadResults{}

	err = z.client.do(req, lbrr)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.BatchRead(): do(): %+v", err)
	}

	return lbrr, nil
}