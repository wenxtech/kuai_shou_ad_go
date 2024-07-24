/*
磁力引擎 copy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// V1ReportCampaignReportResponseDataDetailsInner struct for V1ReportCampaignReportResponseDataDetailsInner
type V1ReportCampaignReportResponseDataDetailsInner struct {
	Charge *float32 `json:"charge,omitempty"`
	Show *int64 `json:"show,omitempty"`
	Aclick *int64 `json:"aclick,omitempty"`
	Bclick *int64 `json:"bclick,omitempty"`
	Share *int64 `json:"share,omitempty"`
	Comment *int64 `json:"comment,omitempty"`
	Like *int64 `json:"like,omitempty"`
	Follow *int64 `json:"follow,omitempty"`
	Report *int64 `json:"report,omitempty"`
	Block *int64 `json:"block,omitempty"`
	Negative *int64 `json:"negative,omitempty"`
	Activation *int64 `json:"activation,omitempty"`
	Submit *int64 `json:"submit,omitempty"`
	AdPhotoPlayed10s *int64 `json:"ad_photo_played_10s,omitempty"`
	AdPhotoPlayed2s *int64 `json:"ad_photo_played_2s,omitempty"`
	AdPhotoPlayed75percent *int64 `json:"ad_photo_played_75percent,omitempty"`
	CancelLike *int64 `json:"cancel_like,omitempty"`
	ClickConversionRatio *int64 `json:"click_conversion_ratio,omitempty"`
	ConversionCost *int64 `json:"conversion_cost,omitempty"`
	ConversionCostByImpression7d *float32 `json:"conversion_cost_by_impression_7d,omitempty"`
	ConversionNum *int64 `json:"conversion_num,omitempty"`
	ConversionNumByImpression7d *int64 `json:"conversion_num_by_impression_7d,omitempty"`
	ConversionNumCost *float32 `json:"conversion_num_cost,omitempty"`
	ConversionRatio *int64 `json:"conversion_ratio,omitempty"`
	ConversionRatioByImpression7d *float32 `json:"conversion_ratio_by_impression_7d,omitempty"`
	DeepConversionCost *int64 `json:"deep_conversion_cost,omitempty"`
	DeepConversionCostByImpression7d *int64 `json:"deep_conversion_cost_by_impression_7d,omitempty"`
	DeepConversionNum *int64 `json:"deep_conversion_num,omitempty"`
	DeepConversionNumByImpression7d *int64 `json:"deep_conversion_num_by_impression_7d,omitempty"`
	DeepConversionRatio *int64 `json:"deep_conversion_ratio,omitempty"`
	DeepConversionRatioByImpression7d *int64 `json:"deep_conversion_ratio_by_impression_7d,omitempty"`
	DownloadCompletedCost *int64 `json:"download_completed_cost,omitempty"`
	DownloadCompletedRatio *int64 `json:"download_completed_ratio,omitempty"`
	DownloadConversionRatio *int64 `json:"download_conversion_ratio,omitempty"`
	DownloadInstalled *int64 `json:"download_installed,omitempty"`
	DownloadStartedCost *int64 `json:"download_started_cost,omitempty"`
	DownloadStartedRatio *int64 `json:"download_started_ratio,omitempty"`
	Event24hStay *int64 `json:"event_24h_stay,omitempty"`
	Event24hStayByConversion *int64 `json:"event_24h_stay_by_conversion,omitempty"`
	Event24hStayByConversionCost *int64 `json:"event_24h_stay_by_conversion_cost,omitempty"`
	Event24hStayByConversionRatio *int64 `json:"event_24h_stay_by_conversion_ratio,omitempty"`
	Event24hStayCost *int64 `json:"event_24h_stay_cost,omitempty"`
	Event24hStayRatio *int64 `json:"event_24h_stay_ratio,omitempty"`
	EventAdWatch10Times *int64 `json:"event_ad_watch_10_times,omitempty"`
	EventAdWatch10TimesCost *int64 `json:"event_ad_watch_10_times_cost,omitempty"`
	EventAdWatch10TimesRatio *int64 `json:"event_ad_watch_10_times_ratio,omitempty"`
	EventAdWatch20Times *int64 `json:"event_ad_watch_20_times,omitempty"`
	EventAdWatch20TimesCost *int64 `json:"event_ad_watch_20_times_cost,omitempty"`
	EventAdWatch20TimesRatio *int64 `json:"event_ad_watch_20_times_ratio,omitempty"`
	EventAdWatch5Times *int64 `json:"event_ad_watch_5_times,omitempty"`
	EventAdWatch5TimesCost *int64 `json:"event_ad_watch_5_times_cost,omitempty"`
	EventAdWatch5TimesRatio *int64 `json:"event_ad_watch_5_times_ratio,omitempty"`
	EventAudition *int64 `json:"event_audition,omitempty"`
	EventConsultationValidRetained *int64 `json:"event_consultation_valid_retained,omitempty"`
	EventConsultationValidRetainedCost *int64 `json:"event_consultation_valid_retained_cost,omitempty"`
	EventConsultationValidRetainedRatio *int64 `json:"event_consultation_valid_retained_ratio,omitempty"`
	EventConversionClickCost *int64 `json:"event_conversion_click_cost,omitempty"`
	EventConversionClickRatio *int64 `json:"event_conversion_click_ratio,omitempty"`
	EventCreditGrantFirstDayApp *int64 `json:"event_credit_grant_first_day_app,omitempty"`
	EventCreditGrantFirstDayAppCost *int64 `json:"event_credit_grant_first_day_app_cost,omitempty"`
	EventCreditGrantFirstDayAppRatio *int64 `json:"event_credit_grant_first_day_app_ratio,omitempty"`
	EventCreditGrantFirstDayLandingPage *int64 `json:"event_credit_grant_first_day_landing_page,omitempty"`
	EventCreditGrantFirstDayLandingPageCost *int64 `json:"event_credit_grant_first_day_landing_page_cost,omitempty"`
	EventCreditGrantFirstDayLandingPageRatio *int64 `json:"event_credit_grant_first_day_landing_page_ratio,omitempty"`
	EventFiveDayStayByConversion *int64 `json:"event_five_day_stay_by_conversion,omitempty"`
	EventFiveDayStayByConversionCost *int64 `json:"event_five_day_stay_by_conversion_cost,omitempty"`
	EventFiveDayStayByConversionRatio *int64 `json:"event_five_day_stay_by_conversion_ratio,omitempty"`
	EventFourDayStayByConversion *int64 `json:"event_four_day_stay_by_conversion,omitempty"`
	EventFourDayStayByConversionCost *int64 `json:"event_four_day_stay_by_conversion_cost,omitempty"`
	EventFourDayStayByConversionRatio *int64 `json:"event_four_day_stay_by_conversion_ratio,omitempty"`
	EventMakingCalls *int64 `json:"event_making_calls,omitempty"`
	EventMakingCallsCost *int64 `json:"event_making_calls_cost,omitempty"`
	EventMakingCallsRatio *int64 `json:"event_making_calls_ratio,omitempty"`
	EventPayPurchaseAmountOneDay *int64 `json:"event_pay_purchase_amount_one_day,omitempty"`
	EventPayPurchaseAmountOneDayByConversion *int64 `json:"event_pay_purchase_amount_one_day_by_conversion,omitempty"`
	EventPayPurchaseAmountOneDayByConversionRoi *int64 `json:"event_pay_purchase_amount_one_day_by_conversion_roi,omitempty"`
	EventPayPurchaseAmountOneDayRoi *int64 `json:"event_pay_purchase_amount_one_day_roi,omitempty"`
	EventPayPurchaseAmountThreeDayByConversion *int64 `json:"event_pay_purchase_amount_three_day_by_conversion,omitempty"`
	EventPayPurchaseAmountThreeDayByConversionRoi *int64 `json:"event_pay_purchase_amount_three_day_by_conversion_roi,omitempty"`
	EventPayPurchaseAmountWeekByConversion *int64 `json:"event_pay_purchase_amount_week_by_conversion,omitempty"`
	EventPayPurchaseAmountWeekByConversionRoi *int64 `json:"event_pay_purchase_amount_week_by_conversion_roi,omitempty"`
	EventPayWeekByConversion *int64 `json:"event_pay_week_by_conversion,omitempty"`
	EventPayWeekByConversionCost *int64 `json:"event_pay_week_by_conversion_cost,omitempty"`
	EventPayWeightedPurchaseAmount *int64 `json:"event_pay_weighted_purchase_amount,omitempty"`
	EventPayWeightedPurchaseAmountFirstDay *int64 `json:"event_pay_weighted_purchase_amount_first_day,omitempty"`
	EventPreComponentConsultationValidRetained *int64 `json:"event_pre_component_consultation_valid_retained,omitempty"`
	EventSixDayStayByConversion *int64 `json:"event_six_day_stay_by_conversion,omitempty"`
	EventSixDayStayByConversionCost *int64 `json:"event_six_day_stay_by_conversion_cost,omitempty"`
	EventSixDayStayByConversionRatio *int64 `json:"event_six_day_stay_by_conversion_ratio,omitempty"`
	EventThreeDayStayByConversion *int64 `json:"event_three_day_stay_by_conversion,omitempty"`
	EventThreeDayStayByConversionCost *int64 `json:"event_three_day_stay_by_conversion_cost,omitempty"`
	EventThreeDayStayByConversionRatio *int64 `json:"event_three_day_stay_by_conversion_ratio,omitempty"`
	EventTwoDayStayByConversion *int64 `json:"event_two_day_stay_by_conversion,omitempty"`
	EventTwoDayStayByConversionCost *int64 `json:"event_two_day_stay_by_conversion_cost,omitempty"`
	EventTwoDayStayByConversionRatio *int64 `json:"event_two_day_stay_by_conversion_ratio,omitempty"`
	EventWechatQrCodeLinkClick *int64 `json:"event_wechat_qr_code_link_click,omitempty"`
	EventWeekStayByConversion *int64 `json:"event_week_stay_by_conversion,omitempty"`
	EventWeekStayByConversionCost *int64 `json:"event_week_stay_by_conversion_cost,omitempty"`
	EventWeekStayByConversionRatio *int64 `json:"event_week_stay_by_conversion_ratio,omitempty"`
	LiveEventGoodsView *int64 `json:"live_event_goods_view,omitempty"`
	LivePlayed3s *int64 `json:"live_played_3s,omitempty"`
	PlayedEnd *int64 `json:"played_end,omitempty"`
	PlayedFiveSeconds *int64 `json:"played_five_seconds,omitempty"`
	PlayedNum *int64 `json:"played_num,omitempty"`
	PlayedThreeSeconds *int64 `json:"played_three_seconds,omitempty"`
	AdScene *string `json:"adScene,omitempty"`
	AdScene *string `json:"ad_scene,omitempty"`
	PlacementType *string `json:"placement_type,omitempty"`
	CancelFollow *int64 `json:"cancel_follow,omitempty"`
	DownloadStarted *int64 `json:"download_started,omitempty"`
	DownloadCompleted *int64 `json:"download_completed,omitempty"`
	StatDate *string `json:"stat_date,omitempty"`
	StatHour *int64 `json:"stat_hour,omitempty"`
	PhotoClick *int64 `json:"photo_click,omitempty"`
	PhotoClickRatio *float32 `json:"photo_click_ratio,omitempty"`
	ActionRatio *float32 `json:"action_ratio,omitempty"`
	Impression1kCost *float32 `json:"impression_1k_cost,omitempty"`
	PhotoClickCost *float32 `json:"photo_click_cost,omitempty"`
	Click1kCost *float32 `json:"click_1k_cost,omitempty"`
	ActionCost *float32 `json:"action_cost,omitempty"`
	EventPayFirstDay *int64 `json:"event_pay_first_day,omitempty"`
	EventPayPurchaseAmountFirstDay *int64 `json:"event_pay_purchase_amount_first_day,omitempty"`
	EventPayFirstDayRoi *int64 `json:"event_pay_first_day_roi,omitempty"`
	EventPay *int64 `json:"event_pay,omitempty"`
	EventPayPurchaseAmount *int64 `json:"event_pay_purchase_amount,omitempty"`
	EventPayPurchaseAmount30DayByConversion *int64 `json:"event_pay_purchase_amount_30_day_by_conversion,omitempty"`
	EventPayPurchaseAmount30DayByConversionRoi *int64 `json:"event_pay_purchase_amount_30_day_by_conversion_roi,omitempty"`
	EventPayRoi *int64 `json:"event_pay_roi,omitempty"`
	EventRegister *int64 `json:"event_register,omitempty"`
	EventRegisterCost *int64 `json:"event_register_cost,omitempty"`
	EventRegisterRatio *int64 `json:"event_register_ratio,omitempty"`
	EventJinJianApp *int64 `json:"event_jin_jian_app,omitempty"`
	EventJinJianAppCost *int64 `json:"event_jin_jian_app_cost,omitempty"`
	EventCreditGrantApp *int64 `json:"event_credit_grant_app,omitempty"`
	EventCreditGrantAppCost *int64 `json:"event_credit_grant_app_cost,omitempty"`
	EventCreditGrantAppRatio *int64 `json:"event_credit_grant_app_ratio,omitempty"`
	EventOrderPaid *int64 `json:"event_order_paid,omitempty"`
	EventOrderPaidPurchaseAmount *int64 `json:"event_order_paid_purchase_amount,omitempty"`
	EventOrderPaidCost *int64 `json:"event_order_paid_cost,omitempty"`
	FormCount *int64 `json:"form_count,omitempty"`
	FormCost *float32 `json:"form_cost,omitempty"`
	FormActionRatio *float32 `json:"form_action_ratio,omitempty"`
	EventJinJianLandingPage *int64 `json:"event_jin_jian_landing_page,omitempty"`
	EventJinJianLandingPageCost *int64 `json:"event_jin_jian_landing_page_cost,omitempty"`
	EventCreditGrantLandingPage *int64 `json:"event_credit_grant_landing_page,omitempty"`
	EventCreditGrantLandingPageCost *int64 `json:"event_credit_grant_landing_page_cost,omitempty"`
	EventCreditGrantLandingRatio *int64 `json:"event_credit_grant_landing_ratio,omitempty"`
	EventNextDayStayCost *int64 `json:"event_next_day_stay_cost,omitempty"`
	EventNextDayStayRatio *int64 `json:"event_next_day_stay_ratio,omitempty"`
	EventNextDayStay *int64 `json:"event_next_day_stay,omitempty"`
	Play3sRatio *float32 `json:"play_3s_ratio,omitempty"`
	EventValidClues *int64 `json:"event_valid_clues,omitempty"`
	EventValidCluesCost *int64 `json:"event_valid_clues_cost,omitempty"`
	AdProductCnt *int64 `json:"ad_product_cnt,omitempty"`
	EventGoodsView *int64 `json:"event_goods_view,omitempty"`
	MerchantRecoFans *int64 `json:"merchant_reco_fans,omitempty"`
	EventOrderAmountRoi *int64 `json:"event_order_amount_roi,omitempty"`
	EventGoodsViewCost *int64 `json:"event_goods_view_cost,omitempty"`
	MerchantRecoFansCost *float32 `json:"merchant_reco_fans_cost,omitempty"`
	EventNewUserPay *int64 `json:"event_new_user_pay,omitempty"`
	EventNewUserPayCost *int64 `json:"event_new_user_pay_cost,omitempty"`
	EventNewUserPayRatio *int64 `json:"event_new_user_pay_ratio,omitempty"`
	EventButtonClick *int64 `json:"event_button_click,omitempty"`
	EventButtonClickCost *int64 `json:"event_button_click_cost,omitempty"`
	EventButtonClickRatio *int64 `json:"event_button_click_ratio,omitempty"`
	Play5sRatio *float32 `json:"play_5s_ratio,omitempty"`
	PlayEndRatio *float32 `json:"play_end_ratio,omitempty"`
	EventOrderPaidRoi *int64 `json:"event_order_paid_roi,omitempty"`
	EventNewUserJinjianApp *int64 `json:"event_new_user_jinjian_app,omitempty"`
	EventNewUserJinjianAppCost *int64 `json:"event_new_user_jinjian_app_cost,omitempty"`
	EventNewUserJinjianAppRoi *int64 `json:"event_new_user_jinjian_app_roi,omitempty"`
	EventNewUserCreditGrantApp *int64 `json:"event_new_user_credit_grant_app,omitempty"`
	EventNewUserCreditGrantAppCost *int64 `json:"event_new_user_credit_grant_app_cost,omitempty"`
	EventNewUserCreditGrantAppRoi *int64 `json:"event_new_user_credit_grant_app_roi,omitempty"`
	EventNewUserJinjianPage *int64 `json:"event_new_user_jinjian_page,omitempty"`
	EventNewUserJinjianPageCost *int64 `json:"event_new_user_jinjian_page_cost,omitempty"`
	EventNewUserJinjianPageRoi *int64 `json:"event_new_user_jinjian_page_roi,omitempty"`
	EventNewUserCreditGrantPage *int64 `json:"event_new_user_credit_grant_page,omitempty"`
	EventNewUserCreditGrantPageCost *int64 `json:"event_new_user_credit_grant_page_cost,omitempty"`
	EventNewUserCreditGrantPageRoi *int64 `json:"event_new_user_credit_grant_page_roi,omitempty"`
	EventAppointForm *int64 `json:"event_appoint_form,omitempty"`
	EventAppointFormCost *int64 `json:"event_appoint_form_cost,omitempty"`
	EventAppointFormRatio *int64 `json:"event_appoint_form_ratio,omitempty"`
	EventAppointJumpClick *int64 `json:"event_appoint_jump_click,omitempty"`
	EventAppointJumpClickCost *int64 `json:"event_appoint_jump_click_cost,omitempty"`
	EventAppointJumpClickRatio *int64 `json:"event_appoint_jump_click_ratio,omitempty"`
	EventDspGiftForm *int64 `json:"event_dsp_gift_form,omitempty"`
	EventAppInvoked *int64 `json:"event_app_invoked,omitempty"`
	EventAppInvokedCost *int64 `json:"event_app_invoked_cost,omitempty"`
	EventAppInvokedRatio *int64 `json:"event_app_invoked_ratio,omitempty"`
	EventAddWechat *int64 `json:"event_add_wechat,omitempty"`
	EventAddWechatCost *int64 `json:"event_add_wechat_cost,omitempty"`
	EventAddWechatRatio *int64 `json:"event_add_wechat_ratio,omitempty"`
	EventMultiConversion *int64 `json:"event_multi_conversion,omitempty"`
	EventMultiConversionRatio *int64 `json:"event_multi_conversion_ratio,omitempty"`
	EventMultiConversionCost *int64 `json:"event_multi_conversion_cost,omitempty"`
	EventWatchAppAd *int64 `json:"event_watch_app_ad,omitempty"`
	EventAdWatchTimes *int64 `json:"event_ad_watch_times,omitempty"`
	EventAdWatchTimesRatio *int64 `json:"event_ad_watch_times_ratio,omitempty"`
	EventAdWatchTimesCost *int64 `json:"event_ad_watch_times_cost,omitempty"`
	EventAddShoppingCart *int64 `json:"event_add_shopping_cart,omitempty"`
	EventAddShoppingCartCost *int64 `json:"event_add_shopping_cart_cost,omitempty"`
	EventGetThrough *int64 `json:"event_get_through,omitempty"`
	EventGetThroughCost *int64 `json:"event_get_through_cost,omitempty"`
	EventGetThroughRatio *int64 `json:"event_get_through_ratio,omitempty"`
	AdPhotoPlayed75percentRatio *float32 `json:"ad_photo_played_75percent_ratio,omitempty"`
	AdPhotoPlayed10sRatio *float32 `json:"ad_photo_played_10s_ratio,omitempty"`
	AdPhotoPlayed2sRatio *float32 `json:"ad_photo_played_2s_ratio,omitempty"`
	EventPhoneGetThrough *int64 `json:"event_phone_get_through,omitempty"`
	EventIntentionConfirmed *int64 `json:"event_intention_confirmed,omitempty"`
	EventWechatConnected *int64 `json:"event_wechat_connected,omitempty"`
	EventOrderSuccessed *int64 `json:"event_order_successed,omitempty"`
	EventPhoneCardActivate *int64 `json:"event_phone_card_activate,omitempty"`
	EventMeasurementHouse *int64 `json:"event_measurement_house,omitempty"`
	EventNextDayStayNew *int64 `json:"event_next_day_stay_new,omitempty"`
	EventNextDayStayNewCost *int64 `json:"event_next_day_stay_new_cost,omitempty"`
	EventNextDayStayNewRatio *int64 `json:"event_next_day_stay_new_ratio,omitempty"`
	AdShow *int64 `json:"ad_show,omitempty"`
	ActionNewRatio *float32 `json:"action_new_ratio,omitempty"`
	EventOutboundCall *int64 `json:"event_outbound_call,omitempty"`
	EventOutboundCallCost *int64 `json:"event_outbound_call_cost,omitempty"`
	EventOutboundCallRatio *int64 `json:"event_outbound_call_ratio,omitempty"`
	KeyAction *int64 `json:"key_action,omitempty"`
	KeyActionCost *int64 `json:"key_action_cost,omitempty"`
	KeyActionRatio *int64 `json:"key_action_ratio,omitempty"`
	EventCreditCardRecheck *int64 `json:"event_credit_card_recheck,omitempty"`
	EventCreditCardRecheckFirstDay *int64 `json:"event_credit_card_recheck_first_day,omitempty"`
	EventNoIntention *int64 `json:"event_no_intention,omitempty"`
	EventMultiPaySevenDayByConversion *int64 `json:"event_multi_pay_seven_day_by_conversion,omitempty"`
	EventMultiPaySevenDayByConversionCost *int64 `json:"event_multi_pay_seven_day_by_conversion_cost,omitempty"`
	McbLeftTimePerfCpaBid *nil `json:"mcb_left_time_perf_cpa_bid,omitempty"`
	McbLeftTimePerfRoiRatio *nil `json:"mcb_left_time_perf_roi_ratio,omitempty"`
	McbAllDayPerfCpaBid *nil `json:"mcb_all_day_perf_cpa_bid,omitempty"`
	McbAllDayPerfRoiRatio *nil `json:"mcb_all_day_perf_roi_ratio,omitempty"`
	LiveRoomAvgPlayedSeconds *int64 `json:"live_room_avg_played_seconds,omitempty"`
	AdLiveShare *int64 `json:"ad_live_share,omitempty"`
	AdLiveComment *int64 `json:"ad_live_comment,omitempty"`
	LivePlayedStarted *int64 `json:"live_played_started,omitempty"`
	LivePlayedStartedCost *int64 `json:"live_played_started_cost,omitempty"`
	AdLiveFollow *int64 `json:"ad_live_follow,omitempty"`
	AdLiveFollowCost *int64 `json:"ad_live_follow_cost,omitempty"`
	SimpleLivePlayedStarted *int64 `json:"simple_live_played_started,omitempty"`
	StandardLivePlayedStarted *int64 `json:"standard_live_played_started,omitempty"`
	ConversionComponentImpression *int64 `json:"conversion_component_impression,omitempty"`
	ConversionComponentClick *int64 `json:"conversion_component_click,omitempty"`
	ConversionComponentRate *int64 `json:"conversion_component_rate,omitempty"`
	AdLandingPageImpression *int64 `json:"ad_landing_page_impression,omitempty"`
	AdAppDownloadHalfImpression *int64 `json:"ad_app_download_half_impression,omitempty"`
	EventDrawCreditLine *int64 `json:"event_draw_credit_line,omitempty"`
	EventActive *int64 `json:"event_active,omitempty"`
	EventOrderSubmitCost *float32 `json:"event_order_submit_cost,omitempty"`
	EventOrderSubmit *int64 `json:"event_order_submit,omitempty"`
	DirectSubmit1dCost *float32 `json:"direct_submit_1d_cost,omitempty"`
	OrderSubmitAmt *int64 `json:"order_submit_amt,omitempty"`
	OrderSubmitRoi *int64 `json:"order_submit_roi,omitempty"`
	EventWeekTotalStayByConversion *int64 `json:"event_week_total_stay_by_conversion,omitempty"`
	EventWeekTotalStayNew *int64 `json:"event_week_total_stay_new,omitempty"`
	EventWeekTotalStayByConversionCost *int64 `json:"event_week_total_stay_by_conversion_cost,omitempty"`
	EventWeekTotalStayNewCost *int64 `json:"event_week_total_stay_new_cost,omitempty"`
	EventWeekTotalStayByConversionRatio *int64 `json:"event_week_total_stay_by_conversion_ratio,omitempty"`
	EventWeekTotalStayNewRatio *int64 `json:"event_week_total_stay_new_ratio,omitempty"`
	RelevanceScoreAvg *nil `json:"relevance_score_avg,omitempty"`
	MinigameIaaPurchaseAmountFirstDay *int64 `json:"minigame_iaa_purchase_amount_first_day,omitempty"`
	MinigameIaaPurchaseAmountThreeDayByConversion *int64 `json:"minigame_iaa_purchase_amount_three_day_by_conversion,omitempty"`
	MinigameIaaPurchaseAmountWeekByConversion *int64 `json:"minigame_iaa_purchase_amount_week_by_conversion,omitempty"`
	MinigameIaaPurchaseAmountFirstDayRoi *int64 `json:"minigame_iaa_purchase_amount_first_day_roi,omitempty"`
	MinigameIaaPurchaseAmountThreeDayByConversionRoi *int64 `json:"minigame_iaa_purchase_amount_three_day_by_conversion_roi,omitempty"`
	MinigameIaaPurchaseAmountWeekByConversionRoi *int64 `json:"minigame_iaa_purchase_amount_week_by_conversion_roi,omitempty"`
	MinigameIaaPurchaseAmount *int64 `json:"minigame_iaa_purchase_amount,omitempty"`
	MinigameIaaPurchaseRoi *int64 `json:"minigame_iaa_purchase_roi,omitempty"`
	T7PaiedAmt *int64 `json:"t7_paied_amt,omitempty"`
	EventEffectiveCustomerAcquisitionCnt *int64 `json:"event_effective_customer_acquisition_cnt,omitempty"`
	EventEffectiveCustomerAcquisitionCost *int64 `json:"event_effective_customer_acquisition_cost,omitempty"`
	EventEffectiveCustomerAcquisitionRatio *int64 `json:"event_effective_customer_acquisition_ratio,omitempty"`
	SupplementConversionCnt *int64 `json:"supplement_conversion_cnt,omitempty"`
	SimpleLiveRoomPlayedSeconds *int64 `json:"simple_live_room_played_seconds,omitempty"`
	StandardLiveRoomPlayedSeconds *int64 `json:"standard_live_room_played_seconds,omitempty"`
	Likes *int64 `json:"likes,omitempty"`
	CampaignId *int64 `json:"campaign_id,omitempty"`
	CampaignName *string `json:"campaign_name,omitempty"`
	Status *int64 `json:"status,omitempty"`
	BackflowCvLower *nil `json:"backflow_cv_lower,omitempty"`
	BackflowCvUpper *nil `json:"backflow_cv_upper,omitempty"`
	ConversionCostLower *nil `json:"conversion_cost_lower,omitempty"`
	ConversionCostUpper *nil `json:"conversion_cost_upper,omitempty"`
	BackflowPayment *nil `json:"backflow_payment,omitempty"`
	BackflowRoi *nil `json:"backflow_roi,omitempty"`
	BackflowSevenPayment *nil `json:"backflow_seven_payment,omitempty"`
	BackflowSevenRoi *nil `json:"backflow_seven_roi,omitempty"`
}


