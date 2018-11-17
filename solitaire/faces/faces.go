package faces

import "github.com/fyne-io/fyne"

// ForCard returns the face resource for the specified card value and suit.
func ForCard(card, suit int) fyne.Resource {
	return faceResources[card-1+(suit*13)]
}

// ForBack returns a face resource for the back of a card
func ForBack() fyne.Resource {
	return back
}

// ForSpace returns a special resource to use when a vacant spot should be indicated
func ForSpace() fyne.Resource {
	return space
}

var faceResources = [52]fyne.Resource{
	cardAC,
	card2C,
	card3C,
	card4C,
	card5C,
	card6C,
	card7C,
	card8C,
	card9C,
	card10C,
	cardJC,
	cardQC,
	cardKC,

	cardAD,
	card2D,
	card3D,
	card4D,
	card5D,
	card6D,
	card7D,
	card8D,
	card9D,
	card10D,
	cardJD,
	cardQD,
	cardKD,

	cardAH,
	card2H,
	card3H,
	card4H,
	card5H,
	card6H,
	card7H,
	card8H,
	card9H,
	card10H,
	cardJH,
	cardQH,
	cardKH,

	cardAS,
	card2S,
	card3S,
	card4S,
	card5S,
	card6S,
	card7S,
	card8S,
	card9S,
	card10S,
	cardJS,
	cardQS,
	cardKS,
}
