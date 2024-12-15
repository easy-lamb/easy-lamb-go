package models

type Config struct {
	LambdaDir         string            `json:"lambdaDir"`
	TerraformDir      string            `json:"terraformDir"`
	TerraformFilename string            `json:"terraformFilename"`
	BuildOutput       string            `json:"buildOutput"`
	DefaultParams     map[string]string `json:"defaultParams"`
	DotenvLocation    string            `json:"dotenvLocation"`
}
