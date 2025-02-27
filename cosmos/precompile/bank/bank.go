// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package bank

import (
	"context"
	"math/big"

	"cosmossdk.io/core/address"

	"github.com/berachain/polaris/contracts/bindings/cosmos/lib"
	bankgenerated "github.com/berachain/polaris/contracts/bindings/cosmos/precompile/bank"
	cosmlib "github.com/berachain/polaris/cosmos/lib"
	"github.com/berachain/polaris/eth/common"
	ethprecompile "github.com/berachain/polaris/eth/core/precompile"
	"github.com/berachain/polaris/eth/core/vm"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// Contract is the precompile contract for the bank module.
type Contract struct {
	ethprecompile.BaseContract

	addressCodec address.Codec
	msgServer    banktypes.MsgServer
	querier      banktypes.QueryServer
	keeper       bankkeeper.Keeper
}

// NewPrecompileContract returns a new instance of the bank precompile contract.
func NewPrecompileContract(
	ak cosmlib.CodecProvider, ms banktypes.MsgServer, qs banktypes.QueryServer,
) *Contract {
	return &Contract{
		BaseContract: ethprecompile.NewBaseContract(
			bankgenerated.BankModuleMetaData.ABI,
			common.BytesToAddress(authtypes.NewModuleAddress(banktypes.ModuleName)),
		),
		addressCodec: ak.AddressCodec(),
		msgServer:    ms,
		querier:      qs,
	}
}

func (c *Contract) CustomValueDecoders() ethprecompile.ValueDecoders {
	return ethprecompile.ValueDecoders{
		banktypes.AttributeKeySender:    c.ConvertAccAddressFromString,
		banktypes.AttributeKeyRecipient: c.ConvertAccAddressFromString,
		banktypes.AttributeKeySpender:   c.ConvertAccAddressFromString,
		banktypes.AttributeKeyReceiver:  c.ConvertAccAddressFromString,
		banktypes.AttributeKeyMinter:    c.ConvertAccAddressFromString,
		banktypes.AttributeKeyBurner:    c.ConvertAccAddressFromString,
	}
}

// GetBalance implements `getBalance(address,string)` method.
func (c *Contract) GetBalance(
	ctx context.Context,
	accountAddress common.Address,
	denom string,
) (*big.Int, error) {
	accAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, accountAddress)
	if err != nil {
		return nil, err
	}

	res, err := c.querier.Balance(ctx, &banktypes.QueryBalanceRequest{
		Address: accAddr,
		Denom:   denom,
	})
	if err != nil {
		return nil, err
	}

	balance := res.GetBalance().Amount
	return balance.BigInt(), nil
}

// GetAllBalances implements `getAllBalances(address)` method.
func (c *Contract) GetAllBalances(
	ctx context.Context,
	accountAddress common.Address,
) ([]lib.CosmosCoin, error) {
	accAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, accountAddress)
	if err != nil {
		return nil, err
	}

	res, err := c.querier.AllBalances(ctx, &banktypes.QueryAllBalancesRequest{
		Address: accAddr,
	})
	if err != nil {
		return nil, err
	}

	return cosmlib.SdkCoinsToEvmCoins(res.Balances), nil
}

// GetSpendableBalanceByDenom implements `getSpendableBalanceByDenom(address,string)` method.
func (c *Contract) GetSpendableBalance(
	ctx context.Context,
	accountAddress common.Address,
	denom string,
) (*big.Int, error) {
	accAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, accountAddress)
	if err != nil {
		return nil, err
	}

	res, err := c.querier.SpendableBalanceByDenom(ctx, &banktypes.QuerySpendableBalanceByDenomRequest{
		Address: accAddr,
		Denom:   denom,
	})
	if err != nil {
		return nil, err
	}

	balance := res.GetBalance().Amount
	return balance.BigInt(), nil
}

