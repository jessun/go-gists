package util

/*
todo: function names need to be modified, the usage of maps should be re-considered and remove where 1=1.
*/
import (
	"fmt"
	"strings"
)

type SQLGen struct {
	dataSQL                      strings.Builder
	countSQL                     strings.Builder
	query                        strings.Builder
	args                         []string
	orders                       map[string]sqlOrder
	fuzzyAndKeys, fuzzyAndValues []string
	fuzzyOrKeys, fuzzyOrValues   []string
	andConditionSQLs             []string
	orConditionSQLGroups         map[string] /*groupID*/ *orConditionGroup
	offset,
	limit uint32
}

// orConditionGroup stores a group of or conditions. e.g., AND (a = ? or b = ?) => a and b are in a group
type orConditionGroup struct {
	conditions, args []string
}

func (s *SQLGen) AddOrConditionGroup(groupID, sql string, args ...string) *SQLGen {
	if condGroup, exists := s.orConditionSQLGroups[groupID]; !exists {
		s.orConditionSQLGroups[groupID] = &orConditionGroup{
			conditions: []string{sql},
			args:       args,
		}
	} else {
		condGroup.conditions = append(condGroup.conditions, sql)
		condGroup.args = append(condGroup.args, args...)
	}
	return s
}

func NewSQLGen(dataSQL, countSQL string) *SQLGen {
	sqlG := &SQLGen{
		dataSQL:              strings.Builder{},
		countSQL:             strings.Builder{},
		orders:               make(map[string]sqlOrder),
		orConditionSQLGroups: make(map[string] /*groupID*/ *orConditionGroup),
	}
	sqlG.dataSQL.WriteString(dataSQL)
	sqlG.countSQL.WriteString(countSQL)
	return sqlG
}

type fuzzyLoc int

type sqlOrder string

const (
	DESC sqlOrder = "DESC"
	ASC  sqlOrder = "ASC"
)

const (
	FuzzyPrefix fuzzyLoc = iota // e.g., "%abc"
	FuzzySuffix                 // e.g., "abc%"
	FuzzyAll                    // e.g., "%abc%"
)

func (s *SQLGen) getFuzzyValue(value string, loc fuzzyLoc) string {
	fuzzyValue := value
	switch loc {
	case FuzzyPrefix:
		fuzzyValue = fmt.Sprintf("%%%v", value)
	case FuzzySuffix:
		fuzzyValue = fmt.Sprintf("%v%%", value)
	case FuzzyAll:
		fuzzyValue = fmt.Sprintf("%%%v%%", value)
	}
	return fuzzyValue
}

func (s *SQLGen) AddFuzzyOrCondition(key, value string, loc fuzzyLoc) *SQLGen {
	fuzzyValue := s.getFuzzyValue(value, loc)
	s.fuzzyOrKeys = append(s.fuzzyOrKeys, key)
	s.fuzzyOrValues = append(s.fuzzyOrValues, fuzzyValue)
	return s
}

func (s *SQLGen) AddFuzzyAndCondition(key, value string, loc fuzzyLoc) *SQLGen {
	fuzzyValue := s.getFuzzyValue(value, loc)
	s.fuzzyAndKeys = append(s.fuzzyAndKeys, key)
	s.fuzzyAndValues = append(s.fuzzyAndValues, fuzzyValue)
	return s
}

func (s *SQLGen) AddWhereCondition(key, op string, value string) *SQLGen {
	s.query.WriteString(fmt.Sprintf(" AND %v %v ? ", key, op))
	s.args = append(s.args, value)
	return s
}

func (s *SQLGen) AddInCondition(key string, values ...string) *SQLGen {
	if len(values) == 0 {
		return s
	}
	placeHolders := make([]string, len(values))
	for i := range values {
		placeHolders[i] = "?"
	}
	s.query.WriteString(fmt.Sprintf(" AND %v in (%s)", key, strings.Join(placeHolders, ",")))
	s.args = append(s.args, values...)
	return s
}

func (s *SQLGen) OrderBy(key, order string) *SQLGen {
	s.orders[key] = sqlOrder(order)
	return s
}

func (s *SQLGen) OffSet(offset int) *SQLGen {
	s.offset = uint32(offset)
	return s
}

func (s *SQLGen) Limit(limit int) *SQLGen {
	s.limit = uint32(limit)
	return s
}

func (s *SQLGen) AddConditionSQL(sql string) *SQLGen {
	s.andConditionSQLs = append(s.andConditionSQLs, sql)
	return s
}

func (s *SQLGen) CountArgs() []string {
	if s.offset != 0 { // if there is an offset, then there must also be a limit
		if len(s.args) <= 2 {
			return nil
		}
		return s.args[:len(s.args)-2]
	} else {
		if s.limit != 0 {
			if len(s.args) <= 1 {
				return nil
			}
			return s.args[:len(s.args)-1]
		}
	}
	return s.args
}

func (s *SQLGen) Args() []string {
	return s.args
}

func (s *SQLGen) GetInterfaceArgs() []interface{} {
	args := make([]interface{}, 0, len(s.args))
	for i := range s.args {
		args = append(args, s.args[i])
	}
	return args
}

func (s *SQLGen) GetInterfaceCountArgs() []interface{} {
	countArgs := s.CountArgs()
	args := make([]interface{}, 0, len(countArgs))
	for _, arg := range countArgs {
		args = append(args, arg)
	}
	return args
}

func (s *SQLGen) GenerateSQL() (countSql string /*countSql*/, querySql string /*querySql*/) {
	var orderBys []string

	fuzzyQueryConcat := func(keys, values []string, op string) string {
		var fuzzyConditions []string
		for i := range keys {
			fuzzyConditions = append(fuzzyConditions, keys[i]+" LIKE "+" ? ")
			s.args = append(s.args, values[i])
		}
		return fmt.Sprintf(" AND (%s)", strings.Join(fuzzyConditions, op))
	}

	if len(s.fuzzyOrKeys) != 0 {
		s.query.WriteString(fuzzyQueryConcat(s.fuzzyOrKeys, s.fuzzyOrValues, " OR "))
	}

	if len(s.fuzzyAndKeys) != 0 {
		s.query.WriteString(fuzzyQueryConcat(s.fuzzyAndKeys, s.fuzzyAndValues, " AND "))
	}

	for _, sql := range s.andConditionSQLs {
		s.query.WriteString(" AND ")
		s.query.WriteString(sql)
	}
	for _, condGroup := range s.orConditionSQLGroups {
		longOrStr := strings.Join(condGroup.conditions, " OR ")
		s.query.WriteString(fmt.Sprintf(" AND ( %s ) ", longOrStr))
		s.args = append(s.args, condGroup.args...)
	}

	queryWithWhere := s.query.String()
	if len(s.orders) != 0 {
		orderBys = make([]string, 0, len(s.orders))
		for k, v := range s.orders {
			orderBys = append(orderBys, fmt.Sprintf("%s %s", k, v))
		}
		orderByStr := strings.Join(orderBys, ",")
		s.query.WriteString(" ORDER BY ")
		s.query.WriteString(orderByStr)
	}

	if s.limit != 0 {
		if s.offset != 0 {
			s.query.WriteString(" limit ?,?")
			s.args = append(s.args, fmt.Sprintf("%v", s.offset), fmt.Sprintf("%v", s.limit))
		} else {
			s.query.WriteString(" limit ?")
			s.args = append(s.args, fmt.Sprintf("%v", s.limit))
		}
	}
	return s.countSQL.String() + queryWithWhere, s.dataSQL.String() + s.query.String()

}
