package product

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/option"
)

// Get get google play iab product
func Get(packageNameID string, skuID string) *androidpublisher.InAppProduct {
	ctx := context.Background()

	jsonFile := "./service-account.json" // You must prepare service-account to authentication on goolge cloud.
	androidpublisherService, err := androidpublisher.NewService(ctx, option.WithCredentialsFile(jsonFile))
	if err != nil {
		log.Fatal(err)
	}

	call := androidpublisherService.Inappproducts.Get(packageNameID, skuID)

	product, err := call.Do()
	fmt.Printf("%+v", product)
	fmt.Printf("%+v", product.Prices)
	fmt.Printf("%+v", product.DefaultPrice)

	if err != nil {
		log.Fatal(err)
	}

	println("create successful :" + product.Sku)
	return product
}

// Insert insert iab-product to GooglePlay
// TODO: insert impl
func Insert(packageNameID string, skuID string) error {
	//	call := androidpublisherService.Inappproducts.Insert(packageNameID, product)
	//	&{DefaultLanguage:ja-JP DefaultPrice:0xc000184730 GracePeriod:P1W Listings:map[ja-JP:{Description:test Title:test ForceSendFields:[] NullFields:[]}] PackageName:****** Prices:map[] PurchaseType:subscription Season:<nil> Sku:test Status:active SubscriptionPeriod:P1M TrialPeriod:P5D

	return nil
}
