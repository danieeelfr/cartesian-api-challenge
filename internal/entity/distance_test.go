package entity

import (
	"testing"
)

const (
	MyUniqueID         = "my_unique_id"
	MyUniqueExternalID = "my_unique_external_id"
	InvalidDocument    = "11111111111"
	MyValidEmailKey    = "user@example.com"
	MyValidPhoneKey    = "+5551912345678"
	MyValidCpfKey      = "11779537069"
	MyValidCnpjKey     = "26896899000166"
	MyValidEvpKey      = "33904379-af67-43f8-b7fc-07b4c7dec0ed"
	MyValidEndToEndID  = "E00038166201907261559y6j6mt9l0pi"
	EmptyValue         = ""
)

func TestDistanceValidate(t *testing.T) {

	// values := &[]struct {
	// 	description string
	// 	input       *Distance
	// 	output      error
	// }{
	// 	{
	// 		description: "For unfilled data, it must fail.",
	// 		input:       &Distance{},
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"id":       ErrorInvalidUniqueID.Error(),
	// 				"payer":    ErrorPayerDataNull.Error(),
	// 				"receiver": ErrorReceiverDataNull.Error(),
	// 				"value":    ErrorPaymentAmount.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For unfilled payer and payee, must fail.",
	// 		input: &Distance{
	// 			ID:       MyUniqueID,
	// 			Payer:    &Customer{},
	// 			Receiver: &Customer{},
	// 			Value:    123.45,
	// 		},
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"payer.document":    ErrorPayerDocumentInvalid.Error(),
	// 				"payer.name":        ErrorPayerNameInvalid.Error(),
	// 				"payer.account":     ErrorPayerAccountNull.Error(),
	// 				"receiver.ispb":     ErrorReceiverISPB.Error(),
	// 				"receiver.document": ErrorReceiverDocumentInvalid.Error(),
	// 				"receiver.name":     ErrorReceiverNameInvalid.Error(),
	// 				"receiver.account":  ErrorReceiverAccountNull.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For customers filled in incorrectly, it must fail.",
	// 		input: &Distance{
	// 			ID: MyUniqueID,
	// 			Payer: &Customer{
	// 				Document: InvalidDocument,
	// 				Name:     "Fulano pagador",
	// 			},
	// 			Receiver: &Customer{
	// 				Ispb:     BexsISPB,
	// 				Document: InvalidDocument,
	// 				Name:     "Fulano recebedor",
	// 			},
	// 			Value: 123.45,
	// 		},
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"payer.document":    ErrorPayerDocumentInvalid.Error(),
	// 				"payer.account":     ErrorPayerAccountNull.Error(),
	// 				"receiver.document": ErrorReceiverDocumentInvalid.Error(),
	// 				"receiver.account":  ErrorReceiverAccountNull.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For unfilled account, it must fail.",
	// 		input: &Distance{
	// 			ID: MyUniqueID,
	// 			Payer: &Customer{
	// 				Type:               IndividualCustomer,
	// 				Document:           "11779537069",
	// 				Name:               "Fulano pagador",
	// 				InstitutionAccount: &InstitutionAccount{},
	// 			},
	// 			Receiver: &Customer{
	// 				Ispb:               BexsISPB,
	// 				Type:               IndividualCustomer,
	// 				Document:           "56985776094",
	// 				Name:               "Fulano recebedor",
	// 				InstitutionAccount: &InstitutionAccount{},
	// 			},
	// 			Value: 123.45,
	// 		},
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"payer.account_agency":    ErrorPayerAgencyInvalid.Error(),
	// 				"payer.account_number":    ErrorPayerAccountNumberInvalid.Error(),
	// 				"receiver.account_agency": ErrorReceiverAgencyInvalid.Error(),
	// 				"receiver.account_number": ErrorReceiverAccountNumberInvalid.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For MANUAL PIX payment with incorrectly filled account, it must fail.",
	// 		input: &Distance{
	// 			ID: MyUniqueID,
	// 			Payer: &Customer{
	// 				Type:     IndividualCustomer,
	// 				Document: "11779537069",
	// 				Name:     "Fulano pagador",
	// 				InstitutionAccount: &InstitutionAccount{
	// 					Type:   InvalidAccount,
	// 					Agency: "1234",
	// 					Number: "12345678",
	// 				},
	// 			},
	// 			Receiver: &Customer{
	// 				Ispb:     BexsISPB,
	// 				Type:     IndividualCustomer,
	// 				Document: "56985776094",
	// 				Name:     "Fulano recebedor",
	// 				InstitutionAccount: &InstitutionAccount{
	// 					Type:   InvalidAccount,
	// 					Agency: "4321",
	// 					Number: "87654321",
	// 				},
	// 			},
	// 			Value: 123.45,
	// 		},
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"payer.account_type":    ErrorPayerAccountTypeInvalid.Error(),
	// 				"receiver.account_type": ErrorReceiverAccountTypeInvalid.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For MANUAL PIX payment with data filled in correctly, it must be successful.",
	// 		input:       createDistance(false, EmptyValue, EmptyValue),
	// 		output:      nil,
	// 	},
	// 	{
	// 		description: "For the payment of PIX PER KEY with the EndToEndID field empty, it must fail.",
	// 		input:       createDistance(false, MyValidEmailKey, EmptyValue),
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"end_to_end_id": ErrorRequiredEndToEndID.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For PIX KEY payment using invalid email as key, it must fail.",
	// 		input:       createDistance(false, "@example.com", MyValidEndToEndID),
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"key": ErrorInvalidEmail.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For PIX KEY payment using invalid phone as key, it must fail.",
	// 		input:       createDistance(false, "+555191234", MyValidEndToEndID),
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"key": ErrorWrongPhoneNumberFormat.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For PIX KEY payment using invalid cpf as key, it must fail.",
	// 		input:       createDistance(false, "11111111111", MyValidEndToEndID),
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"key": ErrorWrongCPFFormat.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For PIX KEY payment using invalid cnpj as key, it must fail.",
	// 		input:       createDistance(false, "22222222222222", MyValidEndToEndID),
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"key": ErrorWrongCNPJFormat.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For PIX KEY payment using invalid evp as key, it must fail.",
	// 		input:       createDistance(false, "abcdefghijklmnopqrstuvwxyz1234567890", MyValidEndToEndID),
	// 		output: fxerrors.NewBusinessErrorWithDetail(
	// 			BadRequestError,
	// 			ValidationMessageFormat,
	// 			map[string]string{
	// 				"key": ErrorInvalidEvpKey.Error(),
	// 			},
	// 		),
	// 	},
	// 	{
	// 		description: "For PIX KEY payment using valid email as key, it must be successful.",
	// 		input:       createDistance(false, MyValidEmailKey, MyValidEndToEndID),
	// 		output:      nil,
	// 	},
	// 	{
	// 		description: "For PIX KEY payment using valid phone as key, it must be successful.",
	// 		input:       createDistance(false, MyValidPhoneKey, MyValidEndToEndID),
	// 		output:      nil,
	// 	},
	// 	{
	// 		description: "For PIX KEY payment using valid cpf as key, it must be successful.",
	// 		input:       createDistance(false, MyValidCpfKey, MyValidEndToEndID),
	// 		output:      nil,
	// 	},
	// 	{
	// 		description: "For PIX KEY payment using valid cnpj as key, it must be successful.",
	// 		input:       createDistance(false, MyValidCnpjKey, MyValidEndToEndID),
	// 		output:      nil,
	// 	},
	// 	{
	// 		description: "For PIX KEY payment using valid evp as key, it must be successful.",
	// 		input:       createDistance(false, MyValidEvpKey, MyValidEndToEndID),
	// 		output:      nil,
	// 	},
	// }

	// for _, v := range *values {
	// 	err := v.input.Validate()
	// 	assert.Equal(t, v.output, err, v.description)
	// }
}

