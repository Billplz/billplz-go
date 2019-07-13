![Deploy to now](https://deploy.now.sh/static/button.svg)

# billplz-go

Go package to use [Billplz API](https://billplz.com/api).

## Setup

Import this package
```
import (
  billplz "github.com/helmiruza/billplz-go"
  // models "github.com/helmiruza/billplz-go/models"
)
```
On the root folder, run `dep ensure`

## Updating the package

On the root folder, run `dep ensure -update`

## Usage & Examples

For POST requests, make sure to import the models from the package.

```
import (
  "fmt"
  billplz "github.com/helmiruza/billplz-go"
  models "github.com/helmiruza/billplz-go/models"
)
```

### 0.0 Init
There are 2 variables need for the init function, which is `ENVIRONMENT` and `APIKEY`. You can get your key from your Billplz dashboard.
`ENVIRONMENT = staging | production`. Setting the correct environment is crucial as Billplz uses different urls for each environment.

```
billplz.Init(ENVIRONMENT, APIKEY)
```
### 1.0 Collections
#### 1.1 Get Collection

```
billplz.Init("staging", "69da23bf-da10-4fda-814d-3ad970035d38")
fmt.Println(billplz.GetCollection("ei3a6mdl"))
```
#### 1.2 Create Collection

```
data := models.Collection{
  Title: "test 4",
  SplitHeader: false,
  SplitPayments: []models.SplitPayment{
    {
      Email: "abc@gmail.com",
      FixedCut: 100,
      StackOrder: 0,
    },
    {
      Email: "abc@gmail.com",
      FixedCut: 200,
      StackOrder: 1,
    },
  },
}

fmt.Println(billplz.CreateCollection(data))
```

### 2.0 Bills
#### 2.1 Create Bill

```
data := models.Bill{
  CollectionId: "gaem247h",
  Email: "helmiruza@gmail.com",
  Mobile: "60193102400",
  Name: "Helmi Ruza",
  Amount: 10000,
  CallbackUrl: "https://billplz.com",
  Description: "Test Bill 123",
  DueAt: "2019-07-12",
}

fmt.Println(billplz.CreateBill(data))
```
#### 2.2 Get Bill
```
fmt.Println(billplz.GetBill("0npozuf0"))
```

### 3.0 Payout Collections
#### 3.1 Create PayoutCollection
```
data := models.PayoutCollection{
  Title: "Collection for Payout",
}

fmt.Println(billplz.CreatePayoutCollection(data))
```
#### 3.2 Get Payout Collection
```
fmt.Println(billplz.GetPayoutCollection(PAYOUT_COLLECTION_ID))
```
### 4.0 Banks
#### 4.1 Get FPX Banks
```
fmt.Println(billplz.GetFpxBanks())
```
#### 4.2 Get Bank Verification
```
fmt.Println(billplz.GetBankVerification(BANK_ACCOUNT_NO))
```
#### 4.2 Verify An Account Number
```
data := models.Bank{
  Name: "Insan Jaya",
  IdNo: "91234567890",
  AccNo: "999988887777",
  Code: "MBBEMYKL",
  Organization: true
}
 
fmt.Println(billplz.CreateBankVerfication(data))
```
