package crm

import (
	"fmt"
)

type Calls interface {
	ListAssociations(query *CallAssociationsQuery, callId string, toObjectType string) (*CallAssociations, error)
	Associate(callId string, toObjectType string, toObjectId string, associationType string) (*Call, error)
	Disassociate(callId string, toObjectType string, toObjectId string, associationType string) (error)
	List(query *CallListQuery) (*CallList, error)
	Create(options *CallCreateOrUpdateOptions) (*Call, error)
	Read(query *CallReadQuery, callId string) (*Call, error)
	Update(options *CallCreateOrUpdateOptions, callId string) (*Call, error)
	Archive(callId string) (error)
	BatchArchive(callIds []string) (error)
	BatchCreate(options *CallBatchCreateOptions) (*CallBatchOutput, error)
	BatchRead(options *CallBatchReadOptions) (*CallBatchOutput, error)
	BatchUpdate(options *CallBatchUpdateOptions) (*CallBatchOutput, error)
	Search(options *CallSearchOptions) (*CallSearchResults, error)
	Merge(options *CallMergeOptions) (*Call, error)
}

type calls struct {
	client *Client
}

type CallAssociationsQuery struct {
	ListAssociationsQuery
}

type CallAssociations struct {
	Results []CallAssociation `json:"results"`
	Pagination
}

type CallAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CallListQuery struct {
	ListQuery
}

type CallList struct {
	Calls      []Call `json:"results"`
	Pagination
}

type Call struct {
	Id         string         `json:"id"`
	Properties CallProperties `json:"properties"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type CallProperties struct {
	CreateDate         string `json:"createdate"`
	HsCallBody         string `json:"hs_call_body,omitempty"`
	HsCallDuration     string `json:"hs_call_duration,omitempty"`
	HsCallFromNumber   string `json:"hs_call_from_number,omitempty"`
	HsCallRecordingUrl string `json:"hs_call_recording_url,omitempty"`
	HsCallStatus       string `json:"hs_call_status,omitempty"`
	HsCallTitle        string `json:"hs_call_title,omitempty"`
	HsCallToNumber     string `json:"hs_call_to_number,omitempty"`
	HsTimestamp        string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId     string `json:"hubspot_owner_id,omitempty"`
}

type CallCreateOrUpdateOptions struct {
	Properties CallCreateOrUpdateProperties `json:"properties"`
}

type CallCreateOrUpdateProperties struct {
	HsCallBody         string `json:"hs_call_body,omitempty"`
	HsCallDuration     string `json:"hs_call_duration,omitempty"`
	HsCallFromNumber   string `json:"hs_call_from_number,omitempty"`
	HsCallRecordingUrl string `json:"hs_call_recording_url,omitempty"`
	HsCallStatus       string `json:"hs_call_status,omitempty"`
	HsCallTitle        string `json:"hs_call_title,omitempty"`
	HsCallToNumber     string `json:"hs_call_to_number,omitempty"`
	HsTimestamp        string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId     string `json:"hubspot_owner_id,omitempty"`
}

type CallReadQuery struct {
	ReadQuery
}

type CallUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}

type CallBatchOutput struct {
	Status      string `json:"status"`
	Results     []Call `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type CallBatchReadOptions struct {
	BatchReadOptions
}

type CallBatchCreateOptions struct {
	Inputs []CallCreateOrUpdateOptions `json:"inputs"`
}

type CallBatchUpdateOptions struct {
	Inputs []CallBatchUpdateProperties `json:"inputs"`
}

type CallBatchUpdateProperties struct {
	Id         string                       `json:"id"`
	Properties CallCreateOrUpdateProperties `json:"properties"`
}

type CallSearchOptions struct {
	SearchOptions
}

type CallSearchResults struct {
	Total      int64  `json:"total"`
	Results    []Call `json:"results"`
	Pagination
}

type CallMergeOptions struct {
	MergeOptions
}

func (z *calls) ListAssociations(query *CallAssociationsQuery, callId string, toObjectType string) (*CallAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/%s/associations/%s", callId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.calls.ListAssociations(): newHttpRequest(): %v", err)
	}

	ca := &CallAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, fmt.Errorf("client.calls.ListAssociations(): do(): %v", err)
	}
	
	return ca, nil
}

