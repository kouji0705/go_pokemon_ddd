package service

import (
	"fmt"
	"pokemon-battle/domain/model"
)

// PokemonService はポケモンの生成と管理を担当します
type PokemonService struct{}

// NewPokemonService は新しいPokemonServiceを作成します
func NewPokemonService() *PokemonService {
	return &PokemonService{}
}

// CreateInitialPokemon は初期ポケモンを作成します
func (s *PokemonService) CreateInitialPokemon() (*model.Pokemon, *model.Pokemon, *model.Move, *model.Move, error) {
	// ピカチュウの作成
	pikachu, err := model.NewPokemon("ピカチュウ", 35, 55, 40, 90)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("ピカチュウの作成に失敗: %w", err)
	}

	// フシギダネの作成
	bulbasaur, err := model.NewPokemon("フシギダネ", 45, 49, 49, 45)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("フシギダネの作成に失敗: %w", err)
	}

	// 技の作成
	thunder, err := model.NewMove("10万ボルト", 90, 100)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("10万ボルトの作成に失敗: %w", err)
	}

	tackle, err := model.NewMove("たいあたり", 40, 100)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("たいあたりの作成に失敗: %w", err)
	}

	// 技をポケモンに覚えさせる
	if !pikachu.AddMove(thunder) {
		return nil, nil, nil, nil, fmt.Errorf("ピカチュウに技を覚えさせることができませんでした")
	}

	if !bulbasaur.AddMove(tackle) {
		return nil, nil, nil, nil, fmt.Errorf("フシギダネに技を覚えさせることができませんでした")
	}

	return pikachu, bulbasaur, thunder, tackle, nil
}
