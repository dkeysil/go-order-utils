// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchange

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ExchangeInitParams is an auto generated low-level Go binding around an user-defined struct.
type ExchangeInitParams struct {
	Admin               common.Address
	Collateral          common.Address
	Ctf                 common.Address
	CtfCollateral       common.Address
	OutcomeTokenFactory common.Address
	ProxyFactory        common.Address
	SafeFactory         common.Address
	FeeReceiver         common.Address
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Salt          *big.Int
	Maker         common.Address
	Signer        common.Address
	TokenId       *big.Int
	MakerAmount   *big.Int
	TakerAmount   *big.Int
	Side          uint8
	SignatureType uint8
	Timestamp     *big.Int
	Metadata      [32]byte
	Builder       [32]byte
	Signature     []byte
}

// OrderStatus is an auto generated low-level Go binding around an user-defined struct.
type OrderStatus struct {
	Filled    bool
	Remaining *big.Int
}

// CTFExchangeMetaData contains all meta data concerning the CTFExchange contract.
var CTFExchangeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structExchangeInitParams\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ctf\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ctfCollateral\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outcomeTokenFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proxyFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"safeFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"PARENT_COLLECTION_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addAdmin\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOperator\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCtf\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCtfCollateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeReceiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxFeeRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderStatus\",\"inputs\":[{\"name\":\"orderHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structOrderStatus\",\"components\":[{\"name\":\"filled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remaining\",\"type\":\"uint248\",\"internalType\":\"uint248\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOutcomeTokenFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProxyFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProxyImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProxyWalletAddress\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSafeFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSafeImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSafeWalletAddress\",\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structOrder\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"takerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"side\",\"type\":\"uint8\",\"internalType\":\"enumSide\"},{\"name\":\"signatureType\",\"type\":\"uint8\",\"internalType\":\"enumSignatureType\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"builder\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"invalidatePreapprovedOrder\",\"inputs\":[{\"name\":\"orderHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAdmin\",\"inputs\":[{\"name\":\"_usr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOperator\",\"inputs\":[{\"name\":\"_usr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isUserPaused\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"matchOrders\",\"inputs\":[{\"name\":\"conditionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"takerOrder\",\"type\":\"tuple\",\"internalType\":\"structOrder\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"takerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"side\",\"type\":\"uint8\",\"internalType\":\"enumSide\"},{\"name\":\"signatureType\",\"type\":\"uint8\",\"internalType\":\"enumSignatureType\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"builder\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"makerOrders\",\"type\":\"tuple[]\",\"internalType\":\"structOrder[]\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"takerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"side\",\"type\":\"uint8\",\"internalType\":\"enumSide\"},{\"name\":\"signatureType\",\"type\":\"uint8\",\"internalType\":\"enumSignatureType\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"builder\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"takerFillAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makerFillAmounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"takerFeeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makerFeeAmounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onERC1155Received\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"orderStatus\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"filled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"remaining\",\"type\":\"uint248\",\"internalType\":\"uint248\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseTrading\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseUser\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preapproveOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structOrder\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"takerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"side\",\"type\":\"uint8\",\"internalType\":\"enumSide\"},{\"name\":\"signatureType\",\"type\":\"uint8\",\"internalType\":\"enumSignatureType\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"builder\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeOperator\",\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOperatorRole\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxFeeRate\",\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUserPauseBlockInterval\",\"inputs\":[{\"name\":\"_interval\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"unpauseTrading\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseUser\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"userPauseBlockInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"userPausedBlockAt\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cashValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateOrder\",\"inputs\":[{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structOrder\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"takerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"side\",\"type\":\"uint8\",\"internalType\":\"enumSide\"},{\"name\":\"signatureType\",\"type\":\"uint8\",\"internalType\":\"enumSignatureType\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"builder\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateOrderSignature\",\"inputs\":[{\"name\":\"orderHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"order\",\"type\":\"tuple\",\"internalType\":\"structOrder\",\"components\":[{\"name\":\"salt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"makerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"takerAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"side\",\"type\":\"uint8\",\"internalType\":\"enumSide\"},{\"name\":\"signatureType\",\"type\":\"uint8\",\"internalType\":\"enumSignatureType\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadata\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"builder\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FeeCharged\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeReceiverUpdated\",\"inputs\":[{\"name\":\"feeReceiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxFeeRateUpdated\",\"inputs\":[{\"name\":\"maxFeeRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewAdmin\",\"inputs\":[{\"name\":\"newAdminAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewOperator\",\"inputs\":[{\"name\":\"newOperatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderFilled\",\"inputs\":[{\"name\":\"orderHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"maker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"side\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumSide\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"makerAmountFilled\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"takerAmountFilled\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"builder\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"metadata\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderPreapprovalInvalidated\",\"inputs\":[{\"name\":\"orderHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderPreapproved\",\"inputs\":[{\"name\":\"orderHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrdersMatched\",\"inputs\":[{\"name\":\"takerOrderHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"takerOrderMaker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"side\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumSide\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"makerAmountFilled\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"takerAmountFilled\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedAdmin\",\"inputs\":[{\"name\":\"removedAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedOperator\",\"inputs\":[{\"name\":\"removedOperator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TradingPaused\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TradingUnpaused\",\"inputs\":[{\"name\":\"pauser\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserPauseBlockIntervalUpdated\",\"inputs\":[{\"name\":\"oldInterval\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newInterval\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserPaused\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"effectivePauseBlock\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserUnpaused\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ComplementaryFillExceedsTakerFill\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExceedsMaxPauseInterval\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeExceedsMaxRate\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FeeExceedsProceeds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LastAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MakingGtRemaining\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxFeeRateExceedsCeiling\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedArrayLengths\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MismatchedTokenIds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoMakerOrders\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAdmin\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCrossing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderAlreadyFilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Paused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TooLittleTokensReceived\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserAlreadyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserIsPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroMakerAmount\",\"inputs\":[]}]",
}

// CTFExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use CTFExchangeMetaData.ABI instead.
var CTFExchangeABI = CTFExchangeMetaData.ABI

// CTFExchange is an auto generated Go binding around an Ethereum contract.
type CTFExchange struct {
	CTFExchangeCaller     // Read-only binding to the contract
	CTFExchangeTransactor // Write-only binding to the contract
	CTFExchangeFilterer   // Log filterer for contract events
}

// CTFExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CTFExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CTFExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CTFExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CTFExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CTFExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CTFExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CTFExchangeSession struct {
	Contract     *CTFExchange      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CTFExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CTFExchangeCallerSession struct {
	Contract *CTFExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CTFExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CTFExchangeTransactorSession struct {
	Contract     *CTFExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CTFExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CTFExchangeRaw struct {
	Contract *CTFExchange // Generic contract binding to access the raw methods on
}

// CTFExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CTFExchangeCallerRaw struct {
	Contract *CTFExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// CTFExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CTFExchangeTransactorRaw struct {
	Contract *CTFExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCTFExchange creates a new instance of CTFExchange, bound to a specific deployed contract.
func NewCTFExchange(address common.Address, backend bind.ContractBackend) (*CTFExchange, error) {
	contract, err := bindCTFExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CTFExchange{CTFExchangeCaller: CTFExchangeCaller{contract: contract}, CTFExchangeTransactor: CTFExchangeTransactor{contract: contract}, CTFExchangeFilterer: CTFExchangeFilterer{contract: contract}}, nil
}

// NewCTFExchangeCaller creates a new read-only instance of CTFExchange, bound to a specific deployed contract.
func NewCTFExchangeCaller(address common.Address, caller bind.ContractCaller) (*CTFExchangeCaller, error) {
	contract, err := bindCTFExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeCaller{contract: contract}, nil
}

// NewCTFExchangeTransactor creates a new write-only instance of CTFExchange, bound to a specific deployed contract.
func NewCTFExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*CTFExchangeTransactor, error) {
	contract, err := bindCTFExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeTransactor{contract: contract}, nil
}

// NewCTFExchangeFilterer creates a new log filterer instance of CTFExchange, bound to a specific deployed contract.
func NewCTFExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*CTFExchangeFilterer, error) {
	contract, err := bindCTFExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeFilterer{contract: contract}, nil
}

// bindCTFExchange binds a generic wrapper to an already deployed contract.
func bindCTFExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CTFExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CTFExchange *CTFExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CTFExchange.Contract.CTFExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CTFExchange *CTFExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CTFExchange.Contract.CTFExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CTFExchange *CTFExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CTFExchange.Contract.CTFExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CTFExchange *CTFExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CTFExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CTFExchange *CTFExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CTFExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CTFExchange *CTFExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CTFExchange.Contract.contract.Transact(opts, method, params...)
}

