// Package model はポケモンバトルのドメインモデルを提供します。
package model

import (
	"errors"
)

// Move はポケモンの技を表す構造体です
type Move struct {
	name     string // 技の名前
	power    int    // 技の威力
	accuracy int    // 技の命中率
}

// NewMove は新しいMoveインスタンスを生成します
func NewMove(name string, power int, accuracy int) (*Move, error) {
	if name == "" {
		return nil, errors.New("技の名前は必須です")
	}
	if power < 0 {
		return nil, errors.New("技の威力は0以上である必要があります")
	}
	if accuracy <= 0 || accuracy > 100 {
		return nil, errors.New("命中率は1から100の間である必要があります")
	}

	return &Move{
		name:     name,
		power:    power,
		accuracy: accuracy,
	}, nil
}

// Name は技の名前を返します
func (m *Move) Name() string {
	return m.name
}

// Power は技の威力を返します
func (m *Move) Power() int {
	return m.power
}

// Accuracy は技の命中率を返します
func (m *Move) Accuracy() int {
	return m.accuracy
}
