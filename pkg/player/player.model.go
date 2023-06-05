package player

type Player struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Level int `json:"level"`
	Experience int `json:"experience"`
}

// type PlayerService interface {
// 	GetAll() ([]Player, error)
// 	GetByID(id int) (Player, error)
// 	Create(player Player) (Player, error)
// 	Update(player Player) (Player, error)
// 	Delete(id int) (Player, error)
// }