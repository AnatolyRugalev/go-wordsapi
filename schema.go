package wordsapi

type WordReference string

const WordAll = WordReference("all")

type PartOfSpeech string

const (
	PoSNoun      = PartOfSpeech("noun")
	PoSVerb      = PartOfSpeech("verb")
	PoSAdjective = PartOfSpeech("adjective")
	PoSAdverb    = PartOfSpeech("adverb")
)

type WordResponse struct {
	Word          string                   `json:"word"`
	Results       []Word                   `json:"results"`
	Syllables     Syllables                `json:"syllables"`
	Pronunciation map[WordReference]string `json:"pronunciation"`
	Frequency     float64                  `json:"frequency"`
}

type WordDefinition struct {
	// Definition the meaning of the word
	Definition string `json:"definition"`
	// PartOfSpeech part of speech of the word
	PartOfSpeech PartOfSpeech `json:"partOfSpeech"`
}

// TODO: split fields into small structs for response reusability
// TODO: document every field and struct
type Word struct {
	WordDefinition
	// Synonims words that can be interchanged for the original word in the same context
	Synonims      []WordReference `json:"synonims"`
	TypeOf        []WordReference `json:"typeOf"`
	HasTypes      []WordReference `json:"hasTypes"`
	PartOf        []WordReference `json:"partOf"`
	HasParts      []WordReference `json:"hasParts"`
	InstanceOf    []WordReference `json:"instanceOf"`
	HasInstances  []WordReference `json:"hasInstances"`
	SimilarTo     []WordReference `json:"similarTo"`
	Also          []WordReference `json:"also"`
	Entails       []WordReference `json:"entails"`
	MemberOf      []WordReference `json:"memberOf"`
	HasMembers    []WordReference `json:"hasMembers"`
	SubstanceOf   []WordReference `json:"substanceOf"`
	HasSubstances []WordReference `json:"hasSubstances"`
	InCategory    []WordReference `json:"inCategory"`
	HasCategories []WordReference `json:"hasCategories"`
	UsageOf       []WordReference `json:"usageOf"`
	HasUsages     []WordReference `json:"hasUsages"`
	InRegion      []WordReference `json:"inRegion"`
	RegionOf      []WordReference `json:"regionOf"`
	PetrainsTo    []WordReference `json:"petrainsTo"`
	Derivation    []WordReference `json:"derivation"`
	Examples      []string        `json:"examples"`
}

type Syllables struct {
	Count int      `json:"count"`
	List  []string `json:"list"`
}
