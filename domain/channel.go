package domain

type Channel struct {
    Name        string  `json:"name"`
    Link        string  `json:"link"`
    Categories  string  `json:"сategory"`
    Subscribers string  `json:"subscribers"`
}

type Channels []Channel

type channelModel interface {
    findAll() Channels
}
