package models

type ScenarioPackage struct {
	Data    string `json:"data"`
	Delay   int    `json:"delay"`   // delay in milliseconds
	IsFinal bool   `json:"isFinal"` // indicates if this is the final package
}

type Scenario struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Packages    []ScenarioPackage `json:"packages"`
}
