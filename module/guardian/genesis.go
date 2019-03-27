package guardian

import (
	"github.com/QOSGroup/qbase/context"
	"github.com/QOSGroup/qos/module/guardian/types"
)

type GenesisState struct {
	Guardians []types.Guardian `json:"guardians"`
}

func NewGenesisState(guardians []types.Guardian) GenesisState {
	return GenesisState{
		Guardians: guardians,
	}
}

func InitGenesis(ctx context.Context, data GenesisState) {
	mapper := GetGuardianMapper(ctx)
	for _, guardian := range data.Guardians {
		mapper.AddGuardian(guardian)
	}
}

func ExportGenesis(ctx context.Context) GenesisState {
	mapper := GetGuardianMapper(ctx)
	iterator := mapper.GuardiansIterator()
	defer iterator.Close()
	var guardians []types.Guardian
	for ; iterator.Valid(); iterator.Next() {
		var guardian types.Guardian
		mapper.GetCodec().MustUnmarshalBinaryBare(iterator.Value(), &guardian)
		guardians = append(guardians, guardian)
	}

	return NewGenesisState(guardians)
}
