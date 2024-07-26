/*
磁力引擎 copy

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// V1ReportMaterialReportResponseDetailsInner struct for V1ReportMaterialReportResponseDetailsInner
type V1ReportMaterialReportResponseDetailsInner struct {
	Charge *float64 `json:"charge,omitempty"`
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
	ClickConversionRatio *float64 `json:"click_conversion_ratio,omitempty"`
	ConversionCost *float64 `json:"conversion_cost,omitempty"`
	ConversionCostByImpression7d *float64 `json:"conversion_cost_by_impression_7d,omitempty"`
	ConversionNum *int64 `json:"conversion_num,omitempty"`
	ConversionNumByImpression7d *int64 `json:"conversion_num_by_impression_7d,omitempty"`
	ConversionNumCost *float64 `json:"conversion_num_cost,omitempty"`
	ConversionRatio *float64 `json:"conversion_ratio,omitempty"`
	ConversionRatioByImpression7d *float64 `json:"conversion_ratio_by_impression_7d,omitempty"`
	DeepConversionCost *float64 `json:"deep_conversion_cost,omitempty"`
	DeepConversionCostByImpression7d *float64 `json:"deep_conversion_cost_by_impression_7d,omitempty"`
	DeepConversionNum *int64 `json:"deep_conversion_num,omitempty"`
	DeepConversionNumByImpression7d *int64 `json:"deep_conversion_num_by_impression_7d,omitempty"`
	DeepConversionRatio *float64 `json:"deep_conversion_ratio,omitempty"`
	DeepConversionRatioByImpression7d *float64 `json:"deep_conversion_ratio_by_impression_7d,omitempty"`
	DownloadCompletedCost *float64 `json:"download_completed_cost,omitempty"`
	DownloadCompletedRatio *float64 `json:"download_completed_ratio,omitempty"`
	DownloadConversionRatio *float64 `json:"download_conversion_ratio,omitempty"`
	DownloadInstalled *int64 `json:"download_installed,omitempty"`
	DownloadStartedCost *float64 `json:"download_started_cost,omitempty"`
	DownloadStartedRatio *float64 `json:"download_started_ratio,omitempty"`
	Event24hStay *int64 `json:"event_24h_stay,omitempty"`
	Event24hStayByConversion *int64 `json:"event_24h_stay_by_conversion,omitempty"`
	Event24hStayByConversionCost *float64 `json:"event_24h_stay_by_conversion_cost,omitempty"`
	Event24hStayByConversionRatio *float64 `json:"event_24h_stay_by_conversion_ratio,omitempty"`
	Event24hStayCost *float64 `json:"event_24h_stay_cost,omitempty"`
	Event24hStayRatio *float64 `json:"event_24h_stay_ratio,omitempty"`
	EventAdWatch10Times *int64 `json:"event_ad_watch_10_times,omitempty"`
	EventAdWatch10TimesCost *float64 `json:"event_ad_watch_10_times_cost,omitempty"`
	EventAdWatch10TimesRatio *float64 `json:"event_ad_watch_10_times_ratio,omitempty"`
	EventAdWatch20Times *int64 `json:"event_ad_watch_20_times,omitempty"`
	EventAdWatch20TimesCost *float64 `json:"event_ad_watch_20_times_cost,omitempty"`
	EventAdWatch20TimesRatio *float64 `json:"event_ad_watch_20_times_ratio,omitempty"`
	EventAdWatch5Times *int64 `json:"event_ad_watch_5_times,omitempty"`
	EventAdWatch5TimesCost *float64 `json:"event_ad_watch_5_times_cost,omitempty"`
	EventAdWatch5TimesRatio *float64 `json:"event_ad_watch_5_times_ratio,omitempty"`
	EventAudition *int64 `json:"event_audition,omitempty"`
	EventConsultationValidRetained *int64 `json:"event_consultation_valid_retained,omitempty"`
	EventConsultationValidRetainedCost *float64 `json:"event_consultation_valid_retained_cost,omitempty"`
	EventConsultationValidRetainedRatio *float64 `json:"event_consultation_valid_retained_ratio,omitempty"`
	EventConversionClickCost *float64 `json:"event_conversion_click_cost,omitempty"`
	EventConversionClickRatio *float64 `json:"event_conversion_click_ratio,omitempty"`
	EventCreditGrantFirstDayApp *int64 `json:"event_credit_grant_first_day_app,omitempty"`
	EventCreditGrantFirstDayAppCost *float64 `json:"event_credit_grant_first_day_app_cost,omitempty"`
	EventCreditGrantFirstDayAppRatio *float64 `json:"event_credit_grant_first_day_app_ratio,omitempty"`
	EventCreditGrantFirstDayLandingPage *int64 `json:"event_credit_grant_first_day_landing_page,omitempty"`
	EventCreditGrantFirstDayLandingPageCost *float64 `json:"event_credit_grant_first_day_landing_page_cost,omitempty"`
	EventCreditGrantFirstDayLandingPageRatio *float64 `json:"event_credit_grant_first_day_landing_page_ratio,omitempty"`
	EventFiveDayStayByConversion *int64 `json:"event_five_day_stay_by_conversion,omitempty"`
	EventFiveDayStayByConversionCost *float64 `json:"event_five_day_stay_by_conversion_cost,omitempty"`
	EventFiveDayStayByConversionRatio *float64 `json:"event_five_day_stay_by_conversion_ratio,omitempty"`
	EventFourDayStayByConversion *int64 `json:"event_four_day_stay_by_conversion,omitempty"`
	EventFourDayStayByConversionCost *float64 `json:"event_four_day_stay_by_conversion_cost,omitempty"`
	EventFourDayStayByConversionRatio *float64 `json:"event_four_day_stay_by_conversion_ratio,omitempty"`
	EventMakingCalls *int64 `json:"event_making_calls,omitempty"`
	EventMakingCallsCost *float64 `json:"event_making_calls_cost,omitempty"`
	EventMakingCallsRatio *float64 `json:"event_making_calls_ratio,omitempty"`
	EventOrderSubmit *int64 `json:"event_order_submit,omitempty"`
	EventPayPurchaseAmountOneDay *float64 `json:"event_pay_purchase_amount_one_day,omitempty"`
	EventPayPurchaseAmountOneDayByConversion *float64 `json:"event_pay_purchase_amount_one_day_by_conversion,omitempty"`
	EventPayPurchaseAmountOneDayByConversionRoi *float64 `json:"event_pay_purchase_amount_one_day_by_conversion_roi,omitempty"`
	EventPayPurchaseAmountOneDayRoi *float64 `json:"event_pay_purchase_amount_one_day_roi,omitempty"`
	EventPayPurchaseAmountThreeDayByConversion *float64 `json:"event_pay_purchase_amount_three_day_by_conversion,omitempty"`
	EventPayPurchaseAmountThreeDayByConversionRoi *float64 `json:"event_pay_purchase_amount_three_day_by_conversion_roi,omitempty"`
	EventPayPurchaseAmountWeekByConversion *float64 `json:"event_pay_purchase_amount_week_by_conversion,omitempty"`
	EventPayPurchaseAmountWeekByConversionRoi *float64 `json:"event_pay_purchase_amount_week_by_conversion_roi,omitempty"`
	EventPayWeekByConversion *int64 `json:"event_pay_week_by_conversion,omitempty"`
	EventPayWeekByConversionCost *float64 `json:"event_pay_week_by_conversion_cost,omitempty"`
	EventPayWeightedPurchaseAmount *float64 `json:"event_pay_weighted_purchase_amount,omitempty"`
	EventPayWeightedPurchaseAmountFirstDay *float64 `json:"event_pay_weighted_purchase_amount_first_day,omitempty"`
	EventPreComponentConsultationValidRetained *int64 `json:"event_pre_component_consultation_valid_retained,omitempty"`
	EventSixDayStayByConversion *int64 `json:"event_six_day_stay_by_conversion,omitempty"`
	EventSixDayStayByConversionCost *float64 `json:"event_six_day_stay_by_conversion_cost,omitempty"`
	EventSixDayStayByConversionRatio *float64 `json:"event_six_day_stay_by_conversion_ratio,omitempty"`
	EventThreeDayStayByConversion *int64 `json:"event_three_day_stay_by_conversion,omitempty"`
	EventThreeDayStayByConversionCost *float64 `json:"event_three_day_stay_by_conversion_cost,omitempty"`
	EventThreeDayStayByConversionRatio *float64 `json:"event_three_day_stay_by_conversion_ratio,omitempty"`
	EventTwoDayStayByConversion *int64 `json:"event_two_day_stay_by_conversion,omitempty"`
	EventTwoDayStayByConversionCost *float64 `json:"event_two_day_stay_by_conversion_cost,omitempty"`
	EventTwoDayStayByConversionRatio *float64 `json:"event_two_day_stay_by_conversion_ratio,omitempty"`
	EventWechatQrCodeLinkClick *int64 `json:"event_wechat_qr_code_link_click,omitempty"`
	EventWeekStay *int64 `json:"event_week_stay,omitempty"`
	EventWeekStayByConversion *int64 `json:"event_week_stay_by_conversion,omitempty"`
	EventWeekStayByConversionCost *float64 `json:"event_week_stay_by_conversion_cost,omitempty"`
	EventWeekStayByConversionRatio *float64 `json:"event_week_stay_by_conversion_ratio,omitempty"`
	EventWeekStayCost *float64 `json:"event_week_stay_cost,omitempty"`
	EventWeekStayRatio *float64 `json:"event_week_stay_ratio,omitempty"`
	LiveEventGoodsView *int64 `json:"live_event_goods_view,omitempty"`
	LivePlayed3s *int64 `json:"live_played_3s,omitempty"`
	PlayedEnd *int64 `json:"played_end,omitempty"`
	PlayedFiveSeconds *int64 `json:"played_five_seconds,omitempty"`
	PlayedNum *int64 `json:"played_num,omitempty"`
	PlayedThreeSeconds *int64 `json:"played_three_seconds,omitempty"`
	AdScene *string `json:"ad_scene,omitempty"`
	PlacementType *string `json:"placement_type,omitempty"`
	CancelFollow *int64 `json:"cancel_follow,omitempty"`
	DownloadStarted *int64 `json:"download_started,omitempty"`
	DownloadCompleted *int64 `json:"download_completed,omitempty"`
	StatDate *string `json:"stat_date,omitempty"`
	StatHour *int64 `json:"stat_hour,omitempty"`
	PhotoClick *int64 `json:"photo_click,omitempty"`
	PhotoClickRatio *float64 `json:"photo_click_ratio,omitempty"`
	ActionRatio *float64 `json:"action_ratio,omitempty"`
	Impression1kCost *float64 `json:"impression_1k_cost,omitempty"`
	PhotoClickCost *float64 `json:"photo_click_cost,omitempty"`
	Click1kCost *float64 `json:"click_1k_cost,omitempty"`
	ActionCost *float64 `json:"action_cost,omitempty"`
	EventPayFirstDay *int64 `json:"event_pay_first_day,omitempty"`
	EventPayPurchaseAmountFirstDay *float64 `json:"event_pay_purchase_amount_first_day,omitempty"`
	EventPayFirstDayRoi *float64 `json:"event_pay_first_day_roi,omitempty"`
	EventPay *int64 `json:"event_pay,omitempty"`
	EventPayPurchaseAmount *float64 `json:"event_pay_purchase_amount,omitempty"`
	EventPayRoi *float64 `json:"event_pay_roi,omitempty"`
	EventRegister *int64 `json:"event_register,omitempty"`
	EventRegisterCost *float64 `json:"event_register_cost,omitempty"`
	EventRegisterRatio *float64 `json:"event_register_ratio,omitempty"`
	EventJinJianApp *int64 `json:"event_jin_jian_app,omitempty"`
	EventJinJianAppCost *float64 `json:"event_jin_jian_app_cost,omitempty"`
	EventCreditGrantApp *int64 `json:"event_credit_grant_app,omitempty"`
	EventCreditGrantAppCost *float64 `json:"event_credit_grant_app_cost,omitempty"`
	EventCreditGrantAppRatio *float64 `json:"event_credit_grant_app_ratio,omitempty"`
	EventOrderPaid *int64 `json:"event_order_paid,omitempty"`
	EventOrderPaidPurchaseAmount *float64 `json:"event_order_paid_purchase_amount,omitempty"`
	EventOrderPaidCost *float64 `json:"event_order_paid_cost,omitempty"`
	FormCount *int64 `json:"form_count,omitempty"`
	FormCost *float64 `json:"form_cost,omitempty"`
	FormActionRatio *float64 `json:"form_action_ratio,omitempty"`
	EventJinJianLandingPage *int64 `json:"event_jin_jian_landing_page,omitempty"`
	EventJinJianLandingPageCost *float64 `json:"event_jin_jian_landing_page_cost,omitempty"`
	EventCreditGrantLandingPage *int64 `json:"event_credit_grant_landing_page,omitempty"`
	EventCreditGrantLandingPageCost *float64 `json:"event_credit_grant_landing_page_cost,omitempty"`
	EventCreditGrantLandingRatio *float64 `json:"event_credit_grant_landing_ratio,omitempty"`
	EventNextDayStayCost *float64 `json:"event_next_day_stay_cost,omitempty"`
	EventNextDayStayRatio *float64 `json:"event_next_day_stay_ratio,omitempty"`
	EventNextDayStay *int64 `json:"event_next_day_stay,omitempty"`
	Play3sRatio *float64 `json:"play_3s_ratio,omitempty"`
	EventValidClues *int64 `json:"event_valid_clues,omitempty"`
	EventValidCluesCost *float64 `json:"event_valid_clues_cost,omitempty"`
	AdProductCnt *int64 `json:"ad_product_cnt,omitempty"`
	EventGoodsView *int64 `json:"event_goods_view,omitempty"`
	MerchantRecoFans *int64 `json:"merchant_reco_fans,omitempty"`
	EventOrderAmountRoi *float64 `json:"event_order_amount_roi,omitempty"`
	EventGoodsViewCost *float64 `json:"event_goods_view_cost,omitempty"`
	MerchantRecoFansCost *float64 `json:"merchant_reco_fans_cost,omitempty"`
	EventNewUserPay *int64 `json:"event_new_user_pay,omitempty"`
	EventNewUserPayCost *float64 `json:"event_new_user_pay_cost,omitempty"`
	EventNewUserPayRatio *float64 `json:"event_new_user_pay_ratio,omitempty"`
	EventButtonClick *int64 `json:"event_button_click,omitempty"`
	EventButtonClickCost *float64 `json:"event_button_click_cost,omitempty"`
	EventButtonClickRatio *float64 `json:"event_button_click_ratio,omitempty"`
	Play5sRatio *float64 `json:"play_5s_ratio,omitempty"`
	PlayEndRatio *float64 `json:"play_end_ratio,omitempty"`
	EventOrderPaidRoi *float64 `json:"event_order_paid_roi,omitempty"`
	EventNewUserJinjianApp *int64 `json:"event_new_user_jinjian_app,omitempty"`
	EventNewUserJinjianAppCost *float64 `json:"event_new_user_jinjian_app_cost,omitempty"`
	EventNewUserJinjianAppRoi *float64 `json:"event_new_user_jinjian_app_roi,omitempty"`
	EventNewUserCreditGrantApp *int64 `json:"event_new_user_credit_grant_app,omitempty"`
	EventNewUserCreditGrantAppCost *float64 `json:"event_new_user_credit_grant_app_cost,omitempty"`
	EventNewUserCreditGrantAppRoi *float64 `json:"event_new_user_credit_grant_app_roi,omitempty"`
	EventNewUserJinjianPage *int64 `json:"event_new_user_jinjian_page,omitempty"`
	EventNewUserJinjianPageCost *float64 `json:"event_new_user_jinjian_page_cost,omitempty"`
	EventNewUserJinjianPageRoi *float64 `json:"event_new_user_jinjian_page_roi,omitempty"`
	EventNewUserCreditGrantPage *int64 `json:"event_new_user_credit_grant_page,omitempty"`
	EventNewUserCreditGrantPageCost *float64 `json:"event_new_user_credit_grant_page_cost,omitempty"`
	EventNewUserCreditGrantPageRoi *float64 `json:"event_new_user_credit_grant_page_roi,omitempty"`
	EventAppointForm *int64 `json:"event_appoint_form,omitempty"`
	EventAppointFormCost *float64 `json:"event_appoint_form_cost,omitempty"`
	EventAppointFormRatio *float64 `json:"event_appoint_form_ratio,omitempty"`
	EventAppointJumpClick *int64 `json:"event_appoint_jump_click,omitempty"`
	EventAppointJumpClickCost *float64 `json:"event_appoint_jump_click_cost,omitempty"`
	EventAppointJumpClickRatio *float64 `json:"event_appoint_jump_click_ratio,omitempty"`
	EventDspGiftForm *int64 `json:"event_dsp_gift_form,omitempty"`
	EventAppInvoked *int64 `json:"event_app_invoked,omitempty"`
	EventAppInvokedCost *float64 `json:"event_app_invoked_cost,omitempty"`
	EventAppInvokedRatio *float64 `json:"event_app_invoked_ratio,omitempty"`
	EventAddWechat *int64 `json:"event_add_wechat,omitempty"`
	EventAddWechatCost *float64 `json:"event_add_wechat_cost,omitempty"`
	EventAddWechatRatio *float64 `json:"event_add_wechat_ratio,omitempty"`
	EventMultiConversion *int64 `json:"event_multi_conversion,omitempty"`
	EventMultiConversionRatio *float64 `json:"event_multi_conversion_ratio,omitempty"`
	EventMultiConversionCost *float64 `json:"event_multi_conversion_cost,omitempty"`
	EventWatchAppAd *int64 `json:"event_watch_app_ad,omitempty"`
	EventAdWatchTimes *int64 `json:"event_ad_watch_times,omitempty"`
	EventAdWatchTimesRatio *float64 `json:"event_ad_watch_times_ratio,omitempty"`
	EventAdWatchTimesCost *float64 `json:"event_ad_watch_times_cost,omitempty"`
	EventAddShoppingCart *int64 `json:"event_add_shopping_cart,omitempty"`
	EventAddShoppingCartCost *float64 `json:"event_add_shopping_cart_cost,omitempty"`
	EventGetThrough *int64 `json:"event_get_through,omitempty"`
	EventGetThroughCost *float64 `json:"event_get_through_cost,omitempty"`
	EventGetThroughRatio *float64 `json:"event_get_through_ratio,omitempty"`
	AdPhotoPlayed75percentRatio *float64 `json:"ad_photo_played_75percent_ratio,omitempty"`
	AdPhotoPlayed10sRatio *float64 `json:"ad_photo_played_10s_ratio,omitempty"`
	AdPhotoPlayed2sRatio *float64 `json:"ad_photo_played_2s_ratio,omitempty"`
	EventPhoneGetThrough *int64 `json:"event_phone_get_through,omitempty"`
	EventIntentionConfirmed *int64 `json:"event_intention_confirmed,omitempty"`
	EventWechatConnected *int64 `json:"event_wechat_connected,omitempty"`
	EventOrderSuccessed *int64 `json:"event_order_successed,omitempty"`
	EventPhoneCardActivate *int64 `json:"event_phone_card_activate,omitempty"`
	EventMeasurementHouse *int64 `json:"event_measurement_house,omitempty"`
	AdShow *float64 `json:"ad_show,omitempty"`
	ActionNewRatio *float64 `json:"action_new_ratio,omitempty"`
	EventOutboundCall *int64 `json:"event_outbound_call,omitempty"`
	EventOutboundCallCost *float64 `json:"event_outbound_call_cost,omitempty"`
	EventOutboundCallRatio *float64 `json:"event_outbound_call_ratio,omitempty"`
	KeyAction *int64 `json:"key_action,omitempty"`
	KeyActionCost *float64 `json:"key_action_cost,omitempty"`
	KeyActionRatio *float64 `json:"key_action_ratio,omitempty"`
	EventCreditCardRecheck *int64 `json:"event_credit_card_recheck,omitempty"`
	EventCreditCardRecheckFirstDay *int64 `json:"event_credit_card_recheck_first_day,omitempty"`
	EventNoIntention *int64 `json:"event_no_intention,omitempty"`
	EventActive *int64 `json:"event_active,omitempty"`
	EventOrderSubmitCost *float64 `json:"event_order_submit_cost,omitempty"`
	OrderSubmitAmt *int64 `json:"order_submit_amt,omitempty"`
	OrderSubmitRoi *float64 `json:"order_submit_roi,omitempty"`
	NativeFlowSoftCostTotal *float64 `json:"native_flow_soft_cost_total,omitempty"`
	NativeCostTotal *float64 `json:"native_cost_total,omitempty"`
	PhotoId *string `json:"photo_id,omitempty"`
	PhotoUrl *string `json:"photo_url,omitempty"`
	CoverUrl *string `json:"cover_url,omitempty"`
	PhotoCaption *string `json:"photo_caption,omitempty"`
	AigcMaterial *bool `json:"aigc_material,omitempty"`
}


