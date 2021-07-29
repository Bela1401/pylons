package config

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	PylonsCoinDenom = "pylon"

	DefaultBasicFee   = 10000
	DefaultPremiumFee = 50000
)

// RequestFieldConfig holds parameters for validating request fields
type RequestFieldConfig struct {
	minNameLength        uint64
	minDescriptionLength uint64
}

func NewRequestFieldConfig() RequestFieldConfig {
	return RequestFieldConfig{
		minNameLength:        8,
		minDescriptionLength: 20,
	}
}

type FeeConfig struct {
	BasicFee   sdk.Coins
	PremiumFee sdk.Coins
}

func NewFeeConfig() FeeConfig {
	return FeeConfig{
		BasicFee:   sdk.Coins{sdk.NewInt64Coin(PylonsCoinDenom, DefaultBasicFee)},
		PremiumFee: sdk.Coins{sdk.NewInt64Coin(PylonsCoinDenom, DefaultPremiumFee)},
	}
}