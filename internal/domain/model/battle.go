// Package model はポケモンバトルのドメインモデルを提供します。
package model

import (
	"math/rand"
)

// BattleEngine はポケモンバトルの進行を制御する構造体です
type BattleEngine struct {
	pokemon1 *Pokemon
	pokemon2 *Pokemon
}

// NewBattleEngine は新しいBattleEngineインスタンスを生成します
func NewBattleEngine(pokemon1, pokemon2 *Pokemon) *BattleEngine {
	return &BattleEngine{
		pokemon1: pokemon1,
		pokemon2: pokemon2,
	}
}

// ExecuteMove は1つの技を実行します
func (b *BattleEngine) ExecuteMove(attacker, defender *Pokemon, move *Move) {
	// 命中判定
	if !b.checkAccuracy(move) {
		return
	}

	// ダメージ計算
	damage := b.calculateDamage(attacker, defender, move)
	defender.TakeDamage(damage)
}

// BattleTurn はターンの処理を実行します
func (b *BattleEngine) BattleTurn(move1, move2 *Move) {
	// 素早さ判定で行動順を決定
	firstPokemon, secondPokemon := b.pokemon1, b.pokemon2
	firstMove, secondMove := move1, move2

	if b.pokemon2.Speed() > b.pokemon1.Speed() {
		firstPokemon, secondPokemon = b.pokemon2, b.pokemon1
		firstMove, secondMove = move2, move1
	}

	// 先攻ポケモンの攻撃
	b.ExecuteMove(firstPokemon, secondPokemon, firstMove)
	if !secondPokemon.IsFainted() {
		// 後攻ポケモンの攻撃
		b.ExecuteMove(secondPokemon, firstPokemon, secondMove)
	}
}

// checkAccuracy は技が命中するかどうかを判定します
func (b *BattleEngine) checkAccuracy(move *Move) bool {
	return rand.Intn(100) < move.Accuracy()
}

// calculateDamage はダメージを計算します
func (b *BattleEngine) calculateDamage(attacker, defender *Pokemon, move *Move) int {
	// 基本的なダメージ計算式
	damage := ((2*50/5 + 2) * move.Power() * attacker.Attack() / defender.Defense() / 50) + 2

	// 乱数補正（85%～100%）
	damage = damage * (85 + rand.Intn(16)) / 100

	return damage
}

// Pokemon1 は1番目のポケモンを返します
func (b *BattleEngine) Pokemon1() *Pokemon {
	return b.pokemon1
}

// Pokemon2 は2番目のポケモンを返します
func (b *BattleEngine) Pokemon2() *Pokemon {
	return b.pokemon2
}
