package csvParse

import (
	"errors"
	"strconv"
)

type Good struct {
	Name  string
	Price float64
	Sales int
	Total float64
}

type AllGoods map[string]Good

func (g AllGoods) GetHighestRevProduct() Good {
	var highest Good
	for _, v := range g {
		if v.Total > highest.Total {
			highest = v
		}
	}
	return highest
}

func (g AllGoods) GetTotal() float64 {
	var total float64
	for _, v := range g {
		total += v.Total
	}
	return total
}

func Parse(strGoods [][]string) (AllGoods, error) {
	goods := make(AllGoods, len(strGoods))
	for _, g := range strGoods {

		s, err := strconv.Atoi(g[2])
		if err != nil {
			return nil, err
		}
		p, err := strconv.ParseFloat(g[1], 32)
		if err != nil {
			return nil, err
		}

		if val, ok := goods[g[0]]; ok {
			val.Sales += s
			if val.Price != p {
				return nil, errors.New("price not equal in good " + g[0]) //looks like price mustn't change in file
			}
			val.Total += p * float64(s)
			goods[g[0]] = val
		} else {
			goods[g[0]] = Good{
				Name:  g[0],
				Price: p,
				Sales: s,
				Total: p * float64(s),
			}
		}
	}
	return goods, nil
}
