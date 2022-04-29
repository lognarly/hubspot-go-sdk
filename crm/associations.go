package crm

import (
	"fmt"
	"strconv"
)

type Associations interface {
	List(fromObjectType string, fromObjectId int64, toObjectType string, query *AssociationListQuery) (*AssociationList, error)
	Create(options *[]AssociationCreateOptions, fromObjectType string, fromObjectId int64, toObjectType string, toObjectId int64) (*AssociationCreateOutput, error)
	Delete(fromObjectType string, fromObjectId int64, toObjectType string, toObjectId int64) (error)
	ReadDefinition(fromObjectType string, toObjectType string) (*AssociationDefinitionOutput, error)
	CreateDefinition(options *AssociationCreateDefinitionOptions, fromObjectType string, toObjectType string) (*AssociationDefinitionOutput, error)
	UpdateDefinition(options *AssociationUpdateDefinitionOptions, fromObjectType string, toObjectType string) (error)
	DeleteDefinition(fromObjectType string, toObjectType string, typeId int64) (error)
}

type associations struct {
	client *Client
}

type AssociationListQuery struct {
	ListAssociationsQuery
}

type AssociationList struct {
	Assocations []AssociationListResult `json:"results"`
	Pagination
}

type AssociationListResult struct {
	ToObjectId       int64             `json:"toObjectId"`
	AssociationTypes []AssociationType `json:"associationTypes"`
}

type HubspotAssociationCategory string

const (
	HubSpotDefined    HubspotAssociationCategory = "HUBSPOT_DEFINED"
	UserDefined       HubspotAssociationCategory = "USER_DEFINED"
	IntegratorDefined HubspotAssociationCategory = "INTEGRATOR_DEFINED"
)

type AssociationType struct {
	Category HubspotAssociationCategory `json:"category,omitempty"`
	TypeId   int64                      `json:"typeId,omitempty"`
	Label    string                     `json:"label,omitempty"`
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
	Labels           []string `json:"labels,omitempty"`
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

func (a *associations) List(fromObjectType string, fromObjectId int64, toObjectType string, query *AssociationListQuery) (*AssociationList, error) {
	u := fmt.Sprintf("crm/v4/objects/%s/%s/associations/%s", fromObjectType, strconv.FormatInt(fromObjectId, 10), toObjectType)
	req, err := a.client.newHttpRequest("GET", u, query)
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

func (a *associations) Create(options *[]AssociationCreateOptions, fromObjectType string, fromObjectId int64, toObjectType string, toObjectId int64) (*AssociationCreateOutput, error) {
	u := fmt.Sprintf("/crm/v4/objects/%s/%s/associations/%s/%s", fromObjectType, strconv.FormatInt(fromObjectId, 10), toObjectType, strconv.FormatInt(toObjectId, 10))
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

func (a *associations) Delete(fromObjectType string, fromObjectId int64, toObjectType string, toObjectId int64) (error) {
	u := fmt.Sprintf("/crm/v4/objects/%s/%s/associations/%s/%s", fromObjectType, strconv.FormatInt(fromObjectId, 10), toObjectType, strconv.FormatInt(toObjectId, 10))
	req, err := a.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.association.Archive(): newHttpRequest(): %v", err)
	}

	return a.client.do(req, nil)
}

func (a *associations) ReadDefinition(fromObjectType string, toObjectType string) (*AssociationDefinitionOutput, error) {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", fromObjectType, toObjectType)
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

func (a *associations) CreateDefinition(options *AssociationCreateDefinitionOptions, fromObjectType string, toObjectType string) (*AssociationDefinitionOutput, error) {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", fromObjectType, toObjectType)
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

func (a *associations) UpdateDefinition(options *AssociationUpdateDefinitionOptions, fromObjectType string, toObjectType string) (error) {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", fromObjectType, toObjectType)
	req, err := a.client.newHttpRequest("PUT", u, options)
	if err != nil {
		return fmt.Errorf("client.association.UpdateDefinition(): newHttpRequest(): %v", err)
	}

	return a.client.do(req, nil)
}

func (a *associations) DeleteDefinition(fromObjectType string, toObjectType string, typeId int64) (error) {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels/%s", fromObjectType, toObjectType, strconv.FormatInt(typeId, 10))
	req, err := a.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.association.UpdateDefinition(): newHttpRequest(): %v", err)
	}

	return a.client.do(req, nil)
}