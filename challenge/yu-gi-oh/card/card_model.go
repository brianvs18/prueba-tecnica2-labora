package card

type Card struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	FrameType   string `json:"frameType"`
	Description string `json:"desc"`
	Atk         int    `json:"atk"`
	Def         int    `json:"def"`
	Level       int    `json:"level"`
	Race        string `json:"race"`
	Attribute   string `json:"attribute"`
}

func (c *Card) IncrementAtk(amount int) {
	c.Atk += amount
}
