package main

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"main/src/admin"
)

func setup_window(title string) *gtk.Window {
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
func setup_popup(width int, height int, title string) *gtk.Window {
	popup, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	popup.SetTitle(title)
	popup.SetDefaultSize(width, height)
	popup.SetPosition(gtk.WIN_POS_CENTER)
	return popup
}

func setup_btn(label string, onClick func()) *gtk.Button {
	btn, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	btn.Connect("clicked", onClick)
	return btn
}

func setup_fixed() *gtk.Fixed {
	fixed, err := gtk.FixedNew()
	if err != nil {
		log.Fatal("Unable to create GtkFixed:", err)
	}
	return fixed
}

func setup_label(text string) *gtk.Label {
	label, err := gtk.LabelNew(text)
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	return label
}

func setup_entry() *gtk.Entry {
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	return entry
}

func setup_gui() {
	win := setup_window("Money Care")

	fixed := setup_fixed()
	fixedSignIn := setup_fixed()
	fixedSignUp := setup_fixed()
	fixedFood := setup_fixed()
	fixedChem := setup_fixed()
	fixedClo := setup_fixed()
	fixedOth := setup_fixed()
	fixedBills := setup_fixed()

	popupSignIn := setup_popup(150, 120, "SignIn")
	popupSignUp := setup_popup(150, 120, "SignUp")
	popupFood := setup_popup(150, 120, "Food")
	popupChem := setup_popup(150, 120, "Chem")
	popupOth := setup_popup(150, 120, "Other")
	popupClo := setup_popup(150, 120, "Clothes")
	popupBills := setup_popup(150, 120, "Bills")

	btSignUp := setup_btn("SIGN UP", func() {
		popupSignUp.ShowAll()
	})
	btSignIn := setup_btn("SIGN IN", func() {
		popupSignIn.ShowAll()
	})
	btFood := setup_btn("FOOD", func() {
		popupFood.ShowAll()
	})
	btChem := setup_btn("CHEM/COSM", func() {
		popupChem.ShowAll()
	})
	btClo := setup_btn("CLOTHES", func() {
		popupClo.ShowAll()
	})
	btBill := setup_btn("BILLS", func() {
		popupBills.ShowAll()
	})
	btOth := setup_btn("OTHER", func() {
		popupOth.ShowAll()
	})

	labBalance := setup_label("Potrosili ste : 0 RSD")
	labUpId := setup_label("ID: ")
	labUpIncome := setup_label("PRIHODI: ")
	labUpOutgoings := setup_label("RASHODI: ")
	labUpGoal := setup_label("CILJ: ")
	labFood := setup_label("Unesite iznos u dinarima: ")
	labChem := setup_label("Unesite iznos u dinarima: ")
	labOth := setup_label("Unesite iznos u dinarima: ")
	labClo := setup_label("Unesite iznos u dinarima: ")
	labBills := setup_label("Unesite iznos u dinarima: ")

	entry := setup_entry()
	entryIn := setup_entry()
	entryUpId := setup_entry()
	entryUpOutgoings := setup_entry()
	entryUpIncome := setup_entry()
	entryUpGoal := setup_entry()
	entryFood := setup_entry()
	entryChem := setup_entry()
	entryOth := setup_entry()
	entryClo := setup_entry()
	entryBills := setup_entry()

	btSignUpOK := setup_btn("OK", func() {
		userService := admin.UserService
		userService.InsertFromEntry(entryUpId, entryUpGoal, entryUpIncome, entryUpOutgoings)

	})

	btSignInOK := setup_btn("OK", func() {

	})
	btFoodOK := setup_btn("OK", func() {
		popupFood.Hide()
	})
	btChemOK := setup_btn("OK", func() {
		popupChem.Hide()
	})
	btCloOK := setup_btn("OK", func() {
		popupClo.Hide()
	})
	btOthOK := setup_btn("OK", func() {
		popupOth.Hide()
	})
	btBillsOK := setup_btn("OK", func() {
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

	fixedSignUp.Put(entryUpId, 100, 10)
	fixedSignUp.Put(entryUpIncome, 100, 50)
	fixedSignUp.Put(entryUpOutgoings, 100, 90)
	fixedSignUp.Put(entryUpGoal, 100, 130)
	fixedSignUp.Put(labUpId, 0, 10)
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
