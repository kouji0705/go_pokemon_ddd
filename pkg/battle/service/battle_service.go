package service

import (
	"context"
	"errors"

	"pokemon-battle/pkg/battle/domain"
	"pokemon-battle/pkg/battle/ports"
)

// BattleService はバトルサービスの実装です
type BattleService struct {
	battles map[string]*domain.Battle
}

// NewBattleService は新しいバトルサービスを作成します
func NewBattleService() *BattleService {
	return &BattleService{
		battles: make(map[string]*domain.Battle),
	}
}

// RegisterBattle はバトルを登録します
func (s *BattleService) RegisterBattle(battle *domain.Battle) error {
	if battle == nil {
		return errors.New("battle cannot be nil")
	}
	if _, exists := s.battles[battle.ID]; exists {
		return errors.New("battle already exists")
	}
	s.battles[battle.ID] = battle
	return nil
}

// ExecuteTurn は1ターンのバトルを実行します
func (s *BattleService) ExecuteTurn(ctx context.Context, battleID string, move1ID string, move2ID string) error {
	battle, exists := s.battles[battleID]
	if !exists {
		return errors.New("battle not found")
	}

	// 技の優先度に基づいて実行順序を決定
	move1 := &domain.Move{ID: move1ID} // 実際には技の詳細情報を取得する必要があります
	move2 := &domain.Move{ID: move2ID}

	// 技の実行
	if err := battle.ExecuteMove(battle.Pokemon1, battle.Pokemon2, move1); err != nil {
		return err
	}

	// 相手のポケモンがまだ戦闘可能な場合、2番目の技を実行
	if battle.Status != domain.BattleStatusFinished {
		if err := battle.ExecuteMove(battle.Pokemon2, battle.Pokemon1, move2); err != nil {
			return err
		}
	}

	return nil
}

// ExecuteMove は1回の攻撃を実行します
func (s *BattleService) ExecuteMove(ctx context.Context, battleID string, attackerID string, defenderID string, moveID string) error {
	battle, exists := s.battles[battleID]
	if !exists {
		return errors.New("battle not found")
	}

	// attackerとdefenderの特定
	var attacker, defender *domain.Pokemon
	if battle.Pokemon1.ID == attackerID {
		attacker = battle.Pokemon1
		defender = battle.Pokemon2
	} else if battle.Pokemon2.ID == attackerID {
		attacker = battle.Pokemon2
		defender = battle.Pokemon1
	} else {
		return errors.New("attacker not found in battle")
	}

	// 技の取得（実際には技のリポジトリなどから取得する必要があります）
	move := &domain.Move{ID: moveID}

	return battle.ExecuteMove(attacker, defender, move)
}

// GetBattleStatus はバトルの現在の状態を取得します
func (s *BattleService) GetBattleStatus(ctx context.Context, battleID string) (*ports.BattleStatus, error) {
	battle, exists := s.battles[battleID]
	if !exists {
		return nil, errors.New("battle not found")
	}

	// ドメインモデルからDTOへの変換
	status := &ports.BattleStatus{
		BattleID: battle.ID,
		Pokemon1: ports.PokemonStatus{
			PokemonID: battle.Pokemon1.ID,
			CurrentHP: battle.Pokemon1.CurrentHP,
			MaxHP:     battle.Pokemon1.MaxHP,
			Status:    battle.Pokemon1.Status,
		},
		Pokemon2: ports.PokemonStatus{
			PokemonID: battle.Pokemon2.ID,
			CurrentHP: battle.Pokemon2.CurrentHP,
			MaxHP:     battle.Pokemon2.MaxHP,
			Status:    battle.Pokemon2.Status,
		},
		IsFinished: battle.Status == domain.BattleStatusFinished,
	}

	// 勝者の決定
	if status.IsFinished {
		var winner string
		if battle.Pokemon1.CurrentHP > 0 {
			winner = battle.Pokemon1.ID
		} else {
			winner = battle.Pokemon2.ID
		}
		status.Winner = &winner
	}

	return status, nil
}
