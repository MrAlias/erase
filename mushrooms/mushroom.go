// Package mushrooms provides mushrooms and information about them.
package mushrooms

var Poisonous = []Mushroom{
	{"amanita muscaria", "fly agaric"},
	{"amanita ocreata", "death angle"},
	{"amanita pantherina", "panther cap"},
	{"galerina marginata", "deadly galerina"},
	{"hypholoma fasciculare", "sulphur tuft"},
	{"omphalotus illudens", "jack-o'lantern mushroom"},
	{"pholiotina rugosa", ""},
}

type Mushroom struct {
	BinomialName string
	CommonName   string
}
