package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/rs/zerolog"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ozonva/ova-place-api/internal/event"
	"github.com/ozonva/ova-place-api/internal/flusher"
	"github.com/ozonva/ova-place-api/internal/metrics"
	"github.com/ozonva/ova-place-api/internal/models"
	"github.com/ozonva/ova-place-api/internal/producer"
	"github.com/ozonva/ova-place-api/internal/repo"
	desc "github.com/ozonva/ova-place-api/pkg/ova-place-api"
)

type api struct {
	desc.UnimplementedOvaPlaceApiV1Server
	repo       repo.Repo
	flusher    flusher.Flusher
	producer   producer.Producer
	cudCounter metrics.CudCounter
	logger     zerolog.Logger
}

// NewOvaPlaceAPI returns desc.OvaPlaceApiV1Server instance.
func NewOvaPlaceAPI(
	repo repo.Repo,
	flusher flusher.Flusher,
	producer producer.Producer,
	cudCounter metrics.CudCounter,
	logger zerolog.Logger,
) desc.OvaPlaceApiV1Server {
	return &api{
		repo:       repo,
		flusher:    flusher,
		producer:   producer,
		cudCounter: cudCounter,
		logger:     logger,
	}
}

func (a *api) CreatePlaceV1(
	ctx context.Context,
	req *desc.CreatePlaceRequestV1,
) (*desc.PlaceV1, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	a.logger.Debug().
		Uint64("UserId", req.UserId).
		Str("Seat", req.Seat).
		Str("Memo", req.Memo).
		Msg("Create place called")

	model := models.Place{
		Memo:   req.Memo,
		Seat:   req.Seat,
		UserID: req.UserId,
	}

	id, err := a.repo.AddEntity(ctx, model)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot AddEntity: %w", err)).Msg("Error from api")
		return nil, status.Error(codes.Internal, "internal error")
	}

	model.ID = id

	eventInstance, err := event.NewEvent("created", model)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot NewEvent: %w", err)).Msg("Error from api")
		return nil, status.Error(codes.Internal, "internal error")
	}

	err = a.producer.Push(ctx, "cud_events", eventInstance)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot Push: %w", err)).Msg("Error from api")
		return nil, status.Error(codes.Internal, "internal error")
	}

	a.cudCounter.SuccessfulCreates.Inc()

	return &desc.PlaceV1{
		PlaceId: id,
		UserId:  model.UserID,
		Seat:    model.Seat,
		Memo:    model.Memo,
	}, nil
}

func (a *api) MultiCreatePlaceV1(
	ctx context.Context,
	req *desc.MultiCreatePlaceRequestV1,
) (*desc.MultiCreatePlaceResponseV1, error) {

	span, ctx := opentracing.StartSpanFromContext(ctx, "multi_create_place")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	a.logger.Debug().
		Msg("Multi create place called")

	places := make([]models.Place, len(req.PlacesCreationData))

	for index := range req.PlacesCreationData {
		places[index] = models.Place{
			UserID: req.PlacesCreationData[index].UserId,
			Seat:   req.PlacesCreationData[index].Seat,
			Memo:   req.PlacesCreationData[index].Memo,
		}

		eventInstance, err := event.NewEvent("create", places[index])
		if err != nil {
			a.logger.Error().Err(fmt.Errorf("cannot NewEvent: %w", err)).Msg("Error from api")
			return nil, status.Error(codes.Internal, "internal error")
		}

		err = a.producer.Push(ctx, "cud_events", eventInstance)
		if err != nil {
			a.logger.Error().Err(fmt.Errorf("cannot Push: %w", err)).Msg("Error from api")
			return nil, status.Error(codes.Internal, "internal error")
		}
	}

	notSaved := a.flusher.Flush(ctx, places)

	notSavedPlaces := make([]*desc.CreatePlaceRequestV1, len(notSaved))

	for index := range notSaved {
		notSavedPlaces[index] = &desc.CreatePlaceRequestV1{
			UserId: notSaved[index].UserID,
			Seat:   notSaved[index].Seat,
			Memo:   notSaved[index].Memo,
		}
	}

	a.cudCounter.SuccessfulCreates.Add(float64(len(places) - len(notSaved)))

	return &desc.MultiCreatePlaceResponseV1{
		NotAdded: notSavedPlaces,
	}, nil
}

