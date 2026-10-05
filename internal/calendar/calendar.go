package calendar

import (
	"fmt"
	"time"
)

type WeekStart string
const ( Monday WeekStart = "monday"; Sunday WeekStart = "sunday" )
var weekdays = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
var months = [...]string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

type Invariants struct { SequentialDates bool `json:"sequentialDates"`; ValidRowCount bool `json:"validRowCount"`; UniqueDates bool `json:"uniqueDates"` }
type Week struct { Dates []string `json:"dates"` }
type Month struct {
	Year int `json:"year"`; Month int `json:"month"`; MonthName string `json:"monthName"`; DaysInMonth int `json:"daysInMonth"`
	FirstWeekday string `json:"firstWeekday"`; WeekStart WeekStart `json:"weekStart"`; LeadingCells int `json:"leadingCells"`; NaturalRows int `json:"naturalRows"`
	Weeks []Week `json:"weeks"`; Invariants Invariants `json:"invariants"`
}
type YearMonth struct { Month int `json:"month"`; MonthName string `json:"monthName"`; StartWeekday string `json:"startWeekday"`; Days int `json:"days"`; MondayRows int `json:"mondayRows"`; SundayRows int `json:"sundayRows"`; ISOStartYear int `json:"isoStartYear"`; ISOEndYear int `json:"isoEndYear"` }

func valid(y, m int) bool { return y >= 1 && y <= 9999 && m >= 1 && m <= 12 }
func IsLeap(y int) bool { return y%400 == 0 || (y%4 == 0 && y%100 != 0) }
func DaysInMonth(y, m int) int { if !valid(y,m) { return 0 }; return time.Date(y,time.Month(m)+1,0,0,0,0,0,time.UTC).Day() }
func parseWeekStart(s string) (WeekStart,error) { if s=="" || s=="monday" { return Monday,nil }; if s=="sunday" { return Sunday,nil }; return "",fmt.Errorf("weekStart must be monday or sunday") }
func Geometry(y,m int, start string) (Month,error) {
	ws,err:=parseWeekStart(start); if err!=nil{return Month{},err}; if !valid(y,m){return Month{},fmt.Errorf("year must be 1..9999 and month 1..12")}
	first:=time.Date(y,time.Month(m),1,0,0,0,0,time.UTC); dow:=int(first.Weekday()); offset:=1; if ws==Sunday {offset=0}
	leading:=(dow-offset+7)%7; days:=DaysInMonth(y,m); rows:=(leading+days+6)/7
	grid:=make([]Week,rows); for i:=range grid { grid[i].Dates=make([]string,7) }
	seen:=map[string]bool{}; seq:=true; idx:=0
	for day:=1;day<=days;day++ { cell:=leading+day-1; d:=time.Date(y,time.Month(m),day,0,0,0,0,time.UTC); value:=d.Format("2006-01-02"); grid[cell/7].Dates[cell%7]=value; if seen[value]{seq=false};seen[value]=true; if idx>0 { prev:=time.Date(y,time.Month(m),day-1,0,0,0,0,time.UTC);if !d.Equal(prev.AddDate(0,0,1)){seq=false} };idx++ }
	return Month{Year:y,Month:m,MonthName:months[m],DaysInMonth:days,FirstWeekday:weekdays[dow],WeekStart:ws,LeadingCells:leading,NaturalRows:rows,Weeks:grid,Invariants:Invariants{seq,rows>=4&&rows<=6,len(seen)==days}},nil
}
func YearMatrix(y int)([]YearMonth,error){ if y<1||y>9998{return nil,fmt.Errorf("year must be 1..9998")}; out:=make([]YearMonth,12);for m:=1;m<=12;m++{a,_:=Geometry(y,m,"monday");b,_:=Geometry(y,m,"sunday"); _,isoStart:=time.Date(y,time.Month(m),1,0,0,0,0,time.UTC).ISOWeek(); _,isoEnd:=time.Date(y,time.Month(m),DaysInMonth(y,m),0,0,0,0,time.UTC).ISOWeek();out[m-1]=YearMonth{m,months[m],a.FirstWeekday,a.DaysInMonth,a.NaturalRows,b.NaturalRows,isoStart,isoEnd}};return out,nil }
func Boundaries(y int)(map[string]any,error){if y<1||y>9998{return nil,fmt.Errorf("year must be 1..9998")}; jan:=time.Date(y,1,1,0,0,0,0,time.UTC);dec:=time.Date(y,12,31,0,0,0,0,time.UTC);jy,jw:=jan.ISOWeek();dy,dw:=dec.ISOWeek();days:=365;if IsLeap(y){days=366};return map[string]any{"year":y,"leapYear":IsLeap(y),"daysInYear":days,"januaryFirstWeekday":weekdays[int(jan.Weekday())],"decemberLastWeekday":weekdays[int(dec.Weekday())],"januaryFirstISOWeek":map[string]int{"year":jy,"week":jw},"decemberLastISOWeek":map[string]int{"year":dy,"week":dw},"isoWeekYearDiffersAtStart":jy!=y,"isoWeekYearDiffersAtEnd":dy!=y},nil }
