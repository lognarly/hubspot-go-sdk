package crm

import (
	"fmt"
)

type LineItems interface {
	ListAssociations(query *LineItemAssociationsQuery, lineItem string, toObjectType string) (*LineItemAssociations, error)
	Associate(lineItemId string, toObjectType string, toObjectId string, associationType string) (*LineItem, error)
	Disassociate(lineItemId string, toObjectType string, toObjectId string, associationType string) (error)
	List(query *LineItemListQuery) (*LineItemList, error)
	Create(options *LineItemCreateOrUpdateOptions) (*LineItem, error)
	Read(query *LineItemReadQuery, lineItemId string) (*LineItem, error)
	Update(lineItemId string, options *LineItemCreateOrUpdateOptions) (*LineItem, error)
	Archive(lineItemId string) (error)
	BatchArchive(lineItemIds []string) (error)
	BatchCreate(options *LineItemBatchCreateOptions) (*LineItemBatchOutput, error)
	BatchRead(options *LineItemBatchReadOptions) (*LineItemBatchOutput, error)
	BatchUpdate(options *LineItemBatchUpdateOptions) (*LineItemBatchOutput, error)
	Search(options *LineItemSearchOptions) (*LineItemSearchResults, error)
	Merge(options *LineItemMergeOptions) (*LineItem, error)
}

type lineItems struct {
	client *Client
}

type LineItemAssociationsQuery struct {
	ListAssociationsQuery
}

type LineItemAssociations struct {
	Results    []LineItemAssociation `json:"results"`
	Pagination
}

type LineItemAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type LineItemListQuery struct {
	ListQuery
}

type LineItemList struct {
	LineItems  []LineItem `json:"results"`
	Pagination
}

type LineItem struct {
	Id         string             `json:"id"`
	Properties LineItemProperties `json:"properties"`
	CreatedAt  string             `json:"createdAt"`
	UpdatedAt  string             `json:"updatedAt"`
	Archived   bool               `json:"archived"`
}

type LineItemCreateOrUpdateOptions struct {
	Properties LineItemProperties `json:"properties"`
}

type LineItemReadQuery struct {
	ReadQuery
}

type LineItemBatchOutput struct {
	Status      string     `json:"status"`
	Results     []LineItem `json:"results"`
	RequestedAt string     `json:"requestedAt"`
	StartedAt   string     `json:"startedAt"`
	CompletedAt string     `json:"completedAt"`
}

type LineItemBatchReadOptions struct {
	BatchReadOptions
}

type LineItemBatchCreateOptions struct {
	Inputs []LineItemCreateOrUpdateOptions `json:"inputs"`
}

type LineItemBatchUpdateOptions struct {
	Inputs []LineItemBatchUpdateProperties `json:"inputs"`
}

type LineItemBatchUpdateProperties struct {
	Id         string             `json:"id"`
	Properties LineItemProperties `json:"properties"`
}

type LineItemSearchOptions struct {
	SearchOptions
}

type LineItemSearchResults struct {
	Total      int64       `json:"total"`
	Results    []LineItem `json:"results"`
	Pagination
}

type LineItemMergeOptions struct {
	MergeOptions
}

func (z *lineItems) ListAssociations(query *LineItemAssociationsQuery, lineItemId string, toObjectType string) (*LineItemAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/%s/associations/%s", lineItemId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.ListAssociations(): newHttpRequest(): %v", err)
	}

	la := &LineItemAssociations{}

	err = z.client.do(req, la)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.ListAssociations(): do(): %v", err)
	}
	
	return la, nil
}

func (z *lineItems) Associate(lineItemId string, toObjectType string, toObjectId string, associationType string) (*LineItem, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/%s/associations/%s/%s/%s", lineItemId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Associate(): newHttpRequest(): %v", err)
	}

	lineItem := &LineItem{}

	err = z.client.do(req, lineItem)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Associate(): do(): %v", err)
	}
	
	return lineItem, nil
}

func (z *lineItems) Disassociate(lineItemId string, toObjectType string, toObjectId string, associationType string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/lineItems/%s/associations/%s/%s/%s", lineItemId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.lineItems.Disassociate(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (l *lineItems) List(query *LineItemListQuery) (*LineItemList, error) {
	u := fmt.Sprintf("crm/v3/objects/line_items")
	req, err := l.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.List(): newHttpRequest(): %v", err)
	}

	pl := &LineItemList{}

	err = l.client.do(req, pl)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.List(): do(): %v", err)
	}
	
	return pl, nil
}

func (l *lineItems) Create(options *LineItemCreateOrUpdateOptions) (*LineItem, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items")
	req, err := l.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Create(): newHttpRequest(): %+v", err)
	}

	lineItem := &LineItem{}

	err = l.client.do(req, lineItem)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Create(): do(): %+v", err)
	}

	return lineItem, nil
}

func (l *lineItems) Read(query *LineItemReadQuery, lineItemId string) (*LineItem, error) {
	u := fmt.Sprintf("crm/v3/objects/line_items/%s", lineItemId)
	req, err := l.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Read(): newHttpRequest(): %v", err)
	}

	lineItem := &LineItem{}

	err = l.client.do(req, lineItem)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Read(): do(): %+v", err)
	}

	return lineItem, nil
}

func (l *lineItems) Update(lineItemId string, options *LineItemCreateOrUpdateOptions) (*LineItem, error) {
	u := fmt.Sprintf("crm/v3/objects/line_items/%s", lineItemId)
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
	u := fmt.Sprintf("crm/v3/objects/line_items/%s", lineItemId)
	req, err := l.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.lineItems.Archive(): newHttpRequest(): %v", err)
	}

	return l.client.do(req, nil)
}

func (l *lineItems) BatchArchive(lineItemIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, lineItemId := range lineItemIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: lineItemId})
	}

	req, err := l.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.lineItems.BatchArchive(): newHttpRequest(): %v", err)
	}

	return l.client.do(req, nil)
}

func (l *lineItems) BatchCreate(options *LineItemBatchCreateOptions) (*LineItemBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/batch/create")
	req, err := l.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.BatchCreate(): newHttpRequest(): %v", err)
	}

	lineItems := &LineItemBatchOutput{}

	err = l.client.do(req, lineItems)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.BatchCreate(): do(): %+v", err)
	}

	return lineItems, nil
}

func (z *lineItems) BatchRead(options *LineItemBatchReadOptions) (*LineItemBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.BatchRead(): newHttpRequest(): %v", err)
	}

	lbrr := &LineItemBatchOutput{}

	err = z.client.do(req, lbrr)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.BatchRead(): do(): %+v", err)
	}

	return lbrr, nil
}

func (z *lineItems) BatchUpdate(options *LineItemBatchUpdateOptions) (*LineItemBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.BatchUpdate(): newHttpRequest(): %v", err)
	}

	lineItems := &LineItemBatchOutput{}

	err = z.client.do(req, lineItems)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.BatchUpdate(): do(): %+v", err)
	}

	return lineItems, nil
}

func (z *lineItems) Search(options *LineItemSearchOptions) (*LineItemSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Search(): newHttpRequest(): %v", err)
	}

	lineItems := &LineItemSearchResults{}

	err = z.client.do(req, lineItems)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Search(): do(): %+v", err)
	}

	return lineItems, nil
}

func (z *lineItems) Merge(options *LineItemMergeOptions) (*LineItem, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/merge")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Merge(): newHttpRequest(): %v", err)
	}

	company := &LineItem{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Merge(): do(): %+v", err)
	}

	return company, nil
}