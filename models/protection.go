package models

type ProtectionMethod struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Effectiveness float64 `json:"effectiveness"`
}

var ProtectionMethods = []ProtectionMethod{
	{"Антивирус", "Защита от вирусов с использованием антивирусного ПО", 0.85},
	{"Фаервол", "Защита сети с помощью фаервола", 0.90},
}
