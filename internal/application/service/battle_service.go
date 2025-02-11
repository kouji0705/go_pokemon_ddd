package service

import "pokemon-battle/internal/domain/model"

type BattleService struct {
	battle *model.BattleEngine
}

func NewBattleService(pokemon1, pokemon2 *model.Pokemon) (*BattleService, error) {
	battle := model.NewBattleEngine(pokemon1, pokemon2)
	return &BattleService{
		battle: battle,
	}, nil
}

func (s *BattleService) ExecuteTurn(move1, move2 *model.Move) error {
	s.battle.BattleTurn(move1, move2)
	return nil
}

// ExecuteMove は1回の攻撃を実行します
func (s *BattleService) ExecuteMove(attacker, defender *model.Pokemon, move *model.Move) error {
	s.battle.ExecuteMove(attacker, defender, move)
	return nil
}
