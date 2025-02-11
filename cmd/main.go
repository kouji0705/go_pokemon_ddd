package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"pokemon-battle/pkg/battle/service"
)

func main() {
	ctx := context.Background()

	// バトルサービスの作成
	battleService := service.NewBattleService()

	// ポケモンの作成
	if err := battleService.CreatePokemon(ctx, "pikachu-1", 100, 90); err != nil {
		log.Fatal(err)
	}
	if err := battleService.CreatePokemon(ctx, "bulbasaur-1", 100, 45); err != nil {
		log.Fatal(err)
	}

	// 技の作成
	if err := battleService.CreateMove(ctx, "thunder-1", 90, 100, 0); err != nil {
		log.Fatal(err)
	}
	if err := battleService.CreateMove(ctx, "tackle-1", 40, 100, 0); err != nil {
		log.Fatal(err)
	}

	// バトルの作成
	if err := battleService.CreateBattle(ctx, "battle-1", "pikachu-1", "bulbasaur-1"); err != nil {
		log.Fatal(err)
	}

	// 初期状態の取得
	status, err := battleService.GetBattleStatus(ctx, "battle-1")
	if err != nil {
		log.Fatal(err)
	}

	// バトル開始
	fmt.Println("バトル開始！")
	fmt.Printf("ピカチュウ (HP: %d) VS フシギダネ (HP: %d)\n",
		status.Pokemon1.CurrentHP,
		status.Pokemon2.CurrentHP)
	fmt.Println("-------------------")

	turn := 1
	for !status.IsFinished {
		fmt.Printf("\nターン%d\n", turn)
		fmt.Println("-------------------")

		// 1番目のポケモンの行動
		if err := battleService.ExecuteMove(ctx, "battle-1", status.Pokemon1.PokemonID, status.Pokemon2.PokemonID, "thunder-1"); err != nil {
			log.Fatal(err)
		}

		// 状態の取得と表示
		status, err = battleService.GetBattleStatus(ctx, "battle-1")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("1. ピカチュウのかみなり！\n")
		fmt.Printf("   ピカチュウの残りHP: %d\n", status.Pokemon1.CurrentHP)
		fmt.Printf("   フシギダネの残りHP: %d\n", status.Pokemon2.CurrentHP)

		time.Sleep(time.Second) // 行動間の待機時間

		// 2番目のポケモンの行動（1番目の攻撃で倒れていない場合）
		if !status.IsFinished {
			if err := battleService.ExecuteMove(ctx, "battle-1", status.Pokemon2.PokemonID, status.Pokemon1.PokemonID, "tackle-1"); err != nil {
				log.Fatal(err)
			}

			status, err = battleService.GetBattleStatus(ctx, "battle-1")
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("2. フシギダネのたいあたり！\n")
			fmt.Printf("   ピカチュウの残りHP: %d\n", status.Pokemon1.CurrentHP)
			fmt.Printf("   フシギダネの残りHP: %d\n", status.Pokemon2.CurrentHP)
		}

		turn++
		time.Sleep(time.Second) // ターン間の待機時間
	}

	// 勝敗の判定と表示
	fmt.Println("\n-------------------")
	if status.Winner != nil {
		if *status.Winner == "pikachu-1" {
			fmt.Println("ピカチュウの勝利！")
		} else {
			fmt.Println("フシギダネの勝利！")
		}
	}
}
