package main

import (
	"fmt"
	"log"
	"pokemon-battle/application/service"
	"pokemon-battle/domain/model"
	"time"
)

// determineTurnOrder は素早さに基づいて行動順序を決定します
func determineTurnOrder(p1, p2 *model.Pokemon, m1, m2 *model.Move) (firstPokemon, secondPokemon *model.Pokemon, firstMove, secondMove *model.Move) {
	if p1.Speed() > p2.Speed() {
		return p1, p2, m1, m2
	}
	return p2, p1, m2, m1
}

func main() {
	// ポケモンサービスの作成
	pokemonService := service.NewPokemonService()

	// ポケモンの作成と技の習得
	pikachu, bulbasaur, thunder, tackle, err := pokemonService.CreateInitialPokemon()
	if err != nil {
		log.Fatal(err)
	}

	// バトルサービスの作成
	battleService, err := service.NewBattleService(pikachu, bulbasaur)
	if err != nil {
		log.Fatal(err)
	}

	// バトル開始
	fmt.Println("バトル開始！")
	fmt.Printf("%s (HP: %d/素早さ: %d) VS %s (HP: %d/素早さ: %d)\n",
		pikachu.Name(), pikachu.CurrentHP(), pikachu.Speed(),
		bulbasaur.Name(), bulbasaur.CurrentHP(), bulbasaur.Speed())
	fmt.Println("-------------------")

	turn := 1
	for !pikachu.IsFainted() && !bulbasaur.IsFainted() {
		fmt.Printf("\nターン%d\n", turn)
		fmt.Println("-------------------")

		// 行動順序の決定
		firstPokemon, secondPokemon, firstMove, secondMove := determineTurnOrder(pikachu, bulbasaur, thunder, tackle)

		// 1番目のポケモンの行動
		fmt.Printf("1. %s の %s！\n", firstPokemon.Name(), firstMove.Name())
		if err := battleService.ExecuteMove(firstPokemon, secondPokemon, firstMove); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("   %s の残りHP: %d\n", pikachu.Name(), pikachu.CurrentHP())
		fmt.Printf("   %s の残りHP: %d\n", bulbasaur.Name(), bulbasaur.CurrentHP())

		time.Sleep(time.Second) // 行動間の待機時間

		// 2番目のポケモンの行動（1番目の攻撃で倒れていない場合）
		if !secondPokemon.IsFainted() {
			fmt.Printf("2. %s の %s！\n", secondPokemon.Name(), secondMove.Name())
			if err := battleService.ExecuteMove(secondPokemon, firstPokemon, secondMove); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("   %s の残りHP: %d\n", pikachu.Name(), pikachu.CurrentHP())
			fmt.Printf("   %s の残りHP: %d\n", bulbasaur.Name(), bulbasaur.CurrentHP())
		}

		turn++
		time.Sleep(time.Second) // ターン間の待機時間
	}

	// 勝敗の判定
	fmt.Println("\n-------------------")
	if pikachu.IsFainted() {
		fmt.Printf("%sの勝利！\n", bulbasaur.Name())
	} else {
		fmt.Printf("%sの勝利！\n", pikachu.Name())
	}
}
