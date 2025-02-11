package model

import "errors"

type Pokemon struct {
	name      string
	maxHP     int
	currentHP int
	attack    int
	defense   int
	speed     int
	moves     []Move
}

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

func (p *Pokemon) Name() string {
	return p.name
}

func (p *Pokemon) CurrentHP() int {
	return p.currentHP
}

func (p *Pokemon) Attack() int {
	return p.attack
}

func (p *Pokemon) Defense() int {
	return p.defense
}

func (p *Pokemon) Speed() int {
	return p.speed
}

func (p *Pokemon) AddMove(move Move) error {
	if len(p.moves) >= 4 {
		return errors.New("技は4つまでしか覚えられません")
	}
	p.moves = append(p.moves, move)
	return nil
}

func (p *Pokemon) TakeDamage(damage int) {
	p.currentHP -= damage
	if p.currentHP < 0 {
		p.currentHP = 0
	}
}

func (p *Pokemon) IsFainted() bool {
	return p.currentHP <= 0
}
