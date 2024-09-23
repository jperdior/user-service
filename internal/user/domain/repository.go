package domain

//go:generate mockery --case=snake --outpkg=domainmocks --output=domainmocks --name=UserRepository

type UserRepository interface {
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	Find(criteria Criteria) ([]*User, int64, error)
	Save(user *User) error
}

type Criteria struct {
	Filters  []Filter
	Sort     string
	SortDir  string
	Page     int
	PageSize int
}

func NewCriteria(filters map[string]interface{}, sort, sortDir string, page, pageSize int) Criteria {
	var criteriaFilters []Filter
	for key, value := range filters {
		if filter, ok := ValidFilters[key]; ok {
			filter.Value = value
			criteriaFilters = append(criteriaFilters, filter)
		}
	}
	return Criteria{
		Filters:  criteriaFilters,
		Sort:     sort,
		SortDir:  sortDir,
		Page:     page,
		PageSize: pageSize,
	}
}

type Filter struct {
	Name      string
	Type      string
	Operation string
	Value     interface{}
}

var ValidFilters = map[string]Filter{
	"id": {
		Name:      "id",
		Type:      "string",
		Operation: "=",
	},
	"name": {
		Name:      "name",
		Type:      "string",
		Operation: "LIKE",
	},
	"email": {
		Name:      "email",
		Type:      "string",
		Operation: "LIKE",
	},
	"role": {
		Name:      "role",
		Type:      "string",
		Operation: "=",
	},
}
