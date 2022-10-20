package crm

import (
	"fmt"
)

type Pipelines interface {
	ListStages(objectType string, pipelineId string) (*PipelineStageList, error)
	CreateStage(objectType string, pipelineId string, options *PipelineStageCreateOrUpdateOptions) (*PipelineStage, error)
	ReadStage(objectType string, pipelineId string, stageId string) (*PipelineStage, error)
	UpdateStage(options *PipelineStageCreateOrUpdateOptions, objectType string, pipelineId string, stageId string) (*PipelineStage, error)
	ReplaceStage(options *PipelineStageCreateOrUpdateOptions, objectType string, pipelineId string, stageId string) (*PipelineStage, error)
	DeleteStage(objectType string, pipelineId string, stageId string) (error)
	List(objectType string) (*PipelineList, error)
	Create(options *PipelineCreateOrUpdateOptions, objectType string) (*Pipeline, error)
	Read(objectType string, pipelineId string) (*Pipeline, error)
	Update(options *PipelineCreateOrUpdateOptions, objectType string, pipelineId string) (*Pipeline, error)
	Replace(options *PipelineCreateOrUpdateOptions, objectType string, pipelineId string) (*Pipeline, error)
	Delete(objectType string, pipelineId string) (error)
	Audit(objectType string, pipelineId string) (*PipelineAuditList, error)
	AuditStage(objectType string, pipelineId string, stageId string) (*PipelineAuditList, error)
}

type pipelines struct {
	client *Client
}

type PipelineStageList struct {
	Results []PipelineStage `json:"results"`
}

type PipelineStage struct {
	Label        string      `json:"label,omitempty"`
	DisplayOrder int64       `json:"displayOrder,omitempty"`
	Metadata     interface{} `json:"metadata,omitempty"`
	CreatedAt    string      `json:"createdAt"`
	UpdatedAt    string      `json:"updatedAt"`
	Archived     bool        `json:"archived"`
	Id           string      `json:"id"`
}

type PipelineStageCreateOrUpdateOptions struct {
	Label        string      `json:"label,omitempty"`
	DisplayOrder int64       `json:"displayOrder,omitempty"`
	Metadata     interface{} `json:"metadata,omitempty"`
}

type PipelineList struct {
	Results []Pipeline `json:"results"`
}

type Pipeline struct {
	Label        string          `json:"label"`
	DisplayOrder string          `json:"displayOrder"`
	CreatedAt    string          `json:"createdAt"`
	UpdatedAt    string          `json:"updatedAt"`
	Archived     bool            `json:"archived"`
	Id           string          `json:"id"`
	Stages       []PipelineStage `json:"stages"`
}

type PipelineCreateOrUpdateOptions struct {
	Label        string                               `json:"label"`
	DisplayOrder int64                                `json:"displayOrder"`
	Stages       []PipelineStageCreateOrUpdateOptions `json:"stages"`
}

type PipelineAuditList struct {
	Results []PipelineAudit `json:"results"`
}

type PipelineAudit struct {
	PortalId   int64       `json:"portalId"`
	Identifier string      `json:"identifier"`
	Action     string      `json:"action"`
	Timestamp  string      `json:"timestamp"`
	Message    string      `json:"message"`
	RawObject  interface{} `json:"rawObject"`
	FromUserId int64       `json:"fromUserId"`
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
		return nil, fmt.Errorf("client.pipelines.ListStages(): do(): %v", err)
	}
	
	return psl, nil
}

func (z *pipelines) CreateStage(objectType string, pipelineId string, options *PipelineStageCreateOrUpdateOptions) (*PipelineStage, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s/stages", objectType, pipelineId)
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.CreateStage(): newHttpRequest(): %v", err)
	}

	ps := &PipelineStage{}

	err = z.client.do(req, ps)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.CreateStage(): do(): %v", err)
	}
	
	return ps, nil
}

func (z *pipelines) ReadStage(objectType string, pipelineId string, stageId string) (*PipelineStage, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s/stages/%s", objectType, pipelineId, stageId)

	req, err := z.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.ReadStage(): newHttpRequest(): %v", err)
	}

	ps := &PipelineStage{}

	err = z.client.do(req, ps)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.ReadStage(): do(): %v", err)
	}
	
	return ps, nil
}

