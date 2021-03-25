package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// Comparator can be used as base model with Comparable
type Comparator struct {
	// Filter can be used if data for comparison must first be loaded
	// if nil presents of data in Value is expected.
	Filter interface{}
	// By tells user.Comparator which field to compare with
	By string
	// Value holds the the key:value data against which the comparison
	// needs to be performed
	Value map[string]interface{}
}

// Comparable holds N items of type ComparableItem which should be checked against the Comparator
type Comparable struct {
	// Filter allows to query data for all ComparableItems
	// can be used if you only have the user uuid and need different fields to compare against
	// if Filter is nil then ALL items must have their value already initialized else they will be skipped
	Filter interface{}
	// filterResults should be an empty slice of bson.M in which the Filter will write the results
	// can be nil if Filter is nil
	filterResults []bson.M
	// By is the key to the value which you want to compare (must be present in the return of the filter)
	By string
	// StorageID is used to map the misses with the correlating storage id
	// for the caller to understand which item failed the comparison
	// example in MongoDB you will get back the object id => StorageID: "_id" in oder to map the miss to an
	// document
	StorageID string
	// Items holds all items which will be compared against the Comparator
	Items []ComparableItem
}

// ComparableItem holds the Value which gets compared with Comparator.Value[By]
type ComparableItem struct {
	// Identifier is used to trace back the item
	// if it was a miss
	Identifier string
	Value      interface{}
}

type CompareResult struct {
	// Hit number of successful comparisons
	Hit uint
	// Miss number of unsuccessful comparisons
	Miss uint
	// MissItems items which failed the comparison
	MissItems []string
	// TruthfulValid is true if all items match
	TruthfulValid bool
	Errors        []error
}

// Compare takes a Comparator and a Comparable type to perform checks wether the items in Comparable match the criteria in
// the base model - the Comparator. If Comparator or Comparable provide a Filter the filter to get the data will first be queried from the storage.
// In return the call receives a *CompareResult in which all invalid/miss-matched items and a total miss-mach count can be found
func (user user) Compare(ctx context.Context, storage storage.Storage, comparator Comparator, comparable Comparable) (int, *CompareResult, error) {

	// CompareResult struct
	var resultSet CompareResult = CompareResult{
		Hit:           0,
		Miss:          0,
		TruthfulValid: true,
		MissItems:     []string{},
	}
	// query for value to compare with
	if comparator.Filter != nil {
		var compValue bson.M
		if err := storage.FindOne(ctx, userDatabase, userCollection, comparator.Filter, &compValue); err != nil {
			return http.StatusInternalServerError, nil, err
		}
		// assign queried value to comparator
		comparator.Value = compValue
	}

	// check if comparator.By exists in Value
	if _, ok := comparator.Value[comparator.By]; !ok {
		return http.StatusBadRequest, nil, fmt.Errorf("could not find key: %v in Value map", comparator.By)
	}

	// check if ComparableItems need to be queried
	// if nil -> Comparable.Items.Value must not be null else they are skipped
	if comparable.Filter != nil {
		// query for Comparable.Item data
		if err := storage.FindMany(ctx, userDatabase, userCollection, comparable.Filter, &comparable.filterResults); err != nil {
			return http.StatusInternalServerError, nil, err
		}

		comparable.Items = make([]ComparableItem, len(comparable.filterResults))
		for i, item := range comparable.filterResults {
			// if storage.result.item does not have the Comparable.By key mark as error and skip
			if _, ok := item[comparable.By]; !ok {
				resultSet.Errors = append(resultSet.Errors, fmt.Errorf("could not find key: %v in comparable.FilterResults: %v", comparable.By, item))
				continue
			}
			// assign Value and Identifier to the item
			comparable.Items[i].Value = item[comparable.By]
			if comparable.StorageID != "" {
				comparable.Items[i].Identifier = item[comparable.StorageID].(string)
			}
		}
	}
	// actual comparison of Comparator value and Comparable.Items
	for _, item := range comparable.Items {
		if item.Value == nil || item.Value != comparator.Value[comparator.By] {
			// mark comparison as dirty
			if resultSet.TruthfulValid {
				resultSet.TruthfulValid = false
			}
			resultSet.Miss++
			resultSet.MissItems = append(resultSet.MissItems, item.Identifier)
			continue
		}
		resultSet.Hit++
	}
	return http.StatusOK, &resultSet, nil
}
