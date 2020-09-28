package commands

import (
	"encoding/hex"
	"encoding/json"
	"github.com/olympus-protocol/ogen/cmd/ogen/initialization"
	"github.com/olympus-protocol/ogen/internal/keystore"
	"github.com/olympus-protocol/ogen/pkg/bls"
	"github.com/olympus-protocol/ogen/pkg/params"
	"github.com/spf13/cobra"
	"io/ioutil"
	"path"
)

var (
	amount  int
	network string
)

var genParamsCmd = &cobra.Command{
	Use:   "genparams",
	Short: "Used to generate parameters for network initialization",
	Long:  `Used to generate parameters for network initialization`,
	Run: func(cmd *cobra.Command, args []string) {

		var netParams *params.ChainParams
		switch network {
		case "mainnet":
			netParams = &params.Mainnet
		default:
			netParams = &params.TestNet
		}

		bls.Initialize(netParams)

		premine := bls.RandKey()

		dirPath := "./cmd/ogen/initialization/"

		ks := keystore.NewKeystore(dirPath, nil)

		err := ks.CreateKeystore()
		if err != nil {
			panic(err)
		}

		validatorsKeys, err := ks.GenerateNewValidatorKey(uint64(amount))
		if err != nil {
			panic(err)
		}

		validators := make([]initialization.Validators, amount)
		for i, key := range validatorsKeys {
			v := initialization.Validators{
				PublicKey: hex.EncodeToString(key.PublicKey().Marshal()),
			}
			if network != "mainnet" {
				v.PrivateKey = hex.EncodeToString(key.Marshal())
			}
			validators[i] = v
		}

		initParams := initialization.NetworkInitialParams{
			Validators:     validators,
			PremineAddress: premine.PublicKey().ToAccount(),
		}

		if network != "mainnet" {
			initParams.PreminePrivateKey = premine.ToWIF()
		}

		b, err := json.Marshal(initParams)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(path.Join(dirPath, network+"_params.json"), b, 0700)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	genParamsCmd.Flags().StringVar(&network, "network", "testnet", "network name to generate params to")
	genParamsCmd.Flags().IntVar(&amount, "amount", 8, "amount of validators to generate")
	genParamsCmd.Flags().IntVar(&amount, "genesistime", 0, "genesis time to start the chain")

	rootCmd.AddCommand(genParamsCmd)
}
