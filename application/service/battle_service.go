package service

import (
	"pokemon-battle/domain/model"
)

type BattleService struct {
	battle *model.BattleEngine
}

func NewBattleService(pokemon1, pokemon2 *model.Pokemon) (*BattleService, error) {
	battle, err := model.NewBattleEngine(pokemon1, pokemon2)
	if err != nil {
		return nil, err
	}

	return &BattleService{
		battle: battle,
	}, nil
}

func (s *BattleService) ExecuteTurn(move1, move2 model.Move) error {
	return s.battle.BattleTurn(move1, move2)
} 