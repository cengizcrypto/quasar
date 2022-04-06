package cli

import (
	"fmt"
	"strconv"

	"github.com/abag/quasarnode/x/intergamm/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channelutils "github.com/cosmos/ibc-go/v3/modules/core/04-channel/client/utils"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSendIbcJoinPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-ibc-join-pool [src-port] [src-channel]",
		Short: "Broadcast message sendIbcJoinPool over IBC",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]
			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			poolId, err := cmd.Flags().GetUint64(FlagPoolId)
			if err != nil {
				return err
			}

			shareAmountOutStr, err := cmd.Flags().GetString(FlagShareAmountOut)
			if err != nil {
				return err
			}

			shareAmountOut, ok := sdk.NewIntFromString(shareAmountOutStr)
			if !ok {
				return fmt.Errorf("invalid share amount out")
			}

			maxAmountsInStrs, err := cmd.Flags().GetStringArray(FlagMaxAmountsIn)
			if err != nil {
				return err
			}

			maxAmountsIn := sdk.Coins{}
			for i := 0; i < len(maxAmountsInStrs); i++ {
				parsed, err := sdk.ParseCoinNormalized(maxAmountsInStrs[i])
				if err != nil {
					return err
				}
				maxAmountsIn = append(maxAmountsIn, parsed)
			}

			msg := types.NewMsgSendIbcJoinPool(
				creator,
				srcPort,
				srcChannel,
				timeoutTimestamp,
				poolId,
				shareAmountOut,
				maxAmountsIn,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetJoinPool())
	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagPoolId)
	_ = cmd.MarkFlagRequired(FlagShareAmountOut)
	_ = cmd.MarkFlagRequired(FlagMaxAmountsIn)

	return cmd
}