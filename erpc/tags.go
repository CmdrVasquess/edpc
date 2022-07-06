package erpc

import (
	"reflect"

	"github.com/fxamacker/cbor/v2"
)

func CBORTags() (tags cbor.TagSet, err error) {
	tagModes := cbor.TagOptions{
		EncTag: cbor.EncTagRequired,
		DecTag: cbor.DecTagRequired,
	}
	tags = cbor.NewTagSet()
	if err = tags.Add(tagModes, reflect.TypeOf(ResponseHeader{}), redirectTag); err != nil {
		return nil, err
	}
	if err = tags.Add(tagModes, reflect.TypeOf(CommanderEvent{}), commanderEventTag); err != nil {
		return nil, err
	}
	if err = tags.Add(tagModes, reflect.TypeOf(DockedEvent{}), dockedEventTag); err != nil {
		return nil, err
	}
	return tags, nil
}

const (
	redirectTag uint64 = (iota + 55800)

	commanderEventTag
	dockedEventTag
)
