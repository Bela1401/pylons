package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/Pylons-tech/pylons/x/pylons/types"
)

func CmdCreateCookbook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-cookbook [id] [name] [description] [developer] [version] [supportEmail] [tier] [costPerBlock] [enabled]",
		Short: "Create a new cookbook",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsName, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsDescription, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsDeveloper, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsVersion, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}
			argsSupportEmail, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}
			argsTier, err := cast.ToInt64E(args[6])
			if err != nil {
				return err
			}
			argsCostPerBlock, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}
			argsEnabled, err := cast.ToBoolE(args[8])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCookbook(clientCtx.GetFromAddress().String(), id, argsName, argsDescription, argsDeveloper, argsVersion, argsSupportEmail, argsTier, argsCostPerBlock, argsEnabled)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateCookbook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-cookbook [id] [name] [description] [developer] [version] [supportEmail] [tier] [costPerBlock] [enabled]",
		Short: "Update a cookbook",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsName, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsDescription, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsDeveloper, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsVersion, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}
			argsSupportEmail, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}
			argsTier, err := cast.ToInt64E(args[6])
			if err != nil {
				return err
			}
			argsCostPerBlock, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}
			argsEnabled, err := cast.ToBoolE(args[8])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateCookbook(clientCtx.GetFromAddress().String(), id, argsName, argsDescription, argsDeveloper, argsVersion, argsSupportEmail, argsTier, argsCostPerBlock, argsEnabled)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
