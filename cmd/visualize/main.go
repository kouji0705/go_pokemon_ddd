package main

import (
	"fmt"
	"pokemon-battle/pkg/battle/domain"
	"pokemon-battle/pkg/battle/ports"
	"pokemon-battle/pkg/battle/service"
)

func main() {
	// このファイルは依存関係を視覚化するためのエントリーポイントとして機能します
	fmt.Println("Package visualization entry point")

	// 依存関係を明示的に示すためのダミー使用
	_ = &domain.Battle{}
	_ = &service.BattleService{}
	var _ ports.BattlePort = &service.BattleService{}
}
