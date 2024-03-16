package models

type CharacterDynamic struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CharacterFixed struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Hp   int    `json:"hp"`
}