func TestDistanceParse(t *testing.T) {

	// type Output struct {
	// 	pay *Distance
	// 	err error
	// }

	// values := &[]struct {
	// 	description string
	// 	input       *Distance
	// 	output      *Output
	// }{
	// 	{
	// 		description: "For any payment order, you must always complete the same fixed fields successfully.",
	// 		input: &Distance{
	// 			Payer: &Customer{},
	// 		},
	// 		output: &Output{
	// 			pay: &Distance{
	// 				InitiationType: ManualInitiation,
	// 				Payer: &Customer{
	// 					Ispb: BexsISPB,
	// 				},
	// 				Schedule:               "",
	// 				Key:                    "",
	// 				EndToEndID:             "",
	// 				ReceiverConciliationID: "",
	// 				InfoBetweenClients:     "",
	// 			},
	// 			err: nil,
	// 		},
	// 	},
	// }

	// for _, v := range *values {
	// 	err := v.input.Parse()
	// 	assert.Equal(t, v.output.pay, v.input, v.description)
	// 	assert.Equal(t, v.output.err, err, v.description)
	// }
}

func TestDistanceString(t *testing.T) {

	// values := &[]struct {
	// 	description string
	// 	input       *Distance
	// 	output      string
	// }{
	// 	{
	// 		description: "For payment order with default parser disabled, must successfully generate text.",
	// 		input:       createDistance(false, EmptyValue, EmptyValue),
	// 		output:      `{"idReqSistemaCliente":"my_unique_id","tpIniciacao":0,"pagador":{"ispb":0,"tpPessoa":0,"cpfCnpj":11779537069,"nome":"Fulano pagador","tpConta":0,"nrAgencia":"1234","nrConta":"12345678"},"recebedor":{"ispb":13059145,"tpPessoa":0,"cpfCnpj":56985776094,"nome":"Fulano recebedor","tpConta":0,"nrAgencia":"4321","nrConta":"87654321"},"valor":123.45}`,
	// 	},
	// 	{
	// 		description: "For payment order with default parser enabled, must successfully generate text.",
	// 		input:       createDistance(true, EmptyValue, EmptyValue),
	// 		output:      `{"id":"my_unique_id","initiation_type":0,"payer":{"ispb":0,"type":0,"document":"11779537069","name":"Fulano pagador","account_type":0,"account_agency":"1234","account_number":"12345678"},"receiver":{"ispb":13059145,"type":0,"document":"56985776094","name":"Fulano recebedor","account_type":0,"account_agency":"4321","account_number":"87654321"},"schedule":"","value":123.45,"key":"","end_to_end_id":"","receiver_conciliation_id":"","info_between_clients":""}`,
	// 	},
	// }

	// for _, v := range *values {
	// 	blob := v.input.String()
	// 	assert.Equal(t, v.output, blob, v.description)
	// }
}

