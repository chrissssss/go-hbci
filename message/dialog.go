package message

import "github.com/mitch000001/go-hbci/segment"

func NewDialogInitializationClientMessage() *DialogInitializationClientMessage {
	d := &DialogInitializationClientMessage{}
	d.BasicMessage = NewBasicMessage(d)
	return d
}

type DialogInitializationClientMessage struct {
	*BasicMessage
	Identification             *segment.IdentificationSegment
	ProcessingPreparation      *segment.ProcessingPreparationSegment
	PublicSigningKeyRequest    *segment.PublicKeyRequestSegment
	PublicEncryptionKeyRequest *segment.PublicKeyRequestSegment
	PublicKeyRequest           *segment.PublicKeyRequestSegment
}

func (d *DialogInitializationClientMessage) HBCISegments() []segment.ClientSegment {
	return []segment.ClientSegment{
		d.Identification,
		d.ProcessingPreparation,
		d.PublicSigningKeyRequest,
		d.PublicEncryptionKeyRequest,
		d.PublicKeyRequest,
	}
}

func (d *DialogInitializationClientMessage) jobs() []segment.Segment {
	return []segment.Segment{
		d.Identification,
		d.ProcessingPreparation,
		d.PublicSigningKeyRequest,
		d.PublicEncryptionKeyRequest,
	}
}

type DialogFinishingMessage struct {
	*BasicMessage
	DialogEnd *segment.DialogEndSegment
}

func (d *DialogFinishingMessage) HBCISegments() []segment.ClientSegment {
	return []segment.ClientSegment{
		d.DialogEnd,
	}
}

func (d *DialogFinishingMessage) jobs() []segment.ClientSegment {
	return []segment.ClientSegment{
		d.DialogEnd,
	}
}

func NewDialogCancellationMessage(messageAcknowledgement *segment.MessageAcknowledgement) *DialogCancellationMessage {
	d := &DialogCancellationMessage{
		MessageAcknowledgements: messageAcknowledgement,
	}
	return d
}

type DialogCancellationMessage struct {
	*BasicMessage
	MessageAcknowledgements *segment.MessageAcknowledgement
}

type AnonymousDialogMessage struct {
	*BasicMessage
	Identification        *segment.IdentificationSegment
	ProcessingPreparation *segment.ProcessingPreparationSegment
}
