package crm

import (
	"fmt"
)

type Meetings interface {
	ListAssociations(query *MeetingAssociationsQuery, meetingId string, toObjectType string) (*MeetingAssociations, error)
	Associate(meetingId string, toObjectType string, toObjectId string, associationType string) (*Meeting, error)
	Disassociate(meetingId string, toObjectType string, toObjectId string, associationType string) (error)
	List(query *MeetingListQuery) (*MeetingList, error)
	Create(options *MeetingCreateOrUpdateOptions) (*Meeting, error)
	Read(query *MeetingReadQuery, meetingId string) (*Meeting, error)
	Update(options *MeetingCreateOrUpdateOptions, meetingId string) (*Meeting, error)
	Archive(meetingId string) (error)
	BatchArchive(meetingIds []string) (error)
	BatchCreate(options *MeetingBatchCreateOptions) (*MeetingBatchOutput, error)
	BatchRead(options *MeetingBatchReadOptions) (*MeetingBatchOutput, error)
	BatchUpdate(options *MeetingBatchUpdateOptions) (*MeetingBatchOutput, error)
	Search(options *MeetingSearchOptions) (*MeetingSearchResults, error)
	Merge(options *MeetingMergeOptions) (*Meeting, error)
}

type meetings struct {
	client *Client
}

type MeetingAssociationsQuery struct {
	ListAssociationsQuery
}

type MeetingAssociations struct {
	Results []MeetingAssociation `json:"results"`
	Pagination
}

type MeetingAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type MeetingListQuery struct {
	ListQuery
}

type MeetingList struct {
	Meetings      []Meeting `json:"results"`
	Pagination
}

type Meeting struct {
	Id         string         `json:"id"`
	Properties MeetingProperties `json:"properties"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type MeetingProperties struct {
	CreateDate             string `json:"createdate"`
	HsInternalMeetingNotes string `json:"hs_internal_meeting_notes,omitempty"`
	HsLastModifiedDate     string `json:"hs_lastmodifieddate"`
	HsMeetingBody          string `json:"hs_meeting_body,omitempty"`
	HsMeetingEndTime       string `json:"hs_meeting_end_time,omitempty"`
	HsMeetingExternalUrl   string `json:"hs_meeting_external_url,omitempty"`
	HsMeetingLocation      string `json:"hs_meeting_location,omitempty"`
	HsMeetingOutcome       string `json:"hs_meeting_outcome,omitempty"`
	HsMeetingStartTime     string `json:"hs_meeting_start_time,omitempty"`
	HsMeetingTitle         string `json:"hs_meeting_title,omitempty"`
	HsTimestamp            string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId         string `json:"hubspot_owner_id,omitempty"`
}

type MeetingCreateOrUpdateOptions struct {
	Properties MeetingCreateOrUpdateProperties `json:"properties"`
}

type MeetingCreateOrUpdateProperties struct {
	HsInternalMeetingNotes string `json:"hs_internal_meeting_notes,omitempty"`
	HsMeetingBody          string `json:"hs_meeting_body,omitempty"`
	HsMeetingEndTime       string `json:"hs_meeting_end_time,omitempty"`
	HsMeetingExternalUrl   string `json:"hs_meeting_external_url,omitempty"`
	HsMeetingLocation      string `json:"hs_meeting_location,omitempty"`
	HsMeetingOutcome       string `json:"hs_meeting_outcome,omitempty"`
	HsMeetingStartTime     string `json:"hs_meeting_start_time,omitempty"`
	HsMeetingTitle         string `json:"hs_meeting_title,omitempty"`
	HsTimestamp            string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId         string `json:"hubspot_owner_id,omitempty"`
}

type MeetingReadQuery struct {
	ReadQuery
}

type MeetingUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}

type MeetingBatchOutput struct {
	Status      string `json:"status"`
	Results     []Meeting `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type MeetingBatchReadOptions struct {
	BatchReadOptions
}

type MeetingBatchCreateOptions struct {
	Inputs []MeetingCreateOrUpdateOptions `json:"inputs"`
}

type MeetingBatchUpdateOptions struct {
	Inputs []MeetingBatchUpdateProperties `json:"inputs"`
}

type MeetingBatchUpdateProperties struct {
	Id         string                       `json:"id"`
	Properties MeetingCreateOrUpdateProperties `json:"properties"`
}

type MeetingSearchOptions struct {
	SearchOptions
}

type MeetingSearchResults struct {
	Total      int64  `json:"total"`
	Results    []Meeting `json:"results"`
	Pagination
}

type MeetingMergeOptions struct {
	MergeOptions
}

func (z *meetings) ListAssociations(query *MeetingAssociationsQuery, meetingId string, toObjectType string) (*MeetingAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/%s/associations/%s", meetingId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.ListAssociations(): newHttpRequest(): %v", err)
	}

	ca := &MeetingAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.ListAssociations(): do(): %v", err)
	}
	
	return ca, nil
}

