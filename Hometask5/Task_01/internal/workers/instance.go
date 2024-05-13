package workers

import (
	"strconv"
	"time"
)

type Worker struct {
	EncName        string
	EncSurname     string
	vacations      [][2]time.Time
	vacDayDuration int
}

func NewWorker(name string, surname string, vac [][2]time.Time) *Worker {
	return &Worker{
		EncName:        name,
		EncSurname:     surname,
		vacations:      vac,
		vacDayDuration: countVacDayDur(vac),
	}
}

func (w *Worker) AddVacationDate(start time.Time, end time.Time) {
	w.vacations = append(w.vacations, [2]time.Time{start, end})
	w.vacDayDuration += int(end.Sub(start).Hours() / 24)
}

func (w *Worker) GetVacDates() [][2]time.Time {
	return w.vacations
}

func countVacDayDur(vacations [][2]time.Time) (res int) {
	for _, vac := range vacations {
		res += int(vac[1].Sub(vac[0]).Hours() / 24)
	}
	return
}

func (w *Worker) GetVacDayDur() int {
	return w.vacDayDuration
}

func (w Worker) String() string {
	return w.EncName + " " + w.EncSurname + ": " + strconv.Itoa(w.vacDayDuration)
}
