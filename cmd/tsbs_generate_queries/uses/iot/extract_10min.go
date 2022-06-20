package iot

import (
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/utils"
	"github.com/timescale/tsbs/pkg/query"
)

// Extract10Min contains info for filling in extract 10 min queries.
type Extract10Min struct {
	core utils.QueryGenerator
}

// NewExtract10Min creates a new extract 10 min query filler.
func NewExtract10Min(core utils.QueryGenerator) utils.QueryFiller {
	return &Extract10Min{
		core: core,
	}
}

// Fill fills in the query.Query with query details.
func (i *Extract10Min) Fill(q query.Query) query.Query {
	fc, ok := i.core.(Extract10MinFiller)
	if !ok {
		common.PanicUnimplementedQuery(i.core)
	}
	fc.Extract10Min(q)
	return q
}
