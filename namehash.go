package resolution

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func NameHash(name string) common.Hash {
	output := common.Hash{}

	if len(name) == 0 {
		return common.Hash{}
	}

	labels := strings.Split(name, ".")

	for _, label := range labels {
		labelSha := crypto.Keccak256Hash([]byte(label))
		output = crypto.Keccak256Hash(output.Bytes(), labelSha.Bytes())
	}

	return output
}
