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
	Actions []ActionFunc
}

type ActionFunc func(ActionContext) error

type ActionContext struct {
	Target string 
	Effect string 
	Amount int 
	Duration int
}

var exampleAttackCard = Card {
	ID: 1,
	Name: "Paw Swipe",
	Description: "A basic attack card",
	Type: Attack,
	Actions: []ActionFunc{
		func (ctx ActionContext) error {
			return nil
		},
	},
}
