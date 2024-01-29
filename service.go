package main

func getPeople(params PeopleGet) ([]Person, error) {
	var people []Person
	offset := (params.Page - 1) * params.Limit
	tx := db.Where(params.Query).Limit(params.Limit).Offset(offset).Find(&people)
	return people, tx.Error
}

func createPerson(personBody PersonCreate) error {
	var err error
	person := Person{
		Name:       personBody.Name,
		Surname:    personBody.Surname,
		Patronymic: personBody.Patronymic,
	}
	person.Age, err = getAge(person.Name)
	if err != nil {
		return err
	}
	person.Gender, err = getGender(person.Name)
	if err != nil {
		return err
	}
	person.Nationality, err = getNationality(person.Name)
	if err != nil {
		return err
	}

	tx := db.Create(&person)
	return tx.Error
}

func updatePerson(personBody PersonUpdate) error {
	person := Person{
		Name:        personBody.Name,
		Surname:     personBody.Surname,
		Patronymic:  personBody.Patronymic,
		Age:         personBody.Age,
		Gender:      personBody.Gender,
		Nationality: personBody.Nationality,
	}
	tx := db.Where("id = ?", personBody.ID).Updates(person)
	return tx.Error
}

func deletePerson(personBody PersonUpdate) error {
	tx := db.Where("id = ?", personBody.ID).Delete(&Person{})
	return tx.Error
}
