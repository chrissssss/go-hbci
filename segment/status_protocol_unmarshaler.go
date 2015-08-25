package segment

import (
	"bytes"
	"fmt"

	"github.com/mitch000001/go-hbci/element"
)

func (s *StatusProtocolResponseSegment) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	header := &element.SegmentHeader{}
	err = header.UnmarshalHBCI(elements[0])
	if err != nil {
		return err
	}
	var segment StatusProtocolResponse
	switch header.Version.Val() {
	case 3:
		segment = &StatusProtocolResponseSegmentV3{}
		err = segment.UnmarshalHBCI(value)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown segment version: %d", header.Version.Val())
	}
	s.StatusProtocolResponse = segment
	return nil
}

func (s *StatusProtocolResponseSegmentV3) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	if len(elements) == 0 {
		return fmt.Errorf("Malformed marshaled value")
	}
	seg, err := SegmentFromHeaderBytes(elements[0], s)
	if err != nil {
		return err
	}
	s.Segment = seg
	if len(elements) > 1 && len(elements[1]) > 0 {
		s.ReferencingMessage = &element.ReferencingMessageDataElement{}
		err = s.ReferencingMessage.UnmarshalHBCI(elements[1])
		if err != nil {
			return err
		}
	}
	if len(elements) > 2 && len(elements[2]) > 0 {
		s.ReferencingSegment = &element.NumberDataElement{}
		err = s.ReferencingSegment.UnmarshalHBCI(elements[2])
		if err != nil {
			return err
		}
	}
	if len(elements) > 3 && len(elements[3]) > 0 {
		s.Date = &element.DateDataElement{}
		err = s.Date.UnmarshalHBCI(elements[3])
		if err != nil {
			return err
		}
	}
	if len(elements) > 4 && len(elements[4]) > 0 {
		s.Time = &element.TimeDataElement{}
		err = s.Time.UnmarshalHBCI(elements[4])
		if err != nil {
			return err
		}
	}
	if len(elements) > 5 && len(elements[5]) > 0 {
		s.Acknowledgement = &element.AcknowledgementDataElement{}
		if len(elements)+1 > 5 {
			err = s.Acknowledgement.UnmarshalHBCI(bytes.Join(elements[5:], []byte("+")))
		} else {
			err = s.Acknowledgement.UnmarshalHBCI(elements[5])
		}
		if err != nil {
			return err
		}
	}
	return nil
}
