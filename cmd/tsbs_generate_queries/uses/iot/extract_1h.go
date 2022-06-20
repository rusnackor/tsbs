package iot

import (
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/utils"
	"github.com/timescale/tsbs/pkg/query"
)

// Extract1H contains info for filling in extract 1 hour queries.
type Extract1H struct {
	core utils.QueryGenerator
}

// NewExtract1H creates a new extract 1 hour query filler.
func NewExtract1H(core utils.QueryGenerator) utils.QueryFiller {
	return &Extract1H{
		core: core,
	}
}

// Fill fills in the query.Query with query details.
func (i *Extract1H) Fill(q query.Query) query.Query {
	fc, ok := i.core.(Extract1HFiller)
	if !ok {
		common.PanicUnimplementedQuery(i.core)
	}
	fc.Extract1H(q)
	return q
}
