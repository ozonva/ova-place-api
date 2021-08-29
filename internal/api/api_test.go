package api_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/ozonva/ova-place-api/internal/api"
	"github.com/ozonva/ova-place-api/internal/models"
	"github.com/ozonva/ova-place-api/mocks"
	desc "github.com/ozonva/ova-place-api/pkg/ova-place-api"
)

var _ = Describe("Api", func() {
	var (
		repoMock *mocks.MockRepo
	)

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		repoMock = mocks.NewMockRepo(ctrl)
		defer ctrl.Finish()
	})

	Describe("Creates place", func() {
		Context("all is ok", func() {
			It("should return place", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				place := models.Place{UserID: 1, Memo: "Aero", Seat: "24G"}

				repoMock.EXPECT().AddEntity(gomock.Eq(place)).Return(uint64(1), nil)

				ctx := context.Background()

				request := desc.CreatePlaceRequestV1{
					UserId: place.UserID, Memo: place.Memo, Seat: place.Seat,
				}

				response, err := apiInstance.CreatePlaceV1(ctx, &request)

				gomega.Expect(response.UserId == place.UserID).To(gomega.BeTrue())
				gomega.Expect(response.Memo == place.Memo).To(gomega.BeTrue())
				gomega.Expect(response.Seat == place.Seat).To(gomega.BeTrue())
				gomega.Expect(response.PlaceId == uint64(1)).To(gomega.BeTrue())
				gomega.Expect(err == nil).To(gomega.BeTrue())
			})
		})

		Context("repo returns error", func() {
			It("should return internal error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				place := models.Place{UserID: 1, Memo: "Aero", Seat: "24G"}

				repoMock.EXPECT().AddEntity(gomock.Eq(place)).Return(uint64(0), errors.New("test error"))

				ctx := context.Background()

				request := desc.CreatePlaceRequestV1{
					UserId: place.UserID, Memo: place.Memo, Seat: place.Seat,
				}

				_, err := apiInstance.CreatePlaceV1(ctx, &request)

				gomega.Expect(err.Error() == "rpc error: code = Internal desc = internal error").To(gomega.BeTrue())
			})
		})
	})

	Describe("Describes place", func() {
		Context("all is ok", func() {
			It("should return place", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				place := models.Place{ID: 1, UserID: 1, Memo: "Aero", Seat: "24G"}

				repoMock.EXPECT().DescribeEntity(gomock.Eq(place.ID)).Return(&place, nil)

				ctx := context.Background()

				request := desc.DescribePlaceRequestV1{
					PlaceId: place.ID,
				}

				response, err := apiInstance.DescribePlaceV1(ctx, &request)

				gomega.Expect(response.UserId == place.UserID).To(gomega.BeTrue())
				gomega.Expect(response.Memo == place.Memo).To(gomega.BeTrue())
				gomega.Expect(response.Seat == place.Seat).To(gomega.BeTrue())
				gomega.Expect(response.PlaceId == place.ID).To(gomega.BeTrue())
				gomega.Expect(err == nil).To(gomega.BeTrue())
			})
		})

		Context("repo returns error", func() {
			It("should return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				place := models.Place{ID: 1, UserID: 1, Memo: "Aero", Seat: "24G"}

				repoMock.EXPECT().DescribeEntity(gomock.Eq(place.ID)).Return(nil, errors.New("test error"))

				ctx := context.Background()

				request := desc.DescribePlaceRequestV1{
					PlaceId: place.ID,
				}

				_, err := apiInstance.DescribePlaceV1(ctx, &request)

				gomega.Expect(err.Error() == "rpc error: code = Internal desc = internal error").To(gomega.BeTrue())
			})
		})

		Context("row not found", func() {
			It("should return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				place := models.Place{ID: 1, UserID: 1, Memo: "Aero", Seat: "24G"}

				repoMock.EXPECT().DescribeEntity(gomock.Eq(place.ID)).Return(nil, errors.New("sql: no rows in result set"))

				ctx := context.Background()

				request := desc.DescribePlaceRequestV1{
					PlaceId: place.ID,
				}

				_, err := apiInstance.DescribePlaceV1(ctx, &request)

				gomega.Expect(err.Error() == "rpc error: code = NotFound desc = not found").To(gomega.BeTrue())
			})
		})
	})

	Describe("Lists places", func() {
		Context("all is ok", func() {
			It("should return places", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				place := models.Place{ID: 1, UserID: 1, Memo: "Aero", Seat: "24G"}

				gomock.InOrder(
					repoMock.EXPECT().TotalCount().Return(uint64(1), nil),
					repoMock.EXPECT().ListEntities(uint64(1), uint64(0)).Return([]models.Place{place}, nil),
				)

				ctx := context.Background()

				request := desc.ListPlacesRequestV1{
					Page: 1, PerPage: 1,
				}

				response, err := apiInstance.ListPlacesV1(ctx, &request)

				gomega.Expect(response.Places[0].UserId == place.UserID).To(gomega.BeTrue())
				gomega.Expect(response.Places[0].Memo == place.Memo).To(gomega.BeTrue())
				gomega.Expect(response.Places[0].Seat == place.Seat).To(gomega.BeTrue())
				gomega.Expect(response.Places[0].PlaceId == place.ID).To(gomega.BeTrue())
				gomega.Expect(response.Pagination.Page == uint64(1)).To(gomega.BeTrue())
				gomega.Expect(response.Pagination.PerPage == uint64(1)).To(gomega.BeTrue())
				gomega.Expect(response.Pagination.Total == uint64(1)).To(gomega.BeTrue())
				gomega.Expect(err == nil).To(gomega.BeTrue())
			})
		})

		Context("repo returns error", func() {
			It("should return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				gomock.InOrder(
					repoMock.EXPECT().TotalCount().Return(uint64(1), nil),
					repoMock.EXPECT().ListEntities(uint64(1), uint64(0)).Return(nil, errors.New("test error")),
				)

				ctx := context.Background()

				request := desc.ListPlacesRequestV1{
					Page: 1, PerPage: 1,
				}

				_, err := apiInstance.ListPlacesV1(ctx, &request)

				gomega.Expect(err.Error() == "rpc error: code = Internal desc = internal error").To(gomega.BeTrue())
			})
		})
	})

	Describe("Updates place", func() {
		Context("all is ok", func() {
			It("should return place", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				place := models.Place{ID: 1, UserID: 1, Memo: "Aero", Seat: "24G"}

				repoMock.EXPECT().UpdateEntity(gomock.Eq(place.ID), gomock.Any()).Return(nil)

				ctx := context.Background()

				request := desc.UpdatePlaceRequestV1{
					PlaceId: place.ID,
					UserId:  place.UserID,
					Seat:    place.Seat,
					Memo:    place.Memo,
				}

				response, err := apiInstance.UpdatePlaceV1(ctx, &request)

				gomega.Expect(response.UserId == place.UserID).To(gomega.BeTrue())
				gomega.Expect(response.Memo == place.Memo).To(gomega.BeTrue())
				gomega.Expect(response.Seat == place.Seat).To(gomega.BeTrue())
				gomega.Expect(response.PlaceId == place.ID).To(gomega.BeTrue())
				gomega.Expect(err == nil).To(gomega.BeTrue())
			})
		})

		Context("repo returns error", func() {
			It("should return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				place := models.Place{ID: 1, UserID: 1, Memo: "Aero", Seat: "24G"}

				repoMock.EXPECT().UpdateEntity(gomock.Eq(place.ID), gomock.Any()).Return(errors.New("test error"))

				ctx := context.Background()

				request := desc.UpdatePlaceRequestV1{
					PlaceId: place.ID,
					UserId:  place.UserID,
					Seat:    place.Seat,
					Memo:    place.Memo,
				}

				_, err := apiInstance.UpdatePlaceV1(ctx, &request)

				gomega.Expect(err.Error() == "rpc error: code = Internal desc = internal error").To(gomega.BeTrue())
			})
		})

		Context("row not found", func() {
			It("should return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				place := models.Place{ID: 1, UserID: 1, Memo: "Aero", Seat: "24G"}

				repoMock.EXPECT().UpdateEntity(gomock.Eq(place.ID), gomock.Any()).Return(errors.New("no rows affected"))

				ctx := context.Background()

				request := desc.UpdatePlaceRequestV1{
					PlaceId: place.ID,
					UserId:  place.UserID,
					Seat:    place.Seat,
					Memo:    place.Memo,
				}

				_, err := apiInstance.UpdatePlaceV1(ctx, &request)

				gomega.Expect(err.Error() == "rpc error: code = NotFound desc = not found").To(gomega.BeTrue())
			})
		})
	})

	Describe("Deletes place", func() {
		Context("all is ok", func() {
			It("should not return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock)

				place := models.Place{ID: 1, UserID: 1, Memo: "Aero", Seat: "24G"}

				repoMock.EXPECT().RemoveEntity(gomock.Eq(place.ID)).Return(nil)

				ctx := context.Background()

				request := desc.RemovePlaceRequestV1{
					PlaceId: place.ID,
				}

				_, err := apiInstance.RemovePlaceV1(ctx, &request)

				gomega.Expect(err == nil).To(gomega.BeTrue())
			})
		})
	})

	Context("repo returns error", func() {
		It("should return error", func() {
			apiInstance := api.NewOvaPlaceApi(repoMock)

			place := models.Place{ID: 1, UserID: 1, Memo: "Aero", Seat: "24G"}

			repoMock.EXPECT().RemoveEntity(gomock.Eq(place.ID)).Return(errors.New("test error"))

			ctx := context.Background()

			request := desc.RemovePlaceRequestV1{
				PlaceId: place.ID,
			}

			_, err := apiInstance.RemovePlaceV1(ctx, &request)

			gomega.Expect(err.Error() == "rpc error: code = Internal desc = internal error").To(gomega.BeTrue())
		})
	})

	Context("row not found", func() {
		It("should return error", func() {
			apiInstance := api.NewOvaPlaceApi(repoMock)

			place := models.Place{ID: 1, UserID: 1, Memo: "Aero", Seat: "24G"}

			repoMock.EXPECT().RemoveEntity(gomock.Eq(place.ID)).Return(errors.New("no rows affected"))

			ctx := context.Background()

			request := desc.RemovePlaceRequestV1{
				PlaceId: place.ID,
			}

			_, err := apiInstance.RemovePlaceV1(ctx, &request)

			gomega.Expect(err.Error() == "rpc error: code = NotFound desc = not found").To(gomega.BeTrue())
		})
	})

})
