package main

import (
	"context"
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var ctx context.Context
var client *mongo.Client
var database *mongo.Database
var transactions *mongo.Collection
var users *mongo.Collection

func main() {

	ctx, _ = context.WithTimeout(context.Background(), 1000*time.Second)//zasto?
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://bogdanis:12345@cluster0.vu61o.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	database = client.Database("MoneyCare")
	//korisniciKolekcija := bazaPodataka.Collection("korisnici")
	transactions = database.Collection("transakcije")
	fmt.Println(transactions)

	//TODO: dodavanje troskova po kategorijama
	//dodaj_trosak(korisnik, kategorija, iznos, opis, datum)
	//		kreiramo novu transakciju i sacuvamo je u bazi (ima jedinstven ID transakcije)
	//nakon toga treba da se azurira celokupno stanje korisnika, tj. da se pozove neka funkcija refresh koja radi sa drugom kolekcijom iz baze
	//
	//TODO: kako azurirati i drugu kolekciju nakon transakcije
	//u trenutku transakcije imamo na raspolaganju korisnicko ime i mesec, pa mozemo da azuriramo dokument sa ta dva (po tome je jedinstveno, npr. bogdanis i april)
	//primer funkcije koja se poziva iz funkcije dodaj_trosak
	//azuriraj_bazu(korisnik, mesec, kategorija_koja_se_azurira, iznos):
	//		npr. za kategoriju hrana dodamo iznos na kolonu hrana i na kolonu rashodi
	//
	//TODO: prikaz balansa: u sredini ekrana prikazuje se mesecni_rashodi/mesecni_prihodi
	//potrebno je pozvati funkciju prikaz(rashodi, prihodi) iz funkcije azuriraj_bazu da bi se azuriralo stanje
	//
	//TODO: pri pozivu funkcije azuriraj_bazu, treba izbaciti upozorenje ako je premasena granica (po nekoj kategoriji ili ukupno) i treba da se zacrveni traka
	//
	//TODO: kako na pocetku meseca sve resetovati na 0
	//
	//TODO: mogucnost azuriranja stalnih rashoda, prihoda, cilja
	//
	//TODO: novi korisnik -> unosi sva 4 i kreira se dokument za njega za tekuci mesec

	//TODO: na signIn provera meseca (da li ga ima u kolekciji) (kao signUp)
	//
	//PROBLEM: na pocetku svakog meseca treba kreirati novi dokument za novi mesec inicijalizovan na 0


	gtk.Init(nil)
	setup_gui()
	//buttonFood.Connect("clicked", func(){
	//	fmt.Println("")
	//	e, err := entry.GetText()
	//	if err!=nil {
	//		panic("nzm")
	//	}
	//	price, err := strconv.ParseFloat(e, 64)
	//	user, err := eUser.GetText()
	//	transakcija1 := Transakcija {
	//		Korisnik: "",
	//		Kategorija: Kategorija{"hrana"},
	//		Opis: "nesto i mleko",
	//		Iznos: price,
	//		Datum: time.Now(),
	//	}
	//
	//	//popUp.ShowAll()
	//	dial.Run()
	//
	//	fmt.Println(transakcija1)
	//	_, err = transakcijeKolekcija.UpdateOne(ctx,
	//		bson.M{"korisnik": user},
	//		bson.D{
	//			{"$set", bson.D{{"iznos", price}}},
	//		},)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//
	//})

	gtk.Main()
}