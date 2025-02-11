package main

import (
	"fmt"
	"log"
	"pokemon-battle/application/service"
	"pokemon-battle/domain/model"
)

func main() {
	// ポケモンの作成
	pikachu, err := model.NewPokemon("ピカチュウ", 35, 55, 40, 90)
	if err != nil {
		log.Fatal(err)
	}

	bulbasaur, err := model.NewPokemon("フシギダネ", 45, 49, 49, 45)
	if err != nil {
		log.Fatal(err)
	}

	// 技の作成
	thunder, err := model.NewMove("10万ボルト", 90, 100)
	if err != nil {
		log.Fatal(err)
	}

	tackle, err := model.NewMove("たいあたり", 40, 100)
	if err != nil {
		log.Fatal(err)
	}

	// 技をポケモンに覚えさせる
	if err := pikachu.AddMove(*thunder); err != nil {
		log.Fatal(err)
	}
	if err := bulbasaur.AddMove(*tackle); err != nil {
		log.Fatal(err)
	}

	// バトルサービスの作成
	battleService, err := service.NewBattleService(pikachu, bulbasaur)
	if err != nil {
		log.Fatal(err)
	}

	// ターンの実行
	if err := battleService.ExecuteTurn(*thunder, *tackle); err != nil {
		log.Fatal(err)
	}

	// 結果の表示
	fmt.Printf("%sの残りHP: %d\n", pikachu.Name(), pikachu.CurrentHP())
	fmt.Printf("%sの残りHP: %d\n", bulbasaur.Name(), bulbasaur.CurrentHP())
}
