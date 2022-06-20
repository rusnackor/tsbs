package iot

import (
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/utils"
	"github.com/timescale/tsbs/pkg/query"
)

// TestQuery contains info for filling in PoC to adding new queries.
type TestQuery struct {
	core utils.QueryGenerator
}

// NewTestQuery creates a new PoC to adding new query filler.
func NewTestQuery(core utils.QueryGenerator) utils.QueryFiller {
	return &TestQuery{
		core: core,
	}
}

// Fill fills in the query.Query with query details.
func (i *TestQuery) Fill(q query.Query) query.Query {
	fc, ok := i.core.(TestQueryFiller)
	if !ok {
		common.PanicUnimplementedQuery(i.core)
	}
	fc.TestQuery(q)
	return q
}
