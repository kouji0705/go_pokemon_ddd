// Package model はポケモンバトルのドメインモデルを提供します。
package model

import (
	"errors"
)

// Pokemon はポケモンを表す構造体です
type Pokemon struct {
	name      string  // ポケモンの名前
	maxHP     int     // 最大HP
	currentHP int     // 現在のHP
	attack    int     // 攻撃力
	defense   int     // 防御力
	speed     int     // 素早さ
	moves     []*Move // 覚えている技
}

// NewPokemon は新しいPokemonインスタンスを生成します
func NewPokemon(name string, maxHP, attack, defense, speed int) (*Pokemon, error) {
	if name == "" {
		return nil, errors.New("ポケモンの名前は必須です")
	}
	if maxHP <= 0 || attack <= 0 || defense <= 0 || speed <= 0 {
		return nil, errors.New("ステータスは0より大きい値である必要があります")
	}

	return &Pokemon{
		name:      name,
		maxHP:     maxHP,
		currentHP: maxHP,
		attack:    attack,
		defense:   defense,
		speed:     speed,
		moves:     make([]*Move, 0),
	}, nil
}

// Name はポケモンの名前を返します
func (p *Pokemon) Name() string {
	return p.name
}

// MaxHP は最大HPを返します
func (p *Pokemon) MaxHP() int {
	return p.maxHP
}

// CurrentHP は現在のHPを返します
func (p *Pokemon) CurrentHP() int {
	return p.currentHP
}

// Attack は攻撃力を返します
func (p *Pokemon) Attack() int {
	return p.attack
}

// Defense は防御力を返します
func (p *Pokemon) Defense() int {
	return p.defense
}

// Speed は素早さを返します
func (p *Pokemon) Speed() int {
	return p.speed
}

// Moves は覚えている技のリストを返します
func (p *Pokemon) Moves() []*Move {
	return p.moves
}

// AddMove は新しい技を追加します
func (p *Pokemon) AddMove(move *Move) bool {
	if len(p.moves) >= 4 {
		return false
	}
	p.moves = append(p.moves, move)
	return true
}

// TakeDamage はダメージを受けた時の処理を行います
func (p *Pokemon) TakeDamage(damage int) {
	p.currentHP -= damage
	if p.currentHP < 0 {
		p.currentHP = 0
	}
}

// IsFainted はひんしになっているかを判定します
func (p *Pokemon) IsFainted() bool {
	return p.currentHP <= 0
}

// Heal は回復処理を行います
func (p *Pokemon) Heal(amount int) {
	p.currentHP += amount
	if p.currentHP > p.maxHP {
		p.currentHP = p.maxHP
	}
}
