package action

import (
	"flag"
)

type operatorRaftConfig struct {
	*config
}

func OperatorRaftConfigAction() Action {
	return &operatorRaftConfig{
		config: &gConfig,
	}
}

func (o *operatorRaftConfig) CommandFlags() *flag.FlagSet {
	return o.newFlagSet(FLAG_DATACENTER, FLAG_OUTPUT, FLAG_STALE)
}

func (o *operatorRaftConfig) Run(args []string) error {
	client, err := o.newOperator()
	if err != nil {
		return err
	}

	queryOpts := o.queryOptions()

	rc, err := client.RaftGetConfiguration(queryOpts)
	if err != nil {
		return err
	}

	return o.Output(rc)
}
