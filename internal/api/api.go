package api

import (
	"context"
	"errors"

	"github.com/ozonva/ova-place-api/internal/models"
	"github.com/ozonva/ova-place-api/internal/repo"
	desc "github.com/ozonva/ova-place-api/pkg/ova-place-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type api struct {
	desc.UnimplementedOvaPlaceApiV1Server
	repo repo.Repo
}

func NewOvaPlaceApi(repo repo.Repo) desc.OvaPlaceApiV1Server {
	return &api{repo: repo}
}

func (a *api) CreatePlaceV1(
	ctx context.Context,
	req *desc.CreatePlaceRequestV1,
) (*desc.PlaceV1, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Uint64("UserId", req.UserId).
		Str("Seat", req.Seat).
		Str("Memo", req.Memo).
		Msg("Create place called")

	model := models.Place{
		Memo:   req.Memo,
		Seat:   req.Seat,
		UserID: req.UserId,
	}

	id, err := a.repo.AddEntity(model)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &desc.PlaceV1{
		PlaceId: id,
		UserId:  model.UserID,
		Seat:    model.Seat,
		Memo:    model.Memo,
	}, nil
}

func (a *api) DescribePlaceV1(
	ctx context.Context,
	req *desc.DescribePlaceRequestV1,
) (*desc.PlaceV1, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Uint64("PlaceId", req.PlaceId).
		Msg("Describe place called")

	place, err := a.repo.DescribeEntity(req.PlaceId)
	if err != nil {
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

	log.Debug().
		Uint64("Page", req.Page).
		Uint64("PerPage", req.PerPage).
		Msg("List place called")

	totalCount, err := a.repo.TotalCount()
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	fetched, err := a.repo.ListEntities(req.PerPage, req.PerPage*(req.Page-1))
	if err != nil {
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

	log.Debug().
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

	err := a.repo.UpdateEntity(req.PlaceId, model)
	if err != nil {
		return nil, mapErrors(err)
	}

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

	log.Debug().
		Uint64("PlaceId", req.PlaceId).
		Msg("Remove place called")

	err := a.repo.RemoveEntity(req.PlaceId)
	if err != nil {
		return nil, mapErrors(err)
	}

	return &emptypb.Empty{}, nil
}

func mapErrors(err error) error {
	if errors.Is(err, &repo.NotFound{}) {
		return status.Error(codes.NotFound, "not found")
	}
	return status.Error(codes.Internal, "internal error")
}
