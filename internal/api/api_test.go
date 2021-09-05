package api_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"

	"github.com/ozonva/ova-place-api/internal/api"
	"github.com/ozonva/ova-place-api/internal/metrics"
	"github.com/ozonva/ova-place-api/internal/models"
	"github.com/ozonva/ova-place-api/internal/repo"
	"github.com/ozonva/ova-place-api/mocks"
	desc "github.com/ozonva/ova-place-api/pkg/ova-place-api"
)

var _ = Describe("Api", func() {
	var (
		repoMock     *mocks.MockRepo
		flusherMock  *mocks.MockFlusher
		producerMock *mocks.MockProducer
		counterMock  *mocks.MockCounter
		cudCounter   metrics.CudCounter
	)

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		repoMock = mocks.NewMockRepo(ctrl)
		flusherMock = mocks.NewMockFlusher(ctrl)
		producerMock = mocks.NewMockProducer(ctrl)
		counterMock = mocks.NewMockCounter(ctrl)
		cudCounter = metrics.NewCudCounter(counterMock, counterMock, counterMock)
		defer ctrl.Finish()
	})

	Describe("Creates place", func() {
		Context("all is ok", func() {
			It("should return place", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				place := models.Place{
					UserID: 1,
					Memo:   "Aero",
					Seat:   "24G",
				}

				gomock.InOrder(
					repoMock.EXPECT().AddEntity(gomock.Eq(place)).Return(uint64(1), nil),
					producerMock.EXPECT().Push(gomock.Eq("cud_events"), gomock.Any()),
					counterMock.EXPECT().Inc(),
				)

				ctx := context.TODO()

				request := desc.CreatePlaceRequestV1{
					UserId: place.UserID,
					Memo:   place.Memo,
					Seat:   place.Seat,
				}

				response, err := apiInstance.CreatePlaceV1(ctx, &request)

				gomega.Expect(response.UserId).To(gomega.Equal(place.UserID))
				gomega.Expect(response.Memo).To(gomega.Equal(place.Memo))
				gomega.Expect(response.Seat).To(gomega.Equal(place.Seat))
				gomega.Expect(response.PlaceId).To(gomega.Equal(uint64(1)))
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		Context("repo returns error", func() {
			It("should return internal error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				place := models.Place{
					UserID: 1,
					Memo:   "Aero",
					Seat:   "24G",
				}

				repoMock.EXPECT().AddEntity(gomock.Eq(place)).Return(uint64(0), errors.New("test error"))

				ctx := context.TODO()

				request := desc.CreatePlaceRequestV1{
					UserId: place.UserID,
					Memo:   place.Memo,
					Seat:   place.Seat,
				}

				_, err := apiInstance.CreatePlaceV1(ctx, &request)

				gomega.Expect(err.Error()).To(gomega.Equal("rpc error: code = Internal desc = internal error"))
			})
		})
	})

	Describe("Multi creates place", func() {
		Context("all is ok", func() {
			It("should not return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				places := []models.Place{
					{UserID: 1, Memo: "Aero", Seat: "24G"},
					{UserID: 1, Memo: "Bus", Seat: "34"},
				}

				ctx := context.TODO()

				gomock.InOrder(
					producerMock.EXPECT().Push(gomock.Eq("cud_events"), gomock.Any()).Times(2),
					flusherMock.EXPECT().Flush(gomock.Any(), gomock.Eq(places)).Return([]models.Place{}),
					counterMock.EXPECT().Add(float64(2)),
				)

				request := desc.MultiCreatePlaceRequestV1{
					PlacesCreationData: []*desc.CreatePlaceRequestV1{
						{
							UserId: places[0].UserID,
							Memo:   places[0].Memo,
							Seat:   places[0].Seat,
						},
						{
							UserId: places[1].UserID,
							Memo:   places[1].Memo,
							Seat:   places[1].Seat,
						},
					},
				}

				_, err := apiInstance.MultiCreatePlaceV1(ctx, &request)

				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		Context("flusher returns not saved places", func() {
			It("should return not saved places", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				places := []models.Place{
					{UserID: 1, Memo: "Aero", Seat: "24G"},
					{UserID: 1, Memo: "Bus", Seat: "34"},
				}

				gomock.InOrder(
					producerMock.EXPECT().Push(gomock.Eq("cud_events"), gomock.Any()).Times(2),
					flusherMock.EXPECT().Flush(gomock.Any(), gomock.Eq(places)).Return([]models.Place{
						places[1],
					}),
					counterMock.EXPECT().Add(float64(1)),
				)

				ctx := context.TODO()

				request := desc.MultiCreatePlaceRequestV1{
					PlacesCreationData: []*desc.CreatePlaceRequestV1{
						{
							UserId: places[0].UserID,
							Memo:   places[0].Memo,
							Seat:   places[0].Seat,
						},
						{
							UserId: places[1].UserID,
							Memo:   places[1].Memo,
							Seat:   places[1].Seat,
						},
					},
				}

				response, err := apiInstance.MultiCreatePlaceV1(ctx, &request)

				gomega.Expect(response.NotAdded[0].UserId).To(gomega.Equal(places[1].UserID))
				gomega.Expect(response.NotAdded[0].Memo).To(gomega.Equal(places[1].Memo))
				gomega.Expect(response.NotAdded[0].Seat).To(gomega.Equal(places[1].Seat))
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	Describe("Describes place", func() {
		Context("all is ok", func() {
			It("should return place", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				place := models.Place{
					ID:     1,
					UserID: 1,
					Memo:   "Aero",
					Seat:   "24G",
				}

				repoMock.EXPECT().DescribeEntity(gomock.Eq(place.ID)).Return(&place, nil)

				ctx := context.TODO()

				request := desc.DescribePlaceRequestV1{
					PlaceId: place.ID,
				}

				response, err := apiInstance.DescribePlaceV1(ctx, &request)

				gomega.Expect(response.UserId).To(gomega.Equal(place.UserID))
				gomega.Expect(response.Memo).To(gomega.Equal(place.Memo))
				gomega.Expect(response.Seat).To(gomega.Equal(place.Seat))
				gomega.Expect(response.PlaceId).To(gomega.Equal(place.ID))
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		Context("repo returns error", func() {
			It("should return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				place := models.Place{
					ID:     1,
					UserID: 1,
					Memo:   "Aero",
					Seat:   "24G",
				}

				repoMock.EXPECT().DescribeEntity(gomock.Eq(place.ID)).Return(nil, errors.New("test error"))

				ctx := context.TODO()

				request := desc.DescribePlaceRequestV1{
					PlaceId: place.ID,
				}

				_, err := apiInstance.DescribePlaceV1(ctx, &request)

				gomega.Expect(err.Error()).To(gomega.Equal("rpc error: code = Internal desc = internal error"))
			})
		})

		Context("row not found", func() {
			It("should return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				place := models.Place{
					ID:     1,
					UserID: 1,
					Memo:   "Aero",
					Seat:   "24G",
				}

				repoMock.EXPECT().DescribeEntity(gomock.Eq(place.ID)).Return(nil, &repo.NotFound{})

				ctx := context.TODO()

				request := desc.DescribePlaceRequestV1{
					PlaceId: place.ID,
				}

				_, err := apiInstance.DescribePlaceV1(ctx, &request)

				gomega.Expect(err.Error()).To(gomega.Equal("rpc error: code = NotFound desc = not found"))
			})
		})
	})

	Describe("Lists places", func() {
		Context("all is ok", func() {
			It("should return places", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				place := models.Place{
					ID:     1,
					UserID: 1,
					Memo:   "Aero",
					Seat:   "24G",
				}

				gomock.InOrder(
					repoMock.EXPECT().TotalCount().Return(uint64(1), nil),
					repoMock.EXPECT().ListEntities(uint64(1), uint64(0)).Return([]models.Place{place}, nil),
				)

				ctx := context.TODO()

				request := desc.ListPlacesRequestV1{
					Page:    1,
					PerPage: 1,
				}

				response, err := apiInstance.ListPlacesV1(ctx, &request)

				gomega.Expect(response.Places[0].UserId).To(gomega.Equal(place.UserID))
				gomega.Expect(response.Places[0].Memo).To(gomega.Equal(place.Memo))
				gomega.Expect(response.Places[0].Seat).To(gomega.Equal(place.Seat))
				gomega.Expect(response.Places[0].PlaceId).To(gomega.Equal(place.ID))
				gomega.Expect(response.Pagination.Page).To(gomega.Equal(uint64(1)))
				gomega.Expect(response.Pagination.PerPage).To(gomega.Equal(uint64(1)))
				gomega.Expect(response.Pagination.Total).To(gomega.Equal(uint64(1)))
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		Context("repo returns error", func() {
			It("should return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				gomock.InOrder(
					repoMock.EXPECT().TotalCount().Return(uint64(1), nil),
					repoMock.EXPECT().ListEntities(uint64(1), uint64(0)).Return(nil, errors.New("test error")),
				)

				ctx := context.TODO()

				request := desc.ListPlacesRequestV1{
					Page:    1,
					PerPage: 1,
				}

				_, err := apiInstance.ListPlacesV1(ctx, &request)

				gomega.Expect(err.Error()).To(gomega.Equal("rpc error: code = Internal desc = internal error"))
			})
		})
	})

	Describe("Updates place", func() {
		Context("all is ok", func() {
			It("should return place", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				place := models.Place{
					ID:     1,
					UserID: 1,
					Memo:   "Aero",
					Seat:   "24G",
				}

				gomock.InOrder(
					repoMock.EXPECT().UpdateEntity(gomock.Eq(place.ID), gomock.Any()).Return(nil),
					producerMock.EXPECT().Push(gomock.Eq("cud_events"), gomock.Any()),
					counterMock.EXPECT().Inc(),
				)

				ctx := context.TODO()

				request := desc.UpdatePlaceRequestV1{
					PlaceId: place.ID,
					UserId:  place.UserID,
					Seat:    place.Seat,
					Memo:    place.Memo,
				}

				response, err := apiInstance.UpdatePlaceV1(ctx, &request)

				gomega.Expect(response.UserId).To(gomega.Equal(place.UserID))
				gomega.Expect(response.Memo).To(gomega.Equal(place.Memo))
				gomega.Expect(response.Seat).To(gomega.Equal(place.Seat))
				gomega.Expect(response.PlaceId).To(gomega.Equal(place.ID))
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		Context("repo returns error", func() {
			It("should return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				place := models.Place{
					ID:     1,
					UserID: 1,
					Memo:   "Aero",
					Seat:   "24G",
				}

				repoMock.EXPECT().UpdateEntity(gomock.Eq(place.ID), gomock.Any()).Return(errors.New("test error"))

				ctx := context.TODO()

				request := desc.UpdatePlaceRequestV1{
					PlaceId: place.ID,
					UserId:  place.UserID,
					Seat:    place.Seat,
					Memo:    place.Memo,
				}

				_, err := apiInstance.UpdatePlaceV1(ctx, &request)

				gomega.Expect(err.Error()).To(gomega.Equal("rpc error: code = Internal desc = internal error"))
			})
		})

		Context("row not found", func() {
			It("should return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				place := models.Place{
					ID:     1,
					UserID: 1,
					Memo:   "Aero",
					Seat:   "24G",
				}

				repoMock.EXPECT().UpdateEntity(gomock.Eq(place.ID), gomock.Any()).Return(&repo.NotFound{})

				ctx := context.TODO()

				request := desc.UpdatePlaceRequestV1{
					PlaceId: place.ID,
					UserId:  place.UserID,
					Seat:    place.Seat,
					Memo:    place.Memo,
				}

				_, err := apiInstance.UpdatePlaceV1(ctx, &request)

				gomega.Expect(err.Error()).To(gomega.Equal("rpc error: code = NotFound desc = not found"))
			})
		})
	})

	Describe("Deletes place", func() {
		Context("all is ok", func() {
			It("should not return error", func() {
				apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

				place := models.Place{
					ID:     1,
					UserID: 1,
					Memo:   "Aero",
					Seat:   "24G",
				}

				gomock.InOrder(
					repoMock.EXPECT().DescribeEntity(gomock.Eq(place.ID)).Return(&place, nil),
					repoMock.EXPECT().RemoveEntity(gomock.Eq(place.ID)).Return(nil),
					producerMock.EXPECT().Push(gomock.Eq("cud_events"), gomock.Any()),
					counterMock.EXPECT().Inc(),
				)

				ctx := context.TODO()

				request := desc.RemovePlaceRequestV1{
					PlaceId: place.ID,
				}

				_, err := apiInstance.RemovePlaceV1(ctx, &request)

				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	Context("repo returns error", func() {
		It("should return error", func() {
			apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

			place := models.Place{
				ID:     1,
				UserID: 1,
				Memo:   "Aero",
				Seat:   "24G",
			}

			repoMock.EXPECT().DescribeEntity(gomock.Eq(place.ID)).Return(nil, errors.New("test error"))

			ctx := context.TODO()

			request := desc.RemovePlaceRequestV1{
				PlaceId: place.ID,
			}

			_, err := apiInstance.RemovePlaceV1(ctx, &request)

			gomega.Expect(err.Error()).To(gomega.Equal("rpc error: code = Internal desc = internal error"))
		})
	})

	Context("row not found", func() {
		It("should return error", func() {
			apiInstance := api.NewOvaPlaceApi(repoMock, flusherMock, producerMock, cudCounter)

			place := models.Place{
				ID:     1,
				UserID: 1,
				Memo:   "Aero",
				Seat:   "24G",
			}

			gomock.InOrder(
				repoMock.EXPECT().DescribeEntity(gomock.Eq(place.ID)).Return(nil, &repo.NotFound{}),
			)

			ctx := context.TODO()

			request := desc.RemovePlaceRequestV1{
				PlaceId: place.ID,
			}

			_, err := apiInstance.RemovePlaceV1(ctx, &request)

			gomega.Expect(err.Error()).To(gomega.Equal("rpc error: code = NotFound desc = not found"))
		})
	})

})
