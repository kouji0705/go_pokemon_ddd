package model

import (
	"errors"
	"math"
)

type BattleEngine struct {
	pokemon1 *Pokemon
	pokemon2 *Pokemon
}

func NewBattleEngine(p1, p2 *Pokemon) (*BattleEngine, error) {
	if p1 == nil || p2 == nil {
		return nil, errors.New("ポケモンが指定されていません")
	}
	return &BattleEngine{
		pokemon1: p1,
		pokemon2: p2,
	}, nil
}

func (b *BattleEngine) ExecuteMove(attacker, defender *Pokemon, move Move) error {
	if attacker.IsFainted() {
		return errors.New("気絶したポケモンは技を使用できません")
	}

	damage := b.calculateDamage(attacker, defender, move)
	defender.TakeDamage(damage)

	return nil
}

func (b *BattleEngine) calculateDamage(attacker, defender *Pokemon, move Move) int {
	// 基本的なダメージ計算式
	baseDamage := (2*1+10)/250.0*float64(attacker.Attack())/float64(defender.Defense())*float64(move.Power()) + 2
	return int(math.Floor(baseDamage))
}

func (b *BattleEngine) BattleTurn(move1, move2 Move) error {
	// 素早さ判定による行動順序の決定
	first, second := b.pokemon1, b.pokemon2
	firstMove, secondMove := move1, move2

	if b.pokemon2.Speed() > b.pokemon1.Speed() {
		first, second = b.pokemon2, b.pokemon1
		firstMove, secondMove = move2, move1
	}

	// 技の実行
	if err := b.ExecuteMove(first, second, firstMove); err != nil {
		return err
	}

	if !second.IsFainted() {
		if err := b.ExecuteMove(second, first, secondMove); err != nil {
			return err
		}
	}

	return nil
}
