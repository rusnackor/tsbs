package iot

import (
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/utils"
	"github.com/timescale/tsbs/pkg/query"
)

// Statistics10Min contains info for filling in statistics 10 min queries.
type Statistics10Min struct {
	core utils.QueryGenerator
}

// NewStatistics10Min creates a new statistics 10 min query filler.
func NewStatistics10Min(core utils.QueryGenerator) utils.QueryFiller {
	return &Statistics10Min{
		core: core,
	}
}

// Fill fills in the query.Query with query details.
func (i *Statistics10Min) Fill(q query.Query) query.Query {
	fc, ok := i.core.(Statistics10MinFiller)
	if !ok {
		common.PanicUnimplementedQuery(i.core)
	}
	fc.Statistics10Min(q)
	return q
}
