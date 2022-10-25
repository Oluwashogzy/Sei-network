package main

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/utils"
)

type Config struct {
	ChainID       string                `json:"chain_id"`
	TxsPerBlock   uint64                `json:"txs_per_block"`
	MsgsPerTx     uint64                `json:"msgs_per_tx"`
	Rounds        uint64                `json:"rounds"`
	MessageType   string                `json:"message_type"`
	PriceDistr    NumericDistribution   `json:"price_distribution"`
	QuantityDistr NumericDistribution   `json:"quantity_distribution"`
	MsgTypeDistr  MsgTypeDistribution   `json:"message_type_distribution"`
	ContractDistr ContractDistributions `json:"contract_distribution"`
}

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	// NOTE: this field will be renamed to Codec
	Marshaler codec.Codec
	TxConfig  client.TxConfig
	Amino     *codec.LegacyAmino
}

type NumericDistribution struct {
	Min         sdk.Dec `json:"min"`
	Max         sdk.Dec `json:"max"`
	NumDistinct int64   `json:"number_of_distinct_values"`
}

func (d *NumericDistribution) Sample() sdk.Dec {
	steps := sdk.NewDec(rand.Int63n(d.NumDistinct))
	return d.Min.Add(d.Max.Sub(d.Min).QuoInt64(d.NumDistinct).Mul(steps))
}

type DexMsgTypeDistribution struct {
	LimitOrderPct  sdk.Dec `json:"limit_order_percentage"`
	MarketOrderPct sdk.Dec `json:"market_order_percentage"`
}
type StakingMsgTypeDistribution struct {
	DelegatePct        sdk.Dec `json:"delegate_percentage"`
	UndelegatePct      sdk.Dec `json:"undelegate_percentage"`
	BeginRedelegatePct sdk.Dec `json:"begin_redelegate_percentage"`
}

type TokenFactoryMsgTypeDistribution struct {
	TokenFactoryMintPct sdk.Dec `json:"token_factory_mint_percentage"`
	TokenFactoryBurnPct sdk.Dec `json:"token_factory_burn_percentage"`
}

type MsgTypeDistribution struct {
	Dex     DexMsgTypeDistribution     `json:"dex"`
	Staking StakingMsgTypeDistribution `json:"staking"`
	TokenFactory TokenFactoryMsgTypeDistribution `json:"tokenfactory"`
}

func (d *MsgTypeDistribution) Sample(subset string) string {
	switch subset {
	case "dex":
		if !d.Dex.LimitOrderPct.Add(d.Dex.MarketOrderPct).Equal(sdk.OneDec()) {
			panic("Distribution percentages must add up to 1")
		}
		randNum := sdk.MustNewDecFromStr(fmt.Sprintf("%f", rand.Float64()))
		if randNum.LT(d.Dex.LimitOrderPct) {
			return "limit"
		}
		return "market"
	case "tokenfactory":
		if !d.TokenFactory.TokenFactoryBurnPct.Add(d.TokenFactory.TokenFactoryMintPct).Equal(sdk.OneDec()) {
			panic("Distribution percentages must add up to 1")
		}
		randNum := sdk.MustNewDecFromStr(fmt.Sprintf("%f", rand.Float64()))
		if randNum.LT(d.TokenFactory.TokenFactoryBurnPct) {
			return "burn"
		}
		return "mint"
	case "staking":
		if !d.Staking.DelegatePct.Add(d.Staking.UndelegatePct).Add(d.Staking.BeginRedelegatePct).Equal(sdk.OneDec()) {
			panic("Distribution percentages must add up to 1")
		}
		randNum := sdk.MustNewDecFromStr(fmt.Sprintf("%f", rand.Float64()))
		if randNum.LT(d.Staking.DelegatePct) {
			return "delegate"
		} else if randNum.LT(d.Staking.DelegatePct.Add(d.Staking.UndelegatePct)) {
			return "undelegate"
		}
		return "begin_redelegate"
	default:
		panic(fmt.Sprintf("Invalid subset=%s of messages passed in", subset))
	}
}

type ContractDistributions []ContractDistribution

func (d *ContractDistributions) Sample() string {
	if !utils.Reduce(*d, func(i ContractDistribution, o sdk.Dec) sdk.Dec { return o.Add(i.Percentage) }, sdk.ZeroDec()).Equal(sdk.OneDec()) {
		panic("Distribution percentages must add up to 1")
	}
	randNum := sdk.MustNewDecFromStr(fmt.Sprintf("%f", rand.Float64()))
	cumPct := sdk.ZeroDec()
	for _, dist := range *d {
		cumPct = cumPct.Add(dist.Percentage)
		if randNum.LTE(cumPct) {
			return dist.ContractAddr
		}
	}
	panic("this should never be triggered")
}

type ContractDistribution struct {
	ContractAddr string  `json:"contract_address"`
	Percentage   sdk.Dec `json:"percentage"`
}