func (z *calls) Associate(callId string, toObjectType string, toObjectId string, associationType string) (*Call, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/%s/associations/%s/%s/%s", callId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Associate(): newHttpRequest(): %v", err)
	}

	call := &Call{}

	err = z.client.do(req, call)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Associate(): do(): %v", err)
	}
	
	return call, nil
}

func (z *calls) Disassociate(callId string, toObjectType string, toObjectId string, associationType string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/%s/associations/%s/%s/%s", callId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.calls.Disassociate(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *calls) List(query *CallListQuery) (*CallList, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.calls.List(): newHttpRequest(): %v", err)
	}

	cl := &CallList{}

	err = z.client.do(req, cl)
	if err != nil {
		return nil, fmt.Errorf("client.calls.List(): do(): %v", err)
	}
	
	return cl, nil
}

func (z *calls) Create(options *CallCreateOrUpdateOptions) (*Call, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Create(): newHttpRequest(): %v", err)
	}

	call := &Call{}

	err = z.client.do(req, call)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Create(): do(): %v", err)
	}
	
	return call, nil
}

func (z *calls) Read(query *CallReadQuery, callId string) (*Call, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", callId)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Read(): newHttpRequest(): %v", err)
	}

	call := &Call{}

	err = z.client.do(req, call)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Read(): do(): %+v", err)
	}

	return call, nil
}

func (z *calls) Update(options *CallCreateOrUpdateOptions, callId string) (*Call, error) {
	u := fmt.Sprintf("crm/v3/objects/calls/%s", callId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Update(): newHttpRequest(): %v", err)
	}

	call := &Call{}

	err = z.client.do(req, call)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Update(): do(): %+v", err)
	}

	return call, nil
}

func (z *calls) Archive(callId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/calls/%s", callId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.calls.Archive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *calls) BatchArchive(callIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, callId := range callIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: callId})
	}

	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.calls.BatchArchive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *calls) BatchCreate(options *CallBatchCreateOptions) (*CallBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.calls.BatchCreate(): newHttpRequest(): %v", err)
	}

	calls := &CallBatchOutput{}

	err = z.client.do(req, calls)
	if err != nil {
		return nil, fmt.Errorf("client.calls.BatchCreate(): do(): %+v", err)
	}

	return calls, nil
}

func (z *calls) BatchRead(options *CallBatchReadOptions) (*CallBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.calls.BatchUpdate(): newHttpRequest(): %v", err)
	}

	calls := &CallBatchOutput{}

	err = z.client.do(req, calls)
	if err != nil {
		return nil, fmt.Errorf("client.calls.BatchUpdate(): do(): %+v", err)
	}

	return calls, nil
}

func (z *calls) BatchUpdate(options *CallBatchUpdateOptions) (*CallBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.calls.BatchUpdate(): newHttpRequest(): %v", err)
	}

	calls := &CallBatchOutput{}

	err = z.client.do(req, calls)
	if err != nil {
		return nil, fmt.Errorf("client.calls.BatchUpdate(): do(): %+v", err)
	}

	return calls, nil
}

func (z *calls) Search(options *CallSearchOptions) (*CallSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Search(): newHttpRequest(): %v", err)
	}

	calls := &CallSearchResults{}

	err = z.client.do(req, calls)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Search(): do(): %+v", err)
	}

	return calls, nil
}

func (z *calls) Merge(options *CallMergeOptions) (*Call, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/merge")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Merge(): newHttpRequest(): %v", err)
	}

	call := &Call{}

	err = z.client.do(req, call)
	if err != nil {
		return nil, fmt.Errorf("client.calls.Merge(): do(): %+v", err)
	}

	return call, nil
}