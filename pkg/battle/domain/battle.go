package domain

import (
	"errors"
)

// Battle はバトルのドメインモデルを表現します
type Battle struct {
	ID       string
	Pokemon1 *Pokemon
	Pokemon2 *Pokemon
	Status   BattleStatus
}

// Pokemon はバトル中のポケモンを表現します
type Pokemon struct {
	ID        string
	CurrentHP int
	MaxHP     int
	Speed     int
	Status    string
	Moves     []Move
}

// Move は技を表現します
type Move struct {
	ID       string
	Power    int
	Accuracy int
	Priority int
}

// BattleStatus はバトルの状態を表現します
type BattleStatus string

const (
	BattleStatusOngoing  BattleStatus = "ONGOING"
	BattleStatusFinished BattleStatus = "FINISHED"
)

// NewBattle は新しいバトルを作成します
func NewBattle(id string, pokemon1, pokemon2 *Pokemon) (*Battle, error) {
	if pokemon1 == nil || pokemon2 == nil {
		return nil, errors.New("pokemon cannot be nil")
	}

	return &Battle{
		ID:       id,
		Pokemon1: pokemon1,
		Pokemon2: pokemon2,
		Status:   BattleStatusOngoing,
	}, nil
}

// ExecuteMove は1回の攻撃を実行します
func (b *Battle) ExecuteMove(attacker, defender *Pokemon, move *Move) error {
	if b.Status == BattleStatusFinished {
		return errors.New("battle is already finished")
	}

	// 技の実行ロジック（簡略化）
	damage := calculateDamage(move.Power)
	defender.CurrentHP -= damage

	// バトル終了判定
	if defender.CurrentHP <= 0 {
		defender.CurrentHP = 0
		b.Status = BattleStatusFinished
	}

	return nil
}

// calculateDamage はダメージ計算を行います（簡略化）
func calculateDamage(power int) int {
	return power // 実際にはもっと複雑な計算が必要
}
