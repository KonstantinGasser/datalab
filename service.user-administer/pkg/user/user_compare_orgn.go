package user

import (
	"context"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_user/pkg/storage"
)

func (user user) CompareOrgn(ctx context.Context, storage storage.Storage, baseObject string, compareWith []string) (bool, []string, errors.ErrApi) {
	// create Comparator as base to compare values with
	// comparator := Comparator{
	// 	LoadData: func() (map[string]interface{}, error) {
	// 		var data map[string]interface{}
	// 		err := storage.FindOne(ctx, userDatabase, userCollection, bson.D{{Key: "_id", Value: baseObject}}, &data)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return data, nil
	// 	},
	// 	By:    "orgn_domain",
	// 	Value: map[string]interface{}{}, // since the Fetch is provided the Value will be assigned with the queried data
	// }
	// // create Comparable from the request data
	// comparable := Comparable{
	// 	LoadData: func() ([]map[string]interface{}, error) {
	// 		var data []map[string]interface{}
	// 		err := storage.FindMany(ctx, userDatabase, userCollection, bson.D{{Key: "_id", Value: bson.M{"$in": compareWith}}}, &data)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return data, nil
	// 	},
	// 	By:        "orgn_domain",
	// 	StorageID: "_id",
	// 	// used to cross-check if all users have been found else comparison will tell false
	// 	ExpectedCount: len(compareWith),
	// }
	// status, comparison, err := user.compareByField(ctx, storage, comparator, comparable)
	// if err != nil {
	// 	return status, false, nil, err
	// }
	// return status, comparison.TruthfulValid, comparison.MissItems, nil
	return false, nil, nil
}
