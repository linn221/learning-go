package seeders

import "fmt"

func Run() {
	fmt.Println("Seeder running..")

	if err := SeedCategories(); err != nil {
		fmt.Println("Error seeding categories: " + err.Error())
	} else {
		fmt.Println("Seeding categories success.")
	}

	if err := SeedTags(); err != nil {
		fmt.Println("Error seeding tags: " + err.Error())
	} else {
		fmt.Println("Seeding tags success.")
	}

	if err := SeedPosts(2000); err != nil {
		fmt.Println("Error seeding posts: " + err.Error())
	} else {
		fmt.Println("Seeding posts success.")
	}
}
