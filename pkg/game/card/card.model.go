package card

type CardType string

const (
	Attack CardType = "attack"
	Defense CardType = "defense"
	Effect CardType = "effect"
)

type Card struct {
	ID int 
	Name string 
	Description string 
	Type CardType 
	Cost int
	Rarity int
	Actions []ActionContext
}


type ActionContext struct {
	Target Target 
	Field Field 
	Type ActionType
	Amount int 
	Duration int
}

type ActionType string

const (
	Damage ActionType = "damage"
	Heal ActionType = "heal"
	Curse ActionType = "curse"
	Buff ActionType = "buff"
)

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

func (c *Card) Play (actions []ActionContext) error {
	for _, action := range actions {
		switch action.Type {
		case Damage:
			// remove action.Amount from action.Target's action.Field
		case Heal:
			// add action.Amount to action.Target's action.Field
		case Curse:
			// add a curse to action.Target that deals action.Amount damage for action.Duration turns
		case Buff:
			// add a buff to action.Target that adds action.Amount to action.Field for action.Duration turns
		}
	}
	return nil
}
