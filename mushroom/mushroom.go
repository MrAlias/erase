// Package mushroom provides mushrooms and information about them.
package mushroom

var Poisonous = []Mushroom{
	{"amanita muscaria", "fly agaric"},
	{"amanita ocreata", "death angle"},
	{"amanita pantherina", "panther cap"},
	{"galerina marginata", "deadly galerina"},
	{"hypholoma fasciculare", "sulphur tuft"},
	{"Omphalotus illudens", "jack-o'lantern mushroom"},
}

type Mushroom struct {
	BinomialName string
	CommonName   string
}
