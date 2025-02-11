package ports

import (
	"context"
)

// BattlePort はバトルドメインの公開インターフェースを定義します
type BattlePort interface {
	// ファクトリーメソッド
	CreateBattle(ctx context.Context, id string, pokemon1ID, pokemon2ID string) error
	CreatePokemon(ctx context.Context, id string, maxHP int, speed int) error
	CreateMove(ctx context.Context, id string, power int, accuracy int, priority int) error

	// バトル操作
	ExecuteTurn(ctx context.Context, battleID string, move1ID string, move2ID string) error
	ExecuteMove(ctx context.Context, battleID string, attackerID string, defenderID string, moveID string) error
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
