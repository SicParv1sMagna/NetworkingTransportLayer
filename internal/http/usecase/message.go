package usecase

import (
	"github.com/SicParv1sMagna/NetworkingTransportLayer/internal/http/constants"
)

// TODO: Refactor
func (uc *UseCase) MessageSegmentation(message string) [][]byte {
	var segments [][]byte

	if len(message) == 0 {
		return segments
	}

	for i := 0; i < len(message); i += constants.SEGMENT_LENGTH {
		end := i + constants.SEGMENT_LENGTH
		if end > len(message) {
			end = len(message)
		}

		segment := make([]byte, end-i)
		copy(segment, []byte(message[i:end]))
		segments = append(segments, segment)
	}

	return segments
}
