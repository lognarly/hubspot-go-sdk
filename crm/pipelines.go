package crm

import (
	"fmt"
)

type Pipelines interface {
	ListStages(objectType string, pipelineId string) (*PipelineStageList, error)
	CreateStage(objectType string, pipelineId string, options *PipelineStageCreateOptions) (*PipelineStage, error)
}

type pipelines struct {
	client *Client
}

type PipelineStageList struct {
	Results []PipelineStage `json:"results"`
}

type PipelineStage struct {
	Label        string      `json:"label"`
	DisplayOrder int64       `json:"displayOrder"`
	Metadata     interface{} `json:"metadata"`
	CreatedAt    string      `json:"createdAt"`
	UpdatedAt    string      `json:"updatedAt"`
	Archived     bool        `json:"archived"`
	Id           string      `json:"id"`
}

type PipelineStageCreateOptions struct {
	Label        string `json:"label"`
	DisplayOrder int64 `json:"displayOrder"`
	Metadata     interface{} `json:"metadata"`
}

func (z *pipelines) ListStages(objectType string, pipelineId string) (*PipelineStageList, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s/stages", objectType, pipelineId)
	req, err := z.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.ListStages(): newHttpRequest(): %v", err)
	}

	psl := &PipelineStageList{}

	err = z.client.do(req, psl)
	if err != nil {
		return nil, fmt.Errorf("client.companies.ListStages(): do(): %v", err)
	}
	
	return psl, nil
}

func (z *pipelines) CreateStage(objectType string, pipelineId string, options *PipelineStageCreateOptions) (*PipelineStage, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s/stages", objectType, pipelineId)
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.CreateStage(): newHttpRequest(): %v", err)
	}

	ps := &PipelineStage{}

	err = z.client.do(req, ps)
	if err != nil {
		return nil, fmt.Errorf("client.companies.CreateStage(): do(): %v", err)
	}
	
	return ps, nil
}