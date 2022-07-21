// Code generated by proroc-gen-graphql, DO NOT EDIT.
package paymentv1alpha1

import (
	"github.com/graphql-go/graphql"
	money "google.golang.org/genproto/googleapis/type/money"
)

var (
	gql__enum_PaymentProvider *graphql.Enum        // enum PaymentProvider in payment/v1alpha1/payment.proto
	gql__type_Order           *graphql.Object      // message Order in payment/v1alpha1/payment.proto
	gql__input_Order          *graphql.InputObject // message Order in payment/v1alpha1/payment.proto
)

func Gql__enum_PaymentProvider() *graphql.Enum {
	if gql__enum_PaymentProvider == nil {
		gql__enum_PaymentProvider = graphql.NewEnum(graphql.EnumConfig{
			Name: "Paymentv1Alpha1_Enum_PaymentProvider",
			Values: graphql.EnumValueConfigMap{
				"PAYMENT_PROVIDER_UNSPECIFIED": &graphql.EnumValueConfig{
					Value: PaymentProvider(0),
				},
				"PAYMENT_PROVIDER_STRIPE": &graphql.EnumValueConfig{
					Value: PaymentProvider(1),
				},
				"PAYMENT_PROVIDER_PAYPAL": &graphql.EnumValueConfig{
					Value: PaymentProvider(2),
				},
				"PAYMENT_PROVIDER_APPLE": &graphql.EnumValueConfig{
					Value: PaymentProvider(3),
				},
			},
		})
	}
	return gql__enum_PaymentProvider
}

func Gql__type_Order() *graphql.Object {
	if gql__type_Order == nil {
		gql__type_Order = graphql.NewObject(graphql.ObjectConfig{
			Name:        "Paymentv1Alpha1_Type_Order",
			Description: `Order represents a monetary order.`,
			Fields: graphql.Fields{
				"order_id": &graphql.Field{
					Type: graphql.String,
				},
				"recipient_id": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"amount": &graphql.Field{
					Type: graphql.NewNonNull(money.Gql__type_Money()),
				},
				"payment_provider": &graphql.Field{
					Type: graphql.NewNonNull(paymentv1alpha1.Gql__enum_PaymentProvider()),
				},
			},
		})
	}
	return gql__type_Order
}

func Gql__input_Order() *graphql.InputObject {
	if gql__input_Order == nil {
		gql__input_Order = graphql.NewInputObject(graphql.InputObjectConfig{
			Name: "Paymentv1Alpha1_Input_Order",
			Fields: graphql.InputObjectConfigFieldMap{
				"order_id": &graphql.InputObjectFieldConfig{
					Type: graphql.String,
				},
				"recipient_id": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"amount": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(money.Gql__input_Money()),
				},
				"payment_provider": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(paymentv1alpha1.Gql__enum_PaymentProvider()),
				},
			},
		})
	}
	return gql__input_Order
}
