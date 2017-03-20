package main

type APIModel struct {
	Version    string               `json:"version"`
	Metadata   Metadata             `json:"metadata"`
	Operations map[string]Operation `json:"operations"`
	Shapes     map[string]Shape     `json:"shapes"`
}

type Metadata struct {
	APIVersion          string `json:"apiVersion"`
	EndpointPrefix      string `json:"endpointPrefix"`
	GlobalEndpoint      string `json:"globalEndpoint,omitempty"`
	ServiceAbbreviation string `json:"serviceAbbreviation"`
	ServiceFullName     string `json:"serviceFullName"`
	SignatureVersion    string `json:"signatureVersion"`
	XMLNamespace        string `json:"xmlNamespace"`
	Protocol            string `json:"protocol"`
}

type Operation struct {
	OpName string    `json:"name"`
	HTTP   HTTP      `json:"http"`
	Input  *ShapeRef `json:"input"`
	Output *ShapeRef `json:"output,omitempty"`
}

type HTTP struct {
	Method     string `json:"method"`
	RequestURI string `json:"requestUri"`
}

type ShapeRef struct {
	ShapeName    string `json:"shape"`
	Location     string `json:"location,omitempty"`
	LocationName string `json:"locationName,omitempty"`
}

type Shape struct {
	ShapeName string              `json:"-"`
	Type      string              `json:"type"`
	Member    *ShapeRef           `json:"member,omitempty"`   // for type=list
	Members   map[string]ShapeRef `json:"members,omitempty"`  // for type=structure
	Required  []string            `json:"required,omitempty"` // for type=structure
	Enum      []string            `json:"enum,omitempty"`     // for type=string
	Key       *ShapeRef           `json:"key,omitempty"`      // for type=map
	Value     *ShapeRef           `json:"value,omitempty"`    // for type=value
}

type DocsModel struct {
	Version    string               `json:"version"`
	Service    string               `json:"service"`
	Operations map[string]string    `json:"operations"`
	Shapes     map[string]DocsShape `json:"shapes"`
}

type DocsShape struct {
	Base string            `json:"base"`
	Refs map[string]string `json:"refs"`
}

type Shapes []Shape

// get position in strint slice case-insensitively
func (shapes Shapes) findShapeByName(shapeName string) *Shape {
	for _, shape := range shapes {
		if shape.ShapeName == shapeName {
			return &shape
		}
	}
	return &Shape{}
}
