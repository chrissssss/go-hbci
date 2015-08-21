package segment

import (
	"strconv"

	"github.com/mitch000001/go-hbci/element"
)

func NewSynchronisationSegmentV2(modus int) *SynchronisationRequestSegment {
	s := &SynchronisationRequestV2{
		SyncModus: element.NewNumber(modus, 1),
	}
	s.ClientSegment = NewBasicSegment(5, s)

	segment := &SynchronisationRequestSegment{
		ClientSegment: s,
	}
	return segment
}

type SynchronisationRequestSegment struct {
	ClientSegment
}

type SynchronisationRequestV2 struct {
	ClientSegment
	// Code | Bedeutung
	// ---------------------------------------------------------
	// 0 ￼ ￼| Neue Kundensystem-ID zurückmelden
	// 1	| Letzte verarbeitete Nachrichtennummer zurückmelden ￼ ￼
	// 2 ￼ ￼| Signatur-ID zurückmelden
	SyncModus *element.NumberDataElement
}

func (s *SynchronisationRequestV2) Version() int         { return 2 }
func (s *SynchronisationRequestV2) ID() string           { return "HKSYN" }
func (s *SynchronisationRequestV2) referencedId() string { return "" }
func (s *SynchronisationRequestV2) sender() string       { return senderUser }

func (s *SynchronisationRequestV2) elements() []element.DataElement {
	return []element.DataElement{
		s.SyncModus,
	}
}

func NewSynchronisationSegmentV3(modus int) *SynchronisationRequestSegment {
	s := &SynchronisationRequestV3{
		SyncModus: element.NewCode(strconv.Itoa(modus), 1, []string{"0", "1", "2"}),
	}
	s.ClientSegment = NewBasicSegment(5, s)

	segment := &SynchronisationRequestSegment{
		ClientSegment: s,
	}
	return segment
}

type SynchronisationRequestV3 struct {
	ClientSegment
	// Code | Bedeutung
	// ---------------------------------------------------------
	// 0 ￼ ￼| Neue Kundensystem-ID zurückmelden
	// 1	| Letzte verarbeitete Nachrichtennummer zurückmelden ￼ ￼
	// 2 ￼ ￼| Signatur-ID zurückmelden
	SyncModus *element.CodeDataElement
}

func (s *SynchronisationRequestV3) Version() int         { return 3 }
func (s *SynchronisationRequestV3) ID() string           { return "HKSYN" }
func (s *SynchronisationRequestV3) referencedId() string { return "" }
func (s *SynchronisationRequestV3) sender() string       { return senderUser }

func (s *SynchronisationRequestV3) elements() []element.DataElement {
	return []element.DataElement{
		s.SyncModus,
	}
}

//go:generate go run ../cmd/unmarshaler/unmarshaler_generator.go -segment SynchronisationResponseSegment

type SynchronisationResponseSegment struct {
	Segment
	ClientSystemID *element.IdentificationDataElement
	MessageNumber  *element.NumberDataElement
	SignatureID    *element.NumberDataElement
}

func (s *SynchronisationResponseSegment) Version() int         { return 3 }
func (s *SynchronisationResponseSegment) ID() string           { return "HISYN" }
func (s *SynchronisationResponseSegment) referencedId() string { return "HKSYN" }
func (s *SynchronisationResponseSegment) sender() string       { return senderBank }

func (s *SynchronisationResponseSegment) elements() []element.DataElement {
	return []element.DataElement{
		s.ClientSystemID,
		s.MessageNumber,
		s.SignatureID,
	}
}
