package builder

import (
	"fmt"
	"log"
	"strings"
)

type SQLQuery struct {
	SelectCols []string
	Table      string
	Joins      []string
	Filters    []string
	GroupBy    []string
	OrderBy    []string
	Limit      int
	Offset     int
}

type SQLBuilder interface {
	Select(cols ...string) SQLBuilder
	From(table string) SQLBuilder
	Join(joinStmt string) SQLBuilder
	Where(cond string) SQLBuilder
	GroupBy(cols ...string) SQLBuilder
	OrderBy(stmt string) SQLBuilder
	Limit(n int) SQLBuilder
	Offset(n int) SQLBuilder
	Build() (string, error)
}

type sqlBuilder struct {
	q   SQLQuery
	err error
}

func NewSQLBuilder() SQLBuilder {
	return &sqlBuilder{q: SQLQuery{Limit: -1}}
}

func (b *sqlBuilder) Select(cols ...string) SQLBuilder {
	if b.err == nil {
		b.q.SelectCols = append(b.q.SelectCols, cols...)
	}
	return b
}

func (b *sqlBuilder) From(table string) SQLBuilder {
	if b.err == nil {
		b.q.Table = table
	}
	return b
}

func (b *sqlBuilder) Join(j string) SQLBuilder {
	if b.err == nil {
		b.q.Joins = append(b.q.Joins, j)
	}
	return b
}

func (b *sqlBuilder) Where(cond string) SQLBuilder {
	if b.err == nil {
		b.q.Filters = append(b.q.Filters, cond)
	}
	return b
}

func (b *sqlBuilder) GroupBy(cols ...string) SQLBuilder {
	if b.err == nil {
		b.q.GroupBy = append(b.q.GroupBy, cols...)
	}
	return b
}

func (b *sqlBuilder) OrderBy(stmt string) SQLBuilder {
	if b.err == nil {
		b.q.OrderBy = append(b.q.OrderBy, stmt)
	}
	return b
}

func (b *sqlBuilder) Limit(n int) SQLBuilder {
	if b.err == nil {
		if n < 0 {
			b.err = fmt.Errorf("limit must be >= 0")
		} else {
			b.q.Limit = n
		}
	}
	return b
}

func (b *sqlBuilder) Offset(n int) SQLBuilder {
	if b.err == nil {
		if n < 0 {
			b.err = fmt.Errorf("offset must be >= 0")
		} else {
			b.q.Offset = n
		}
	}
	return b
}

func (b *sqlBuilder) Build() (string, error) {
	if b.err != nil {
		return "", b.err
	}
	if b.q.Table == "" {
		return "", fmt.Errorf("FROM table not specified")
	}
	parts := []string{"SELECT", strings.Join(b.q.SelectCols, ", "), "FROM", b.q.Table}
	parts = append(parts, b.q.Joins...)
	if len(b.q.Filters) > 0 {
		parts = append(parts, "WHERE "+strings.Join(b.q.Filters, " AND "))
	}
	if len(b.q.GroupBy) > 0 {
		parts = append(parts, "GROUP BY "+strings.Join(b.q.GroupBy, ", "))
	}
	if len(b.q.OrderBy) > 0 {
		parts = append(parts, "ORDER BY "+strings.Join(b.q.OrderBy, ", "))
	}
	if b.q.Limit >= 0 {
		parts = append(parts, fmt.Sprintf("LIMIT %d", b.q.Limit))
	}
	if b.q.Offset > 0 {
		parts = append(parts, fmt.Sprintf("OFFSET %d", b.q.Offset))
	}
	return strings.Join(parts, " "), nil
}

func RunSQLBuilderDemo() {
	sql, err := NewSQLBuilder().
		Select("u.id", "u.name", "SUM(o.amount) AS total").
		From("users u").
		Join("INNER JOIN orders o ON o.user_id = u.id").
		Where("u.active = TRUE").
		GroupBy("u.id", "u.name").
		OrderBy("total DESC").
		Limit(100).
		Offset(0).
		Build()
	if err != nil {
		log.Fatal(err)
	}
	// run sql...
	fmt.Printf("Built SQL Query: %s\n", sql)
}
