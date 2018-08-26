package models

import (
	"github.com/shkatovdm/domain"
)

type postgresChannelModel struct {}

func (pchm *postgresChannelModel) findAll() domain.Channels {
	var channels domain.Channels
	var channel domain.Channel

	channels[1] = channel
	return channels
}
