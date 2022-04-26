package main

import (
	//"encoding/json"
)

type ListOptions struct {
	Limit    int    `url:"limit,omitempty"`
	After    string `url:"after,omitempty"`
	Archived bool   `url:"archived,omitempty"`
}

type Pagination struct {
	Paging Paging `json:"paging"`
}

type Paging struct {
	Next Next `json:"next"`
}

type Next struct {
	After string `json:"after"`
}

type HubSpotError struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
	CorrelationId string `json:"correlationId"`
}

type FilterOperator string

const (
	EqualTo            FilterOperator = "EQ"
	NotEqualTo         FilterOperator = "NEQ"
	LessThan           FilterOperator = "LT"
	LessThanEqualTo    FilterOperator = "LTE"
	GreaterThan        FilterOperator = "GT"
	GreaterThanEqualTo FilterOperator = "GTE"
	Between            FilterOperator = "BETWEEN"
	In                 FilterOperator = "IN"
	NotIn              FilterOperator = "NOT_IN"
	HasProperty        FilterOperator = "HAS_PROPERTY"
	NotHasProperty     FilterOperator = "NOT_HAS_PROPERTY"
	ContainsToken      FilterOperator = "CONTAINS_TOKEN"
	NotContainsToken   FilterOperator = "NOT_CONTAINS_TOKEN"
)

type SearchOptions struct {
	FilterGroups []FilterGroups `json:"filterGroups"`
	Sorts        []string       `json:"sorts"`
	Query        string         `json:"query"`
	Properties   []string       `json:"properties"`
}

type FilterGroups struct {
	Filters []Filters `json:"filters"`
}

type Filters struct {
	Value        string                `json:"value,omitempty"`
	Values       []string              `json:"values,omitempty"`
	PropertyName string                `json:"propertyName,omitempty"`
	Operator     FilterOperator        `json:"operator"`
}

type BatchInputs struct {
	Inputs []BatchInput `json:"inputs"`
}

type BatchInput struct {
	Id string `json:"id"`
}

type MergeOptions struct {
	PrimaryObjectId string `json:"primaryObjectId"`
	ObjectIdToMerge string `json:"objectIdToMerge"`
}