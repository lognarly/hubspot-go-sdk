package crm

type DealCalcPrefernce string

const (
	HsTcv    DealCalcPrefernce = "TCV"
	HsArr    DealCalcPrefernce = "ARR"
	HsMrr    DealCalcPrefernce = "MRR"
	HsCustom DealCalcPrefernce = "CUSTOM"
)

type DealManualForecastCategory string

const (
	HsOmit       DealManualForecastCategory = "OMIT"
	HsPipeline   DealManualForecastCategory = "PIPELINE"
	HsBestCase   DealManualForecastCategory = "BEST_CASE"
	HsCommit     DealManualForecastCategory = "COMMIT"
	HsClosed     DealManualForecastCategory = "CLOSED"
	HsClosedLost DealManualForecastCategory = "CLOSED_LOST"
)

type DealPriority string

const (
	HsLow    DealPriority = "low"
	HsMedium DealPriority = "medium"
	HsHigh   DealPriority = "high"
)

type DealType string

const (
	NewBusiness      DealType = "newbusiness"
	ExistingBusiness DealType = "existingbusiness"
)

type DealProperties struct {
	AmountInHomeCurrency                   int64                      `json:"amount_in_home_currency,omitempty"`
	CreatedByApi                           bool                       `json:"createdbyapi,omitempty"` //This is specific to an attribute used by my company
	DaysToClose                            int64                      `json:"days_to_close,omitempty"`
	DealCurrencyCode                       string                     `json:"deal_currency_code,omitempty"`
	HsAcv                                  int64                      `json:"hs_acv,omitempty"`
	HsArr                                  int64                      `json:"hs_arr,omitempty"`
	HsCampaign                             string                     `json:"hs_campaign,omitempty"`
	HsClosedAmount                         int64                      `json:"hs_closed_amount,omitempty"`
	HsClosedAmountInHomeCurrency           float64                    `json:"hs_closed_amount_in_home_currency,omitempty"`
	HsCreatedByUserId                      int64                      `json:"hs_created_by_user_id,omitempty"`
	HsDealAmountCalculationPreference      DealCalcPrefernce          `json:"hs_deal_amount_calculation_preference,omitempty"`
	HsDealStageProbability                 int64                      `json:"hs_deal_stage_probability,omitempty"`
	HsDealStageProbabilityShadow           int64                      `json:"hs_deal_stage_probability_shadow,omitempty"`
	HsForecastAmount                       int64                      `json:"hs_forecast_amount,omitempty"`
	HsForecastProbability                  int64                      `json:"hs_forecast_probability,omitempty"`
	HsIsClosed                             bool                       `json:"hs_is_closed,omitempty"`
	HsIsClosedWon                          bool                       `json:"hs_is_closed_won,omitempty"`
	HsIsDealSplit                          bool                       `json:"hs_is_deal_split,omitempty"`
	HsLastmodifieddate                     string                     `json:"hs_lastmodifieddate,omitempty"`
	HsLikelihoodToClose                    int64                      `json:"hs_likelihood_to_close,omitempty"`
	HsManualForecastCategory 			   DealManualForecastCategory `json:"hs_manual_forecast_category,omitempty"`
	HsMrr 								   int64                      `json:"hs_mrr,omitempty"`
	HsNextStep 							   string                     `json:"hs_next_step,omitempty"`
	HsNumAssociatedActiveDealRegistrations int64                      `json:"hs_num_associated_active_deal_registrations,omitempty"`
	HsNumAssociatedDealRegistrations       int64                      `json:"hs_num_associated_deal_registrations,omitempty"`
	HsNumAssociatedDealSplits              int64                      `json:"hs_num_associated_deal_splits,omitempty"`
	HsNumTargetAccounts                    int64                      `json:"hs_num_target_accounts,omitempty"`
	HsObjectId                             int64                      `json:"hs_object_id,omitempty"`
	HsPredictedAmount                      int64                      `json:"hs_predicted_amount,omitempty"`
	HsPredictedAmountInHomeCurrency        float64                    `json:"hs_predicted_amount_in_home_currency,omitempty"`
	HsPriority                             DealPriority               `json:"hs_priority,omitempty"`
	HsProjectedAmount                      float64                    `json:"hs_projected_amount,omitempty"`
	HsProjectedAmountInHomeCurrency        float64                    `json:"hs_projected_amount_in_home_currency,omitempty"`
	HsTcv                                  int64                      `json:"hs_tcv,omitempty"`
	HsUniqueCreationKey                    string                     `json:"hs_unique_creation_key,omitempty"`
	HsUpdatedByUserId                      string                     `json:"hs_updated_by_user_id,omitempty"`
	HubspotOwnerAssigneddate               string                     `json:"hubspot_owner_assigneddate,omitempty"`
	Dealname                               string                     `json:"dealname,omitempty"`
	Amount                                 int64                      `json:"amount,omitempty"`
	Dealstage                              string                     `json:"dealstage,omitempty"`
	Pipeline                               string                     `json:"pipeline,omitempty"`
	Closedate                              string                     `json:"closedate,omitempty"`
	Createdate                             string                     `json:"createdate,omitempty"`
	EngagementsLastMeetingBooked           string                     `json:"engagements_last_meeting_booked,omitempty"`
	EngagementsLastMeetingBookedCampaign   string                     `json:"engagements_last_meeting_booked_campaign,omitempty"`
	EngagementsLastMeetingBookedMedium     string                     `json:"engagements_last_meeting_booked_medium,omitempty"`
	EngagementsLastMeetingBookedSource     string                     `json:"engagements_last_meeting_booked_source,omitempty"`
	HsLatestMeetingActivity                string                     `json:"hs_latest_meeting_activity,omitempty"`
	HsSalesEmailLastReplied                string                     `json:"hs_sales_email_last_replied,omitempty"`
	HubspotOwnerId                         string                     `json:"hubspot_owner_id,omitempty"`
	NotesLastContacted                     string                     `json:"notes_last_contacted,omitempty"`
	NotesLastUpdated                       string                     `json:"notes_last_updated,omitempty"`
	NotesNextActivityDate                  string                     `json:"notes_next_activity_date,omitempty"`
	NumContactedNotes                      int64                      `json:"num_contacted_notes,omitempty"`
	NumNotes                               int64                      `json:"num_notes,omitempty"`
	HsCreatedate                           string                     `json:"hs_createdate,omitempty"`
	HubspotTeamId                          string                     `json:"hubspot_team_id,omitempty"`
	Dealtype                               DealType                   `json:"dealtype,omitempty"`
	HsAllOwnerIds                          string                     `json:"hs_all_owner_ids,omitempty"`
	Description                            string                     `json:"description,omitempty"`
	HsAllTeamIds                           string                     `json:"hs_all_team_ids,omitempty"`
	HsAllAccessibleTeamIds                 string                     `json:"hs_all_accessible_team_ids,omitempty"`
	NumAssociatedContacts                  int64                      `json:"num_associated_contacts,omitempty"`
	ClosedLostReason                       string                     `json:"closed_lost_reason,omitempty"`
	ClosedWonReason                        string                     `json:"closed_won_reason,omitempty"`
}