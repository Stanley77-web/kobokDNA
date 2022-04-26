package Models

type Disease struct {
	DiseaseName string `json:"disease_name" bson:"name_disease" binding:"required"`
	SequenceDNA string `json:"disease_sequence_DNA" bson:"sequence_DNA" binding:"required"`
}
