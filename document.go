package gosolr

type SolrField struct {
	name  string
	value interface{}
	boost float32
}

func (f *SolrField) Name() string {
	return f.name
}

func (f *SolrField) Value() interface{} {
	return f.value
}

func (f *SolrField) Boost() float32 {
	return f.boost
}

type SolrDocument struct {
	boost  float32
	fields map[string]*SolrField
}

func NewSolrDocument() *SolrDocument {
	return &SolrDocument{
		boost:  float32(1.0),
		fields: make(map[string]*SolrField),
	}
}

func (doc *SolrDocument) Add(name string, value interface{}, boost float32) {
	if found := doc.fields[name]; found == nil || found.value == nil {
		field := &SolrField{name, value, boost}
		doc.fields[name] = field
	} else {
		found.value = value
		found.boost = boost
	}
}

func (s *SolrDocument) Fields() map[string]*SolrField {
	return s.fields
}

func (doc *SolrDocument) Get(name string) *SolrField {
	return doc.fields[name]
}
