package mapper

import (
	domain "github.com/SmartCityFlensburg/green-space-management/internal/entities/tree"
	response "github.com/SmartCityFlensburg/green-space-management/internal/service/entities/tree"
	repo "github.com/SmartCityFlensburg/green-space-management/internal/storage/entities/tree"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// goverter:converter
// goverter:extend PrimitiveIDToString StringToPrimitiveID
type TreeMapper interface {
	ToEntity(src *domain.Tree) *repo.TreeEntity
  FromEntity(src *repo.TreeEntity) *domain.Tree
  FromEntityList(src []*repo.TreeEntity) []*domain.Tree

	ToResponse(src *domain.Tree) *response.TreeResponse
  ToResponseList(src []*domain.Tree) []*response.TreeResponse
  FromResponse(src *response.TreeResponse) *domain.Tree
}

func PrimitiveIDToString(id primitive.ObjectID) string {
  return id.Hex()
}

func StringToPrimitiveID(id string) primitive.ObjectID {
  objID, _ := primitive.ObjectIDFromHex(id)
  return objID
}



