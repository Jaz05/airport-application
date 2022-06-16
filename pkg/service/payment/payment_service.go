package payment

import (
	"airport/pkg/service/queries"
	"errors"
	"fmt"
)

func ProcessPayment(cardNumber int64, securityNumber int, expirationDate string) error {

	// call payment api
	cardValidationFetch := queries.FakeFetch(fmt.Sprintf("api/bank/card_number=%d", cardNumber))
	cardPaymentFetch := queries.FakeFetch(fmt.Sprintf("api/payment/card_number=%d", cardNumber))

	_, err := queries.FanInFetch(cardValidationFetch, cardPaymentFetch)
	if err != nil {
		return errors.New("there was an error processing your payment")
	}

	return nil
}
