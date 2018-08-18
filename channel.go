package tlgstat

type Channel struct {
    Name         string  `json:"name"`
    Link         string  `json:"link"`
    Subscribers  uint    `json:"сategory"`
}

type Channels []Channel