//go:build !race
// +build !race

package saver_test

import (
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozonva/ova-place-api/internal/models"
	"github.com/ozonva/ova-place-api/internal/saver"
	"github.com/ozonva/ova-place-api/mocks"
)

var _ = Describe("Saver", func() {

	var (
		flusherMock *mocks.MockFlusher
		places      []models.Place
	)

	BeforeEach(func() {
		ctrl := gomock.NewController(GinkgoT())
		flusherMock = mocks.NewMockFlusher(ctrl)
		defer ctrl.Finish()

		places = []models.Place{
			{UserID: 1, Memo: "Aero", Seat: "24G"},
			{UserID: 1, Memo: "Bus", Seat: "34"},
		}
	})

	Describe("Save places", func() {
		Context("close method is called before the saving timeout", func() {
			It("saving occurs on close method calling", func() {
				saverInstance := saver.NewSaver(2, time.Second*2, flusherMock)

				flusherMock.EXPECT().Flush(gomock.Eq(places[0:2])).Return([]models.Place{})

				Expect(saverInstance.Save(places[0])).To(BeNil())
				Expect(saverInstance.Save(places[1])).To(BeNil())
				Expect(saverInstance.Close()).To(BeNil())
			})
		})

		Context("close method is called after first saving timeout", func() {
			It("saving one element occurs by timeout", func() {
				saverInstance := saver.NewSaver(2, time.Millisecond*1, flusherMock)

				flusherMock.EXPECT().Flush(gomock.Eq(places[0:1])).Return([]models.Place{})
				flusherMock.EXPECT().Flush(gomock.Eq(places[1:2])).Return([]models.Place{})

				Expect(saverInstance.Save(places[0])).To(BeNil())
				time.Sleep(time.Millisecond * 4)
				Expect(saverInstance.Save(places[1])).To(BeNil())
				Expect(saverInstance.Close()).To(BeNil())
			})
		})

		Context("the buffer is full", func() {
			It("second save should return an error", func() {
				saverInstance := saver.NewSaver(1, time.Second*2, flusherMock)

				flusherMock.EXPECT().Flush(gomock.Eq(places[0:1])).Return([]models.Place{})

				Expect(saverInstance.Save(places[0])).To(BeNil())
				Expect(saverInstance.Save(places[1])).To(Not(BeNil()))
				Expect(saverInstance.Close()).To(BeNil())
			})
		})

		Context("when closed, there are unsaved entities", func() {
			It("close method calling should return an error", func() {
				saverInstance := saver.NewSaver(2, time.Second*2, flusherMock)

				flusherMock.EXPECT().Flush(gomock.Eq(places[0:2])).Return(places[0:1])

				Expect(saverInstance.Save(places[0])).To(BeNil())
				Expect(saverInstance.Save(places[1])).To(BeNil())
				Expect(saverInstance.Close()).To(Not(BeNil()))
			})
		})
	})
})
