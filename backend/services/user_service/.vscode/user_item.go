package user

// user implements the User interface
type user struct{}

// InsertNew inserts a new UserItem into the mongoDB. Checks before that the user name is not already taken if so returns
// an http.StatusBadeRequest, error
// func (user user) InsertNew(ctx context.Context, storage storage.Storage, userItem UserItem) (int, error) {
// 	// check if username is taken
// 	taken, err := storage.Exists(ctx, userDatabase, userCollection, bson.M{"_id": userItem.UUID})
// 	if err != nil {
// 		return http.StatusInternalServerError, err
// 	}
// 	if taken { // username exists but must be unique
// 		return http.StatusBadGateway, errors.New("username already taken")
// 	}
// 	// inserts new user in storage
// 	if err := storage.InsertOne(ctx, userDatabase, userCollection, userItem); err != nil {
// 		return http.StatusInternalServerError, err
// 	}
// 	return http.StatusOK, nil
// }

// // Authenticate verifies that a user performing a login exists in the database and has provided the correct
// // login credentials (username, password)
// func (user user) Authenticate(ctx context.Context, storage storage.Storage, username, password string) (int, *UserItemAuth, error) {

// 	var authedUser UserItemAuth
// 	if err := storage.FindOne(ctx, userDatabase, userCollection, bson.M{"username": username}, &authedUser); err != nil {
// 		if err == mongo.ErrNoDocuments { // as specified by the mongoClient.FindOne func
// 			return http.StatusForbidden, nil, nil
// 		}
// 		return http.StatusInternalServerError, nil, err
// 	}
// 	if !hash.CheckPasswordHash(password, authedUser.Password) {
// 		return http.StatusForbidden, nil, nil
// 	}
// 	return http.StatusOK, &authedUser, nil
// }

// // Update updates the allowed fields of the user record
// func (user user) Update(ctx context.Context, storage storage.Storage, userItem UserItemUpdateable) (int, error) {

// 	updateQuery := bson.D{
// 		{
// 			"$set", bson.D{
// 				{"first_name", userItem.FirstName},
// 				{"last_name", userItem.LastName},
// 				{"orgn_position", userItem.OrgnPosition},
// 			},
// 		},
// 	}
// 	if err := storage.UpdateOne(ctx, userDatabase, userCollection, bson.M{"_id": userItem.UUID}, updateQuery); err != nil {
// 		return http.StatusInternalServerError, err
// 	}
// 	return http.StatusOK, nil
// }

// // GetByIDs looks up all the records that match the provided UUIDs
// func (user user) GetByIDs(ctx context.Context, storage storage.Storage, UUIDs []string) (int, []UserItem, error) {

// 	var userList []UserItem
// 	if err := storage.FindMany(ctx, userDatabase, userCollection, bson.M{"_id": bson.M{"$in": UUIDs}}, &userList); err != nil {
// 		return http.StatusInternalServerError, nil, err
// 	}
// 	return http.StatusOK, userList, nil
// }

// // GetByID collects all user details for the given user uuid
// func (user user) GetByID(ctx context.Context, storage storage.Storage, UUID string) (int, UserItem, error) {
// 	var userItem UserItem
// 	if err := storage.FindOne(ctx, userDatabase, userCollection, bson.M{"_id": UUID}, &userItem); err != nil {
// 		return http.StatusInternalServerError, UserItem{}, nil
// 	}
// 	return http.StatusOK, userItem, nil
// }

// // Compare compares a slice of Comparable with the given Comparator. Compare results will be returned in a
// // CompareResult.
// func (user user) Compare(ctx context.Context, storage storage.Storage, comparator Comparator, comparable Comparable) (*CompareResult, error) {

// 	// CompareResult struct
// 	var resultSet CompareResult = CompareResult{
// 		InvalidCount: 0,
// 		InvalidList:  []ComparableItem{},
// 	}
// 	// query for value to compare with
// 	if comparator.Filter != nil {
// 		var compValue bson.M
// 		if err := storage.FindOne(ctx, userDatabase, userCollection, comparator.Filter, &compValue); err != nil {
// 			panic(err)
// 			return nil, err
// 		}
// 		// assign queried value to comparator
// 		comparator.Value = compValue
// 	}

// 	// check if comparator.By exists in Value
// 	if _, ok := comparator.Value[comparator.By]; !ok {
// 		return nil, fmt.Errorf("could not find key: %v in Value map", comparator.By)
// 	}

// 	// check if ComparableItems need to be queried
// 	if comparable.Filter != nil {
// 		if err := storage.FindMany(ctx, userDatabase, userCollection, comparable.Filter, &comparable.FilterResults); err != nil {
// 			panic(err)
// 			return nil, err
// 		}
// 		// assign results to each comparableItem
// 		if comparable.Items == nil {
// 			comparable.Items = make([]ComparableItem, len(comparable.FilterResults))
// 		}
// 		for i, item := range comparable.FilterResults {
// 			if _, ok := item[comparable.By]; !ok {
// 				resultSet.Errors = append(resultSet.Errors, fmt.Errorf("could not find key: %v in comparable.FilterResults: %v", comparable.By, item))
// 				continue
// 			}
// 			comparable.Items[i].Value = item[comparable.By]
// 			fmt.Printf("Item: %v\n", comparable.Items[i])
// 		}
// 	}

// 	for _, item := range comparable.Items {
// 		if item.Value == nil || item.Value != comparator.Value[comparator.By] {
// 			resultSet.InvalidCount++
// 			resultSet.InvalidList = append(resultSet.InvalidList, item)
// 		}
// 	}
// 	return &resultSet, nil
// }
