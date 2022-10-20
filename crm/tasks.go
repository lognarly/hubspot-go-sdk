package crm

import (
	"fmt"
)

type Tasks interface {
	ListAssociations(query *TaskAssociationsQuery, taskId string, toObjectType string) (*TaskAssociations, error)
	Associate(taskId string, toObjectType string, toObjectId string, associationType string) (*Task, error)
	Disassociate(taskId string, toObjectType string, toObjectId string, associationType string) (error)
	List(query *TaskListQuery) (*TaskList, error)
	Create(options *TaskCreateOrUpdateOptions) (*Task, error)
	Read(query *TaskReadQuery, taskId string) (*Task, error)
	Update(options *TaskCreateOrUpdateOptions, taskId string) (*Task, error)
	Archive(taskId string) (error)
	BatchArchive(taskIds []string) (error)
	BatchCreate(options *TaskBatchCreateOptions) (*TaskBatchOutput, error)
	BatchRead(options *TaskBatchReadOptions) (*TaskBatchOutput, error)
	BatchUpdate(options *TaskBatchUpdateOptions) (*TaskBatchOutput, error)
	Search(options *TaskSearchOptions) (*TaskSearchResults, error)
	Merge(options *TaskMergeOptions) (*Task, error)
}

type tasks struct {
	client *Client
}

type TaskAssociationsQuery struct {
	ListAssociationsQuery
}

type TaskAssociations struct {
	Results []TaskAssociation `json:"results"`
	Pagination
}

type TaskAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type TaskListQuery struct {
	ListQuery
}

type TaskList struct {
	Tasks      []Task `json:"results"`
	Pagination
}

type Task struct {
	Id         string         `json:"id"`
	Properties TaskProperties `json:"properties"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type TaskProperties struct {
	CreateDate         string `json:"createdate"`
	HsLastModifiedDate string `json:"hs_lastmodifieddate"`
	HsTaskBody         string `json:"hs_task_body"`
	HsTaskPriority     string `json:"hs_task_priority"`
	HsTaskStatus       string `json:"hs_task_status"`
	HsTaskSubject      string `json:"hs_task_subject"`
	HsTimestamp        string `json:"hs_timestamp"`
	HubSpotOwnerId     string `json:"hubspot_owner_id"`
}

type TaskCreateOrUpdateOptions struct {
	Properties TaskCreateOrUpdateProperties `json:"properties"`
}

type TaskCreateOrUpdateProperties struct {
	HsTaskBody         string `json:"hs_task_body"`
	HsTaskPriority     string `json:"hs_task_priority"`
	HsTaskStatus       string `json:"hs_task_status"`
	HsTaskSubject      string `json:"hs_task_subject"`
	HsTimestamp        string `json:"hs_timestamp"`
	HubSpotOwnerId     string `json:"hubspot_owner_id"`
}

type TaskReadQuery struct {
	ReadQuery
}

type TaskUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}

type TaskBatchOutput struct {
	Status      string `json:"status"`
	Results     []Task `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type TaskBatchReadOptions struct {
	BatchReadOptions
}

type TaskBatchCreateOptions struct {
	Inputs []TaskCreateOrUpdateOptions `json:"inputs"`
}

type TaskBatchUpdateOptions struct {
	Inputs []TaskBatchUpdateProperties `json:"inputs"`
}

type TaskBatchUpdateProperties struct {
	Id         string                       `json:"id"`
	Properties TaskCreateOrUpdateProperties `json:"properties"`
}

type TaskSearchOptions struct {
	SearchOptions
}

type TaskSearchResults struct {
	Total      int64  `json:"total"`
	Results    []Task `json:"results"`
	Pagination
}

type TaskMergeOptions struct {
	MergeOptions
}

func (z *tasks) ListAssociations(query *TaskAssociationsQuery, taskId string, toObjectType string) (*TaskAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/%s/associations/%s", taskId, toObjectType)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.ListAssociations(): newHttpRequest(): %v", err)
	}

	ca := &TaskAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.ListAssociations(): do(): %v", err)
	}
	
	return ca, nil
}

