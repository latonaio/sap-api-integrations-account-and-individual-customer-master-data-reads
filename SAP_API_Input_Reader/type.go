package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema      string `json:"api_schema"`
	MaterialCode   string `json:"material_code"`
	Plant_Supplier string `json:"plant/supplier"`
	Stock          string `json:"stock"`
	DocumentType   string `json:"document_type"`
	DocumentNo     string `json:"document_no"`
	PlannedDate    string `json:"planned_date"`
	ValidatedDate  string `json:"validated_date"`
	Deleted        bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	ContactCollection struct {
		ObjectID                              string `json:"ObjectID"`
		ContactID                             string `json:"ContactID"`
		ContactUUID                           string `json:"ContactUUID"`
		ExternalID                            string `json:"ExternalID"`
		ExternalSystem                        string `json:"ExternalSystem"`
		StatusCode                            string `json:"StatusCode"`
		StatusCodeText                        string `json:"StatusCodeText"`
		TitleCode                             string `json:"TitleCode"`
		TitleCodeText                         string `json:"TitleCodeText"`
		AcademicTitleCode                     string `json:"AcademicTitleCode"`
		AcademicTitleCodeText                 string `json:"AcademicTitleCodeText"`
		AdditionalAcademicTitleCode           string `json:"AdditionalAcademicTitleCode"`
		AdditionalAcademicTitleCodeText       string `json:"AdditionalAcademicTitleCodeText"`
		NamePrefixCode                        string `json:"NamePrefixCode"`
		NamePrefixCodeText                    string `json:"NamePrefixCodeText"`
		FirstName                             string `json:"FirstName"`
		LastName                              string `json:"LastName"`
		AdditionalFamilyName                  string `json:"AdditionalFamilyName"`
		Initials                              string `json:"Initials"`
		MiddleName                            string `json:"MiddleName"`
		Name                                  string `json:"Name"`
		GenderCode                            string `json:"GenderCode"`
		GenderCodeText                        string `json:"GenderCodeText"`
		MaritalStatusCode                     string `json:"MaritalStatusCode"`
		MaritalStatusCodeText                 string `json:"MaritalStatusCodeText"`
		LanguageCode                          string `json:"LanguageCode"`
		LanguageCodeText                      string `json:"LanguageCodeText"`
		NickName                              string `json:"NickName"`
		BirthDate                             string `json:"BirthDate"`
		BirthName                             string `json:"BirthName"`
		ContactPermissionCode                 string `json:"ContactPermissionCode"`
		ContactPermissionCodeText             string `json:"ContactPermissionCodeText"`
		ProfessionCode                        string `json:"ProfessionCode"`
		ProfessionCodeText                    string `json:"ProfessionCodeText"`
		PerceptionOfCompanyCode               string `json:"PerceptionOfCompanyCode"`
		PerceptionOfCompanyCodeText           string `json:"PerceptionOfCompanyCodeText"`
		DeviatingFullName                     string `json:"DeviatingFullName"`
		AccountID                             string `json:"AccountID"`
		AccountUUID                           string `json:"AccountUUID"`
		AccountFormattedName                  string `json:"AccountFormattedName"`
		Building                              string `json:"Building"`
		Floor                                 string `json:"Floor"`
		Room                                  string `json:"Room"`
		JobTitle                              string `json:"JobTitle"`
		FunctionCode                          string `json:"FunctionCode"`
		FunctionCodeText                      string `json:"FunctionCodeText"`
		DepartmentCode                        string `json:"DepartmentCode"`
		DepartmentCodeText                    string `json:"DepartmentCodeText"`
		Department                            string `json:"Department"`
		VIPContactCode                        string `json:"VIPContactCode"`
		VIPContactCodeText                    string `json:"VIPContactCodeText"`
		Phone                                 string `json:"Phone"`
		Mobile                                string `json:"Mobile"`
		Fax                                   string `json:"Fax"`
		Email                                 string `json:"Email"`
		EmailInvalidIndicator                 bool   `json:"EmailInvalidIndicator"`
		BestReachedByCode                     string `json:"BestReachedByCode"`
		BestReachedByCodeText                 string `json:"BestReachedByCodeText"`
		FormattedPostalAddressDescription     string `json:"FormattedPostalAddressDescription"`
		BusinessAddressCountryCode            string `json:"BusinessAddressCountryCode"`
		BusinessAddressCountryCodeText        string `json:"BusinessAddressCountryCodeText"`
		BusinessAddressStateCodeTextUpdatable string `json:"BusinessAddressStateCodeTextUpdatable"`
		BusinessAddressHouseNumber            string `json:"BusinessAddressHouseNumber"`
		BusinessAddressStreet                 string `json:"BusinessAddressStreet"`
		BusinessAddressCity                   string `json:"BusinessAddressCity"`
		BusinessAddressStreetPostalCode       string `json:"BusinessAddressStreetPostalCode"`
		BusinessAddressStateCode              string `json:"BusinessAddressStateCode"`
		BusinessAddressStateCodeText          string `json:"BusinessAddressStateCodeText"`
		CreationOn                            string `json:"CreationOn"`
		CreatedBy                             string `json:"CreatedBy"`
		CreatedByIdentityUUID                 string `json:"CreatedByIdentityUUID"`
		ChangedOn                             string `json:"ChangedOn"`
		ChangedBy                             string `json:"ChangedBy"`
		ChangedByIdentityUUID                 string `json:"ChangedByIdentityUUID"`
		ContactOwnerID                        string `json:"ContactOwnerID"`
		ContactOwnerUUID                      string `json:"ContactOwnerUUID"`
		NormalisedPhone                       string `json:"NormalisedPhone"`
		NormalisedMobile                      string `json:"NormalisedMobile"`
		EntityLastChangedOn                   string `json:"EntityLastChangedOn"`
		ETag                                  string `json:"ETag"`
		ContactIsContactPersonFor             struct {
			ObjectID                string `json:"ObjectID"`
			ParentObjectID          string `json:"ParentObjectID"`
			ETag                    string `json:"ETag"`
			ContactID               string `json:"ContactID"`
			AccountID               string `json:"AccountID"`
			AccountFormattedName    string `json:"AccountFormattedName"`
			ReverseMainIndicator    bool   `json:"ReverseMainIndicator"`
			DepartmentCode          string `json:"DepartmentCode"`
			DepartmentCodeText      string `json:"DepartmentCodeText"`
			FunctionCode            string `json:"FunctionCode"`
			FunctionCodeText        string `json:"FunctionCodeText"`
			VIPReasonCode           string `json:"VIPReasonCode"`
			VIPReasonCodeText       string `json:"VIPReasonCodeText"`
			JobTitle                string `json:"JobTitle"`
			Department              string `json:"Department"`
			Building                string `json:"Building"`
			Floor                   string `json:"Floor"`
			Room                    string `json:"Room"`
			Email                   string `json:"Email"`
			EmailInvalidIndicator   bool   `json:"EmailInvalidIndicator"`
			Fax                     string `json:"Fax"`
			Mobile                  string `json:"Mobile"`
			Phone                   string `json:"Phone"`
			BestReachedByCode       string `json:"BestReachedByCode"`
			BestReachedByCodeText   string `json:"BestReachedByCodeText"`
			OrganisationAddressUUID string `json:"OrganisationAddressUUID"`
			EntityLastChangedOn     string `json:"EntityLastChangedOn"`
			CorporateAccount        struct {
				ObjectID                          string `json:"ObjectID"`
				AccountID                         string `json:"AccountID"`
				UUID                              string `json:"UUID"`
				ExternalID                        string `json:"ExternalID"`
				ExternalSystem                    string `json:"ExternalSystem"`
				RoleCode                          string `json:"RoleCode"`
				RoleCodeText                      string `json:"RoleCodeText"`
				LifeCycleStatusCode               string `json:"LifeCycleStatusCode"`
				LifeCycleStatusCodeText           string `json:"LifeCycleStatusCodeText"`
				Dunsid                            string `json:"DUNSID"`
				LegalFormCode                     string `json:"LegalFormCode"`
				LegalFormCodeText                 string `json:"LegalFormCodeText"`
				CustomerABCClassificationCode     string `json:"CustomerABCClassificationCode"`
				CustomerABCClassificationCodeText string `json:"CustomerABCClassificationCodeText"`
				NielsenRegionCode                 string `json:"NielsenRegionCode"`
				NielsenRegionCodeText             string `json:"NielsenRegionCodeText"`
				IndustrialSectorCode              string `json:"IndustrialSectorCode"`
				IndustrialSectorCodeText          string `json:"IndustrialSectorCodeText"`
				ContactPermissionCode             string `json:"ContactPermissionCode"`
				ContactPermissionCodeText         string `json:"ContactPermissionCodeText"`
				BusinessPartnerFormattedName      string `json:"BusinessPartnerFormattedName"`
				Name                              string `json:"Name"`
				AdditionalName                    string `json:"AdditionalName"`
				AdditionalName2                   string `json:"AdditionalName2"`
				AdditionalName3                   string `json:"AdditionalName3"`
				CurrentDefaultAddressUUID         string `json:"CurrentDefaultAddressUUID"`
				FormattedPostalAddressDescription string `json:"FormattedPostalAddressDescription"`
				CountryCode                       string `json:"CountryCode"`
				CountryCodeText                   string `json:"CountryCodeText"`
				StateCode                         string `json:"StateCode"`
				StateCodeText                     string `json:"StateCodeText"`
				CareOfName                        string `json:"CareOfName"`
				AddressLine1                      string `json:"AddressLine1"`
				AddressLine2                      string `json:"AddressLine2"`
				HouseNumber                       string `json:"HouseNumber"`
				AdditionalHouseNumber             string `json:"AdditionalHouseNumber"`
				Street                            string `json:"Street"`
				AddressLine4                      string `json:"AddressLine4"`
				AddressLine5                      string `json:"AddressLine5"`
				District                          string `json:"District"`
				City                              string `json:"City"`
				DifferentCity                     string `json:"DifferentCity"`
				StreetPostalCode                  string `json:"StreetPostalCode"`
				County                            string `json:"County"`
				CompanyPostalCode                 string `json:"CompanyPostalCode"`
				POBoxIndicator                    bool   `json:"POBoxIndicator"`
				POBox                             string `json:"POBox"`
				POBoxPostalCode                   string `json:"POBoxPostalCode"`
				POBoxDeviatingCountryCode         string `json:"POBoxDeviatingCountryCode"`
				POBoxDeviatingCountryCodeText     string `json:"POBoxDeviatingCountryCodeText"`
				POBoxDeviatingRegionCode          string `json:"POBoxDeviatingRegionCode"`
				POBoxDeviatingRegionCodeText      string `json:"POBoxDeviatingRegionCodeText"`
				POBoxDeviatingCity                string `json:"POBoxDeviatingCity"`
				TimeZoneCode                      string `json:"TimeZoneCode"`
				TimeZoneCodeText                  string `json:"TimeZoneCodeText"`
				Building                          string `json:"Building"`
				Floor                             string `json:"Floor"`
				Room                              string `json:"Room"`
				Phone                             string `json:"Phone"`
				NormalisedPhone                   string `json:"NormalisedPhone"`
				Mobile                            string `json:"Mobile"`
				NormalisedMobile                  string `json:"NormalisedMobile"`
				Fax                               string `json:"Fax"`
				Email                             string `json:"Email"`
				WebSite                           string `json:"WebSite"`
				LanguageCode                      string `json:"LanguageCode"`
				LanguageCodeText                  string `json:"LanguageCodeText"`
				BestReachedByCode                 string `json:"BestReachedByCode"`
				BestReachedByCodeText             string `json:"BestReachedByCodeText"`
				OrderBlockingReasonCode           string `json:"OrderBlockingReasonCode"`
				OrderBlockingReasonCodeText       string `json:"OrderBlockingReasonCodeText"`
				DeliveryBlockingReasonCode        string `json:"DeliveryBlockingReasonCode"`
				DeliveryBlockingReasonCodeText    string `json:"DeliveryBlockingReasonCodeText"`
				BillingBlockingReasonCode         string `json:"BillingBlockingReasonCode"`
				BillingBlockingReasonCodeText     string `json:"BillingBlockingReasonCodeText"`
				SalesSupportBlockingIndicator     bool   `json:"SalesSupportBlockingIndicator"`
				LegalCompetenceIndicator          bool   `json:"LegalCompetenceIndicator"`
				RecommendedVisitingFrequency      string `json:"RecommendedVisitingFrequency"`
				VisitDuration                     string `json:"VisitDuration"`
				LastVisitingDate                  string `json:"LastVisitingDate"`
				NextVisitingDate                  string `json:"NextVisitingDate"`
				LatestRecommendedVisitingDate     string `json:"LatestRecommendedVisitingDate"`
				OwnerID                           string `json:"OwnerID"`
				OwnerUUID                         string `json:"OwnerUUID"`
				ParentAccountID                   string `json:"ParentAccountID"`
				CreationOn                        string `json:"CreationOn"`
				CreatedBy                         string `json:"CreatedBy"`
				CreatedByIdentityUUID             string `json:"CreatedByIdentityUUID"`
				ChangedOn                         string `json:"ChangedOn"`
				ChangedBy                         string `json:"ChangedBy"`
				ChangedByIdentityUUID             string `json:"ChangedByIdentityUUID"`
				EntityLastChangedOn               string `json:"EntityLastChangedOn"`
				ETag                              string `json:"ETag"`
				Mul1                              string `json:"mul1"`
			} `json:"CorporateAccount"`
		} `json:"ContactIsContactPersonFor"`
	} `json:"ContactCollection"`
	EmployeeBasicData struct {
		ObjectID                         string `json:"ObjectID"`
		ETag                             string `json:"ETag"`
		EmployeeID                       string `json:"EmployeeID"`
		EmployeeUUID                     string `json:"EmployeeUUID"`
		UserID                           string `json:"UserID"`
		IdentityUUID                     string `json:"IdentityUUID"`
		BusinessPartnerID                string `json:"BusinessPartnerID"`
		CurrentInternalEmployeeIndicator string `json:"CurrentInternalEmployeeIndicator"`
		CurrentExternalEmployeeIndicator string `json:"CurrentExternalEmployeeIndicator"`
		FormattedName                    string `json:"FormattedName"`
		TitleCode                        string `json:"TitleCode"`
		AcademicTitleCode                string `json:"AcademicTitleCode"`
		FirstName                        string `json:"FirstName"`
		MiddleName                       string `json:"MiddleName"`
		LastName                         string `json:"LastName"`
		SecondLastName                   string `json:"SecondLastName"`
		NickName                         string `json:"NickName"`
		GenderCode                       string `json:"GenderCode"`
		LanguageCode                     string `json:"LanguageCode"`
		FormattedAddress                 string `json:"FormattedAddress"`
		RegionCode                       string `json:"RegionCode"`
		AddressLine1                     string `json:"AddressLine1"`
		AddressLine2                     string `json:"AddressLine2"`
		HouseNumber                      string `json:"HouseNumber"`
		Street                           string `json:"Street"`
		AddressLine4                     string `json:"AddressLine4"`
		AddressLine5                     string `json:"AddressLine5"`
		City                             string `json:"City"`
		PostalCode                       string `json:"PostalCode"`
		Phone                            string `json:"Phone"`
		Mobile                           string `json:"Mobile"`
		Fax                              string `json:"Fax"`
		Email                            string `json:"Email"`
		UserValidityStartDate            string `json:"UserValidityStartDate"`
		UserValidityEndDate              string `json:"UserValidityEndDate"`
		UserPasswordPolicyCode           string `json:"UserPasswordPolicyCode"`
		UserLockedIndicator              string `json:"UserLockedIndicator"`
		TimeZoneCode                     string `json:"TimeZoneCode"`
		ManagerUUID                      string `json:"ManagerUUID"`
		ManagerFormattedName             string `json:"ManagerFormattedName"`
		JobName                          string `json:"JobName"`
		CreatedOn                        string `json:"CreatedOn"`
		CreatedBy                        string `json:"CreatedBy"`
		ChangedOn                        string `json:"ChangedOn"`
		ChangedBy                        string `json:"ChangedBy"`
		EntityLastChangedOn              string `json:"EntityLastChangedOn"`
	} `json:"EmployeeBasicData"`
	IndividualCustomerCollection struct {
		ObjectID                          string `json:"ObjectID"`
		CustomerID                        string `json:"CustomerID"`
		UUID                              string `json:"UUID"`
		ExternalID                        string `json:"ExternalID"`
		ExternalSystem                    string `json:"ExternalSystem"`
		RoleCode                          string `json:"RoleCode"`
		RoleCodeText                      string `json:"RoleCodeText"`
		LifeCycleStatusCode               string `json:"LifeCycleStatusCode"`
		LifeCycleStatusCodeText           string `json:"LifeCycleStatusCodeText"`
		CustomerABCClassificationCode     string `json:"CustomerABCClassificationCode"`
		CustomerABCClassificationCodeText string `json:"CustomerABCClassificationCodeText"`
		ContactPermissionCode             string `json:"ContactPermissionCode"`
		ContactPermissionCodeText         string `json:"ContactPermissionCodeText"`
		TitleCode                         string `json:"TitleCode"`
		TitleCodeText                     string `json:"TitleCodeText"`
		AcademicTitleCode                 string `json:"AcademicTitleCode"`
		AcademicTitleCodeText             string `json:"AcademicTitleCodeText"`
		FirstName                         string `json:"FirstName"`
		MiddleName                        string `json:"MiddleName"`
		LastName                          string `json:"LastName"`
		AdditionalLastName                string `json:"AdditionalLastName"`
		Initials                          string `json:"Initials"`
		NickName                          string `json:"NickName"`
		GenderCode                        string `json:"GenderCode"`
		GenderCodeText                    string `json:"GenderCodeText"`
		NamePrefixCode                    string `json:"NamePrefixCode"`
		NamePrefixCodeText                string `json:"NamePrefixCodeText"`
		MaritalStatusCode                 string `json:"MaritalStatusCode"`
		MaritalStatusCodeText             string `json:"MaritalStatusCodeText"`
		LanguageCode                      string `json:"LanguageCode"`
		LanguageCodeText                  string `json:"LanguageCodeText"`
		BirthName                         string `json:"BirthName"`
		BirthDate                         string `json:"BirthDate"`
		NationalityCountryCode            string `json:"NationalityCountryCode"`
		NationalityCountryCodeText        string `json:"NationalityCountryCodeText"`
		ProfessionCode                    string `json:"ProfessionCode"`
		ProfessionCodeText                string `json:"ProfessionCodeText"`
		FormattedName                     string `json:"FormattedName"`
		FormattedPostalAddressDescription string `json:"FormattedPostalAddressDescription"`
		CountryCode                       string `json:"CountryCode"`
		CountryCodeText                   string `json:"CountryCodeText"`
		StateCode                         string `json:"StateCode"`
		StateCodeText                     string `json:"StateCodeText"`
		CareOfName                        string `json:"CareOfName"`
		AddressLine1                      string `json:"AddressLine1"`
		AddressLine2                      string `json:"AddressLine2"`
		HouseNumber                       string `json:"HouseNumber"`
		AdditionalHouseNumber             string `json:"AdditionalHouseNumber"`
		Street                            string `json:"Street"`
		AddressLine4                      string `json:"AddressLine4"`
		AddressLine5                      string `json:"AddressLine5"`
		District                          string `json:"District"`
		City                              string `json:"City"`
		DifferentCity                     string `json:"DifferentCity"`
		StreetPostalCode                  string `json:"StreetPostalCode"`
		County                            string `json:"County"`
		POBoxIndicator                    string `json:"POBoxIndicator"`
		POBox                             string `json:"POBox"`
		POBoxPostalCode                   string `json:"POBoxPostalCode"`
		POBoxDeviatingCountryCode         string `json:"POBoxDeviatingCountryCode"`
		POBoxDeviatingCountryCodeText     string `json:"POBoxDeviatingCountryCodeText"`
		POBoxDeviatingStateCode           string `json:"POBoxDeviatingStateCode"`
		POBoxDeviatingStateCodeText       string `json:"POBoxDeviatingStateCodeText"`
		POBoxDeviatingCity                string `json:"POBoxDeviatingCity"`
		TimeZoneCode                      string `json:"TimeZoneCode"`
		TimeZoneCodeText                  string `json:"TimeZoneCodeText"`
		TaxJurisdictionCode               string `json:"TaxJurisdictionCode"`
		TaxJurisdictionCodeText           string `json:"TaxJurisdictionCodeText"`
		Building                          string `json:"Building"`
		Floor                             string `json:"Floor"`
		Room                              string `json:"Room"`
		Phone                             string `json:"Phone"`
		NormalisedPhone                   string `json:"NormalisedPhone"`
		Mobile                            string `json:"Mobile"`
		NormalisedMobile                  string `json:"NormalisedMobile"`
		Fax                               string `json:"Fax"`
		Email                             string `json:"Email"`
		EmailInvalidIndicator             string `json:"EmailInvalidIndicator"`
		WebSite                           string `json:"WebSite"`
		BestReachedByCode                 string `json:"BestReachedByCode"`
		BestReachedByCodeText             string `json:"BestReachedByCodeText"`
		OrderBlockingReasonCode           string `json:"OrderBlockingReasonCode"`
		OrderBlockingReasonCodeText       string `json:"OrderBlockingReasonCodeText"`
		DeliveryBlockingReasonCode        string `json:"DeliveryBlockingReasonCode"`
		DeliveryBlockingReasonCodeText    string `json:"DeliveryBlockingReasonCodeText"`
		BillingBlockingReasonCode         string `json:"BillingBlockingReasonCode"`
		BillingBlockingReasonCodeText     string `json:"BillingBlockingReasonCodeText"`
		SalesSupportBlockingIndicator     string `json:"SalesSupportBlockingIndicator"`
		RecommendedVisitingFrequency      string `json:"RecommendedVisitingFrequency"`
		VisitDuration                     string `json:"VisitDuration"`
		LastVisitingDate                  string `json:"LastVisitingDate"`
		NextVisitingDate                  string `json:"NextVisitingDate"`
		LatestRecommendedVisitingDate     string `json:"LatestRecommendedVisitingDate"`
		OwnerID                           string `json:"OwnerID"`
		OwnerUUID                         string `json:"OwnerUUID"`
		CreationOn                        string `json:"CreationOn"`
		CreatedBy                         string `json:"CreatedBy"`
		CreatedByIdentityUUID             string `json:"CreatedByIdentityUUID"`
		ChangedOn                         string `json:"ChangedOn"`
		ChangedBy                         string `json:"ChangedBy"`
		ChangedByIdentityUUID             string `json:"ChangedByIdentityUUID"`
		EntityLastChangedOn               string `json:"EntityLastChangedOn"`
		ETag                              string `json:"ETag"`
		Mul1                              string `json:"Mul1"`
		IndividualCustomerAddress         struct {
			ObjectID                                    string `json:"ObjectID"`
			ParentObjectID                              string `json:"ParentObjectID"`
			CustomerID                                  string `json:"CustomerID"`
			MainIndicator                               string `json:"MainIndicator"`
			ShipTo                                      string `json:"ShipTo"`
			DefaultShipTo                               string `json:"DefaultShipTo"`
			BillTo                                      string `json:"BillTo"`
			DefaultBillTo                               string `json:"DefaultBillTo"`
			FormattedPostalAddressDescription           string `json:"FormattedPostalAddressDescription"`
			FormattedAddressFirstLineDescription        string `json:"FormattedAddressFirstLineDescription"`
			FormattedAddressSecondLineDescription       string `json:"FormattedAddressSecondLineDescription"`
			FormattedAddressThirdLineDescription        string `json:"FormattedAddressThirdLineDescription"`
			FormattedAddressFourthLineDescription       string `json:"FormattedAddressFourthLineDescription"`
			FormattedPostalAddressFirstLineDescription  string `json:"FormattedPostalAddressFirstLineDescription"`
			FormattedPostalAddressSecondLineDescription string `json:"FormattedPostalAddressSecondLineDescription"`
			FormattedPostalAddressThirdLineDescription  string `json:"FormattedPostalAddressThirdLineDescription"`
			CountryCode                                 string `json:"CountryCode"`
			CountryCodeText                             string `json:"CountryCodeText"`
			StateCode                                   string `json:"StateCode"`
			StateCodeText                               string `json:"StateCodeText"`
			CareOfName                                  string `json:"CareOfName"`
			AddressLine1                                string `json:"AddressLine1"`
			AddressLine2                                string `json:"AddressLine2"`
			HouseNumber                                 string `json:"HouseNumber"`
			AdditionalHouseNumber                       string `json:"AdditionalHouseNumber"`
			Street                                      string `json:"Street"`
			AddressLine4                                string `json:"AddressLine4"`
			AddressLine5                                string `json:"AddressLine5"`
			District                                    string `json:"District"`
			City                                        string `json:"City"`
			DifferentCity                               string `json:"DifferentCity"`
			StreetPostalCode                            string `json:"StreetPostalCode"`
			County                                      string `json:"County"`
			POBoxIndicator                              string `json:"POBoxIndicator"`
			POBox                                       string `json:"POBox"`
			POBoxPostalCode                             string `json:"POBoxPostalCode"`
			POBoxDeviatingCountryCode                   string `json:"POBoxDeviatingCountryCode"`
			POBoxDeviatingCountryCodeText               string `json:"POBoxDeviatingCountryCodeText"`
			POBoxDeviatingStateCode                     string `json:"POBoxDeviatingStateCode"`
			POBoxDeviatingStateCodeText                 string `json:"POBoxDeviatingStateCodeText"`
			POBoxDeviatingCity                          string `json:"POBoxDeviatingCity"`
			TimeZoneCode                                string `json:"TimeZoneCode"`
			TimeZoneCodeText                            string `json:"TimeZoneCodeText"`
			Latitude                                    string `json:"Latitude"`
			Longitude                                   string `json:"Longitude"`
			Building                                    string `json:"Building"`
			Floor                                       string `json:"Floor"`
			Room                                        string `json:"Room"`
			Phone                                       string `json:"Phone"`
			NormalisedPhone                             string `json:"NormalisedPhone"`
			Mobile                                      string `json:"Mobile"`
			NormalisedMobile                            string `json:"NormalisedMobile"`
			Fax                                         string `json:"Fax"`
			Email                                       string `json:"Email"`
			EmailInvalidIndicator                       string `json:"EmailInvalidIndicator"`
			WebSite                                     string `json:"WebSite"`
			BestReachedByCode                           string `json:"BestReachedByCode"`
			BestReachedByCodeText                       string `json:"BestReachedByCodeText"`
			ETag                                        string `json:"ETag"`
		} `json:"IndividualCustomerAddress"`
	} `json:"IndividualCustomerCollection"`
	APISchema   string   `json:"api_schema"`
	Accepter    []string `json:"accepter"`
	ContactCode string   `json:"contact_code"`
	Deleted     bool     `json:"deleted"`
}
