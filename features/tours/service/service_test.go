package service

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"
	"wanderer/features/airlines"
	"wanderer/features/facilities"
	"wanderer/features/locations"
	"wanderer/features/tours"
	"wanderer/features/tours/mocks"

	"github.com/stretchr/testify/assert"
)

func TestTourServiceCreate(t *testing.T) {
	repo := mocks.NewRepository(t)
	srv := NewTourService(repo)
	ctx := context.Background()

	data := tours.Tour{
		Title:       "Jepang Winter Golden Route & Mount Fuji",
		Description: "Everything feels extra spectacular in Dubai—from the ultra-modern Burj Khalifa to the souks and malls filled with gold and jewelry vendors. It`s a place where if you can dream it, you can do it: Whether that means skiing indoors, dune-surfing in the desert, or zip-lining above the city. But it`s not all glitz and adrenaline-pumping action. Stroll through the winding alleys of Al Fahidi Historical Neighborhood to see what Dubai was like during the mid-19th century. Or visit the Jumeirah Mosque (one of the few mosques open to non-Muslims) and learn about Emirati culture. Spot some street art on Jumeirah Beach Road and grab a bite at a shawarma shop, or spend the day hunting for spices and perfume then round things out with a Michelin-starred meal. You can really do it all and we`ve got more recs, below.",
		Price:       30000000,
		AdminFee:    5000,
		Discount:    10,
		Start:       time.Now(),
		Finish:      time.Now().Add(time.Hour * 48),
		Quota:       25,
		Thumbnail: tours.File{
			Raw: strings.NewReader("case image"),
		},
		Picture: []tours.File{
			{Raw: strings.NewReader("case image")},
			{Raw: strings.NewReader("case image")},
			{Raw: strings.NewReader("case image")},
		},
		Itinerary: []tours.Itinerary{
			{Location: "location 1", Description: "description 1"},
			{Location: "location 2", Description: "description 2"},
			{Location: "location 3", Description: "description 3"},
		},
		Facility: []facilities.Facility{
			{Id: 1},
			{Id: 2},
		},
		Airline: airlines.Airline{
			Id: 3,
		},
		Location: locations.Location{
			Id: 1,
		},
	}

	t.Run("invalid title", func(t *testing.T) {
		caseData := data
		caseData.Title = ""

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "title")
	})

	t.Run("invalid description", func(t *testing.T) {
		caseData := data
		caseData.Description = ""

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "description")
	})

	t.Run("invalid price", func(t *testing.T) {
		caseData := data
		caseData.Price = 0

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "price")
	})

	t.Run("invalid start date", func(t *testing.T) {
		caseData := data
		caseData.Start = time.Time{}

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "start date")
	})

	t.Run("invalid finish date", func(t *testing.T) {
		caseData := data
		caseData.Finish = time.Time{}

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "finish date")
	})

	t.Run("invalid quota", func(t *testing.T) {
		caseData := data
		caseData.Quota = 0

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "quota")
	})

	t.Run("invalid thumbnail", func(t *testing.T) {
		caseData := data
		caseData.Thumbnail.Raw = nil

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "thumbnail")
	})

	t.Run("invalid itinerary", func(t *testing.T) {
		caseData := data
		caseData.Itinerary = nil

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "itinerary")
	})

	t.Run("invalid location", func(t *testing.T) {
		caseData := data
		caseData.Location.Id = 0

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "location")
	})

	t.Run("invalid airline", func(t *testing.T) {
		caseData := data
		caseData.Airline.Id = 0

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "validate")
		assert.ErrorContains(t, err, "airline")
	})

	t.Run("error from repository", func(t *testing.T) {
		caseData := data

		repo.On("Create", ctx, caseData).Return(errors.New("some error from repository")).Once()

		err := srv.Create(ctx, caseData)

		assert.ErrorContains(t, err, "some error from repository")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		caseData := data

		repo.On("Create", ctx, caseData).Return(nil).Once()

		err := srv.Create(ctx, caseData)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}
