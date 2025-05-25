package dto

type CreateFavoritePokemonRequest struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Notes     string `json:"notes" binding:"required,max=30"`
	Sprite    string `json:"sprite"`
}

type UpdateFavoritePokemonRequest struct {
	Name     string `json:"name""`
	Type     string `json:"type"`
	Notes    string `json:"notes" binding:"required,max=30"`
}
