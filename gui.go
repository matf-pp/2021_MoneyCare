package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"main/src/admin"
	"main/src/plots"
	"main/src/services"
	"math"
	"strconv"
	"time"
)

var entryUpIncomeAmount = -1
var currentUserMean = make(map[string]float64)
var username string
var userID primitive.ObjectID
var userMonth = time.Now().Month()
var userTime = time.Now()

var imgPie *gtk.Image
var imgGraph *gtk.Image

func init() {
	entryUpIncomeAmount = -1
}

func mean(xs [30]float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(5)
}

func setCurrentUserMean() {
	var currentUser int
	if entryUpIncomeAmount != -1 {
		if entryUpIncomeAmount <= 50000 {
			currentUser = 0
		}
		if entryUpIncomeAmount >= 50000 && entryUpIncomeAmount <= 150000 {
			currentUser = 1
		}
		if entryUpIncomeAmount >= 150000 {
			currentUser = 2
		}
	}
	var xs [30]float64

	idFood, _ := admin.CategoryService.FindOne("food")
	idClot, _ := admin.CategoryService.FindOne("clothes")
	idChem, _ := admin.CategoryService.FindOne("chem")
	idOther, _ := admin.CategoryService.FindOne("other")
	idBills, _ := admin.CategoryService.FindOne("bills")

	if currentUser == 0 {

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID0[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idFood.ID)
		}

		currentUserMean["food"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID0[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idClot.ID)
		}

		currentUserMean["clothes"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID0[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idChem.ID)
		}

		currentUserMean["chem"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID0[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idBills.ID)
		}

		currentUserMean["bills"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID0[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idOther.ID)
		}

		currentUserMean["other"] = mean(xs)

	}

	if currentUser == 1 {

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID1[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idFood.ID)
		}

		currentUserMean["food"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID1[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idClot.ID)
		}

		currentUserMean["clothes"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID1[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idChem.ID)
		}

		currentUserMean["chem"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID1[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idBills.ID)
		}

		currentUserMean["bills"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID1[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idOther.ID)
		}

		currentUserMean["other"] = mean(xs)

	}

	if currentUser == 2 {
		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID2[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idFood.ID)
		}

		currentUserMean["food"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID2[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idClot.ID)
		}

		currentUserMean["clothes"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID2[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idChem.ID)
		}

		currentUserMean["chem"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID2[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idBills.ID)
		}

		currentUserMean["bills"] = mean(xs)
		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID2[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idOther.ID)
		}

		currentUserMean["other"] = mean(xs)
	}
}

func addSpending(cat string, entry *gtk.Entry) {
	categoryService := admin.CategoryService
	categoryID, err := categoryService.FindOne(cat)
	if err != nil {
		panic(err)
	}
	spendingService := admin.SpendingService
	spendingService.InsertFromEntry(userID, categoryID.ID, entry)
}

func showWarning(cat string, eps float64, popup *gtk.Window) {

	categoryService := admin.CategoryService
	catID, err := categoryService.FindOne(cat)
	if err != nil {
		panic(err)
	}
	spendingService := admin.SpendingService
	s := spendingService.FindUsersSpendingByCategory(userID, catID.ID)
	if math.Abs(s-currentUserMean[cat]) < eps || s > currentUserMean[cat] {
		popup.ShowAll()
	}
}

func showWarningForGoal(percent float64, popupGoal *gtk.Window) {
	userService := admin.UserService
	user, err := userService.FindOne(username)
	if err != nil {
		panic(err)
	}
	goal := user.Goal
	income := user.Income
	p := goal * percent / 100.00
	spendingService := admin.SpendingService
	spend := spendingService.FindUsersSpending(userID)
	if spend > (income-goal)-p {
		popupGoal.ShowAll()
	}

}

func showBalance(service *services.SpendingService, label *gtk.Label, pb *gtk.ProgressBar) {
	spent := service.FindUsersSpendingByMonth(userID, userMonth)
	s := fmt.Sprint("You've spent : ", spent)
	label.SetText(s)

	userService := admin.UserService
	us, err := userService.FindOne(username)
	if err != nil {
		panic(err)
	}
	total := us.Income
	x := (100.00 * spent / total) / 100.00
	pb.SetFraction(x)
}

