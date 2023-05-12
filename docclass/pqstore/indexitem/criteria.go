package indexitem

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-batis/sqlmapper"
)

/*
 * Criteria
 */
type FilterBuilder struct {
	fb *sqlmapper.FilterBuilder
}

func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{fb: &sqlmapper.FilterBuilder{}}
}

func (ub *FilterBuilder) OrderBy(ob string) *FilterBuilder {
	ub.fb.OrderBy(ob)
	return ub
}

func (ub *FilterBuilder) Or() *FilterBuilder {
	ub.fb.Or()
	return ub
}

func (ub *FilterBuilder) Build() sqlmapper.Filter {
	return ub.fb.Build()
}

func (ub *FilterBuilder) AndClassIdEqualTo(aClassId Max30Text) *FilterBuilder {
	ub.fb.And(sqlmapper.Criterion{Type: sqlmapper.SingleValue, Condition: "class_id = ", Value: aClassId})
	return ub
}

func (ub *FilterBuilder) AndNdxIdEqualTo(aNdxId int32) *FilterBuilder {
	ub.fb.And(sqlmapper.Criterion{Type: sqlmapper.SingleValue, Condition: "ndx_id = ", Value: aNdxId})
	return ub
}