func (z *meetings) Associate(meetingId string, toObjectType string, toObjectId string, associationType string) (*Meeting, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/%s/associations/%s/%s/%s", meetingId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Associate(): newHttpRequest(): %v", err)
	}

	meeting := &Meeting{}

	err = z.client.do(req, meeting)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Associate(): do(): %v", err)
	}
	
	return meeting, nil
}

func (z *meetings) Disassociate(meetingId string, toObjectType string, toObjectId string, associationType string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/%s/associations/%s/%s/%s", meetingId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.meetings.Disassociate(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *meetings) List(query *MeetingListQuery) (*MeetingList, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.List(): newHttpRequest(): %v", err)
	}

	ml := &MeetingList{}

	err = z.client.do(req, ml)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.List(): do(): %v", err)
	}
	
	return ml, nil
}

func (z *meetings) Create(options *MeetingCreateOrUpdateOptions) (*Meeting, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Create(): newHttpRequest(): %v", err)
	}

	meeting := &Meeting{}

	err = z.client.do(req, meeting)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Create(): do(): %v", err)
	}
	
	return meeting, nil
}

func (z *meetings) Read(query *MeetingReadQuery, meetingId string) (*Meeting, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", meetingId)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Read(): newHttpRequest(): %v", err)
	}

	meeting := &Meeting{}

	err = z.client.do(req, meeting)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Read(): do(): %+v", err)
	}

	return meeting, nil
}

func (z *meetings) Update(options *MeetingCreateOrUpdateOptions, meetingId string) (*Meeting, error) {
	u := fmt.Sprintf("crm/v3/objects/meetings/%s", meetingId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Update(): newHttpRequest(): %v", err)
	}

	meeting := &Meeting{}

	err = z.client.do(req, meeting)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Update(): do(): %+v", err)
	}

	return meeting, nil
}

func (z *meetings) Archive(meetingId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/meetings/%s", meetingId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.meetings.Archive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *meetings) BatchArchive(meetingIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, meetingId := range meetingIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: meetingId})
	}

	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.meetings.BatchArchive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *meetings) BatchCreate(options *MeetingBatchCreateOptions) (*MeetingBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.BatchCreate(): newHttpRequest(): %v", err)
	}

	meetings := &MeetingBatchOutput{}

	err = z.client.do(req, meetings)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.BatchCreate(): do(): %+v", err)
	}

	return meetings, nil
}

func (z *meetings) BatchRead(options *MeetingBatchReadOptions) (*MeetingBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.BatchUpdate(): newHttpRequest(): %v", err)
	}

	meetings := &MeetingBatchOutput{}

	err = z.client.do(req, meetings)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.BatchUpdate(): do(): %+v", err)
	}

	return meetings, nil
}

func (z *meetings) BatchUpdate(options *MeetingBatchUpdateOptions) (*MeetingBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.BatchUpdate(): newHttpRequest(): %v", err)
	}

	meetings := &MeetingBatchOutput{}

	err = z.client.do(req, meetings)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.BatchUpdate(): do(): %+v", err)
	}

	return meetings, nil
}

func (z *meetings) Search(options *MeetingSearchOptions) (*MeetingSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Search(): newHttpRequest(): %v", err)
	}

	meetings := &MeetingSearchResults{}

	err = z.client.do(req, meetings)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Search(): do(): %+v", err)
	}

	return meetings, nil
}

func (z *meetings) Merge(options *MeetingMergeOptions) (*Meeting, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/merge")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Merge(): newHttpRequest(): %v", err)
	}

	meeting := &Meeting{}

	err = z.client.do(req, meeting)
	if err != nil {
		return nil, fmt.Errorf("client.meetings.Merge(): do(): %+v", err)
	}

	return meeting, nil
}