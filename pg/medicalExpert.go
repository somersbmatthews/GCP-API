package pg

type MedicalExpert struct {
	UserID                string `gorm:"unique"`
	Email                 string
	Name                  string
	Specialty             string
	Degree                string
	Verified              bool
	EmailVerificationCode string
}
