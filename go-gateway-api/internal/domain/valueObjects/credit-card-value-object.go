package valueObjects

import (
	"errors"
	"time"
)

type CreditCard struct {
	Number          string
	CVV             string
	ExpirationMonth int
	ExpirationYear  int
	CardHolderName  string
}

func NewCreditCard(number, cvv, cardHolderName string, expirationMonth, expirationYear int) (*CreditCard, error) {
	if number == "" || cvv == "" || cardHolderName == "" {
		return nil, errors.New("credit card details cannot be empty")
	}

	if expirationMonth < 1 || expirationMonth > 12 {
		return nil, errors.New("invalid expiration month")
	}

	if expirationYear < time.Now().Year() {
		return nil, errors.New("invalid expiration year")
	}

	if len(number) < 13 || len(number) > 19 {
		return nil, errors.New("invalid credit card number length")
	}

	if len(cvv) < 3 || len(cvv) > 4 {
		return nil, errors.New("invalid CVV length")
	}

	return &CreditCard{
		Number:          number,
		CVV:             cvv,
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationYear,
		CardHolderName:  cardHolderName,
	}, nil
}

func (c *CreditCard) GetLastDigits() string {
	if len(c.Number) < 4 {
		return c.Number
	}
	return c.Number[len(c.Number)-4:]
}