func TestDistanceMarshalJSON(t *testing.T) {

	// type Output struct {
	// 	blob []byte
	// 	err  error
	// }

	// values := &[]struct {
	// 	description string
	// 	input       *Distance
	// 	output      *Output
	// }{
	// 	{
	// 		description: "For payment order with default parser disabled, must successfully generate blob.",
	// 		input:       createDistance(false, EmptyValue, EmptyValue),
	// 		output: &Output{
	// 			blob: []byte(`{"idReqSistemaCliente":"my_unique_id","tpIniciacao":0,"pagador":{"ispb":0,"tpPessoa":0,"cpfCnpj":11779537069,"nome":"Fulano pagador","tpConta":0,"nrAgencia":"1234","nrConta":"12345678"},"recebedor":{"ispb":13059145,"tpPessoa":0,"cpfCnpj":56985776094,"nome":"Fulano recebedor","tpConta":0,"nrAgencia":"4321","nrConta":"87654321"},"valor":123.45}`),
	// 			err:  nil,
	// 		},
	// 	},
	// 	{
	// 		description: "For payment order with default parser enabled, must successfully generate blob.",
	// 		input:       createDistance(true, EmptyValue, EmptyValue),
	// 		output: &Output{
	// 			blob: []byte(`{"id":"my_unique_id","initiation_type":0,"payer":{"ispb":0,"type":0,"document":"11779537069","name":"Fulano pagador","account_type":0,"account_agency":"1234","account_number":"12345678"},"receiver":{"ispb":13059145,"type":0,"document":"56985776094","name":"Fulano recebedor","account_type":0,"account_agency":"4321","account_number":"87654321"},"schedule":"","value":123.45,"key":"","end_to_end_id":"","receiver_conciliation_id":"","info_between_clients":""}`),
	// 			err:  nil,
	// 		},
	// 	},
	// }

	// for _, v := range *values {
	// 	blob, err := v.input.MarshalJSON()
	// 	assert.Equal(t, string(v.output.blob), string(blob), v.description)
	// 	assert.Equal(t, v.output.err, err, v.description)
	// }
}

func TestDistanceResponseString(t *testing.T) {

	// values := &[]struct {
	// 	description string
	// 	input       *DistanceResponse
	// 	output      string
	// }{
	// 	{
	// 		description: "For payment order response, must successfully generate text.",
	// 		input: &DistanceResponse{
	// 			ID:            MyUniqueID,
	// 			ExternalID:    MyUniqueExternalID,
	// 			EndToEndID:    "",
	// 			OperationTime: "",
	// 		},
	// 		output: `{"id":"my_unique_id","external_id":"my_unique_external_id","end_to_end_id":""}`,
	// 	},
	// }

	// for _, v := range *values {
	// 	blob := v.input.String()
	// 	assert.Equal(t, v.output, blob, v.description)
	// }
}

func TestDistanceResponseUnmarshalJSON(t *testing.T) {

	// type Output struct {
	// 	resp *DistanceResponse
	// 	err  error
	// }

	// values := &[]struct {
	// 	description string
	// 	input       []byte
	// 	output      *Output
	// }{
	// 	{
	// 		description: "For blob the payment request response, you must complete the structure successfully.",
	// 		input:       []byte(`{"idReqSistemaCliente":"my_unique_id","idReqJdPi":"my_unique_external_id","endToEndId":"","dtHrReqJdPi":""}`),
	// 		output: &Output{
	// 			resp: &DistanceResponse{
	// 				ID:            MyUniqueID,
	// 				ExternalID:    MyUniqueExternalID,
	// 				EndToEndID:    "",
	// 				OperationTime: "",
	// 			},
	// 			err: nil,
	// 		},
	// 	},
	// }

	// for _, v := range *values {
	// 	resp := new(DistanceResponse)
	// 	err := resp.UnmarshalJSON(v.input)
	// 	assert.Equal(t, v.output.resp, resp, v.description)
	// 	assert.Equal(t, v.output.err, err, v.description)
	// }
}
