// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// MarketStorageFileMeta is an auto generated low-level Go binding around an user-defined struct.
type MarketStorageFileMeta struct {
	Root *big.Int
	Uri  string
}

// FileMarketMetaData contains all meta data concerning the FileMarket contract.
var FileMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"name\":\"CANCEL_PENALTY_MAX_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CANCEL_PENALTY_MIN_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CHALLENGE_WINDOW_BLOCKS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CancellationPenaltyDistributed\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"penaltyAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nodeCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChallengeSlotsScaled\",\"inputs\":[{\"name\":\"oldCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClientCompensationAccrued\",\"inputs\":[{\"name\":\"client\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClientCompensationBpsUpdated\",\"inputs\":[{\"name\":\"oldBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"ExpiredSlotsProcessed\",\"inputs\":[{\"name\":\"processedCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reporter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"ForcedOrderExits\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"totalFreed\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"KeyLeakReported\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reporter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"function\",\"name\":\"MAX_CHALLENGE_SLOTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_CLIENT_COMPENSATION_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FORCED_EXITS_PER_SWEEP\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PROOF_FAILURE_SLASH_MULTIPLIER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_REPORTER_REWARD_BPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_SWEEP_PER_CALL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_CHALLENGE_SLOTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_PROOF_FAILURE_SLASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"NodeQuit\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeSlashed\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"OnDemandChallengeExpired\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnDemandChallengeIssued\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"deadlineBlock\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnDemandProofSubmitted\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderCancelled\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"refundAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderCompleted\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderFulfilled\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderPlaced\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"numChunks\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"periods\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"replicas\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderUnderReplicated\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"currentFilled\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"desiredReplicas\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProofFailureSlashMultiplierUpdated\",\"inputs\":[{\"name\":\"oldMultiplier\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newMultiplier\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RefundQueued\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RefundWithdrawn\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReporterRewardAccrued\",\"inputs\":[{\"name\":\"reporter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"rewardAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"slashedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReporterRewardBpsUpdated\",\"inputs\":[{\"name\":\"oldBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReporterRewardsClaimed\",\"inputs\":[{\"name\":\"reporter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsCalculated\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"periods\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardsClaimed\",\"inputs\":[{\"name\":\"node\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlashAuthorityUpdated\",\"inputs\":[{\"name\":\"authority\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotChallengeIssued\",\"inputs\":[{\"name\":\"slotIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"orderId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"challengedNode\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"deadlineBlock\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotDeactivated\",\"inputs\":[{\"name\":\"slotIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotExpired\",\"inputs\":[{\"name\":\"slotIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"failedNode\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"slashAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotProofSubmitted\",\"inputs\":[{\"name\":\"slotIndex\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"commitment\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlotsActivated\",\"inputs\":[{\"name\":\"activatedCount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"function\",\"name\":\"activateSlots\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeOrders\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregateActiveEscrow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aggregateActiveWithdrawn\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOrder\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"challengeNode\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"challengeSlots\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengedNode\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"deadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"randomness\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeSlotsInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeStartBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challengeableOrders\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimReporterRewards\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimRewards\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cleanupCursor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clientCompensationBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"completeExpiredOrder\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_extension\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeOrder\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"_commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extension\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fspVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIFspVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"genesisTs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveOrdersCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveOrdersPage\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"orderIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"total\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllSlotInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"orderIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"challengedNodes\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"randomnesses\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"deadlineBlocks\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"isExpired\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChallengeableOrdersCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChallengeableOrdersPage\",\"inputs\":[{\"name\":\"offset\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"orderIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"total\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getClaimableRewards\",\"inputs\":[{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"claimable\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFinancialStats\",\"inputs\":[],\"outputs\":[{\"name\":\"totalContractBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalEscrowHeld\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalRewardsPaid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"averageOrderValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalStakeValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalStats\",\"inputs\":[],\"outputs\":[{\"name\":\"totalOrders\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activeOrdersCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalEscrowLocked\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalNodes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCapacityStaked\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalCapacityUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"activeChallengeSlots\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentPeriod_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentBlock_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengeableOrdersCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeChallengeStatus\",\"inputs\":[{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"activeChallenges\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeEarningsInfo\",\"inputs\":[{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalEarned\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimable\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lastClaimPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOrderEarnings\",\"inputs\":[{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNodeOrders\",\"inputs\":[{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderDetails\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uri_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"root_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"numChunks_\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"periods_\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"replicas_\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"filled_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderEscrowInfo\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"totalEscrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paidToNodes\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remainingEscrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderFinancials\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"escrow_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawn_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startPeriod_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"expired_\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nodes_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderNodes\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProofSystemStats\",\"inputs\":[],\"outputs\":[{\"name\":\"activeSlotsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"idleSlotsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiredSlotsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentBlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengeWindowBlocks\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengeableOrdersCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalSlotsCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecentOrders\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"orderIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"owners\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"numChunks\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"periods\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"},{\"name\":\"replicas\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"},{\"name\":\"filled\",\"type\":\"uint8[]\",\"internalType\":\"uint8[]\"},{\"name\":\"escrows\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"isActive\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReporterEarningsInfo\",\"inputs\":[{\"name\":\"_reporter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"earned\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pending\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlashRedistributionStats\",\"inputs\":[],\"outputs\":[{\"name\":\"totalReceived\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalBurned\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalRewards\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalClientComp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSlotInfo\",\"inputs\":[{\"name\":\"_slotIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challengedNode\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"randomness\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadlineBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isExpired\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalSeedRandomness\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasUnresolvedProofObligation\",\"inputs\":[{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nodeStaking\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_poiVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fspVerifier\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_keyleakVerifier\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChallengeable\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOrderExpired\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyleakVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyLeakVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lifetimeEscrowDeposited\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lifetimeRewardsPaid\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextOrderId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeActiveChallengeCount\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeEarnings\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeLastClaimPeriod\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeOrderEarnings\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeOrderStartTimestamp\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodePendingRewards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeStaking\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractNodeStaking\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeToOrders\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeWithdrawn\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"numChallengeSlots\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onDemandChallenges\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deadlineBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"randomness\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"challenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fileRoot\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orderActiveChallengeCount\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orderEscrowWithdrawn\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orderIndexInActive\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orderIndexInChallengeable\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orderToNodes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orders\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"file\",\"type\":\"tuple\",\"internalType\":\"structMarketStorage.FileMeta\",\"components\":[{\"name\":\"root\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"numChunks\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"periods\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"replicas\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"filled\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"startPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"escrow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingRefunds\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"placeOrder\",\"inputs\":[{\"name\":\"_file\",\"type\":\"tuple\",\"internalType\":\"structMarketStorage.FileMeta\",\"components\":[{\"name\":\"root\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"uri\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"_numChunks\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_periods\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"_replicas\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_pricePerChunkPerPeriod\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fspProof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"}],\"outputs\":[{\"name\":\"orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"poiVerifier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIPoiVerifier\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processExpiredOnDemandChallenge\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processExpiredSlots\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proofFailureSlashMultiplier\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quitOrder\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"reportKeyLeak\",\"inputs\":[{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_proof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reporterEarnings\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reporterPendingRewards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reporterRewardBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reporterWithdrawn\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChallengeStartBlock\",\"inputs\":[{\"name\":\"_block\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setClientCompensationBps\",\"inputs\":[{\"name\":\"_newBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProofFailureSlashMultiplier\",\"inputs\":[{\"name\":\"_newMultiplier\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReporterRewardBps\",\"inputs\":[{\"name\":\"_newBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlashAuthority\",\"inputs\":[{\"name\":\"_authority\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashAuthorities\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slashNode\",\"inputs\":[{\"name\":\"_node\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_slashAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_reason\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitOnDemandProof\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"_commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitProof\",\"inputs\":[{\"name\":\"_slotIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"uint256[4]\",\"internalType\":\"uint256[4]\"},{\"name\":\"_commitment\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sweepCursor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalBurnedFromSlash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalCancellationPenalties\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalClientCompensation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalReporterRewards\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSlashedReceived\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawRefund\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// FileMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use FileMarketMetaData.ABI instead.
var FileMarketABI = FileMarketMetaData.ABI

// FileMarket is an auto generated Go binding around an Ethereum contract.
type FileMarket struct {
	FileMarketCaller     // Read-only binding to the contract
	FileMarketTransactor // Write-only binding to the contract
	FileMarketFilterer   // Log filterer for contract events
}

// FileMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileMarketSession struct {
	Contract     *FileMarket       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileMarketCallerSession struct {
	Contract *FileMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FileMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileMarketTransactorSession struct {
	Contract     *FileMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FileMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileMarketRaw struct {
	Contract *FileMarket // Generic contract binding to access the raw methods on
}

// FileMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileMarketCallerRaw struct {
	Contract *FileMarketCaller // Generic read-only contract binding to access the raw methods on
}

// FileMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileMarketTransactorRaw struct {
	Contract *FileMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileMarket creates a new instance of FileMarket, bound to a specific deployed contract.
func NewFileMarket(address common.Address, backend bind.ContractBackend) (*FileMarket, error) {
	contract, err := bindFileMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FileMarket{FileMarketCaller: FileMarketCaller{contract: contract}, FileMarketTransactor: FileMarketTransactor{contract: contract}, FileMarketFilterer: FileMarketFilterer{contract: contract}}, nil
}

// NewFileMarketCaller creates a new read-only instance of FileMarket, bound to a specific deployed contract.
func NewFileMarketCaller(address common.Address, caller bind.ContractCaller) (*FileMarketCaller, error) {
	contract, err := bindFileMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileMarketCaller{contract: contract}, nil
}

// NewFileMarketTransactor creates a new write-only instance of FileMarket, bound to a specific deployed contract.
func NewFileMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*FileMarketTransactor, error) {
	contract, err := bindFileMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileMarketTransactor{contract: contract}, nil
}

// NewFileMarketFilterer creates a new log filterer instance of FileMarket, bound to a specific deployed contract.
func NewFileMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*FileMarketFilterer, error) {
	contract, err := bindFileMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileMarketFilterer{contract: contract}, nil
}

// bindFileMarket binds a generic wrapper to an already deployed contract.
func bindFileMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FileMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileMarket *FileMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FileMarket.Contract.FileMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileMarket *FileMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileMarket.Contract.FileMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileMarket *FileMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileMarket.Contract.FileMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FileMarket *FileMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FileMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FileMarket *FileMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FileMarket *FileMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FileMarket.Contract.contract.Transact(opts, method, params...)
}

// CANCELPENALTYMAXBPS is a free data retrieval call binding the contract method 0x9fd76abe.
//
// Solidity: function CANCEL_PENALTY_MAX_BPS() view returns(uint256)
func (_FileMarket *FileMarketCaller) CANCELPENALTYMAXBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "CANCEL_PENALTY_MAX_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CANCELPENALTYMAXBPS is a free data retrieval call binding the contract method 0x9fd76abe.
//
// Solidity: function CANCEL_PENALTY_MAX_BPS() view returns(uint256)
func (_FileMarket *FileMarketSession) CANCELPENALTYMAXBPS() (*big.Int, error) {
	return _FileMarket.Contract.CANCELPENALTYMAXBPS(&_FileMarket.CallOpts)
}

// CANCELPENALTYMAXBPS is a free data retrieval call binding the contract method 0x9fd76abe.
//
// Solidity: function CANCEL_PENALTY_MAX_BPS() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) CANCELPENALTYMAXBPS() (*big.Int, error) {
	return _FileMarket.Contract.CANCELPENALTYMAXBPS(&_FileMarket.CallOpts)
}

// CANCELPENALTYMINBPS is a free data retrieval call binding the contract method 0xcd051260.
//
// Solidity: function CANCEL_PENALTY_MIN_BPS() view returns(uint256)
func (_FileMarket *FileMarketCaller) CANCELPENALTYMINBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "CANCEL_PENALTY_MIN_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CANCELPENALTYMINBPS is a free data retrieval call binding the contract method 0xcd051260.
//
// Solidity: function CANCEL_PENALTY_MIN_BPS() view returns(uint256)
func (_FileMarket *FileMarketSession) CANCELPENALTYMINBPS() (*big.Int, error) {
	return _FileMarket.Contract.CANCELPENALTYMINBPS(&_FileMarket.CallOpts)
}

// CANCELPENALTYMINBPS is a free data retrieval call binding the contract method 0xcd051260.
//
// Solidity: function CANCEL_PENALTY_MIN_BPS() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) CANCELPENALTYMINBPS() (*big.Int, error) {
	return _FileMarket.Contract.CANCELPENALTYMINBPS(&_FileMarket.CallOpts)
}

// CHALLENGEWINDOWBLOCKS is a free data retrieval call binding the contract method 0x8a1213ca.
//
// Solidity: function CHALLENGE_WINDOW_BLOCKS() view returns(uint256)
func (_FileMarket *FileMarketCaller) CHALLENGEWINDOWBLOCKS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "CHALLENGE_WINDOW_BLOCKS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CHALLENGEWINDOWBLOCKS is a free data retrieval call binding the contract method 0x8a1213ca.
//
// Solidity: function CHALLENGE_WINDOW_BLOCKS() view returns(uint256)
func (_FileMarket *FileMarketSession) CHALLENGEWINDOWBLOCKS() (*big.Int, error) {
	return _FileMarket.Contract.CHALLENGEWINDOWBLOCKS(&_FileMarket.CallOpts)
}

// CHALLENGEWINDOWBLOCKS is a free data retrieval call binding the contract method 0x8a1213ca.
//
// Solidity: function CHALLENGE_WINDOW_BLOCKS() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) CHALLENGEWINDOWBLOCKS() (*big.Int, error) {
	return _FileMarket.Contract.CHALLENGEWINDOWBLOCKS(&_FileMarket.CallOpts)
}

// MAXCHALLENGESLOTS is a free data retrieval call binding the contract method 0x34b09aa8.
//
// Solidity: function MAX_CHALLENGE_SLOTS() view returns(uint256)
func (_FileMarket *FileMarketCaller) MAXCHALLENGESLOTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "MAX_CHALLENGE_SLOTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCHALLENGESLOTS is a free data retrieval call binding the contract method 0x34b09aa8.
//
// Solidity: function MAX_CHALLENGE_SLOTS() view returns(uint256)
func (_FileMarket *FileMarketSession) MAXCHALLENGESLOTS() (*big.Int, error) {
	return _FileMarket.Contract.MAXCHALLENGESLOTS(&_FileMarket.CallOpts)
}

// MAXCHALLENGESLOTS is a free data retrieval call binding the contract method 0x34b09aa8.
//
// Solidity: function MAX_CHALLENGE_SLOTS() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) MAXCHALLENGESLOTS() (*big.Int, error) {
	return _FileMarket.Contract.MAXCHALLENGESLOTS(&_FileMarket.CallOpts)
}

// MAXCLIENTCOMPENSATIONBPS is a free data retrieval call binding the contract method 0x1fab58cc.
//
// Solidity: function MAX_CLIENT_COMPENSATION_BPS() view returns(uint256)
func (_FileMarket *FileMarketCaller) MAXCLIENTCOMPENSATIONBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "MAX_CLIENT_COMPENSATION_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCLIENTCOMPENSATIONBPS is a free data retrieval call binding the contract method 0x1fab58cc.
//
// Solidity: function MAX_CLIENT_COMPENSATION_BPS() view returns(uint256)
func (_FileMarket *FileMarketSession) MAXCLIENTCOMPENSATIONBPS() (*big.Int, error) {
	return _FileMarket.Contract.MAXCLIENTCOMPENSATIONBPS(&_FileMarket.CallOpts)
}

// MAXCLIENTCOMPENSATIONBPS is a free data retrieval call binding the contract method 0x1fab58cc.
//
// Solidity: function MAX_CLIENT_COMPENSATION_BPS() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) MAXCLIENTCOMPENSATIONBPS() (*big.Int, error) {
	return _FileMarket.Contract.MAXCLIENTCOMPENSATIONBPS(&_FileMarket.CallOpts)
}

// MAXFORCEDEXITSPERSWEEP is a free data retrieval call binding the contract method 0xfa1e5094.
//
// Solidity: function MAX_FORCED_EXITS_PER_SWEEP() view returns(uint256)
func (_FileMarket *FileMarketCaller) MAXFORCEDEXITSPERSWEEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "MAX_FORCED_EXITS_PER_SWEEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFORCEDEXITSPERSWEEP is a free data retrieval call binding the contract method 0xfa1e5094.
//
// Solidity: function MAX_FORCED_EXITS_PER_SWEEP() view returns(uint256)
func (_FileMarket *FileMarketSession) MAXFORCEDEXITSPERSWEEP() (*big.Int, error) {
	return _FileMarket.Contract.MAXFORCEDEXITSPERSWEEP(&_FileMarket.CallOpts)
}

// MAXFORCEDEXITSPERSWEEP is a free data retrieval call binding the contract method 0xfa1e5094.
//
// Solidity: function MAX_FORCED_EXITS_PER_SWEEP() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) MAXFORCEDEXITSPERSWEEP() (*big.Int, error) {
	return _FileMarket.Contract.MAXFORCEDEXITSPERSWEEP(&_FileMarket.CallOpts)
}

// MAXPROOFFAILURESLASHMULTIPLIER is a free data retrieval call binding the contract method 0x15ddb6fe.
//
// Solidity: function MAX_PROOF_FAILURE_SLASH_MULTIPLIER() view returns(uint256)
func (_FileMarket *FileMarketCaller) MAXPROOFFAILURESLASHMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "MAX_PROOF_FAILURE_SLASH_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPROOFFAILURESLASHMULTIPLIER is a free data retrieval call binding the contract method 0x15ddb6fe.
//
// Solidity: function MAX_PROOF_FAILURE_SLASH_MULTIPLIER() view returns(uint256)
func (_FileMarket *FileMarketSession) MAXPROOFFAILURESLASHMULTIPLIER() (*big.Int, error) {
	return _FileMarket.Contract.MAXPROOFFAILURESLASHMULTIPLIER(&_FileMarket.CallOpts)
}

// MAXPROOFFAILURESLASHMULTIPLIER is a free data retrieval call binding the contract method 0x15ddb6fe.
//
// Solidity: function MAX_PROOF_FAILURE_SLASH_MULTIPLIER() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) MAXPROOFFAILURESLASHMULTIPLIER() (*big.Int, error) {
	return _FileMarket.Contract.MAXPROOFFAILURESLASHMULTIPLIER(&_FileMarket.CallOpts)
}

// MAXREPORTERREWARDBPS is a free data retrieval call binding the contract method 0x7f027a77.
//
// Solidity: function MAX_REPORTER_REWARD_BPS() view returns(uint256)
func (_FileMarket *FileMarketCaller) MAXREPORTERREWARDBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "MAX_REPORTER_REWARD_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREPORTERREWARDBPS is a free data retrieval call binding the contract method 0x7f027a77.
//
// Solidity: function MAX_REPORTER_REWARD_BPS() view returns(uint256)
func (_FileMarket *FileMarketSession) MAXREPORTERREWARDBPS() (*big.Int, error) {
	return _FileMarket.Contract.MAXREPORTERREWARDBPS(&_FileMarket.CallOpts)
}

// MAXREPORTERREWARDBPS is a free data retrieval call binding the contract method 0x7f027a77.
//
// Solidity: function MAX_REPORTER_REWARD_BPS() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) MAXREPORTERREWARDBPS() (*big.Int, error) {
	return _FileMarket.Contract.MAXREPORTERREWARDBPS(&_FileMarket.CallOpts)
}

// MAXSWEEPPERCALL is a free data retrieval call binding the contract method 0x54506b97.
//
// Solidity: function MAX_SWEEP_PER_CALL() view returns(uint256)
func (_FileMarket *FileMarketCaller) MAXSWEEPPERCALL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "MAX_SWEEP_PER_CALL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSWEEPPERCALL is a free data retrieval call binding the contract method 0x54506b97.
//
// Solidity: function MAX_SWEEP_PER_CALL() view returns(uint256)
func (_FileMarket *FileMarketSession) MAXSWEEPPERCALL() (*big.Int, error) {
	return _FileMarket.Contract.MAXSWEEPPERCALL(&_FileMarket.CallOpts)
}

// MAXSWEEPPERCALL is a free data retrieval call binding the contract method 0x54506b97.
//
// Solidity: function MAX_SWEEP_PER_CALL() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) MAXSWEEPPERCALL() (*big.Int, error) {
	return _FileMarket.Contract.MAXSWEEPPERCALL(&_FileMarket.CallOpts)
}

// MINCHALLENGESLOTS is a free data retrieval call binding the contract method 0x9fc27649.
//
// Solidity: function MIN_CHALLENGE_SLOTS() view returns(uint256)
func (_FileMarket *FileMarketCaller) MINCHALLENGESLOTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "MIN_CHALLENGE_SLOTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCHALLENGESLOTS is a free data retrieval call binding the contract method 0x9fc27649.
//
// Solidity: function MIN_CHALLENGE_SLOTS() view returns(uint256)
func (_FileMarket *FileMarketSession) MINCHALLENGESLOTS() (*big.Int, error) {
	return _FileMarket.Contract.MINCHALLENGESLOTS(&_FileMarket.CallOpts)
}

// MINCHALLENGESLOTS is a free data retrieval call binding the contract method 0x9fc27649.
//
// Solidity: function MIN_CHALLENGE_SLOTS() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) MINCHALLENGESLOTS() (*big.Int, error) {
	return _FileMarket.Contract.MINCHALLENGESLOTS(&_FileMarket.CallOpts)
}

// MINPROOFFAILURESLASH is a free data retrieval call binding the contract method 0xbab0096a.
//
// Solidity: function MIN_PROOF_FAILURE_SLASH() view returns(uint256)
func (_FileMarket *FileMarketCaller) MINPROOFFAILURESLASH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "MIN_PROOF_FAILURE_SLASH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINPROOFFAILURESLASH is a free data retrieval call binding the contract method 0xbab0096a.
//
// Solidity: function MIN_PROOF_FAILURE_SLASH() view returns(uint256)
func (_FileMarket *FileMarketSession) MINPROOFFAILURESLASH() (*big.Int, error) {
	return _FileMarket.Contract.MINPROOFFAILURESLASH(&_FileMarket.CallOpts)
}

// MINPROOFFAILURESLASH is a free data retrieval call binding the contract method 0xbab0096a.
//
// Solidity: function MIN_PROOF_FAILURE_SLASH() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) MINPROOFFAILURESLASH() (*big.Int, error) {
	return _FileMarket.Contract.MINPROOFFAILURESLASH(&_FileMarket.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FileMarket *FileMarketCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FileMarket *FileMarketSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FileMarket.Contract.UPGRADEINTERFACEVERSION(&_FileMarket.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_FileMarket *FileMarketCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _FileMarket.Contract.UPGRADEINTERFACEVERSION(&_FileMarket.CallOpts)
}

// ActiveOrders is a free data retrieval call binding the contract method 0xc2b5431f.
//
// Solidity: function activeOrders(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCaller) ActiveOrders(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "activeOrders", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveOrders is a free data retrieval call binding the contract method 0xc2b5431f.
//
// Solidity: function activeOrders(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketSession) ActiveOrders(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.ActiveOrders(&_FileMarket.CallOpts, arg0)
}

// ActiveOrders is a free data retrieval call binding the contract method 0xc2b5431f.
//
// Solidity: function activeOrders(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) ActiveOrders(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.ActiveOrders(&_FileMarket.CallOpts, arg0)
}

// AggregateActiveEscrow is a free data retrieval call binding the contract method 0x995a1413.
//
// Solidity: function aggregateActiveEscrow() view returns(uint256)
func (_FileMarket *FileMarketCaller) AggregateActiveEscrow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "aggregateActiveEscrow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AggregateActiveEscrow is a free data retrieval call binding the contract method 0x995a1413.
//
// Solidity: function aggregateActiveEscrow() view returns(uint256)
func (_FileMarket *FileMarketSession) AggregateActiveEscrow() (*big.Int, error) {
	return _FileMarket.Contract.AggregateActiveEscrow(&_FileMarket.CallOpts)
}

// AggregateActiveEscrow is a free data retrieval call binding the contract method 0x995a1413.
//
// Solidity: function aggregateActiveEscrow() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) AggregateActiveEscrow() (*big.Int, error) {
	return _FileMarket.Contract.AggregateActiveEscrow(&_FileMarket.CallOpts)
}

// AggregateActiveWithdrawn is a free data retrieval call binding the contract method 0xd1e95fb1.
//
// Solidity: function aggregateActiveWithdrawn() view returns(uint256)
func (_FileMarket *FileMarketCaller) AggregateActiveWithdrawn(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "aggregateActiveWithdrawn")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AggregateActiveWithdrawn is a free data retrieval call binding the contract method 0xd1e95fb1.
//
// Solidity: function aggregateActiveWithdrawn() view returns(uint256)
func (_FileMarket *FileMarketSession) AggregateActiveWithdrawn() (*big.Int, error) {
	return _FileMarket.Contract.AggregateActiveWithdrawn(&_FileMarket.CallOpts)
}

// AggregateActiveWithdrawn is a free data retrieval call binding the contract method 0xd1e95fb1.
//
// Solidity: function aggregateActiveWithdrawn() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) AggregateActiveWithdrawn() (*big.Int, error) {
	return _FileMarket.Contract.AggregateActiveWithdrawn(&_FileMarket.CallOpts)
}

