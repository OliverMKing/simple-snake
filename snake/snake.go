package snake

import "simple-snake/model"

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

func (s *snake) Move(model.GameReq) *model.MoveResp {
	return &model.MoveResp{
		Move: "right",
	}
}
