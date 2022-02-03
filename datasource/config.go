package datasource

import (
	"github.com/shifty11/cosmos-gov/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type lensConf struct {
	DefaultChain string `yaml:"default_chain"`
	Chains       map[string]struct {
		NotUsed map[string]struct{}
	}
}

func ReadLensConfig(filename string) []string {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Sugar.Panicf("in file %q: %v", filename, err)
	}

	config := &lensConf{}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		log.Sugar.Panicf("in file %q: %v", filename, err)
	}

	var chains []string
	for key := range config.Chains {
		chains = append(chains, key)
	}
	return chains
}