func (a *api) DescribePlaceV1(
	ctx context.Context,
	req *desc.DescribePlaceRequestV1,
) (*desc.PlaceV1, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	a.logger.Debug().
		Uint64("PlaceId", req.PlaceId).
		Msg("Describe place called")

	place, err := a.repo.DescribeEntity(ctx, req.PlaceId)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot DescribeEntity: %w", err)).Msg("Error from api")
		return nil, mapErrors(err)
	}

	return &desc.PlaceV1{
		PlaceId: req.PlaceId,
		UserId:  place.UserID,
		Seat:    place.Seat,
		Memo:    place.Memo,
	}, nil
}

func (a *api) ListPlacesV1(
	ctx context.Context,
	req *desc.ListPlacesRequestV1,
) (*desc.ListPlacesResponseV1, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	a.logger.Debug().
		Uint64("Page", req.Page).
		Uint64("PerPage", req.PerPage).
		Msg("List place called")

	totalCount, err := a.repo.TotalCount(ctx)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot TotalCount: %w", err)).Msg("Error from api")
		return nil, status.Error(codes.Internal, "internal error")
	}

	fetched, err := a.repo.ListEntities(ctx, req.PerPage, req.PerPage*(req.Page-1))
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot ListEntities: %w", err)).Msg("Error from api")
		return nil, status.Error(codes.Internal, "internal error")
	}

	places := make([]*desc.PlaceV1, len(fetched))

	for index := range fetched {
		places[index] = &desc.PlaceV1{
			PlaceId: fetched[index].ID,
			UserId:  fetched[index].UserID,
			Seat:    fetched[index].Seat,
			Memo:    fetched[index].Memo,
		}
	}

	return &desc.ListPlacesResponseV1{
		Places: places,
		Pagination: &desc.PaginationV1{
			Page:    req.Page,
			PerPage: req.PerPage,
			Total:   totalCount,
		},
	}, nil
}

func (a *api) UpdatePlaceV1(
	ctx context.Context,
	req *desc.UpdatePlaceRequestV1,
) (*desc.PlaceV1, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	a.logger.Debug().
		Uint64("UserId", req.UserId).
		Uint64("PlaceId", req.PlaceId).
		Str("Seat", req.Seat).
		Str("Memo", req.Memo).
		Msg("Update place called")

	model := models.Place{
		Memo:   req.Memo,
		Seat:   req.Seat,
		UserID: req.UserId,
	}

	err := a.repo.UpdateEntity(ctx, req.PlaceId, model)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot UpdateEntity: %w", err)).Msg("Error from api")
		return nil, mapErrors(err)
	}

	eventInstance, err := event.NewEvent("updated", model)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot NewEvent: %w", err)).Msg("Error from api")
		return nil, status.Error(codes.Internal, "internal error")
	}

	err = a.producer.Push(ctx, "cud_events", eventInstance)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot Push: %w", err)).Msg("Error from api")
		return nil, status.Error(codes.Internal, "internal error")
	}

	a.cudCounter.SuccessfulUpdates.Inc()

	return &desc.PlaceV1{
		PlaceId: req.PlaceId,
		UserId:  model.UserID,
		Seat:    model.Seat,
		Memo:    model.Memo,
	}, nil
}

func (a *api) RemovePlaceV1(
	ctx context.Context,
	req *desc.RemovePlaceRequestV1,
) (*emptypb.Empty, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	a.logger.Debug().
		Uint64("PlaceId", req.PlaceId).
		Msg("Remove place called")

	model, err := a.repo.DescribeEntity(ctx, req.PlaceId)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot DescribeEntity: %w", err)).Msg("Error from api")
		return nil, mapErrors(err)
	}

	err = a.repo.RemoveEntity(ctx, req.PlaceId)
	if err != nil {
		a.logger.Error().Err(err).Msg("Error from api")
		return nil, mapErrors(err)
	}

	eventInstance, err := event.NewEvent("deleted", *model)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot NewEvent: %w", err)).Msg("Error from api")
		return nil, status.Error(codes.Internal, "internal error")
	}

	err = a.producer.Push(ctx, "cud_events", eventInstance)
	if err != nil {
		a.logger.Error().Err(fmt.Errorf("cannot Push: %w", err)).Msg("Error from api")
		return nil, status.Error(codes.Internal, "internal error")
	}

	a.cudCounter.SuccessfulDeletes.Inc()

	return &emptypb.Empty{}, nil
}

func mapErrors(err error) error {
	if errors.Is(err, &repo.NotFound{}) {
		return status.Error(codes.NotFound, "not found")
	}
	return status.Error(codes.Internal, "internal error")
}
