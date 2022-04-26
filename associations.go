package main

import (
	"fmt"
	"strconv"
)

type Associations interface {
	List(options *AssociationListOptions, query *AssociationListQuery) (*AssociationList, error)
	Create(options *[]AssociationCreateOptions, query *AssociationCreateOrDeleteQuery) (*AssociationCreateOutput, error)
	Delete(query *AssociationCreateOrDeleteQuery) (error)
	ReadDefinition(query *AssociationDefinitionQuery) (*AssociationDefinitionOutput, error)
	CreateDefinition(options *AssociationCreateDefinitionOptions, query *AssociationDefinitionQuery) (*AssociationDefinitionOutput, error)
	UpdateDefinition(options *AssociationUpdateDefinitionOptions, query *AssociationDefinitionQuery) (error)
}

type associations struct {
	client *Client
}

type AssociationListOptions struct {
	ListOptions
}

type AssociationListQuery struct {
	FromObjectType  string
	FromObjectId    int64
	ToObjectType    string
}

type AssociationList struct {
	Assocations []AssociationListResult `json:"results"`
	Pagination
}

type AssociationListResult struct {
	ToObjectId       int               `json:"toObjectId"`
	AssociationTypes []AssociationType `json:"associationTypes"`
}

type HubspotAssociationCategory string

const (
	HubSpotDefined    HubspotAssociationCategory = "HUBSPOT_DEFINED"
	UserDefined       HubspotAssociationCategory = "USER_DEFINED"
	IntegratorDefined HubspotAssociationCategory = "INTEGRATOR_DEFINED"
)

type AssociationType struct {
	Category HubspotAssociationCategory `json:"category"`
	TypeId   int64                      `json:"typeId"`
	Label    string                     `json:"label"`
}

type AssociationCreateOrDeleteQuery struct {
	FromObjectType  string
	FromObjectId    int64
	ToObjectType    string
	ToObjectId      int64
}

type AssociationCreateOptions struct {
	Category HubspotAssociationCategory `json:"associationCategory"`
	TypeId   int64                      `json:"associationTypeId"`
}

type AssociationCreateOutput struct {
	FromObjectTypeId string   `json:"fromObjectTypeId"`
	FromObjectId     int64    `json:"fromObjectId"`
	ToObjectTypeId   string   `json:"toObjectTypeId"`
	ToObjectId       int64    `json:"toObjectId"`
	Labels           []string `json:"labels"`
}

type AssociationDefinitionQuery struct {
	FromObjectType string
	ToObjectType   string
}

type AssociationDefinitionOutput struct {
	Results []AssociationDefinition `json:"results"`
}

type AssociationDefinition struct {
	Category HubspotAssociationCategory `json:"category"`
	TypeId   int                        `json:"typeId"`
	Label    string                     `json:"label"`
}

type AssociationCreateDefinitionOptions struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

type AssociationUpdateDefinitionOptions struct {
	Label string `json:"label"`
	TypeId int64 `json:"associationTypeId"`
}

type AssocationDeleteDefinitionQuery struct {
	FromObjectType string 
	ToObjectType   string 
	TypeId         int64  
}

func (a *associations) List(options *AssociationListOptions, query *AssociationListQuery) (*AssociationList, error) {
	u := fmt.Sprintf("crm/v4/objects/%s/%s/associations/%s", query.FromObjectType, strconv.FormatInt(query.FromObjectId, 10), query.ToObjectType)
	req, err := a.client.newHttpRequest("GET", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.association.List(): newHttpRequest(): %v", err)
	}

	al := &AssociationList{}

	err = a.client.do(req, al)
	if err != nil {
		return nil, fmt.Errorf("client.association.List(): do(): %v", err)
	}

	return al, nil
}

func (a *associations) Create(options *[]AssociationCreateOptions, query *AssociationCreateOrDeleteQuery) (*AssociationCreateOutput, error) {
	u := fmt.Sprintf("/crm/v4/objects/%s/%s/associations/%s/%s", query.FromObjectType, strconv.FormatInt(query.FromObjectId, 10), query.ToObjectType, strconv.FormatInt(query.ToObjectId, 10))
	req, err := a.client.newHttpRequest("PUT", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.association.Create(): newHttpRequest(): %v", err)
	}

	aco := &AssociationCreateOutput{}

	err = a.client.do(req, aco)
	if err != nil {
		return nil, fmt.Errorf("client.association.Create(): do(): %v", err)
	}

	return aco, nil
}

func (a *associations) Delete(query *AssociationCreateOrDeleteQuery) (error) {
	u := fmt.Sprintf("/crm/v4/objects/%s/%s/associations/%s/%s", query.FromObjectType, strconv.FormatInt(query.FromObjectId, 10), query.ToObjectType, strconv.FormatInt(query.ToObjectId, 10))
	req, err := a.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.association.Archive(): newHttpRequest(): %v", err)
	}

	return a.client.do(req, nil)
}

func (a *associations) ReadDefinition(query *AssociationDefinitionQuery) (*AssociationDefinitionOutput, error) {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", query.FromObjectType, query.ToObjectType)
	req, err := a.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.association.ReadDefinition(): newHttpRequest(): %v", err)
	}

	ard := &AssociationDefinitionOutput{}

	err = a.client.do(req, ard)
	if err != nil {
		return nil, fmt.Errorf("client.association.ReadDefinition(): do(): %v", err)
	}

	return ard, nil
}

func (a *associations) CreateDefinition(options *AssociationCreateDefinitionOptions, query *AssociationDefinitionQuery) (*AssociationDefinitionOutput, error) {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", query.FromObjectType, query.ToObjectType)
	req, err := a.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.association.CreateDefinition(): newHttpRequest(): %v", err)
	}

	ard := &AssociationDefinitionOutput{}

	err = a.client.do(req, ard)
	if err != nil {
		return nil, fmt.Errorf("client.association.CreateDefinition(): do(): %v", err)
	}

	return ard, nil
}

func (a *associations) UpdateDefinition(options *AssociationUpdateDefinitionOptions, query *AssociationDefinitionQuery) (error) {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", query.FromObjectType, query.ToObjectType)
	req, err := a.client.newHttpRequest("PUT", u, options)
	if err != nil {
		return fmt.Errorf("client.association.UpdateDefinition(): newHttpRequest(): %v", err)
	}

	return a.client.do(req, nil)
}

func (a *associations) DeleteDefinition(query *AssocationDeleteDefinitionQuery) (error) {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels/%s", query.FromObjectType, query.ToObjectType, strconv.FormatInt(query.TypeId, 10))
	req, err := a.client.newHttpRequest("PUT", u, query)
	if err != nil {
		return fmt.Errorf("client.association.UpdateDefinition(): newHttpRequest(): %v", err)
	}

	return a.client.do(req, nil)
}