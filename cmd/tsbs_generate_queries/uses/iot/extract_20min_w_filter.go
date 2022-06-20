package iot

import (
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/utils"
	"github.com/timescale/tsbs/pkg/query"
)

// Extract20MinWfilter contains info for filling in extract 20 min with filter queries.
type Extract20MinWfilter struct {
	core utils.QueryGenerator
}

// NewExtract20MinWfilter creates a new extract 20 min with filter query filler.
func NewExtract20MinWfilter(core utils.QueryGenerator) utils.QueryFiller {
	return &Extract20MinWfilter{
		core: core,
	}
}

// Fill fills in the query.Query with query details.
func (i *Extract20MinWfilter) Fill(q query.Query) query.Query {
	fc, ok := i.core.(Extract20MinWfilterFiller)
	if !ok {
		common.PanicUnimplementedQuery(i.core)
	}
	fc.Extract20MinWfilter(q)
	return q
}
