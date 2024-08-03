package config

import "github.com/joao-gabriel-cruz/debora-api/prisma/db"

func ConnectDatabase() (*db.PrismaClient, error) {
	// Connect to database
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return nil, err
	}
	return client, nil
}
