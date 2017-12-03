package main

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// publishableKey := os.Getenv("PUBLISHABLE_KEY")
	stripe.Key = os.Getenv("SECRET_KEY")
	charge_description := os.Getenv("CHARGE_DESCRIPTION")
	currency := os.Getenv("CURRENCY")
	if currency == "" { currency = "usd" }


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Submit a Stripe charge token to /charge")
	})

	http.HandleFunc("/charge", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		switch r.Method {
		case "OPTIONS":

		case "POST":
			r.ParseForm()

			fmt.Println(r.Form)
			fmt.Println(r.Form.Get("email"))

			customerParams := &stripe.CustomerParams{Email: r.Form.Get("email")}
			customerParams.SetSource(r.Form.Get("token"))

			newCustomer, err := customer.New(customerParams)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			amount, err := strconv.Atoi(r.Form.Get("amount"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			btcaddress, err := strconv.QuoteToASCII(r.Form.Get("btcaddress"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "BTC Address is : ", r.Form.Get("btcaddress"))

			chargeParams := &stripe.ChargeParams{
				Amount:   uint64(amount),
				Currency: stripe.Currency(currency),
				Desc:     charge_description,
				Customer: newCustomer.ID,
			}

			if _, err := charge.New(chargeParams); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "Charge completed successfully!")
		}
	})

	port := os.Getenv("PORT")
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
