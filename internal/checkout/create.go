package checkout

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

// TODO: Receber na requisição o product_id
// TODO: Pegar os dados desse product_id na api de products
// TODO: Completar no CheckoutSessionParams as informações do produto
func Create(c *fiber.Ctx) error {
	stripe.Key = os.Getenv("STRIPE_KEY")
	domain := "http://localhost:8085"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("price_1PFpyBRvKFn3fK2YYTpl4l7q"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/success"),
		CancelURL:  stripe.String(domain + "/cancel"),
	}

	s, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
		return c.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return c.Redirect(s.URL, http.StatusSeeOther)

}
