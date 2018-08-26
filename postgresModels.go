package tlgstat

type postgresChannelModel struct {}

func (pchm *postgresChannelModel) findAll() Channels {
	var channels Channels
	var channel Channel

	channels[1] = channel
	return channels
}

