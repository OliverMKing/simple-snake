package snake

import (
	"fmt"
	"math/rand"
	"simple-snake/model"
)

const (
	apiVersion = "1"
	author     = "OliverMKing"
	lightRed   = "#FE4365"
	tonugeHead = "tongue"
	blockTail  = "block-bum"
)

type Snake interface {
	Info() *model.InfoResp
	Move(model.GameReq) *model.MoveResp
}

type snake struct{}

var _ Snake = &snake{}

func New() Snake {
	return &snake{}
}

func (s *snake) Info() *model.InfoResp {
	return &model.InfoResp{
		ApiVersion: apiVersion,
		Author:     author,
		Color:      lightRed,
		Head:       tonugeHead,
		Tail:       blockTail,
	}
}

func (s *snake) Move(m model.GameReq) *model.MoveResp {
	id := m.You.Id

	moves := validMoves(id, model.GameReqSet(m))
	if len(moves) == 0 {
		fmt.Println("No valid moves")
		return &model.MoveResp{
			Move: model.Down,
		}
	}

	move := moves[rand.Intn(len(moves))]
	return &model.MoveResp{
		Move: move,
	}
}
