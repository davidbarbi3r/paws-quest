package card

type CardType string

const (
	Attack CardType = "attack"
	Defense CardType = "defense"
	Heal CardType = "heal"
	Curse CardType = "curse"
	Buff CardType = "buff"
)

type Card struct {
	ID int 
	Name string 
	Description string 
	Type CardType 
	Cost int
	Rarity int
	Actions []ActionFunc
}

type ActionFunc func(ActionContext) error

type ActionContext struct {
	Target Target 
	Field Field 
	Amount int 
	Duration int
	Card *Card
}

type Field int

const (
	Health Field = iota
	Stamina
	Speed
	Strength
	Agility
	Intelligence
	HandSize
)

type Target interface {}

type Rarity int

const (
	Common Rarity = iota
	Uncommon
	Rare
)

func (c *Card) Play(ctx ActionContext) error {
	for _, action := range c.Actions {
		err := action(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
