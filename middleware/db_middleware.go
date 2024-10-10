package middleware

import "saved_messages_classifier/classifier"

func PGMiddleware(dbService *classifier.Queries) {
	return func(c *fiber.Ctx) error {
		// Check if the database connection is already set in the context
		if gormU := c.Locals("gorm"); gormU != nil {
			if _, ok := gormU.(*gorm.DB); ok {
				return c.Next()
			} else {
				log.Println("Database connection already set in the context but not of type *gorm.DB")

			}

		}

		// otherwise initialize the database connection
		c.Locals("gorm", dbService.DB)
		return c.Next()
	}
}
