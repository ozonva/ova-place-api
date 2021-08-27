package api

import (
	"context"
	desc "github.com/ozonva/ova-place-api/pkg/ova-place-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOvaPlaceApiV1Server
}

func NewOvaPlaceApi() desc.OvaPlaceApiV1Server {
	return &api{}
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

	return &desc.PlaceV1{}, nil
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

	return &desc.PlaceV1{}, nil
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

	return &desc.ListPlacesResponseV1{}, nil
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

	return &desc.PlaceV1{}, nil
}

func (a *api) RemovePlaceV1(
	ctx context.Context,
	req *desc.RemovePlaceRequestV1,
) (*desc.EmptyV1, error) {

	if err := req.Validate(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	log.Debug().
		Uint64("PlaceId", req.PlaceId).
		Msg("Remove place called")

	return &desc.EmptyV1{}, nil
}