// ChallengeSlots is a free data retrieval call binding the contract method 0x66319061.
//
// Solidity: function challengeSlots(uint256 ) view returns(uint256 orderId, address challengedNode, uint64 deadlineBlock, uint256 randomness)
func (_FileMarket *FileMarketCaller) ChallengeSlots(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OrderId        *big.Int
	ChallengedNode common.Address
	DeadlineBlock  uint64
	Randomness     *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "challengeSlots", arg0)

	outstruct := new(struct {
		OrderId        *big.Int
		ChallengedNode common.Address
		DeadlineBlock  uint64
		Randomness     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChallengedNode = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.DeadlineBlock = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Randomness = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ChallengeSlots is a free data retrieval call binding the contract method 0x66319061.
//
// Solidity: function challengeSlots(uint256 ) view returns(uint256 orderId, address challengedNode, uint64 deadlineBlock, uint256 randomness)
func (_FileMarket *FileMarketSession) ChallengeSlots(arg0 *big.Int) (struct {
	OrderId        *big.Int
	ChallengedNode common.Address
	DeadlineBlock  uint64
	Randomness     *big.Int
}, error) {
	return _FileMarket.Contract.ChallengeSlots(&_FileMarket.CallOpts, arg0)
}

// ChallengeSlots is a free data retrieval call binding the contract method 0x66319061.
//
// Solidity: function challengeSlots(uint256 ) view returns(uint256 orderId, address challengedNode, uint64 deadlineBlock, uint256 randomness)
func (_FileMarket *FileMarketCallerSession) ChallengeSlots(arg0 *big.Int) (struct {
	OrderId        *big.Int
	ChallengedNode common.Address
	DeadlineBlock  uint64
	Randomness     *big.Int
}, error) {
	return _FileMarket.Contract.ChallengeSlots(&_FileMarket.CallOpts, arg0)
}

// ChallengeSlotsInitialized is a free data retrieval call binding the contract method 0x7c178b38.
//
// Solidity: function challengeSlotsInitialized() view returns(bool)
func (_FileMarket *FileMarketCaller) ChallengeSlotsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "challengeSlotsInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChallengeSlotsInitialized is a free data retrieval call binding the contract method 0x7c178b38.
//
// Solidity: function challengeSlotsInitialized() view returns(bool)
func (_FileMarket *FileMarketSession) ChallengeSlotsInitialized() (bool, error) {
	return _FileMarket.Contract.ChallengeSlotsInitialized(&_FileMarket.CallOpts)
}

// ChallengeSlotsInitialized is a free data retrieval call binding the contract method 0x7c178b38.
//
// Solidity: function challengeSlotsInitialized() view returns(bool)
func (_FileMarket *FileMarketCallerSession) ChallengeSlotsInitialized() (bool, error) {
	return _FileMarket.Contract.ChallengeSlotsInitialized(&_FileMarket.CallOpts)
}

// ChallengeStartBlock is a free data retrieval call binding the contract method 0x4df3830e.
//
// Solidity: function challengeStartBlock() view returns(uint256)
func (_FileMarket *FileMarketCaller) ChallengeStartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "challengeStartBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeStartBlock is a free data retrieval call binding the contract method 0x4df3830e.
//
// Solidity: function challengeStartBlock() view returns(uint256)
func (_FileMarket *FileMarketSession) ChallengeStartBlock() (*big.Int, error) {
	return _FileMarket.Contract.ChallengeStartBlock(&_FileMarket.CallOpts)
}

// ChallengeStartBlock is a free data retrieval call binding the contract method 0x4df3830e.
//
// Solidity: function challengeStartBlock() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) ChallengeStartBlock() (*big.Int, error) {
	return _FileMarket.Contract.ChallengeStartBlock(&_FileMarket.CallOpts)
}

// ChallengeableOrders is a free data retrieval call binding the contract method 0x47960aab.
//
// Solidity: function challengeableOrders(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCaller) ChallengeableOrders(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "challengeableOrders", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeableOrders is a free data retrieval call binding the contract method 0x47960aab.
//
// Solidity: function challengeableOrders(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketSession) ChallengeableOrders(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.ChallengeableOrders(&_FileMarket.CallOpts, arg0)
}

// ChallengeableOrders is a free data retrieval call binding the contract method 0x47960aab.
//
// Solidity: function challengeableOrders(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) ChallengeableOrders(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.ChallengeableOrders(&_FileMarket.CallOpts, arg0)
}

// CleanupCursor is a free data retrieval call binding the contract method 0x7c851766.
//
// Solidity: function cleanupCursor() view returns(uint256)
func (_FileMarket *FileMarketCaller) CleanupCursor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "cleanupCursor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CleanupCursor is a free data retrieval call binding the contract method 0x7c851766.
//
// Solidity: function cleanupCursor() view returns(uint256)
func (_FileMarket *FileMarketSession) CleanupCursor() (*big.Int, error) {
	return _FileMarket.Contract.CleanupCursor(&_FileMarket.CallOpts)
}

// CleanupCursor is a free data retrieval call binding the contract method 0x7c851766.
//
// Solidity: function cleanupCursor() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) CleanupCursor() (*big.Int, error) {
	return _FileMarket.Contract.CleanupCursor(&_FileMarket.CallOpts)
}

// ClientCompensationBps is a free data retrieval call binding the contract method 0x2fae131e.
//
// Solidity: function clientCompensationBps() view returns(uint256)
func (_FileMarket *FileMarketCaller) ClientCompensationBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "clientCompensationBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClientCompensationBps is a free data retrieval call binding the contract method 0x2fae131e.
//
// Solidity: function clientCompensationBps() view returns(uint256)
func (_FileMarket *FileMarketSession) ClientCompensationBps() (*big.Int, error) {
	return _FileMarket.Contract.ClientCompensationBps(&_FileMarket.CallOpts)
}

// ClientCompensationBps is a free data retrieval call binding the contract method 0x2fae131e.
//
// Solidity: function clientCompensationBps() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) ClientCompensationBps() (*big.Int, error) {
	return _FileMarket.Contract.ClientCompensationBps(&_FileMarket.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_FileMarket *FileMarketCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_FileMarket *FileMarketSession) CurrentEpoch() (*big.Int, error) {
	return _FileMarket.Contract.CurrentEpoch(&_FileMarket.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) CurrentEpoch() (*big.Int, error) {
	return _FileMarket.Contract.CurrentEpoch(&_FileMarket.CallOpts)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_FileMarket *FileMarketCaller) CurrentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "currentPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_FileMarket *FileMarketSession) CurrentPeriod() (*big.Int, error) {
	return _FileMarket.Contract.CurrentPeriod(&_FileMarket.CallOpts)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) CurrentPeriod() (*big.Int, error) {
	return _FileMarket.Contract.CurrentPeriod(&_FileMarket.CallOpts)
}

// Extension is a free data retrieval call binding the contract method 0x2d5537b0.
//
// Solidity: function extension() view returns(address)
func (_FileMarket *FileMarketCaller) Extension(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "extension")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Extension is a free data retrieval call binding the contract method 0x2d5537b0.
//
// Solidity: function extension() view returns(address)
func (_FileMarket *FileMarketSession) Extension() (common.Address, error) {
	return _FileMarket.Contract.Extension(&_FileMarket.CallOpts)
}

// Extension is a free data retrieval call binding the contract method 0x2d5537b0.
//
// Solidity: function extension() view returns(address)
func (_FileMarket *FileMarketCallerSession) Extension() (common.Address, error) {
	return _FileMarket.Contract.Extension(&_FileMarket.CallOpts)
}

// FspVerifier is a free data retrieval call binding the contract method 0xd5501099.
//
// Solidity: function fspVerifier() view returns(address)
func (_FileMarket *FileMarketCaller) FspVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "fspVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FspVerifier is a free data retrieval call binding the contract method 0xd5501099.
//
// Solidity: function fspVerifier() view returns(address)
func (_FileMarket *FileMarketSession) FspVerifier() (common.Address, error) {
	return _FileMarket.Contract.FspVerifier(&_FileMarket.CallOpts)
}

// FspVerifier is a free data retrieval call binding the contract method 0xd5501099.
//
// Solidity: function fspVerifier() view returns(address)
func (_FileMarket *FileMarketCallerSession) FspVerifier() (common.Address, error) {
	return _FileMarket.Contract.FspVerifier(&_FileMarket.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_FileMarket *FileMarketCaller) GenesisTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "genesisTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_FileMarket *FileMarketSession) GenesisTs() (*big.Int, error) {
	return _FileMarket.Contract.GenesisTs(&_FileMarket.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) GenesisTs() (*big.Int, error) {
	return _FileMarket.Contract.GenesisTs(&_FileMarket.CallOpts)
}

// GetActiveOrdersCount is a free data retrieval call binding the contract method 0xdadcfd9a.
//
// Solidity: function getActiveOrdersCount() view returns(uint256)
func (_FileMarket *FileMarketCaller) GetActiveOrdersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getActiveOrdersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveOrdersCount is a free data retrieval call binding the contract method 0xdadcfd9a.
//
// Solidity: function getActiveOrdersCount() view returns(uint256)
func (_FileMarket *FileMarketSession) GetActiveOrdersCount() (*big.Int, error) {
	return _FileMarket.Contract.GetActiveOrdersCount(&_FileMarket.CallOpts)
}

// GetActiveOrdersCount is a free data retrieval call binding the contract method 0xdadcfd9a.
//
// Solidity: function getActiveOrdersCount() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) GetActiveOrdersCount() (*big.Int, error) {
	return _FileMarket.Contract.GetActiveOrdersCount(&_FileMarket.CallOpts)
}

