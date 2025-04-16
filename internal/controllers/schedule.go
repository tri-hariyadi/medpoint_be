package controllers

import (
	"github.com/sev-2/raiden"
	"medpoint/internal/models"
)

type ScheduleController struct {
	raiden.ControllerBase
	Http  string `path:"/schedules" type:"rest"`
	Model models.Schedules
}
