
package cli

import (
	"fmt"
	_ "github.com/fatih/color"
	"github.com/filecoin-project/go-address"
	_ "github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/types"
	miner2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
	"os"
	"strconv"
	"text/tabwriter"
)

var WinPostSetCmd = &cli.Command{
	Name:  "proving",
	Usage: "WinPost check",
	Subcommands: []*cli.Command{
		provingCheckWinCmd,
	},
}



var provingCheckWinCmd = &cli.Command{
	Name:      "check",
	Usage:     "Check sectors provable",
	ArgsUsage: "<deadlineIdx>",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "only-bad",
			Usage: "print only bad sectors",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  "full",
			Usage: "run deadline full check",
		},
	},
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return xerrors.Errorf("must pass deadline index")
		}

		dlIdx, err := strconv.ParseUint(cctx.Args().Get(0), 10, 64)
		if err != nil {
			return xerrors.Errorf("could not parse deadline index: %w", err)
		}

		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		sapi, scloser, err := GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer scloser()

		ctx := ReqContext(cctx)

		addr, err := sapi.ActorAddress(ctx)
		if err != nil {
			return err
		}

		mid, err := address.IDFromAddress(addr)
		if err != nil {
			return err
		}
		fmt.Println("mid",mid)

		info, err := api.StateMinerInfo(ctx, addr, types.EmptyTSK)
		if err != nil {
			return err
		}
		fmt.Println("StateMinerInfo",info)

		partitions, err := api.StateMinerPartitions(ctx, addr, dlIdx, types.EmptyTSK)
		if err != nil {
			return err
		}
		fmt.Println("StateMinerPartitions",partitions)

		tw := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
		_, _ = fmt.Fprintln(tw, "deadline\tpartition\tsector\tstatus")

		///step1
		log.Warnf("prepare CHECK WINDOW POST -------------------- %v", dlIdx)
		proofs, err := CheckWindowPoSt(cctx, dlIdx)
		if err != nil {
				return err
		}
		fmt.Printf("Deadline: %v\n", proofs)
		return tw.Flush()
	},
}


func CheckWindowPoSt(ctx *cli.Context, deadline uint64) ([]miner2.SubmitWindowedPoStParams, error) {
	minerApi, closer, err := GetStorageMinerAPI(ctx)
	if err != nil {
		return nil, err
	}
	defer closer()

	return minerApi.CheckWindowPoSt(ctx.Context, deadline)
}
