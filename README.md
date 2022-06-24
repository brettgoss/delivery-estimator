# delivery-estimator

Estimates package delivery time of day using historical data


## Contributing
1. `go run ./cmd/server/main.go`

## To Do

- Create data store
- Create data input method
  - Simple form with dropdown for postal code maybe?

## Data Structures

### Deliveries
```json
[
  {
    "expectedDelivery": null,
    "actualDelivery": "2022-03-04T1500",
    "deliveryType": "standardMail",
    "postalCode": "V1V1V1",
    "service": "canadaPost"
  },
  {
    "expectedDelivery": "2022-03-04",
    "actualDelivery": "2022-03-04T1500",
    "deliveryType": "expressPost",
    "postalCode": "V1V1V1",
    "service": "canadaPost"
  }
  {
    "expectedDelivery": "2022-03-05",
    "actualDelivery": "2022-03-04T1500",
    "deliveryType": "expressPost",
    "postalCode": "V1V1V1",
    "service": "UPS"
  }
]
```

### Delivery Locations
Keyed by postal code to reduce geocoding API usage

```json
{
    "V1V1V1" {
        "city": "Victoria",
        "province": "BC",
    }
}
```