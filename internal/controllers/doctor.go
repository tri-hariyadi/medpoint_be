package controllers

import (
	"github.com/sev-2/raiden"
	"medpoint/internal/models"
)

type DoctorController struct {
	raiden.ControllerBase
	Http  string `path:"/doctors" type:"rest"`
	Model models.Doctors
}
