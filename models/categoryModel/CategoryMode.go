package categoryModel

import (
	"go-web-native/config"
	"go-web-native/entities"
)

func GetAll() []entities.Category {

	res, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		panic(err)
	}

	defer res.Close()

	var categories []entities.Category

	for res.Next() {
		var category entities.Category
		if err := res.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
			INSERT INTO categories (name, created_at, updated_at) 
			VALUES (?, ?, ?)
		`, category.Name, category.CreatedAt, category.UpdatedAt)

	if err != nil {
		panic(err)
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}

	return lastId > 0

}

func Detail(id int) entities.Category {
	res := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = ?`, id)
	var category entities.Category
	if err := res.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec("UPDATE categories SET name = ?, updated_at = ? WHERE id = ?", category.Name, category.UpdatedAt, id)
	if err != nil {
		panic(err)
	}
	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}
