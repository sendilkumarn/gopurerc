package common

type COLOR struct {
	BLACK  string
	GRAY   string
	WHITE  string
	PURPLE string
	GREEN  string
	RED    string
	ORANGE string
}

var Color = COLOR{
	/** In use or free. */
	BLACK: "BLACK",
	/** Possible member of cycle. */
	GRAY: "GRAY",
	/** Member of cycle. */
	WHITE: "WHITE",
	/** Possible root of cycle. */
	PURPLE: "PURPLE",
	/** Acyclic. */
	GREEN: "GREEN",
	/** Candidate cycle undergoing Σ-computation. */
	RED: "RED",
	/** Candidate cycle awaiting epoch boundary. */
	ORANGE: "ORANGE",
}
