package gui

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"main/src/admin"
	"main/src/services"
)

var Username string
var UserID primitive.ObjectID

func addSpending(cat string, entry *gtk.Entry) {
	categoryService := admin.CategoryService
	categoryId, err := categoryService.FindOne(cat)
	if err != nil {
		panic(err)
	}
	spendingService := admin.SpendingService
	spendingService.InsertFromEntry(UserID, categoryId.ID, entry)
}

func showBalance(service *services.SpendingService, label *gtk.Label, pb *gtk.ProgressBar) {
	spent := service.FindUsersSpending(UserID)
	s := fmt.Sprint("Potrosili ste: ", spent)
	label.SetText(s)

	userService := admin.UserService
	us, err := userService.FindOne(Username)
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
	spent := service.FindUsersSpendingByCategory(UserID, catID.ID)
	fmt.Println(spent)
	s := fmt.Sprint("Potrosili ste: ", spent)
	label.SetText(s)

	userService := admin.UserService
	us, err := userService.FindOne(Username)
	if err != nil {
		panic(err)
	}
	total := us.Income
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

func SetupGui() {
	win := setupWindow("Money Care")
	fixed := setupFixed()
	fixedSignIn := setupFixed()
	fixedSignUp := setupFixed()
	fixedFood := setupFixed()
	fixedChem := setupFixed()
	fixedClo := setupFixed()
	fixedOth := setupFixed()
	fixedBills := setupFixed()

	popupSignIn := setupPopup(150, 120, "SignIn")
	popupSignUp := setupPopup(150, 120, "SignUp")
	popupFood := setupPopup(150, 120, "Food")
	popupChem := setupPopup(150, 120, "Chem")
	popupOth := setupPopup(150, 120, "Other")
	popupClo := setupPopup(150, 120, "Clothes")
	popupBills := setupPopup(150, 120, "Bills")

	btSignUp := setupBtn("SIGN UP", func() {
		popupSignUp.ShowAll()
	})
	btSignIn := setupBtn("SIGN IN", func() {
		popupSignIn.ShowAll()
	})
	btFood := setupBtn("FOOD", func() {
		popupFood.ShowAll()
	})
	btChem := setupBtn("CHEM/COSM", func() {
		popupChem.ShowAll()
	})
	btClo := setupBtn("CLOTHES", func() {
		popupClo.ShowAll()
	})
	btBill := setupBtn("BILLS", func() {
		popupBills.ShowAll()
	})
	btOth := setupBtn("OTHER", func() {
		popupOth.ShowAll()
	})
	btHist := setupBtn("HISTOGRAM", func() {

	})
	btPieChar := setupBtn("PIECHART", func() {

	})
	btGraph := setupBtn("GRAPH", func() {

	})

	labBalance := setupLabel("balance")
	labelFoodEx := setupLabel("balance")
	labelBillsEx := setupLabel("balance")
	labelChemEx := setupLabel("balance")
	labelOtherEx := setupLabel("balance")
	labelClothesEx := setupLabel("balance")

	labUpID := setupLabel("ID: ")
	labUpIncome := setupLabel("PRIHODI: ")
	labUpOutgoings := setupLabel("RASHODI: ")
	labUpGoal := setupLabel("CILJ: ")
	labFood := setupLabel("Unesite iznos u dinarima: ")
	labChem := setupLabel("Unesite iznos u dinarima: ")
	labOth := setupLabel("Unesite iznos u dinarima: ")
	labClo := setupLabel("Unesite iznos u dinarima: ")
	labBills := setupLabel("Unesite iznos u dinarima: ")

	pbFood := setupProgressBar()
	pbChem := setupProgressBar()
	pbClo := setupProgressBar()
	pbBill := setupProgressBar()
	pbOth := setupProgressBar()
	pb := setupProgressBar()

	//entry := setupEntry()
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
		userService := admin.UserService
		userService.InsertFromEntry(entryUpID, entryUpGoal, entryUpIncome, entryUpOutgoings)
		uname, err := entryUpID.GetText()
		if err != nil {
			panic(err)
		}
		Username = uname
		user, err := userService.FindOne(uname)
		UserID = user.ID

		popupSignUp.Hide()
	})

	btSignInOK := setupBtn("OK", func() {
		userService := admin.UserService
		userService.FindOneFromEntry(entryIn)
		uname, err := entryIn.GetText()
		if err != nil {
			panic(err)
		}
		Username = uname
		user, err := userService.FindOne(uname)
		UserID = user.ID

		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)

		popupSignIn.Hide()
	})

	btFoodOK := setupBtn("OK", func() {
		addSpending("food", entryFood)
		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)

		popupFood.Hide()
	})
	btChemOK := setupBtn("OK", func() {
		addSpending("chem", entryChem)
		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
		popupChem.Hide()
	})
	btCloOK := setupBtn("OK", func() {
		addSpending("clothes", entryClo)
		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
		popupClo.Hide()
	})
	btOthOK := setupBtn("OK", func() {
		addSpending("other", entryOth)
		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
		popupOth.Hide()
	})
	btBillsOK := setupBtn("OK", func() {
		addSpending("bills", entryBills)
		showBalanceForAll(labBalance, labelFoodEx, labelClothesEx, labelChemEx, labelOtherEx, labelBillsEx, pb, pbFood, pbClo, pbChem, pbOth, pbBill)
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
	fixed.Put(pb,470,500 )

	fixed.Put(btSignUp, 10, 10)
	fixed.Put(btSignIn, 110, 10)
	fixed.Put(btFood, 100, 100)
	fixed.Put(btChem, 100, 200)
	fixed.Put(btClo, 100, 300)
	fixed.Put(btBill, 100, 400)
	fixed.Put(btOth, 100, 500)
	fixed.Put(btHist,640,10)
	fixed.Put(btPieChar,532,10)
	fixed.Put(btGraph,441,10)

	fixed.Put(labBalance, 520, 450)
	//fixed.Put(entry, 335, 380)
	//fixed.Put(popupIn,12,30)

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
	fixedFood.Put(btFoodClose,80,80)
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

	win.Add(fixed)
	win.ShowAll()
}
