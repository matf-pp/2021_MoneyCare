package main

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

func showBalance(service *services.SpendingService, label *gtk.Label) {

	spent := service.FindUsersSpending(UserID)
	s := fmt.Sprint("Potrosili ste: ", spent)
	label.SetText(s)
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
		log.Fatal("Unable to create label:", err)
	}
	return entry
}

func setupGui() {
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

	labBalance := setupLabel("")
	labUpID := setupLabel("ID: ")
	labUpIncome := setupLabel("PRIHODI: ")
	labUpOutgoings := setupLabel("RASHODI: ")
	labUpGoal := setupLabel("CILJ: ")
	labFood := setupLabel("Unesite iznos u dinarima: ")
	labChem := setupLabel("Unesite iznos u dinarima: ")
	labOth := setupLabel("Unesite iznos u dinarima: ")
	labClo := setupLabel("Unesite iznos u dinarima: ")
	labBills := setupLabel("Unesite iznos u dinarima: ")

	entry := setupEntry()
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

		spendingService := admin.SpendingService
		showBalance(spendingService, labBalance)

		popupSignIn.Hide()
	})

	btFoodOK := setupBtn("OK", func() {
		addSpending("food", entryFood)
		spendingService := admin.SpendingService
		showBalance(spendingService, labBalance)
		popupFood.Hide()
	})
	btChemOK := setupBtn("OK", func() {
		addSpending("chem", entryChem)
		spendingService := admin.SpendingService
		showBalance(spendingService, labBalance)
		popupChem.Hide()
	})
	btCloOK := setupBtn("OK", func() {
		addSpending("clothes", entryClo)
		spendingService := admin.SpendingService
		showBalance(spendingService, labBalance)
		popupClo.Hide()
	})
	btOthOK := setupBtn("OK", func() {
		addSpending("other", entryOth)
		spendingService := admin.SpendingService
		showBalance(spendingService, labBalance)
		popupOth.Hide()
	})
	btBillsOK := setupBtn("OK", func() {
		addSpending("bills", entryBills)
		spendingService := admin.SpendingService
		showBalance(spendingService, labBalance)
		popupBills.Hide()
	})

	sb, err := gtk.SpinButtonNewWithRange(0, 1, 0.1)
	if err != nil {
		log.Fatal("Unable to create spin button:", err)
	}
	pbFood, err := gtk.ProgressBarNew()
	if err != nil {
		log.Fatal("Unable to create progress bar:", err)
	}

	pbChem, err := gtk.ProgressBarNew()
	if err != nil {
		log.Fatal("Unable to create progress bar:", err)
	}
	pbClo, err := gtk.ProgressBarNew()
	if err != nil {
		log.Fatal("Unable to create progress bar:", err)
	}
	pbBill, err := gtk.ProgressBarNew()
	if err != nil {
		log.Fatal("Unable to create progress bar:", err)
	}
	pbOth, err := gtk.ProgressBarNew()
	if err != nil {
		log.Fatal("Unable to create progress bar:", err)
	}

	pbOth.SetOpacity(0.1)

	pbEntry, err := gtk.ProgressBarNew()
	if err != nil {
		log.Fatal("Unable to create progress bar:", err)
	}

	fixed.Put(sb, 450, 580)
	fixed.Put(pbFood, 100, 350)
	fixed.Put(pbChem, 600, 350)
	fixed.Put(pbClo, 200, 550)
	fixed.Put(pbBill, 500, 550)
	fixed.Put(pbOth, 350, 200)
	fixed.Put(pbEntry, 340, 420)
	fixed.Put(btSignUp, 10, 10)
	fixed.Put(btSignIn, 110, 10)
	fixed.Put(btFood, 100, 300)
	fixed.Put(btChem, 600, 300)
	fixed.Put(btClo, 200, 500)
	fixed.Put(btBill, 500, 500)
	fixed.Put(btOth, 350, 150)
	fixed.Put(labBalance, 350, 300)
	fixed.Put(entry, 335, 380)
	//fixed.Put(popupIn,12,30)

	fixedSignIn.Put(btSignInOK, 10, 70)
	fixedSignIn.Put(entryIn, 0, 30)

	fixedSignUp.Put(entryUpID, 100, 10)
	fixedSignUp.Put(entryUpIncome, 100, 50)
	fixedSignUp.Put(entryUpOutgoings, 100, 90)
	fixedSignUp.Put(entryUpGoal, 100, 130)
	fixedSignUp.Put(labUpID, 0, 10)
	fixedSignUp.Put(labUpIncome, 0, 50)
	fixedSignUp.Put(labUpOutgoings, 0, 90)
	fixedSignUp.Put(labUpGoal, 0, 130)
	fixedSignUp.Put(btSignUpOK, 50, 160)

	fixedFood.Put(btFoodOK, 50, 80)
	fixedFood.Put(entryFood, 0, 40)
	fixedFood.Put(labFood, 10, 10)

	fixedChem.Put(btChemOK, 50, 80)
	fixedChem.Put(entryChem, 0, 40)
	fixedChem.Put(labChem, 10, 10)

	fixedClo.Put(btCloOK, 50, 80)
	fixedClo.Put(entryClo, 0, 40)
	fixedClo.Put(labClo, 10, 10)

	fixedOth.Put(btOthOK, 50, 80)
	fixedOth.Put(entryOth, 0, 40)
	fixedOth.Put(labOth, 10, 10)

	fixedBills.Put(btBillsOK, 50, 80)
	fixedBills.Put(entryBills, 0, 40)
	fixedBills.Put(labBills, 10, 10)

	sb.Connect("value-changed", func(sb *gtk.SpinButton, pb *gtk.ProgressBar) {
		pb.SetFraction(sb.GetValue() / 1)
	}, pbFood)

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