// GetActiveOrdersPage is a free data retrieval call binding the contract method 0x429de3f0.
//
// Solidity: function getActiveOrdersPage(uint256 offset, uint256 limit) view returns(uint256[] orderIds, uint256 total)
func (_FileMarket *FileMarketCaller) GetActiveOrdersPage(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	OrderIds []*big.Int
	Total    *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getActiveOrdersPage", offset, limit)

	outstruct := new(struct {
		OrderIds []*big.Int
		Total    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderIds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetActiveOrdersPage is a free data retrieval call binding the contract method 0x429de3f0.
//
// Solidity: function getActiveOrdersPage(uint256 offset, uint256 limit) view returns(uint256[] orderIds, uint256 total)
func (_FileMarket *FileMarketSession) GetActiveOrdersPage(offset *big.Int, limit *big.Int) (struct {
	OrderIds []*big.Int
	Total    *big.Int
}, error) {
	return _FileMarket.Contract.GetActiveOrdersPage(&_FileMarket.CallOpts, offset, limit)
}

// GetActiveOrdersPage is a free data retrieval call binding the contract method 0x429de3f0.
//
// Solidity: function getActiveOrdersPage(uint256 offset, uint256 limit) view returns(uint256[] orderIds, uint256 total)
func (_FileMarket *FileMarketCallerSession) GetActiveOrdersPage(offset *big.Int, limit *big.Int) (struct {
	OrderIds []*big.Int
	Total    *big.Int
}, error) {
	return _FileMarket.Contract.GetActiveOrdersPage(&_FileMarket.CallOpts, offset, limit)
}

// GetAllSlotInfo is a free data retrieval call binding the contract method 0xaa7cb178.
//
// Solidity: function getAllSlotInfo() view returns(uint256[] orderIds, address[] challengedNodes, uint256[] randomnesses, uint256[] deadlineBlocks, bool[] isExpired)
func (_FileMarket *FileMarketCaller) GetAllSlotInfo(opts *bind.CallOpts) (struct {
	OrderIds        []*big.Int
	ChallengedNodes []common.Address
	Randomnesses    []*big.Int
	DeadlineBlocks  []*big.Int
	IsExpired       []bool
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getAllSlotInfo")

	outstruct := new(struct {
		OrderIds        []*big.Int
		ChallengedNodes []common.Address
		Randomnesses    []*big.Int
		DeadlineBlocks  []*big.Int
		IsExpired       []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderIds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.ChallengedNodes = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Randomnesses = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.DeadlineBlocks = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	outstruct.IsExpired = *abi.ConvertType(out[4], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetAllSlotInfo is a free data retrieval call binding the contract method 0xaa7cb178.
//
// Solidity: function getAllSlotInfo() view returns(uint256[] orderIds, address[] challengedNodes, uint256[] randomnesses, uint256[] deadlineBlocks, bool[] isExpired)
func (_FileMarket *FileMarketSession) GetAllSlotInfo() (struct {
	OrderIds        []*big.Int
	ChallengedNodes []common.Address
	Randomnesses    []*big.Int
	DeadlineBlocks  []*big.Int
	IsExpired       []bool
}, error) {
	return _FileMarket.Contract.GetAllSlotInfo(&_FileMarket.CallOpts)
}

// GetAllSlotInfo is a free data retrieval call binding the contract method 0xaa7cb178.
//
// Solidity: function getAllSlotInfo() view returns(uint256[] orderIds, address[] challengedNodes, uint256[] randomnesses, uint256[] deadlineBlocks, bool[] isExpired)
func (_FileMarket *FileMarketCallerSession) GetAllSlotInfo() (struct {
	OrderIds        []*big.Int
	ChallengedNodes []common.Address
	Randomnesses    []*big.Int
	DeadlineBlocks  []*big.Int
	IsExpired       []bool
}, error) {
	return _FileMarket.Contract.GetAllSlotInfo(&_FileMarket.CallOpts)
}

// GetChallengeableOrdersCount is a free data retrieval call binding the contract method 0x4bad3e7e.
//
// Solidity: function getChallengeableOrdersCount() view returns(uint256)
func (_FileMarket *FileMarketCaller) GetChallengeableOrdersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getChallengeableOrdersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChallengeableOrdersCount is a free data retrieval call binding the contract method 0x4bad3e7e.
//
// Solidity: function getChallengeableOrdersCount() view returns(uint256)
func (_FileMarket *FileMarketSession) GetChallengeableOrdersCount() (*big.Int, error) {
	return _FileMarket.Contract.GetChallengeableOrdersCount(&_FileMarket.CallOpts)
}

// GetChallengeableOrdersCount is a free data retrieval call binding the contract method 0x4bad3e7e.
//
// Solidity: function getChallengeableOrdersCount() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) GetChallengeableOrdersCount() (*big.Int, error) {
	return _FileMarket.Contract.GetChallengeableOrdersCount(&_FileMarket.CallOpts)
}

// GetChallengeableOrdersPage is a free data retrieval call binding the contract method 0x4b9b2894.
//
// Solidity: function getChallengeableOrdersPage(uint256 offset, uint256 limit) view returns(uint256[] orderIds, uint256 total)
func (_FileMarket *FileMarketCaller) GetChallengeableOrdersPage(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	OrderIds []*big.Int
	Total    *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getChallengeableOrdersPage", offset, limit)

	outstruct := new(struct {
		OrderIds []*big.Int
		Total    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderIds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Total = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetChallengeableOrdersPage is a free data retrieval call binding the contract method 0x4b9b2894.
//
// Solidity: function getChallengeableOrdersPage(uint256 offset, uint256 limit) view returns(uint256[] orderIds, uint256 total)
func (_FileMarket *FileMarketSession) GetChallengeableOrdersPage(offset *big.Int, limit *big.Int) (struct {
	OrderIds []*big.Int
	Total    *big.Int
}, error) {
	return _FileMarket.Contract.GetChallengeableOrdersPage(&_FileMarket.CallOpts, offset, limit)
}

// GetChallengeableOrdersPage is a free data retrieval call binding the contract method 0x4b9b2894.
//
// Solidity: function getChallengeableOrdersPage(uint256 offset, uint256 limit) view returns(uint256[] orderIds, uint256 total)
func (_FileMarket *FileMarketCallerSession) GetChallengeableOrdersPage(offset *big.Int, limit *big.Int) (struct {
	OrderIds []*big.Int
	Total    *big.Int
}, error) {
	return _FileMarket.Contract.GetChallengeableOrdersPage(&_FileMarket.CallOpts, offset, limit)
}

// GetClaimableRewards is a free data retrieval call binding the contract method 0x308e401e.
//
// Solidity: function getClaimableRewards(address _node) view returns(uint256 claimable)
func (_FileMarket *FileMarketCaller) GetClaimableRewards(opts *bind.CallOpts, _node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getClaimableRewards", _node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableRewards is a free data retrieval call binding the contract method 0x308e401e.
//
// Solidity: function getClaimableRewards(address _node) view returns(uint256 claimable)
func (_FileMarket *FileMarketSession) GetClaimableRewards(_node common.Address) (*big.Int, error) {
	return _FileMarket.Contract.GetClaimableRewards(&_FileMarket.CallOpts, _node)
}

// GetClaimableRewards is a free data retrieval call binding the contract method 0x308e401e.
//
// Solidity: function getClaimableRewards(address _node) view returns(uint256 claimable)
func (_FileMarket *FileMarketCallerSession) GetClaimableRewards(_node common.Address) (*big.Int, error) {
	return _FileMarket.Contract.GetClaimableRewards(&_FileMarket.CallOpts, _node)
}

// GetFinancialStats is a free data retrieval call binding the contract method 0x4143a954.
//
// Solidity: function getFinancialStats() view returns(uint256 totalContractBalance, uint256 totalEscrowHeld, uint256 totalRewardsPaid, uint256 averageOrderValue, uint256 totalStakeValue)
func (_FileMarket *FileMarketCaller) GetFinancialStats(opts *bind.CallOpts) (struct {
	TotalContractBalance *big.Int
	TotalEscrowHeld      *big.Int
	TotalRewardsPaid     *big.Int
	AverageOrderValue    *big.Int
	TotalStakeValue      *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getFinancialStats")

	outstruct := new(struct {
		TotalContractBalance *big.Int
		TotalEscrowHeld      *big.Int
		TotalRewardsPaid     *big.Int
		AverageOrderValue    *big.Int
		TotalStakeValue      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalContractBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalEscrowHeld = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalRewardsPaid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AverageOrderValue = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalStakeValue = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFinancialStats is a free data retrieval call binding the contract method 0x4143a954.
//
// Solidity: function getFinancialStats() view returns(uint256 totalContractBalance, uint256 totalEscrowHeld, uint256 totalRewardsPaid, uint256 averageOrderValue, uint256 totalStakeValue)
func (_FileMarket *FileMarketSession) GetFinancialStats() (struct {
	TotalContractBalance *big.Int
	TotalEscrowHeld      *big.Int
	TotalRewardsPaid     *big.Int
	AverageOrderValue    *big.Int
	TotalStakeValue      *big.Int
}, error) {
	return _FileMarket.Contract.GetFinancialStats(&_FileMarket.CallOpts)
}

// GetFinancialStats is a free data retrieval call binding the contract method 0x4143a954.
//
// Solidity: function getFinancialStats() view returns(uint256 totalContractBalance, uint256 totalEscrowHeld, uint256 totalRewardsPaid, uint256 averageOrderValue, uint256 totalStakeValue)
func (_FileMarket *FileMarketCallerSession) GetFinancialStats() (struct {
	TotalContractBalance *big.Int
	TotalEscrowHeld      *big.Int
	TotalRewardsPaid     *big.Int
	AverageOrderValue    *big.Int
	TotalStakeValue      *big.Int
}, error) {
	return _FileMarket.Contract.GetFinancialStats(&_FileMarket.CallOpts)
}

// GetGlobalStats is a free data retrieval call binding the contract method 0x6b4169c3.
//
// Solidity: function getGlobalStats() view returns(uint256 totalOrders, uint256 activeOrdersCount, uint256 totalEscrowLocked, uint256 totalNodes, uint256 totalCapacityStaked, uint256 totalCapacityUsed, uint256 activeChallengeSlots, uint256 currentPeriod_, uint256 currentBlock_, uint256 challengeableOrdersCount)
func (_FileMarket *FileMarketCaller) GetGlobalStats(opts *bind.CallOpts) (struct {
	TotalOrders              *big.Int
	ActiveOrdersCount        *big.Int
	TotalEscrowLocked        *big.Int
	TotalNodes               *big.Int
	TotalCapacityStaked      *big.Int
	TotalCapacityUsed        *big.Int
	ActiveChallengeSlots     *big.Int
	CurrentPeriod            *big.Int
	CurrentBlock             *big.Int
	ChallengeableOrdersCount *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getGlobalStats")

	outstruct := new(struct {
		TotalOrders              *big.Int
		ActiveOrdersCount        *big.Int
		TotalEscrowLocked        *big.Int
		TotalNodes               *big.Int
		TotalCapacityStaked      *big.Int
		TotalCapacityUsed        *big.Int
		ActiveChallengeSlots     *big.Int
		CurrentPeriod            *big.Int
		CurrentBlock             *big.Int
		ChallengeableOrdersCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalOrders = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ActiveOrdersCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalEscrowLocked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalNodes = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalCapacityStaked = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalCapacityUsed = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ActiveChallengeSlots = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CurrentPeriod = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.CurrentBlock = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.ChallengeableOrdersCount = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGlobalStats is a free data retrieval call binding the contract method 0x6b4169c3.
//
// Solidity: function getGlobalStats() view returns(uint256 totalOrders, uint256 activeOrdersCount, uint256 totalEscrowLocked, uint256 totalNodes, uint256 totalCapacityStaked, uint256 totalCapacityUsed, uint256 activeChallengeSlots, uint256 currentPeriod_, uint256 currentBlock_, uint256 challengeableOrdersCount)
func (_FileMarket *FileMarketSession) GetGlobalStats() (struct {
	TotalOrders              *big.Int
	ActiveOrdersCount        *big.Int
	TotalEscrowLocked        *big.Int
	TotalNodes               *big.Int
	TotalCapacityStaked      *big.Int
	TotalCapacityUsed        *big.Int
	ActiveChallengeSlots     *big.Int
	CurrentPeriod            *big.Int
	CurrentBlock             *big.Int
	ChallengeableOrdersCount *big.Int
}, error) {
	return _FileMarket.Contract.GetGlobalStats(&_FileMarket.CallOpts)
}

// GetGlobalStats is a free data retrieval call binding the contract method 0x6b4169c3.
//
// Solidity: function getGlobalStats() view returns(uint256 totalOrders, uint256 activeOrdersCount, uint256 totalEscrowLocked, uint256 totalNodes, uint256 totalCapacityStaked, uint256 totalCapacityUsed, uint256 activeChallengeSlots, uint256 currentPeriod_, uint256 currentBlock_, uint256 challengeableOrdersCount)
func (_FileMarket *FileMarketCallerSession) GetGlobalStats() (struct {
	TotalOrders              *big.Int
	ActiveOrdersCount        *big.Int
	TotalEscrowLocked        *big.Int
	TotalNodes               *big.Int
	TotalCapacityStaked      *big.Int
	TotalCapacityUsed        *big.Int
	ActiveChallengeSlots     *big.Int
	CurrentPeriod            *big.Int
	CurrentBlock             *big.Int
	ChallengeableOrdersCount *big.Int
}, error) {
	return _FileMarket.Contract.GetGlobalStats(&_FileMarket.CallOpts)
}

// GetNodeChallengeStatus is a free data retrieval call binding the contract method 0xe0fffae2.
//
// Solidity: function getNodeChallengeStatus(address _node) view returns(uint256 activeChallenges)
func (_FileMarket *FileMarketCaller) GetNodeChallengeStatus(opts *bind.CallOpts, _node common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getNodeChallengeStatus", _node)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeChallengeStatus is a free data retrieval call binding the contract method 0xe0fffae2.
//
// Solidity: function getNodeChallengeStatus(address _node) view returns(uint256 activeChallenges)
func (_FileMarket *FileMarketSession) GetNodeChallengeStatus(_node common.Address) (*big.Int, error) {
	return _FileMarket.Contract.GetNodeChallengeStatus(&_FileMarket.CallOpts, _node)
}

// GetNodeChallengeStatus is a free data retrieval call binding the contract method 0xe0fffae2.
//
// Solidity: function getNodeChallengeStatus(address _node) view returns(uint256 activeChallenges)
func (_FileMarket *FileMarketCallerSession) GetNodeChallengeStatus(_node common.Address) (*big.Int, error) {
	return _FileMarket.Contract.GetNodeChallengeStatus(&_FileMarket.CallOpts, _node)
}

// GetNodeEarningsInfo is a free data retrieval call binding the contract method 0xaf50e95e.
//
// Solidity: function getNodeEarningsInfo(address _node) view returns(uint256 totalEarned, uint256 withdrawn, uint256 claimable, uint256 lastClaimPeriod)
func (_FileMarket *FileMarketCaller) GetNodeEarningsInfo(opts *bind.CallOpts, _node common.Address) (struct {
	TotalEarned     *big.Int
	Withdrawn       *big.Int
	Claimable       *big.Int
	LastClaimPeriod *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getNodeEarningsInfo", _node)

	outstruct := new(struct {
		TotalEarned     *big.Int
		Withdrawn       *big.Int
		Claimable       *big.Int
		LastClaimPeriod *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalEarned = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Withdrawn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Claimable = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastClaimPeriod = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNodeEarningsInfo is a free data retrieval call binding the contract method 0xaf50e95e.
//
// Solidity: function getNodeEarningsInfo(address _node) view returns(uint256 totalEarned, uint256 withdrawn, uint256 claimable, uint256 lastClaimPeriod)
func (_FileMarket *FileMarketSession) GetNodeEarningsInfo(_node common.Address) (struct {
	TotalEarned     *big.Int
	Withdrawn       *big.Int
	Claimable       *big.Int
	LastClaimPeriod *big.Int
}, error) {
	return _FileMarket.Contract.GetNodeEarningsInfo(&_FileMarket.CallOpts, _node)
}

// GetNodeEarningsInfo is a free data retrieval call binding the contract method 0xaf50e95e.
//
// Solidity: function getNodeEarningsInfo(address _node) view returns(uint256 totalEarned, uint256 withdrawn, uint256 claimable, uint256 lastClaimPeriod)
func (_FileMarket *FileMarketCallerSession) GetNodeEarningsInfo(_node common.Address) (struct {
	TotalEarned     *big.Int
	Withdrawn       *big.Int
	Claimable       *big.Int
	LastClaimPeriod *big.Int
}, error) {
	return _FileMarket.Contract.GetNodeEarningsInfo(&_FileMarket.CallOpts, _node)
}

// GetNodeOrderEarnings is a free data retrieval call binding the contract method 0xce3bf480.
//
// Solidity: function getNodeOrderEarnings(address _node, uint256 _orderId) view returns(uint256)
func (_FileMarket *FileMarketCaller) GetNodeOrderEarnings(opts *bind.CallOpts, _node common.Address, _orderId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getNodeOrderEarnings", _node, _orderId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOrderEarnings is a free data retrieval call binding the contract method 0xce3bf480.
//
// Solidity: function getNodeOrderEarnings(address _node, uint256 _orderId) view returns(uint256)
func (_FileMarket *FileMarketSession) GetNodeOrderEarnings(_node common.Address, _orderId *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.GetNodeOrderEarnings(&_FileMarket.CallOpts, _node, _orderId)
}

// GetNodeOrderEarnings is a free data retrieval call binding the contract method 0xce3bf480.
//
// Solidity: function getNodeOrderEarnings(address _node, uint256 _orderId) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) GetNodeOrderEarnings(_node common.Address, _orderId *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.GetNodeOrderEarnings(&_FileMarket.CallOpts, _node, _orderId)
}

// GetNodeOrders is a free data retrieval call binding the contract method 0xe1623463.
//
// Solidity: function getNodeOrders(address _node) view returns(uint256[])
func (_FileMarket *FileMarketCaller) GetNodeOrders(opts *bind.CallOpts, _node common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getNodeOrders", _node)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNodeOrders is a free data retrieval call binding the contract method 0xe1623463.
//
// Solidity: function getNodeOrders(address _node) view returns(uint256[])
func (_FileMarket *FileMarketSession) GetNodeOrders(_node common.Address) ([]*big.Int, error) {
	return _FileMarket.Contract.GetNodeOrders(&_FileMarket.CallOpts, _node)
}

// GetNodeOrders is a free data retrieval call binding the contract method 0xe1623463.
//
// Solidity: function getNodeOrders(address _node) view returns(uint256[])
func (_FileMarket *FileMarketCallerSession) GetNodeOrders(_node common.Address) ([]*big.Int, error) {
	return _FileMarket.Contract.GetNodeOrders(&_FileMarket.CallOpts, _node)
}

// GetOrderDetails is a free data retrieval call binding the contract method 0xec7dd7bb.
//
// Solidity: function getOrderDetails(uint256 _orderId) view returns(address owner_, string uri_, uint256 root_, uint32 numChunks_, uint16 periods_, uint8 replicas_, uint8 filled_)
func (_FileMarket *FileMarketCaller) GetOrderDetails(opts *bind.CallOpts, _orderId *big.Int) (struct {
	Owner     common.Address
	Uri       string
	Root      *big.Int
	NumChunks uint32
	Periods   uint16
	Replicas  uint8
	Filled    uint8
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getOrderDetails", _orderId)

	outstruct := new(struct {
		Owner     common.Address
		Uri       string
		Root      *big.Int
		NumChunks uint32
		Periods   uint16
		Replicas  uint8
		Filled    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Uri = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Root = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NumChunks = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Periods = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.Replicas = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Filled = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetOrderDetails is a free data retrieval call binding the contract method 0xec7dd7bb.
//
// Solidity: function getOrderDetails(uint256 _orderId) view returns(address owner_, string uri_, uint256 root_, uint32 numChunks_, uint16 periods_, uint8 replicas_, uint8 filled_)
func (_FileMarket *FileMarketSession) GetOrderDetails(_orderId *big.Int) (struct {
	Owner     common.Address
	Uri       string
	Root      *big.Int
	NumChunks uint32
	Periods   uint16
	Replicas  uint8
	Filled    uint8
}, error) {
	return _FileMarket.Contract.GetOrderDetails(&_FileMarket.CallOpts, _orderId)
}

// GetOrderDetails is a free data retrieval call binding the contract method 0xec7dd7bb.
//
// Solidity: function getOrderDetails(uint256 _orderId) view returns(address owner_, string uri_, uint256 root_, uint32 numChunks_, uint16 periods_, uint8 replicas_, uint8 filled_)
func (_FileMarket *FileMarketCallerSession) GetOrderDetails(_orderId *big.Int) (struct {
	Owner     common.Address
	Uri       string
	Root      *big.Int
	NumChunks uint32
	Periods   uint16
	Replicas  uint8
	Filled    uint8
}, error) {
	return _FileMarket.Contract.GetOrderDetails(&_FileMarket.CallOpts, _orderId)
}

// GetOrderEscrowInfo is a free data retrieval call binding the contract method 0x695cbac0.
//
// Solidity: function getOrderEscrowInfo(uint256 _orderId) view returns(uint256 totalEscrow, uint256 paidToNodes, uint256 remainingEscrow)
func (_FileMarket *FileMarketCaller) GetOrderEscrowInfo(opts *bind.CallOpts, _orderId *big.Int) (struct {
	TotalEscrow     *big.Int
	PaidToNodes     *big.Int
	RemainingEscrow *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getOrderEscrowInfo", _orderId)

	outstruct := new(struct {
		TotalEscrow     *big.Int
		PaidToNodes     *big.Int
		RemainingEscrow *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalEscrow = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PaidToNodes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RemainingEscrow = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOrderEscrowInfo is a free data retrieval call binding the contract method 0x695cbac0.
//
// Solidity: function getOrderEscrowInfo(uint256 _orderId) view returns(uint256 totalEscrow, uint256 paidToNodes, uint256 remainingEscrow)
func (_FileMarket *FileMarketSession) GetOrderEscrowInfo(_orderId *big.Int) (struct {
	TotalEscrow     *big.Int
	PaidToNodes     *big.Int
	RemainingEscrow *big.Int
}, error) {
	return _FileMarket.Contract.GetOrderEscrowInfo(&_FileMarket.CallOpts, _orderId)
}

// GetOrderEscrowInfo is a free data retrieval call binding the contract method 0x695cbac0.
//
// Solidity: function getOrderEscrowInfo(uint256 _orderId) view returns(uint256 totalEscrow, uint256 paidToNodes, uint256 remainingEscrow)
func (_FileMarket *FileMarketCallerSession) GetOrderEscrowInfo(_orderId *big.Int) (struct {
	TotalEscrow     *big.Int
	PaidToNodes     *big.Int
	RemainingEscrow *big.Int
}, error) {
	return _FileMarket.Contract.GetOrderEscrowInfo(&_FileMarket.CallOpts, _orderId)
}

// GetOrderFinancials is a free data retrieval call binding the contract method 0xefa83c79.
//
// Solidity: function getOrderFinancials(uint256 _orderId) view returns(uint256 escrow_, uint256 withdrawn_, uint64 startPeriod_, bool expired_, address[] nodes_)
func (_FileMarket *FileMarketCaller) GetOrderFinancials(opts *bind.CallOpts, _orderId *big.Int) (struct {
	Escrow      *big.Int
	Withdrawn   *big.Int
	StartPeriod uint64
	Expired     bool
	Nodes       []common.Address
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getOrderFinancials", _orderId)

	outstruct := new(struct {
		Escrow      *big.Int
		Withdrawn   *big.Int
		StartPeriod uint64
		Expired     bool
		Nodes       []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Escrow = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Withdrawn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartPeriod = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.Expired = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Nodes = *abi.ConvertType(out[4], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetOrderFinancials is a free data retrieval call binding the contract method 0xefa83c79.
//
// Solidity: function getOrderFinancials(uint256 _orderId) view returns(uint256 escrow_, uint256 withdrawn_, uint64 startPeriod_, bool expired_, address[] nodes_)
func (_FileMarket *FileMarketSession) GetOrderFinancials(_orderId *big.Int) (struct {
	Escrow      *big.Int
	Withdrawn   *big.Int
	StartPeriod uint64
	Expired     bool
	Nodes       []common.Address
}, error) {
	return _FileMarket.Contract.GetOrderFinancials(&_FileMarket.CallOpts, _orderId)
}

// GetOrderFinancials is a free data retrieval call binding the contract method 0xefa83c79.
//
// Solidity: function getOrderFinancials(uint256 _orderId) view returns(uint256 escrow_, uint256 withdrawn_, uint64 startPeriod_, bool expired_, address[] nodes_)
func (_FileMarket *FileMarketCallerSession) GetOrderFinancials(_orderId *big.Int) (struct {
	Escrow      *big.Int
	Withdrawn   *big.Int
	StartPeriod uint64
	Expired     bool
	Nodes       []common.Address
}, error) {
	return _FileMarket.Contract.GetOrderFinancials(&_FileMarket.CallOpts, _orderId)
}

// GetOrderNodes is a free data retrieval call binding the contract method 0x615f7cfc.
//
// Solidity: function getOrderNodes(uint256 _orderId) view returns(address[])
func (_FileMarket *FileMarketCaller) GetOrderNodes(opts *bind.CallOpts, _orderId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getOrderNodes", _orderId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOrderNodes is a free data retrieval call binding the contract method 0x615f7cfc.
//
// Solidity: function getOrderNodes(uint256 _orderId) view returns(address[])
func (_FileMarket *FileMarketSession) GetOrderNodes(_orderId *big.Int) ([]common.Address, error) {
	return _FileMarket.Contract.GetOrderNodes(&_FileMarket.CallOpts, _orderId)
}

// GetOrderNodes is a free data retrieval call binding the contract method 0x615f7cfc.
//
// Solidity: function getOrderNodes(uint256 _orderId) view returns(address[])
func (_FileMarket *FileMarketCallerSession) GetOrderNodes(_orderId *big.Int) ([]common.Address, error) {
	return _FileMarket.Contract.GetOrderNodes(&_FileMarket.CallOpts, _orderId)
}

// GetProofSystemStats is a free data retrieval call binding the contract method 0xad38c6c0.
//
// Solidity: function getProofSystemStats() view returns(uint256 activeSlotsCount, uint256 idleSlotsCount, uint256 expiredSlotsCount, uint256 currentBlockNumber, uint256 challengeWindowBlocks, uint256 challengeableOrdersCount, uint256 totalSlotsCount)
func (_FileMarket *FileMarketCaller) GetProofSystemStats(opts *bind.CallOpts) (struct {
	ActiveSlotsCount         *big.Int
	IdleSlotsCount           *big.Int
	ExpiredSlotsCount        *big.Int
	CurrentBlockNumber       *big.Int
	ChallengeWindowBlocks    *big.Int
	ChallengeableOrdersCount *big.Int
	TotalSlotsCount          *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getProofSystemStats")

	outstruct := new(struct {
		ActiveSlotsCount         *big.Int
		IdleSlotsCount           *big.Int
		ExpiredSlotsCount        *big.Int
		CurrentBlockNumber       *big.Int
		ChallengeWindowBlocks    *big.Int
		ChallengeableOrdersCount *big.Int
		TotalSlotsCount          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ActiveSlotsCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.IdleSlotsCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ExpiredSlotsCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentBlockNumber = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ChallengeWindowBlocks = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ChallengeableOrdersCount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalSlotsCount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetProofSystemStats is a free data retrieval call binding the contract method 0xad38c6c0.
//
// Solidity: function getProofSystemStats() view returns(uint256 activeSlotsCount, uint256 idleSlotsCount, uint256 expiredSlotsCount, uint256 currentBlockNumber, uint256 challengeWindowBlocks, uint256 challengeableOrdersCount, uint256 totalSlotsCount)
func (_FileMarket *FileMarketSession) GetProofSystemStats() (struct {
	ActiveSlotsCount         *big.Int
	IdleSlotsCount           *big.Int
	ExpiredSlotsCount        *big.Int
	CurrentBlockNumber       *big.Int
	ChallengeWindowBlocks    *big.Int
	ChallengeableOrdersCount *big.Int
	TotalSlotsCount          *big.Int
}, error) {
	return _FileMarket.Contract.GetProofSystemStats(&_FileMarket.CallOpts)
}

// GetProofSystemStats is a free data retrieval call binding the contract method 0xad38c6c0.
//
// Solidity: function getProofSystemStats() view returns(uint256 activeSlotsCount, uint256 idleSlotsCount, uint256 expiredSlotsCount, uint256 currentBlockNumber, uint256 challengeWindowBlocks, uint256 challengeableOrdersCount, uint256 totalSlotsCount)
func (_FileMarket *FileMarketCallerSession) GetProofSystemStats() (struct {
	ActiveSlotsCount         *big.Int
	IdleSlotsCount           *big.Int
	ExpiredSlotsCount        *big.Int
	CurrentBlockNumber       *big.Int
	ChallengeWindowBlocks    *big.Int
	ChallengeableOrdersCount *big.Int
	TotalSlotsCount          *big.Int
}, error) {
	return _FileMarket.Contract.GetProofSystemStats(&_FileMarket.CallOpts)
}

// GetRecentOrders is a free data retrieval call binding the contract method 0x4811bc4b.
//
// Solidity: function getRecentOrders(uint256 count) view returns(uint256[] orderIds, address[] owners, uint32[] numChunks, uint16[] periods, uint8[] replicas, uint8[] filled, uint256[] escrows, bool[] isActive)
func (_FileMarket *FileMarketCaller) GetRecentOrders(opts *bind.CallOpts, count *big.Int) (struct {
	OrderIds  []*big.Int
	Owners    []common.Address
	NumChunks []uint32
	Periods   []uint16
	Replicas  []uint8
	Filled    []uint8
	Escrows   []*big.Int
	IsActive  []bool
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getRecentOrders", count)

	outstruct := new(struct {
		OrderIds  []*big.Int
		Owners    []common.Address
		NumChunks []uint32
		Periods   []uint16
		Replicas  []uint8
		Filled    []uint8
		Escrows   []*big.Int
		IsActive  []bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderIds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Owners = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.NumChunks = *abi.ConvertType(out[2], new([]uint32)).(*[]uint32)
	outstruct.Periods = *abi.ConvertType(out[3], new([]uint16)).(*[]uint16)
	outstruct.Replicas = *abi.ConvertType(out[4], new([]uint8)).(*[]uint8)
	outstruct.Filled = *abi.ConvertType(out[5], new([]uint8)).(*[]uint8)
	outstruct.Escrows = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)
	outstruct.IsActive = *abi.ConvertType(out[7], new([]bool)).(*[]bool)

	return *outstruct, err

}

// GetRecentOrders is a free data retrieval call binding the contract method 0x4811bc4b.
//
// Solidity: function getRecentOrders(uint256 count) view returns(uint256[] orderIds, address[] owners, uint32[] numChunks, uint16[] periods, uint8[] replicas, uint8[] filled, uint256[] escrows, bool[] isActive)
func (_FileMarket *FileMarketSession) GetRecentOrders(count *big.Int) (struct {
	OrderIds  []*big.Int
	Owners    []common.Address
	NumChunks []uint32
	Periods   []uint16
	Replicas  []uint8
	Filled    []uint8
	Escrows   []*big.Int
	IsActive  []bool
}, error) {
	return _FileMarket.Contract.GetRecentOrders(&_FileMarket.CallOpts, count)
}

// GetRecentOrders is a free data retrieval call binding the contract method 0x4811bc4b.
//
// Solidity: function getRecentOrders(uint256 count) view returns(uint256[] orderIds, address[] owners, uint32[] numChunks, uint16[] periods, uint8[] replicas, uint8[] filled, uint256[] escrows, bool[] isActive)
func (_FileMarket *FileMarketCallerSession) GetRecentOrders(count *big.Int) (struct {
	OrderIds  []*big.Int
	Owners    []common.Address
	NumChunks []uint32
	Periods   []uint16
	Replicas  []uint8
	Filled    []uint8
	Escrows   []*big.Int
	IsActive  []bool
}, error) {
	return _FileMarket.Contract.GetRecentOrders(&_FileMarket.CallOpts, count)
}

// GetReporterEarningsInfo is a free data retrieval call binding the contract method 0x33f19658.
//
// Solidity: function getReporterEarningsInfo(address _reporter) view returns(uint256 earned, uint256 withdrawn, uint256 pending)
func (_FileMarket *FileMarketCaller) GetReporterEarningsInfo(opts *bind.CallOpts, _reporter common.Address) (struct {
	Earned    *big.Int
	Withdrawn *big.Int
	Pending   *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getReporterEarningsInfo", _reporter)

	outstruct := new(struct {
		Earned    *big.Int
		Withdrawn *big.Int
		Pending   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Earned = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Withdrawn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Pending = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReporterEarningsInfo is a free data retrieval call binding the contract method 0x33f19658.
//
// Solidity: function getReporterEarningsInfo(address _reporter) view returns(uint256 earned, uint256 withdrawn, uint256 pending)
func (_FileMarket *FileMarketSession) GetReporterEarningsInfo(_reporter common.Address) (struct {
	Earned    *big.Int
	Withdrawn *big.Int
	Pending   *big.Int
}, error) {
	return _FileMarket.Contract.GetReporterEarningsInfo(&_FileMarket.CallOpts, _reporter)
}

// GetReporterEarningsInfo is a free data retrieval call binding the contract method 0x33f19658.
//
// Solidity: function getReporterEarningsInfo(address _reporter) view returns(uint256 earned, uint256 withdrawn, uint256 pending)
func (_FileMarket *FileMarketCallerSession) GetReporterEarningsInfo(_reporter common.Address) (struct {
	Earned    *big.Int
	Withdrawn *big.Int
	Pending   *big.Int
}, error) {
	return _FileMarket.Contract.GetReporterEarningsInfo(&_FileMarket.CallOpts, _reporter)
}

// GetSlashRedistributionStats is a free data retrieval call binding the contract method 0x87ca98a2.
//
// Solidity: function getSlashRedistributionStats() view returns(uint256 totalReceived, uint256 totalBurned, uint256 totalRewards, uint256 currentBps, uint256 totalClientComp)
func (_FileMarket *FileMarketCaller) GetSlashRedistributionStats(opts *bind.CallOpts) (struct {
	TotalReceived   *big.Int
	TotalBurned     *big.Int
	TotalRewards    *big.Int
	CurrentBps      *big.Int
	TotalClientComp *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getSlashRedistributionStats")

	outstruct := new(struct {
		TotalReceived   *big.Int
		TotalBurned     *big.Int
		TotalRewards    *big.Int
		CurrentBps      *big.Int
		TotalClientComp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalReceived = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalBurned = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalRewards = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentBps = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalClientComp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSlashRedistributionStats is a free data retrieval call binding the contract method 0x87ca98a2.
//
// Solidity: function getSlashRedistributionStats() view returns(uint256 totalReceived, uint256 totalBurned, uint256 totalRewards, uint256 currentBps, uint256 totalClientComp)
func (_FileMarket *FileMarketSession) GetSlashRedistributionStats() (struct {
	TotalReceived   *big.Int
	TotalBurned     *big.Int
	TotalRewards    *big.Int
	CurrentBps      *big.Int
	TotalClientComp *big.Int
}, error) {
	return _FileMarket.Contract.GetSlashRedistributionStats(&_FileMarket.CallOpts)
}

// GetSlashRedistributionStats is a free data retrieval call binding the contract method 0x87ca98a2.
//
// Solidity: function getSlashRedistributionStats() view returns(uint256 totalReceived, uint256 totalBurned, uint256 totalRewards, uint256 currentBps, uint256 totalClientComp)
func (_FileMarket *FileMarketCallerSession) GetSlashRedistributionStats() (struct {
	TotalReceived   *big.Int
	TotalBurned     *big.Int
	TotalRewards    *big.Int
	CurrentBps      *big.Int
	TotalClientComp *big.Int
}, error) {
	return _FileMarket.Contract.GetSlashRedistributionStats(&_FileMarket.CallOpts)
}

// GetSlotInfo is a free data retrieval call binding the contract method 0xbe20f9ac.
//
// Solidity: function getSlotInfo(uint256 _slotIndex) view returns(uint256 orderId, address challengedNode, uint256 randomness, uint256 deadlineBlock, bool isExpired)
func (_FileMarket *FileMarketCaller) GetSlotInfo(opts *bind.CallOpts, _slotIndex *big.Int) (struct {
	OrderId        *big.Int
	ChallengedNode common.Address
	Randomness     *big.Int
	DeadlineBlock  *big.Int
	IsExpired      bool
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "getSlotInfo", _slotIndex)

	outstruct := new(struct {
		OrderId        *big.Int
		ChallengedNode common.Address
		Randomness     *big.Int
		DeadlineBlock  *big.Int
		IsExpired      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ChallengedNode = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Randomness = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.DeadlineBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsExpired = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// GetSlotInfo is a free data retrieval call binding the contract method 0xbe20f9ac.
//
// Solidity: function getSlotInfo(uint256 _slotIndex) view returns(uint256 orderId, address challengedNode, uint256 randomness, uint256 deadlineBlock, bool isExpired)
func (_FileMarket *FileMarketSession) GetSlotInfo(_slotIndex *big.Int) (struct {
	OrderId        *big.Int
	ChallengedNode common.Address
	Randomness     *big.Int
	DeadlineBlock  *big.Int
	IsExpired      bool
}, error) {
	return _FileMarket.Contract.GetSlotInfo(&_FileMarket.CallOpts, _slotIndex)
}

// GetSlotInfo is a free data retrieval call binding the contract method 0xbe20f9ac.
//
// Solidity: function getSlotInfo(uint256 _slotIndex) view returns(uint256 orderId, address challengedNode, uint256 randomness, uint256 deadlineBlock, bool isExpired)
func (_FileMarket *FileMarketCallerSession) GetSlotInfo(_slotIndex *big.Int) (struct {
	OrderId        *big.Int
	ChallengedNode common.Address
	Randomness     *big.Int
	DeadlineBlock  *big.Int
	IsExpired      bool
}, error) {
	return _FileMarket.Contract.GetSlotInfo(&_FileMarket.CallOpts, _slotIndex)
}

// GlobalSeedRandomness is a free data retrieval call binding the contract method 0xef7c9c8f.
//
// Solidity: function globalSeedRandomness() view returns(uint256)
func (_FileMarket *FileMarketCaller) GlobalSeedRandomness(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "globalSeedRandomness")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalSeedRandomness is a free data retrieval call binding the contract method 0xef7c9c8f.
//
// Solidity: function globalSeedRandomness() view returns(uint256)
func (_FileMarket *FileMarketSession) GlobalSeedRandomness() (*big.Int, error) {
	return _FileMarket.Contract.GlobalSeedRandomness(&_FileMarket.CallOpts)
}

// GlobalSeedRandomness is a free data retrieval call binding the contract method 0xef7c9c8f.
//
// Solidity: function globalSeedRandomness() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) GlobalSeedRandomness() (*big.Int, error) {
	return _FileMarket.Contract.GlobalSeedRandomness(&_FileMarket.CallOpts)
}

// HasUnresolvedProofObligation is a free data retrieval call binding the contract method 0xf1d39a2a.
//
// Solidity: function hasUnresolvedProofObligation(address _node) view returns(bool)
func (_FileMarket *FileMarketCaller) HasUnresolvedProofObligation(opts *bind.CallOpts, _node common.Address) (bool, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "hasUnresolvedProofObligation", _node)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUnresolvedProofObligation is a free data retrieval call binding the contract method 0xf1d39a2a.
//
// Solidity: function hasUnresolvedProofObligation(address _node) view returns(bool)
func (_FileMarket *FileMarketSession) HasUnresolvedProofObligation(_node common.Address) (bool, error) {
	return _FileMarket.Contract.HasUnresolvedProofObligation(&_FileMarket.CallOpts, _node)
}

// HasUnresolvedProofObligation is a free data retrieval call binding the contract method 0xf1d39a2a.
//
// Solidity: function hasUnresolvedProofObligation(address _node) view returns(bool)
func (_FileMarket *FileMarketCallerSession) HasUnresolvedProofObligation(_node common.Address) (bool, error) {
	return _FileMarket.Contract.HasUnresolvedProofObligation(&_FileMarket.CallOpts, _node)
}

// IsChallengeable is a free data retrieval call binding the contract method 0x98dd89cf.
//
// Solidity: function isChallengeable(uint256 ) view returns(bool)
func (_FileMarket *FileMarketCaller) IsChallengeable(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "isChallengeable", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChallengeable is a free data retrieval call binding the contract method 0x98dd89cf.
//
// Solidity: function isChallengeable(uint256 ) view returns(bool)
func (_FileMarket *FileMarketSession) IsChallengeable(arg0 *big.Int) (bool, error) {
	return _FileMarket.Contract.IsChallengeable(&_FileMarket.CallOpts, arg0)
}

// IsChallengeable is a free data retrieval call binding the contract method 0x98dd89cf.
//
// Solidity: function isChallengeable(uint256 ) view returns(bool)
func (_FileMarket *FileMarketCallerSession) IsChallengeable(arg0 *big.Int) (bool, error) {
	return _FileMarket.Contract.IsChallengeable(&_FileMarket.CallOpts, arg0)
}

// IsOrderExpired is a free data retrieval call binding the contract method 0x59c69313.
//
// Solidity: function isOrderExpired(uint256 _orderId) view returns(bool)
func (_FileMarket *FileMarketCaller) IsOrderExpired(opts *bind.CallOpts, _orderId *big.Int) (bool, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "isOrderExpired", _orderId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderExpired is a free data retrieval call binding the contract method 0x59c69313.
//
// Solidity: function isOrderExpired(uint256 _orderId) view returns(bool)
func (_FileMarket *FileMarketSession) IsOrderExpired(_orderId *big.Int) (bool, error) {
	return _FileMarket.Contract.IsOrderExpired(&_FileMarket.CallOpts, _orderId)
}

// IsOrderExpired is a free data retrieval call binding the contract method 0x59c69313.
//
// Solidity: function isOrderExpired(uint256 _orderId) view returns(bool)
func (_FileMarket *FileMarketCallerSession) IsOrderExpired(_orderId *big.Int) (bool, error) {
	return _FileMarket.Contract.IsOrderExpired(&_FileMarket.CallOpts, _orderId)
}

// KeyleakVerifier is a free data retrieval call binding the contract method 0xdf4daf65.
//
// Solidity: function keyleakVerifier() view returns(address)
func (_FileMarket *FileMarketCaller) KeyleakVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "keyleakVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyleakVerifier is a free data retrieval call binding the contract method 0xdf4daf65.
//
// Solidity: function keyleakVerifier() view returns(address)
func (_FileMarket *FileMarketSession) KeyleakVerifier() (common.Address, error) {
	return _FileMarket.Contract.KeyleakVerifier(&_FileMarket.CallOpts)
}

// KeyleakVerifier is a free data retrieval call binding the contract method 0xdf4daf65.
//
// Solidity: function keyleakVerifier() view returns(address)
func (_FileMarket *FileMarketCallerSession) KeyleakVerifier() (common.Address, error) {
	return _FileMarket.Contract.KeyleakVerifier(&_FileMarket.CallOpts)
}

// LifetimeEscrowDeposited is a free data retrieval call binding the contract method 0x02e79283.
//
// Solidity: function lifetimeEscrowDeposited() view returns(uint256)
func (_FileMarket *FileMarketCaller) LifetimeEscrowDeposited(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "lifetimeEscrowDeposited")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LifetimeEscrowDeposited is a free data retrieval call binding the contract method 0x02e79283.
//
// Solidity: function lifetimeEscrowDeposited() view returns(uint256)
func (_FileMarket *FileMarketSession) LifetimeEscrowDeposited() (*big.Int, error) {
	return _FileMarket.Contract.LifetimeEscrowDeposited(&_FileMarket.CallOpts)
}

// LifetimeEscrowDeposited is a free data retrieval call binding the contract method 0x02e79283.
//
// Solidity: function lifetimeEscrowDeposited() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) LifetimeEscrowDeposited() (*big.Int, error) {
	return _FileMarket.Contract.LifetimeEscrowDeposited(&_FileMarket.CallOpts)
}

// LifetimeRewardsPaid is a free data retrieval call binding the contract method 0x9bd9c558.
//
// Solidity: function lifetimeRewardsPaid() view returns(uint256)
func (_FileMarket *FileMarketCaller) LifetimeRewardsPaid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "lifetimeRewardsPaid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LifetimeRewardsPaid is a free data retrieval call binding the contract method 0x9bd9c558.
//
// Solidity: function lifetimeRewardsPaid() view returns(uint256)
func (_FileMarket *FileMarketSession) LifetimeRewardsPaid() (*big.Int, error) {
	return _FileMarket.Contract.LifetimeRewardsPaid(&_FileMarket.CallOpts)
}

// LifetimeRewardsPaid is a free data retrieval call binding the contract method 0x9bd9c558.
//
// Solidity: function lifetimeRewardsPaid() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) LifetimeRewardsPaid() (*big.Int, error) {
	return _FileMarket.Contract.LifetimeRewardsPaid(&_FileMarket.CallOpts)
}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_FileMarket *FileMarketCaller) NextOrderId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "nextOrderId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_FileMarket *FileMarketSession) NextOrderId() (*big.Int, error) {
	return _FileMarket.Contract.NextOrderId(&_FileMarket.CallOpts)
}

// NextOrderId is a free data retrieval call binding the contract method 0x2a58b330.
//
// Solidity: function nextOrderId() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) NextOrderId() (*big.Int, error) {
	return _FileMarket.Contract.NextOrderId(&_FileMarket.CallOpts)
}

// NodeActiveChallengeCount is a free data retrieval call binding the contract method 0xb90824f0.
//
// Solidity: function nodeActiveChallengeCount(address ) view returns(uint256)
func (_FileMarket *FileMarketCaller) NodeActiveChallengeCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "nodeActiveChallengeCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeActiveChallengeCount is a free data retrieval call binding the contract method 0xb90824f0.
//
// Solidity: function nodeActiveChallengeCount(address ) view returns(uint256)
func (_FileMarket *FileMarketSession) NodeActiveChallengeCount(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodeActiveChallengeCount(&_FileMarket.CallOpts, arg0)
}

// NodeActiveChallengeCount is a free data retrieval call binding the contract method 0xb90824f0.
//
// Solidity: function nodeActiveChallengeCount(address ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) NodeActiveChallengeCount(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodeActiveChallengeCount(&_FileMarket.CallOpts, arg0)
}

// NodeEarnings is a free data retrieval call binding the contract method 0x10641175.
//
// Solidity: function nodeEarnings(address ) view returns(uint256)
func (_FileMarket *FileMarketCaller) NodeEarnings(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "nodeEarnings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeEarnings is a free data retrieval call binding the contract method 0x10641175.
//
// Solidity: function nodeEarnings(address ) view returns(uint256)
func (_FileMarket *FileMarketSession) NodeEarnings(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodeEarnings(&_FileMarket.CallOpts, arg0)
}

// NodeEarnings is a free data retrieval call binding the contract method 0x10641175.
//
// Solidity: function nodeEarnings(address ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) NodeEarnings(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodeEarnings(&_FileMarket.CallOpts, arg0)
}

// NodeLastClaimPeriod is a free data retrieval call binding the contract method 0x869d3f17.
//
// Solidity: function nodeLastClaimPeriod(address ) view returns(uint256)
func (_FileMarket *FileMarketCaller) NodeLastClaimPeriod(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "nodeLastClaimPeriod", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeLastClaimPeriod is a free data retrieval call binding the contract method 0x869d3f17.
//
// Solidity: function nodeLastClaimPeriod(address ) view returns(uint256)
func (_FileMarket *FileMarketSession) NodeLastClaimPeriod(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodeLastClaimPeriod(&_FileMarket.CallOpts, arg0)
}

// NodeLastClaimPeriod is a free data retrieval call binding the contract method 0x869d3f17.
//
// Solidity: function nodeLastClaimPeriod(address ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) NodeLastClaimPeriod(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodeLastClaimPeriod(&_FileMarket.CallOpts, arg0)
}

// NodeOrderEarnings is a free data retrieval call binding the contract method 0xa00c6484.
//
// Solidity: function nodeOrderEarnings(uint256 , address ) view returns(uint256)
func (_FileMarket *FileMarketCaller) NodeOrderEarnings(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "nodeOrderEarnings", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeOrderEarnings is a free data retrieval call binding the contract method 0xa00c6484.
//
// Solidity: function nodeOrderEarnings(uint256 , address ) view returns(uint256)
func (_FileMarket *FileMarketSession) NodeOrderEarnings(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodeOrderEarnings(&_FileMarket.CallOpts, arg0, arg1)
}

// NodeOrderEarnings is a free data retrieval call binding the contract method 0xa00c6484.
//
// Solidity: function nodeOrderEarnings(uint256 , address ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) NodeOrderEarnings(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodeOrderEarnings(&_FileMarket.CallOpts, arg0, arg1)
}

// NodeOrderStartTimestamp is a free data retrieval call binding the contract method 0xdcfd44e8.
//
// Solidity: function nodeOrderStartTimestamp(address , uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCaller) NodeOrderStartTimestamp(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "nodeOrderStartTimestamp", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeOrderStartTimestamp is a free data retrieval call binding the contract method 0xdcfd44e8.
//
// Solidity: function nodeOrderStartTimestamp(address , uint256 ) view returns(uint256)
func (_FileMarket *FileMarketSession) NodeOrderStartTimestamp(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.NodeOrderStartTimestamp(&_FileMarket.CallOpts, arg0, arg1)
}

// NodeOrderStartTimestamp is a free data retrieval call binding the contract method 0xdcfd44e8.
//
// Solidity: function nodeOrderStartTimestamp(address , uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) NodeOrderStartTimestamp(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.NodeOrderStartTimestamp(&_FileMarket.CallOpts, arg0, arg1)
}

// NodePendingRewards is a free data retrieval call binding the contract method 0x08d55762.
//
// Solidity: function nodePendingRewards(address ) view returns(uint256)
func (_FileMarket *FileMarketCaller) NodePendingRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "nodePendingRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodePendingRewards is a free data retrieval call binding the contract method 0x08d55762.
//
// Solidity: function nodePendingRewards(address ) view returns(uint256)
func (_FileMarket *FileMarketSession) NodePendingRewards(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodePendingRewards(&_FileMarket.CallOpts, arg0)
}

// NodePendingRewards is a free data retrieval call binding the contract method 0x08d55762.
//
// Solidity: function nodePendingRewards(address ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) NodePendingRewards(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodePendingRewards(&_FileMarket.CallOpts, arg0)
}

// NodeStaking is a free data retrieval call binding the contract method 0xaa70d54b.
//
// Solidity: function nodeStaking() view returns(address)
func (_FileMarket *FileMarketCaller) NodeStaking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "nodeStaking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeStaking is a free data retrieval call binding the contract method 0xaa70d54b.
//
// Solidity: function nodeStaking() view returns(address)
func (_FileMarket *FileMarketSession) NodeStaking() (common.Address, error) {
	return _FileMarket.Contract.NodeStaking(&_FileMarket.CallOpts)
}

// NodeStaking is a free data retrieval call binding the contract method 0xaa70d54b.
//
// Solidity: function nodeStaking() view returns(address)
func (_FileMarket *FileMarketCallerSession) NodeStaking() (common.Address, error) {
	return _FileMarket.Contract.NodeStaking(&_FileMarket.CallOpts)
}

// NodeToOrders is a free data retrieval call binding the contract method 0x13491d32.
//
// Solidity: function nodeToOrders(address , uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCaller) NodeToOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "nodeToOrders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeToOrders is a free data retrieval call binding the contract method 0x13491d32.
//
// Solidity: function nodeToOrders(address , uint256 ) view returns(uint256)
func (_FileMarket *FileMarketSession) NodeToOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.NodeToOrders(&_FileMarket.CallOpts, arg0, arg1)
}

// NodeToOrders is a free data retrieval call binding the contract method 0x13491d32.
//
// Solidity: function nodeToOrders(address , uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) NodeToOrders(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.NodeToOrders(&_FileMarket.CallOpts, arg0, arg1)
}

// NodeWithdrawn is a free data retrieval call binding the contract method 0x5c2521ea.
//
// Solidity: function nodeWithdrawn(address ) view returns(uint256)
func (_FileMarket *FileMarketCaller) NodeWithdrawn(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "nodeWithdrawn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodeWithdrawn is a free data retrieval call binding the contract method 0x5c2521ea.
//
// Solidity: function nodeWithdrawn(address ) view returns(uint256)
func (_FileMarket *FileMarketSession) NodeWithdrawn(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodeWithdrawn(&_FileMarket.CallOpts, arg0)
}

// NodeWithdrawn is a free data retrieval call binding the contract method 0x5c2521ea.
//
// Solidity: function nodeWithdrawn(address ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) NodeWithdrawn(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.NodeWithdrawn(&_FileMarket.CallOpts, arg0)
}

// NumChallengeSlots is a free data retrieval call binding the contract method 0xf32b57d8.
//
// Solidity: function numChallengeSlots() view returns(uint256)
func (_FileMarket *FileMarketCaller) NumChallengeSlots(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "numChallengeSlots")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumChallengeSlots is a free data retrieval call binding the contract method 0xf32b57d8.
//
// Solidity: function numChallengeSlots() view returns(uint256)
func (_FileMarket *FileMarketSession) NumChallengeSlots() (*big.Int, error) {
	return _FileMarket.Contract.NumChallengeSlots(&_FileMarket.CallOpts)
}

// NumChallengeSlots is a free data retrieval call binding the contract method 0xf32b57d8.
//
// Solidity: function numChallengeSlots() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) NumChallengeSlots() (*big.Int, error) {
	return _FileMarket.Contract.NumChallengeSlots(&_FileMarket.CallOpts)
}

// OnDemandChallenges is a free data retrieval call binding the contract method 0xb802ad32.
//
// Solidity: function onDemandChallenges(bytes32 ) view returns(uint64 deadlineBlock, uint256 randomness, address challenger, uint256 fileRoot)
func (_FileMarket *FileMarketCaller) OnDemandChallenges(opts *bind.CallOpts, arg0 [32]byte) (struct {
	DeadlineBlock uint64
	Randomness    *big.Int
	Challenger    common.Address
	FileRoot      *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "onDemandChallenges", arg0)

	outstruct := new(struct {
		DeadlineBlock uint64
		Randomness    *big.Int
		Challenger    common.Address
		FileRoot      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DeadlineBlock = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Randomness = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Challenger = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.FileRoot = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OnDemandChallenges is a free data retrieval call binding the contract method 0xb802ad32.
//
// Solidity: function onDemandChallenges(bytes32 ) view returns(uint64 deadlineBlock, uint256 randomness, address challenger, uint256 fileRoot)
func (_FileMarket *FileMarketSession) OnDemandChallenges(arg0 [32]byte) (struct {
	DeadlineBlock uint64
	Randomness    *big.Int
	Challenger    common.Address
	FileRoot      *big.Int
}, error) {
	return _FileMarket.Contract.OnDemandChallenges(&_FileMarket.CallOpts, arg0)
}

// OnDemandChallenges is a free data retrieval call binding the contract method 0xb802ad32.
//
// Solidity: function onDemandChallenges(bytes32 ) view returns(uint64 deadlineBlock, uint256 randomness, address challenger, uint256 fileRoot)
func (_FileMarket *FileMarketCallerSession) OnDemandChallenges(arg0 [32]byte) (struct {
	DeadlineBlock uint64
	Randomness    *big.Int
	Challenger    common.Address
	FileRoot      *big.Int
}, error) {
	return _FileMarket.Contract.OnDemandChallenges(&_FileMarket.CallOpts, arg0)
}

// OrderActiveChallengeCount is a free data retrieval call binding the contract method 0xa351e119.
//
// Solidity: function orderActiveChallengeCount(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCaller) OrderActiveChallengeCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "orderActiveChallengeCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderActiveChallengeCount is a free data retrieval call binding the contract method 0xa351e119.
//
// Solidity: function orderActiveChallengeCount(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketSession) OrderActiveChallengeCount(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.OrderActiveChallengeCount(&_FileMarket.CallOpts, arg0)
}

// OrderActiveChallengeCount is a free data retrieval call binding the contract method 0xa351e119.
//
// Solidity: function orderActiveChallengeCount(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) OrderActiveChallengeCount(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.OrderActiveChallengeCount(&_FileMarket.CallOpts, arg0)
}

// OrderEscrowWithdrawn is a free data retrieval call binding the contract method 0x03909140.
//
// Solidity: function orderEscrowWithdrawn(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCaller) OrderEscrowWithdrawn(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "orderEscrowWithdrawn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderEscrowWithdrawn is a free data retrieval call binding the contract method 0x03909140.
//
// Solidity: function orderEscrowWithdrawn(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketSession) OrderEscrowWithdrawn(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.OrderEscrowWithdrawn(&_FileMarket.CallOpts, arg0)
}

// OrderEscrowWithdrawn is a free data retrieval call binding the contract method 0x03909140.
//
// Solidity: function orderEscrowWithdrawn(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) OrderEscrowWithdrawn(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.OrderEscrowWithdrawn(&_FileMarket.CallOpts, arg0)
}

// OrderIndexInActive is a free data retrieval call binding the contract method 0x8a25b8ad.
//
// Solidity: function orderIndexInActive(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCaller) OrderIndexInActive(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "orderIndexInActive", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderIndexInActive is a free data retrieval call binding the contract method 0x8a25b8ad.
//
// Solidity: function orderIndexInActive(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketSession) OrderIndexInActive(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.OrderIndexInActive(&_FileMarket.CallOpts, arg0)
}

// OrderIndexInActive is a free data retrieval call binding the contract method 0x8a25b8ad.
//
// Solidity: function orderIndexInActive(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) OrderIndexInActive(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.OrderIndexInActive(&_FileMarket.CallOpts, arg0)
}

// OrderIndexInChallengeable is a free data retrieval call binding the contract method 0xd12efff1.
//
// Solidity: function orderIndexInChallengeable(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCaller) OrderIndexInChallengeable(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "orderIndexInChallengeable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderIndexInChallengeable is a free data retrieval call binding the contract method 0xd12efff1.
//
// Solidity: function orderIndexInChallengeable(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketSession) OrderIndexInChallengeable(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.OrderIndexInChallengeable(&_FileMarket.CallOpts, arg0)
}

// OrderIndexInChallengeable is a free data retrieval call binding the contract method 0xd12efff1.
//
// Solidity: function orderIndexInChallengeable(uint256 ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) OrderIndexInChallengeable(arg0 *big.Int) (*big.Int, error) {
	return _FileMarket.Contract.OrderIndexInChallengeable(&_FileMarket.CallOpts, arg0)
}

// OrderToNodes is a free data retrieval call binding the contract method 0x4e51245d.
//
// Solidity: function orderToNodes(uint256 , uint256 ) view returns(address)
func (_FileMarket *FileMarketCaller) OrderToNodes(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "orderToNodes", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderToNodes is a free data retrieval call binding the contract method 0x4e51245d.
//
// Solidity: function orderToNodes(uint256 , uint256 ) view returns(address)
func (_FileMarket *FileMarketSession) OrderToNodes(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _FileMarket.Contract.OrderToNodes(&_FileMarket.CallOpts, arg0, arg1)
}

// OrderToNodes is a free data retrieval call binding the contract method 0x4e51245d.
//
// Solidity: function orderToNodes(uint256 , uint256 ) view returns(address)
func (_FileMarket *FileMarketCallerSession) OrderToNodes(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _FileMarket.Contract.OrderToNodes(&_FileMarket.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address owner, (uint256,string) file, uint32 numChunks, uint16 periods, uint8 replicas, uint256 price, uint8 filled, uint64 startPeriod, uint256 escrow)
func (_FileMarket *FileMarketCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner       common.Address
	File        MarketStorageFileMeta
	NumChunks   uint32
	Periods     uint16
	Replicas    uint8
	Price       *big.Int
	Filled      uint8
	StartPeriod uint64
	Escrow      *big.Int
}, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		Owner       common.Address
		File        MarketStorageFileMeta
		NumChunks   uint32
		Periods     uint16
		Replicas    uint8
		Price       *big.Int
		Filled      uint8
		StartPeriod uint64
		Escrow      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.File = *abi.ConvertType(out[1], new(MarketStorageFileMeta)).(*MarketStorageFileMeta)
	outstruct.NumChunks = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Periods = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.Replicas = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Filled = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.StartPeriod = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.Escrow = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address owner, (uint256,string) file, uint32 numChunks, uint16 periods, uint8 replicas, uint256 price, uint8 filled, uint64 startPeriod, uint256 escrow)
func (_FileMarket *FileMarketSession) Orders(arg0 *big.Int) (struct {
	Owner       common.Address
	File        MarketStorageFileMeta
	NumChunks   uint32
	Periods     uint16
	Replicas    uint8
	Price       *big.Int
	Filled      uint8
	StartPeriod uint64
	Escrow      *big.Int
}, error) {
	return _FileMarket.Contract.Orders(&_FileMarket.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(address owner, (uint256,string) file, uint32 numChunks, uint16 periods, uint8 replicas, uint256 price, uint8 filled, uint64 startPeriod, uint256 escrow)
func (_FileMarket *FileMarketCallerSession) Orders(arg0 *big.Int) (struct {
	Owner       common.Address
	File        MarketStorageFileMeta
	NumChunks   uint32
	Periods     uint16
	Replicas    uint8
	Price       *big.Int
	Filled      uint8
	StartPeriod uint64
	Escrow      *big.Int
}, error) {
	return _FileMarket.Contract.Orders(&_FileMarket.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FileMarket *FileMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FileMarket *FileMarketSession) Owner() (common.Address, error) {
	return _FileMarket.Contract.Owner(&_FileMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FileMarket *FileMarketCallerSession) Owner() (common.Address, error) {
	return _FileMarket.Contract.Owner(&_FileMarket.CallOpts)
}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_FileMarket *FileMarketCaller) PendingRefunds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "pendingRefunds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_FileMarket *FileMarketSession) PendingRefunds(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.PendingRefunds(&_FileMarket.CallOpts, arg0)
}

// PendingRefunds is a free data retrieval call binding the contract method 0xb613b114.
//
// Solidity: function pendingRefunds(address ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) PendingRefunds(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.PendingRefunds(&_FileMarket.CallOpts, arg0)
}

// PoiVerifier is a free data retrieval call binding the contract method 0xd67a5395.
//
// Solidity: function poiVerifier() view returns(address)
func (_FileMarket *FileMarketCaller) PoiVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "poiVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoiVerifier is a free data retrieval call binding the contract method 0xd67a5395.
//
// Solidity: function poiVerifier() view returns(address)
func (_FileMarket *FileMarketSession) PoiVerifier() (common.Address, error) {
	return _FileMarket.Contract.PoiVerifier(&_FileMarket.CallOpts)
}

// PoiVerifier is a free data retrieval call binding the contract method 0xd67a5395.
//
// Solidity: function poiVerifier() view returns(address)
func (_FileMarket *FileMarketCallerSession) PoiVerifier() (common.Address, error) {
	return _FileMarket.Contract.PoiVerifier(&_FileMarket.CallOpts)
}

// ProofFailureSlashMultiplier is a free data retrieval call binding the contract method 0x65a5a216.
//
// Solidity: function proofFailureSlashMultiplier() view returns(uint256)
func (_FileMarket *FileMarketCaller) ProofFailureSlashMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "proofFailureSlashMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProofFailureSlashMultiplier is a free data retrieval call binding the contract method 0x65a5a216.
//
// Solidity: function proofFailureSlashMultiplier() view returns(uint256)
func (_FileMarket *FileMarketSession) ProofFailureSlashMultiplier() (*big.Int, error) {
	return _FileMarket.Contract.ProofFailureSlashMultiplier(&_FileMarket.CallOpts)
}

// ProofFailureSlashMultiplier is a free data retrieval call binding the contract method 0x65a5a216.
//
// Solidity: function proofFailureSlashMultiplier() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) ProofFailureSlashMultiplier() (*big.Int, error) {
	return _FileMarket.Contract.ProofFailureSlashMultiplier(&_FileMarket.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FileMarket *FileMarketCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FileMarket *FileMarketSession) ProxiableUUID() ([32]byte, error) {
	return _FileMarket.Contract.ProxiableUUID(&_FileMarket.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_FileMarket *FileMarketCallerSession) ProxiableUUID() ([32]byte, error) {
	return _FileMarket.Contract.ProxiableUUID(&_FileMarket.CallOpts)
}

// ReporterEarnings is a free data retrieval call binding the contract method 0x258c58fa.
//
// Solidity: function reporterEarnings(address ) view returns(uint256)
func (_FileMarket *FileMarketCaller) ReporterEarnings(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "reporterEarnings", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReporterEarnings is a free data retrieval call binding the contract method 0x258c58fa.
//
// Solidity: function reporterEarnings(address ) view returns(uint256)
func (_FileMarket *FileMarketSession) ReporterEarnings(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.ReporterEarnings(&_FileMarket.CallOpts, arg0)
}

// ReporterEarnings is a free data retrieval call binding the contract method 0x258c58fa.
//
// Solidity: function reporterEarnings(address ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) ReporterEarnings(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.ReporterEarnings(&_FileMarket.CallOpts, arg0)
}

// ReporterPendingRewards is a free data retrieval call binding the contract method 0x9eae7c47.
//
// Solidity: function reporterPendingRewards(address ) view returns(uint256)
func (_FileMarket *FileMarketCaller) ReporterPendingRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "reporterPendingRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReporterPendingRewards is a free data retrieval call binding the contract method 0x9eae7c47.
//
// Solidity: function reporterPendingRewards(address ) view returns(uint256)
func (_FileMarket *FileMarketSession) ReporterPendingRewards(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.ReporterPendingRewards(&_FileMarket.CallOpts, arg0)
}

// ReporterPendingRewards is a free data retrieval call binding the contract method 0x9eae7c47.
//
// Solidity: function reporterPendingRewards(address ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) ReporterPendingRewards(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.ReporterPendingRewards(&_FileMarket.CallOpts, arg0)
}

// ReporterRewardBps is a free data retrieval call binding the contract method 0xcb93f86c.
//
// Solidity: function reporterRewardBps() view returns(uint256)
func (_FileMarket *FileMarketCaller) ReporterRewardBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "reporterRewardBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReporterRewardBps is a free data retrieval call binding the contract method 0xcb93f86c.
//
// Solidity: function reporterRewardBps() view returns(uint256)
func (_FileMarket *FileMarketSession) ReporterRewardBps() (*big.Int, error) {
	return _FileMarket.Contract.ReporterRewardBps(&_FileMarket.CallOpts)
}

// ReporterRewardBps is a free data retrieval call binding the contract method 0xcb93f86c.
//
// Solidity: function reporterRewardBps() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) ReporterRewardBps() (*big.Int, error) {
	return _FileMarket.Contract.ReporterRewardBps(&_FileMarket.CallOpts)
}

// ReporterWithdrawn is a free data retrieval call binding the contract method 0x0762759a.
//
// Solidity: function reporterWithdrawn(address ) view returns(uint256)
func (_FileMarket *FileMarketCaller) ReporterWithdrawn(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "reporterWithdrawn", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReporterWithdrawn is a free data retrieval call binding the contract method 0x0762759a.
//
// Solidity: function reporterWithdrawn(address ) view returns(uint256)
func (_FileMarket *FileMarketSession) ReporterWithdrawn(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.ReporterWithdrawn(&_FileMarket.CallOpts, arg0)
}

// ReporterWithdrawn is a free data retrieval call binding the contract method 0x0762759a.
//
// Solidity: function reporterWithdrawn(address ) view returns(uint256)
func (_FileMarket *FileMarketCallerSession) ReporterWithdrawn(arg0 common.Address) (*big.Int, error) {
	return _FileMarket.Contract.ReporterWithdrawn(&_FileMarket.CallOpts, arg0)
}

// SlashAuthorities is a free data retrieval call binding the contract method 0x72c71c4b.
//
// Solidity: function slashAuthorities(address ) view returns(bool)
func (_FileMarket *FileMarketCaller) SlashAuthorities(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "slashAuthorities", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SlashAuthorities is a free data retrieval call binding the contract method 0x72c71c4b.
//
// Solidity: function slashAuthorities(address ) view returns(bool)
func (_FileMarket *FileMarketSession) SlashAuthorities(arg0 common.Address) (bool, error) {
	return _FileMarket.Contract.SlashAuthorities(&_FileMarket.CallOpts, arg0)
}

// SlashAuthorities is a free data retrieval call binding the contract method 0x72c71c4b.
//
// Solidity: function slashAuthorities(address ) view returns(bool)
func (_FileMarket *FileMarketCallerSession) SlashAuthorities(arg0 common.Address) (bool, error) {
	return _FileMarket.Contract.SlashAuthorities(&_FileMarket.CallOpts, arg0)
}

// SweepCursor is a free data retrieval call binding the contract method 0xe3c5998f.
//
// Solidity: function sweepCursor() view returns(uint256)
func (_FileMarket *FileMarketCaller) SweepCursor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "sweepCursor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SweepCursor is a free data retrieval call binding the contract method 0xe3c5998f.
//
// Solidity: function sweepCursor() view returns(uint256)
func (_FileMarket *FileMarketSession) SweepCursor() (*big.Int, error) {
	return _FileMarket.Contract.SweepCursor(&_FileMarket.CallOpts)
}

// SweepCursor is a free data retrieval call binding the contract method 0xe3c5998f.
//
// Solidity: function sweepCursor() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) SweepCursor() (*big.Int, error) {
	return _FileMarket.Contract.SweepCursor(&_FileMarket.CallOpts)
}

// TotalBurnedFromSlash is a free data retrieval call binding the contract method 0x319d785e.
//
// Solidity: function totalBurnedFromSlash() view returns(uint256)
func (_FileMarket *FileMarketCaller) TotalBurnedFromSlash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "totalBurnedFromSlash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBurnedFromSlash is a free data retrieval call binding the contract method 0x319d785e.
//
// Solidity: function totalBurnedFromSlash() view returns(uint256)
func (_FileMarket *FileMarketSession) TotalBurnedFromSlash() (*big.Int, error) {
	return _FileMarket.Contract.TotalBurnedFromSlash(&_FileMarket.CallOpts)
}

// TotalBurnedFromSlash is a free data retrieval call binding the contract method 0x319d785e.
//
// Solidity: function totalBurnedFromSlash() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) TotalBurnedFromSlash() (*big.Int, error) {
	return _FileMarket.Contract.TotalBurnedFromSlash(&_FileMarket.CallOpts)
}

// TotalCancellationPenalties is a free data retrieval call binding the contract method 0x673f9465.
//
// Solidity: function totalCancellationPenalties() view returns(uint256)
func (_FileMarket *FileMarketCaller) TotalCancellationPenalties(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "totalCancellationPenalties")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCancellationPenalties is a free data retrieval call binding the contract method 0x673f9465.
//
// Solidity: function totalCancellationPenalties() view returns(uint256)
func (_FileMarket *FileMarketSession) TotalCancellationPenalties() (*big.Int, error) {
	return _FileMarket.Contract.TotalCancellationPenalties(&_FileMarket.CallOpts)
}

// TotalCancellationPenalties is a free data retrieval call binding the contract method 0x673f9465.
//
// Solidity: function totalCancellationPenalties() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) TotalCancellationPenalties() (*big.Int, error) {
	return _FileMarket.Contract.TotalCancellationPenalties(&_FileMarket.CallOpts)
}

// TotalClientCompensation is a free data retrieval call binding the contract method 0xdf7a5c24.
//
// Solidity: function totalClientCompensation() view returns(uint256)
func (_FileMarket *FileMarketCaller) TotalClientCompensation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "totalClientCompensation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClientCompensation is a free data retrieval call binding the contract method 0xdf7a5c24.
//
// Solidity: function totalClientCompensation() view returns(uint256)
func (_FileMarket *FileMarketSession) TotalClientCompensation() (*big.Int, error) {
	return _FileMarket.Contract.TotalClientCompensation(&_FileMarket.CallOpts)
}

// TotalClientCompensation is a free data retrieval call binding the contract method 0xdf7a5c24.
//
// Solidity: function totalClientCompensation() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) TotalClientCompensation() (*big.Int, error) {
	return _FileMarket.Contract.TotalClientCompensation(&_FileMarket.CallOpts)
}

// TotalReporterRewards is a free data retrieval call binding the contract method 0xfa4596c9.
//
// Solidity: function totalReporterRewards() view returns(uint256)
func (_FileMarket *FileMarketCaller) TotalReporterRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "totalReporterRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReporterRewards is a free data retrieval call binding the contract method 0xfa4596c9.
//
// Solidity: function totalReporterRewards() view returns(uint256)
func (_FileMarket *FileMarketSession) TotalReporterRewards() (*big.Int, error) {
	return _FileMarket.Contract.TotalReporterRewards(&_FileMarket.CallOpts)
}

// TotalReporterRewards is a free data retrieval call binding the contract method 0xfa4596c9.
//
// Solidity: function totalReporterRewards() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) TotalReporterRewards() (*big.Int, error) {
	return _FileMarket.Contract.TotalReporterRewards(&_FileMarket.CallOpts)
}

// TotalSlashedReceived is a free data retrieval call binding the contract method 0xa7598a61.
//
// Solidity: function totalSlashedReceived() view returns(uint256)
func (_FileMarket *FileMarketCaller) TotalSlashedReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FileMarket.contract.Call(opts, &out, "totalSlashedReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSlashedReceived is a free data retrieval call binding the contract method 0xa7598a61.
//
// Solidity: function totalSlashedReceived() view returns(uint256)
func (_FileMarket *FileMarketSession) TotalSlashedReceived() (*big.Int, error) {
	return _FileMarket.Contract.TotalSlashedReceived(&_FileMarket.CallOpts)
}

// TotalSlashedReceived is a free data retrieval call binding the contract method 0xa7598a61.
//
// Solidity: function totalSlashedReceived() view returns(uint256)
func (_FileMarket *FileMarketCallerSession) TotalSlashedReceived() (*big.Int, error) {
	return _FileMarket.Contract.TotalSlashedReceived(&_FileMarket.CallOpts)
}

// ActivateSlots is a paid mutator transaction binding the contract method 0x53e3c7a1.
//
// Solidity: function activateSlots() returns()
func (_FileMarket *FileMarketTransactor) ActivateSlots(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "activateSlots")
}

// ActivateSlots is a paid mutator transaction binding the contract method 0x53e3c7a1.
//
// Solidity: function activateSlots() returns()
func (_FileMarket *FileMarketSession) ActivateSlots() (*types.Transaction, error) {
	return _FileMarket.Contract.ActivateSlots(&_FileMarket.TransactOpts)
}

// ActivateSlots is a paid mutator transaction binding the contract method 0x53e3c7a1.
//
// Solidity: function activateSlots() returns()
func (_FileMarket *FileMarketTransactorSession) ActivateSlots() (*types.Transaction, error) {
	return _FileMarket.Contract.ActivateSlots(&_FileMarket.TransactOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _orderId) returns()
func (_FileMarket *FileMarketTransactor) CancelOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "cancelOrder", _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _orderId) returns()
func (_FileMarket *FileMarketSession) CancelOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.CancelOrder(&_FileMarket.TransactOpts, _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x514fcac7.
//
// Solidity: function cancelOrder(uint256 _orderId) returns()
func (_FileMarket *FileMarketTransactorSession) CancelOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.CancelOrder(&_FileMarket.TransactOpts, _orderId)
}

// ChallengeNode is a paid mutator transaction binding the contract method 0xabf8f87c.
//
// Solidity: function challengeNode(uint256 _orderId, address _node) returns()
func (_FileMarket *FileMarketTransactor) ChallengeNode(opts *bind.TransactOpts, _orderId *big.Int, _node common.Address) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "challengeNode", _orderId, _node)
}

// ChallengeNode is a paid mutator transaction binding the contract method 0xabf8f87c.
//
// Solidity: function challengeNode(uint256 _orderId, address _node) returns()
func (_FileMarket *FileMarketSession) ChallengeNode(_orderId *big.Int, _node common.Address) (*types.Transaction, error) {
	return _FileMarket.Contract.ChallengeNode(&_FileMarket.TransactOpts, _orderId, _node)
}

// ChallengeNode is a paid mutator transaction binding the contract method 0xabf8f87c.
//
// Solidity: function challengeNode(uint256 _orderId, address _node) returns()
func (_FileMarket *FileMarketTransactorSession) ChallengeNode(_orderId *big.Int, _node common.Address) (*types.Transaction, error) {
	return _FileMarket.Contract.ChallengeNode(&_FileMarket.TransactOpts, _orderId, _node)
}

// ClaimReporterRewards is a paid mutator transaction binding the contract method 0x1dd1924f.
//
// Solidity: function claimReporterRewards() returns()
func (_FileMarket *FileMarketTransactor) ClaimReporterRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "claimReporterRewards")
}

// ClaimReporterRewards is a paid mutator transaction binding the contract method 0x1dd1924f.
//
// Solidity: function claimReporterRewards() returns()
func (_FileMarket *FileMarketSession) ClaimReporterRewards() (*types.Transaction, error) {
	return _FileMarket.Contract.ClaimReporterRewards(&_FileMarket.TransactOpts)
}

// ClaimReporterRewards is a paid mutator transaction binding the contract method 0x1dd1924f.
//
// Solidity: function claimReporterRewards() returns()
func (_FileMarket *FileMarketTransactorSession) ClaimReporterRewards() (*types.Transaction, error) {
	return _FileMarket.Contract.ClaimReporterRewards(&_FileMarket.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_FileMarket *FileMarketTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_FileMarket *FileMarketSession) ClaimRewards() (*types.Transaction, error) {
	return _FileMarket.Contract.ClaimRewards(&_FileMarket.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_FileMarket *FileMarketTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _FileMarket.Contract.ClaimRewards(&_FileMarket.TransactOpts)
}

// CompleteExpiredOrder is a paid mutator transaction binding the contract method 0x4ee7de66.
//
// Solidity: function completeExpiredOrder(uint256 _orderId) returns()
func (_FileMarket *FileMarketTransactor) CompleteExpiredOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "completeExpiredOrder", _orderId)
}

// CompleteExpiredOrder is a paid mutator transaction binding the contract method 0x4ee7de66.
//
// Solidity: function completeExpiredOrder(uint256 _orderId) returns()
func (_FileMarket *FileMarketSession) CompleteExpiredOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.CompleteExpiredOrder(&_FileMarket.TransactOpts, _orderId)
}

// CompleteExpiredOrder is a paid mutator transaction binding the contract method 0x4ee7de66.
//
// Solidity: function completeExpiredOrder(uint256 _orderId) returns()
func (_FileMarket *FileMarketTransactorSession) CompleteExpiredOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.CompleteExpiredOrder(&_FileMarket.TransactOpts, _orderId)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0xa1a870d9.
//
// Solidity: function executeOrder(uint256 _orderId, uint256[4] _proof, bytes32 _commitment) returns()
func (_FileMarket *FileMarketTransactor) ExecuteOrder(opts *bind.TransactOpts, _orderId *big.Int, _proof [4]*big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "executeOrder", _orderId, _proof, _commitment)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0xa1a870d9.
//
// Solidity: function executeOrder(uint256 _orderId, uint256[4] _proof, bytes32 _commitment) returns()
func (_FileMarket *FileMarketSession) ExecuteOrder(_orderId *big.Int, _proof [4]*big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _FileMarket.Contract.ExecuteOrder(&_FileMarket.TransactOpts, _orderId, _proof, _commitment)
}

// ExecuteOrder is a paid mutator transaction binding the contract method 0xa1a870d9.
//
// Solidity: function executeOrder(uint256 _orderId, uint256[4] _proof, bytes32 _commitment) returns()
func (_FileMarket *FileMarketTransactorSession) ExecuteOrder(_orderId *big.Int, _proof [4]*big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _FileMarket.Contract.ExecuteOrder(&_FileMarket.TransactOpts, _orderId, _proof, _commitment)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _nodeStaking, address _poiVerifier, address _fspVerifier, address _keyleakVerifier) returns()
func (_FileMarket *FileMarketTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _nodeStaking common.Address, _poiVerifier common.Address, _fspVerifier common.Address, _keyleakVerifier common.Address) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "initialize", _owner, _nodeStaking, _poiVerifier, _fspVerifier, _keyleakVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _nodeStaking, address _poiVerifier, address _fspVerifier, address _keyleakVerifier) returns()
func (_FileMarket *FileMarketSession) Initialize(_owner common.Address, _nodeStaking common.Address, _poiVerifier common.Address, _fspVerifier common.Address, _keyleakVerifier common.Address) (*types.Transaction, error) {
	return _FileMarket.Contract.Initialize(&_FileMarket.TransactOpts, _owner, _nodeStaking, _poiVerifier, _fspVerifier, _keyleakVerifier)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _nodeStaking, address _poiVerifier, address _fspVerifier, address _keyleakVerifier) returns()
func (_FileMarket *FileMarketTransactorSession) Initialize(_owner common.Address, _nodeStaking common.Address, _poiVerifier common.Address, _fspVerifier common.Address, _keyleakVerifier common.Address) (*types.Transaction, error) {
	return _FileMarket.Contract.Initialize(&_FileMarket.TransactOpts, _owner, _nodeStaking, _poiVerifier, _fspVerifier, _keyleakVerifier)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x678d66a9.
//
// Solidity: function placeOrder((uint256,string) _file, uint32 _numChunks, uint16 _periods, uint8 _replicas, uint256 _pricePerChunkPerPeriod, uint256[4] _fspProof) payable returns(uint256 orderId)
func (_FileMarket *FileMarketTransactor) PlaceOrder(opts *bind.TransactOpts, _file MarketStorageFileMeta, _numChunks uint32, _periods uint16, _replicas uint8, _pricePerChunkPerPeriod *big.Int, _fspProof [4]*big.Int) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "placeOrder", _file, _numChunks, _periods, _replicas, _pricePerChunkPerPeriod, _fspProof)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x678d66a9.
//
// Solidity: function placeOrder((uint256,string) _file, uint32 _numChunks, uint16 _periods, uint8 _replicas, uint256 _pricePerChunkPerPeriod, uint256[4] _fspProof) payable returns(uint256 orderId)
func (_FileMarket *FileMarketSession) PlaceOrder(_file MarketStorageFileMeta, _numChunks uint32, _periods uint16, _replicas uint8, _pricePerChunkPerPeriod *big.Int, _fspProof [4]*big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.PlaceOrder(&_FileMarket.TransactOpts, _file, _numChunks, _periods, _replicas, _pricePerChunkPerPeriod, _fspProof)
}

// PlaceOrder is a paid mutator transaction binding the contract method 0x678d66a9.
//
// Solidity: function placeOrder((uint256,string) _file, uint32 _numChunks, uint16 _periods, uint8 _replicas, uint256 _pricePerChunkPerPeriod, uint256[4] _fspProof) payable returns(uint256 orderId)
func (_FileMarket *FileMarketTransactorSession) PlaceOrder(_file MarketStorageFileMeta, _numChunks uint32, _periods uint16, _replicas uint8, _pricePerChunkPerPeriod *big.Int, _fspProof [4]*big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.PlaceOrder(&_FileMarket.TransactOpts, _file, _numChunks, _periods, _replicas, _pricePerChunkPerPeriod, _fspProof)
}

// ProcessExpiredOnDemandChallenge is a paid mutator transaction binding the contract method 0xac57aeb6.
//
// Solidity: function processExpiredOnDemandChallenge(uint256 _orderId, address _node) returns()
func (_FileMarket *FileMarketTransactor) ProcessExpiredOnDemandChallenge(opts *bind.TransactOpts, _orderId *big.Int, _node common.Address) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "processExpiredOnDemandChallenge", _orderId, _node)
}

// ProcessExpiredOnDemandChallenge is a paid mutator transaction binding the contract method 0xac57aeb6.
//
// Solidity: function processExpiredOnDemandChallenge(uint256 _orderId, address _node) returns()
func (_FileMarket *FileMarketSession) ProcessExpiredOnDemandChallenge(_orderId *big.Int, _node common.Address) (*types.Transaction, error) {
	return _FileMarket.Contract.ProcessExpiredOnDemandChallenge(&_FileMarket.TransactOpts, _orderId, _node)
}

// ProcessExpiredOnDemandChallenge is a paid mutator transaction binding the contract method 0xac57aeb6.
//
// Solidity: function processExpiredOnDemandChallenge(uint256 _orderId, address _node) returns()
func (_FileMarket *FileMarketTransactorSession) ProcessExpiredOnDemandChallenge(_orderId *big.Int, _node common.Address) (*types.Transaction, error) {
	return _FileMarket.Contract.ProcessExpiredOnDemandChallenge(&_FileMarket.TransactOpts, _orderId, _node)
}

// ProcessExpiredSlots is a paid mutator transaction binding the contract method 0xdbbbe766.
//
// Solidity: function processExpiredSlots() returns()
func (_FileMarket *FileMarketTransactor) ProcessExpiredSlots(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "processExpiredSlots")
}

// ProcessExpiredSlots is a paid mutator transaction binding the contract method 0xdbbbe766.
//
// Solidity: function processExpiredSlots() returns()
func (_FileMarket *FileMarketSession) ProcessExpiredSlots() (*types.Transaction, error) {
	return _FileMarket.Contract.ProcessExpiredSlots(&_FileMarket.TransactOpts)
}

// ProcessExpiredSlots is a paid mutator transaction binding the contract method 0xdbbbe766.
//
// Solidity: function processExpiredSlots() returns()
func (_FileMarket *FileMarketTransactorSession) ProcessExpiredSlots() (*types.Transaction, error) {
	return _FileMarket.Contract.ProcessExpiredSlots(&_FileMarket.TransactOpts)
}

// QuitOrder is a paid mutator transaction binding the contract method 0xd2926b21.
//
// Solidity: function quitOrder(uint256 _orderId) returns()
func (_FileMarket *FileMarketTransactor) QuitOrder(opts *bind.TransactOpts, _orderId *big.Int) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "quitOrder", _orderId)
}

// QuitOrder is a paid mutator transaction binding the contract method 0xd2926b21.
//
// Solidity: function quitOrder(uint256 _orderId) returns()
func (_FileMarket *FileMarketSession) QuitOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.QuitOrder(&_FileMarket.TransactOpts, _orderId)
}

// QuitOrder is a paid mutator transaction binding the contract method 0xd2926b21.
//
// Solidity: function quitOrder(uint256 _orderId) returns()
func (_FileMarket *FileMarketTransactorSession) QuitOrder(_orderId *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.QuitOrder(&_FileMarket.TransactOpts, _orderId)
}

// ReportKeyLeak is a paid mutator transaction binding the contract method 0x7d9d6998.
//
// Solidity: function reportKeyLeak(address _node, bytes _proof) returns()
func (_FileMarket *FileMarketTransactor) ReportKeyLeak(opts *bind.TransactOpts, _node common.Address, _proof []byte) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "reportKeyLeak", _node, _proof)
}

// ReportKeyLeak is a paid mutator transaction binding the contract method 0x7d9d6998.
//
// Solidity: function reportKeyLeak(address _node, bytes _proof) returns()
func (_FileMarket *FileMarketSession) ReportKeyLeak(_node common.Address, _proof []byte) (*types.Transaction, error) {
	return _FileMarket.Contract.ReportKeyLeak(&_FileMarket.TransactOpts, _node, _proof)
}

// ReportKeyLeak is a paid mutator transaction binding the contract method 0x7d9d6998.
//
// Solidity: function reportKeyLeak(address _node, bytes _proof) returns()
func (_FileMarket *FileMarketTransactorSession) ReportKeyLeak(_node common.Address, _proof []byte) (*types.Transaction, error) {
	return _FileMarket.Contract.ReportKeyLeak(&_FileMarket.TransactOpts, _node, _proof)
}

// SetChallengeStartBlock is a paid mutator transaction binding the contract method 0x9cf4aad2.
//
// Solidity: function setChallengeStartBlock(uint256 _block) returns()
func (_FileMarket *FileMarketTransactor) SetChallengeStartBlock(opts *bind.TransactOpts, _block *big.Int) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "setChallengeStartBlock", _block)
}

// SetChallengeStartBlock is a paid mutator transaction binding the contract method 0x9cf4aad2.
//
// Solidity: function setChallengeStartBlock(uint256 _block) returns()
func (_FileMarket *FileMarketSession) SetChallengeStartBlock(_block *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.SetChallengeStartBlock(&_FileMarket.TransactOpts, _block)
}

// SetChallengeStartBlock is a paid mutator transaction binding the contract method 0x9cf4aad2.
//
// Solidity: function setChallengeStartBlock(uint256 _block) returns()
func (_FileMarket *FileMarketTransactorSession) SetChallengeStartBlock(_block *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.SetChallengeStartBlock(&_FileMarket.TransactOpts, _block)
}

// SetClientCompensationBps is a paid mutator transaction binding the contract method 0x4a1d3a24.
//
// Solidity: function setClientCompensationBps(uint256 _newBps) returns()
func (_FileMarket *FileMarketTransactor) SetClientCompensationBps(opts *bind.TransactOpts, _newBps *big.Int) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "setClientCompensationBps", _newBps)
}

// SetClientCompensationBps is a paid mutator transaction binding the contract method 0x4a1d3a24.
//
// Solidity: function setClientCompensationBps(uint256 _newBps) returns()
func (_FileMarket *FileMarketSession) SetClientCompensationBps(_newBps *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.SetClientCompensationBps(&_FileMarket.TransactOpts, _newBps)
}

// SetClientCompensationBps is a paid mutator transaction binding the contract method 0x4a1d3a24.
//
// Solidity: function setClientCompensationBps(uint256 _newBps) returns()
func (_FileMarket *FileMarketTransactorSession) SetClientCompensationBps(_newBps *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.SetClientCompensationBps(&_FileMarket.TransactOpts, _newBps)
}

// SetProofFailureSlashMultiplier is a paid mutator transaction binding the contract method 0x62c10b17.
//
// Solidity: function setProofFailureSlashMultiplier(uint256 _newMultiplier) returns()
func (_FileMarket *FileMarketTransactor) SetProofFailureSlashMultiplier(opts *bind.TransactOpts, _newMultiplier *big.Int) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "setProofFailureSlashMultiplier", _newMultiplier)
}

// SetProofFailureSlashMultiplier is a paid mutator transaction binding the contract method 0x62c10b17.
//
// Solidity: function setProofFailureSlashMultiplier(uint256 _newMultiplier) returns()
func (_FileMarket *FileMarketSession) SetProofFailureSlashMultiplier(_newMultiplier *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.SetProofFailureSlashMultiplier(&_FileMarket.TransactOpts, _newMultiplier)
}

// SetProofFailureSlashMultiplier is a paid mutator transaction binding the contract method 0x62c10b17.
//
// Solidity: function setProofFailureSlashMultiplier(uint256 _newMultiplier) returns()
func (_FileMarket *FileMarketTransactorSession) SetProofFailureSlashMultiplier(_newMultiplier *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.SetProofFailureSlashMultiplier(&_FileMarket.TransactOpts, _newMultiplier)
}

// SetReporterRewardBps is a paid mutator transaction binding the contract method 0x220c4e70.
//
// Solidity: function setReporterRewardBps(uint256 _newBps) returns()
func (_FileMarket *FileMarketTransactor) SetReporterRewardBps(opts *bind.TransactOpts, _newBps *big.Int) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "setReporterRewardBps", _newBps)
}

// SetReporterRewardBps is a paid mutator transaction binding the contract method 0x220c4e70.
//
// Solidity: function setReporterRewardBps(uint256 _newBps) returns()
func (_FileMarket *FileMarketSession) SetReporterRewardBps(_newBps *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.SetReporterRewardBps(&_FileMarket.TransactOpts, _newBps)
}

// SetReporterRewardBps is a paid mutator transaction binding the contract method 0x220c4e70.
//
// Solidity: function setReporterRewardBps(uint256 _newBps) returns()
func (_FileMarket *FileMarketTransactorSession) SetReporterRewardBps(_newBps *big.Int) (*types.Transaction, error) {
	return _FileMarket.Contract.SetReporterRewardBps(&_FileMarket.TransactOpts, _newBps)
}

// SetSlashAuthority is a paid mutator transaction binding the contract method 0x73f1cb7c.
//
// Solidity: function setSlashAuthority(address _authority, bool _allowed) returns()
func (_FileMarket *FileMarketTransactor) SetSlashAuthority(opts *bind.TransactOpts, _authority common.Address, _allowed bool) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "setSlashAuthority", _authority, _allowed)
}

// SetSlashAuthority is a paid mutator transaction binding the contract method 0x73f1cb7c.
//
// Solidity: function setSlashAuthority(address _authority, bool _allowed) returns()
func (_FileMarket *FileMarketSession) SetSlashAuthority(_authority common.Address, _allowed bool) (*types.Transaction, error) {
	return _FileMarket.Contract.SetSlashAuthority(&_FileMarket.TransactOpts, _authority, _allowed)
}

// SetSlashAuthority is a paid mutator transaction binding the contract method 0x73f1cb7c.
//
// Solidity: function setSlashAuthority(address _authority, bool _allowed) returns()
func (_FileMarket *FileMarketTransactorSession) SetSlashAuthority(_authority common.Address, _allowed bool) (*types.Transaction, error) {
	return _FileMarket.Contract.SetSlashAuthority(&_FileMarket.TransactOpts, _authority, _allowed)
}

// SlashNode is a paid mutator transaction binding the contract method 0x31968899.
//
// Solidity: function slashNode(address _node, uint256 _slashAmount, string _reason) returns()
func (_FileMarket *FileMarketTransactor) SlashNode(opts *bind.TransactOpts, _node common.Address, _slashAmount *big.Int, _reason string) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "slashNode", _node, _slashAmount, _reason)
}

// SlashNode is a paid mutator transaction binding the contract method 0x31968899.
//
// Solidity: function slashNode(address _node, uint256 _slashAmount, string _reason) returns()
func (_FileMarket *FileMarketSession) SlashNode(_node common.Address, _slashAmount *big.Int, _reason string) (*types.Transaction, error) {
	return _FileMarket.Contract.SlashNode(&_FileMarket.TransactOpts, _node, _slashAmount, _reason)
}

// SlashNode is a paid mutator transaction binding the contract method 0x31968899.
//
// Solidity: function slashNode(address _node, uint256 _slashAmount, string _reason) returns()
func (_FileMarket *FileMarketTransactorSession) SlashNode(_node common.Address, _slashAmount *big.Int, _reason string) (*types.Transaction, error) {
	return _FileMarket.Contract.SlashNode(&_FileMarket.TransactOpts, _node, _slashAmount, _reason)
}

// SubmitOnDemandProof is a paid mutator transaction binding the contract method 0x5589d60a.
//
// Solidity: function submitOnDemandProof(uint256 _orderId, uint256[4] _proof, bytes32 _commitment) returns()
func (_FileMarket *FileMarketTransactor) SubmitOnDemandProof(opts *bind.TransactOpts, _orderId *big.Int, _proof [4]*big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "submitOnDemandProof", _orderId, _proof, _commitment)
}

// SubmitOnDemandProof is a paid mutator transaction binding the contract method 0x5589d60a.
//
// Solidity: function submitOnDemandProof(uint256 _orderId, uint256[4] _proof, bytes32 _commitment) returns()
func (_FileMarket *FileMarketSession) SubmitOnDemandProof(_orderId *big.Int, _proof [4]*big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _FileMarket.Contract.SubmitOnDemandProof(&_FileMarket.TransactOpts, _orderId, _proof, _commitment)
}

// SubmitOnDemandProof is a paid mutator transaction binding the contract method 0x5589d60a.
//
// Solidity: function submitOnDemandProof(uint256 _orderId, uint256[4] _proof, bytes32 _commitment) returns()
func (_FileMarket *FileMarketTransactorSession) SubmitOnDemandProof(_orderId *big.Int, _proof [4]*big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _FileMarket.Contract.SubmitOnDemandProof(&_FileMarket.TransactOpts, _orderId, _proof, _commitment)
}

// SubmitProof is a paid mutator transaction binding the contract method 0xe68f31e3.
//
// Solidity: function submitProof(uint256 _slotIndex, uint256[4] _proof, bytes32 _commitment) returns()
func (_FileMarket *FileMarketTransactor) SubmitProof(opts *bind.TransactOpts, _slotIndex *big.Int, _proof [4]*big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "submitProof", _slotIndex, _proof, _commitment)
}

// SubmitProof is a paid mutator transaction binding the contract method 0xe68f31e3.
//
// Solidity: function submitProof(uint256 _slotIndex, uint256[4] _proof, bytes32 _commitment) returns()
func (_FileMarket *FileMarketSession) SubmitProof(_slotIndex *big.Int, _proof [4]*big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _FileMarket.Contract.SubmitProof(&_FileMarket.TransactOpts, _slotIndex, _proof, _commitment)
}

// SubmitProof is a paid mutator transaction binding the contract method 0xe68f31e3.
//
// Solidity: function submitProof(uint256 _slotIndex, uint256[4] _proof, bytes32 _commitment) returns()
func (_FileMarket *FileMarketTransactorSession) SubmitProof(_slotIndex *big.Int, _proof [4]*big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _FileMarket.Contract.SubmitProof(&_FileMarket.TransactOpts, _slotIndex, _proof, _commitment)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_FileMarket *FileMarketTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_FileMarket *FileMarketSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _FileMarket.Contract.TransferOwnership(&_FileMarket.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_FileMarket *FileMarketTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _FileMarket.Contract.TransferOwnership(&_FileMarket.TransactOpts, _newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FileMarket *FileMarketTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FileMarket *FileMarketSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FileMarket.Contract.UpgradeToAndCall(&_FileMarket.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_FileMarket *FileMarketTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _FileMarket.Contract.UpgradeToAndCall(&_FileMarket.TransactOpts, newImplementation, data)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x110f8874.
//
// Solidity: function withdrawRefund() returns()
func (_FileMarket *FileMarketTransactor) WithdrawRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileMarket.contract.Transact(opts, "withdrawRefund")
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x110f8874.
//
// Solidity: function withdrawRefund() returns()
func (_FileMarket *FileMarketSession) WithdrawRefund() (*types.Transaction, error) {
	return _FileMarket.Contract.WithdrawRefund(&_FileMarket.TransactOpts)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x110f8874.
//
// Solidity: function withdrawRefund() returns()
func (_FileMarket *FileMarketTransactorSession) WithdrawRefund() (*types.Transaction, error) {
	return _FileMarket.Contract.WithdrawRefund(&_FileMarket.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FileMarket *FileMarketTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FileMarket.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FileMarket *FileMarketSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FileMarket.Contract.Fallback(&_FileMarket.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FileMarket *FileMarketTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FileMarket.Contract.Fallback(&_FileMarket.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FileMarket *FileMarketTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FileMarket.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FileMarket *FileMarketSession) Receive() (*types.Transaction, error) {
	return _FileMarket.Contract.Receive(&_FileMarket.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FileMarket *FileMarketTransactorSession) Receive() (*types.Transaction, error) {
	return _FileMarket.Contract.Receive(&_FileMarket.TransactOpts)
}

// FileMarketCancellationPenaltyDistributedIterator is returned from FilterCancellationPenaltyDistributed and is used to iterate over the raw logs and unpacked data for CancellationPenaltyDistributed events raised by the FileMarket contract.
type FileMarketCancellationPenaltyDistributedIterator struct {
	Event *FileMarketCancellationPenaltyDistributed // Event containing the contract specifics and raw log

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
func (it *FileMarketCancellationPenaltyDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketCancellationPenaltyDistributed)
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
		it.Event = new(FileMarketCancellationPenaltyDistributed)
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
func (it *FileMarketCancellationPenaltyDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketCancellationPenaltyDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketCancellationPenaltyDistributed represents a CancellationPenaltyDistributed event raised by the FileMarket contract.
type FileMarketCancellationPenaltyDistributed struct {
	OrderId       *big.Int
	PenaltyAmount *big.Int
	NodeCount     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCancellationPenaltyDistributed is a free log retrieval operation binding the contract event 0xc3ed9acdaa58da46e7cd1e5fe6ad7d78c6cf8a1fd776846beb025ddb7bf08000.
//
// Solidity: event CancellationPenaltyDistributed(uint256 indexed orderId, uint256 penaltyAmount, uint256 nodeCount)
func (_FileMarket *FileMarketFilterer) FilterCancellationPenaltyDistributed(opts *bind.FilterOpts, orderId []*big.Int) (*FileMarketCancellationPenaltyDistributedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "CancellationPenaltyDistributed", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketCancellationPenaltyDistributedIterator{contract: _FileMarket.contract, event: "CancellationPenaltyDistributed", logs: logs, sub: sub}, nil
}

// WatchCancellationPenaltyDistributed is a free log subscription operation binding the contract event 0xc3ed9acdaa58da46e7cd1e5fe6ad7d78c6cf8a1fd776846beb025ddb7bf08000.
//
// Solidity: event CancellationPenaltyDistributed(uint256 indexed orderId, uint256 penaltyAmount, uint256 nodeCount)
func (_FileMarket *FileMarketFilterer) WatchCancellationPenaltyDistributed(opts *bind.WatchOpts, sink chan<- *FileMarketCancellationPenaltyDistributed, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "CancellationPenaltyDistributed", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketCancellationPenaltyDistributed)
				if err := _FileMarket.contract.UnpackLog(event, "CancellationPenaltyDistributed", log); err != nil {
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

// ParseCancellationPenaltyDistributed is a log parse operation binding the contract event 0xc3ed9acdaa58da46e7cd1e5fe6ad7d78c6cf8a1fd776846beb025ddb7bf08000.
//
// Solidity: event CancellationPenaltyDistributed(uint256 indexed orderId, uint256 penaltyAmount, uint256 nodeCount)
func (_FileMarket *FileMarketFilterer) ParseCancellationPenaltyDistributed(log types.Log) (*FileMarketCancellationPenaltyDistributed, error) {
	event := new(FileMarketCancellationPenaltyDistributed)
	if err := _FileMarket.contract.UnpackLog(event, "CancellationPenaltyDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketChallengeSlotsScaledIterator is returned from FilterChallengeSlotsScaled and is used to iterate over the raw logs and unpacked data for ChallengeSlotsScaled events raised by the FileMarket contract.
type FileMarketChallengeSlotsScaledIterator struct {
	Event *FileMarketChallengeSlotsScaled // Event containing the contract specifics and raw log

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
func (it *FileMarketChallengeSlotsScaledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketChallengeSlotsScaled)
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
		it.Event = new(FileMarketChallengeSlotsScaled)
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
func (it *FileMarketChallengeSlotsScaledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketChallengeSlotsScaledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketChallengeSlotsScaled represents a ChallengeSlotsScaled event raised by the FileMarket contract.
type FileMarketChallengeSlotsScaled struct {
	OldCount *big.Int
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChallengeSlotsScaled is a free log retrieval operation binding the contract event 0x9914f1952e5a1f5b850acc0bb516eaf87ad1f31474736a2d93d5af41ec6ce6bf.
//
// Solidity: event ChallengeSlotsScaled(uint256 oldCount, uint256 newCount)
func (_FileMarket *FileMarketFilterer) FilterChallengeSlotsScaled(opts *bind.FilterOpts) (*FileMarketChallengeSlotsScaledIterator, error) {

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "ChallengeSlotsScaled")
	if err != nil {
		return nil, err
	}
	return &FileMarketChallengeSlotsScaledIterator{contract: _FileMarket.contract, event: "ChallengeSlotsScaled", logs: logs, sub: sub}, nil
}

// WatchChallengeSlotsScaled is a free log subscription operation binding the contract event 0x9914f1952e5a1f5b850acc0bb516eaf87ad1f31474736a2d93d5af41ec6ce6bf.
//
// Solidity: event ChallengeSlotsScaled(uint256 oldCount, uint256 newCount)
func (_FileMarket *FileMarketFilterer) WatchChallengeSlotsScaled(opts *bind.WatchOpts, sink chan<- *FileMarketChallengeSlotsScaled) (event.Subscription, error) {

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "ChallengeSlotsScaled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketChallengeSlotsScaled)
				if err := _FileMarket.contract.UnpackLog(event, "ChallengeSlotsScaled", log); err != nil {
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

// ParseChallengeSlotsScaled is a log parse operation binding the contract event 0x9914f1952e5a1f5b850acc0bb516eaf87ad1f31474736a2d93d5af41ec6ce6bf.
//
// Solidity: event ChallengeSlotsScaled(uint256 oldCount, uint256 newCount)
func (_FileMarket *FileMarketFilterer) ParseChallengeSlotsScaled(log types.Log) (*FileMarketChallengeSlotsScaled, error) {
	event := new(FileMarketChallengeSlotsScaled)
	if err := _FileMarket.contract.UnpackLog(event, "ChallengeSlotsScaled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketClientCompensationAccruedIterator is returned from FilterClientCompensationAccrued and is used to iterate over the raw logs and unpacked data for ClientCompensationAccrued events raised by the FileMarket contract.
type FileMarketClientCompensationAccruedIterator struct {
	Event *FileMarketClientCompensationAccrued // Event containing the contract specifics and raw log

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
func (it *FileMarketClientCompensationAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketClientCompensationAccrued)
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
		it.Event = new(FileMarketClientCompensationAccrued)
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
func (it *FileMarketClientCompensationAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketClientCompensationAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketClientCompensationAccrued represents a ClientCompensationAccrued event raised by the FileMarket contract.
type FileMarketClientCompensationAccrued struct {
	Client  common.Address
	Amount  *big.Int
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClientCompensationAccrued is a free log retrieval operation binding the contract event 0x2d0475ca2e6724882218f8f67bb5587ba2f6b3e1618b800a78c8cf87028fc5c7.
//
// Solidity: event ClientCompensationAccrued(address indexed client, uint256 amount, uint256 orderId)
func (_FileMarket *FileMarketFilterer) FilterClientCompensationAccrued(opts *bind.FilterOpts, client []common.Address) (*FileMarketClientCompensationAccruedIterator, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "ClientCompensationAccrued", clientRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketClientCompensationAccruedIterator{contract: _FileMarket.contract, event: "ClientCompensationAccrued", logs: logs, sub: sub}, nil
}

// WatchClientCompensationAccrued is a free log subscription operation binding the contract event 0x2d0475ca2e6724882218f8f67bb5587ba2f6b3e1618b800a78c8cf87028fc5c7.
//
// Solidity: event ClientCompensationAccrued(address indexed client, uint256 amount, uint256 orderId)
func (_FileMarket *FileMarketFilterer) WatchClientCompensationAccrued(opts *bind.WatchOpts, sink chan<- *FileMarketClientCompensationAccrued, client []common.Address) (event.Subscription, error) {

	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "ClientCompensationAccrued", clientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketClientCompensationAccrued)
				if err := _FileMarket.contract.UnpackLog(event, "ClientCompensationAccrued", log); err != nil {
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

// ParseClientCompensationAccrued is a log parse operation binding the contract event 0x2d0475ca2e6724882218f8f67bb5587ba2f6b3e1618b800a78c8cf87028fc5c7.
//
// Solidity: event ClientCompensationAccrued(address indexed client, uint256 amount, uint256 orderId)
func (_FileMarket *FileMarketFilterer) ParseClientCompensationAccrued(log types.Log) (*FileMarketClientCompensationAccrued, error) {
	event := new(FileMarketClientCompensationAccrued)
	if err := _FileMarket.contract.UnpackLog(event, "ClientCompensationAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketClientCompensationBpsUpdatedIterator is returned from FilterClientCompensationBpsUpdated and is used to iterate over the raw logs and unpacked data for ClientCompensationBpsUpdated events raised by the FileMarket contract.
type FileMarketClientCompensationBpsUpdatedIterator struct {
	Event *FileMarketClientCompensationBpsUpdated // Event containing the contract specifics and raw log

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
func (it *FileMarketClientCompensationBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketClientCompensationBpsUpdated)
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
		it.Event = new(FileMarketClientCompensationBpsUpdated)
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
func (it *FileMarketClientCompensationBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketClientCompensationBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketClientCompensationBpsUpdated represents a ClientCompensationBpsUpdated event raised by the FileMarket contract.
type FileMarketClientCompensationBpsUpdated struct {
	OldBps *big.Int
	NewBps *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClientCompensationBpsUpdated is a free log retrieval operation binding the contract event 0x2c1965abc1a15203a0dd33a6f74e08e6a83bc49c98927be64355f2f3299f840f.
//
// Solidity: event ClientCompensationBpsUpdated(uint256 oldBps, uint256 newBps)
func (_FileMarket *FileMarketFilterer) FilterClientCompensationBpsUpdated(opts *bind.FilterOpts) (*FileMarketClientCompensationBpsUpdatedIterator, error) {

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "ClientCompensationBpsUpdated")
	if err != nil {
		return nil, err
	}
	return &FileMarketClientCompensationBpsUpdatedIterator{contract: _FileMarket.contract, event: "ClientCompensationBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchClientCompensationBpsUpdated is a free log subscription operation binding the contract event 0x2c1965abc1a15203a0dd33a6f74e08e6a83bc49c98927be64355f2f3299f840f.
//
// Solidity: event ClientCompensationBpsUpdated(uint256 oldBps, uint256 newBps)
func (_FileMarket *FileMarketFilterer) WatchClientCompensationBpsUpdated(opts *bind.WatchOpts, sink chan<- *FileMarketClientCompensationBpsUpdated) (event.Subscription, error) {

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "ClientCompensationBpsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketClientCompensationBpsUpdated)
				if err := _FileMarket.contract.UnpackLog(event, "ClientCompensationBpsUpdated", log); err != nil {
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

// ParseClientCompensationBpsUpdated is a log parse operation binding the contract event 0x2c1965abc1a15203a0dd33a6f74e08e6a83bc49c98927be64355f2f3299f840f.
//
// Solidity: event ClientCompensationBpsUpdated(uint256 oldBps, uint256 newBps)
func (_FileMarket *FileMarketFilterer) ParseClientCompensationBpsUpdated(log types.Log) (*FileMarketClientCompensationBpsUpdated, error) {
	event := new(FileMarketClientCompensationBpsUpdated)
	if err := _FileMarket.contract.UnpackLog(event, "ClientCompensationBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketExpiredSlotsProcessedIterator is returned from FilterExpiredSlotsProcessed and is used to iterate over the raw logs and unpacked data for ExpiredSlotsProcessed events raised by the FileMarket contract.
type FileMarketExpiredSlotsProcessedIterator struct {
	Event *FileMarketExpiredSlotsProcessed // Event containing the contract specifics and raw log

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
func (it *FileMarketExpiredSlotsProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketExpiredSlotsProcessed)
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
		it.Event = new(FileMarketExpiredSlotsProcessed)
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
func (it *FileMarketExpiredSlotsProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketExpiredSlotsProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketExpiredSlotsProcessed represents a ExpiredSlotsProcessed event raised by the FileMarket contract.
type FileMarketExpiredSlotsProcessed struct {
	ProcessedCount *big.Int
	Reporter       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExpiredSlotsProcessed is a free log retrieval operation binding the contract event 0x2243b3b8bc650df05b78f40549c8c709f116651ff3b3d87276eca1fd01e66be5.
//
// Solidity: event ExpiredSlotsProcessed(uint256 processedCount, address indexed reporter)
func (_FileMarket *FileMarketFilterer) FilterExpiredSlotsProcessed(opts *bind.FilterOpts, reporter []common.Address) (*FileMarketExpiredSlotsProcessedIterator, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "ExpiredSlotsProcessed", reporterRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketExpiredSlotsProcessedIterator{contract: _FileMarket.contract, event: "ExpiredSlotsProcessed", logs: logs, sub: sub}, nil
}

// WatchExpiredSlotsProcessed is a free log subscription operation binding the contract event 0x2243b3b8bc650df05b78f40549c8c709f116651ff3b3d87276eca1fd01e66be5.
//
// Solidity: event ExpiredSlotsProcessed(uint256 processedCount, address indexed reporter)
func (_FileMarket *FileMarketFilterer) WatchExpiredSlotsProcessed(opts *bind.WatchOpts, sink chan<- *FileMarketExpiredSlotsProcessed, reporter []common.Address) (event.Subscription, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "ExpiredSlotsProcessed", reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketExpiredSlotsProcessed)
				if err := _FileMarket.contract.UnpackLog(event, "ExpiredSlotsProcessed", log); err != nil {
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

// ParseExpiredSlotsProcessed is a log parse operation binding the contract event 0x2243b3b8bc650df05b78f40549c8c709f116651ff3b3d87276eca1fd01e66be5.
//
// Solidity: event ExpiredSlotsProcessed(uint256 processedCount, address indexed reporter)
func (_FileMarket *FileMarketFilterer) ParseExpiredSlotsProcessed(log types.Log) (*FileMarketExpiredSlotsProcessed, error) {
	event := new(FileMarketExpiredSlotsProcessed)
	if err := _FileMarket.contract.UnpackLog(event, "ExpiredSlotsProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketForcedOrderExitsIterator is returned from FilterForcedOrderExits and is used to iterate over the raw logs and unpacked data for ForcedOrderExits events raised by the FileMarket contract.
type FileMarketForcedOrderExitsIterator struct {
	Event *FileMarketForcedOrderExits // Event containing the contract specifics and raw log

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
func (it *FileMarketForcedOrderExitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketForcedOrderExits)
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
		it.Event = new(FileMarketForcedOrderExits)
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
func (it *FileMarketForcedOrderExitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketForcedOrderExitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketForcedOrderExits represents a ForcedOrderExits event raised by the FileMarket contract.
type FileMarketForcedOrderExits struct {
	Node       common.Address
	OrderIds   []*big.Int
	TotalFreed uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterForcedOrderExits is a free log retrieval operation binding the contract event 0x8f2894130ea2030245ae8931b2570b7e55e1f78fbf170338c8b67b4f3f12750c.
//
// Solidity: event ForcedOrderExits(address indexed node, uint256[] orderIds, uint64 totalFreed)
func (_FileMarket *FileMarketFilterer) FilterForcedOrderExits(opts *bind.FilterOpts, node []common.Address) (*FileMarketForcedOrderExitsIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "ForcedOrderExits", nodeRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketForcedOrderExitsIterator{contract: _FileMarket.contract, event: "ForcedOrderExits", logs: logs, sub: sub}, nil
}

// WatchForcedOrderExits is a free log subscription operation binding the contract event 0x8f2894130ea2030245ae8931b2570b7e55e1f78fbf170338c8b67b4f3f12750c.
//
// Solidity: event ForcedOrderExits(address indexed node, uint256[] orderIds, uint64 totalFreed)
func (_FileMarket *FileMarketFilterer) WatchForcedOrderExits(opts *bind.WatchOpts, sink chan<- *FileMarketForcedOrderExits, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "ForcedOrderExits", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketForcedOrderExits)
				if err := _FileMarket.contract.UnpackLog(event, "ForcedOrderExits", log); err != nil {
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

// ParseForcedOrderExits is a log parse operation binding the contract event 0x8f2894130ea2030245ae8931b2570b7e55e1f78fbf170338c8b67b4f3f12750c.
//
// Solidity: event ForcedOrderExits(address indexed node, uint256[] orderIds, uint64 totalFreed)
func (_FileMarket *FileMarketFilterer) ParseForcedOrderExits(log types.Log) (*FileMarketForcedOrderExits, error) {
	event := new(FileMarketForcedOrderExits)
	if err := _FileMarket.contract.UnpackLog(event, "ForcedOrderExits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FileMarket contract.
type FileMarketInitializedIterator struct {
	Event *FileMarketInitialized // Event containing the contract specifics and raw log

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
func (it *FileMarketInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketInitialized)
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
		it.Event = new(FileMarketInitialized)
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
func (it *FileMarketInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketInitialized represents a Initialized event raised by the FileMarket contract.
type FileMarketInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FileMarket *FileMarketFilterer) FilterInitialized(opts *bind.FilterOpts) (*FileMarketInitializedIterator, error) {

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FileMarketInitializedIterator{contract: _FileMarket.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FileMarket *FileMarketFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FileMarketInitialized) (event.Subscription, error) {

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketInitialized)
				if err := _FileMarket.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FileMarket *FileMarketFilterer) ParseInitialized(log types.Log) (*FileMarketInitialized, error) {
	event := new(FileMarketInitialized)
	if err := _FileMarket.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketKeyLeakReportedIterator is returned from FilterKeyLeakReported and is used to iterate over the raw logs and unpacked data for KeyLeakReported events raised by the FileMarket contract.
type FileMarketKeyLeakReportedIterator struct {
	Event *FileMarketKeyLeakReported // Event containing the contract specifics and raw log

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
func (it *FileMarketKeyLeakReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketKeyLeakReported)
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
		it.Event = new(FileMarketKeyLeakReported)
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
func (it *FileMarketKeyLeakReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketKeyLeakReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketKeyLeakReported represents a KeyLeakReported event raised by the FileMarket contract.
type FileMarketKeyLeakReported struct {
	Node        common.Address
	Reporter    common.Address
	SlashAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterKeyLeakReported is a free log retrieval operation binding the contract event 0x594c1ba5adf566d661cd3f871ab024db28906aa6a5c794b28e1636569282542e.
//
// Solidity: event KeyLeakReported(address indexed node, address indexed reporter, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) FilterKeyLeakReported(opts *bind.FilterOpts, node []common.Address, reporter []common.Address) (*FileMarketKeyLeakReportedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "KeyLeakReported", nodeRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketKeyLeakReportedIterator{contract: _FileMarket.contract, event: "KeyLeakReported", logs: logs, sub: sub}, nil
}

// WatchKeyLeakReported is a free log subscription operation binding the contract event 0x594c1ba5adf566d661cd3f871ab024db28906aa6a5c794b28e1636569282542e.
//
// Solidity: event KeyLeakReported(address indexed node, address indexed reporter, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) WatchKeyLeakReported(opts *bind.WatchOpts, sink chan<- *FileMarketKeyLeakReported, node []common.Address, reporter []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "KeyLeakReported", nodeRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketKeyLeakReported)
				if err := _FileMarket.contract.UnpackLog(event, "KeyLeakReported", log); err != nil {
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

// ParseKeyLeakReported is a log parse operation binding the contract event 0x594c1ba5adf566d661cd3f871ab024db28906aa6a5c794b28e1636569282542e.
//
// Solidity: event KeyLeakReported(address indexed node, address indexed reporter, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) ParseKeyLeakReported(log types.Log) (*FileMarketKeyLeakReported, error) {
	event := new(FileMarketKeyLeakReported)
	if err := _FileMarket.contract.UnpackLog(event, "KeyLeakReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketNodeQuitIterator is returned from FilterNodeQuit and is used to iterate over the raw logs and unpacked data for NodeQuit events raised by the FileMarket contract.
type FileMarketNodeQuitIterator struct {
	Event *FileMarketNodeQuit // Event containing the contract specifics and raw log

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
func (it *FileMarketNodeQuitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketNodeQuit)
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
		it.Event = new(FileMarketNodeQuit)
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
func (it *FileMarketNodeQuitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketNodeQuitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketNodeQuit represents a NodeQuit event raised by the FileMarket contract.
type FileMarketNodeQuit struct {
	OrderId     *big.Int
	Node        common.Address
	SlashAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeQuit is a free log retrieval operation binding the contract event 0x9287e72144e797bd772ff1552d9478064b166948d2d3e60ec1b24837d57ec311.
//
// Solidity: event NodeQuit(uint256 indexed orderId, address indexed node, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) FilterNodeQuit(opts *bind.FilterOpts, orderId []*big.Int, node []common.Address) (*FileMarketNodeQuitIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "NodeQuit", orderIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketNodeQuitIterator{contract: _FileMarket.contract, event: "NodeQuit", logs: logs, sub: sub}, nil
}

// WatchNodeQuit is a free log subscription operation binding the contract event 0x9287e72144e797bd772ff1552d9478064b166948d2d3e60ec1b24837d57ec311.
//
// Solidity: event NodeQuit(uint256 indexed orderId, address indexed node, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) WatchNodeQuit(opts *bind.WatchOpts, sink chan<- *FileMarketNodeQuit, orderId []*big.Int, node []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "NodeQuit", orderIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketNodeQuit)
				if err := _FileMarket.contract.UnpackLog(event, "NodeQuit", log); err != nil {
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

// ParseNodeQuit is a log parse operation binding the contract event 0x9287e72144e797bd772ff1552d9478064b166948d2d3e60ec1b24837d57ec311.
//
// Solidity: event NodeQuit(uint256 indexed orderId, address indexed node, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) ParseNodeQuit(log types.Log) (*FileMarketNodeQuit, error) {
	event := new(FileMarketNodeQuit)
	if err := _FileMarket.contract.UnpackLog(event, "NodeQuit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketNodeSlashedIterator is returned from FilterNodeSlashed and is used to iterate over the raw logs and unpacked data for NodeSlashed events raised by the FileMarket contract.
type FileMarketNodeSlashedIterator struct {
	Event *FileMarketNodeSlashed // Event containing the contract specifics and raw log

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
func (it *FileMarketNodeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketNodeSlashed)
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
		it.Event = new(FileMarketNodeSlashed)
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
func (it *FileMarketNodeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketNodeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketNodeSlashed represents a NodeSlashed event raised by the FileMarket contract.
type FileMarketNodeSlashed struct {
	Node        common.Address
	SlashAmount *big.Int
	Reason      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNodeSlashed is a free log retrieval operation binding the contract event 0x992bebd8a6e574d4b441fbc6155cf5486a16f70f9003d8189f63374675c5022b.
//
// Solidity: event NodeSlashed(address indexed node, uint256 slashAmount, string reason)
func (_FileMarket *FileMarketFilterer) FilterNodeSlashed(opts *bind.FilterOpts, node []common.Address) (*FileMarketNodeSlashedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "NodeSlashed", nodeRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketNodeSlashedIterator{contract: _FileMarket.contract, event: "NodeSlashed", logs: logs, sub: sub}, nil
}

// WatchNodeSlashed is a free log subscription operation binding the contract event 0x992bebd8a6e574d4b441fbc6155cf5486a16f70f9003d8189f63374675c5022b.
//
// Solidity: event NodeSlashed(address indexed node, uint256 slashAmount, string reason)
func (_FileMarket *FileMarketFilterer) WatchNodeSlashed(opts *bind.WatchOpts, sink chan<- *FileMarketNodeSlashed, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "NodeSlashed", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketNodeSlashed)
				if err := _FileMarket.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
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

// ParseNodeSlashed is a log parse operation binding the contract event 0x992bebd8a6e574d4b441fbc6155cf5486a16f70f9003d8189f63374675c5022b.
//
// Solidity: event NodeSlashed(address indexed node, uint256 slashAmount, string reason)
func (_FileMarket *FileMarketFilterer) ParseNodeSlashed(log types.Log) (*FileMarketNodeSlashed, error) {
	event := new(FileMarketNodeSlashed)
	if err := _FileMarket.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketOnDemandChallengeExpiredIterator is returned from FilterOnDemandChallengeExpired and is used to iterate over the raw logs and unpacked data for OnDemandChallengeExpired events raised by the FileMarket contract.
type FileMarketOnDemandChallengeExpiredIterator struct {
	Event *FileMarketOnDemandChallengeExpired // Event containing the contract specifics and raw log

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
func (it *FileMarketOnDemandChallengeExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketOnDemandChallengeExpired)
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
		it.Event = new(FileMarketOnDemandChallengeExpired)
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
func (it *FileMarketOnDemandChallengeExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketOnDemandChallengeExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketOnDemandChallengeExpired represents a OnDemandChallengeExpired event raised by the FileMarket contract.
type FileMarketOnDemandChallengeExpired struct {
	OrderId     *big.Int
	Node        common.Address
	SlashAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOnDemandChallengeExpired is a free log retrieval operation binding the contract event 0x6621d07e53f9fb20fbbc721c265500a8fc8ee5fe77f1ec68857751916629c21d.
//
// Solidity: event OnDemandChallengeExpired(uint256 indexed orderId, address indexed node, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) FilterOnDemandChallengeExpired(opts *bind.FilterOpts, orderId []*big.Int, node []common.Address) (*FileMarketOnDemandChallengeExpiredIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "OnDemandChallengeExpired", orderIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketOnDemandChallengeExpiredIterator{contract: _FileMarket.contract, event: "OnDemandChallengeExpired", logs: logs, sub: sub}, nil
}

// WatchOnDemandChallengeExpired is a free log subscription operation binding the contract event 0x6621d07e53f9fb20fbbc721c265500a8fc8ee5fe77f1ec68857751916629c21d.
//
// Solidity: event OnDemandChallengeExpired(uint256 indexed orderId, address indexed node, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) WatchOnDemandChallengeExpired(opts *bind.WatchOpts, sink chan<- *FileMarketOnDemandChallengeExpired, orderId []*big.Int, node []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "OnDemandChallengeExpired", orderIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketOnDemandChallengeExpired)
				if err := _FileMarket.contract.UnpackLog(event, "OnDemandChallengeExpired", log); err != nil {
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

// ParseOnDemandChallengeExpired is a log parse operation binding the contract event 0x6621d07e53f9fb20fbbc721c265500a8fc8ee5fe77f1ec68857751916629c21d.
//
// Solidity: event OnDemandChallengeExpired(uint256 indexed orderId, address indexed node, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) ParseOnDemandChallengeExpired(log types.Log) (*FileMarketOnDemandChallengeExpired, error) {
	event := new(FileMarketOnDemandChallengeExpired)
	if err := _FileMarket.contract.UnpackLog(event, "OnDemandChallengeExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketOnDemandChallengeIssuedIterator is returned from FilterOnDemandChallengeIssued and is used to iterate over the raw logs and unpacked data for OnDemandChallengeIssued events raised by the FileMarket contract.
type FileMarketOnDemandChallengeIssuedIterator struct {
	Event *FileMarketOnDemandChallengeIssued // Event containing the contract specifics and raw log

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
func (it *FileMarketOnDemandChallengeIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketOnDemandChallengeIssued)
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
		it.Event = new(FileMarketOnDemandChallengeIssued)
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
func (it *FileMarketOnDemandChallengeIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketOnDemandChallengeIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketOnDemandChallengeIssued represents a OnDemandChallengeIssued event raised by the FileMarket contract.
type FileMarketOnDemandChallengeIssued struct {
	OrderId       *big.Int
	Node          common.Address
	Challenger    common.Address
	DeadlineBlock *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOnDemandChallengeIssued is a free log retrieval operation binding the contract event 0xf6bacafda998fd1622a9429a25dd847507ba1e556577a003510d743a400593f4.
//
// Solidity: event OnDemandChallengeIssued(uint256 indexed orderId, address indexed node, address challenger, uint256 deadlineBlock)
func (_FileMarket *FileMarketFilterer) FilterOnDemandChallengeIssued(opts *bind.FilterOpts, orderId []*big.Int, node []common.Address) (*FileMarketOnDemandChallengeIssuedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "OnDemandChallengeIssued", orderIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketOnDemandChallengeIssuedIterator{contract: _FileMarket.contract, event: "OnDemandChallengeIssued", logs: logs, sub: sub}, nil
}

// WatchOnDemandChallengeIssued is a free log subscription operation binding the contract event 0xf6bacafda998fd1622a9429a25dd847507ba1e556577a003510d743a400593f4.
//
// Solidity: event OnDemandChallengeIssued(uint256 indexed orderId, address indexed node, address challenger, uint256 deadlineBlock)
func (_FileMarket *FileMarketFilterer) WatchOnDemandChallengeIssued(opts *bind.WatchOpts, sink chan<- *FileMarketOnDemandChallengeIssued, orderId []*big.Int, node []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "OnDemandChallengeIssued", orderIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketOnDemandChallengeIssued)
				if err := _FileMarket.contract.UnpackLog(event, "OnDemandChallengeIssued", log); err != nil {
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

// ParseOnDemandChallengeIssued is a log parse operation binding the contract event 0xf6bacafda998fd1622a9429a25dd847507ba1e556577a003510d743a400593f4.
//
// Solidity: event OnDemandChallengeIssued(uint256 indexed orderId, address indexed node, address challenger, uint256 deadlineBlock)
func (_FileMarket *FileMarketFilterer) ParseOnDemandChallengeIssued(log types.Log) (*FileMarketOnDemandChallengeIssued, error) {
	event := new(FileMarketOnDemandChallengeIssued)
	if err := _FileMarket.contract.UnpackLog(event, "OnDemandChallengeIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketOnDemandProofSubmittedIterator is returned from FilterOnDemandProofSubmitted and is used to iterate over the raw logs and unpacked data for OnDemandProofSubmitted events raised by the FileMarket contract.
type FileMarketOnDemandProofSubmittedIterator struct {
	Event *FileMarketOnDemandProofSubmitted // Event containing the contract specifics and raw log

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
func (it *FileMarketOnDemandProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketOnDemandProofSubmitted)
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
		it.Event = new(FileMarketOnDemandProofSubmitted)
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
func (it *FileMarketOnDemandProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketOnDemandProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketOnDemandProofSubmitted represents a OnDemandProofSubmitted event raised by the FileMarket contract.
type FileMarketOnDemandProofSubmitted struct {
	OrderId    *big.Int
	Node       common.Address
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOnDemandProofSubmitted is a free log retrieval operation binding the contract event 0x551d83d0a8c10f67a56904f6f0bf4e4fe79403fd791df322accceeb2675512f7.
//
// Solidity: event OnDemandProofSubmitted(uint256 indexed orderId, address indexed node, bytes32 commitment)
func (_FileMarket *FileMarketFilterer) FilterOnDemandProofSubmitted(opts *bind.FilterOpts, orderId []*big.Int, node []common.Address) (*FileMarketOnDemandProofSubmittedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "OnDemandProofSubmitted", orderIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketOnDemandProofSubmittedIterator{contract: _FileMarket.contract, event: "OnDemandProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchOnDemandProofSubmitted is a free log subscription operation binding the contract event 0x551d83d0a8c10f67a56904f6f0bf4e4fe79403fd791df322accceeb2675512f7.
//
// Solidity: event OnDemandProofSubmitted(uint256 indexed orderId, address indexed node, bytes32 commitment)
func (_FileMarket *FileMarketFilterer) WatchOnDemandProofSubmitted(opts *bind.WatchOpts, sink chan<- *FileMarketOnDemandProofSubmitted, orderId []*big.Int, node []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "OnDemandProofSubmitted", orderIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketOnDemandProofSubmitted)
				if err := _FileMarket.contract.UnpackLog(event, "OnDemandProofSubmitted", log); err != nil {
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

// ParseOnDemandProofSubmitted is a log parse operation binding the contract event 0x551d83d0a8c10f67a56904f6f0bf4e4fe79403fd791df322accceeb2675512f7.
//
// Solidity: event OnDemandProofSubmitted(uint256 indexed orderId, address indexed node, bytes32 commitment)
func (_FileMarket *FileMarketFilterer) ParseOnDemandProofSubmitted(log types.Log) (*FileMarketOnDemandProofSubmitted, error) {
	event := new(FileMarketOnDemandProofSubmitted)
	if err := _FileMarket.contract.UnpackLog(event, "OnDemandProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the FileMarket contract.
type FileMarketOrderCancelledIterator struct {
	Event *FileMarketOrderCancelled // Event containing the contract specifics and raw log

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
func (it *FileMarketOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketOrderCancelled)
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
		it.Event = new(FileMarketOrderCancelled)
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
func (it *FileMarketOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketOrderCancelled represents a OrderCancelled event raised by the FileMarket contract.
type FileMarketOrderCancelled struct {
	OrderId      *big.Int
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xaa2a006d186dc89a3d6851834c85d16226329ac40bcad6531a5756a38ed1c90f.
//
// Solidity: event OrderCancelled(uint256 indexed orderId, uint256 refundAmount)
func (_FileMarket *FileMarketFilterer) FilterOrderCancelled(opts *bind.FilterOpts, orderId []*big.Int) (*FileMarketOrderCancelledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "OrderCancelled", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketOrderCancelledIterator{contract: _FileMarket.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xaa2a006d186dc89a3d6851834c85d16226329ac40bcad6531a5756a38ed1c90f.
//
// Solidity: event OrderCancelled(uint256 indexed orderId, uint256 refundAmount)
func (_FileMarket *FileMarketFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *FileMarketOrderCancelled, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "OrderCancelled", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketOrderCancelled)
				if err := _FileMarket.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0xaa2a006d186dc89a3d6851834c85d16226329ac40bcad6531a5756a38ed1c90f.
//
// Solidity: event OrderCancelled(uint256 indexed orderId, uint256 refundAmount)
func (_FileMarket *FileMarketFilterer) ParseOrderCancelled(log types.Log) (*FileMarketOrderCancelled, error) {
	event := new(FileMarketOrderCancelled)
	if err := _FileMarket.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketOrderCompletedIterator is returned from FilterOrderCompleted and is used to iterate over the raw logs and unpacked data for OrderCompleted events raised by the FileMarket contract.
type FileMarketOrderCompletedIterator struct {
	Event *FileMarketOrderCompleted // Event containing the contract specifics and raw log

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
func (it *FileMarketOrderCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketOrderCompleted)
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
		it.Event = new(FileMarketOrderCompleted)
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
func (it *FileMarketOrderCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketOrderCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketOrderCompleted represents a OrderCompleted event raised by the FileMarket contract.
type FileMarketOrderCompleted struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCompleted is a free log retrieval operation binding the contract event 0x02a7d4d9472af7118644884b9a0ee443540c0027e938dd0aa35be8ecbe946c0a.
//
// Solidity: event OrderCompleted(uint256 indexed orderId)
func (_FileMarket *FileMarketFilterer) FilterOrderCompleted(opts *bind.FilterOpts, orderId []*big.Int) (*FileMarketOrderCompletedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "OrderCompleted", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketOrderCompletedIterator{contract: _FileMarket.contract, event: "OrderCompleted", logs: logs, sub: sub}, nil
}

// WatchOrderCompleted is a free log subscription operation binding the contract event 0x02a7d4d9472af7118644884b9a0ee443540c0027e938dd0aa35be8ecbe946c0a.
//
// Solidity: event OrderCompleted(uint256 indexed orderId)
func (_FileMarket *FileMarketFilterer) WatchOrderCompleted(opts *bind.WatchOpts, sink chan<- *FileMarketOrderCompleted, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "OrderCompleted", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketOrderCompleted)
				if err := _FileMarket.contract.UnpackLog(event, "OrderCompleted", log); err != nil {
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

// ParseOrderCompleted is a log parse operation binding the contract event 0x02a7d4d9472af7118644884b9a0ee443540c0027e938dd0aa35be8ecbe946c0a.
//
// Solidity: event OrderCompleted(uint256 indexed orderId)
func (_FileMarket *FileMarketFilterer) ParseOrderCompleted(log types.Log) (*FileMarketOrderCompleted, error) {
	event := new(FileMarketOrderCompleted)
	if err := _FileMarket.contract.UnpackLog(event, "OrderCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the FileMarket contract.
type FileMarketOrderFulfilledIterator struct {
	Event *FileMarketOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *FileMarketOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketOrderFulfilled)
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
		it.Event = new(FileMarketOrderFulfilled)
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
func (it *FileMarketOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketOrderFulfilled represents a OrderFulfilled event raised by the FileMarket contract.
type FileMarketOrderFulfilled struct {
	OrderId *big.Int
	Node    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0xe5f48e2a4a303d196ad814389aa54cf34e7b1e404e740599d005907a2ca0cd6c.
//
// Solidity: event OrderFulfilled(uint256 indexed orderId, address indexed node)
func (_FileMarket *FileMarketFilterer) FilterOrderFulfilled(opts *bind.FilterOpts, orderId []*big.Int, node []common.Address) (*FileMarketOrderFulfilledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "OrderFulfilled", orderIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketOrderFulfilledIterator{contract: _FileMarket.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0xe5f48e2a4a303d196ad814389aa54cf34e7b1e404e740599d005907a2ca0cd6c.
//
// Solidity: event OrderFulfilled(uint256 indexed orderId, address indexed node)
func (_FileMarket *FileMarketFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *FileMarketOrderFulfilled, orderId []*big.Int, node []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "OrderFulfilled", orderIdRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketOrderFulfilled)
				if err := _FileMarket.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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

// ParseOrderFulfilled is a log parse operation binding the contract event 0xe5f48e2a4a303d196ad814389aa54cf34e7b1e404e740599d005907a2ca0cd6c.
//
// Solidity: event OrderFulfilled(uint256 indexed orderId, address indexed node)
func (_FileMarket *FileMarketFilterer) ParseOrderFulfilled(log types.Log) (*FileMarketOrderFulfilled, error) {
	event := new(FileMarketOrderFulfilled)
	if err := _FileMarket.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketOrderPlacedIterator is returned from FilterOrderPlaced and is used to iterate over the raw logs and unpacked data for OrderPlaced events raised by the FileMarket contract.
type FileMarketOrderPlacedIterator struct {
	Event *FileMarketOrderPlaced // Event containing the contract specifics and raw log

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
func (it *FileMarketOrderPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketOrderPlaced)
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
		it.Event = new(FileMarketOrderPlaced)
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
func (it *FileMarketOrderPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketOrderPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketOrderPlaced represents a OrderPlaced event raised by the FileMarket contract.
type FileMarketOrderPlaced struct {
	OrderId   *big.Int
	Owner     common.Address
	NumChunks uint32
	Periods   uint16
	Replicas  uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderPlaced is a free log retrieval operation binding the contract event 0x95b2702f91601cb754f51a9ccf9452e095fffb2bcf738f5fc59f170d86f3f678.
//
// Solidity: event OrderPlaced(uint256 indexed orderId, address indexed owner, uint32 numChunks, uint16 periods, uint8 replicas)
func (_FileMarket *FileMarketFilterer) FilterOrderPlaced(opts *bind.FilterOpts, orderId []*big.Int, owner []common.Address) (*FileMarketOrderPlacedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "OrderPlaced", orderIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketOrderPlacedIterator{contract: _FileMarket.contract, event: "OrderPlaced", logs: logs, sub: sub}, nil
}

// WatchOrderPlaced is a free log subscription operation binding the contract event 0x95b2702f91601cb754f51a9ccf9452e095fffb2bcf738f5fc59f170d86f3f678.
//
// Solidity: event OrderPlaced(uint256 indexed orderId, address indexed owner, uint32 numChunks, uint16 periods, uint8 replicas)
func (_FileMarket *FileMarketFilterer) WatchOrderPlaced(opts *bind.WatchOpts, sink chan<- *FileMarketOrderPlaced, orderId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "OrderPlaced", orderIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketOrderPlaced)
				if err := _FileMarket.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
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

// ParseOrderPlaced is a log parse operation binding the contract event 0x95b2702f91601cb754f51a9ccf9452e095fffb2bcf738f5fc59f170d86f3f678.
//
// Solidity: event OrderPlaced(uint256 indexed orderId, address indexed owner, uint32 numChunks, uint16 periods, uint8 replicas)
func (_FileMarket *FileMarketFilterer) ParseOrderPlaced(log types.Log) (*FileMarketOrderPlaced, error) {
	event := new(FileMarketOrderPlaced)
	if err := _FileMarket.contract.UnpackLog(event, "OrderPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketOrderUnderReplicatedIterator is returned from FilterOrderUnderReplicated and is used to iterate over the raw logs and unpacked data for OrderUnderReplicated events raised by the FileMarket contract.
type FileMarketOrderUnderReplicatedIterator struct {
	Event *FileMarketOrderUnderReplicated // Event containing the contract specifics and raw log

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
func (it *FileMarketOrderUnderReplicatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketOrderUnderReplicated)
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
		it.Event = new(FileMarketOrderUnderReplicated)
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
func (it *FileMarketOrderUnderReplicatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketOrderUnderReplicatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketOrderUnderReplicated represents a OrderUnderReplicated event raised by the FileMarket contract.
type FileMarketOrderUnderReplicated struct {
	OrderId         *big.Int
	CurrentFilled   uint8
	DesiredReplicas uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderUnderReplicated is a free log retrieval operation binding the contract event 0x569170f803879a4e984d6eb7e9c2e66fd69956c8722f9e6a6777b0f950cfb3d1.
//
// Solidity: event OrderUnderReplicated(uint256 indexed orderId, uint8 currentFilled, uint8 desiredReplicas)
func (_FileMarket *FileMarketFilterer) FilterOrderUnderReplicated(opts *bind.FilterOpts, orderId []*big.Int) (*FileMarketOrderUnderReplicatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "OrderUnderReplicated", orderIdRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketOrderUnderReplicatedIterator{contract: _FileMarket.contract, event: "OrderUnderReplicated", logs: logs, sub: sub}, nil
}

// WatchOrderUnderReplicated is a free log subscription operation binding the contract event 0x569170f803879a4e984d6eb7e9c2e66fd69956c8722f9e6a6777b0f950cfb3d1.
//
// Solidity: event OrderUnderReplicated(uint256 indexed orderId, uint8 currentFilled, uint8 desiredReplicas)
func (_FileMarket *FileMarketFilterer) WatchOrderUnderReplicated(opts *bind.WatchOpts, sink chan<- *FileMarketOrderUnderReplicated, orderId []*big.Int) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "OrderUnderReplicated", orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketOrderUnderReplicated)
				if err := _FileMarket.contract.UnpackLog(event, "OrderUnderReplicated", log); err != nil {
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

// ParseOrderUnderReplicated is a log parse operation binding the contract event 0x569170f803879a4e984d6eb7e9c2e66fd69956c8722f9e6a6777b0f950cfb3d1.
//
// Solidity: event OrderUnderReplicated(uint256 indexed orderId, uint8 currentFilled, uint8 desiredReplicas)
func (_FileMarket *FileMarketFilterer) ParseOrderUnderReplicated(log types.Log) (*FileMarketOrderUnderReplicated, error) {
	event := new(FileMarketOrderUnderReplicated)
	if err := _FileMarket.contract.UnpackLog(event, "OrderUnderReplicated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FileMarket contract.
type FileMarketOwnershipTransferredIterator struct {
	Event *FileMarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FileMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketOwnershipTransferred)
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
		it.Event = new(FileMarketOwnershipTransferred)
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
func (it *FileMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketOwnershipTransferred represents a OwnershipTransferred event raised by the FileMarket contract.
type FileMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FileMarket *FileMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FileMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketOwnershipTransferredIterator{contract: _FileMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FileMarket *FileMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FileMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketOwnershipTransferred)
				if err := _FileMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FileMarket *FileMarketFilterer) ParseOwnershipTransferred(log types.Log) (*FileMarketOwnershipTransferred, error) {
	event := new(FileMarketOwnershipTransferred)
	if err := _FileMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketProofFailureSlashMultiplierUpdatedIterator is returned from FilterProofFailureSlashMultiplierUpdated and is used to iterate over the raw logs and unpacked data for ProofFailureSlashMultiplierUpdated events raised by the FileMarket contract.
type FileMarketProofFailureSlashMultiplierUpdatedIterator struct {
	Event *FileMarketProofFailureSlashMultiplierUpdated // Event containing the contract specifics and raw log

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
func (it *FileMarketProofFailureSlashMultiplierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketProofFailureSlashMultiplierUpdated)
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
		it.Event = new(FileMarketProofFailureSlashMultiplierUpdated)
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
func (it *FileMarketProofFailureSlashMultiplierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketProofFailureSlashMultiplierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketProofFailureSlashMultiplierUpdated represents a ProofFailureSlashMultiplierUpdated event raised by the FileMarket contract.
type FileMarketProofFailureSlashMultiplierUpdated struct {
	OldMultiplier *big.Int
	NewMultiplier *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProofFailureSlashMultiplierUpdated is a free log retrieval operation binding the contract event 0x700bd13236dc88e9575fd685afabfbd2527c01c5a7a7a133fdfa747cde90743c.
//
// Solidity: event ProofFailureSlashMultiplierUpdated(uint256 oldMultiplier, uint256 newMultiplier)
func (_FileMarket *FileMarketFilterer) FilterProofFailureSlashMultiplierUpdated(opts *bind.FilterOpts) (*FileMarketProofFailureSlashMultiplierUpdatedIterator, error) {

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "ProofFailureSlashMultiplierUpdated")
	if err != nil {
		return nil, err
	}
	return &FileMarketProofFailureSlashMultiplierUpdatedIterator{contract: _FileMarket.contract, event: "ProofFailureSlashMultiplierUpdated", logs: logs, sub: sub}, nil
}

// WatchProofFailureSlashMultiplierUpdated is a free log subscription operation binding the contract event 0x700bd13236dc88e9575fd685afabfbd2527c01c5a7a7a133fdfa747cde90743c.
//
// Solidity: event ProofFailureSlashMultiplierUpdated(uint256 oldMultiplier, uint256 newMultiplier)
func (_FileMarket *FileMarketFilterer) WatchProofFailureSlashMultiplierUpdated(opts *bind.WatchOpts, sink chan<- *FileMarketProofFailureSlashMultiplierUpdated) (event.Subscription, error) {

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "ProofFailureSlashMultiplierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketProofFailureSlashMultiplierUpdated)
				if err := _FileMarket.contract.UnpackLog(event, "ProofFailureSlashMultiplierUpdated", log); err != nil {
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

// ParseProofFailureSlashMultiplierUpdated is a log parse operation binding the contract event 0x700bd13236dc88e9575fd685afabfbd2527c01c5a7a7a133fdfa747cde90743c.
//
// Solidity: event ProofFailureSlashMultiplierUpdated(uint256 oldMultiplier, uint256 newMultiplier)
func (_FileMarket *FileMarketFilterer) ParseProofFailureSlashMultiplierUpdated(log types.Log) (*FileMarketProofFailureSlashMultiplierUpdated, error) {
	event := new(FileMarketProofFailureSlashMultiplierUpdated)
	if err := _FileMarket.contract.UnpackLog(event, "ProofFailureSlashMultiplierUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketRefundQueuedIterator is returned from FilterRefundQueued and is used to iterate over the raw logs and unpacked data for RefundQueued events raised by the FileMarket contract.
type FileMarketRefundQueuedIterator struct {
	Event *FileMarketRefundQueued // Event containing the contract specifics and raw log

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
func (it *FileMarketRefundQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketRefundQueued)
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
		it.Event = new(FileMarketRefundQueued)
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
func (it *FileMarketRefundQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketRefundQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketRefundQueued represents a RefundQueued event raised by the FileMarket contract.
type FileMarketRefundQueued struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundQueued is a free log retrieval operation binding the contract event 0x03346f418777de37a432a76ace912eeb3d1d63333a4f8c59df9ae07eb0973215.
//
// Solidity: event RefundQueued(address indexed recipient, uint256 amount)
func (_FileMarket *FileMarketFilterer) FilterRefundQueued(opts *bind.FilterOpts, recipient []common.Address) (*FileMarketRefundQueuedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "RefundQueued", recipientRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketRefundQueuedIterator{contract: _FileMarket.contract, event: "RefundQueued", logs: logs, sub: sub}, nil
}

// WatchRefundQueued is a free log subscription operation binding the contract event 0x03346f418777de37a432a76ace912eeb3d1d63333a4f8c59df9ae07eb0973215.
//
// Solidity: event RefundQueued(address indexed recipient, uint256 amount)
func (_FileMarket *FileMarketFilterer) WatchRefundQueued(opts *bind.WatchOpts, sink chan<- *FileMarketRefundQueued, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "RefundQueued", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketRefundQueued)
				if err := _FileMarket.contract.UnpackLog(event, "RefundQueued", log); err != nil {
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

// ParseRefundQueued is a log parse operation binding the contract event 0x03346f418777de37a432a76ace912eeb3d1d63333a4f8c59df9ae07eb0973215.
//
// Solidity: event RefundQueued(address indexed recipient, uint256 amount)
func (_FileMarket *FileMarketFilterer) ParseRefundQueued(log types.Log) (*FileMarketRefundQueued, error) {
	event := new(FileMarketRefundQueued)
	if err := _FileMarket.contract.UnpackLog(event, "RefundQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketRefundWithdrawnIterator is returned from FilterRefundWithdrawn and is used to iterate over the raw logs and unpacked data for RefundWithdrawn events raised by the FileMarket contract.
type FileMarketRefundWithdrawnIterator struct {
	Event *FileMarketRefundWithdrawn // Event containing the contract specifics and raw log

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
func (it *FileMarketRefundWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketRefundWithdrawn)
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
		it.Event = new(FileMarketRefundWithdrawn)
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
func (it *FileMarketRefundWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketRefundWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketRefundWithdrawn represents a RefundWithdrawn event raised by the FileMarket contract.
type FileMarketRefundWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefundWithdrawn is a free log retrieval operation binding the contract event 0x3d97f39b86d061200a7834082f5926e58ec10fd85a9d6930f497729d5e6cc35c.
//
// Solidity: event RefundWithdrawn(address indexed recipient, uint256 amount)
func (_FileMarket *FileMarketFilterer) FilterRefundWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*FileMarketRefundWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "RefundWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketRefundWithdrawnIterator{contract: _FileMarket.contract, event: "RefundWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRefundWithdrawn is a free log subscription operation binding the contract event 0x3d97f39b86d061200a7834082f5926e58ec10fd85a9d6930f497729d5e6cc35c.
//
// Solidity: event RefundWithdrawn(address indexed recipient, uint256 amount)
func (_FileMarket *FileMarketFilterer) WatchRefundWithdrawn(opts *bind.WatchOpts, sink chan<- *FileMarketRefundWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "RefundWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketRefundWithdrawn)
				if err := _FileMarket.contract.UnpackLog(event, "RefundWithdrawn", log); err != nil {
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

// ParseRefundWithdrawn is a log parse operation binding the contract event 0x3d97f39b86d061200a7834082f5926e58ec10fd85a9d6930f497729d5e6cc35c.
//
// Solidity: event RefundWithdrawn(address indexed recipient, uint256 amount)
func (_FileMarket *FileMarketFilterer) ParseRefundWithdrawn(log types.Log) (*FileMarketRefundWithdrawn, error) {
	event := new(FileMarketRefundWithdrawn)
	if err := _FileMarket.contract.UnpackLog(event, "RefundWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketReporterRewardAccruedIterator is returned from FilterReporterRewardAccrued and is used to iterate over the raw logs and unpacked data for ReporterRewardAccrued events raised by the FileMarket contract.
type FileMarketReporterRewardAccruedIterator struct {
	Event *FileMarketReporterRewardAccrued // Event containing the contract specifics and raw log

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
func (it *FileMarketReporterRewardAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketReporterRewardAccrued)
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
		it.Event = new(FileMarketReporterRewardAccrued)
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
func (it *FileMarketReporterRewardAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketReporterRewardAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketReporterRewardAccrued represents a ReporterRewardAccrued event raised by the FileMarket contract.
type FileMarketReporterRewardAccrued struct {
	Reporter      common.Address
	RewardAmount  *big.Int
	SlashedAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterReporterRewardAccrued is a free log retrieval operation binding the contract event 0x5e831ec38c5df7467ace5ea04e7d150c124e53f7caf551ae2a78d07e8f9c1d31.
//
// Solidity: event ReporterRewardAccrued(address indexed reporter, uint256 rewardAmount, uint256 slashedAmount)
func (_FileMarket *FileMarketFilterer) FilterReporterRewardAccrued(opts *bind.FilterOpts, reporter []common.Address) (*FileMarketReporterRewardAccruedIterator, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "ReporterRewardAccrued", reporterRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketReporterRewardAccruedIterator{contract: _FileMarket.contract, event: "ReporterRewardAccrued", logs: logs, sub: sub}, nil
}

// WatchReporterRewardAccrued is a free log subscription operation binding the contract event 0x5e831ec38c5df7467ace5ea04e7d150c124e53f7caf551ae2a78d07e8f9c1d31.
//
// Solidity: event ReporterRewardAccrued(address indexed reporter, uint256 rewardAmount, uint256 slashedAmount)
func (_FileMarket *FileMarketFilterer) WatchReporterRewardAccrued(opts *bind.WatchOpts, sink chan<- *FileMarketReporterRewardAccrued, reporter []common.Address) (event.Subscription, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "ReporterRewardAccrued", reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketReporterRewardAccrued)
				if err := _FileMarket.contract.UnpackLog(event, "ReporterRewardAccrued", log); err != nil {
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

// ParseReporterRewardAccrued is a log parse operation binding the contract event 0x5e831ec38c5df7467ace5ea04e7d150c124e53f7caf551ae2a78d07e8f9c1d31.
//
// Solidity: event ReporterRewardAccrued(address indexed reporter, uint256 rewardAmount, uint256 slashedAmount)
func (_FileMarket *FileMarketFilterer) ParseReporterRewardAccrued(log types.Log) (*FileMarketReporterRewardAccrued, error) {
	event := new(FileMarketReporterRewardAccrued)
	if err := _FileMarket.contract.UnpackLog(event, "ReporterRewardAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketReporterRewardBpsUpdatedIterator is returned from FilterReporterRewardBpsUpdated and is used to iterate over the raw logs and unpacked data for ReporterRewardBpsUpdated events raised by the FileMarket contract.
type FileMarketReporterRewardBpsUpdatedIterator struct {
	Event *FileMarketReporterRewardBpsUpdated // Event containing the contract specifics and raw log

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
func (it *FileMarketReporterRewardBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketReporterRewardBpsUpdated)
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
		it.Event = new(FileMarketReporterRewardBpsUpdated)
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
func (it *FileMarketReporterRewardBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketReporterRewardBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketReporterRewardBpsUpdated represents a ReporterRewardBpsUpdated event raised by the FileMarket contract.
type FileMarketReporterRewardBpsUpdated struct {
	OldBps *big.Int
	NewBps *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReporterRewardBpsUpdated is a free log retrieval operation binding the contract event 0xbd21ba39bf0fe019fb62f92f1751db64b170f99a0d70857e0b338da8d169cb7a.
//
// Solidity: event ReporterRewardBpsUpdated(uint256 oldBps, uint256 newBps)
func (_FileMarket *FileMarketFilterer) FilterReporterRewardBpsUpdated(opts *bind.FilterOpts) (*FileMarketReporterRewardBpsUpdatedIterator, error) {

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "ReporterRewardBpsUpdated")
	if err != nil {
		return nil, err
	}
	return &FileMarketReporterRewardBpsUpdatedIterator{contract: _FileMarket.contract, event: "ReporterRewardBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchReporterRewardBpsUpdated is a free log subscription operation binding the contract event 0xbd21ba39bf0fe019fb62f92f1751db64b170f99a0d70857e0b338da8d169cb7a.
//
// Solidity: event ReporterRewardBpsUpdated(uint256 oldBps, uint256 newBps)
func (_FileMarket *FileMarketFilterer) WatchReporterRewardBpsUpdated(opts *bind.WatchOpts, sink chan<- *FileMarketReporterRewardBpsUpdated) (event.Subscription, error) {

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "ReporterRewardBpsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketReporterRewardBpsUpdated)
				if err := _FileMarket.contract.UnpackLog(event, "ReporterRewardBpsUpdated", log); err != nil {
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

// ParseReporterRewardBpsUpdated is a log parse operation binding the contract event 0xbd21ba39bf0fe019fb62f92f1751db64b170f99a0d70857e0b338da8d169cb7a.
//
// Solidity: event ReporterRewardBpsUpdated(uint256 oldBps, uint256 newBps)
func (_FileMarket *FileMarketFilterer) ParseReporterRewardBpsUpdated(log types.Log) (*FileMarketReporterRewardBpsUpdated, error) {
	event := new(FileMarketReporterRewardBpsUpdated)
	if err := _FileMarket.contract.UnpackLog(event, "ReporterRewardBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketReporterRewardsClaimedIterator is returned from FilterReporterRewardsClaimed and is used to iterate over the raw logs and unpacked data for ReporterRewardsClaimed events raised by the FileMarket contract.
type FileMarketReporterRewardsClaimedIterator struct {
	Event *FileMarketReporterRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *FileMarketReporterRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketReporterRewardsClaimed)
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
		it.Event = new(FileMarketReporterRewardsClaimed)
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
func (it *FileMarketReporterRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketReporterRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketReporterRewardsClaimed represents a ReporterRewardsClaimed event raised by the FileMarket contract.
type FileMarketReporterRewardsClaimed struct {
	Reporter common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReporterRewardsClaimed is a free log retrieval operation binding the contract event 0x7516bba11e49aa0761597694cac32c946a64065c9578fc7df200fc4c12f7e0d2.
//
// Solidity: event ReporterRewardsClaimed(address indexed reporter, uint256 amount)
func (_FileMarket *FileMarketFilterer) FilterReporterRewardsClaimed(opts *bind.FilterOpts, reporter []common.Address) (*FileMarketReporterRewardsClaimedIterator, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "ReporterRewardsClaimed", reporterRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketReporterRewardsClaimedIterator{contract: _FileMarket.contract, event: "ReporterRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchReporterRewardsClaimed is a free log subscription operation binding the contract event 0x7516bba11e49aa0761597694cac32c946a64065c9578fc7df200fc4c12f7e0d2.
//
// Solidity: event ReporterRewardsClaimed(address indexed reporter, uint256 amount)
func (_FileMarket *FileMarketFilterer) WatchReporterRewardsClaimed(opts *bind.WatchOpts, sink chan<- *FileMarketReporterRewardsClaimed, reporter []common.Address) (event.Subscription, error) {

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "ReporterRewardsClaimed", reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketReporterRewardsClaimed)
				if err := _FileMarket.contract.UnpackLog(event, "ReporterRewardsClaimed", log); err != nil {
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

// ParseReporterRewardsClaimed is a log parse operation binding the contract event 0x7516bba11e49aa0761597694cac32c946a64065c9578fc7df200fc4c12f7e0d2.
//
// Solidity: event ReporterRewardsClaimed(address indexed reporter, uint256 amount)
func (_FileMarket *FileMarketFilterer) ParseReporterRewardsClaimed(log types.Log) (*FileMarketReporterRewardsClaimed, error) {
	event := new(FileMarketReporterRewardsClaimed)
	if err := _FileMarket.contract.UnpackLog(event, "ReporterRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketRewardsCalculatedIterator is returned from FilterRewardsCalculated and is used to iterate over the raw logs and unpacked data for RewardsCalculated events raised by the FileMarket contract.
type FileMarketRewardsCalculatedIterator struct {
	Event *FileMarketRewardsCalculated // Event containing the contract specifics and raw log

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
func (it *FileMarketRewardsCalculatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketRewardsCalculated)
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
		it.Event = new(FileMarketRewardsCalculated)
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
func (it *FileMarketRewardsCalculatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketRewardsCalculatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketRewardsCalculated represents a RewardsCalculated event raised by the FileMarket contract.
type FileMarketRewardsCalculated struct {
	Node    common.Address
	Amount  *big.Int
	Periods *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewardsCalculated is a free log retrieval operation binding the contract event 0xbb3f529bf65429c4e1220ae6b7c10c40a76b917975cb0cdf6396e91b515bcb1b.
//
// Solidity: event RewardsCalculated(address indexed node, uint256 amount, uint256 periods)
func (_FileMarket *FileMarketFilterer) FilterRewardsCalculated(opts *bind.FilterOpts, node []common.Address) (*FileMarketRewardsCalculatedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "RewardsCalculated", nodeRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketRewardsCalculatedIterator{contract: _FileMarket.contract, event: "RewardsCalculated", logs: logs, sub: sub}, nil
}

// WatchRewardsCalculated is a free log subscription operation binding the contract event 0xbb3f529bf65429c4e1220ae6b7c10c40a76b917975cb0cdf6396e91b515bcb1b.
//
// Solidity: event RewardsCalculated(address indexed node, uint256 amount, uint256 periods)
func (_FileMarket *FileMarketFilterer) WatchRewardsCalculated(opts *bind.WatchOpts, sink chan<- *FileMarketRewardsCalculated, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "RewardsCalculated", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketRewardsCalculated)
				if err := _FileMarket.contract.UnpackLog(event, "RewardsCalculated", log); err != nil {
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

// ParseRewardsCalculated is a log parse operation binding the contract event 0xbb3f529bf65429c4e1220ae6b7c10c40a76b917975cb0cdf6396e91b515bcb1b.
//
// Solidity: event RewardsCalculated(address indexed node, uint256 amount, uint256 periods)
func (_FileMarket *FileMarketFilterer) ParseRewardsCalculated(log types.Log) (*FileMarketRewardsCalculated, error) {
	event := new(FileMarketRewardsCalculated)
	if err := _FileMarket.contract.UnpackLog(event, "RewardsCalculated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the FileMarket contract.
type FileMarketRewardsClaimedIterator struct {
	Event *FileMarketRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *FileMarketRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketRewardsClaimed)
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
		it.Event = new(FileMarketRewardsClaimed)
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
func (it *FileMarketRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketRewardsClaimed represents a RewardsClaimed event raised by the FileMarket contract.
type FileMarketRewardsClaimed struct {
	Node   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed node, uint256 amount)
func (_FileMarket *FileMarketFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, node []common.Address) (*FileMarketRewardsClaimedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "RewardsClaimed", nodeRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketRewardsClaimedIterator{contract: _FileMarket.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed node, uint256 amount)
func (_FileMarket *FileMarketFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *FileMarketRewardsClaimed, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "RewardsClaimed", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketRewardsClaimed)
				if err := _FileMarket.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed node, uint256 amount)
func (_FileMarket *FileMarketFilterer) ParseRewardsClaimed(log types.Log) (*FileMarketRewardsClaimed, error) {
	event := new(FileMarketRewardsClaimed)
	if err := _FileMarket.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketSlashAuthorityUpdatedIterator is returned from FilterSlashAuthorityUpdated and is used to iterate over the raw logs and unpacked data for SlashAuthorityUpdated events raised by the FileMarket contract.
type FileMarketSlashAuthorityUpdatedIterator struct {
	Event *FileMarketSlashAuthorityUpdated // Event containing the contract specifics and raw log

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
func (it *FileMarketSlashAuthorityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketSlashAuthorityUpdated)
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
		it.Event = new(FileMarketSlashAuthorityUpdated)
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
func (it *FileMarketSlashAuthorityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketSlashAuthorityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketSlashAuthorityUpdated represents a SlashAuthorityUpdated event raised by the FileMarket contract.
type FileMarketSlashAuthorityUpdated struct {
	Authority common.Address
	Allowed   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashAuthorityUpdated is a free log retrieval operation binding the contract event 0xebbbe5164023c08e353c5df99a80431bc5ab20f72444bb10eb805e35971f402c.
//
// Solidity: event SlashAuthorityUpdated(address indexed authority, bool allowed)
func (_FileMarket *FileMarketFilterer) FilterSlashAuthorityUpdated(opts *bind.FilterOpts, authority []common.Address) (*FileMarketSlashAuthorityUpdatedIterator, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "SlashAuthorityUpdated", authorityRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketSlashAuthorityUpdatedIterator{contract: _FileMarket.contract, event: "SlashAuthorityUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashAuthorityUpdated is a free log subscription operation binding the contract event 0xebbbe5164023c08e353c5df99a80431bc5ab20f72444bb10eb805e35971f402c.
//
// Solidity: event SlashAuthorityUpdated(address indexed authority, bool allowed)
func (_FileMarket *FileMarketFilterer) WatchSlashAuthorityUpdated(opts *bind.WatchOpts, sink chan<- *FileMarketSlashAuthorityUpdated, authority []common.Address) (event.Subscription, error) {

	var authorityRule []interface{}
	for _, authorityItem := range authority {
		authorityRule = append(authorityRule, authorityItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "SlashAuthorityUpdated", authorityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketSlashAuthorityUpdated)
				if err := _FileMarket.contract.UnpackLog(event, "SlashAuthorityUpdated", log); err != nil {
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

// ParseSlashAuthorityUpdated is a log parse operation binding the contract event 0xebbbe5164023c08e353c5df99a80431bc5ab20f72444bb10eb805e35971f402c.
//
// Solidity: event SlashAuthorityUpdated(address indexed authority, bool allowed)
func (_FileMarket *FileMarketFilterer) ParseSlashAuthorityUpdated(log types.Log) (*FileMarketSlashAuthorityUpdated, error) {
	event := new(FileMarketSlashAuthorityUpdated)
	if err := _FileMarket.contract.UnpackLog(event, "SlashAuthorityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketSlotChallengeIssuedIterator is returned from FilterSlotChallengeIssued and is used to iterate over the raw logs and unpacked data for SlotChallengeIssued events raised by the FileMarket contract.
type FileMarketSlotChallengeIssuedIterator struct {
	Event *FileMarketSlotChallengeIssued // Event containing the contract specifics and raw log

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
func (it *FileMarketSlotChallengeIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketSlotChallengeIssued)
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
		it.Event = new(FileMarketSlotChallengeIssued)
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
func (it *FileMarketSlotChallengeIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketSlotChallengeIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketSlotChallengeIssued represents a SlotChallengeIssued event raised by the FileMarket contract.
type FileMarketSlotChallengeIssued struct {
	SlotIndex      *big.Int
	OrderId        *big.Int
	ChallengedNode common.Address
	DeadlineBlock  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSlotChallengeIssued is a free log retrieval operation binding the contract event 0x7f9f229e5aaf8ec305fad64c32b3013867b7adbd12394356515ab61a723482f0.
//
// Solidity: event SlotChallengeIssued(uint256 indexed slotIndex, uint256 orderId, address challengedNode, uint256 deadlineBlock)
func (_FileMarket *FileMarketFilterer) FilterSlotChallengeIssued(opts *bind.FilterOpts, slotIndex []*big.Int) (*FileMarketSlotChallengeIssuedIterator, error) {

	var slotIndexRule []interface{}
	for _, slotIndexItem := range slotIndex {
		slotIndexRule = append(slotIndexRule, slotIndexItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "SlotChallengeIssued", slotIndexRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketSlotChallengeIssuedIterator{contract: _FileMarket.contract, event: "SlotChallengeIssued", logs: logs, sub: sub}, nil
}

// WatchSlotChallengeIssued is a free log subscription operation binding the contract event 0x7f9f229e5aaf8ec305fad64c32b3013867b7adbd12394356515ab61a723482f0.
//
// Solidity: event SlotChallengeIssued(uint256 indexed slotIndex, uint256 orderId, address challengedNode, uint256 deadlineBlock)
func (_FileMarket *FileMarketFilterer) WatchSlotChallengeIssued(opts *bind.WatchOpts, sink chan<- *FileMarketSlotChallengeIssued, slotIndex []*big.Int) (event.Subscription, error) {

	var slotIndexRule []interface{}
	for _, slotIndexItem := range slotIndex {
		slotIndexRule = append(slotIndexRule, slotIndexItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "SlotChallengeIssued", slotIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketSlotChallengeIssued)
				if err := _FileMarket.contract.UnpackLog(event, "SlotChallengeIssued", log); err != nil {
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

// ParseSlotChallengeIssued is a log parse operation binding the contract event 0x7f9f229e5aaf8ec305fad64c32b3013867b7adbd12394356515ab61a723482f0.
//
// Solidity: event SlotChallengeIssued(uint256 indexed slotIndex, uint256 orderId, address challengedNode, uint256 deadlineBlock)
func (_FileMarket *FileMarketFilterer) ParseSlotChallengeIssued(log types.Log) (*FileMarketSlotChallengeIssued, error) {
	event := new(FileMarketSlotChallengeIssued)
	if err := _FileMarket.contract.UnpackLog(event, "SlotChallengeIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketSlotDeactivatedIterator is returned from FilterSlotDeactivated and is used to iterate over the raw logs and unpacked data for SlotDeactivated events raised by the FileMarket contract.
type FileMarketSlotDeactivatedIterator struct {
	Event *FileMarketSlotDeactivated // Event containing the contract specifics and raw log

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
func (it *FileMarketSlotDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketSlotDeactivated)
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
		it.Event = new(FileMarketSlotDeactivated)
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
func (it *FileMarketSlotDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketSlotDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketSlotDeactivated represents a SlotDeactivated event raised by the FileMarket contract.
type FileMarketSlotDeactivated struct {
	SlotIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlotDeactivated is a free log retrieval operation binding the contract event 0x6b62697cd5a501bb3486f9301575ea9f7c0a176bd7cb1ee6ddefef8d355d6569.
//
// Solidity: event SlotDeactivated(uint256 indexed slotIndex)
func (_FileMarket *FileMarketFilterer) FilterSlotDeactivated(opts *bind.FilterOpts, slotIndex []*big.Int) (*FileMarketSlotDeactivatedIterator, error) {

	var slotIndexRule []interface{}
	for _, slotIndexItem := range slotIndex {
		slotIndexRule = append(slotIndexRule, slotIndexItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "SlotDeactivated", slotIndexRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketSlotDeactivatedIterator{contract: _FileMarket.contract, event: "SlotDeactivated", logs: logs, sub: sub}, nil
}

// WatchSlotDeactivated is a free log subscription operation binding the contract event 0x6b62697cd5a501bb3486f9301575ea9f7c0a176bd7cb1ee6ddefef8d355d6569.
//
// Solidity: event SlotDeactivated(uint256 indexed slotIndex)
func (_FileMarket *FileMarketFilterer) WatchSlotDeactivated(opts *bind.WatchOpts, sink chan<- *FileMarketSlotDeactivated, slotIndex []*big.Int) (event.Subscription, error) {

	var slotIndexRule []interface{}
	for _, slotIndexItem := range slotIndex {
		slotIndexRule = append(slotIndexRule, slotIndexItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "SlotDeactivated", slotIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketSlotDeactivated)
				if err := _FileMarket.contract.UnpackLog(event, "SlotDeactivated", log); err != nil {
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

// ParseSlotDeactivated is a log parse operation binding the contract event 0x6b62697cd5a501bb3486f9301575ea9f7c0a176bd7cb1ee6ddefef8d355d6569.
//
// Solidity: event SlotDeactivated(uint256 indexed slotIndex)
func (_FileMarket *FileMarketFilterer) ParseSlotDeactivated(log types.Log) (*FileMarketSlotDeactivated, error) {
	event := new(FileMarketSlotDeactivated)
	if err := _FileMarket.contract.UnpackLog(event, "SlotDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketSlotExpiredIterator is returned from FilterSlotExpired and is used to iterate over the raw logs and unpacked data for SlotExpired events raised by the FileMarket contract.
type FileMarketSlotExpiredIterator struct {
	Event *FileMarketSlotExpired // Event containing the contract specifics and raw log

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
func (it *FileMarketSlotExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketSlotExpired)
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
		it.Event = new(FileMarketSlotExpired)
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
func (it *FileMarketSlotExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketSlotExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketSlotExpired represents a SlotExpired event raised by the FileMarket contract.
type FileMarketSlotExpired struct {
	SlotIndex   *big.Int
	FailedNode  common.Address
	SlashAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSlotExpired is a free log retrieval operation binding the contract event 0x2914300d42c351bbf44f48d58b7e670ce3e053bab12e182a7940509f448f5c28.
//
// Solidity: event SlotExpired(uint256 indexed slotIndex, address indexed failedNode, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) FilterSlotExpired(opts *bind.FilterOpts, slotIndex []*big.Int, failedNode []common.Address) (*FileMarketSlotExpiredIterator, error) {

	var slotIndexRule []interface{}
	for _, slotIndexItem := range slotIndex {
		slotIndexRule = append(slotIndexRule, slotIndexItem)
	}
	var failedNodeRule []interface{}
	for _, failedNodeItem := range failedNode {
		failedNodeRule = append(failedNodeRule, failedNodeItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "SlotExpired", slotIndexRule, failedNodeRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketSlotExpiredIterator{contract: _FileMarket.contract, event: "SlotExpired", logs: logs, sub: sub}, nil
}

// WatchSlotExpired is a free log subscription operation binding the contract event 0x2914300d42c351bbf44f48d58b7e670ce3e053bab12e182a7940509f448f5c28.
//
// Solidity: event SlotExpired(uint256 indexed slotIndex, address indexed failedNode, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) WatchSlotExpired(opts *bind.WatchOpts, sink chan<- *FileMarketSlotExpired, slotIndex []*big.Int, failedNode []common.Address) (event.Subscription, error) {

	var slotIndexRule []interface{}
	for _, slotIndexItem := range slotIndex {
		slotIndexRule = append(slotIndexRule, slotIndexItem)
	}
	var failedNodeRule []interface{}
	for _, failedNodeItem := range failedNode {
		failedNodeRule = append(failedNodeRule, failedNodeItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "SlotExpired", slotIndexRule, failedNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketSlotExpired)
				if err := _FileMarket.contract.UnpackLog(event, "SlotExpired", log); err != nil {
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

// ParseSlotExpired is a log parse operation binding the contract event 0x2914300d42c351bbf44f48d58b7e670ce3e053bab12e182a7940509f448f5c28.
//
// Solidity: event SlotExpired(uint256 indexed slotIndex, address indexed failedNode, uint256 slashAmount)
func (_FileMarket *FileMarketFilterer) ParseSlotExpired(log types.Log) (*FileMarketSlotExpired, error) {
	event := new(FileMarketSlotExpired)
	if err := _FileMarket.contract.UnpackLog(event, "SlotExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketSlotProofSubmittedIterator is returned from FilterSlotProofSubmitted and is used to iterate over the raw logs and unpacked data for SlotProofSubmitted events raised by the FileMarket contract.
type FileMarketSlotProofSubmittedIterator struct {
	Event *FileMarketSlotProofSubmitted // Event containing the contract specifics and raw log

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
func (it *FileMarketSlotProofSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketSlotProofSubmitted)
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
		it.Event = new(FileMarketSlotProofSubmitted)
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
func (it *FileMarketSlotProofSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketSlotProofSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketSlotProofSubmitted represents a SlotProofSubmitted event raised by the FileMarket contract.
type FileMarketSlotProofSubmitted struct {
	SlotIndex  *big.Int
	Prover     common.Address
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSlotProofSubmitted is a free log retrieval operation binding the contract event 0x2bd855824677a92d22bd2899c194d2fce6f3ee66645cfc378c349d4148c5367e.
//
// Solidity: event SlotProofSubmitted(uint256 indexed slotIndex, address indexed prover, bytes32 commitment)
func (_FileMarket *FileMarketFilterer) FilterSlotProofSubmitted(opts *bind.FilterOpts, slotIndex []*big.Int, prover []common.Address) (*FileMarketSlotProofSubmittedIterator, error) {

	var slotIndexRule []interface{}
	for _, slotIndexItem := range slotIndex {
		slotIndexRule = append(slotIndexRule, slotIndexItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "SlotProofSubmitted", slotIndexRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketSlotProofSubmittedIterator{contract: _FileMarket.contract, event: "SlotProofSubmitted", logs: logs, sub: sub}, nil
}

// WatchSlotProofSubmitted is a free log subscription operation binding the contract event 0x2bd855824677a92d22bd2899c194d2fce6f3ee66645cfc378c349d4148c5367e.
//
// Solidity: event SlotProofSubmitted(uint256 indexed slotIndex, address indexed prover, bytes32 commitment)
func (_FileMarket *FileMarketFilterer) WatchSlotProofSubmitted(opts *bind.WatchOpts, sink chan<- *FileMarketSlotProofSubmitted, slotIndex []*big.Int, prover []common.Address) (event.Subscription, error) {

	var slotIndexRule []interface{}
	for _, slotIndexItem := range slotIndex {
		slotIndexRule = append(slotIndexRule, slotIndexItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "SlotProofSubmitted", slotIndexRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketSlotProofSubmitted)
				if err := _FileMarket.contract.UnpackLog(event, "SlotProofSubmitted", log); err != nil {
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

// ParseSlotProofSubmitted is a log parse operation binding the contract event 0x2bd855824677a92d22bd2899c194d2fce6f3ee66645cfc378c349d4148c5367e.
//
// Solidity: event SlotProofSubmitted(uint256 indexed slotIndex, address indexed prover, bytes32 commitment)
func (_FileMarket *FileMarketFilterer) ParseSlotProofSubmitted(log types.Log) (*FileMarketSlotProofSubmitted, error) {
	event := new(FileMarketSlotProofSubmitted)
	if err := _FileMarket.contract.UnpackLog(event, "SlotProofSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketSlotsActivatedIterator is returned from FilterSlotsActivated and is used to iterate over the raw logs and unpacked data for SlotsActivated events raised by the FileMarket contract.
type FileMarketSlotsActivatedIterator struct {
	Event *FileMarketSlotsActivated // Event containing the contract specifics and raw log

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
func (it *FileMarketSlotsActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketSlotsActivated)
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
		it.Event = new(FileMarketSlotsActivated)
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
func (it *FileMarketSlotsActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketSlotsActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketSlotsActivated represents a SlotsActivated event raised by the FileMarket contract.
type FileMarketSlotsActivated struct {
	ActivatedCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSlotsActivated is a free log retrieval operation binding the contract event 0xc36d933a8e45c2ff7f0d8359aa2de2770d3c8f9bd96f9a9f28db85d4717ac33c.
//
// Solidity: event SlotsActivated(uint256 activatedCount)
func (_FileMarket *FileMarketFilterer) FilterSlotsActivated(opts *bind.FilterOpts) (*FileMarketSlotsActivatedIterator, error) {

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "SlotsActivated")
	if err != nil {
		return nil, err
	}
	return &FileMarketSlotsActivatedIterator{contract: _FileMarket.contract, event: "SlotsActivated", logs: logs, sub: sub}, nil
}

// WatchSlotsActivated is a free log subscription operation binding the contract event 0xc36d933a8e45c2ff7f0d8359aa2de2770d3c8f9bd96f9a9f28db85d4717ac33c.
//
// Solidity: event SlotsActivated(uint256 activatedCount)
func (_FileMarket *FileMarketFilterer) WatchSlotsActivated(opts *bind.WatchOpts, sink chan<- *FileMarketSlotsActivated) (event.Subscription, error) {

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "SlotsActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketSlotsActivated)
				if err := _FileMarket.contract.UnpackLog(event, "SlotsActivated", log); err != nil {
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

// ParseSlotsActivated is a log parse operation binding the contract event 0xc36d933a8e45c2ff7f0d8359aa2de2770d3c8f9bd96f9a9f28db85d4717ac33c.
//
// Solidity: event SlotsActivated(uint256 activatedCount)
func (_FileMarket *FileMarketFilterer) ParseSlotsActivated(log types.Log) (*FileMarketSlotsActivated, error) {
	event := new(FileMarketSlotsActivated)
	if err := _FileMarket.contract.UnpackLog(event, "SlotsActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileMarketUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the FileMarket contract.
type FileMarketUpgradedIterator struct {
	Event *FileMarketUpgraded // Event containing the contract specifics and raw log

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
func (it *FileMarketUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileMarketUpgraded)
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
		it.Event = new(FileMarketUpgraded)
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
func (it *FileMarketUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileMarketUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileMarketUpgraded represents a Upgraded event raised by the FileMarket contract.
type FileMarketUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FileMarket *FileMarketFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FileMarketUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FileMarket.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FileMarketUpgradedIterator{contract: _FileMarket.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FileMarket *FileMarketFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FileMarketUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _FileMarket.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileMarketUpgraded)
				if err := _FileMarket.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_FileMarket *FileMarketFilterer) ParseUpgraded(log types.Log) (*FileMarketUpgraded, error) {
	event := new(FileMarketUpgraded)
	if err := _FileMarket.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
