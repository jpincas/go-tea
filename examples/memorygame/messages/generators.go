package messages

import gotea "github.com/jpincas/go-tea"

const (
	FlipCard_      = "FlipCard"
	FlipAllBack_   = "FlipAllBack"
	RemoveMatches_ = "RemoveMatches"
)

// Message generator

func FlipCard(index int) gotea.Message {
	return gotea.Message{
		FuncCode:  FlipCard_,
		Arguments: index,
	}
}

func FlipAllBack() gotea.Message {
	return gotea.Message{
		FuncCode:  FlipAllBack_,
		Arguments: nil,
	}
}

func RemoveMatches() gotea.Message {
	return gotea.Message{
		FuncCode:  RemoveMatches_,
		Arguments: nil,
	}
}
