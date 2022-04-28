package crm
/*
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
	AmountInHomeCurrency                   int64                      `json:"amount_in_home_currency"`
	DaysToClose                            int64                      `json:"days_to_close"`
	DealCurrencyCode                       string                     `json:"deal_currency_code"`
	HsAcv                                  int64                      `json:"hs_acv"`
	HsArr                                  int64                      `json:"hs_arr"`
	HsCampaign                             string                     `json:"hs_campaign"`
	HsClosedAmount                         int64                      `json:"hs_closed_amount"`
	HsClosedAmountInHomeCurrency           float64                    `json:"hs_closed_amount_in_home_currency"`
	HsCreatedByUserId                      int64                      `json:"hs_created_by_user_id"`
	HsDealAmountCalculationPreference      DealCalcPrefernce          `json:"hs_deal_amount_calculation_preference"`
	HsDealStageProbability                 int64                      `json:"hs_deal_stage_probability"`
	HsDealStageProbabilityShadow           int64                      `json:"hs_deal_stage_probability_shadow"`
	HsForecastAmount                       int64                      `json:"hs_forecast_amount"`
	HsForecastProbability                  int64                      `json:"hs_forecast_probability"`
	HsIsClosed                             bool                       `json:"hs_is_closed"`
	HsIsClosedWon                          bool                       `json:"hs_is_closed_won"`
	HsIsDealSplit                          bool                       `json:"hs_is_deal_split"`
	HsLastmodifieddate                     string                     `json:"hs_lastmodifieddate"`
	HsLikelihoodToClose                    int64                      `json:"hs_likelihood_to_close"`
	HsManualForecastCategory 			   DealManualForecastCategory `json:"hs_manual_forecast_category"`
	HsMrr 								   int64                      `json:"hs_mrr"`
	HsNextStep 							   string                     `json:"hs_next_step"`
	HsNumAssociatedActiveDealRegistrations int64                      `json:"hs_num_associated_active_deal_registrations"`
	HsNumAssociatedDealRegistrations       int64                      `json:"hs_num_associated_deal_registrations"`
	HsNumAssociatedDealSplits              int64                      `json:"hs_num_associated_deal_splits"`
	HsNumTargetAccounts                    int64                      `json:"hs_num_target_accounts"`
	HsObjectId                             int64                      `json:"hs_object_id"`
	HsPredictedAmount                      int64                      `json:"hs_predicted_amount"`
	HsPredictedAmountInHomeCurrency        float64                    `json:"hs_predicted_amount_in_home_currency"`
	HsPriority                             DealPriority               `json:"hs_priority"`
	HsProjectedAmount                      float64                    `json:"hs_projected_amount"`
	HsProjectedAmountInHomeCurrency        float64                    `json:"hs_projected_amount_in_home_currency"`
	HsTcv                                  int64                      `json:"hs_tcv"`
	HsUniqueCreationKey                    string                     `json:"hs_unique_creation_key"`
	HsUpdatedByUserId                      string                     `json:"hs_updated_by_user_id"`
	HubspotOwnerAssigneddate               string                     `json:"hubspot_owner_assigneddate"`
	Dealname                               string                     `json:"dealname"`
	Amount                                 int64                      `json:"amount"`
	Dealstage                              string                     `json:"dealstage"`
	Pipeline                               string                     `json:"pipeline"`
	Closedate                              string                     `json:"closedate"`
	Createdate                             string                     `json:"createdate"`
	EngagementsLastMeetingBooked           string                     `json:"engagements_last_meeting_booked"`
	EngagementsLastMeetingBookedCampaign   string                     `json:"engagements_last_meeting_booked_campaign"`
	EngagementsLastMeetingBookedMedium     string                     `json:"engagements_last_meeting_booked_medium"`
	EngagementsLastMeetingBookedSource     string                     `json:"engagements_last_meeting_booked_source"`
	HsLatestMeetingActivity                string                     `json:"hs_latest_meeting_activity"`
	HsSalesEmailLastReplied                string                     `json:"hs_sales_email_last_replied"`
	HubspotOwnerId                         string                     `json:"hubspot_owner_id"`
	NotesLastContacted                     string                     `json:"notes_last_contacted"`
	NotesLastUpdated                       string                     `json:"notes_last_updated"`
	NotesNextActivityDate                  string                     `json:"notes_next_activity_date"`
	NumContactedNotes                      int64                      `json:"num_contacted_notes"`
	NumNotes                               int64                      `json:"num_notes"`
	HsCreatedate                           string                     `json:"hs_createdate"`
	HubspotTeamId                          string                     `json:"hubspot_team_id"`
	Dealtype                               DealType                   `json:"dealtype"`
	HsAllOwnerIds                          string                     `json:"hs_all_owner_ids"`
	Description                            string                     `json:"description"`
	HsAllTeamIds                           string                     `json:"hs_all_team_ids"`
	HsAllAccessibleTeamIds                 string                     `json:"hs_all_accessible_team_ids"`
	NumAssociatedContacts                  int64                      `json:"num_associated_contacts"`
	ClosedLostReason                       string                     `json:"closed_lost_reason"`
	ClosedWonReason                        string                     `json:"closed_won_reason"`
}*/