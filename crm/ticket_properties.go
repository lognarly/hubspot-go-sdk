package crm

type TicketProperties struct {
	ClosedDate                            string `json:"closed_date,omitempty"`
	CreatedBy                             string `json:"created_by,omitempty"`
	Createdate                            string `json:"createdate,omitempty"`
	FirstAgentReplyDate                   string `json:"first_agent_reply_date,omitempty"`
	HsAllAssignedBusinessUnitIds          string `json:"hs_all_assigned_business_unit_ids,omitempty"`
	HsAutoGeneratedFromThreadId           string `json:"hs_auto_generated_from_thread_id,omitempty"`
	HsConversationsOriginatingMessageId   string `json:"hs_conversations_originating_message_id,omitempty"`
	HsConversationsOriginatingThreadId    string `json:"hs_conversations_originating_thread_id,omitempty"`
	HsCreatedByUserId                     string `json:"hs_created_by_user_id,omitempty"`
	HsCreatedate                          string `json:"hs_createdate,omitempty"`
	HsCustomInbox                         string `json:"hs_custom_inbox,omitempty"`
	HsExternalObjectIds                   string `json:"hs_external_object_ids,omitempty"`
	HsFeedbackLastCesFollowUp             string `json:"hs_feedback_last_ces_follow_up,omitempty"`
	HsFeedbackLastCesRating               string `json:"hs_feedback_last_ces_rating,omitempty"`
	HsFeedbackLastSurveyDate              string `json:"hs_feedback_last_survey_date,omitempty"`
	HsFileUpload                          string `json:"hs_file_upload,omitempty"`
	HsFirstAgentMessageSentAt             string `json:"hs_first_agent_message_sent_at,omitempty"`
	HsLastEmailActivity                   string `json:"hs_last_email_activity,omitempty"`
	HsLastEmailDate                       string `json:"hs_last_email_date,omitempty"`
	HsLastMessageReceivedAt               string `json:"hs_last_message_received_at,omitempty"`
	HsLastMessageSentAt                   string `json:"hs_last_message_sent_at,omitempty"`
	HsLastactivitydate                    string `json:"hs_lastactivitydate,omitempty"`
	HsLastcontacted                       string `json:"hs_lastcontacted,omitempty"`
	HsLastmodifieddate                    string `json:"hs_lastmodifieddate,omitempty"`
	HsLatestMessageSeenByAgentIds         string `json:"hs_latest_message_seen_by_agent_ids,omitempty"`
	HsMergedObjectIds                     string `json:"hs_merged_object_ids,omitempty"`
	HsMsteamsMessageId                    string `json:"hs_msteams_message_id,omitempty"`
	HsNextactivitydate                    string `json:"hs_nextactivitydate,omitempty"`
	HsNumAssociatedCompanies              int64  `json:"hs_num_associated_companies,omitempty"`
	HsNumTimesContacted                   int64  `json:"hs_num_times_contacted,omitempty"`
	HsObjectId                            string `json:"hs_object_id,omitempty"`
	HsOriginatingEmailEngagementId        int64  `json:"hs_originating_email_engagement_id,omitempty"`
	HsPipeline                            string `json:"hs_pipeline,omitempty"`
	HsPipelineStage                       string `json:"hs_pipeline_stage,omitempty"`
	HsResolution                          string `json:"hs_resolution,omitempty"`
	HsThreadIdsToRestore                  string `json:"hs_thread_ids_to_restore,omitempty"`
	HsTicketCategory                      string `json:"hs_ticket_category,omitempty"`
	HsTicketId                            string `json:"hs_ticket_id,omitempty"`
	HsTicketPriority                      string `json:"hs_ticket_priority,omitempty"`
	HsTimeToCloseSlaAt                    string `json:"hs_time_to_close_sla_at,omitempty"`
	HsTimeToCloseSlaStatus                string `json:"hs_time_to_close_sla_status,omitempty"`
	HsTimeToFirstResponseSlaAt            string `json:"hs_time_to_first_response_sla_at,omitempty"`
	HsTimeToFirstResponseSlaStatus        string `json:"hs_time_to_first_response_sla_status,omitempty"`
	HsUniqueCreationKey                   string `json:"hs_unique_creation_key,omitempty"`
	HsUpdatedByUserId                     string `json:"hs_updated_by_user_id,omitempty"`
	HsUserIdsOfAllNotificationFollowers   string `json:"hs_user_ids_of_all_notification_followers,omitempty"`
	HsUserIdsOfAllNotificationUnfollowers string `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`
	HsUserIdsOfAllOwners                  string `json:"hs_user_ids_of_all_owners,omitempty"`
	HubspotOwnerAssigneddate              string `json:"hubspot_owner_assigneddate,omitempty"`
	LastEngagementDate                    string `json:"last_engagement_date,omitempty"`
	LastReplyDate                         string `json:"last_reply_date,omitempty"`
	NpsFollowUpAnswer                     string `json:"nps_follow_up_answer,omitempty"`
	NpsFollowUpQuestionVersion            string `json:"nps_follow_up_question_version,omitempty"`
	NpsScore                              string `json:"nps_score,omitempty"`
	SourceThreadId                        string `json:"source_thread_id,omitempty"`
	TimeToClose                           int64  `json:"time_to_close,omitempty"`
	TimeToFirstAgentReply                 int64  `json:"time_to_first_agent_reply,omitempty"`
	Subject                               string `json:"subject,omitempty"`
	Content                               string `json:"content,omitempty"`
	SourceType                            string `json:"source_type,omitempty"`
	SourceRef                             string `json:"source_ref,omitempty"`
	Tags                                  string `json:"tags,omitempty"`
	HsSalesEmailLastReplied               string `json:"hs_sales_email_last_replied,omitempty"`
	HubspotOwnerId                        string `json:"hubspot_owner_id,omitempty"`
	NotesLastContacted                    string `json:"notes_last_contacted,omitempty"`
	NotesLastUpdated                      string `json:"notes_last_updated,omitempty"`
	NotesNextActivityDate                 string `json:"notes_next_activity_date,omitempty"`
	NumContactedNotes                     int64  `json:"num_contacted_notes,omitempty"`
	NumNotes                              int64  `json:"num_notes,omitempty"`
	HubspotTeamId                         string `json:"hubspot_team_id,omitempty"`
	HsAllOwnerIds                         string `json:"hs_all_owner_ids,omitempty"`
	HsAllTeamIds                          string `json:"hs_all_team_ids,omitempty"`
	HsAllAccessibleTeamIds                string `json:"hs_all_accessible_team_ids,omitempty"`
}