package solitaire

func (g *Game) ruleCanMoveToBuild(build *Stack, card *Card) bool {
	if len(build.Cards) == 0 {
		return card.Value == 1
	}

	top := build.Top()
	return card.Suit == top.Suit && card.Value == top.Value+1
}
