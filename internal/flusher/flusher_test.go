package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozonva/ova-place-api/internal/flusher"
	"github.com/ozonva/ova-place-api/internal/models"
	"github.com/ozonva/ova-place-api/mocks"
)

var _ = Describe("Flusher", func() {
	var (
		repoMock *mocks.MockRepo
		places   []models.Place
	)
	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		repoMock = mocks.NewMockRepo(ctrl)
		defer ctrl.Finish()

		places = []models.Place{
			{UserID: 1, Memo: "Aero", Seat: "24G"},
			{UserID: 1, Memo: "Bus", Seat: "34"},
			{UserID: 1, Memo: "Train", Seat: "4 71"},
			{UserID: 1, Memo: "Aero", Seat: "34G"},
		}
	})

	Describe("Flush places", func() {
		Context("all is ok and the batch size is two", func() {
			It("should return nil", func() {
				flusherInstance := flusher.NewFlusher(2, repoMock)

				gomock.InOrder(
					repoMock.EXPECT().AddEntities(gomock.Eq(places[0:2])).Return(nil),
					repoMock.EXPECT().AddEntities(gomock.Eq(places[2:4])).Return(nil),
				)

				Expect(flusherInstance.Flush(places) == nil).To(BeTrue())
			})
		})

		Context("all is ok and the batch size is three", func() {
			It("should return nil", func() {
				flusherInstance := flusher.NewFlusher(3, repoMock)

				gomock.InOrder(
					repoMock.EXPECT().AddEntities(gomock.Eq(places[0:3])).Return(nil),
					repoMock.EXPECT().AddEntities(gomock.Eq(places[3:4])).Return(nil),
				)

				Expect(flusherInstance.Flush(places) == nil).To(BeTrue())
			})
		})

		Context("when batch size is invalid", func() {
			It("should return all inputted places", func() {
				flusherInstance := flusher.NewFlusher(0, repoMock)

				Expect(flusherInstance.Flush(places)).To(Equal(places))
			})
		})

		Context("repo returns error", func() {
			It("should return all inputted places", func() {
				flusherInstance := flusher.NewFlusher(2, repoMock)

				gomock.InOrder(
					repoMock.EXPECT().AddEntities(gomock.Eq(places[0:2])).Return(errors.New("some error")),
					repoMock.EXPECT().AddEntities(gomock.Eq(places[2:4])).Return(errors.New("some error")),
				)

				Expect(flusherInstance.Flush(places)).To(Equal(places))
			})
		})

		Context("one of the repo calls returns an error", func() {
			It("should return all inputted places", func() {
				flusherInstance := flusher.NewFlusher(2, repoMock)

				gomock.InOrder(
					repoMock.EXPECT().AddEntities(gomock.Eq(places[0:2])).Return(errors.New("some error")),
					repoMock.EXPECT().AddEntities(gomock.Eq(places[2:4])).Return(nil),
				)

				Expect(flusherInstance.Flush(places)).To(Equal(places[0:2]))
			})
		})
	})
})
