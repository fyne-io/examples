package solitaire

func (g *Game) ruleCanMoveToBuild(build *Stack, card *Card) bool {
	if len(build.Cards) == 0 {
		return card.Value == 1
	}

	top := build.Top()
	return card.Suit == top.Suit && card.Value == top.Value+1
}

func (g *Game) ruleCanMoveToStack(stack *Stack, card *Card) bool {
	if len(stack.Cards) == 0 {
		return card.Value == ValueKing
	}

	top := stack.Top()
	if top.Color() == card.Color() {
		return false
	}
	return card.Value == top.Value-1
}
