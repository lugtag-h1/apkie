package apkie

type component struct {
	Name         string `xml:"name,attr"`
	Exported     string `xml:"exported,attr"`
	Authorities  string `xml:"authorities,attr"`
	IntentFilter []struct {
	} `xml:"intent-filter"`
}

type manifest struct {
	Application struct {
		Activity      []component `xml:"activity"`
		Service       []component `xml:"service"`
		Receiver      []component `xml:"receiver"`
		ActivityAlias []component `xml:"activity-alias"`
		Provider      []component `xml:"provider"`
	} `xml:"application"`
}

type Component interface {
	GetName() string
	IsExported() bool
}

type ComponentInfo struct {
	Name       string
	IsExported bool
}

func (c component) GetName() string {
	return c.Name
}

func (c component) GetAuthorities() string {
	return c.Authorities
}

func (c component) IsExported() bool {
	return c.Exported == "true" ||
		(c.Exported == "" && len(c.IntentFilter) > 0)
}
