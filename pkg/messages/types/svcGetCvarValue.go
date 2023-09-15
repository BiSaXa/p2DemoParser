package messages

import "github.com/pektezol/bitreader"

type SvcGetCvarValue struct {
	Cookie   string
	CvarName string
}

func ParseSvcGetCvarValue(reader *bitreader.Reader) SvcGetCvarValue {
	svcGetCvarValue := SvcGetCvarValue{
		Cookie:   reader.TryReadStringLength(4),
		CvarName: reader.TryReadString(),
	}
	return svcGetCvarValue
}
