package service

import (
	"context"
	"errors"

	"pokemon-battle/pkg/battle/domain"
	"pokemon-battle/pkg/battle/ports"
)

// BattleService はバトルサービスの実装です
type BattleService struct {
	battles  map[string]*domain.Battle
	moves    map[string]*domain.Move
	pokemons map[string]*domain.Pokemon
}

// NewBattleService は新しいバトルサービスを作成します
func NewBattleService() *BattleService {
	return &BattleService{
		battles:  make(map[string]*domain.Battle),
		moves:    make(map[string]*domain.Move),
		pokemons: make(map[string]*domain.Pokemon),
	}
}

// CreatePokemon はポケモンを作成します
func (s *BattleService) CreatePokemon(ctx context.Context, id string, maxHP int, speed int) error {
	if _, exists := s.pokemons[id]; exists {
		return errors.New("pokemon already exists")
	}

	pokemon := &domain.Pokemon{
		ID:        id,
		CurrentHP: maxHP,
		MaxHP:     maxHP,
		Speed:     speed,
		Status:    "NORMAL",
		Moves:     []domain.Move{},
	}

	s.pokemons[id] = pokemon
	return nil
}

// CreateMove は技を作成します
func (s *BattleService) CreateMove(ctx context.Context, id string, power int, accuracy int, priority int) error {
	if _, exists := s.moves[id]; exists {
		return errors.New("move already exists")
	}

	move := &domain.Move{
		ID:       id,
		Power:    power,
		Accuracy: accuracy,
		Priority: priority,
	}

	s.moves[id] = move
	return nil
}

// CreateBattle はバトルを作成します
func (s *BattleService) CreateBattle(ctx context.Context, id string, pokemon1ID, pokemon2ID string) error {
	pokemon1, exists := s.pokemons[pokemon1ID]
	if !exists {
		return errors.New("pokemon1 not found")
	}

	pokemon2, exists := s.pokemons[pokemon2ID]
	if !exists {
		return errors.New("pokemon2 not found")
	}

	battle, err := domain.NewBattle(id, pokemon1, pokemon2)
	if err != nil {
		return err
	}

	s.battles[id] = battle
	return nil
}

// ExecuteTurn は1ターンのバトルを実行します
func (s *BattleService) ExecuteTurn(ctx context.Context, battleID string, move1ID string, move2ID string) error {
	battle, exists := s.battles[battleID]
	if !exists {
		return errors.New("battle not found")
	}

	// 技の取得
	move1, exists := s.moves[move1ID]
	if !exists {
		return errors.New("move1 not found")
	}
	move2, exists := s.moves[move2ID]
	if !exists {
		return errors.New("move2 not found")
	}

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

	// 技の取得
	move, exists := s.moves[moveID]
	if !exists {
		return errors.New("move not found")
	}

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
