package iot

import (
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/uses/common"
	"github.com/timescale/tsbs/cmd/tsbs_generate_queries/utils"
	"github.com/timescale/tsbs/pkg/query"
)

// Extract1Tag contains info for filling in extract 1 tag queries.
type Extract1Tag struct {
	core utils.QueryGenerator
}

// NewExtract1Tag creates a new extract 1 tag query filler.
func NewExtract1TagH(core utils.QueryGenerator) utils.QueryFiller {
	return &Extract1Tag{
		core: core,
	}
}

// Fill fills in the query.Query with query details.
func (i *Extract1Tag) Fill(q query.Query) query.Query {
	fc, ok := i.core.(Extract1TagFiller)
	if !ok {
		common.PanicUnimplementedQuery(i.core)
	}
	fc.Extract1Tag(q)
	return q
}
