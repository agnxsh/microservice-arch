package data

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB


//New is a function that creates an instance of the data package, the return type is the Model
//which embeds all the types we want to be available in the application
func New (dbPool *sql.DB) Models {
	db = dbPool
	return Models {
		User: User{},
	}
}

//Models is the type for this package. Note that any model that is included as a member 
//in this type is available to us throughout the application, anywhere that the app variable is used
//provided that the model is also added in the New function
type Models struct {
	User user
}

type User struct {
	ID			int 		`json:"id"`
	Email 		string 		`json:"email"`
	FirstName 	string 		`json:"first_name,omitempty"`
	LastName 	string 		`json:"last_name,omitempty"`
	Password 	string		`json:"-"`
	Active		int 		`json:"active"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`	
}

//GetAll returns a slice of all users, sorted by last name
func (u *User) GetAll() ([]*User, error){
	ctx, cancel:= context.WithTimeout(context.Background(), dbTimeout)

	defer cancel()

	query:= `select id, email, first_name, last_name, password, user_active, created_at, updated_at 
	from users order by last_name`

	rows, err:= db.QueryContext((ctx,query))
	if err!= nil {
		return nil, err
	}

	defer rows.Close()
	var users [] *User

	for rows.Next() {
		var user User
		err:= rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err!= nil {
			log.Println("Error in scanning dataset", err)
			return nil, err
		}

		return users, nil
	}
}

//GetByEmail returns one user by email
func (u *User) GetByEmail (email string) (*User, error)[
	ctx, cancel := context.WithTimeout(context.Background, dbTimeout)
	defer cancel()

	query :=  `select id, email, first_name, last_name, password, user_active, created_at, updated_at`
]