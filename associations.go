package main

import (
	"fmt"
)

type Associations interface {
	List()
	Create()
	Delete()
}

type AssociationQuery string

type AssociationListOptions struct {
	ListOptions
}

type AssociationList struct {
	Assocations []Association `json:"results"`
}

type Association struct {
	ToObjectId int `json:"toObjectId"`
	
}