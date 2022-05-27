// Package mushrooms provides mushrooms and information about them.
package mushrooms

var Tasty = []Mushroom{
	{"agaricus bisporus", "button mushroom"},
	{"boletus edulis", "porcini"},
	{"cantharellus cibarius", "chanterelle"},
	{"clitocybe nuda", "blewit"},
	{"lactarius deliciosus", "saffron milk cap"},
	{"morchella esculenta", "morel"},
	{"pleurotus ostreatus", "oyster mushroom"},
	{"tricholoma matsutake", "matsutake"},
}

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
