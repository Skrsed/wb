func trash() {
	payment := &domain.Payment{
		Transaction:  "testdata",
		RequestId:    "testdata",
		Currency:     "testdata",
		Provider:     "testdata",
		Amount:       12345,
		PaymentDt:    12345,
		Bank:         "testdata",
		DeliveryCost: 9876,
		GoodsTotal:   9876,
		CustomFee:    9876,
	}

	delivery := &domain.Delivery{
		ID:       "delivery details",
		Name:     "delivery details",
		Phone:    "delivery details",
		Zip:      "delivery details",
		City:     "delivery details",
		Address:  "delivery details",
		Region:   "delivery details",
		Email:    "delivery details",
		OrderUid: "delivery details",
	}

	items := make([]*domain.Item, 3)
	items[0] = &domain.Item{
		ID:          1,
		ChrtId:      2,
		TrackNumber: "item test 1",
		Price:       "item test 1",
		Rid:         "item test 1",
		Name:        "item test 1",
		Sale:        3,
		Size:        4,
		TotalPrice:  5,
		NmId:        6,
		Brand:       "item test 1",
		Status:      7,
	}
	items[1] = &domain.Item{
		ID:          1,
		ChrtId:      2,
		TrackNumber: "item test 2",
		Price:       "item test 2",
		Rid:         "item test 2",
		Name:        "item test 2",
		Sale:        3,
		Size:        4,
		TotalPrice:  5,
		NmId:        6,
		Brand:       "item test 2",
		Status:      7,
	}
	items[2] = &domain.Item{
		ID:          1,
		ChrtId:      2,
		TrackNumber: "item test 3",
		Price:       "item test 3",
		Rid:         "item test 3",
		Name:        "item test 3",
		Sale:        3,
		Size:        4,
		TotalPrice:  5,
		NmId:        6,
		Brand:       "item test 2",
		Status:      7,
	}

	order := &domain.Order{
		Uid:               "uuid string",
		TrackNumber:       "track number",
		Entry:             "some entity",
		Delivery:          *delivery,
		Payment:           *payment,
		Items:             items,
		Locale:            "en",
		InternalSignature: "signature",
		CustomerId:        "customer id",
		DeliveryService:   "service",
		Shardkey:          1,
		SmId:              2,
		DateCreated:       time.Now(),
		OofShard:          1,
	}

	pgRepository := pgRepository.NewPostgresRepository(db)
	order, err := pgRepository.CreateOrderCascade(context.Background(), order)

	fmt.Println(order)
	fmt.Println(err)
}