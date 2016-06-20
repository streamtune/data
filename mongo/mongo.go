package mongo

import (
	"github.com/streamtune/data"
	"gopkg.in/mgo.v2"
)

// ApplyPageable will apply the given pageable object to provided query parameters, returning the updated query
func ApplyPageable(pageable *data.Pageable, query *mgo.Query) *mgo.Query {
	return applySort(pageable.Sort.Orders, query.Skip(pageable.Offset()).Limit(pageable.Size))
}

func applySort(ordering []data.Order, query *mgo.Query) *mgo.Query {
	fields := make([]string, len(ordering))
	for i, clause := range ordering {
		fields[i] = clause.Property
		if clause.Direction == data.Desc {
			fields[i] = "-" + fields[i]
		}
	}
	return query.Sort(fields...)
}
