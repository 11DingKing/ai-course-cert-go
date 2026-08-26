package pagination

type Page struct{ Limit, Offset int }

func Normalize(limit, offset int) Page {
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}
	return Page{limit, offset}
}
