package datasource

import (
	"fmt"
	"github.com/shifty11/cosmos-gov/log"
	"github.com/strangelove-ventures/lens/cmd"
	"strings"
)

func addOrUpdateChainInLensConfig(chainName string) {
	query := fmt.Sprintf("chains add %v", chainName)
	rootCmd := cmd.NewRootCmd()
	rootCmd.SetArgs(strings.Fields(query))
	log.Sugar.Debugf("Add/update chain %v in lens config: 'lens %v'", chainName, query)
	err := rootCmd.Execute()
	if err != nil {
		log.Sugar.Errorf("Error while executing query '%v': %v", query, err)
	}
}

func removeChainFromLensConfig(chainName string) {
	query := fmt.Sprintf("chains delete %v", chainName)
	rootCmd := cmd.NewRootCmd()
	rootCmd.SetArgs(strings.Fields(query))
	log.Sugar.Debugf("Remove chain %v from lens config: 'lens %v'", chainName, query)
	err := rootCmd.Execute()
	if err != nil {
		log.Sugar.Errorf("Error while executing query '%v': %v", query, err)
	}
}