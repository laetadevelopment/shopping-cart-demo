package v1

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/laetadevelopment/shopping-cart-demo/microservices/items/pkg/api/v1"
)

const (
	apiVersion = "v1"
)

type itemServiceServer struct {
	db *mongo.Client
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// NewItemServiceServer connects to MongoDB
func NewItemServiceServer(db *mongo.Client) v1.ItemServiceServer {
	return &itemServiceServer{db: db}
}

func (s *itemServiceServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

// Create a new item in MongoDB
func (s *itemServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c := s.db.Database("items").Collection("item")

	p := v1.Item{
		Id:          uuid.NewV1().String(),
		Url:         req.Item.Url,
		Title:       req.Item.Title,
		Created:     ptypes.TimestampNow(),
		Updated:     ptypes.TimestampNow(),
	}

	_, err := c.InsertOne(ctx, &p)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to insert into Items.item-> "+err.Error())
	}

	return &v1.CreateResponse{
		Api: apiVersion,
	}, nil
}

// Read a item in MongoDB
func (s *itemServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c := s.db.Database("items").Collection("item")

	filter := bson.D{{"id", req.Id}}

	var p v1.Item

	err := c.FindOne(context.TODO(), filter).Decode(&p)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to find document-> "+err.Error())
	}

	return &v1.ReadResponse{
		Api:     apiVersion,
		Item: &p,
	}, nil

}

// Update a item in MongoDB
func (s *itemServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c := s.db.Database("items").Collection("item")

	filter := bson.D{{"id", req.Item.Id}}
	update := bson.D{
		{"$inc", bson.D{
			{"url", req.Item.Url},
			{"title", req.Item.Title},
			{"updated", ptypes.TimestampNow()},
		}},
	}

	updateResult, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to update document-> "+err.Error())
	}

	u := updateResult.ModifiedCount

	return &v1.UpdateResponse{
		Api:     apiVersion,
		Updated: u,
	}, nil
}

// Delete a item in MongoDB
func (s *itemServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c := s.db.Database("items").Collection("item")

	filter := bson.D{{"id", req.Id}}

	deleteResult, err := c.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to delete document-> "+err.Error())
	}

	d := deleteResult.DeletedCount

	return &v1.DeleteResponse{
		Api:     apiVersion,
		Deleted: d,
	}, nil
}

// List all items available via MongoDB Client
func (s *itemServiceServer) List(ctx context.Context, req *v1.ListRequest) (*v1.ListResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	c := s.db.Database("items").Collection("item")

	findOptions := options.Find()
	var list []*v1.Item

	cur, err := c.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return nil, status.Error(codes.Unknown, "failed to find documents in Items.item-> "+err.Error())
	}

	for cur.Next(context.TODO()) {
		var elem v1.Item
		err := cur.Decode(&elem)
		if err != nil {
			return nil, status.Error(codes.Unknown, "failed to decode document-> "+err.Error())
		}

		list = append(list, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, status.Error(codes.Unknown, "failed reading documents-> "+err.Error())
	}

	cur.Close(context.TODO())

	return &v1.ListResponse{
		Api:  apiVersion,
		Data: list,
	}, nil
}