func (z *tasks) Associate(taskId string, toObjectType string, toObjectId string, associationType string) (*Task, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/%s/associations/%s/%s/%s", taskId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Associate(): newHttpRequest(): %v", err)
	}

	task := &Task{}

	err = z.client.do(req, task)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Associate(): do(): %v", err)
	}
	
	return task, nil
}

func (z *tasks) Disassociate(taskId string, toObjectType string, toObjectId string, associationType string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/%s/associations/%s/%s/%s", taskId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.tasks.Disassociate(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *tasks) List(query *TaskListQuery) (*TaskList, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks")
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.List(): newHttpRequest(): %v", err)
	}

	tl := &TaskList{}

	err = z.client.do(req, tl)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.List(): do(): %v", err)
	}
	
	return tl, nil
}

func (z *tasks) Create(options *TaskCreateOrUpdateOptions) (*Task, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Create(): newHttpRequest(): %v", err)
	}

	task := &Task{}

	err = z.client.do(req, task)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Create(): do(): %v", err)
	}
	
	return task, nil
}

func (z *tasks) Read(query *TaskReadQuery, taskId string) (*Task, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", taskId)
	req, err := z.client.newHttpRequest("GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Read(): newHttpRequest(): %v", err)
	}

	task := &Task{}

	err = z.client.do(req, task)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Read(): do(): %+v", err)
	}

	return task, nil
}

func (z *tasks) Update(options *TaskCreateOrUpdateOptions, taskId string) (*Task, error) {
	u := fmt.Sprintf("crm/v3/objects/tasks/%s", taskId)
	req, err := z.client.newHttpRequest("PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Update(): newHttpRequest(): %v", err)
	}

	task := &Task{}

	err = z.client.do(req, task)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Update(): do(): %+v", err)
	}

	return task, nil
}

func (z *tasks) Archive(taskId string) (error) {
	u := fmt.Sprintf("crm/v3/objects/tasks/%s", taskId)
	req, err := z.client.newHttpRequest("DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.tasks.Archive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *tasks) BatchArchive(taskIds []string) (error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/batch/archive")

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, taskId := range taskIds{
		options.Inputs = append(options.Inputs, BatchInput{Id: taskId})
	}

	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return fmt.Errorf("client.tasks.BatchArchive(): newHttpRequest(): %v", err)
	}

	return z.client.do(req, nil)
}

func (z *tasks) BatchCreate(options *TaskBatchCreateOptions) (*TaskBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/batch/create")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.BatchCreate(): newHttpRequest(): %v", err)
	}

	tasks := &TaskBatchOutput{}

	err = z.client.do(req, tasks)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.BatchCreate(): do(): %+v", err)
	}

	return tasks, nil
}

func (z *tasks) BatchRead(options *TaskBatchReadOptions) (*TaskBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/batch/read")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.BatchUpdate(): newHttpRequest(): %v", err)
	}

	tasks := &TaskBatchOutput{}

	err = z.client.do(req, tasks)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.BatchUpdate(): do(): %+v", err)
	}

	return tasks, nil
}

func (z *tasks) BatchUpdate(options *TaskBatchUpdateOptions) (*TaskBatchOutput, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/batch/update")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.BatchUpdate(): newHttpRequest(): %v", err)
	}

	tasks := &TaskBatchOutput{}

	err = z.client.do(req, tasks)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.BatchUpdate(): do(): %+v", err)
	}

	return tasks, nil
}

func (z *tasks) Search(options *TaskSearchOptions) (*TaskSearchResults, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/search")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Search(): newHttpRequest(): %v", err)
	}

	tasks := &TaskSearchResults{}

	err = z.client.do(req, tasks)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Search(): do(): %+v", err)
	}

	return tasks, nil
}

func (z *tasks) Merge(options *TaskMergeOptions) (*Task, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/merge")
	req, err := z.client.newHttpRequest("POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Merge(): newHttpRequest(): %v", err)
	}

	task := &Task{}

	err = z.client.do(req, task)
	if err != nil {
		return nil, fmt.Errorf("client.tasks.Merge(): do(): %+v", err)
	}

	return task, nil
}