func (z *pipelines) UpdateStage(options *PipelineStageCreateOrUpdateOptions, objectType string, pipelineId string, stageId string) (*PipelineStage, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s/stages/%s", objectType, pipelineId, stageId)

	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.UpdateStage(): newHttpRequest(): %v", err)
	}

	ps := &PipelineStage{}

	err = z.client.do(req, ps)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.UpdateStage(): do(): %v", err)
	}
	
	return ps, nil
}

func (z *pipelines) ReplaceStage(options *PipelineStageCreateOrUpdateOptions, objectType string, pipelineId string, stageId string) (*PipelineStage, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s/stages/%s", objectType, pipelineId, stageId)

	req, err := z.client.newHttpRequest("PUT", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.ReplaceStage(): newHttpRequest(): %v", err)
	}

	ps := &PipelineStage{}

	err = z.client.do(req, ps)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.ReplaceStage(): do(): %v", err)
	}
	
	return ps, nil
}

func (z *pipelines) DeleteStage(objectType string, pipelineId string, stageId string) (error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s/stages/%s", objectType, pipelineId, stageId)

	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.pipelines.DeleteStage(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *pipelines) List(objectType string) (*PipelineList, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s", objectType)

	req, err := z.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.List(): newHttpRequest(): %v", err)
	}

	pl := &PipelineList{}

	err = z.client.do(req, pl)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.List(): do(): %v", err)
	}
	
	return pl, nil
}

func (z *pipelines) Create(options *PipelineCreateOrUpdateOptions, objectType string) (*Pipeline, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s", objectType)

	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Create(): newHttpRequest(): %v", err)
	}

	pipeline := &Pipeline{}

	err = z.client.do(req, pipeline)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Create(): do(): %v", err)
	}
	
	return pipeline, nil
}

func (z *pipelines) Read(objectType string, pipelineId string) (*Pipeline, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s", objectType, pipelineId)

	req, err := z.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Read(): newHttpRequest(): %v", err)
	}

	pipeline := &Pipeline{}

	err = z.client.do(req, pipeline)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Read(): do(): %v", err)
	}
	
	return pipeline, nil
}

func (z *pipelines) Update(options *PipelineCreateOrUpdateOptions, objectType string, pipelineId string) (*Pipeline, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s", objectType, pipelineId)

	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Update(): newHttpRequest(): %v", err)
	}

	pipeline := &Pipeline{}

	err = z.client.do(req, pipeline)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Update(): do(): %v", err)
	}
	
	return pipeline, nil
}

func (z *pipelines) Replace(options *PipelineCreateOrUpdateOptions, objectType string, pipelineId string) (*Pipeline, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s", objectType, pipelineId)
	// Add validateReferencesBeforeDelete boolean in URL??
	req, err := z.client.newHttpRequest("PUT", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Replace(): newHttpRequest(): %v", err)
	}

	pipeline := &Pipeline{}

	err = z.client.do(req, pipeline)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Replace(): do(): %v", err)
	}
	
	return pipeline, nil
}

func (z *pipelines) Delete(objectType string, pipelineId string) (error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s", objectType, pipelineId)

	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.pipelines.Delete(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *pipelines) Audit(objectType string, pipelineId string) (*PipelineAuditList, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s/audit", objectType, pipelineId)

	req, err := z.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Audit(): newHttpRequest(): %v", err)
	}

	pal := &PipelineAuditList{}

	err = z.client.do(req, pal)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Audit(): do(): %v", err)
	}
	
	return pal, nil
}

func (z *pipelines) AuditStage(objectType string, pipelineId string, stageId string) (*PipelineAuditList, error) {
	u := fmt.Sprintf("/crm/v3/pipelines/%s/%s/stages/%s/audit", objectType, pipelineId, stageId)

	req, err := z.client.newHttpRequest("GET", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Audit(): newHttpRequest(): %v", err)
	}

	pal := &PipelineAuditList{}

	err = z.client.do(req, pal)
	if err != nil {
		return nil, fmt.Errorf("client.pipelines.Audit(): do(): %v", err)
	}
	
	return pal, nil
}