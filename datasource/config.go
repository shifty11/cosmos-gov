package datasource

import (
	"fmt"
	"github.com/shifty11/cosmos-gov/log"
	"github.com/strangelove-ventures/lens/cmd"
	"go.uber.org/zap"
	"strings"
)

func addOrUpdateChainInLensConfig(chainName string) {
	query := fmt.Sprintf("chains add %v", chainName)
	rootCmd := cmd.NewRootCmd(log.Sugar.Desugar(), zap.NewAtomicLevel(), nil)
	rootCmd.SetArgs(strings.Fields(query))
	log.Sugar.Debugf("Add/update chain %v in lens config: 'lens %v'", chainName, query)
	err := rootCmd.Execute()
	if err != nil {
		log.Sugar.Debugf("Error while executing query '%v': %v", query, err)
	}
}
