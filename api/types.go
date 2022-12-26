package api

type Set struct {
	ID                 int         `json:"id,omitempty"`
	Timestamp          int64       `json:"timestamp,omitempty"`
	LastModified       int64       `json:"lastModified,omitempty"`
	PublishedTimestamp int64       `json:"publishedTimestamp,omitempty"`
	CreatorID          int         `json:"creatorId,omitempty"`
	WordLang           string      `json:"wordLang,omitempty"`
	DefLang            string      `json:"defLang,omitempty"`
	Title              string      `json:"title,omitempty"`
	PasswordUse        bool        `json:"passwordUse,omitempty"`
	PasswordEdit       bool        `json:"passwordEdit,omitempty"`
	AccessType         int         `json:"accessType,omitempty"`
	AccessCodePrefix   interface{} `json:"accessCodePrefix,omitempty"`
	Description        string      `json:"description,omitempty"`
	NumTerms           int         `json:"numTerms,omitempty"`
	HasImages          bool        `json:"hasImages,omitempty"`
	ParentID           int         `json:"parentId,omitempty"`
	CreationSource     int         `json:"creationSource,omitempty"`
	PrivacyLockStatus  int         `json:"privacyLockStatus,omitempty"`
	HasDiagrams        bool        `json:"hasDiagrams,omitempty"`
	WebURL             string      `json:"_webUrl,omitempty"`
	ThumbnailURL       interface{} `json:"_thumbnailUrl,omitempty"`
	Price              interface{} `json:"price,omitempty"`
	McqCount           int         `json:"mcqCount,omitempty"`
	PurchasableType    int         `json:"purchasableType,omitempty"`
}

type Term struct {
	ID                          int64       `json:"id,omitempty"`
	Word                        string      `json:"word,omitempty"`
	WordTtsURL                  string      `json:"_wordTtsUrl,omitempty"`
	WordSlowTtsURL              string      `json:"_wordSlowTtsUrl,omitempty"`
	WordAudioURL                string      `json:"_wordAudioUrl,omitempty"`
	Definition                  string      `json:"definition,omitempty"`
	DefinitionTtsURL            string      `json:"_definitionTtsUrl,omitempty"`
	DefinitionSlowTtsURL        string      `json:"_definitionSlowTtsUrl,omitempty"`
	DefinitionAudioURL          string      `json:"_definitionAudioUrl,omitempty"`
	ImageURL                    string      `json:"_imageUrl,omitempty"`
	SetID                       int         `json:"setId,omitempty"`
	Rank                        int         `json:"rank,omitempty"`
	LastModified                int         `json:"lastModified,omitempty"`
	WordCustomAudioID           interface{} `json:"wordCustomAudioId,omitempty"`
	DefinitionCustomAudioID     interface{} `json:"definitionCustomAudioId,omitempty"`
	DefinitionImageID           interface{} `json:"definitionImageId,omitempty"`
	DefinitionRichText          interface{} `json:"definitionRichText,omitempty"`
	WordRichText                interface{} `json:"wordRichText,omitempty"`
	DefinitionCustomDistractors interface{} `json:"definitionCustomDistractors,omitempty"`
	WordCustomDistractors       interface{} `json:"wordCustomDistractors,omitempty"`
}

type User struct {
	ID                  int    `json:"id,omitempty"`
	Username            string `json:"username,omitempty"`
	Timestamp           int    `json:"timestamp,omitempty"`
	LastModified        int    `json:"lastModified,omitempty"`
	Type                int    `json:"type,omitempty"`
	IsLocked            bool   `json:"isLocked,omitempty"`
	ImageURL            string `json:"_imageUrl,omitempty"`
	TimeZone            string `json:"timeZone,omitempty"`
	NumClassMemberships int    `json:"_numClassMemberships,omitempty"`
}

type Models struct {
	Set  []Set  `json:"set,omitempty"`
	Term []Term `json:"term,omitempty"`
	User []User `json:"user,omitempty"`
}

type Paging struct {
	Total   int    `json:"total,omitempty"`
	Page    int    `json:"page,omitempty"`
	PerPage int    `json:"perPage,omitempty"`
	Token   string `json:"token,omitempty"`
}

type Response struct {
	Models Models `json:"models,omitempty"`
	Paging Paging `json:"paging,omitempty"`
}

type GenericResponse struct {
	Responses []Response `json:"responses,omitempty"`
}