// GetSpendableBalances implements `getAllSpendableBalances(address)` method.
func (c *Contract) GetAllSpendableBalances(
	ctx context.Context,
	accountAddress common.Address,
) ([]lib.CosmosCoin, error) {
	accAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, accountAddress)
	if err != nil {
		return nil, err
	}

	res, err := c.querier.SpendableBalances(ctx, &banktypes.QuerySpendableBalancesRequest{
		Address: accAddr,
	})
	if err != nil {
		return nil, err
	}

	return cosmlib.SdkCoinsToEvmCoins(res.Balances), nil
}

// GetSupplyOf implements `getSupply(string)` method.
func (c *Contract) GetSupply(
	ctx context.Context,
	denom string,
) (*big.Int, error) {
	res, err := c.querier.SupplyOf(ctx, &banktypes.QuerySupplyOfRequest{
		Denom: denom,
	})
	if err != nil {
		return nil, err
	}

	supply := res.GetAmount().Amount
	return supply.BigInt(), nil
}

// GetTotalSupply implements `getAllSupply()` method.
func (c *Contract) GetAllSupply(
	ctx context.Context,
) ([]lib.CosmosCoin, error) {
	// todo: add pagination here
	res, err := c.querier.TotalSupply(ctx, &banktypes.QueryTotalSupplyRequest{})
	if err != nil {
		return nil, err
	}

	return cosmlib.SdkCoinsToEvmCoins(res.Supply), nil
}

// Mint new eth method. Needs to be restricted to prove pre-compile.
func (c *Contract) Mint(
	ctx context.Context,
	toAddress common.Address,
	coins any,
) (bool, error) {
	amount, err := cosmlib.ExtractCoinsFromInput(coins)
	if err != nil {
		return false, err
	}
	caller, err := cosmlib.StringFromEthAddress(
		c.addressCodec, vm.UnwrapPolarContext(ctx).MsgSender(),
	)
	if err != nil {
		return false, err
	}

	// NOTE: Hardcoded the prover contract address to make sure only the
	// prover contract can call this method. First deployed contract after network init.
	proverContract := "0x18Df82C7E422A42D47345Ed86B0E935E9718eBda"

	if caller != proverContract {
		return false, nil // TODO: return a proper error
	}

	toAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, toAddress)
	if err != nil {
		return false, err
	}

	// NOTE: defined in /e2e/testapp/entrypoint.sh
	mintStash := "cosmos1yrene6g2zwjttemf0c65fscg8w8c55w58yh8rl"

	// NOTE: using msgSend is a hack, since checkTx is skipped where the sender is checked
	// therefore we can send from anyone's account here, but need an account with eth
	// to send it from
	_, err = c.msgServer.Send(ctx, &banktypes.MsgSend{
		FromAddress: mintStash,
		ToAddress:   toAddr,
		Amount:      amount,
	})
	return err == nil, err
}

// Send implements `send(address,(uint256,string)[])` method.
func (c *Contract) Send(
	ctx context.Context,
	toAddress common.Address,
	coins any,
) (bool, error) {
	amount, err := cosmlib.ExtractCoinsFromInput(coins)
	if err != nil {
		return false, err
	}
	caller, err := cosmlib.StringFromEthAddress(
		c.addressCodec, vm.UnwrapPolarContext(ctx).MsgSender(),
	)
	if err != nil {
		return false, err
	}
	toAddr, err := cosmlib.StringFromEthAddress(c.addressCodec, toAddress)
	if err != nil {
		return false, err
	}

	_, err = c.msgServer.Send(ctx, &banktypes.MsgSend{
		FromAddress: caller,
		ToAddress:   toAddr,
		Amount:      amount,
	})
	return err == nil, err
}

// ConvertAccAddressFromString converts a Cosmos string representing a account address to a
// common.Address.
func (c *Contract) ConvertAccAddressFromString(attributeValue string) (any, error) {
	// extract the sdk.AccAddress from string value as common.Address
	return cosmlib.EthAddressFromString(c.addressCodec, attributeValue)
}
