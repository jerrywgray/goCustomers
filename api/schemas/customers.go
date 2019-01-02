package schemas

import (
	"github.com/jinzhu/gorm"
)

// Customers is the ORM struct of our Customers database
type Customers struct {
	gorm.Model
}