// PARENTCOLLECTIONID is a free data retrieval call binding the contract method 0x7afda8b8.
//
// Solidity: function PARENT_COLLECTION_ID() view returns(bytes32)
func (_CTFExchange *CTFExchangeCaller) PARENTCOLLECTIONID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "PARENT_COLLECTION_ID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PARENTCOLLECTIONID is a free data retrieval call binding the contract method 0x7afda8b8.
//
// Solidity: function PARENT_COLLECTION_ID() view returns(bytes32)
func (_CTFExchange *CTFExchangeSession) PARENTCOLLECTIONID() ([32]byte, error) {
	return _CTFExchange.Contract.PARENTCOLLECTIONID(&_CTFExchange.CallOpts)
}

// PARENTCOLLECTIONID is a free data retrieval call binding the contract method 0x7afda8b8.
//
// Solidity: function PARENT_COLLECTION_ID() view returns(bytes32)
func (_CTFExchange *CTFExchangeCallerSession) PARENTCOLLECTIONID() ([32]byte, error) {
	return _CTFExchange.Contract.PARENTCOLLECTIONID(&_CTFExchange.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CTFExchange *CTFExchangeCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CTFExchange *CTFExchangeSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CTFExchange.Contract.Eip712Domain(&_CTFExchange.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_CTFExchange *CTFExchangeCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _CTFExchange.Contract.Eip712Domain(&_CTFExchange.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetCollateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getCollateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_CTFExchange *CTFExchangeSession) GetCollateral() (common.Address, error) {
	return _CTFExchange.Contract.GetCollateral(&_CTFExchange.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetCollateral() (common.Address, error) {
	return _CTFExchange.Contract.GetCollateral(&_CTFExchange.CallOpts)
}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetCtf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getCtf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_CTFExchange *CTFExchangeSession) GetCtf() (common.Address, error) {
	return _CTFExchange.Contract.GetCtf(&_CTFExchange.CallOpts)
}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetCtf() (common.Address, error) {
	return _CTFExchange.Contract.GetCtf(&_CTFExchange.CallOpts)
}

// GetCtfCollateral is a free data retrieval call binding the contract method 0x03cee3df.
//
// Solidity: function getCtfCollateral() view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetCtfCollateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getCtfCollateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCtfCollateral is a free data retrieval call binding the contract method 0x03cee3df.
//
// Solidity: function getCtfCollateral() view returns(address)
func (_CTFExchange *CTFExchangeSession) GetCtfCollateral() (common.Address, error) {
	return _CTFExchange.Contract.GetCtfCollateral(&_CTFExchange.CallOpts)
}

// GetCtfCollateral is a free data retrieval call binding the contract method 0x03cee3df.
//
// Solidity: function getCtfCollateral() view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetCtfCollateral() (common.Address, error) {
	return _CTFExchange.Contract.GetCtfCollateral(&_CTFExchange.CallOpts)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeReceiver is a free data retrieval call binding the contract method 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address)
func (_CTFExchange *CTFExchangeSession) GetFeeReceiver() (common.Address, error) {
	return _CTFExchange.Contract.GetFeeReceiver(&_CTFExchange.CallOpts)
}

// GetFeeReceiver is a free data retrieval call binding the contract method 0xe8a35392.
//
// Solidity: function getFeeReceiver() view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetFeeReceiver() (common.Address, error) {
	return _CTFExchange.Contract.GetFeeReceiver(&_CTFExchange.CallOpts)
}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() view returns(uint256)
func (_CTFExchange *CTFExchangeCaller) GetMaxFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getMaxFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() view returns(uint256)
func (_CTFExchange *CTFExchangeSession) GetMaxFeeRate() (*big.Int, error) {
	return _CTFExchange.Contract.GetMaxFeeRate(&_CTFExchange.CallOpts)
}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() view returns(uint256)
func (_CTFExchange *CTFExchangeCallerSession) GetMaxFeeRate() (*big.Int, error) {
	return _CTFExchange.Contract.GetMaxFeeRate(&_CTFExchange.CallOpts)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint248))
func (_CTFExchange *CTFExchangeCaller) GetOrderStatus(opts *bind.CallOpts, orderHash [32]byte) (OrderStatus, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getOrderStatus", orderHash)

	if err != nil {
		return *new(OrderStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(OrderStatus)).(*OrderStatus)

	return out0, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint248))
func (_CTFExchange *CTFExchangeSession) GetOrderStatus(orderHash [32]byte) (OrderStatus, error) {
	return _CTFExchange.Contract.GetOrderStatus(&_CTFExchange.CallOpts, orderHash)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint248))
func (_CTFExchange *CTFExchangeCallerSession) GetOrderStatus(orderHash [32]byte) (OrderStatus, error) {
	return _CTFExchange.Contract.GetOrderStatus(&_CTFExchange.CallOpts, orderHash)
}

// GetOutcomeTokenFactory is a free data retrieval call binding the contract method 0x29cf67f2.
//
// Solidity: function getOutcomeTokenFactory() view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetOutcomeTokenFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getOutcomeTokenFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOutcomeTokenFactory is a free data retrieval call binding the contract method 0x29cf67f2.
//
// Solidity: function getOutcomeTokenFactory() view returns(address)
func (_CTFExchange *CTFExchangeSession) GetOutcomeTokenFactory() (common.Address, error) {
	return _CTFExchange.Contract.GetOutcomeTokenFactory(&_CTFExchange.CallOpts)
}

// GetOutcomeTokenFactory is a free data retrieval call binding the contract method 0x29cf67f2.
//
// Solidity: function getOutcomeTokenFactory() view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetOutcomeTokenFactory() (common.Address, error) {
	return _CTFExchange.Contract.GetOutcomeTokenFactory(&_CTFExchange.CallOpts)
}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getProxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_CTFExchange *CTFExchangeSession) GetProxyFactory() (common.Address, error) {
	return _CTFExchange.Contract.GetProxyFactory(&_CTFExchange.CallOpts)
}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetProxyFactory() (common.Address, error) {
	return _CTFExchange.Contract.GetProxyFactory(&_CTFExchange.CallOpts)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x90e4b720.
//
// Solidity: function getProxyImplementation() view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetProxyImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getProxyImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x90e4b720.
//
// Solidity: function getProxyImplementation() view returns(address)
func (_CTFExchange *CTFExchangeSession) GetProxyImplementation() (common.Address, error) {
	return _CTFExchange.Contract.GetProxyImplementation(&_CTFExchange.CallOpts)
}

// GetProxyImplementation is a free data retrieval call binding the contract method 0x90e4b720.
//
// Solidity: function getProxyImplementation() view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetProxyImplementation() (common.Address, error) {
	return _CTFExchange.Contract.GetProxyImplementation(&_CTFExchange.CallOpts)
}

// GetProxyWalletAddress is a free data retrieval call binding the contract method 0x58d8b6bb.
//
// Solidity: function getProxyWalletAddress(address _addr) view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetProxyWalletAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getProxyWalletAddress", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyWalletAddress is a free data retrieval call binding the contract method 0x58d8b6bb.
//
// Solidity: function getProxyWalletAddress(address _addr) view returns(address)
func (_CTFExchange *CTFExchangeSession) GetProxyWalletAddress(_addr common.Address) (common.Address, error) {
	return _CTFExchange.Contract.GetProxyWalletAddress(&_CTFExchange.CallOpts, _addr)
}

// GetProxyWalletAddress is a free data retrieval call binding the contract method 0x58d8b6bb.
//
// Solidity: function getProxyWalletAddress(address _addr) view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetProxyWalletAddress(_addr common.Address) (common.Address, error) {
	return _CTFExchange.Contract.GetProxyWalletAddress(&_CTFExchange.CallOpts, _addr)
}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetSafeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getSafeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_CTFExchange *CTFExchangeSession) GetSafeFactory() (common.Address, error) {
	return _CTFExchange.Contract.GetSafeFactory(&_CTFExchange.CallOpts)
}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetSafeFactory() (common.Address, error) {
	return _CTFExchange.Contract.GetSafeFactory(&_CTFExchange.CallOpts)
}

// GetSafeImplementation is a free data retrieval call binding the contract method 0xe3b59000.
//
// Solidity: function getSafeImplementation() view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetSafeImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getSafeImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeImplementation is a free data retrieval call binding the contract method 0xe3b59000.
//
// Solidity: function getSafeImplementation() view returns(address)
func (_CTFExchange *CTFExchangeSession) GetSafeImplementation() (common.Address, error) {
	return _CTFExchange.Contract.GetSafeImplementation(&_CTFExchange.CallOpts)
}

// GetSafeImplementation is a free data retrieval call binding the contract method 0xe3b59000.
//
// Solidity: function getSafeImplementation() view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetSafeImplementation() (common.Address, error) {
	return _CTFExchange.Contract.GetSafeImplementation(&_CTFExchange.CallOpts)
}

// GetSafeWalletAddress is a free data retrieval call binding the contract method 0x70bf48e5.
//
// Solidity: function getSafeWalletAddress(address _addr) view returns(address)
func (_CTFExchange *CTFExchangeCaller) GetSafeWalletAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "getSafeWalletAddress", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeWalletAddress is a free data retrieval call binding the contract method 0x70bf48e5.
//
// Solidity: function getSafeWalletAddress(address _addr) view returns(address)
func (_CTFExchange *CTFExchangeSession) GetSafeWalletAddress(_addr common.Address) (common.Address, error) {
	return _CTFExchange.Contract.GetSafeWalletAddress(&_CTFExchange.CallOpts, _addr)
}

