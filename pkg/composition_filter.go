package flocksofblocks

type Filter struct {
	include                       []string
	excludeExternalSharedChannels bool
	excludeBotUsers               bool

	optionals filterOptions
}

type filterOptions struct {
	Include                       bool
	ExcludeExternalSharedChannels bool
	ExcludeBotUsers               bool
}

func NewFilter() Filter {
	return Filter{
		include:                       []string{},
		excludeExternalSharedChannels: false,
		excludeBotUsers:               false,
		optionals: filterOptions{
			Include:                       false,
			ExcludeExternalSharedChannels: false,
			ExcludeBotUsers:               false,
		},
	}
}

// abstracted type
type filterAbstraction struct {
	Include                       []string
	ExcludeExternalSharedChannels bool
	ExcludeBotUsers               bool

	Optionals filterOptions
}

// add include to filter
func (f *Filter) addInclude(include string) {
	f.include = append(f.include, include)
	f.optionals.Include = true
}

// AddInclude add filter string
func (f Filter) AddInclude(include string) Filter {
	f.addInclude(include)
	return f
}

// IncludeIM Add IM to include filter
func (f Filter) IncludeIM() Filter {
	return f.AddInclude("im")
}

// IncludeMPIM Add MPIM to include filter
func (f Filter) IncludeMPIM() Filter {
	return f.AddInclude("mpim")
}

// IncludePrivate Add private to include filter
func (f Filter) IncludePrivate() Filter {
	return f.AddInclude("private")
}

// IncludePublic Add public to include filter
func (f Filter) IncludePublic() Filter {
	return f.AddInclude("public")
}

// ClearInclude clear include
func (f *Filter) clearInclude() {
	f.include = []string{}
	f.optionals.Include = false
}

// ClearInclude clear include
func (f Filter) ClearInclude() Filter {
	f.clearInclude()
	return f
}

// set exclude external shared channels
func (f *Filter) setExcludeExternalSharedChannels(excludeExternalSharedChannels bool) {
	f.excludeExternalSharedChannels = excludeExternalSharedChannels
	f.optionals.ExcludeExternalSharedChannels = excludeExternalSharedChannels
}

// ExcludeExternalSharedChannels exclude external shared channels
func (f Filter) ExcludeExternalSharedChannels() Filter {
	f.setExcludeExternalSharedChannels(true)
	return f
}

// UnsetExcludeExternalSharedChannels unset exclude external shared channels
func (f Filter) UnsetExcludeExternalSharedChannels() Filter {
	f.setExcludeExternalSharedChannels(false)
	return f
}

// set exclude bot users
func (f *Filter) setExcludeBotUsers(excludeBotUsers bool) {
	f.excludeBotUsers = excludeBotUsers
	f.optionals.ExcludeBotUsers = excludeBotUsers
}

// ExcludeBotUsers exclude bot users
func (f Filter) ExcludeBotUsers() Filter {
	f.setExcludeBotUsers(true)
	return f
}

// UnsetExcludeBotUsers unset exclude bot users
func (f Filter) UnsetExcludeBotUsers() Filter {
	f.setExcludeBotUsers(false)
	return f
}

// abstraction
func (f Filter) abstraction() filterAbstraction {
	return filterAbstraction{
		Include:                       removeDuplicateString(f.include),
		ExcludeExternalSharedChannels: f.excludeExternalSharedChannels,
		ExcludeBotUsers:               f.excludeBotUsers,
		Optionals:                     f.optionals,
	}
}

func (f filterAbstraction) empty() bool {
	if f.Optionals.Include {
		return false
	}
	if f.Optionals.ExcludeExternalSharedChannels {
		return false
	}
	if f.Optionals.ExcludeBotUsers {
		return false
	}
	return true
}

// template
func (f filterAbstraction) Template() string {
	if f.empty() {
		return ""
	}
	return `"filter": {
	{{if .Optionals.Include}}"include": [{{range $index, $include := .Include}}{{if $index}},{{end}}"{{ $include}}"{{end}}]{{if .Optionals.ExcludeExternalSharedChannels}},{{end}}{{end}}{{if .Optionals.ExcludeExternalSharedChannels}}
	"exclude_external_shared_channels": {{.ExcludeExternalSharedChannels}}{{if .Optionals.ExcludeBotUsers }},{{end}}{{end}}{{if .Optionals.ExcludeBotUsers }}
	"exclude_bot_users": {{.ExcludeBotUsers}}{{end}}
}`
}

// Render method
func (f Filter) Render() string {
	return Render(f.abstraction())
}

// im, mpim, private, and public.
