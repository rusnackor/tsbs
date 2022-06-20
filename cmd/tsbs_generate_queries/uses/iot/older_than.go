package iot

import (
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/utils"
	"github.com/timescale/tsbs/pkg/query"
)

// OlderThan contains info for filling in older than queries.
type OlderThan struct {
	core utils.QueryGenerator
}

// NewOlderThan creates a new older than query filler.
func NewOlderThan(core utils.QueryGenerator) utils.QueryFiller {
	return &OlderThan{
		core: core,
	}
}

// Fill fills in the query.Query with query details.
func (i *OlderThan) Fill(q query.Query) query.Query {
	fc, ok := i.core.(OlderThanFiller)
	if !ok {
		common.PanicUnimplementedQuery(i.core)
	}
	fc.OlderThan(q)
	return q
}
