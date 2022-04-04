package services

import (
	"context"
	"fmt"
	"github.com/multycloud/multy/api/converter"
	"github.com/multycloud/multy/api/deploy"
	"github.com/multycloud/multy/api/errors"
	"github.com/multycloud/multy/api/proto/commonpb"
	"github.com/multycloud/multy/api/proto/configpb"
	"github.com/multycloud/multy/api/proto/resourcespb"
	"github.com/multycloud/multy/api/util"
	"github.com/multycloud/multy/db"
	"google.golang.org/protobuf/proto"
)

type CreateRequest[Arg proto.Message] interface {
	GetResource() Arg
	proto.Message
}

type UpdateRequest[Arg proto.Message] interface {
	GetResource() Arg
	WithResourceId
}

type WithResourceId interface {
	GetResourceId() string
	proto.Message
}

type Service[Arg proto.Message, OutT proto.Message] struct {
	Db         *db.Database
	Converters converter.ResourceConverters[Arg, OutT]
}

func WrappingErrors[InT any, OutT any](f func(context.Context, InT) (OutT, error)) func(context.Context, InT) (OutT, error) {
	return func(ctx context.Context, in InT) (OutT, error) {
		out, err := f(ctx, in)
		if err != nil {
			return out, errors.InternalServerError(err)
		}
		return out, nil
	}
}

func (s Service[Arg, OutT]) Create(ctx context.Context, in CreateRequest[Arg]) (OutT, error) {
	return WrappingErrors(s.create)(ctx, in)
}

func (s Service[Arg, OutT]) create(ctx context.Context, in CreateRequest[Arg]) (OutT, error) {
	fmt.Println("Service create")
	userId, err := util.ExtractUserId(ctx)
	if err != nil {
		return *new(OutT), err
	}
	c, err := s.Db.LoadUserConfig(userId)
	if err != nil {
		return *new(OutT), err
	}
	resource, err := util.InsertIntoConfig(in.GetResource(), c)
	if err != nil {
		return *new(OutT), err
	}

	fmt.Printf("Deploying %s\n", resource.ResourceId)
	_, err = deploy.Deploy(ctx, c, nil, resource)
	if err != nil {
		return *new(OutT), err
	}

	err = s.Db.StoreUserConfig(c)
	if err != nil {
		return *new(OutT), err
	}
	return s.Read(ctx, &resourcespb.ReadVirtualNetworkRequest{ResourceId: resource.ResourceId})
}

func (s Service[Arg, OutT]) Read(ctx context.Context, in WithResourceId) (OutT, error) {
	return WrappingErrors(s.read)(ctx, in)
}

func (s Service[Arg, OutT]) read(ctx context.Context, in WithResourceId) (OutT, error) {
	fmt.Printf("Service read: %s\n", in.GetResourceId())
	userId, err := util.ExtractUserId(ctx)
	if err != nil {
		return *new(OutT), err
	}

	c, err := s.Db.LoadUserConfig(userId)
	if err != nil {
		return *new(OutT), err
	}
	for _, r := range c.Resources {
		if r.ResourceId == in.GetResourceId() {
			converted, err := r.ResourceArgs.ResourceArgs.UnmarshalNew()
			if err != nil {
				return *new(OutT), err
			}
			err = deploy.MaybeInit(c.UserId)
			if err != nil {
				return *new(OutT), err
			}
			state, err := deploy.GetState(c.UserId)
			if err != nil {
				return *new(OutT), err
			}
			return s.Converters.Convert(in.GetResourceId(), converted.(Arg), state)
		}
	}

	return *new(OutT), errors.ResourceNotFound(in.GetResourceId())
}

func (s Service[Arg, OutT]) Update(ctx context.Context, in UpdateRequest[Arg]) (OutT, error) {
	return WrappingErrors(s.update)(ctx, in)
}

func (s Service[Arg, OutT]) update(ctx context.Context, in UpdateRequest[Arg]) (OutT, error) {
	fmt.Printf("Service update: %s\n", in.GetResourceId())
	userId, err := util.ExtractUserId(ctx)
	if err != nil {
		return *new(OutT), err
	}
	c, err := s.Db.LoadUserConfig(userId)
	if err != nil {
		return *new(OutT), err
	}
	previousResource := find(c, in.GetResourceId())
	err = util.UpdateInConfig(c, in.GetResourceId(), in.GetResource())
	if err != nil {
		return *new(OutT), err
	}

	_, err = deploy.Deploy(ctx, c, previousResource, find(c, in.GetResourceId()))
	if err != nil {
		return *new(OutT), err
	}

	err = s.Db.StoreUserConfig(c)
	if err != nil {
		return *new(OutT), err
	}
	return s.Read(ctx, in)
}

func (s Service[Arg, OutT]) Delete(ctx context.Context, in WithResourceId) (*commonpb.Empty, error) {
	return WrappingErrors(s.delete)(ctx, in)
}

func (s Service[Arg, OutT]) delete(ctx context.Context, in WithResourceId) (*commonpb.Empty, error) {
	fmt.Printf("Service delete: %s\n", in.GetResourceId())
	userId, err := util.ExtractUserId(ctx)
	if err != nil {
		return nil, err
	}
	c, err := s.Db.LoadUserConfig(userId)
	if err != nil {
		return nil, err
	}
	previousResource := find(c, in.GetResourceId())
	err = util.DeleteResourceFromConfig(c, in.GetResourceId())
	if err != nil {
		return nil, err
	}

	_, err = deploy.Deploy(ctx, c, previousResource, nil)
	if err != nil {
		return nil, err
	}

	err = s.Db.StoreUserConfig(c)
	if err != nil {
		return nil, err
	}
	return &commonpb.Empty{}, nil
}

func find(c *configpb.Config, resourceId string) *configpb.Resource {
	for _, r := range c.Resources {
		if r.ResourceId == resourceId {
			return r
		}
	}
	return nil
}
