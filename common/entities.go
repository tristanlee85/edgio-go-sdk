package common

import "time"

type Searchable interface {
	GetKey() string
	GetName() string
	GetSlug() string
}

type Filterable interface {
	Searchable
	Property | Env | Variable
}

type FilterListParams[T Filterable] struct {
	Needle   string
	Haystack []T
}

type Org struct {
	Searchable

	ID   string `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

type Variable struct {
	Searchable

	ID        string
	Key       string
	Value     string
	Secret    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetKey returns the Key value of the Variable.
func (v Variable) GetKey() string {
	return v.Key
}

// GetName returns the Name value of the Variable
// It returns an empty string because the Property struct does not have a Name field.
func (v Variable) GetName() string {
	return ""
}

// GetSlug returns the Slug value of the Variable
// It returns an empty string because the Property struct does not have a Slug field.
func (v Variable) GetSlug() string {
	return ""
}

type Property struct {
	Searchable

	ID        string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetKey returns the Key value of the Property
// It returns an empty string because the Property struct does not have a Key field.
func (p Property) GetKey() string {
	return ""
}

// GetName returns the Name value of the Property
// It returns an empty string because the Property struct does not have a Name field.
func (p Property) GetName() string {
	return ""
}

// GetSlug returns the Slug value of the Property.
func (p Property) GetSlug() string {
	return p.Slug
}

type Env struct {
	Searchable

	ID                       string
	Name                     string
	LegacyAccNumber          string
	DefaultDomainName        string
	DNSDomainName            string
	CanMembersDeploy         bool
	OnlyMaintainersCanDeploy bool
	HTTPRequestLogging       bool
	PciCompliance            bool
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

// GetKey returns the Key value of the Env
// It returns an empty string because the Env struct does not have a Key field.
func (e Env) GetKey() string {
	return ""
}

// GetName returns the Name value of the Env.
func (e Env) GetName() string {
	return e.Name
}

// GetSlug returns the Slug value of the Env
// It returns an empty string because the Env struct does not have a Slug field.
func (e Env) GetSlug() string {
	return ""
}
