package main

import (
	"fmt"
)

type Products interface {
	List(query *ProductListQuery) (*ProductList, error)
	Create(options *ProductCreateOptions) (*Product, error)
	Read(productId string) (*Product, error)
	Update(productId string, options *ProductUpdateOptions) (*Product, error)
	Archive(productId string) (error)
	BatchArchive(productIds []string) (error)
	BatchCreate(options *ProductBatchCreateOptions) (*ProductBatchCreateOutput, error)
}

type products struct {
	client *Client
}

type ProductQuery string

type ProductListQuery struct {
	ListQuery
}

type ProductList struct {
	Products      []Product `json:"results"`
	Pagination
}

type Product struct {
	ID         string            `json:"id"`
	Properties ProductProperties `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

type ProductProperties struct {
	CreateDate       string `json:"createdate"`
	Description      string `json:"description"`
	LastModifiedDate string `json:"hs_lastmodifieddate"`
	HSObjectID       string `json:"hs_object_id"`
	Name             string `json:"name"`
	Price            string `json:"price"`
	SKU              string `json:"hs_sku"`
}

type ProductCreateOptions struct {
	Properties ProductCreateOrUpdateProperties `json:"properties"`
}

type ProductCreateOrUpdateProperties struct {
	Description      string `json:"description"`
	Name             string `json:"name"`
	Price            string `json:"price"`
	SKU              string `json:"hs_sku"`
}

type ProductUpdateOptions struct {
	Properties ProductCreateOrUpdateProperties `json:"properties"`
}

type ProductBatchCreateOptions struct {
	Inputs []ProductCreateOptions `json:"inputs"`
}

type ProductBatchCreateOutput struct {
	Status  string    `json:"status"`
	Results []Product `json:"results"`
}

func (z *products) List(query *ProductListQuery) (*ProductList, error) {
	u := fmt.Sprintf("crm/v3/objects/products")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.product.List(): newHttpRequest(): %v", err)
	}

	pl := &ProductList{}

	err = z.client.do(req, pl)
	if err != nil {
		return nil, fmt.Errorf("client.product.List(): do(): %v", err)
	}
	
	return pl, nil
}

func (z *products) Create(options *ProductCreateOptions) (*Product, error) {
	u := fmt.Sprintf("/crm/v3/objects/products")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.product.Create(): newHttpRequest(): %+v", err)
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, fmt.Errorf("client.product.Create(): do(): %+v", err)
	}

	return product, nil
}

func (z *products) Read(productId string) (*Product, error) {
	u := fmt.Sprintf("crm/v3/objects/products/%s", productId)
	req, err := z.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.product.Read(): newHttpRequest(): %v", err)
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, fmt.Errorf("client.product.Read(): do(): %+v", err)
	}

	return product, nil
}

func (z *products) Update(productId string, options *ProductUpdateOptions) (*Product, error) {
	u := fmt.Sprintf("crm/v3/objects/products/%s", productId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.product.Update(): newHttpRequest(): %v", err)
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, fmt.Errorf("client.product.Update(): do(): %+v", err)
	}

	return product, nil
}

func (z *products) Archive(productId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/products/%s", productId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.product.Archive(): newHttpRequest(): %v", err)
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
		return fmt.Errorf("client.product.BatchArchive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *products) BatchCreate(options *ProductBatchCreateOptions) (*ProductBatchCreateOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/products/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.product.BatchCreate(): newHttpRequest(): %v", err)
	}

	products := &ProductBatchCreateOutput{}

	err = z.client.do(req, products)
	if err != nil {
		return nil, fmt.Errorf("client.product.BatchCreate(): do(): %+v", err)
	}

	return products, nil
}