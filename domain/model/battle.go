// Package model はポケモンバトルのドメインモデルを提供します。
package model

import (
	"errors"
	"math"
)

// BattleEngine はポケモンバトルのコアロジックを実装する構造体です。
// 2体のポケモン間の対戦を管理し、ダメージ計算やターン進行を処理します。
type BattleEngine struct {
	pokemon1 *Pokemon // 1番目のポケモン
	pokemon2 *Pokemon // 2番目のポケモン
}

// NewBattleEngine は2体のポケモンによる対戦を管理するBattleEngineを生成します。
// 引数として渡された2体のポケモンのいずれかがnilの場合、エラーを返します。
//
// Parameters:
//   - p1: 1体目のポケモン。nilであってはいけません。
//   - p2: 2体目のポケモン。nilであってはいけません。
//
// Returns:
//   - *BattleEngine: 生成されたバトルエンジン
//   - error: ポケモンが指定されていない場合はエラーを返します
func NewBattleEngine(p1, p2 *Pokemon) (*BattleEngine, error) {
	if p1 == nil || p2 == nil {
		return nil, errors.New("ポケモンが指定されていません")
	}
	return &BattleEngine{
		pokemon1: p1,
		pokemon2: p2,
	}, nil
}

// ExecuteMove は指定された技を使用してポケモン間の攻撃を実行します。
// 攻撃側のポケモンが気絶している場合はエラーを返します。
//
// Parameters:
//   - attacker: 攻撃側のポケモン
//   - defender: 防御側のポケモン
//   - move: 使用する技
//
// Returns:
//   - error: 攻撃側が気絶している場合にエラーを返します
func (b *BattleEngine) ExecuteMove(attacker, defender *Pokemon, move Move) error {
	if attacker.IsFainted() {
		return errors.New("気絶したポケモンは技を使用できません")
	}

	damage := b.calculateDamage(attacker, defender, move)
	defender.TakeDamage(damage)

	return nil
}

// calculateDamage は技のダメージを計算します。
// 基本的なダメージ計算式を使用し、攻撃力、防御力、技の威力を考慮します。
//
// Parameters:
//   - attacker: 攻撃側のポケモン
//   - defender: 防御側のポケモン
//   - move: 使用する技
//
// Returns:
//   - int: 計算されたダメージ値
func (b *BattleEngine) calculateDamage(attacker, defender *Pokemon, move Move) int {
	baseDamage := (2*1+10)/250.0*float64(attacker.Attack())/float64(defender.Defense())*float64(move.Power()) + 2
	return int(math.Floor(baseDamage))
}

// BattleTurn は1ターンの対戦を実行します。
// 素早さ判定により行動順序を決定し、両方のポケモンの技を実行します。
// 2番目のポケモンが気絶している場合、その技は実行されません。
//
// Parameters:
//   - move1: 1番目のポケモンの技
//   - move2: 2番目のポケモンの技
//
// Returns:
//   - error: 技の実行中にエラーが発生した場合にエラーを返します
func (b *BattleEngine) BattleTurn(move1, move2 Move) error {
	first, second := b.pokemon1, b.pokemon2
	firstMove, secondMove := move1, move2

	if b.pokemon2.Speed() > b.pokemon1.Speed() {
		first, second = b.pokemon2, b.pokemon1
		firstMove, secondMove = move2, move1
	}

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
