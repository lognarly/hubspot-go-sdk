package main

import (
	"fmt"
)

type LineItems interface {
	List(options *LineItemListOptions) (*LineItemList, error)
	Create(options *LineItemCreateOptions) (*LineItem, error)
	Read(lineItemId string) (*LineItem, error)
	Update(lineItemId string, options *LineItemUpdateOptions) (*LineItem, error)
	Archive(lineItemId string) (error)
	BatchArchive(options *LineItemBatchArchiveOptions) (error)
	BatchCreate(options *LineItemBatchCreateOptions) (*LineItemBatchCreateOutput, error)
}

type lineItems struct {
	client *Client
}

type LineItemQuery string

type LineItemListOptions struct {
	ListOptions
	Properties  LineItemQuery `url:"properties,omitempty"`
}

type LineItemList struct {
	LineItems      []LineItem `json:"results"`
	Pagination
}

type LineItem struct {
	ID string `json:"id"`
	Properties LineItemProperties `json:"properties"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
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

type LineItemBatchArchiveOptions struct {
	Inputs []ArchiveLineItems `json:"inputs"`
}

type ArchiveLineItems struct {
	LineItemId string `json:"id"`
}

type LineItemBatchCreateOptions struct {
	Inputs []LineItemCreateOptions `json:"inputs"`
}

type LineItemBatchCreateOutput struct {
	Status  string    `json:"status"`
	Results []LineItem `json:"results"`
}

func (l *lineItems) List(options *LineItemListOptions) (*LineItemList, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems")
	req, err := l.client.newHttpRequest("GET", u, options)
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

func (l *lineItems) Delete(lineItemId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", lineItemId)
	req, err := l.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.lineItem.Delete(): newHttpRequest(): %v", err)
	}

	return l.client.do(req, nil)
}

func (l *lineItems) BatchArchive(options *LineItemBatchArchiveOptions) (error) {
	u := fmt.Sprintf("/crm/v3/objects/lineitems/batch/archive")
	req, err := l.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.lineItem.BatchArchive(): newHttpRequest(): %v", err)
	}

	return l.client.do(req, nil)
}

func (l *lineItems) BatchCreate(options *LineItemBatchCreateOptions) (*LineItemBatchCreateOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/lineitems/batch/create")
	req, err := l.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.BatchCreate(): newHttpRequest(): %v", err)
	}

	lineItems := &LineItemBatchCreateOutput{}

	err = l.client.do(req, lineItems)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.BatchCreate(): do(): %+v", err)
	}

	return lineItems, nil
}