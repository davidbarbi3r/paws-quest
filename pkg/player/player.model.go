package player

type Player struct {
	ID int 
	Name string 
	Email string 
	Password string 
	Level int 
	Experience int 
}

// type PlayerService interface {
// 	GetAll() ([]Player, error)
// 	GetByID(id int) (Player, error)
// 	Create(player Player) (Player, error)
// 	Update(player Player) (Player, error)
// 	Delete(id int) (Player, error)
// }