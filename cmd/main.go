package main

import (
	"fmt"
	"log"
	"pokemon-battle/application/service"
	"pokemon-battle/domain/model"
)

// createPokemon はポケモンを作成し、技を覚えさせます
func createPokemon() (*model.Pokemon, *model.Pokemon, *model.Move, *model.Move, error) {
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

func main() {
	// ポケモンの作成と技の習得
	pikachu, bulbasaur, thunder, tackle, err := createPokemon()
	if err != nil {
		log.Fatal(err)
	}

	// バトルサービスの作成
	battleService, err := service.NewBattleService(pikachu, bulbasaur)
	if err != nil {
		log.Fatal(err)
	}

	// ターンの実行
	if err := battleService.ExecuteTurn(thunder, tackle); err != nil {
		log.Fatal(err)
	}

	// 結果の表示
	fmt.Printf("%sの残りHP: %d\n", pikachu.Name(), pikachu.CurrentHP())
	fmt.Printf("%sの残りHP: %d\n", bulbasaur.Name(), bulbasaur.CurrentHP())
}
