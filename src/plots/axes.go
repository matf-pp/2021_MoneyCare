package plots

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"main/src/admin"
	"os"
	"time"

	chart "github.com/wcharczuk/go-chart/v2"
)

func DrawChart(userID primitive.ObjectID, time2 time.Time) {
	/*
	   This is an example of using the `TimeSeries` to automatically coerce time.Time values into a continuous xrange.
	   Note: chart.TimeSeries implements `ValueFormatterProvider` and as a result gives the XAxis the appropriate formatter to use for the ticks.
	*/
	spendingService := admin.SpendingService

	day0 := time.Now().AddDate(0, 0, -6)
	day1 := time.Now().AddDate(0, 0, -5)
	day2 := time.Now().AddDate(0, 0, -4)
	day3 := time.Now().AddDate(0, 0, -3)
	day4 := time.Now().AddDate(0, 0, -2)
	day5 := time.Now().AddDate(0, 0, -1)
	day6 := time.Now()

	amount0 := spendingService.FindUsersSpendingByDate(userID, day0)
	fmt.Println(amount0)
	amount1 := spendingService.FindUsersSpendingByDate(userID, day1)
	fmt.Println(amount1)
	amount2 := spendingService.FindUsersSpendingByDate(userID, day2)
	amount3 := spendingService.FindUsersSpendingByDate(userID, day3)
	amount4 := spendingService.FindUsersSpendingByDate(userID, day4)
	amount5 := spendingService.FindUsersSpendingByDate(userID, day5)
	amount6 := spendingService.FindUsersSpendingByDate(userID, day6)

	graph := chart.Chart{
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: []time.Time{
					day0,
					day1,
					day2,
					day3,
					day4,
					day5,
					day6,
				},
				YValues: []float64{amount0, amount1, amount2, amount3, amount4, amount5, amount6},
			},
		},
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}

//package plots
//
//import (
//	"fmt"
//	chart "github.com/wcharczuk/go-chart/v2"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"main/src/admin"
//	"os"
//	"time"
//)
//
//func Axes(userID primitive.ObjectID, time2 time.Time) {
//
//	/*
//	   The below will draw the same chart as the `basic` example, except with both the x and y axes turned on.
//	   In this case, both the x and y axis ticks are generated automatically, the x and y ranges are established automatically, the canvas "box" is adjusted to fit the space the axes occupy so as not to clip.
//	*/
//
//	spendingService := admin.SpendingService
//
//	day0 := time2
//	day1 := time2.Add((-24)*time.Hour)
//	day2 := time2.Add(2*(-24)*time.Hour)
//	day3 := time2.Add(3*(-24)*time.Hour)
//	day4 := time2.Add(4*(-24)*time.Hour)
//	day5 := time2.Add(5*(-24)*time.Hour)
//	day6 := time2.Add(6*(-24)*time.Hour)
//
//	amount0 := spendingService.FindUsersSpendingByDate(userID, day0)
//	fmt.Println(amount0)
//	amount1 := spendingService.FindUsersSpendingByDate(userID, day1)
//	fmt.Println(amount1)
//	amount2 := spendingService.FindUsersSpendingByDate(userID, day2)
//	amount3 := spendingService.FindUsersSpendingByDate(userID, day3)
//	amount4 := spendingService.FindUsersSpendingByDate(userID, day4)
//	amount5 := spendingService.FindUsersSpendingByDate(userID, day5)
//	amount6 := spendingService.FindUsersSpendingByDate(userID, day6)
//
//
//	graph := chart.Chart{
//		Series: []chart.Series{
//			chart.ContinuousSeries{
//				Style: chart.Style{
//					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
//					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
//				},
//				XValues: []float64{6, 5, 4, 3, 2, 1, 0},
//				YValues: []float64{amount0, amount1, amount2, amount3, amount4, amount5, amount6},
//			},
//		},
//	}
//
//	f, _ := os.Create("output.png")
//	defer f.Close()
//	graph.Render(chart.PNG, f)
//}
