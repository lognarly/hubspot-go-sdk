package main

import (
	//"encoding/json"
)

type ListOptions struct {
	Limit int    `url:"limit,omitempty"`
	After string `url:"after,omitempty"`
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