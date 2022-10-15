package snake

import "simple-snake/model"

type path func(id string) (model.Move, error)
