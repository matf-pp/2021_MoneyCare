package main

import (
	"github.com/gotk3/gotk3/gtk"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"main/src/admin"
	_ "main/src/services"
)

var currentuser int
var xs [30]float64

func mean(xs [30]float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
	//TODO: dodavanje troskova po kategorijama
	//TODO: kako azurirati i drugu kolekciju nakon transakcije
	//TODO: prikaz balansa: u sredini ekrana prikazuje se mesecni_rashodi/mesecni_prihodi
	//TODO: pri pozivu funkcije azuriraj_bazu, treba izbaciti upozorenje ako je premasena granica (po nekoj kategoriji ili ukupno) i treba da se zacrveni traka
	//TODO: kako na pocetku meseca sve resetovati na 0
	//TODO: mogucnost azuriranja stalnih rashoda, prihoda, cilja
	//TODO: novi korisnik -> unosi sva 4 i kreira se dokument za njega za tekuci mesec
	//TODO: na signIn provera meseca (da li ga ima u kolekciji) (kao signUp)

	admin.SetupSeed()
	gtk.Init(nil)
	SetupGui()
	gtk.Main()

	if EntryUpIncomeAmount != -1 {
		if EntryUpIncomeAmount <= 50000 {
			currentuser = 0
		}
		if EntryUpIncomeAmount >= 50000 && EntryUpIncomeAmount <= 15000 {
			currentuser = 1
		}
		if EntryUpIncomeAmount >= 150000 {
			currentuser = 2
		}

	}

	currentusermean := make(map[string]float64)
	idfood, _ := admin.CategoryService.FindOne("food")
	idclot, _ := admin.CategoryService.FindOne("clothes")
	idchem, _ := admin.CategoryService.FindOne("chem")
	idother, _ := admin.CategoryService.FindOne("other")
	idbills, _ := admin.CategoryService.FindOne("bills")

	if currentuser == 0 {

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID0[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idfood.ID)
		}

		currentusermean["food"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID0[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idclot.ID)
		}

		currentusermean["clothes"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID0[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idchem.ID)
		}

		currentusermean["chem"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID0[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idbills.ID)
		}

		currentusermean["bills"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID0[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idother.ID)
		}

		currentusermean["othes"] = mean(xs)

	}

	if currentuser == 1 {

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID1[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idfood.ID)
		}

		currentusermean["food"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID1[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idclot.ID)
		}

		currentusermean["clothes"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID1[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idchem.ID)
		}

		currentusermean["chem"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID1[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idbills.ID)
		}

		currentusermean["bills"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID1[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idother.ID)
		}

		currentusermean["othes"] = mean(xs)

	}

	if currentuser == 2 {

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID2[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idfood.ID)
		}

		currentusermean["food"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID2[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idclot.ID)
		}

		currentusermean["clothes"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID2[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idchem.ID)
		}

		currentusermean["chem"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID2[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idbills.ID)
		}

		currentusermean["bills"] = mean(xs)

		for i := 0; i < 5; i++ {
			idp, _ := admin.UserService.FindOne(admin.UsersID2[i])
			xs[i] = admin.SpendingService.FindUsersSpendingByCategory(idp.ID, idother.ID)
		}

		currentusermean["othes"] = mean(xs)

	}

	//plots.PieChart()
}
