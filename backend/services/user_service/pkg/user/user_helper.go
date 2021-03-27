package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/storage"
)

// Comparator can be used as base model with Comparable
type Comparator struct {
	// Fetch is a injected function which can be set if data must be loaded first.
	// for example if the data must first be collected from a data store or else
	Fetch func() (map[string]interface{}, error)

	// By tells user.Comparator which field to compare with
	By string

	// Value holds the the key:value data against which the comparison
	// needs to be performed
	Value map[string]interface{}
}

// Comparable holds N items of type ComparableItem which should be checked against the Comparator
type Comparable struct {
	// Fetch can be used if the data to compare with the Comparator must first
	// be fetched from somewhere else (like a data store)
	// the returned map[string]interface{} must include the KEY which is set in Comparable.By
	// else the Comparison will fail
	Fetch func() ([]map[string]interface{}, error)

	// filterResults will hold the fetched data from the Comparable.Fetch function
	fetchResults []map[string]interface{}

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
// the base model - the Comparator. In case the data for both first needs to be loaded the injected Fetch
// func will be executed to retrieve the data.
// In return the caller receives a *CompareResult with all hits and misses.
func (user user) Compare(ctx context.Context, storage storage.Storage, comparator Comparator, comparable Comparable) (int, *CompareResult, error) {

	// CompareResult struct
	var resultSet CompareResult = CompareResult{
		Hit:           0,
		Miss:          0,
		TruthfulValid: true,
		MissItems:     []string{},
	}
	// execute func Fetch if data for comparison needs to be loaded first
	if comparator.Fetch != nil {
		var err error
		comparator.Value, err = comparator.Fetch()
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
	}
	// check if comparator.By exists in Value
	if _, ok := comparator.Value[comparator.By]; !ok {
		return http.StatusBadRequest, nil, fmt.Errorf("could not find key: %v in Value map", comparator.By)
	}

	// check if ComparableItems need to be queried
	// if nil -> Comparable.Items.Value must not be null else they are skipped
	if comparable.Fetch != nil {
		var err error
		comparable.fetchResults, err = comparable.Fetch()
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
		comparable.Items = make([]ComparableItem, len(comparable.fetchResults))
		for i, item := range comparable.fetchResults {
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
