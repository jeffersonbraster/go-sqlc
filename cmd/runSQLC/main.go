package main

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jeffersonbraster/go-sqlc/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/courses")
	if err != nil {
		panic(err)
	}

	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Category 1",
	// 	Description: sql.NullString{String: "Description 1"},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)

	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "8124d169-12f9-41c0-a602-7eb13b81ec39",
	// 	Name:        "um nome",
	// 	Description: sql.NullString{String: "aqui tem uma desc", Valid: true},
	// })

	// categories, err := queries.ListCategories(ctx)

	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	err = queries.DeleteCategory(ctx, "8124d169-12f9-41c0-a602-7eb13b81ec39")

	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

}