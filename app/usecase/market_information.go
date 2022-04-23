package usecase

import "automatic-trade/app/api"

type MarketInformationUseCase struct {
	marketInformationRequest IMarketInformationRequest
}

type IMarketInformationRequest interface {
	GetMarketInformation() (*[]api.MarketInformation, error)
}

func NewMarketInformationUseCase(marketInformationRequest IMarketInformationRequest) MarketInformationUseCase {
	return MarketInformationUseCase{marketInformationRequest: marketInformationRequest}
}

func (useCase MarketInformationUseCase) GetMarketInformation() (*[]api.MarketInformation, error) {
	marketInformation, err := useCase.marketInformationRequest.GetMarketInformation()
	if err != nil {
		return nil, err
	}

	return marketInformation, nil
}
