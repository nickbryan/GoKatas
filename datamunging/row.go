package datamunging

type Row struct {
	Day      string
	Min, Max float64
}

func (r *Row) Spread() float64 {
	return r.Max - r.Min
}

type Rows []Row

func (rs *Rows) MinSpread() *Row {
	var ms *Row
	for _, r := range *rs {
		if ms == nil || r.Spread() <= ms.Spread() {
			ms = &r
		}
	}
	return ms
}
