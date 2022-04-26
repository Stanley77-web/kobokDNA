package Models

type Result struct {
	ID              int    `json:"id" bson:"id"`
	Date            string `json:"date" bson:"date"`
	UserName        string `json:"user_name" bson:"user_name"`
	DiseaseName     string `json:"disease_name" bson:"disease_name"`
	SimilarityScore int    `json:"similarity_score" bson:"similarity"`
	Status          bool   `json:"status" bson:"status"`
}
