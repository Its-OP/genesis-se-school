package application

import (
	"btcRate/coin/domain"
	"btcRate/common/application"
	"time"
)

type ICoinClientFactory interface {
	CreateClient() ICoinClient
}

type ICoinClient interface {
	GetRate(currency string, coin string) (*SpotPrice, error)
}

type IValidator[T any] interface {
	Validate(T) error
}

type CoinService struct {
	coinClient        ICoinClient
	coinValidator     IValidator[string]
	currencyValidator IValidator[string]
	logger            application.ILogger
}

type SpotPrice struct {
	Amount    float64
	Timestamp time.Time
}

func NewCoinService(factory ICoinClientFactory, coinValidator IValidator[string], currencyValidator IValidator[string], logger application.ILogger) *CoinService {
	coinClient := factory.CreateClient()

	return &CoinService{coinClient: coinClient, coinValidator: coinValidator, currencyValidator: currencyValidator, logger: logger}
}

func (c *CoinService) GetCurrentRate(currency string, coin string) (*domain.Price, error) {
	err := c.validateParameters(currency, coin)
	if err != nil {
		return nil, err
	}

	price, err := c.coinClient.GetRate(currency, coin)

	if err != nil {
		return nil, err
	}

	c.logger.Info("conversion rate was fetched", "coin", coin, "amount", price.Amount, "currency", currency)

	return &domain.Price{
		Amount:    price.Amount,
		Currency:  currency,
		Timestamp: price.Timestamp,
	}, nil
}

func (c *CoinService) validateParameters(currency string, coin string) error {
	err := c.currencyValidator.Validate(currency)
	if err != nil {
		return err
	}

	err = c.coinValidator.Validate(coin)
	if err != nil {
		return err
	}

	return nil
}