// GetSafeWalletAddress is a free data retrieval call binding the contract method 0x70bf48e5.
//
// Solidity: function getSafeWalletAddress(address _addr) view returns(address)
func (_CTFExchange *CTFExchangeCallerSession) GetSafeWalletAddress(_addr common.Address) (common.Address, error) {
	return _CTFExchange.Contract.GetSafeWalletAddress(&_CTFExchange.CallOpts, _addr)
}

// HashOrder is a free data retrieval call binding the contract method 0x3d861a4d.
//
// Solidity: function hashOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns(bytes32)
func (_CTFExchange *CTFExchangeCaller) HashOrder(opts *bind.CallOpts, order Order) ([32]byte, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x3d861a4d.
//
// Solidity: function hashOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns(bytes32)
func (_CTFExchange *CTFExchangeSession) HashOrder(order Order) ([32]byte, error) {
	return _CTFExchange.Contract.HashOrder(&_CTFExchange.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x3d861a4d.
//
// Solidity: function hashOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns(bytes32)
func (_CTFExchange *CTFExchangeCallerSession) HashOrder(order Order) ([32]byte, error) {
	return _CTFExchange.Contract.HashOrder(&_CTFExchange.CallOpts, order)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _usr) view returns(bool)
func (_CTFExchange *CTFExchangeCaller) IsAdmin(opts *bind.CallOpts, _usr common.Address) (bool, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "isAdmin", _usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _usr) view returns(bool)
func (_CTFExchange *CTFExchangeSession) IsAdmin(_usr common.Address) (bool, error) {
	return _CTFExchange.Contract.IsAdmin(&_CTFExchange.CallOpts, _usr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address _usr) view returns(bool)
func (_CTFExchange *CTFExchangeCallerSession) IsAdmin(_usr common.Address) (bool, error) {
	return _CTFExchange.Contract.IsAdmin(&_CTFExchange.CallOpts, _usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _usr) view returns(bool)
func (_CTFExchange *CTFExchangeCaller) IsOperator(opts *bind.CallOpts, _usr common.Address) (bool, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "isOperator", _usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _usr) view returns(bool)
func (_CTFExchange *CTFExchangeSession) IsOperator(_usr common.Address) (bool, error) {
	return _CTFExchange.Contract.IsOperator(&_CTFExchange.CallOpts, _usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address _usr) view returns(bool)
func (_CTFExchange *CTFExchangeCallerSession) IsOperator(_usr common.Address) (bool, error) {
	return _CTFExchange.Contract.IsOperator(&_CTFExchange.CallOpts, _usr)
}

// IsUserPaused is a free data retrieval call binding the contract method 0x28872101.
//
// Solidity: function isUserPaused(address user) view returns(bool)
func (_CTFExchange *CTFExchangeCaller) IsUserPaused(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "isUserPaused", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserPaused is a free data retrieval call binding the contract method 0x28872101.
//
// Solidity: function isUserPaused(address user) view returns(bool)
func (_CTFExchange *CTFExchangeSession) IsUserPaused(user common.Address) (bool, error) {
	return _CTFExchange.Contract.IsUserPaused(&_CTFExchange.CallOpts, user)
}

// IsUserPaused is a free data retrieval call binding the contract method 0x28872101.
//
// Solidity: function isUserPaused(address user) view returns(bool)
func (_CTFExchange *CTFExchangeCallerSession) IsUserPaused(user common.Address) (bool, error) {
	return _CTFExchange.Contract.IsUserPaused(&_CTFExchange.CallOpts, user)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool filled, uint248 remaining)
func (_CTFExchange *CTFExchangeCaller) OrderStatus(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Filled    bool
	Remaining *big.Int
}, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "orderStatus", arg0)

	outstruct := new(struct {
		Filled    bool
		Remaining *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Filled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Remaining = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool filled, uint248 remaining)
func (_CTFExchange *CTFExchangeSession) OrderStatus(arg0 [32]byte) (struct {
	Filled    bool
	Remaining *big.Int
}, error) {
	return _CTFExchange.Contract.OrderStatus(&_CTFExchange.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool filled, uint248 remaining)
func (_CTFExchange *CTFExchangeCallerSession) OrderStatus(arg0 [32]byte) (struct {
	Filled    bool
	Remaining *big.Int
}, error) {
	return _CTFExchange.Contract.OrderStatus(&_CTFExchange.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CTFExchange *CTFExchangeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CTFExchange *CTFExchangeSession) Paused() (bool, error) {
	return _CTFExchange.Contract.Paused(&_CTFExchange.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CTFExchange *CTFExchangeCallerSession) Paused() (bool, error) {
	return _CTFExchange.Contract.Paused(&_CTFExchange.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_CTFExchange *CTFExchangeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_CTFExchange *CTFExchangeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CTFExchange.Contract.SupportsInterface(&_CTFExchange.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_CTFExchange *CTFExchangeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CTFExchange.Contract.SupportsInterface(&_CTFExchange.CallOpts, interfaceId)
}

// UserPauseBlockInterval is a free data retrieval call binding the contract method 0xe3bf917a.
//
// Solidity: function userPauseBlockInterval() view returns(uint256)
func (_CTFExchange *CTFExchangeCaller) UserPauseBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "userPauseBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPauseBlockInterval is a free data retrieval call binding the contract method 0xe3bf917a.
//
// Solidity: function userPauseBlockInterval() view returns(uint256)
func (_CTFExchange *CTFExchangeSession) UserPauseBlockInterval() (*big.Int, error) {
	return _CTFExchange.Contract.UserPauseBlockInterval(&_CTFExchange.CallOpts)
}

// UserPauseBlockInterval is a free data retrieval call binding the contract method 0xe3bf917a.
//
// Solidity: function userPauseBlockInterval() view returns(uint256)
func (_CTFExchange *CTFExchangeCallerSession) UserPauseBlockInterval() (*big.Int, error) {
	return _CTFExchange.Contract.UserPauseBlockInterval(&_CTFExchange.CallOpts)
}

// UserPausedBlockAt is a free data retrieval call binding the contract method 0x234d81b9.
//
// Solidity: function userPausedBlockAt(address ) view returns(uint256)
func (_CTFExchange *CTFExchangeCaller) UserPausedBlockAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "userPausedBlockAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserPausedBlockAt is a free data retrieval call binding the contract method 0x234d81b9.
//
// Solidity: function userPausedBlockAt(address ) view returns(uint256)
func (_CTFExchange *CTFExchangeSession) UserPausedBlockAt(arg0 common.Address) (*big.Int, error) {
	return _CTFExchange.Contract.UserPausedBlockAt(&_CTFExchange.CallOpts, arg0)
}

// UserPausedBlockAt is a free data retrieval call binding the contract method 0x234d81b9.
//
// Solidity: function userPausedBlockAt(address ) view returns(uint256)
func (_CTFExchange *CTFExchangeCallerSession) UserPausedBlockAt(arg0 common.Address) (*big.Int, error) {
	return _CTFExchange.Contract.UserPausedBlockAt(&_CTFExchange.CallOpts, arg0)
}

// ValidateFee is a free data retrieval call binding the contract method 0x0ffea65d.
//
// Solidity: function validateFee(uint256 fee, uint256 cashValue) view returns()
func (_CTFExchange *CTFExchangeCaller) ValidateFee(opts *bind.CallOpts, fee *big.Int, cashValue *big.Int) error {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "validateFee", fee, cashValue)

	if err != nil {
		return err
	}

	return err

}

// ValidateFee is a free data retrieval call binding the contract method 0x0ffea65d.
//
// Solidity: function validateFee(uint256 fee, uint256 cashValue) view returns()
func (_CTFExchange *CTFExchangeSession) ValidateFee(fee *big.Int, cashValue *big.Int) error {
	return _CTFExchange.Contract.ValidateFee(&_CTFExchange.CallOpts, fee, cashValue)
}

// ValidateFee is a free data retrieval call binding the contract method 0x0ffea65d.
//
// Solidity: function validateFee(uint256 fee, uint256 cashValue) view returns()
func (_CTFExchange *CTFExchangeCallerSession) ValidateFee(fee *big.Int, cashValue *big.Int) error {
	return _CTFExchange.Contract.ValidateFee(&_CTFExchange.CallOpts, fee, cashValue)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x088170cb.
//
// Solidity: function validateOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_CTFExchange *CTFExchangeCaller) ValidateOrder(opts *bind.CallOpts, order Order) error {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "validateOrder", order)

	if err != nil {
		return err
	}

	return err

}

// ValidateOrder is a free data retrieval call binding the contract method 0x088170cb.
//
// Solidity: function validateOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_CTFExchange *CTFExchangeSession) ValidateOrder(order Order) error {
	return _CTFExchange.Contract.ValidateOrder(&_CTFExchange.CallOpts, order)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x088170cb.
//
// Solidity: function validateOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_CTFExchange *CTFExchangeCallerSession) ValidateOrder(order Order) error {
	return _CTFExchange.Contract.ValidateOrder(&_CTFExchange.CallOpts, order)
}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xf5db3e4b.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_CTFExchange *CTFExchangeCaller) ValidateOrderSignature(opts *bind.CallOpts, orderHash [32]byte, order Order) error {
	var out []interface{}
	err := _CTFExchange.contract.Call(opts, &out, "validateOrderSignature", orderHash, order)

	if err != nil {
		return err
	}

	return err

}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xf5db3e4b.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_CTFExchange *CTFExchangeSession) ValidateOrderSignature(orderHash [32]byte, order Order) error {
	return _CTFExchange.Contract.ValidateOrderSignature(&_CTFExchange.CallOpts, orderHash, order)
}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xf5db3e4b.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) view returns()
func (_CTFExchange *CTFExchangeCallerSession) ValidateOrderSignature(orderHash [32]byte, order Order) error {
	return _CTFExchange.Contract.ValidateOrderSignature(&_CTFExchange.CallOpts, orderHash, order)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_CTFExchange *CTFExchangeTransactor) AddAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "addAdmin", _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_CTFExchange *CTFExchangeSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CTFExchange.Contract.AddAdmin(&_CTFExchange.TransactOpts, _admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _admin) returns()
func (_CTFExchange *CTFExchangeTransactorSession) AddAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CTFExchange.Contract.AddAdmin(&_CTFExchange.TransactOpts, _admin)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_CTFExchange *CTFExchangeTransactor) AddOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "addOperator", _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_CTFExchange *CTFExchangeSession) AddOperator(_operator common.Address) (*types.Transaction, error) {
	return _CTFExchange.Contract.AddOperator(&_CTFExchange.TransactOpts, _operator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address _operator) returns()
func (_CTFExchange *CTFExchangeTransactorSession) AddOperator(_operator common.Address) (*types.Transaction, error) {
	return _CTFExchange.Contract.AddOperator(&_CTFExchange.TransactOpts, _operator)
}

// InvalidatePreapprovedOrder is a paid mutator transaction binding the contract method 0x437f1994.
//
// Solidity: function invalidatePreapprovedOrder(bytes32 orderHash) returns()
func (_CTFExchange *CTFExchangeTransactor) InvalidatePreapprovedOrder(opts *bind.TransactOpts, orderHash [32]byte) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "invalidatePreapprovedOrder", orderHash)
}

// InvalidatePreapprovedOrder is a paid mutator transaction binding the contract method 0x437f1994.
//
// Solidity: function invalidatePreapprovedOrder(bytes32 orderHash) returns()
func (_CTFExchange *CTFExchangeSession) InvalidatePreapprovedOrder(orderHash [32]byte) (*types.Transaction, error) {
	return _CTFExchange.Contract.InvalidatePreapprovedOrder(&_CTFExchange.TransactOpts, orderHash)
}

// InvalidatePreapprovedOrder is a paid mutator transaction binding the contract method 0x437f1994.
//
// Solidity: function invalidatePreapprovedOrder(bytes32 orderHash) returns()
func (_CTFExchange *CTFExchangeTransactorSession) InvalidatePreapprovedOrder(orderHash [32]byte) (*types.Transaction, error) {
	return _CTFExchange.Contract.InvalidatePreapprovedOrder(&_CTFExchange.TransactOpts, orderHash)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x3c2b4399.
//
// Solidity: function matchOrders(bytes32 conditionId, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) takerOrder, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 takerFeeAmount, uint256[] makerFeeAmounts) returns()
func (_CTFExchange *CTFExchangeTransactor) MatchOrders(opts *bind.TransactOpts, conditionId [32]byte, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, takerFeeAmount *big.Int, makerFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "matchOrders", conditionId, takerOrder, makerOrders, takerFillAmount, makerFillAmounts, takerFeeAmount, makerFeeAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x3c2b4399.
//
// Solidity: function matchOrders(bytes32 conditionId, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) takerOrder, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 takerFeeAmount, uint256[] makerFeeAmounts) returns()
func (_CTFExchange *CTFExchangeSession) MatchOrders(conditionId [32]byte, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, takerFeeAmount *big.Int, makerFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _CTFExchange.Contract.MatchOrders(&_CTFExchange.TransactOpts, conditionId, takerOrder, makerOrders, takerFillAmount, makerFillAmounts, takerFeeAmount, makerFeeAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0x3c2b4399.
//
// Solidity: function matchOrders(bytes32 conditionId, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) takerOrder, (uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts, uint256 takerFeeAmount, uint256[] makerFeeAmounts) returns()
func (_CTFExchange *CTFExchangeTransactorSession) MatchOrders(conditionId [32]byte, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int, takerFeeAmount *big.Int, makerFeeAmounts []*big.Int) (*types.Transaction, error) {
	return _CTFExchange.Contract.MatchOrders(&_CTFExchange.TransactOpts, conditionId, takerOrder, makerOrders, takerFillAmount, makerFillAmounts, takerFeeAmount, makerFeeAmounts)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_CTFExchange *CTFExchangeTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_CTFExchange *CTFExchangeSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CTFExchange.Contract.OnERC1155BatchReceived(&_CTFExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_CTFExchange *CTFExchangeTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CTFExchange.Contract.OnERC1155BatchReceived(&_CTFExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_CTFExchange *CTFExchangeTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_CTFExchange *CTFExchangeSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CTFExchange.Contract.OnERC1155Received(&_CTFExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_CTFExchange *CTFExchangeTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _CTFExchange.Contract.OnERC1155Received(&_CTFExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_CTFExchange *CTFExchangeTransactor) PauseTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "pauseTrading")
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_CTFExchange *CTFExchangeSession) PauseTrading() (*types.Transaction, error) {
	return _CTFExchange.Contract.PauseTrading(&_CTFExchange.TransactOpts)
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_CTFExchange *CTFExchangeTransactorSession) PauseTrading() (*types.Transaction, error) {
	return _CTFExchange.Contract.PauseTrading(&_CTFExchange.TransactOpts)
}

// PauseUser is a paid mutator transaction binding the contract method 0xc0c3132c.
//
// Solidity: function pauseUser() returns()
func (_CTFExchange *CTFExchangeTransactor) PauseUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "pauseUser")
}

// PauseUser is a paid mutator transaction binding the contract method 0xc0c3132c.
//
// Solidity: function pauseUser() returns()
func (_CTFExchange *CTFExchangeSession) PauseUser() (*types.Transaction, error) {
	return _CTFExchange.Contract.PauseUser(&_CTFExchange.TransactOpts)
}

// PauseUser is a paid mutator transaction binding the contract method 0xc0c3132c.
//
// Solidity: function pauseUser() returns()
func (_CTFExchange *CTFExchangeTransactorSession) PauseUser() (*types.Transaction, error) {
	return _CTFExchange.Contract.PauseUser(&_CTFExchange.TransactOpts)
}

// PreapproveOrder is a paid mutator transaction binding the contract method 0xe3a5ced5.
//
// Solidity: function preapproveOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) returns()
func (_CTFExchange *CTFExchangeTransactor) PreapproveOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "preapproveOrder", order)
}

// PreapproveOrder is a paid mutator transaction binding the contract method 0xe3a5ced5.
//
// Solidity: function preapproveOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) returns()
func (_CTFExchange *CTFExchangeSession) PreapproveOrder(order Order) (*types.Transaction, error) {
	return _CTFExchange.Contract.PreapproveOrder(&_CTFExchange.TransactOpts, order)
}

// PreapproveOrder is a paid mutator transaction binding the contract method 0xe3a5ced5.
//
// Solidity: function preapproveOrder((uint256,address,address,uint256,uint256,uint256,uint8,uint8,uint256,bytes32,bytes32,bytes) order) returns()
func (_CTFExchange *CTFExchangeTransactorSession) PreapproveOrder(order Order) (*types.Transaction, error) {
	return _CTFExchange.Contract.PreapproveOrder(&_CTFExchange.TransactOpts, order)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_CTFExchange *CTFExchangeTransactor) RemoveAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "removeAdmin", _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_CTFExchange *CTFExchangeSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CTFExchange.Contract.RemoveAdmin(&_CTFExchange.TransactOpts, _admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _admin) returns()
func (_CTFExchange *CTFExchangeTransactorSession) RemoveAdmin(_admin common.Address) (*types.Transaction, error) {
	return _CTFExchange.Contract.RemoveAdmin(&_CTFExchange.TransactOpts, _admin)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_CTFExchange *CTFExchangeTransactor) RemoveOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "removeOperator", _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_CTFExchange *CTFExchangeSession) RemoveOperator(_operator common.Address) (*types.Transaction, error) {
	return _CTFExchange.Contract.RemoveOperator(&_CTFExchange.TransactOpts, _operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address _operator) returns()
func (_CTFExchange *CTFExchangeTransactorSession) RemoveOperator(_operator common.Address) (*types.Transaction, error) {
	return _CTFExchange.Contract.RemoveOperator(&_CTFExchange.TransactOpts, _operator)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_CTFExchange *CTFExchangeTransactor) RenounceOperatorRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "renounceOperatorRole")
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_CTFExchange *CTFExchangeSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _CTFExchange.Contract.RenounceOperatorRole(&_CTFExchange.TransactOpts)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_CTFExchange *CTFExchangeTransactorSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _CTFExchange.Contract.RenounceOperatorRole(&_CTFExchange.TransactOpts)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address receiver) returns()
func (_CTFExchange *CTFExchangeTransactor) SetFeeReceiver(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "setFeeReceiver", receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address receiver) returns()
func (_CTFExchange *CTFExchangeSession) SetFeeReceiver(receiver common.Address) (*types.Transaction, error) {
	return _CTFExchange.Contract.SetFeeReceiver(&_CTFExchange.TransactOpts, receiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address receiver) returns()
func (_CTFExchange *CTFExchangeTransactorSession) SetFeeReceiver(receiver common.Address) (*types.Transaction, error) {
	return _CTFExchange.Contract.SetFeeReceiver(&_CTFExchange.TransactOpts, receiver)
}

// SetMaxFeeRate is a paid mutator transaction binding the contract method 0x8cda96de.
//
// Solidity: function setMaxFeeRate(uint256 rate) returns()
func (_CTFExchange *CTFExchangeTransactor) SetMaxFeeRate(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "setMaxFeeRate", rate)
}

// SetMaxFeeRate is a paid mutator transaction binding the contract method 0x8cda96de.
//
// Solidity: function setMaxFeeRate(uint256 rate) returns()
func (_CTFExchange *CTFExchangeSession) SetMaxFeeRate(rate *big.Int) (*types.Transaction, error) {
	return _CTFExchange.Contract.SetMaxFeeRate(&_CTFExchange.TransactOpts, rate)
}

// SetMaxFeeRate is a paid mutator transaction binding the contract method 0x8cda96de.
//
// Solidity: function setMaxFeeRate(uint256 rate) returns()
func (_CTFExchange *CTFExchangeTransactorSession) SetMaxFeeRate(rate *big.Int) (*types.Transaction, error) {
	return _CTFExchange.Contract.SetMaxFeeRate(&_CTFExchange.TransactOpts, rate)
}

// SetUserPauseBlockInterval is a paid mutator transaction binding the contract method 0xcd4d8cb4.
//
// Solidity: function setUserPauseBlockInterval(uint256 _interval) returns()
func (_CTFExchange *CTFExchangeTransactor) SetUserPauseBlockInterval(opts *bind.TransactOpts, _interval *big.Int) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "setUserPauseBlockInterval", _interval)
}

// SetUserPauseBlockInterval is a paid mutator transaction binding the contract method 0xcd4d8cb4.
//
// Solidity: function setUserPauseBlockInterval(uint256 _interval) returns()
func (_CTFExchange *CTFExchangeSession) SetUserPauseBlockInterval(_interval *big.Int) (*types.Transaction, error) {
	return _CTFExchange.Contract.SetUserPauseBlockInterval(&_CTFExchange.TransactOpts, _interval)
}

// SetUserPauseBlockInterval is a paid mutator transaction binding the contract method 0xcd4d8cb4.
//
// Solidity: function setUserPauseBlockInterval(uint256 _interval) returns()
func (_CTFExchange *CTFExchangeTransactorSession) SetUserPauseBlockInterval(_interval *big.Int) (*types.Transaction, error) {
	return _CTFExchange.Contract.SetUserPauseBlockInterval(&_CTFExchange.TransactOpts, _interval)
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_CTFExchange *CTFExchangeTransactor) UnpauseTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "unpauseTrading")
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_CTFExchange *CTFExchangeSession) UnpauseTrading() (*types.Transaction, error) {
	return _CTFExchange.Contract.UnpauseTrading(&_CTFExchange.TransactOpts)
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_CTFExchange *CTFExchangeTransactorSession) UnpauseTrading() (*types.Transaction, error) {
	return _CTFExchange.Contract.UnpauseTrading(&_CTFExchange.TransactOpts)
}

