package model

import (
	"testing"
)

func TestNewMove(t *testing.T) {
	tests := []struct {
		name       string
		moveName   string
		power      int
		accuracy   int
		wantErr    bool
		errMessage string
	}{
		{
			name:     "正常なケース",
			moveName: "でんこうせっか",
			power:    40,
			accuracy: 100,
			wantErr:  false,
		},
		{
			name:       "技名が空の場合",
			moveName:   "",
			power:      40,
			accuracy:   100,
			wantErr:    true,
			errMessage: "技の名前は必須です",
		},
		{
			name:       "威力が負の値の場合",
			moveName:   "でんこうせっか",
			power:      -1,
			accuracy:   100,
			wantErr:    true,
			errMessage: "技の威力は0以上である必要があります",
		},
		{
			name:       "命中率が0以下の場合",
			moveName:   "でんこうせっか",
			power:      40,
			accuracy:   0,
			wantErr:    true,
			errMessage: "命中率は1から100の間である必要があります",
		},
		{
			name:       "命中率が100より大きい場合",
			moveName:   "でんこうせっか",
			power:      40,
			accuracy:   101,
			wantErr:    true,
			errMessage: "命中率は1から100の間である必要があります",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			move, err := NewMove(tt.moveName, tt.power, tt.accuracy)

			// エラーチェック
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// エラーメッセージのチェック
			if tt.wantErr && err.Error() != tt.errMessage {
				t.Errorf("NewMove() error message = %v, want %v", err.Error(), tt.errMessage)
				return
			}

			// 正常系の場合、値のチェック
			if !tt.wantErr {
				if move.Name() != tt.moveName {
					t.Errorf("move.Name() = %v, want %v", move.Name(), tt.moveName)
				}
				if move.Power() != tt.power {
					t.Errorf("move.Power() = %v, want %v", move.Power(), tt.power)
				}
				if move.Accuracy() != tt.accuracy {
					t.Errorf("move.Accuracy() = %v, want %v", move.Accuracy(), tt.accuracy)
				}
			}
		})
	}
}

func TestMove_Getters(t *testing.T) {
	// テスト用の技を作成
	move, err := NewMove("でんこうせっか", 40, 100)
	if err != nil {
		t.Fatalf("Failed to create test move: %v", err)
	}

	// Name()のテスト
	t.Run("Name getter", func(t *testing.T) {
		if got := move.Name(); got != "でんこうせっか" {
			t.Errorf("Move.Name() = %v, want %v", got, "でんこうせっか")
		}
	})

	// Power()のテスト
	t.Run("Power getter", func(t *testing.T) {
		if got := move.Power(); got != 40 {
			t.Errorf("Move.Power() = %v, want %v", got, 40)
		}
	})

	// Accuracy()のテスト
	t.Run("Accuracy getter", func(t *testing.T) {
		if got := move.Accuracy(); got != 100 {
			t.Errorf("Move.Accuracy() = %v, want %v", got, 100)
		}
	})
}