func showBalanceByCat(service *services.SpendingService, label *gtk.Label, pb *gtk.ProgressBar, cat string) {
	categoryService := admin.CategoryService
	catID, err := categoryService.FindOne(cat)
	if err != nil {
		panic(err)
	}
	spent := service.FindUsersSpendingByCategoryByMonth(userID, catID.ID, userMonth)
	s := fmt.Sprint("You've spent : ", spent)
	label.SetText(s)
	total := currentUserMean[cat]
	x := (100.00 * spent / total) / 100.00
	pb.SetFraction(x)
}

func showBalanceForAll(labelAll *gtk.Label, labelFood *gtk.Label, labelClothes *gtk.Label, labelChem *gtk.Label, labelOther *gtk.Label, labelBills *gtk.Label, pbAll *gtk.ProgressBar, pbFood *gtk.ProgressBar, pbClothes *gtk.ProgressBar, pbChem *gtk.ProgressBar, pbOther *gtk.ProgressBar, pbBills *gtk.ProgressBar) {
	spendingService := admin.SpendingService
	showBalance(spendingService, labelAll, pbAll)
	showBalanceByCat(spendingService, labelFood, pbFood, "food")
	showBalanceByCat(spendingService, labelBills, pbBills, "bills")
	showBalanceByCat(spendingService, labelClothes, pbClothes, "clothes")
	showBalanceByCat(spendingService, labelOther, pbOther, "other")
	showBalanceByCat(spendingService, labelChem, pbChem, "chem")
}

func setupWindow(title string) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(800, 600)
	win.SetPosition(gtk.WIN_POS_CENTER)
	return win
}
func setupPopup(width int, height int, title string) *gtk.Window {
	popup, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	popup.SetTitle(title)
	popup.SetDefaultSize(width, height)
	popup.SetPosition(gtk.WIN_POS_CENTER)
	return popup
}

func setupBtn(label string, onClick func()) *gtk.Button {
	btn, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	btn.Connect("clicked", onClick)
	return btn
}

func setupFixed() *gtk.Fixed {
	fixed, err := gtk.FixedNew()
	if err != nil {
		log.Fatal("Unable to create GtkFixed:", err)
	}
	return fixed
}

func setupLabel(text string) *gtk.Label {
	label, err := gtk.LabelNew(text)
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	return label
}

func setupEntry() *gtk.Entry {
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry:", err)
	}
	return entry
}

func setupProgressBar() *gtk.ProgressBar {
	pb, err := gtk.ProgressBarNew()
	if err != nil {
		log.Fatal("Unable to create progress bar:", err)
	}
	return pb
}

func showUsername(uname string, label *gtk.Label) {
	label.SetText(uname)
}

