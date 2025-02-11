package ports

import (
	"context"
)

// BattlePort はバトルドメインの公開インターフェースを定義します
type BattlePort interface {
	// ExecuteTurn は1ターンのバトルを実行します
	ExecuteTurn(ctx context.Context, battleID string, move1ID string, move2ID string) error

	// ExecuteMove は1回の攻撃を実行します
	ExecuteMove(ctx context.Context, battleID string, attackerID string, defenderID string, moveID string) error

	// GetBattleStatus はバトルの現在の状態を取得します
	GetBattleStatus(ctx context.Context, battleID string) (*BattleStatus, error)
}

// BattleStatus はバトルの状態を表現します
type BattleStatus struct {
	BattleID   string
	Pokemon1   PokemonStatus
	Pokemon2   PokemonStatus
	IsFinished bool
	Winner     *string
}

// PokemonStatus はポケモンの状態を表現します
type PokemonStatus struct {
	PokemonID string
	CurrentHP int
	MaxHP     int
	Status    string
}
