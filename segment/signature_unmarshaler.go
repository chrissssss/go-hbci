package segment

import (
	"bytes"
	"fmt"

	"github.com/mitch000001/go-hbci/element"
)

func (s *SignatureHeaderSegment) UnmarshalHBCI(value []byte) error {
	elements, err := ExtractElements(value)
	if err != nil {
		return err
	}
	header := &element.SegmentHeader{}
	err = header.UnmarshalHBCI(elements[0])
	if err != nil {
		return err
	}
	var segment signatureHeaderSegment
	switch header.Version.Val() {
	case 3:
		segment = &SignatureHeaderV3{}
		err = segment.UnmarshalHBCI(value)
		if err != nil {
			return err
		}
	case 4:
		segment = &SignatureHeaderSegmentV4{}
		err = segment.UnmarshalHBCI(value)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unknown segment version: %d", header.Version.Val())
	}
	s.signatureHeaderSegment = segment
	return nil
}

func (s *SignatureHeaderV3) UnmarshalHBCI(value []byte) error {
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
	s.ClientSegment = seg
	if len(elements) > 1 && len(elements[1]) > 0 {
		s.SecurityFunction = &element.AlphaNumericDataElement{}
		err = s.SecurityFunction.UnmarshalHBCI(elements[1])
		if err != nil {
			return err
		}
	}
	if len(elements) > 2 && len(elements[2]) > 0 {
		s.SecurityControlRef = &element.AlphaNumericDataElement{}
		err = s.SecurityControlRef.UnmarshalHBCI(elements[2])
		if err != nil {
			return err
		}
	}
	if len(elements) > 3 && len(elements[3]) > 0 {
		s.SecurityApplicationRange = &element.AlphaNumericDataElement{}
		err = s.SecurityApplicationRange.UnmarshalHBCI(elements[3])
		if err != nil {
			return err
		}
	}
	if len(elements) > 4 && len(elements[4]) > 0 {
		s.SecuritySupplierRole = &element.AlphaNumericDataElement{}
		err = s.SecuritySupplierRole.UnmarshalHBCI(elements[4])
		if err != nil {
			return err
		}
	}
	if len(elements) > 5 && len(elements[5]) > 0 {
		s.SecurityID = &element.SecurityIdentificationDataElement{}
		err = s.SecurityID.UnmarshalHBCI(elements[5])
		if err != nil {
			return err
		}
	}
	if len(elements) > 6 && len(elements[6]) > 0 {
		s.SecurityRefNumber = &element.NumberDataElement{}
		err = s.SecurityRefNumber.UnmarshalHBCI(elements[6])
		if err != nil {
			return err
		}
	}
	if len(elements) > 7 && len(elements[7]) > 0 {
		s.SecurityDate = &element.SecurityDateDataElement{}
		err = s.SecurityDate.UnmarshalHBCI(elements[7])
		if err != nil {
			return err
		}
	}
	if len(elements) > 8 && len(elements[8]) > 0 {
		s.HashAlgorithm = &element.HashAlgorithmDataElement{}
		err = s.HashAlgorithm.UnmarshalHBCI(elements[8])
		if err != nil {
			return err
		}
	}
	if len(elements) > 9 && len(elements[9]) > 0 {
		s.SignatureAlgorithm = &element.SignatureAlgorithmDataElement{}
		err = s.SignatureAlgorithm.UnmarshalHBCI(elements[9])
		if err != nil {
			return err
		}
	}
	if len(elements) > 10 && len(elements[10]) > 0 {
		s.KeyName = &element.KeyNameDataElement{}
		err = s.KeyName.UnmarshalHBCI(elements[10])
		if err != nil {
			return err
		}
	}
	if len(elements) > 11 && len(elements[11]) > 0 {
		s.Certificate = &element.CertificateDataElement{}
		if len(elements)+1 > 11 {
			err = s.Certificate.UnmarshalHBCI(bytes.Join(elements[11:], []byte("+")))
		} else {
			err = s.Certificate.UnmarshalHBCI(elements[11])
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SignatureHeaderSegmentV4) UnmarshalHBCI(value []byte) error {
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
	s.ClientSegment = seg
	if len(elements) > 1 && len(elements[1]) > 0 {
		s.SecurityProfile = &element.SecurityProfileDataElement{}
		err = s.SecurityProfile.UnmarshalHBCI(elements[1])
		if err != nil {
			return err
		}
	}
	if len(elements) > 2 && len(elements[2]) > 0 {
		s.SecurityFunction = &element.CodeDataElement{}
		err = s.SecurityFunction.UnmarshalHBCI(elements[2])
		if err != nil {
			return err
		}
	}
	if len(elements) > 3 && len(elements[3]) > 0 {
		s.SecurityControlRef = &element.AlphaNumericDataElement{}
		err = s.SecurityControlRef.UnmarshalHBCI(elements[3])
		if err != nil {
			return err
		}
	}
	if len(elements) > 4 && len(elements[4]) > 0 {
		s.SecurityApplicationRange = &element.CodeDataElement{}
		err = s.SecurityApplicationRange.UnmarshalHBCI(elements[4])
		if err != nil {
			return err
		}
	}
	if len(elements) > 5 && len(elements[5]) > 0 {
		s.SecuritySupplierRole = &element.CodeDataElement{}
		err = s.SecuritySupplierRole.UnmarshalHBCI(elements[5])
		if err != nil {
			return err
		}
	}
	if len(elements) > 6 && len(elements[6]) > 0 {
		s.SecurityID = &element.SecurityIdentificationDataElement{}
		err = s.SecurityID.UnmarshalHBCI(elements[6])
		if err != nil {
			return err
		}
	}
	if len(elements) > 7 && len(elements[7]) > 0 {
		s.SecurityRefNumber = &element.NumberDataElement{}
		err = s.SecurityRefNumber.UnmarshalHBCI(elements[7])
		if err != nil {
			return err
		}
	}
	if len(elements) > 8 && len(elements[8]) > 0 {
		s.SecurityDate = &element.SecurityDateDataElement{}
		err = s.SecurityDate.UnmarshalHBCI(elements[8])
		if err != nil {
			return err
		}
	}
	if len(elements) > 9 && len(elements[9]) > 0 {
		s.HashAlgorithm = &element.HashAlgorithmDataElement{}
		err = s.HashAlgorithm.UnmarshalHBCI(elements[9])
		if err != nil {
			return err
		}
	}
	if len(elements) > 10 && len(elements[10]) > 0 {
		s.SignatureAlgorithm = &element.SignatureAlgorithmDataElement{}
		err = s.SignatureAlgorithm.UnmarshalHBCI(elements[10])
		if err != nil {
			return err
		}
	}
	if len(elements) > 11 && len(elements[11]) > 0 {
		s.KeyName = &element.KeyNameDataElement{}
		err = s.KeyName.UnmarshalHBCI(elements[11])
		if err != nil {
			return err
		}
	}
	if len(elements) > 12 && len(elements[12]) > 0 {
		s.Certificate = &element.CertificateDataElement{}
		if len(elements)+1 > 12 {
			err = s.Certificate.UnmarshalHBCI(bytes.Join(elements[12:], []byte("+")))
		} else {
			err = s.Certificate.UnmarshalHBCI(elements[12])
		}
		if err != nil {
			return err
		}
	}
	return nil
}
