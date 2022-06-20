package iot

import (
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/utils"
	"github.com/timescale/tsbs/pkg/query"
)

// Statistics24H contains info for filling in statistics 24 hours queries.
type Statistics24H struct {
	core utils.QueryGenerator
}

// NewStatistics24H creates a new statistics 24 hours query filler.
func NewStatistics24H(core utils.QueryGenerator) utils.QueryFiller {
	return &Statistics24H{
		core: core,
	}
}

// Fill fills in the query.Query with query details.
func (i *Statistics24H) Fill(q query.Query) query.Query {
	fc, ok := i.core.(Statistics24HFiller)
	if !ok {
		common.PanicUnimplementedQuery(i.core)
	}
	fc.Statistics24H(q)
	return q
}
