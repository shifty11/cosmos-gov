// Code generated by entc, DO NOT EDIT.

package proposal

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/shifty11/cosmos-gov/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdatedTime applies equality check predicate on the "updated_time" field. It's identical to UpdatedTimeEQ.
func UpdatedTime(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedTime), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ProposalID applies equality check predicate on the "proposal_id" field. It's identical to ProposalIDEQ.
func ProposalID(v uint64) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProposalID), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// VotingStartTime applies equality check predicate on the "voting_start_time" field. It's identical to VotingStartTimeEQ.
func VotingStartTime(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVotingStartTime), v))
	})
}

// VotingEndTime applies equality check predicate on the "voting_end_time" field. It's identical to VotingEndTimeEQ.
func VotingEndTime(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVotingEndTime), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateTime)))
	})
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateTime)))
	})
}

// UpdatedTimeEQ applies the EQ predicate on the "updated_time" field.
func UpdatedTimeEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedTime), v))
	})
}

// UpdatedTimeNEQ applies the NEQ predicate on the "updated_time" field.
func UpdatedTimeNEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedTime), v))
	})
}

// UpdatedTimeIn applies the In predicate on the "updated_time" field.
func UpdatedTimeIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedTime), v...))
	})
}

// UpdatedTimeNotIn applies the NotIn predicate on the "updated_time" field.
func UpdatedTimeNotIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedTime), v...))
	})
}

// UpdatedTimeGT applies the GT predicate on the "updated_time" field.
func UpdatedTimeGT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedTime), v))
	})
}

// UpdatedTimeGTE applies the GTE predicate on the "updated_time" field.
func UpdatedTimeGTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedTime), v))
	})
}

// UpdatedTimeLT applies the LT predicate on the "updated_time" field.
func UpdatedTimeLT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedTime), v))
	})
}

// UpdatedTimeLTE applies the LTE predicate on the "updated_time" field.
func UpdatedTimeLTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedTime), v))
	})
}

// UpdatedTimeIsNil applies the IsNil predicate on the "updated_time" field.
func UpdatedTimeIsNil() predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedTime)))
	})
}

// UpdatedTimeNotNil applies the NotNil predicate on the "updated_time" field.
func UpdatedTimeNotNil() predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedTime)))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// ProposalIDEQ applies the EQ predicate on the "proposal_id" field.
func ProposalIDEQ(v uint64) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProposalID), v))
	})
}

// ProposalIDNEQ applies the NEQ predicate on the "proposal_id" field.
func ProposalIDNEQ(v uint64) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProposalID), v))
	})
}

// ProposalIDIn applies the In predicate on the "proposal_id" field.
func ProposalIDIn(vs ...uint64) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProposalID), v...))
	})
}

// ProposalIDNotIn applies the NotIn predicate on the "proposal_id" field.
func ProposalIDNotIn(vs ...uint64) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProposalID), v...))
	})
}

// ProposalIDGT applies the GT predicate on the "proposal_id" field.
func ProposalIDGT(v uint64) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProposalID), v))
	})
}

// ProposalIDGTE applies the GTE predicate on the "proposal_id" field.
func ProposalIDGTE(v uint64) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProposalID), v))
	})
}

// ProposalIDLT applies the LT predicate on the "proposal_id" field.
func ProposalIDLT(v uint64) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProposalID), v))
	})
}

// ProposalIDLTE applies the LTE predicate on the "proposal_id" field.
func ProposalIDLTE(v uint64) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProposalID), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// VotingStartTimeEQ applies the EQ predicate on the "voting_start_time" field.
func VotingStartTimeEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVotingStartTime), v))
	})
}

// VotingStartTimeNEQ applies the NEQ predicate on the "voting_start_time" field.
func VotingStartTimeNEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVotingStartTime), v))
	})
}

// VotingStartTimeIn applies the In predicate on the "voting_start_time" field.
func VotingStartTimeIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVotingStartTime), v...))
	})
}

// VotingStartTimeNotIn applies the NotIn predicate on the "voting_start_time" field.
func VotingStartTimeNotIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVotingStartTime), v...))
	})
}

// VotingStartTimeGT applies the GT predicate on the "voting_start_time" field.
func VotingStartTimeGT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVotingStartTime), v))
	})
}

// VotingStartTimeGTE applies the GTE predicate on the "voting_start_time" field.
func VotingStartTimeGTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVotingStartTime), v))
	})
}

// VotingStartTimeLT applies the LT predicate on the "voting_start_time" field.
func VotingStartTimeLT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVotingStartTime), v))
	})
}

// VotingStartTimeLTE applies the LTE predicate on the "voting_start_time" field.
func VotingStartTimeLTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVotingStartTime), v))
	})
}

// VotingEndTimeEQ applies the EQ predicate on the "voting_end_time" field.
func VotingEndTimeEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVotingEndTime), v))
	})
}

// VotingEndTimeNEQ applies the NEQ predicate on the "voting_end_time" field.
func VotingEndTimeNEQ(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVotingEndTime), v))
	})
}

// VotingEndTimeIn applies the In predicate on the "voting_end_time" field.
func VotingEndTimeIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVotingEndTime), v...))
	})
}

// VotingEndTimeNotIn applies the NotIn predicate on the "voting_end_time" field.
func VotingEndTimeNotIn(vs ...time.Time) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVotingEndTime), v...))
	})
}

// VotingEndTimeGT applies the GT predicate on the "voting_end_time" field.
func VotingEndTimeGT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVotingEndTime), v))
	})
}

// VotingEndTimeGTE applies the GTE predicate on the "voting_end_time" field.
func VotingEndTimeGTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVotingEndTime), v))
	})
}

// VotingEndTimeLT applies the LT predicate on the "voting_end_time" field.
func VotingEndTimeLT(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVotingEndTime), v))
	})
}

// VotingEndTimeLTE applies the LTE predicate on the "voting_end_time" field.
func VotingEndTimeLTE(v time.Time) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVotingEndTime), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Proposal {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Proposal(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// HasChain applies the HasEdge predicate on the "chain" edge.
func HasChain() predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChainTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChainTable, ChainColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChainWith applies the HasEdge predicate on the "chain" edge with a given conditions (other predicates).
func HasChainWith(preds ...predicate.Chain) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChainInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ChainTable, ChainColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Proposal) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Proposal) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Proposal) predicate.Proposal {
	return predicate.Proposal(func(s *sql.Selector) {
		p(s.Not())
	})
}
