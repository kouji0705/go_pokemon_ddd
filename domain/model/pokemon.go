// Package model はポケモンバトルのドメインモデルを提供します。
package model

import "errors"

// Pokemon はポケモンの基本的な属性と状態を表現する構造体です。
// 名前、HP、攻撃力、防御力、素早さなどの基本ステータスと、
// 覚えている技を管理します。
type Pokemon struct {
	name      string // ポケモンの名前
	maxHP     int    // 最大HP
	currentHP int    // 現在のHP
	attack    int    // 攻撃力
	defense   int    // 防御力
	speed     int    // 素早さ
	moves     []Move // 覚えている技
}

// NewPokemon は新しいポケモンインスタンスを生成します。
// 名前が空の場合、またはステータスが0以下の場合はエラーを返します。
//
// Parameters:
//   - name: ポケモンの名前
//   - hp: 最大HP（1以上である必要があります）
//   - attack: 攻撃力（1以上である必要があります）
//   - defense: 防御力（1以上である必要があります）
//   - speed: 素早さ（1以上である必要があります）
//
// Returns:
//   - *Pokemon: 生成されたポケモンインスタンス
//   - error: パラメータが不正な場合のエラー
func NewPokemon(name string, hp, attack, defense, speed int) (*Pokemon, error) {
	if name == "" {
		return nil, errors.New("ポケモンの名前は必須です")
	}
	if hp <= 0 || attack <= 0 || defense <= 0 || speed <= 0 {
		return nil, errors.New("ステータスは0より大きい値である必要があります")
	}

	return &Pokemon{
		name:      name,
		maxHP:     hp,
		currentHP: hp,
		attack:    attack,
		defense:   defense,
		speed:     speed,
		moves:     make([]Move, 0, 4),
	}, nil
}

// Name はポケモンの名前を返します。
//
// Returns:
//   - string: ポケモンの名前
func (p *Pokemon) Name() string {
	return p.name
}

// CurrentHP はポケモンの現在のHPを返します。
//
// Returns:
//   - int: 現在のHP値
func (p *Pokemon) CurrentHP() int {
	return p.currentHP
}

// Attack はポケモンの攻撃力を返します。
//
// Returns:
//   - int: 攻撃力の値
func (p *Pokemon) Attack() int {
	return p.attack
}

// Defense はポケモンの防御力を返します。
//
// Returns:
//   - int: 防御力の値
func (p *Pokemon) Defense() int {
	return p.defense
}

// Speed はポケモンの素早さを返します。
//
// Returns:
//   - int: 素早さの値
func (p *Pokemon) Speed() int {
	return p.speed
}

// AddMove はポケモンに新しい技を追加します。
// すでに4つの技を覚えている場合はエラーを返します。
//
// Parameters:
//   - move: 追加する技
//
// Returns:
//   - error: 技の追加に失敗した場合のエラー
func (p *Pokemon) AddMove(move Move) error {
	if len(p.moves) >= 4 {
		return errors.New("技は4つまでしか覚えられません")
	}
	p.moves = append(p.moves, move)
	return nil
}

// TakeDamage はポケモンにダメージを与えます。
// HPが0未満になる場合は0に設定されます。
//
// Parameters:
//   - damage: 与えるダメージ量
func (p *Pokemon) TakeDamage(damage int) {
	p.currentHP -= damage
	if p.currentHP < 0 {
		p.currentHP = 0
	}
}

// IsFainted はポケモンが気絶状態（HP0）かどうかを返します。
//
// Returns:
//   - bool: 気絶している場合はtrue、そうでない場合はfalse
func (p *Pokemon) IsFainted() bool {
	return p.currentHP <= 0
}
