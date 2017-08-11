package kolpa

// PaymentCard function returns a random payment card number.
// A convenience function, same as g.GenericGenerator("payment_card").
// Sample Output: 24208918642114291
func (g *Generator) PaymentCard() string {
	return g.GenericGenerator("payment_card")
}

// PaymentCard function returns a random master card number.
// A convenience function, same as g.GenericGenerator("payment_master_card").
// Sample Output: 5154214251151243
func (g *Generator) MasterCard() string {
	return g.GenericGenerator("payment_mastercard")
}

// PaymentCard function returns a random master card number.
// A convenience function, same as g.GenericGenerator("payment_visa_card").
// Sample Output: 4292198623077593
func (g *Generator) VisaCard() string {
	return g.GenericGenerator("payment_visa_card")
}
