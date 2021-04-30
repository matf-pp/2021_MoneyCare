package main

import (
	"github.com/gotk3/gotk3/gtk"
	_ "go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	//TODO: dodavanje troskova po kategorijama
	//TODO: kako azurirati i drugu kolekciju nakon transakcije
	//TODO: prikaz balansa: u sredini ekrana prikazuje se mesecni_rashodi/mesecni_prihodi
	//TODO: pri pozivu funkcije azuriraj_bazu, treba izbaciti upozorenje ako je premasena granica (po nekoj kategoriji ili ukupno) i treba da se zacrveni traka
	//TODO: kako na pocetku meseca sve resetovati na 0
	//TODO: mogucnost azuriranja stalnih rashoda, prihoda, cilja
	//TODO: novi korisnik -> unosi sva 4 i kreira se dokument za njega za tekuci mesec
	//TODO: na signIn provera meseca (da li ga ima u kolekciji) (kao signUp)

	gtk.Init(nil)
	setup_gui()
	gtk.Main()
}
