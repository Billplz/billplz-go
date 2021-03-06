# billplz-go

Go package to use [Billplz API](https://billplz.com/api).

## Setup

Import this package
```go
import (
  billplz "github.com/billplz/billplz-go"
  // models "github.com/billplz/billplz-go/models"
)
```
On the root folder, run `dep ensure`

## Updating the package

On the root folder, run `dep ensure -update`

## Usage & Examples

For POST requests, make sure to import the models from the package.

```go
import (
  "fmt"
  billplz "github.com/billplz/billplz-go"
  models "github.com/billplz/billplz-go/models"
)
```

### 0.0 Init
There are 2 variables need for the init function, which is `ENVIRONMENT` and `APIKEY`.

You can get your key from your Billplz dashboard.
`ENVIRONMENT = staging | production`. Setting the correct environment is crucial as Billplz uses different urls for each environment.

```go
billplz.Init(ENVIRONMENT, APIKEY)
```

### 1.0 Collections
#### 1.1 Get Collection
Reference: [API V4 Get Collection](https://www.billplz.com/api#get-a-collection37)

```go
billplz.Init("staging", "69da23bf-da10-4fda-814d-3ad970035d38")

r, err := billplz.GetCollection("ei3a6mdl")
if len(err.Error.Type) > 0 {
  fmt.Printf("%+v\n", err)
  return
}
fmt.Printf("%+v\n", r)

```
#### 1.2 Create Collection
Reference: [API V4 Create Collection](https://www.billplz.com/api#get-a-collection36)

`Collection` struct:

| Parameters        | Data Type         | Required          |
| ----------------- | ----------------- | ----------------- |
| Title             | string            | true              |
| SplitHeader       | string            | false             |
| SplitPayments     | array             | false             |

`SplitPayment` struct:

| Parameters        | Type              | Required          |
| ----------------- | ----------------- | ----------------- |
| Email             | string            | true              |
| FixedCut          | integer           | false             |
| VariableCut       | string            | false             |
| StackOrder        | integer           | true              |

```go
billplz.Init("staging", "69da23bf-da10-4fda-814d-3ad970035d38")
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

r, err := billplz.CreateCollection(data)
if len(err.Error.Type) > 0 {
  fmt.Printf("%+v\n", err)
  return
}
fmt.Printf("%+v\n", r)
```

### 2.0 Bills
#### 2.1 Create Bill
Reference: [API V3 Create Bill](https://www.billplz.com/api#create-a-bill)

`Bill` struct:

| Parameters        | Type              | Required          |
| ----------------- | ----------------- | ----------------- |
| CollectionId      | string            | true              |
| Email             | string            | true              |
| Mobile            | string            | true              |
| Name              | string            | true              |
| Amount            | integer           | true              |
| CallbackUrl       | string            | true              |
| Description       | string            | true              |
| DueAt             | string            | true              |
| RedirectUrl       | integer           | false             |
| Reference1Label   | string            | false             |
| Reference1        | string            | false             |
| Reference2Label   | string            | false             |
| Reference2        | string            | false             |

```go
billplz.Init("staging", "69da23bf-da10-4fda-814d-3ad970035d38")
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

resp, err := billplz.CreateBill(data)
if len(err.Error.Type) > 0 {
  fmt.Printf("%+v\n", err)
  return
}
fmt.Printf("%+v\n", resp)
```
#### 2.2 Get Bill
Reference: [API V3 Get Bill](https://www.billplz.com/api#v3-get-a-bill)

```go
billplz.Init("staging", "69da23bf-da10-4fda-814d-3ad970035d38")

resp, err := billplz.GetBill("0npozuf0")
if len(err.Error.Type) > 0 {
  fmt.Printf("%+v\n", err)
  return
}
fmt.Printf("%+v\n", resp)
```

### 3.0 Payout Collections
#### 3.1 Create PayoutCollection
Reference: [API V4 Create Payout Collection](https://www.billplz.com/api#create-a-payout-collection)

`PayoutCollection` struct:

| Parameters        | Type              | Required          |
| ----------------- | ----------------- | ----------------- |
| Title             | string            | true              |

```go
billplz.Init("staging", "69da23bf-da10-4fda-814d-3ad970035d38")
data := models.PayoutCollection{
  Title: "Collection for Payout",
}

fmt.Println(billplz.CreatePayoutCollection(data))
```
#### 3.2 Get Payout Collection
Reference: [API V4 Get Payout Collection](https://www.billplz.com/api#get-a-payout-collection)

```go
billplz.Init("staging", "69da23bf-da10-4fda-814d-3ad970035d38")
fmt.Println(billplz.GetPayoutCollection(PAYOUT_COLLECTION_ID))
```
### 4.0 Banks
#### 4.1 Get FPX Banks
Reference: [API V3 Get FPX Banks](https://www.billplz.com/api#get-fpx-banks)

This list can be used to bypass Billplz's bill page.

```go
billplz.Init("staging", "69da23bf-da10-4fda-814d-3ad970035d38")
fmt.Println(billplz.GetFpxBanks())
```

#### 4.2 Get Bank Verification Status
Reference: [API V3 Get Bank Account](https://www.billplz.com/api#get-a-bank-account)

```go
billplz.Init("staging", "69da23bf-da10-4fda-814d-3ad970035d38")
fmt.Println(billplz.GetBankVerification(BANK_ACCOUNT_NO))
```
#### 4.3 Verify A New Bank Account
Reference: [API V3 Create Bank Account](https://www.billplz.com/api#create-a-bank-account)

This process is required for Mass Payment Instructions (Payout). In order to create a successful payout instruction, each recipient's account should be verified.

`Bank` struct:

| Parameters        | Type              | Required          |
| ----------------- | ----------------- | ----------------- |
| Name              | string            | true              |
| IdNo              | string            | true              |
| AccNo             | string            | true              |
| Code              | string            | true              |
| Organization      | boolean           | true              |

```go
billplz.Init("staging", "69da23bf-da10-4fda-814d-3ad970035d38")
data := models.Bank{
  Name: "Insan Jaya",
  IdNo: "91234567890",
  AccNo: "999988887777",
  Code: "MBBEMYKL",
  Organization: true
}

fmt.Println(billplz.CreateBankVerfication(data))
```
## Reference
[Billplz's API Documentation](https://billplz.com/api)

## Support / Issues
Facebook: [Billplz Dev Jam](https://www.facebook.com/groups/billplzdevjam/)