// UnpauseUser is a paid mutator transaction binding the contract method 0x4cce30c9.
//
// Solidity: function unpauseUser() returns()
func (_CTFExchange *CTFExchangeTransactor) UnpauseUser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CTFExchange.contract.Transact(opts, "unpauseUser")
}

// UnpauseUser is a paid mutator transaction binding the contract method 0x4cce30c9.
//
// Solidity: function unpauseUser() returns()
func (_CTFExchange *CTFExchangeSession) UnpauseUser() (*types.Transaction, error) {
	return _CTFExchange.Contract.UnpauseUser(&_CTFExchange.TransactOpts)
}

// UnpauseUser is a paid mutator transaction binding the contract method 0x4cce30c9.
//
// Solidity: function unpauseUser() returns()
func (_CTFExchange *CTFExchangeTransactorSession) UnpauseUser() (*types.Transaction, error) {
	return _CTFExchange.Contract.UnpauseUser(&_CTFExchange.TransactOpts)
}

// CTFExchangeFeeChargedIterator is returned from FilterFeeCharged and is used to iterate over the raw logs and unpacked data for FeeCharged events raised by the CTFExchange contract.
type CTFExchangeFeeChargedIterator struct {
	Event *CTFExchangeFeeCharged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeFeeCharged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeFeeCharged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeFeeCharged represents a FeeCharged event raised by the CTFExchange contract.
type CTFExchangeFeeCharged struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeCharged is a free log retrieval operation binding the contract event 0x55bb3cade9d43b798a4fe5ffdd05024b2d7870df53920673bfc7e68047cd0ab1.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 amount)
func (_CTFExchange *CTFExchangeFilterer) FilterFeeCharged(opts *bind.FilterOpts, receiver []common.Address) (*CTFExchangeFeeChargedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "FeeCharged", receiverRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeFeeChargedIterator{contract: _CTFExchange.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

// WatchFeeCharged is a free log subscription operation binding the contract event 0x55bb3cade9d43b798a4fe5ffdd05024b2d7870df53920673bfc7e68047cd0ab1.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 amount)
func (_CTFExchange *CTFExchangeFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *CTFExchangeFeeCharged, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "FeeCharged", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeFeeCharged)
				if err := _CTFExchange.contract.UnpackLog(event, "FeeCharged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeCharged is a log parse operation binding the contract event 0x55bb3cade9d43b798a4fe5ffdd05024b2d7870df53920673bfc7e68047cd0ab1.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 amount)
func (_CTFExchange *CTFExchangeFilterer) ParseFeeCharged(log types.Log) (*CTFExchangeFeeCharged, error) {
	event := new(CTFExchangeFeeCharged)
	if err := _CTFExchange.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeFeeReceiverUpdatedIterator is returned from FilterFeeReceiverUpdated and is used to iterate over the raw logs and unpacked data for FeeReceiverUpdated events raised by the CTFExchange contract.
type CTFExchangeFeeReceiverUpdatedIterator struct {
	Event *CTFExchangeFeeReceiverUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeFeeReceiverUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeFeeReceiverUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeFeeReceiverUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeFeeReceiverUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeFeeReceiverUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeFeeReceiverUpdated represents a FeeReceiverUpdated event raised by the CTFExchange contract.
type CTFExchangeFeeReceiverUpdated struct {
	FeeReceiver common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeeReceiverUpdated is a free log retrieval operation binding the contract event 0x27aae5db36d94179909d019ae0b1ac7c16d96d953148f63c0f6a0a9c8ead79ee.
//
// Solidity: event FeeReceiverUpdated(address indexed feeReceiver)
func (_CTFExchange *CTFExchangeFilterer) FilterFeeReceiverUpdated(opts *bind.FilterOpts, feeReceiver []common.Address) (*CTFExchangeFeeReceiverUpdatedIterator, error) {

	var feeReceiverRule []interface{}
	for _, feeReceiverItem := range feeReceiver {
		feeReceiverRule = append(feeReceiverRule, feeReceiverItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "FeeReceiverUpdated", feeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeFeeReceiverUpdatedIterator{contract: _CTFExchange.contract, event: "FeeReceiverUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeReceiverUpdated is a free log subscription operation binding the contract event 0x27aae5db36d94179909d019ae0b1ac7c16d96d953148f63c0f6a0a9c8ead79ee.
//
// Solidity: event FeeReceiverUpdated(address indexed feeReceiver)
func (_CTFExchange *CTFExchangeFilterer) WatchFeeReceiverUpdated(opts *bind.WatchOpts, sink chan<- *CTFExchangeFeeReceiverUpdated, feeReceiver []common.Address) (event.Subscription, error) {

	var feeReceiverRule []interface{}
	for _, feeReceiverItem := range feeReceiver {
		feeReceiverRule = append(feeReceiverRule, feeReceiverItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "FeeReceiverUpdated", feeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeFeeReceiverUpdated)
				if err := _CTFExchange.contract.UnpackLog(event, "FeeReceiverUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeReceiverUpdated is a log parse operation binding the contract event 0x27aae5db36d94179909d019ae0b1ac7c16d96d953148f63c0f6a0a9c8ead79ee.
//
// Solidity: event FeeReceiverUpdated(address indexed feeReceiver)
func (_CTFExchange *CTFExchangeFilterer) ParseFeeReceiverUpdated(log types.Log) (*CTFExchangeFeeReceiverUpdated, error) {
	event := new(CTFExchangeFeeReceiverUpdated)
	if err := _CTFExchange.contract.UnpackLog(event, "FeeReceiverUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeMaxFeeRateUpdatedIterator is returned from FilterMaxFeeRateUpdated and is used to iterate over the raw logs and unpacked data for MaxFeeRateUpdated events raised by the CTFExchange contract.
type CTFExchangeMaxFeeRateUpdatedIterator struct {
	Event *CTFExchangeMaxFeeRateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeMaxFeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeMaxFeeRateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeMaxFeeRateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeMaxFeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeMaxFeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeMaxFeeRateUpdated represents a MaxFeeRateUpdated event raised by the CTFExchange contract.
type CTFExchangeMaxFeeRateUpdated struct {
	MaxFeeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMaxFeeRateUpdated is a free log retrieval operation binding the contract event 0xe380d7c3967dd06cc7c01db8b17332a1d806fd18f63206dcbd12aaef455c7ff2.
//
// Solidity: event MaxFeeRateUpdated(uint256 maxFeeRate)
func (_CTFExchange *CTFExchangeFilterer) FilterMaxFeeRateUpdated(opts *bind.FilterOpts) (*CTFExchangeMaxFeeRateUpdatedIterator, error) {

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "MaxFeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &CTFExchangeMaxFeeRateUpdatedIterator{contract: _CTFExchange.contract, event: "MaxFeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxFeeRateUpdated is a free log subscription operation binding the contract event 0xe380d7c3967dd06cc7c01db8b17332a1d806fd18f63206dcbd12aaef455c7ff2.
//
// Solidity: event MaxFeeRateUpdated(uint256 maxFeeRate)
func (_CTFExchange *CTFExchangeFilterer) WatchMaxFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *CTFExchangeMaxFeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "MaxFeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeMaxFeeRateUpdated)
				if err := _CTFExchange.contract.UnpackLog(event, "MaxFeeRateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxFeeRateUpdated is a log parse operation binding the contract event 0xe380d7c3967dd06cc7c01db8b17332a1d806fd18f63206dcbd12aaef455c7ff2.
//
// Solidity: event MaxFeeRateUpdated(uint256 maxFeeRate)
func (_CTFExchange *CTFExchangeFilterer) ParseMaxFeeRateUpdated(log types.Log) (*CTFExchangeMaxFeeRateUpdated, error) {
	event := new(CTFExchangeMaxFeeRateUpdated)
	if err := _CTFExchange.contract.UnpackLog(event, "MaxFeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the CTFExchange contract.
type CTFExchangeNewAdminIterator struct {
	Event *CTFExchangeNewAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeNewAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeNewAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeNewAdmin represents a NewAdmin event raised by the CTFExchange contract.
type CTFExchangeNewAdmin struct {
	NewAdminAddress common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) FilterNewAdmin(opts *bind.FilterOpts, newAdminAddress []common.Address, admin []common.Address) (*CTFExchangeNewAdminIterator, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeNewAdminIterator{contract: _CTFExchange.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *CTFExchangeNewAdmin, newAdminAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeNewAdmin)
				if err := _CTFExchange.contract.UnpackLog(event, "NewAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) ParseNewAdmin(log types.Log) (*CTFExchangeNewAdmin, error) {
	event := new(CTFExchangeNewAdmin)
	if err := _CTFExchange.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeNewOperatorIterator is returned from FilterNewOperator and is used to iterate over the raw logs and unpacked data for NewOperator events raised by the CTFExchange contract.
type CTFExchangeNewOperatorIterator struct {
	Event *CTFExchangeNewOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeNewOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeNewOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeNewOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeNewOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeNewOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeNewOperator represents a NewOperator event raised by the CTFExchange contract.
type CTFExchangeNewOperator struct {
	NewOperatorAddress common.Address
	Admin              common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewOperator is a free log retrieval operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) FilterNewOperator(opts *bind.FilterOpts, newOperatorAddress []common.Address, admin []common.Address) (*CTFExchangeNewOperatorIterator, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeNewOperatorIterator{contract: _CTFExchange.contract, event: "NewOperator", logs: logs, sub: sub}, nil
}

// WatchNewOperator is a free log subscription operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) WatchNewOperator(opts *bind.WatchOpts, sink chan<- *CTFExchangeNewOperator, newOperatorAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeNewOperator)
				if err := _CTFExchange.contract.UnpackLog(event, "NewOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOperator is a log parse operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) ParseNewOperator(log types.Log) (*CTFExchangeNewOperator, error) {
	event := new(CTFExchangeNewOperator)
	if err := _CTFExchange.contract.UnpackLog(event, "NewOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the CTFExchange contract.
type CTFExchangeOrderFilledIterator struct {
	Event *CTFExchangeOrderFilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeOrderFilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeOrderFilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeOrderFilled represents a OrderFilled event raised by the CTFExchange contract.
type CTFExchangeOrderFilled struct {
	OrderHash         [32]byte
	Maker             common.Address
	Taker             common.Address
	Side              uint8
	TokenId           *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Fee               *big.Int
	Builder           [32]byte
	Metadata          [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xd543adfd945773f1a62f74f0ee55a5e3b9b1a28262980ba90b1a89f2ea84d8ee.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee, bytes32 builder, bytes32 metadata)
func (_CTFExchange *CTFExchangeFilterer) FilterOrderFilled(opts *bind.FilterOpts, orderHash [][32]byte, maker []common.Address, taker []common.Address) (*CTFExchangeOrderFilledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeOrderFilledIterator{contract: _CTFExchange.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xd543adfd945773f1a62f74f0ee55a5e3b9b1a28262980ba90b1a89f2ea84d8ee.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee, bytes32 builder, bytes32 metadata)
func (_CTFExchange *CTFExchangeFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *CTFExchangeOrderFilled, orderHash [][32]byte, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeOrderFilled)
				if err := _CTFExchange.contract.UnpackLog(event, "OrderFilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderFilled is a log parse operation binding the contract event 0xd543adfd945773f1a62f74f0ee55a5e3b9b1a28262980ba90b1a89f2ea84d8ee.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee, bytes32 builder, bytes32 metadata)
func (_CTFExchange *CTFExchangeFilterer) ParseOrderFilled(log types.Log) (*CTFExchangeOrderFilled, error) {
	event := new(CTFExchangeOrderFilled)
	if err := _CTFExchange.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeOrderPreapprovalInvalidatedIterator is returned from FilterOrderPreapprovalInvalidated and is used to iterate over the raw logs and unpacked data for OrderPreapprovalInvalidated events raised by the CTFExchange contract.
type CTFExchangeOrderPreapprovalInvalidatedIterator struct {
	Event *CTFExchangeOrderPreapprovalInvalidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeOrderPreapprovalInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeOrderPreapprovalInvalidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeOrderPreapprovalInvalidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeOrderPreapprovalInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeOrderPreapprovalInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeOrderPreapprovalInvalidated represents a OrderPreapprovalInvalidated event raised by the CTFExchange contract.
type CTFExchangeOrderPreapprovalInvalidated struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderPreapprovalInvalidated is a free log retrieval operation binding the contract event 0xb766aa470f20b094f26a9a14ea5bf63a60af51703c15776e2e739b6a0428adf6.
//
// Solidity: event OrderPreapprovalInvalidated(bytes32 indexed orderHash)
func (_CTFExchange *CTFExchangeFilterer) FilterOrderPreapprovalInvalidated(opts *bind.FilterOpts, orderHash [][32]byte) (*CTFExchangeOrderPreapprovalInvalidatedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "OrderPreapprovalInvalidated", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeOrderPreapprovalInvalidatedIterator{contract: _CTFExchange.contract, event: "OrderPreapprovalInvalidated", logs: logs, sub: sub}, nil
}

// WatchOrderPreapprovalInvalidated is a free log subscription operation binding the contract event 0xb766aa470f20b094f26a9a14ea5bf63a60af51703c15776e2e739b6a0428adf6.
//
// Solidity: event OrderPreapprovalInvalidated(bytes32 indexed orderHash)
func (_CTFExchange *CTFExchangeFilterer) WatchOrderPreapprovalInvalidated(opts *bind.WatchOpts, sink chan<- *CTFExchangeOrderPreapprovalInvalidated, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "OrderPreapprovalInvalidated", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeOrderPreapprovalInvalidated)
				if err := _CTFExchange.contract.UnpackLog(event, "OrderPreapprovalInvalidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderPreapprovalInvalidated is a log parse operation binding the contract event 0xb766aa470f20b094f26a9a14ea5bf63a60af51703c15776e2e739b6a0428adf6.
//
// Solidity: event OrderPreapprovalInvalidated(bytes32 indexed orderHash)
func (_CTFExchange *CTFExchangeFilterer) ParseOrderPreapprovalInvalidated(log types.Log) (*CTFExchangeOrderPreapprovalInvalidated, error) {
	event := new(CTFExchangeOrderPreapprovalInvalidated)
	if err := _CTFExchange.contract.UnpackLog(event, "OrderPreapprovalInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeOrderPreapprovedIterator is returned from FilterOrderPreapproved and is used to iterate over the raw logs and unpacked data for OrderPreapproved events raised by the CTFExchange contract.
type CTFExchangeOrderPreapprovedIterator struct {
	Event *CTFExchangeOrderPreapproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeOrderPreapprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeOrderPreapproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeOrderPreapproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeOrderPreapprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeOrderPreapprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeOrderPreapproved represents a OrderPreapproved event raised by the CTFExchange contract.
type CTFExchangeOrderPreapproved struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderPreapproved is a free log retrieval operation binding the contract event 0xe92c22722d9c284034b6c9f5aaec018edb3e593c0e084900b6b9d390a1182a0b.
//
// Solidity: event OrderPreapproved(bytes32 indexed orderHash)
func (_CTFExchange *CTFExchangeFilterer) FilterOrderPreapproved(opts *bind.FilterOpts, orderHash [][32]byte) (*CTFExchangeOrderPreapprovedIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "OrderPreapproved", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeOrderPreapprovedIterator{contract: _CTFExchange.contract, event: "OrderPreapproved", logs: logs, sub: sub}, nil
}

// WatchOrderPreapproved is a free log subscription operation binding the contract event 0xe92c22722d9c284034b6c9f5aaec018edb3e593c0e084900b6b9d390a1182a0b.
//
// Solidity: event OrderPreapproved(bytes32 indexed orderHash)
func (_CTFExchange *CTFExchangeFilterer) WatchOrderPreapproved(opts *bind.WatchOpts, sink chan<- *CTFExchangeOrderPreapproved, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "OrderPreapproved", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeOrderPreapproved)
				if err := _CTFExchange.contract.UnpackLog(event, "OrderPreapproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderPreapproved is a log parse operation binding the contract event 0xe92c22722d9c284034b6c9f5aaec018edb3e593c0e084900b6b9d390a1182a0b.
//
// Solidity: event OrderPreapproved(bytes32 indexed orderHash)
func (_CTFExchange *CTFExchangeFilterer) ParseOrderPreapproved(log types.Log) (*CTFExchangeOrderPreapproved, error) {
	event := new(CTFExchangeOrderPreapproved)
	if err := _CTFExchange.contract.UnpackLog(event, "OrderPreapproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the CTFExchange contract.
type CTFExchangeOrdersMatchedIterator struct {
	Event *CTFExchangeOrdersMatched // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeOrdersMatched)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeOrdersMatched)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeOrdersMatched represents a OrdersMatched event raised by the CTFExchange contract.
type CTFExchangeOrdersMatched struct {
	TakerOrderHash    [32]byte
	TakerOrderMaker   common.Address
	Side              uint8
	TokenId           *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x174b3811690657c217184f89418266767c87e4805d09680c39fc9c031c0cab7c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_CTFExchange *CTFExchangeFilterer) FilterOrdersMatched(opts *bind.FilterOpts, takerOrderHash [][32]byte, takerOrderMaker []common.Address) (*CTFExchangeOrdersMatchedIterator, error) {

	var takerOrderHashRule []interface{}
	for _, takerOrderHashItem := range takerOrderHash {
		takerOrderHashRule = append(takerOrderHashRule, takerOrderHashItem)
	}
	var takerOrderMakerRule []interface{}
	for _, takerOrderMakerItem := range takerOrderMaker {
		takerOrderMakerRule = append(takerOrderMakerRule, takerOrderMakerItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "OrdersMatched", takerOrderHashRule, takerOrderMakerRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeOrdersMatchedIterator{contract: _CTFExchange.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x174b3811690657c217184f89418266767c87e4805d09680c39fc9c031c0cab7c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_CTFExchange *CTFExchangeFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *CTFExchangeOrdersMatched, takerOrderHash [][32]byte, takerOrderMaker []common.Address) (event.Subscription, error) {

	var takerOrderHashRule []interface{}
	for _, takerOrderHashItem := range takerOrderHash {
		takerOrderHashRule = append(takerOrderHashRule, takerOrderHashItem)
	}
	var takerOrderMakerRule []interface{}
	for _, takerOrderMakerItem := range takerOrderMaker {
		takerOrderMakerRule = append(takerOrderMakerRule, takerOrderMakerItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "OrdersMatched", takerOrderHashRule, takerOrderMakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeOrdersMatched)
				if err := _CTFExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrdersMatched is a log parse operation binding the contract event 0x174b3811690657c217184f89418266767c87e4805d09680c39fc9c031c0cab7c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint8 side, uint256 tokenId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_CTFExchange *CTFExchangeFilterer) ParseOrdersMatched(log types.Log) (*CTFExchangeOrdersMatched, error) {
	event := new(CTFExchangeOrdersMatched)
	if err := _CTFExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the CTFExchange contract.
type CTFExchangeRemovedAdminIterator struct {
	Event *CTFExchangeRemovedAdmin // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeRemovedAdmin)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeRemovedAdmin)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeRemovedAdmin represents a RemovedAdmin event raised by the CTFExchange contract.
type CTFExchangeRemovedAdmin struct {
	RemovedAdmin common.Address
	Admin        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) FilterRemovedAdmin(opts *bind.FilterOpts, removedAdmin []common.Address, admin []common.Address) (*CTFExchangeRemovedAdminIterator, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeRemovedAdminIterator{contract: _CTFExchange.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *CTFExchangeRemovedAdmin, removedAdmin []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeRemovedAdmin)
				if err := _CTFExchange.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemovedAdmin is a log parse operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) ParseRemovedAdmin(log types.Log) (*CTFExchangeRemovedAdmin, error) {
	event := new(CTFExchangeRemovedAdmin)
	if err := _CTFExchange.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeRemovedOperatorIterator is returned from FilterRemovedOperator and is used to iterate over the raw logs and unpacked data for RemovedOperator events raised by the CTFExchange contract.
type CTFExchangeRemovedOperatorIterator struct {
	Event *CTFExchangeRemovedOperator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeRemovedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeRemovedOperator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeRemovedOperator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeRemovedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeRemovedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeRemovedOperator represents a RemovedOperator event raised by the CTFExchange contract.
type CTFExchangeRemovedOperator struct {
	RemovedOperator common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRemovedOperator is a free log retrieval operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) FilterRemovedOperator(opts *bind.FilterOpts, removedOperator []common.Address, admin []common.Address) (*CTFExchangeRemovedOperatorIterator, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeRemovedOperatorIterator{contract: _CTFExchange.contract, event: "RemovedOperator", logs: logs, sub: sub}, nil
}

// WatchRemovedOperator is a free log subscription operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) WatchRemovedOperator(opts *bind.WatchOpts, sink chan<- *CTFExchangeRemovedOperator, removedOperator []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeRemovedOperator)
				if err := _CTFExchange.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemovedOperator is a log parse operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_CTFExchange *CTFExchangeFilterer) ParseRemovedOperator(log types.Log) (*CTFExchangeRemovedOperator, error) {
	event := new(CTFExchangeRemovedOperator)
	if err := _CTFExchange.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeTradingPausedIterator is returned from FilterTradingPaused and is used to iterate over the raw logs and unpacked data for TradingPaused events raised by the CTFExchange contract.
type CTFExchangeTradingPausedIterator struct {
	Event *CTFExchangeTradingPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeTradingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeTradingPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeTradingPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeTradingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeTradingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeTradingPaused represents a TradingPaused event raised by the CTFExchange contract.
type CTFExchangeTradingPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingPaused is a free log retrieval operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_CTFExchange *CTFExchangeFilterer) FilterTradingPaused(opts *bind.FilterOpts, pauser []common.Address) (*CTFExchangeTradingPausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "TradingPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeTradingPausedIterator{contract: _CTFExchange.contract, event: "TradingPaused", logs: logs, sub: sub}, nil
}

// WatchTradingPaused is a free log subscription operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_CTFExchange *CTFExchangeFilterer) WatchTradingPaused(opts *bind.WatchOpts, sink chan<- *CTFExchangeTradingPaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "TradingPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeTradingPaused)
				if err := _CTFExchange.contract.UnpackLog(event, "TradingPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTradingPaused is a log parse operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_CTFExchange *CTFExchangeFilterer) ParseTradingPaused(log types.Log) (*CTFExchangeTradingPaused, error) {
	event := new(CTFExchangeTradingPaused)
	if err := _CTFExchange.contract.UnpackLog(event, "TradingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeTradingUnpausedIterator is returned from FilterTradingUnpaused and is used to iterate over the raw logs and unpacked data for TradingUnpaused events raised by the CTFExchange contract.
type CTFExchangeTradingUnpausedIterator struct {
	Event *CTFExchangeTradingUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeTradingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeTradingUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeTradingUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeTradingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeTradingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeTradingUnpaused represents a TradingUnpaused event raised by the CTFExchange contract.
type CTFExchangeTradingUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingUnpaused is a free log retrieval operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_CTFExchange *CTFExchangeFilterer) FilterTradingUnpaused(opts *bind.FilterOpts, pauser []common.Address) (*CTFExchangeTradingUnpausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "TradingUnpaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeTradingUnpausedIterator{contract: _CTFExchange.contract, event: "TradingUnpaused", logs: logs, sub: sub}, nil
}

// WatchTradingUnpaused is a free log subscription operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_CTFExchange *CTFExchangeFilterer) WatchTradingUnpaused(opts *bind.WatchOpts, sink chan<- *CTFExchangeTradingUnpaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "TradingUnpaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeTradingUnpaused)
				if err := _CTFExchange.contract.UnpackLog(event, "TradingUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTradingUnpaused is a log parse operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_CTFExchange *CTFExchangeFilterer) ParseTradingUnpaused(log types.Log) (*CTFExchangeTradingUnpaused, error) {
	event := new(CTFExchangeTradingUnpaused)
	if err := _CTFExchange.contract.UnpackLog(event, "TradingUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeUserPauseBlockIntervalUpdatedIterator is returned from FilterUserPauseBlockIntervalUpdated and is used to iterate over the raw logs and unpacked data for UserPauseBlockIntervalUpdated events raised by the CTFExchange contract.
type CTFExchangeUserPauseBlockIntervalUpdatedIterator struct {
	Event *CTFExchangeUserPauseBlockIntervalUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeUserPauseBlockIntervalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeUserPauseBlockIntervalUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeUserPauseBlockIntervalUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeUserPauseBlockIntervalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeUserPauseBlockIntervalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeUserPauseBlockIntervalUpdated represents a UserPauseBlockIntervalUpdated event raised by the CTFExchange contract.
type CTFExchangeUserPauseBlockIntervalUpdated struct {
	OldInterval *big.Int
	NewInterval *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserPauseBlockIntervalUpdated is a free log retrieval operation binding the contract event 0x8c8acf678b7cd311e3b5768c92794d63943684862fdea390856e14d9e2a9ef88.
//
// Solidity: event UserPauseBlockIntervalUpdated(uint256 oldInterval, uint256 newInterval)
func (_CTFExchange *CTFExchangeFilterer) FilterUserPauseBlockIntervalUpdated(opts *bind.FilterOpts) (*CTFExchangeUserPauseBlockIntervalUpdatedIterator, error) {

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "UserPauseBlockIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return &CTFExchangeUserPauseBlockIntervalUpdatedIterator{contract: _CTFExchange.contract, event: "UserPauseBlockIntervalUpdated", logs: logs, sub: sub}, nil
}

// WatchUserPauseBlockIntervalUpdated is a free log subscription operation binding the contract event 0x8c8acf678b7cd311e3b5768c92794d63943684862fdea390856e14d9e2a9ef88.
//
// Solidity: event UserPauseBlockIntervalUpdated(uint256 oldInterval, uint256 newInterval)
func (_CTFExchange *CTFExchangeFilterer) WatchUserPauseBlockIntervalUpdated(opts *bind.WatchOpts, sink chan<- *CTFExchangeUserPauseBlockIntervalUpdated) (event.Subscription, error) {

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "UserPauseBlockIntervalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeUserPauseBlockIntervalUpdated)
				if err := _CTFExchange.contract.UnpackLog(event, "UserPauseBlockIntervalUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserPauseBlockIntervalUpdated is a log parse operation binding the contract event 0x8c8acf678b7cd311e3b5768c92794d63943684862fdea390856e14d9e2a9ef88.
//
// Solidity: event UserPauseBlockIntervalUpdated(uint256 oldInterval, uint256 newInterval)
func (_CTFExchange *CTFExchangeFilterer) ParseUserPauseBlockIntervalUpdated(log types.Log) (*CTFExchangeUserPauseBlockIntervalUpdated, error) {
	event := new(CTFExchangeUserPauseBlockIntervalUpdated)
	if err := _CTFExchange.contract.UnpackLog(event, "UserPauseBlockIntervalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeUserPausedIterator is returned from FilterUserPaused and is used to iterate over the raw logs and unpacked data for UserPaused events raised by the CTFExchange contract.
type CTFExchangeUserPausedIterator struct {
	Event *CTFExchangeUserPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeUserPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeUserPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeUserPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeUserPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeUserPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeUserPaused represents a UserPaused event raised by the CTFExchange contract.
type CTFExchangeUserPaused struct {
	User                common.Address
	EffectivePauseBlock *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUserPaused is a free log retrieval operation binding the contract event 0xa3e76126f19eb25001b29726d2a9502b6377938633d2d6a955107dd442e7a14a.
//
// Solidity: event UserPaused(address indexed user, uint256 effectivePauseBlock)
func (_CTFExchange *CTFExchangeFilterer) FilterUserPaused(opts *bind.FilterOpts, user []common.Address) (*CTFExchangeUserPausedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "UserPaused", userRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeUserPausedIterator{contract: _CTFExchange.contract, event: "UserPaused", logs: logs, sub: sub}, nil
}

// WatchUserPaused is a free log subscription operation binding the contract event 0xa3e76126f19eb25001b29726d2a9502b6377938633d2d6a955107dd442e7a14a.
//
// Solidity: event UserPaused(address indexed user, uint256 effectivePauseBlock)
func (_CTFExchange *CTFExchangeFilterer) WatchUserPaused(opts *bind.WatchOpts, sink chan<- *CTFExchangeUserPaused, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "UserPaused", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeUserPaused)
				if err := _CTFExchange.contract.UnpackLog(event, "UserPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserPaused is a log parse operation binding the contract event 0xa3e76126f19eb25001b29726d2a9502b6377938633d2d6a955107dd442e7a14a.
//
// Solidity: event UserPaused(address indexed user, uint256 effectivePauseBlock)
func (_CTFExchange *CTFExchangeFilterer) ParseUserPaused(log types.Log) (*CTFExchangeUserPaused, error) {
	event := new(CTFExchangeUserPaused)
	if err := _CTFExchange.contract.UnpackLog(event, "UserPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTFExchangeUserUnpausedIterator is returned from FilterUserUnpaused and is used to iterate over the raw logs and unpacked data for UserUnpaused events raised by the CTFExchange contract.
type CTFExchangeUserUnpausedIterator struct {
	Event *CTFExchangeUserUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CTFExchangeUserUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTFExchangeUserUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CTFExchangeUserUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CTFExchangeUserUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTFExchangeUserUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTFExchangeUserUnpaused represents a UserUnpaused event raised by the CTFExchange contract.
type CTFExchangeUserUnpaused struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserUnpaused is a free log retrieval operation binding the contract event 0x1419d4111b5c8636aecff843bf618525f4f8e1aa6898a14357021d68dde8af12.
//
// Solidity: event UserUnpaused(address indexed user)
func (_CTFExchange *CTFExchangeFilterer) FilterUserUnpaused(opts *bind.FilterOpts, user []common.Address) (*CTFExchangeUserUnpausedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CTFExchange.contract.FilterLogs(opts, "UserUnpaused", userRule)
	if err != nil {
		return nil, err
	}
	return &CTFExchangeUserUnpausedIterator{contract: _CTFExchange.contract, event: "UserUnpaused", logs: logs, sub: sub}, nil
}

// WatchUserUnpaused is a free log subscription operation binding the contract event 0x1419d4111b5c8636aecff843bf618525f4f8e1aa6898a14357021d68dde8af12.
//
// Solidity: event UserUnpaused(address indexed user)
func (_CTFExchange *CTFExchangeFilterer) WatchUserUnpaused(opts *bind.WatchOpts, sink chan<- *CTFExchangeUserUnpaused, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _CTFExchange.contract.WatchLogs(opts, "UserUnpaused", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTFExchangeUserUnpaused)
				if err := _CTFExchange.contract.UnpackLog(event, "UserUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserUnpaused is a log parse operation binding the contract event 0x1419d4111b5c8636aecff843bf618525f4f8e1aa6898a14357021d68dde8af12.
//
// Solidity: event UserUnpaused(address indexed user)
func (_CTFExchange *CTFExchangeFilterer) ParseUserUnpaused(log types.Log) (*CTFExchangeUserUnpaused, error) {
	event := new(CTFExchangeUserUnpaused)
	if err := _CTFExchange.contract.UnpackLog(event, "UserUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
