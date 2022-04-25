package Modules

type TestData struct {
	UserName    string `json:"user_name" binding:"required"`
	DNASequence string `json:"user_DNA_sequence" binding:"required"`
	DiseaseName string `json:"disease_name" binding:"required"`
	Method      string `json:"method" binding:"required"`
}
