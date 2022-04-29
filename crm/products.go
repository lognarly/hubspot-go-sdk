package crm

import (
	"fmt"
)

type Products interface {
	ListAssociations(query *ProductAssociationsQuery, productId string, toObjectType string) (*ProductAssociations, error)
	Associate(productId string, toObjectType string, toObjectId string, associationType string) (*Product, error)
	Disassociate(productId string, toObjectType string, toObjectId string, associationType string) (error)
	List(query *ProductListQuery) (*ProductList, error)
	Create(options *ProductCreateOrUpdateOptions) (*Product, error)
	Read(query *ProductReadQuery, productId string) (*Product, error)
	Update(productId string, options *ProductCreateOrUpdateOptions) (*Product, error)
	Archive(productId string) (error)
	BatchArchive(productIds []string) (error)
	BatchCreate(options *ProductBatchCreateOptions) (*ProductBatchOutput, error)
	BatchRead(options *ProductBatchReadOptions) (*ProductBatchOutput, error)
	BatchUpdate(options *ProductBatchUpdateOptions) (*ProductBatchOutput, error)
	Search(options *ProductSearchOptions) (*ProductSearchResults, error)
	Merge(options *ProductMergeOptions) (*Product, error)
}

type products struct {
	client *Client
}

type ProductAssociationsQuery struct {
	ListAssociationsQuery
}

type ProductAssociations struct {
	Results []ProductAssociation `json:"results"`
	Pagination
}

type ProductAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type ProductListQuery struct {
	ListQuery
}

type ProductList struct {
	Products      []Product `json:"results"`
	Pagination
}

type Product struct {
	Id         string            `json:"id"`
	Properties ProductProperties `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

type ProductCreateOrUpdateOptions struct {
	Properties ProductProperties `json:"properties"`
}

type ProductReadQuery struct {
	ReadQuery
}

type ProductBatchOutput struct {
	Status      string    `json:"status"`
	Results     []Product `json:"results"`
	RequestedAt string    `json:"requestedAt"`
	StartedAt   string    `json:"startedAt"`
	CompletedAt string    `json:"completedAt"`
}

type ProductBatchReadOptions struct {
	BatchReadOptions
}

type ProductBatchCreateOptions struct {
	Inputs []ProductCreateOrUpdateOptions `json:"inputs"`
}

type ProductBatchUpdateOptions struct {
	Inputs []ProductBatchUpdateProperties `json:"inputs"`
}

type ProductBatchUpdateProperties struct {
	Id         string            `json:"id"`
	Properties ProductProperties `json:"properties"`
}

type ProductSearchOptions struct {
	SearchOptions
}

type ProductSearchResults struct {
	Total      int64     `json:"total"`
	Results    []Product `json:"results"`
	Pagination
}

type ProductMergeOptions struct {
	MergeOptions
}

func (z *products) ListAssociations(query *ProductAssociationsQuery, productId string, toObjectType string) (*ProductAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/products/%s/associations/%s", productId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.products.ListAssociations(): newHttpRequest(): %v", err)
	}

	pa := &ProductAssociations{}

	err = z.client.do(req, pa)
	if err != nil {
		return nil, fmt.Errorf("client.products.ListAssociations(): do(): %v", err)
	}
	
	return pa, nil
}

func (z *products) Associate(productId string, toObjectType string, toObjectId string, associationType string) (*Product, error) {
	u := fmt.Sprintf("/crm/v3/objects/products/%s/associations/%s/%s/%s", productId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.products.Associate(): newHttpRequest(): %v", err)
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, fmt.Errorf("client.products.Associate(): do(): %v", err)
	}
	
	return product, nil
}

func (z *products) Disassociate(productId string, toObjectType string, toObjectId string, associationType string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/products/%s/associations/%s/%s/%s", productId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.products.Disassociate(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *products) List(query *ProductListQuery) (*ProductList, error) {
	u := fmt.Sprintf("crm/v3/objects/products")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.products.List(): newHttpRequest(): %v", err)
	}

	pl := &ProductList{}

	err = z.client.do(req, pl)
	if err != nil {
		return nil, fmt.Errorf("client.products.List(): do(): %v", err)
	}
	
	return pl, nil
}

func (z *products) Create(options *ProductCreateOrUpdateOptions) (*Product, error) {
	u := fmt.Sprintf("/crm/v3/objects/products")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.products.Create(): newHttpRequest(): %+v", err)
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, fmt.Errorf("client.products.Create(): do(): %+v", err)
	}

	return product, nil
}

func (z *products) Read(query *ProductReadQuery, productId string) (*Product, error) {
	u := fmt.Sprintf("crm/v3/objects/products/%s", productId)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.products.Read(): newHttpRequest(): %v", err)
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, fmt.Errorf("client.products.Read(): do(): %+v", err)
	}

	return product, nil
}

func (z *products) Update(productId string, options *ProductCreateOrUpdateOptions) (*Product, error) {
	u := fmt.Sprintf("crm/v3/objects/products/%s", productId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.products.Update(): newHttpRequest(): %v", err)
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, fmt.Errorf("client.products.Update(): do(): %+v", err)
	}

	return product, nil
}

func (z *products) Archive(productId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/products/%s", productId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.products.Archive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *products) BatchArchive(productIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/products/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, productId := range productIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: productId})
	}

	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.products.BatchArchive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *products) BatchCreate(options *ProductBatchCreateOptions) (*ProductBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/products/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.products.BatchCreate(): newHttpRequest(): %v", err)
	}

	products := &ProductBatchOutput{}

	err = z.client.do(req, products)
	if err != nil {
		return nil, fmt.Errorf("client.products.BatchCreate(): do(): %+v", err)
	}

	return products, nil
}

func (z *products) BatchRead(options *ProductBatchReadOptions) (*ProductBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/products/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.products.BatchUpdate(): newHttpRequest(): %v", err)
	}

	products := &ProductBatchOutput{}

	err = z.client.do(req, products)
	if err != nil {
		return nil, fmt.Errorf("client.products.BatchUpdate(): do(): %+v", err)
	}

	return products, nil
}

func (z *products) BatchUpdate(options *ProductBatchUpdateOptions) (*ProductBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/products/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.products.BatchUpdate(): newHttpRequest(): %v", err)
	}

	products := &ProductBatchOutput{}

	err = z.client.do(req, products)
	if err != nil {
		return nil, fmt.Errorf("client.products.BatchUpdate(): do(): %+v", err)
	}

	return products, nil
}

func (z *products) Search(options *ProductSearchOptions) (*ProductSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/products/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.products.Search(): newHttpRequest(): %v", err)
	}

	products := &ProductSearchResults{}

	err = z.client.do(req, products)
	if err != nil {
		return nil, fmt.Errorf("client.products.Search(): do(): %+v", err)
	}

	return products, nil
}

func (z *products) Merge(options *ProductMergeOptions) (*Product, error) {
	u := fmt.Sprintf("/crm/v3/objects/products/merge")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.products.Merge(): newHttpRequest(): %v", err)
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, fmt.Errorf("client.products.Merge(): do(): %+v", err)
	}

	return product, nil
}