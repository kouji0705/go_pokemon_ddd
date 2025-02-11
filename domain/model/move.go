package model

import "errors"

type Move struct {
    name     string
    power    int
    accuracy int
}

func NewMove(name string, power, accuracy int) (*Move, error) {
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

func (m *Move) Name() string {
    return m.name
}

func (m *Move) Power() int {
    return m.power
}

func (m *Move) Accuracy() int {
    return m.accuracy
} 