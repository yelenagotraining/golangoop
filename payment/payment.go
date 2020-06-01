package payment

type PaymentOption interface {
	ProcessPayment(float32) bool
}