func setupGui() {

	cal, err := gtk.CalendarNew()
	if err != nil {
		panic(err)
	}

	win := setupWindow("Money Care")
	fixed := setupFixed()
	fixedSignIn := setupFixed()
	fixedSignUp := setupFixed()
	fixedFood := setupFixed()
	fixedChem := setupFixed()
	fixedClo := setupFixed()
	fixedOth := setupFixed()
	fixedBills := setupFixed()
	fixedWarning := setupFixed()
	fixedWarningGoal := setupFixed()

	popupSignIn := setupPopup(150, 120, "SignIn")
	popupSignUp := setupPopup(150, 120, "SignUp")
	popupFood := setupPopup(150, 120, "Food")
	popupChem := setupPopup(150, 120, "Chem")
	popupOth := setupPopup(150, 120, "Other")
	popupClo := setupPopup(150, 120, "Clothes")
	popupBills := setupPopup(150, 120, "Bills")
	popupWarningGoal := setupPopup(15, 120, "Goal")
	popupWarning := setupPopup(150, 120, "Warning")

	btWarningGoalOk := setupBtn("OK", func() {
		popupWarningGoal.Hide()
	})

	btWarningOK := setupBtn("OK", func() {
		popupWarning.Hide()
	})

	btSignUp := setupBtn("SIGN UP", func() {
		popupSignUp.ShowAll()
	})
	btSignIn := setupBtn("SIGN IN", func() {
		popupSignIn.ShowAll()
	})
	btFood := setupBtn("FOOD", func() {
		if username == "" {
			dial, err := gtk.DialogNew()
			if err != nil {
				panic(err)
			}
			dial.AddButton("PLEASE SIGN IN", gtk.RESPONSE_OK)
			dial.SetTitle("Please")
			dial.Show()
		} else {
			popupFood.ShowAll()
		}
	})
	btChem := setupBtn("CHEM/COSM", func() {
		if username == "" {
			dial, err := gtk.DialogNew()
			if err != nil {
				panic(err)
			}
			dial.AddButton("PLEASE SIGN IN", gtk.RESPONSE_OK)
			dial.SetTitle("Please")
			dial.Show()
		} else {
			popupChem.ShowAll()
		}
	})
	btClo := setupBtn("CLOTHES", func() {
		if username == "" {
			dial, err := gtk.DialogNew()
			if err != nil {
				panic(err)
			}
			dial.AddButton("PLEASE SIGN IN", gtk.RESPONSE_OK)
			dial.SetTitle("Please")
			dial.Show()
		} else {
			popupClo.ShowAll()
		}
	})
	btBill := setupBtn("BILLS", func() {
		if username == "" {
			dial, err := gtk.DialogNew()
			if err != nil {
				panic(err)
			}
			dial.AddButton("PLEASE SIGN IN", gtk.RESPONSE_OK)
			dial.SetTitle("Please")
			dial.Show()
		} else {
			popupBills.ShowAll()
		}
	})
	btOth := setupBtn("OTHER", func() {
		if username == "" {
			dial, err := gtk.DialogNew()
			if err != nil {
				panic(err)
			}
			dial.AddButton("PLEASE SIGN IN", gtk.RESPONSE_OK)
			dial.SetTitle("Please")
			dial.Show()
		} else {
			popupOth.ShowAll()
		}
	})

	btPieChart := setupBtn("PIECHART", func() {
		if username == "" {
			dial, err := gtk.DialogNew()
			if err != nil {
				panic(err)
			}
			dial.AddButton("PLEASE SIGN IN", gtk.RESPONSE_OK)
			dial.SetTitle("Please")
			dial.Show()
		} else {
			if imgPie != nil && imgPie.IsVisible() {
				imgPie.Hide()
			}
			plots.PieChart(userID, userMonth)
			imgPie, err = gtk.ImageNewFromFile("piechart.png")
			if err != nil {
				panic(err)
			}
			fixed.Put(imgPie, 441, 60)
			if imgGraph != nil && imgGraph.IsVisible() {
				imgGraph.Hide()
			}
			imgPie.Show()
		}
	})

	btGraph := setupBtn("GRAPH", func() {
		if username == "" {
			dial, err := gtk.DialogNew()
			if err != nil {
				panic(err)
			}
			dial.AddButton("PLEASE SIGN IN", gtk.RESPONSE_OK)
			dial.SetTitle("Please")
			dial.Show()
		} else {
			if imgGraph != nil && imgGraph.IsVisible() {
				imgGraph.Hide()
			}
			plots.DrawChart(userID, userTime)
			imgGraph, err = gtk.ImageNewFromFile("graph.png")
			if err != nil {
				panic(err)
			}
			fixed.Put(imgGraph, 441, 60)
			if imgPie != nil && imgPie.IsVisible() {
				imgPie.Hide()
			}
			imgGraph.Show()
		}
	})

	labBalance := setupLabel("balance")
	labelFoodEx := setupLabel("balance")
	labelBillsEx := setupLabel("balance")
	labelChemEx := setupLabel("balance")
	labelOtherEx := setupLabel("balance")
	labelClothesEx := setupLabel("balance")

	labUpID := setupLabel("USERNAME: ")
	labUpIncome := setupLabel("INCOME: ")
	labUpOutgoings := setupLabel("OUTGOINGS: ")
	labUpGoal := setupLabel("GOAL: ")
	labFood := setupLabel("Enter amount: ")
	labChem := setupLabel("Enter amount: ")
	labOth := setupLabel("Enter amount: ")
	labClo := setupLabel("Enter amount: ")
	labBills := setupLabel("Enter amount: ")
	labWarning := setupLabel("Spending too much\non this category!")
	labWarningGoal := setupLabel("Remember your goal!")

	labUsername := setupLabel("USERNAME")

	pbFood := setupProgressBar()
	pbChem := setupProgressBar()
	pbClo := setupProgressBar()
	pbBill := setupProgressBar()
	pbOth := setupProgressBar()
	pb := setupProgressBar()

	entryIn := setupEntry()
	entryUpID := setupEntry()
	entryUpOutgoings := setupEntry()
	entryUpIncome := setupEntry()
	entryUpGoal := setupEntry()
	entryFood := setupEntry()
	entryChem := setupEntry()
	entryOth := setupEntry()
	entryClo := setupEntry()
	entryBills := setupEntry()

	btSignUpOK := setupBtn("OK", func() {
		enID, err := entryUpID.GetText()
		if err != nil {
			panic(err)
		}

		enGoal, err := entryUpGoal.GetText()
		if err != nil {
			panic(err)
		}
		enIncome, err := entryUpIncome.GetText()
		if err != nil {
			panic(err)
		}
		enOutgoings, err := entryUpOutgoings.GetText()
		if err != nil {
			panic(err)
		}

		userService := admin.UserService
		us, err := userService.FindOne(enID)

		if err != mongo.ErrNoDocuments {
			entryUpID.SetText("USERNAME TAKEN")
		} else if enID == "" {
			entryUpID.SetText("Enter username")
		} else if enGoal == "" {
			entryUpGoal.SetText("Enter goal")
		} else if enIncome == "" {
			entryUpIncome.SetText("Enter income")
		} else if enOutgoings == "" {
			entryUpOutgoings.SetText("Enter outgoings")
		} else {

			userService.InsertFromEntry(entryUpID, entryUpGoal, entryUpIncome, entryUpOutgoings)

			str, err := entryUpIncome.GetText()
			entryUpIncomeAmount, _ = strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			username = us.Username
			user, err := userService.FindOne(us.Username)
			userID = user.ID
			popupSignUp.Hide()
		}

	})

	btSignInOK := setupBtn("OK", func() {
		userService := admin.UserService
		uname, err := entryIn.GetText()
		if err != nil {
			panic(err)
		}
		user, err := userService.FindOne(uname)

		if err == mongo.ErrNoDocuments {
			entryIn.SetText("DOESN'T EXIST")
		} else {

			username = uname
			userID = user.ID

			showUsername(username, labUsername)

			entryUpIncomeAmount = int(user.Income)
			setCurrentUserMean()

			showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
			popupSignIn.Hide()
		}
	})

	btFoodOK := setupBtn("OK", func() {

		addSpending("food", entryFood)
		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
		showWarning("food", 100.00, popupWarning)
		showWarningForGoal(30, popupWarningGoal)
		popupFood.Hide()

	})
	btChemOK := setupBtn("OK", func() {
		addSpending("chem", entryChem)
		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
		showWarning("chem", 100.00, popupWarning)
		showWarningForGoal(30, popupWarningGoal)
		popupChem.Hide()
	})
	btCloOK := setupBtn("OK", func() {
		addSpending("clothes", entryClo)
		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
		showWarning("clothes", 100.00, popupWarning)
		showWarningForGoal(30, popupWarningGoal)
		popupClo.Hide()
	})
	btOthOK := setupBtn("OK", func() {
		addSpending("other", entryOth)
		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
		showWarning("other", 50.00, popupWarning)
		showWarningForGoal(30, popupWarningGoal)
		popupOth.Hide()
	})
	btBillsOK := setupBtn("OK", func() {
		addSpending("bills", entryBills)
		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
		showWarning("bills", 50.00, popupWarning)
		showWarningForGoal(30, popupWarningGoal)
		popupBills.Hide()
	})
	btFoodClose := setupBtn("CLOSE", func() {

		popupFood.Hide()
	})
	btChemClose := setupBtn("CLOSE", func() {

		popupChem.Hide()
	})
	btCloClose := setupBtn("CLOSE", func() {

		popupClo.Hide()
	})
	btOthClose := setupBtn("CLOSE", func() {

		popupOth.Hide()
	})
	btBillsClose := setupBtn("CLOSE", func() {

		popupBills.Hide()
	})
	btSignInClose := setupBtn("CLOSE", func() {

		popupSignIn.Hide()
	})
	btSignUpClose := setupBtn("CLOSE", func() {

		popupSignUp.Hide()
	})

	btCalOK := setupBtn("OK", func() {
		if username == "" {
			dial, err := gtk.DialogNew()
			if err != nil {
				panic(err)
			}
			dial.AddButton("PLEASE SIGN IN", gtk.RESPONSE_OK)
			dial.SetTitle("Please")
			dial.Show()
		} else {
			y, m, d := cal.GetDate()
			userTime = time.Date(int(y), time.Month(m+1), int(d), 12, 12, 12, 12, time.Local)
			userMonth = time.Month(m + 1)
			showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
		}
	})

	fixed.Put(labelFoodEx, 250, 100)
	fixed.Put(labelChemEx, 250, 200)
	fixed.Put(labelClothesEx, 250, 300)
	fixed.Put(labelBillsEx, 250, 400)
	fixed.Put(labelOtherEx, 250, 500)

	fixed.Put(pbFood, 100, 150)
	fixed.Put(pbChem, 100, 250)
	fixed.Put(pbClo, 100, 350)
	fixed.Put(pbBill, 100, 450)
	fixed.Put(pbOth, 100, 550)
	fixed.Put(pb, 400, 500)

	fixed.Put(btSignUp, 10, 10)
	fixed.Put(btSignIn, 110, 10)
	fixed.Put(btFood, 100, 100)
	fixed.Put(btChem, 100, 200)
	fixed.Put(btClo, 100, 300)
	fixed.Put(btBill, 100, 400)
	fixed.Put(btOth, 100, 500)
	fixed.Put(btPieChart, 532, 10)
	fixed.Put(btGraph, 441, 10)

	fixed.Put(cal, 600, 500)
	fixed.Put(btCalOK, 600, 450)

	fixed.Put(labBalance, 450, 450)
	fixed.Put(labUsername, 250, 15)

	fixedSignIn.Put(btSignInOK, 10, 70)
	fixedSignIn.Put(btSignInClose, 80, 70)
	fixedSignIn.Put(entryIn, 0, 30)

	fixedSignUp.Put(entryUpID, 100, 10)
	fixedSignUp.Put(entryUpIncome, 100, 50)
	fixedSignUp.Put(entryUpOutgoings, 100, 90)
	fixedSignUp.Put(entryUpGoal, 100, 130)
	fixedSignUp.Put(labUpID, 0, 10)
	fixedSignUp.Put(labUpIncome, 0, 50)
	fixedSignUp.Put(labUpOutgoings, 0, 90)
	fixedSignUp.Put(labUpGoal, 0, 130)
	fixedSignUp.Put(btSignUpOK, 20, 180)
	fixedSignUp.Put(btSignUpClose, 130, 180)

	fixedFood.Put(btFoodOK, 20, 80)
	fixedFood.Put(btFoodClose, 80, 80)
	fixedFood.Put(entryFood, 0, 40)
	fixedFood.Put(labFood, 10, 10)

	fixedChem.Put(btChemOK, 20, 80)
	fixedChem.Put(btChemClose, 80, 80)
	fixedChem.Put(entryChem, 0, 40)
	fixedChem.Put(labChem, 10, 10)

	fixedClo.Put(btCloOK, 20, 80)
	fixedClo.Put(btCloClose, 80, 80)
	fixedClo.Put(entryClo, 0, 40)
	fixedClo.Put(labClo, 10, 10)

	fixedOth.Put(btOthOK, 20, 80)
	fixedOth.Put(btOthClose, 80, 80)
	fixedOth.Put(entryOth, 0, 40)
	fixedOth.Put(labOth, 10, 10)

	fixedBills.Put(btBillsOK, 20, 80)
	fixedBills.Put(btBillsClose, 80, 80)
	fixedBills.Put(entryBills, 0, 40)
	fixedBills.Put(labBills, 10, 10)

	popupSignIn.Add(fixedSignIn)
	popupSignUp.Add(fixedSignUp)
	popupFood.Add(fixedFood)
	popupClo.Add(fixedClo)
	popupChem.Add(fixedChem)
	popupBills.Add(fixedBills)
	popupOth.Add(fixedOth)

	fixedWarning.Put(labWarning, 10, 10)
	fixedWarning.Put(btWarningOK, 30, 80)
	popupWarning.Add(fixedWarning)

	fixedWarningGoal.Put(labWarningGoal, 10, 10)
	fixedWarningGoal.Put(btWarningGoalOk, 30, 80)
	popupWarningGoal.Add(fixedWarningGoal)

	win.Add(fixed)
	win.ShowAll()
}
