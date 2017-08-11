package main

import (
  "fmt"
  "github.com/stripe/stripe-go"
  "github.com/stripe/stripe-go/charge"
  "github.com/stripe/stripe-go/customer"
  "net/http"
  "os"
)

func main() {
  // publishableKey := os.Getenv("PUBLISHABLE_KEY")
  stripe.Key = os.Getenv("SECRET_KEY")

	http.HandleFunc("/charge", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		customerParams := &stripe.CustomerParams{Email: r.Form.Get("stripeEmail")}
		customerParams.SetSource(r.Form.Get("stripeToken"))

		newCustomer, err := customer.New(customerParams)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		chargeParams := &stripe.ChargeParams{
			Amount:   500,
			Currency: "usd",
			Desc:     "Donation to Jacob",
			Customer: newCustomer.ID,
		}

		if _, err := charge.New(chargeParams); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Charge completed successfully!")
	})

	http.ListenAndServe(":5000", nil)
}