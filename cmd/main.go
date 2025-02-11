package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"pokemon-battle/pkg/battle/domain"
	"pokemon-battle/pkg/battle/service"
)

// determineTurnOrder は素早さに基づいて行動順序を決定します
func determineTurnOrder(p1, p2 *domain.Pokemon, m1, m2 *domain.Move) (firstPokemon, secondPokemon *domain.Pokemon, firstMove, secondMove *domain.Move) {
	// 素早さが同じ場合はp1が先攻（実際には乱数を使用するべき）
	if p1.Speed >= p2.Speed {
		return p1, p2, m1, m2
	}
	return p2, p1, m2, m1
}

func main() {
	ctx := context.Background()

	// バトルサービスの作成
	battleService := service.NewBattleService()

	// ポケモンの作成
	pikachu := &domain.Pokemon{
		ID:        "pikachu-1",
		CurrentHP: 100,
		MaxHP:     100,
		Speed:     90,
		Status:    "NORMAL",
	}

	bulbasaur := &domain.Pokemon{
		ID:        "bulbasaur-1",
		CurrentHP: 100,
		MaxHP:     100,
		Speed:     45,
		Status:    "NORMAL",
	}

	// 技の作成
	thunder := &domain.Move{
		ID:       "thunder-1",
		Power:    90,
		Accuracy: 100,
		Priority: 0,
	}

	tackle := &domain.Move{
		ID:       "tackle-1",
		Power:    40,
		Accuracy: 100,
		Priority: 0,
	}

	// バトルの作成
	battle, err := domain.NewBattle("battle-1", pikachu, bulbasaur)
	if err != nil {
		log.Fatal(err)
	}

	// バトルの登録
	if err := battleService.RegisterBattle(battle); err != nil {
		log.Fatal(err)
	}

	// バトル開始
	fmt.Println("バトル開始！")
	fmt.Printf("ピカチュウ (HP: %d/素早さ: %d) VS フシギダネ (HP: %d/素早さ: %d)\n",
		pikachu.CurrentHP, pikachu.Speed,
		bulbasaur.CurrentHP, bulbasaur.Speed)
	fmt.Println("-------------------")

	turn := 1
	for battle.Status != domain.BattleStatusFinished {
		fmt.Printf("\nターン%d\n", turn)
		fmt.Println("-------------------")

		// 行動順序の決定
		firstPokemon, secondPokemon, firstMove, secondMove := determineTurnOrder(pikachu, bulbasaur, thunder, tackle)

		// ターンの実行
		if err := battleService.ExecuteMove(ctx, battle.ID, firstPokemon.ID, secondPokemon.ID, firstMove.ID); err != nil {
			log.Fatal(err)
		}

		// 状態の取得と表示
		status, err := battleService.GetBattleStatus(ctx, battle.ID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("1. %sの攻撃！\n", firstPokemon.ID)
		fmt.Printf("   ピカチュウの残りHP: %d\n", status.Pokemon1.CurrentHP)
		fmt.Printf("   フシギダネの残りHP: %d\n", status.Pokemon2.CurrentHP)

		time.Sleep(time.Second) // 行動間の待機時間

		// 2番目のポケモンの行動（1番目の攻撃で倒れていない場合）
		if !status.IsFinished {
			if err := battleService.ExecuteMove(ctx, battle.ID, secondPokemon.ID, firstPokemon.ID, secondMove.ID); err != nil {
				log.Fatal(err)
			}

			status, err = battleService.GetBattleStatus(ctx, battle.ID)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("2. %sの攻撃！\n", secondPokemon.ID)
			fmt.Printf("   ピカチュウの残りHP: %d\n", status.Pokemon1.CurrentHP)
			fmt.Printf("   フシギダネの残りHP: %d\n", status.Pokemon2.CurrentHP)
		}

		turn++
		time.Sleep(time.Second) // ターン間の待機時間
	}

	// 最終状態の取得
	finalStatus, err := battleService.GetBattleStatus(ctx, battle.ID)
	if err != nil {
		log.Fatal(err)
	}

	// 勝敗の判定と表示
	fmt.Println("\n-------------------")
	if finalStatus.Winner != nil {
		if *finalStatus.Winner == pikachu.ID {
			fmt.Println("ピカチュウの勝利！")
		} else {
			fmt.Println("フシギダネの勝利！")
		}
	}
}
