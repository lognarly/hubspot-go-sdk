package crm

type ProductProperties struct {
	Amount                                float64 `json:"amount"`
	Createdate                            string  `json:"createdate"`
	Description                           string  `json:"description"`
	Discount                              float64 `json:"discount"`
	HsAllAccessibleTeamIds                string  `json:"hs_all_accessible_team_ids"`
	HsAllAssignedBusinessUnitIds          string  `json:"hs_all_assigned_business_unit_ids"`
	HsAllOwnerIds                         string  `json:"hs_all_owner_ids"`
	HsAllTeamIds                          string  `json:"hs_all_team_ids"`
	HsAvatarFilemanagerKey                string  `json:"hs_avatar_filemanager_key"`
	HsCostOfGoodsSold                     float64 `json:"hs_cost_of_goods_sold"`
	HsCreatedByUserId                     string  `json:"hs_created_by_user_id"`
	HsCreatedate                          string  `json:"hs_createdate"`
	HsDiscountPercentage                  float64 `json:"hs_discount_percentage"`
	HsFolderId                            string  `json:"hs_folder_id"`
	HsImages                              string  `json:"hs_images"`
	HsLastmodifieddate                    string  `json:"hs_lastmodifieddate"`
	HsMergedObjectIds                     string  `json:"hs_merged_object_ids"`
	HsObjectId                            string  `json:"hs_object_id"`
	HsProductType                         string  `json:"hs_product_type"`
	HsRecurringBillingPeriod              string  `json:"hs_recurring_billing_period"`
	HsRecurringBillingStartDate           string  `json:"hs_recurring_billing_start_date"`
	HsSku                                 string  `json:"hs_sku"`
	HsUniqueCreationKey                   string  `json:"hs_unique_creation_key"`
	HsUpdatedByUserId                     string  `json:"hs_updated_by_user_id"`
	HsUrl                                 string  `json:"hs_url"`
	HsUserIdsOfAllNotificationFollowers   string  `json:"hs_user_ids_of_all_notification_followers"`
	HsUserIdsOfAllNotificationUnfollowers string  `json:"hs_user_ids_of_all_notification_unfollowers"`
	HsUserIdsOfAllOwners                  string  `json:"hs_user_ids_of_all_owners"`
	HubspotOwnerAssigneddate              string  `json:"hubspot_owner_assigneddate"`
	HubspotOwnerId                        string  `json:"hubspot_owner_id"`
	HubspotTeamId                         string  `json:"hubspot_team_id"`
	Name                                  string  `json:"name"`
	Price                                 float64 `json:"price"`
	Quantity                              int64   `json:"quantity"`
	Recurringbillingfrequency             string  `json:"recurringbillingfrequency"`
	Tax                                   float64 `json:"tax"`
}