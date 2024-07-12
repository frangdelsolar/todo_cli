package contractor

type CreateAgreementInput struct {
	ContractorId       string
	AgreementType      string
	FrequencyType      string
	FrequencyDay       string
	FrequencyMonth     string
	FrequencyDayOfWeek string
	StartDate          string
	EndDate            string
	Installments       string
	FixedAmount        string
	CurrencyCode       string
	requestedBy        string
}

func CreateAgreement(input *CreateAgreementInput) (*Agreement, error) {
	return nil, nil
}
