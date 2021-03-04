package poker

//type Game struct {
//	alerter BlindAlerter
//	store   PlayerStore
//}
//
//func NewGame(alerter BlindAlerter, store PlayerStore) *Game {
//	return &Game{
//		alerter: alerter,
//		store:   store,
//	}
//}
//
//func (p *Game) Start(numberOfPlayers int) {
//	blindInCrement := time.Duration(5+numberOfPlayers) * time.Minute
//
//	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
//	blindTime := 0 * time.Second
//	for _, blind := range blinds {
//		p.alerter.ScheduleAlertAt(blindTime, blind)
//		blindTime = blindTime + blindInCrement
//	}
//}
//
//func (p *Game) Finish(winner string) {
//	p.store.RecordWin(winner)
//}
type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}
