package model

//   - -----------------------------------------
//     Configuration for service
//   - -----------------------------------------
type Config struct {
	DBConnectionString    string
	FinancialsApplyURL    string
	FinancialsRollbackURL string
	AuthenticationURL     string
	FraudURL              string
	BookkepingApplyURL    string
	BookkepingUpdateURL   string
}

func NewConfig() *Config {
	return &Config{}
}
