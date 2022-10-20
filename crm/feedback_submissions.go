package crm

import (
    "fmt"
)

type FeedbackSubmissions interface {
    ListAssociations(feedbackSubmissionId string, toObjectType string, query *FeedbackSubmissionListAssociationQuery) (*FeedbackSubmissionAssociations, error)
    List(query *FeedbackSubmissionListQuery) (*FeedbackSubmissionList, error)
    Read(feedbackSubmissionId string, query *FeedbackSubmissionReadQuery) (*FeedbackSubmission, error)
    BatchRead(options *FeedbackSubmissionBatchReadOptions) (*FeedbackSubmissionBatchReadResults, error)
    Search(options *FeedbackSubmissionSearchOptions) (*FeedbackSubmissionSearchResults, error)
}

type feedbackSubmissions struct {
    client *Client
}

type FeedbackSubmissionListAssociationQuery struct {
	ListAssociationsQuery
}

type FeedbackSubmissionAssociations struct {
	Results    []FeedbackSubmissionAssociation `json:"results"`
	Pagination
}

type FeedbackSubmissionAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type FeedbackSubmissionListQuery struct {
	ListQuery
}

type FeedbackSubmissionList struct {
	Results []FeedbackSubmission `json:"results"`
	Pagination
}

type FeedbackSubmission struct {
	Id         string                       `json:"id"`
	Properties FeedbackSubmissionProperties `json:"properties"`
}

type FeedbackSubmissionProperties struct {
	HsContent        string `json:"hs_content,omitempty"`
	HsIngestionId    string `json:"hs_ingestion_id,omitempty"`
	HsResponseGroup  string `json:"hs_response_group,omitempty"`
	HsSubmissionName string `json:"hs_submission_name,omitempty"`
	HsSurveyChannel  string `json:"hs_survey_channel,omitempty"`
	HsSurveyId       string `json:"hs_survey_id,omitempty"`
	HsSurveyName     string `json:"hs_survey_name,omitempty"`
	HsSurveyType     string `json:"hs_survey_type,omitempty"`
	HsValue          string `json:"hs_value,omitempty"`
}

type FeedbackSubmissionReadQuery struct {
	Properties            []string `url:"properties,omitempty"`
	PropertiesWithHistory []string `url:"propertiesWithHistory,omitempty"`
	Associations          []string `url:"associations,omitempty"`
	Archived              bool     `url:"archived,omitempty"`
	IdProperty            string   `url:"idProperty,omitempty"`
}

type FeedbackSubmissionBatchReadOptions struct {
	BatchReadOptions
}

type FeedbackSubmissionBatchReadResults struct {
	Status      string               `json:"status"`
	Results     []FeedbackSubmission `json:"results"`
	RequestedAt string               `json:"requestedAt"`
	StartedAt   string               `json:"startedAt"`
	UpdatedAt   string               `json:"updatedAt"`
	Archived    bool                 `json:"archived"`
}

type FeedbackSubmissionSearchOptions struct {
	SearchOptions
}

type FeedbackSubmissionSearchResults struct {
	Total      int64                `json:"total"`
	Results    []FeedbackSubmission `json:"results"`
	Pagination
}

func (z *feedbackSubmissions) ListAssociations(feedbackSubmissionId string, toObjectType string, query *FeedbackSubmissionListAssociationQuery) (*FeedbackSubmissionAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/feedback_submissions/%s/associations/%s", feedbackSubmissionId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.feedbackSubmissions.ListAssociations(): newHttpRequest(): %v", err)
	}

	fsa := &FeedbackSubmissionAssociations{}

	err = z.client.do(req, fsa)
	if err != nil {
		return nil, fmt.Errorf("client.feedbackSubmissions.ListAssociations(): do(): %v", err)
	}
	
	return fsa, nil
}

func (z *feedbackSubmissions) List(query *FeedbackSubmissionListQuery) (*FeedbackSubmissionList, error) {
	u := fmt.Sprintf("/crm/v3/objects/feedback_submissions")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.feedbackSubmissions.List(): newHttpRequest(): %v", err)
	}

	fsl := &FeedbackSubmissionList{}

	err = z.client.do(req, fsl)
	if err != nil {
		return nil, fmt.Errorf("client.feedbackSubmissions.List(): do(): %v", err)
	}
	
	return fsl, nil
}

func (z *feedbackSubmissions) Read(feedbackSubmissionId string, query *FeedbackSubmissionReadQuery) (*FeedbackSubmission, error) {
	u := fmt.Sprintf("/crm/v3/objects/feedback_submissions/%s", feedbackSubmissionId)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.feedbackSubmissions.Read(): newHttpRequest(): %v", err)
	}

	fs := &FeedbackSubmission{}

	err = z.client.do(req, fs)
	if err != nil {
		return nil, fmt.Errorf("client.feedbackSubmissions.Read(): do(): %v", err)
	}
	
	return fs, nil
}

func (z *feedbackSubmissions) BatchRead(options *FeedbackSubmissionBatchReadOptions) (*FeedbackSubmissionBatchReadResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/feedback_submissions/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.feedbackSubmissions.BatchRead(): newHttpRequest(): %v", err)
	}

	fsbrr := &FeedbackSubmissionBatchReadResults{}

	err = z.client.do(req, fsbrr)
	if err != nil {
		return nil, fmt.Errorf("client.feedbackSubmissions.BatchRead(): do(): %v", err)
	}
	
	return fsbrr, nil
}

func (z *feedbackSubmissions) Search(options *FeedbackSubmissionSearchOptions) (*FeedbackSubmissionSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/feedback_submissions/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.feedbackSubmissions.Search(): newHttpRequest(): %v", err)
	}

	fsso := &FeedbackSubmissionSearchResults{}

	err = z.client.do(req, fsso)
	if err != nil {
		return nil, fmt.Errorf("client.feedbackSubmissions.Search(): do(): %v", err)
	}
	
	return fsso, nil
}