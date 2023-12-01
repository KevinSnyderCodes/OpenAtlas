package jsonapi

import "errors"

var ErrDocumentMustContainTopLevelMember = errors.New("a document MUST contain at least one of the following top-level members: data, errors, meta")
var ErrDocumentMustNotContainBothDataAndErrors = errors.New("a document MUST NOT contain both top-level members data and errors")
var ErrDocumentMustNotContainIncludedWithoutData = errors.New("a document MUST NOT contain a top-level included member if the document does not contain a top-level data member")
var ErrDocumentDataInvalidType = errors.New("a document's data member MUST be either: a resource object, an array of resource objects, or null")

var ErrLinkInvalidType = errors.New("a link MUST be represented as either: a string containing the link's URL, a link object, or null")

var ErrLinkObjectMustContainHref = errors.New("a link object MUST contain the following member: href")
var ErrLinkObjectHrefLangInvalidType = errors.New("a link object's hreflang member MUST be a string or an array of strings")

var ErrResourceMustContainID = errors.New("a resource object MUST contain an id member")
var ErrResourceMustContainType = errors.New("a resource object MUST contain a type member")

var ErrRelationshipMustContainLinksDataOrMeta = errors.New("a relationship object MUST contain at least one of the following top-level members: links, data, meta")

var ErrErrorMustContainTopLevelMember = errors.New("an error object MUST contain at least one of the following top-level members: id, links, status, code, title, detail, source, meta")
