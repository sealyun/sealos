package database

import (
	"context"
	"fmt"
	accountv1 "github.com/labring/sealos/controllers/account/api/v1"
	"github.com/labring/sealos/controllers/pkg/common"
	"github.com/labring/sealos/pkg/utils/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"math"
	"strings"
	"time"
)

const (
	DefaultDBName       = "sealos-resources"
	DefaultMeteringConn = "metering"
	DefaultMonitorConn  = "monitor"
	DefaultBillingConn  = "billing"
	DefaultPricesConn   = "prices"
)

const (
	MongoURL      = "MONGO_URI"
	MongoUsername = "MONGO_USERNAME"
	MongoPassword = "MONGO_PASSWORD"
)

type MongoDB struct {
	Url          string
	Client       *mongo.Client
	DBName       string
	MonitorConn  string
	MeteringConn string
	BillingConn  string
	PricesConn   string
}

type AccountBalanceSpecBSON struct {
	OrderID string          `json:"order_id" bson:"order_id"`
	Owner   string          `json:"owner" bson:"owner"`
	Time    time.Time       `json:"time" bson:"time"`
	Type    accountv1.Type  `json:"type" bson:"type"`
	Costs   accountv1.Costs `json:"costs,omitempty" bson:"costs,omitempty"`
	Amount  int64           `json:"amount,omitempty" bson:"amount"`
}

func (m *MongoDB) Disconnect(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}

func (m *MongoDB) GetPrices() *accountv1.PriceQuery {
	//collection := mongoClient.Database(SealosResourcesDBName).Collection(SealosPricesCollectionName)
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//cursor, err := collection.Find(ctx, bson.M{})
	//if err != nil {
	//	return nil, fmt.Errorf("get all prices error: %v", err)
	//}
	//var prices []Price
	//if err = cursor.All(ctx, &prices); err != nil {
	//	return nil, fmt.Errorf("get all prices error: %v", err)
	//}
	//var pricesMap = make(map[string]Price, len(prices))
	//for i := range prices {
	//	pricesMap[strings.ToLower(prices[i].Property)] = prices[i]
	//}
	//return pricesMap, nil
	m.getMonitorCollection()
	return nil
}

func (m *MongoDB) SaveBillingsWithAccountBalance(accountBalanceSpec *accountv1.AccountBalanceSpec) error {

	// Time    metav1.Time `json:"time" bson:"time"`
	// time字段如果为time.Time类型无法转换为json crd，所以使用metav1.Time，但是使用metav1.Time无法插入到mongo中，所以需要转换为time.Time

	accountBalanceTime := accountBalanceSpec.Time.Time

	// Create BSON document
	accountBalanceDoc := bson.M{
		"order_id": accountBalanceSpec.OrderID,
		"owner":    accountBalanceSpec.Owner,
		"time":     accountBalanceTime.UTC(),
		"type":     accountBalanceSpec.Type,
		"costs":    accountBalanceSpec.Costs,
		"amount":   accountBalanceSpec.Amount,
	}
	_, err := m.getBillingCollection().InsertOne(context.Background(), accountBalanceDoc)
	return err
}

func (m *MongoDB) GetMeteringOwnerTimeResult(queryTime time.Time, queryCategories, queryProperties []string, queryOwner string) (*MeteringOwnerTimeResult, error) {
	pipeline := bson.A{
		bson.D{{Key: "$match", Value: bson.M{
			"time":     queryTime,
			"category": bson.M{"$in": queryCategories},
			"property": bson.M{"$in": queryProperties},
		}}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id":           bson.M{"property": "$property"},
			"propertyTotal": bson.M{"$sum": "$amount"},
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":           0,
			"property":      "$_id.property",
			"propertyTotal": 1,
		}}},
		bson.D{{Key: "$group", Value: bson.M{
			"_id":         nil,
			"amountTotal": bson.M{"$sum": "$propertyTotal"},
			"costs":       bson.M{"$push": bson.M{"k": "$property", "v": "$propertyTotal"}},
		}}},
		bson.D{{Key: "$addFields", Value: bson.M{
			"owner":  queryOwner,
			"time":   queryTime,
			"amount": "$amountTotal",
			"costs":  bson.M{"$arrayToObject": "$costs"},
		}}},
	}

	/*
		db.metering.aggregate([
		{ $match:
		  { time: queryTime, category:
		     { $in: ["ns-gxqoxr8s"] }, property: { $in: ["cpu", "memory", "storage"] } } },
		{ $group: { _id: { property: "$property" }, propertyTotal: { $sum: "$amount" } } },
		{ $project: { _id: 0, property: "$_id.property", propertyTotal: 1 } },
		{ $group: { _id: null, amountTotal: { $sum: "$propertyTotal" }, costs: { $push: { k: "$property", v: "$propertyTotal" } } } },
		{ $addFields: { orderId: "111111111", own: queryOwn, time: queryTime, type: 0, amount: "$amountTotal", costs: { $arrayToObject: "$costs" } } },
		{ $out: "results1" }]);
	*/
	cursor, err := m.getMeteringCollection().Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if cursor.Next(context.Background()) {
		var result MeteringOwnerTimeResult
		if err = cursor.Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, nil
}

func (m *MongoDB) InsertMonitor(ctx context.Context, monitors ...*common.Monitor) error {
	if len(monitors) == 0 {
		return nil
	}
	var manyMonitor []interface{}
	for i := range monitors {
		manyMonitor = append(manyMonitor, monitors[i])
	}
	_, err := m.getMonitorCollection().InsertMany(ctx, manyMonitor)
	return err
}

func (m *MongoDB) GetAllPricesMap() (map[string]common.Price, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := m.getPricesCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("get all prices error: %v", err)
	}
	var prices []common.Price
	if err = cursor.All(ctx, &prices); err != nil {
		return nil, fmt.Errorf("get all prices error: %v", err)
	}
	var pricesMap = make(map[string]common.Price, len(prices))
	for i := range prices {
		pricesMap[strings.ToLower(prices[i].Property)] = prices[i]
	}
	return pricesMap, nil
}

func (m *MongoDB) GenerateMeteringData(startTime, endTime time.Time, prices map[string]common.Price) error {
	filter := bson.M{
		"time": bson.M{
			"$gte": startTime,
			"$lt":  endTime,
		},
	}
	cursor, err := m.getMonitorCollection().Find(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("find monitors error: %v", err)
	}
	defer cursor.Close(context.Background())

	meteringMap := make(map[string]map[string]int64)
	countMap := make(map[string]map[string]int64)
	updateTimeMap := make(map[string]map[string]*time.Time)

	for cursor.Next(context.Background()) {
		var monitor common.Monitor
		if err := cursor.Decode(&monitor); err != nil {
			return fmt.Errorf("decode monitor error: %v", err)
		}

		if _, ok := updateTimeMap[monitor.Category]; !ok {
			updateTimeMap[monitor.Category] = make(map[string]*time.Time)
		}
		if _, ok := updateTimeMap[monitor.Category][monitor.Property]; !ok {
			lastUpdateTime, err := m.GetUpdateTimeForCategoryAndPropertyFromMetering(monitor.Category, monitor.Property)
			if err != nil {
				logger.Debug(err, "get latest update time failed", "category", monitor.Category, "property", monitor.Property)
			}
			updateTimeMap[monitor.Category][monitor.Property] = &lastUpdateTime
		}
		lastUpdateTime := updateTimeMap[monitor.Category][monitor.Property].UTC()

		if /* skip last update lte 1 hour*/ lastUpdateTime.Before(startTime) || lastUpdateTime.Equal(startTime) {
			if _, ok := meteringMap[monitor.Category]; !ok {
				meteringMap[monitor.Category] = make(map[string]int64)
				countMap[monitor.Category] = make(map[string]int64)
			}

			meteringMap[monitor.Category][monitor.Property] += monitor.Value
			countMap[monitor.Category][monitor.Property]++
			continue
		}
		logger.Debug("Info", "skip metering", "category", monitor.Category, "property", monitor.Property, "lastUpdateTime", updateTimeMap[monitor.Category][monitor.Property].UTC(), "startTime", startTime)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("cursor error: %v", err)
	}
	eg, _ := errgroup.WithContext(context.Background())

	for category, propertyMap := range meteringMap {
		for property, totalValue := range propertyMap {
			count := countMap[category][property]
			if count < 60 {
				count = 60
			}
			unitValue := math.Ceil(float64(totalValue) / float64(count))
			metering := &common.Metering{
				Category: category,
				Property: property,
				Time:     endTime,
				Amount:   int64(unitValue * float64(prices[property].Price)),
				Value:    int64(unitValue),
				Status:   0,
				//Detail:   "",
			}
			_category, _property := category, property
			eg.Go(func() error {
				_, err := m.getMeteringCollection().InsertOne(context.Background(), metering)
				if err != nil {
					//TODO if insert failed, should todo?
					logger.Error(err, "insert metering data failed", "category", _category, "property", _property)
				}
				return err
			})
		}
	}
	return eg.Wait()
}

func (m *MongoDB) GetUpdateTimeForCategoryAndPropertyFromMetering(category string, property string) (time.Time, error) {
	filter := bson.M{"category": category, "property": property}
	// sort by time desc
	opts := options.FindOne().SetSort(bson.M{"time": -1})

	var result struct {
		Time time.Time `bson:"time"`
	}
	err := m.getMeteringCollection().FindOne(context.Background(), filter, opts).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No documents match the filter. Handle this case accordingly.
			return time.Time{}, nil
		}
		return time.Time{}, err
	}
	return result.Time, nil
}

func (m *MongoDB) queryBillingRecordsByOrderId(billingRecordQuery *accountv1.BillingRecordQuery, owner string) error {
	if billingRecordQuery.Spec.OrderID == "" {
		return fmt.Errorf("order id is empty")
	}
	billingColl := m.getBillingCollection()
	matchStage := bson.D{
		{"$match", bson.D{
			{"order_id", billingRecordQuery.Spec.OrderID},
			{"owner", owner},
		}},
	}
	var billingRecords []accountv1.AccountBalanceSpec
	ctx := context.Background()

	cursor, err := billingColl.Aggregate(ctx, bson.A{matchStage})
	if err != nil {
		return fmt.Errorf("failed to execute aggregate query: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var bsonRecord AccountBalanceSpecBSON
		if err := cursor.Decode(&bsonRecord); err != nil {
			return fmt.Errorf("failed to decode billing record: %w", err)
		}
		billingRecord := accountv1.AccountBalanceSpec{
			OrderID: bsonRecord.OrderID,
			Owner:   bsonRecord.Owner,
			Time:    metav1.NewTime(bsonRecord.Time),
			Type:    bsonRecord.Type,
			Costs:   bsonRecord.Costs,
			Amount:  bsonRecord.Amount,
		}
		billingRecords = append(billingRecords, billingRecord)
	}

	billingRecordQuery.Status.Items = billingRecords
	return nil
}

func (m *MongoDB) QueryBillingRecords(billingRecordQuery *accountv1.BillingRecordQuery, owner string) (err error) {
	if billingRecordQuery.Spec.OrderID != "" {
		return m.queryBillingRecordsByOrderId(billingRecordQuery, owner)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	billingColl := m.getBillingCollection()
	timeMatchValue := bson.D{{"$gte", billingRecordQuery.Spec.StartTime.Time}, {"$lte", billingRecordQuery.Spec.EndTime.Time}}
	matchStage := bson.D{
		{"$match", bson.D{
			{"time", timeMatchValue},
			{"owner", owner},
		}},
	}

	if billingRecordQuery.Spec.Type != -1 {
		matchStage = bson.D{
			{"$match", bson.D{
				{"time", timeMatchValue},
				{"owner", owner},
				{"type", billingRecordQuery.Spec.Type},
			}},
		}
	}

	// Pipeline for getting the paginated data
	pipeline := bson.A{
		matchStage,
		bson.D{{"$sort", bson.D{{"time", -1}}}},
		bson.D{{"$skip", (billingRecordQuery.Spec.Page - 1) * billingRecordQuery.Spec.PageSize}},
		bson.D{{"$limit", billingRecordQuery.Spec.PageSize}},
	}

	pipelineAll := bson.A{
		matchStage,
	}

	pipelineCountAndAmount := bson.A{
		bson.D{{"$match", bson.D{
			{"time", timeMatchValue},
			{"owner", owner},
			{"type", accountv1.Consumption},
		}}},
		bson.D{{"$addFields", bson.D{
			{"costsArray", bson.D{{"$objectToArray", "$costs"}}},
		}}},
		bson.D{{"$unwind", "$costsArray"}},
		bson.D{{"$group", bson.D{
			{"_id", bson.D{
				{"type", "$type"},
				{"key", "$costsArray.k"},
			}},
			{"total", bson.D{{"$sum", "$costsArray.v"}}},
			{"count", bson.D{{"$sum", 1}}},
		}}},
	}

	pipelineRechargeAmount := bson.A{
		bson.D{{"$match", bson.D{
			{"time", timeMatchValue},
			{"owner", owner},
			{"type", accountv1.Recharge},
		}}},
		bson.D{{"$group", bson.D{
			{"_id", nil},
			{"totalRechargeAmount", bson.D{{"$sum", "$amount"}}},
			{"count", bson.D{{"$sum", 1}}},
		}}},
	}

	cursor, err := billingColl.Aggregate(ctx, pipeline)
	if err != nil {
		return fmt.Errorf("failed to execute aggregate query: %w", err)
	}
	defer cursor.Close(ctx)

	var billingRecords []accountv1.AccountBalanceSpec
	for cursor.Next(ctx) {
		var bsonRecord AccountBalanceSpecBSON
		if err := cursor.Decode(&bsonRecord); err != nil {
			return fmt.Errorf("failed to decode billing record: %w", err)
		}
		billingRecord := accountv1.AccountBalanceSpec{
			OrderID: bsonRecord.OrderID,
			Owner:   bsonRecord.Owner,
			Time:    metav1.NewTime(bsonRecord.Time),
			Type:    bsonRecord.Type,
			Costs:   bsonRecord.Costs,
			Amount:  bsonRecord.Amount,
		}
		billingRecords = append(billingRecords, billingRecord)
	}

	totalCount := 0

	// 总数量
	cursorAll, err := billingColl.Aggregate(ctx, pipelineAll)
	if err != nil {
		return fmt.Errorf("failed to execute aggregate all query: %w", err)
	}
	totalCount = cursorAll.RemainingBatchLength()
	cursorAll.Close(ctx)

	// 消费总金额Costs Executing the second pipeline for getting the total count, recharge and deduction amount
	cursorCountAndAmount, err := billingColl.Aggregate(ctx, pipelineCountAndAmount)
	if err != nil {
		return fmt.Errorf("failed to execute aggregate query for count and amount: %w", err)
	}
	defer cursorCountAndAmount.Close(ctx)

	totalDeductionAmount := make(map[string]int64)
	totalRechargeAmount := int64(0)

	for cursorCountAndAmount.Next(ctx) {
		var result struct {
			ID struct {
				Type int    `bson:"type"`
				Key  string `bson:"key"`
			} `bson:"_id"`
			Total int64 `bson:"total"`
		}
		if err := cursorCountAndAmount.Decode(&result); err != nil {
			return fmt.Errorf("failed to decode billing record: %w", err)
		}
		if result.ID.Type == 0 {
			totalDeductionAmount[result.ID.Key] = result.Total
		}
	}

	// 充值总金额
	cursorRechargeAmount, err := billingColl.Aggregate(ctx, pipelineRechargeAmount)
	if err != nil {
		return fmt.Errorf("failed to execute aggregate query for recharge amount: %w", err)
	}
	defer cursorRechargeAmount.Close(ctx)

	for cursorRechargeAmount.Next(ctx) {
		var result struct {
			TotalRechargeAmount int64 `bson:"totalRechargeAmount"`
			Count               int   `bson:"count"`
		}
		if err := cursorRechargeAmount.Decode(&result); err != nil {
			return fmt.Errorf("failed to decode recharge amount record: %w", err)
		}
		totalRechargeAmount = result.TotalRechargeAmount
	}

	totalPages := (totalCount + billingRecordQuery.Spec.PageSize - 1) / billingRecordQuery.Spec.PageSize
	billingRecordQuery.Status.Items, billingRecordQuery.Status.PageLength,
		billingRecordQuery.Status.RechargeAmount, billingRecordQuery.Status.DeductionAmount = billingRecords, totalPages, totalRechargeAmount, totalDeductionAmount
	return nil
}

func (m *MongoDB) getMeteringCollection() *mongo.Collection {
	return m.Client.Database(m.DBName).Collection(m.MeteringConn)
}

func (m *MongoDB) getMonitorCollection() *mongo.Collection {
	return m.Client.Database(m.DBName).Collection(m.MonitorConn)
}

func (m *MongoDB) getPricesCollection() *mongo.Collection {
	return m.Client.Database(m.DBName).Collection(m.PricesConn)
}

func (m *MongoDB) getBillingCollection() *mongo.Collection {
	return m.Client.Database(m.DBName).Collection(m.BillingConn)
}

func (m *MongoDB) CreateBillingTimeSeriesIfNotExist() error {
	return m.CreateTimeSeriesIfNotExist(m.DBName, m.BillingConn)
}

// CreateMonitorTimeSeriesIfNotExist creates the time series table for monitor
func (m *MongoDB) CreateMonitorTimeSeriesIfNotExist() error {
	return m.CreateTimeSeriesIfNotExist(m.DBName, m.MonitorConn)
}

// CreateMeteringTimeSeriesIfNotExist creates the time series table for metering
func (m *MongoDB) CreateMeteringTimeSeriesIfNotExist() error {
	return m.CreateTimeSeriesIfNotExist(m.DBName, m.MeteringConn)
}

func (m *MongoDB) CreateTimeSeriesIfNotExist(dbName, collectionName string) error {
	// Check if the collection already exists
	collections, err := m.Client.Database(dbName).ListCollectionNames(context.Background(), bson.M{"name": collectionName})
	if err != nil {
		return err
	}

	// If the collection does not exist, create it
	if len(collections) == 0 {
		cmd := bson.D{
			{Key: "create", Value: collectionName},
			{Key: "timeseries", Value: bson.D{{Key: "timeField", Value: "time"}}},
		}
		err = m.Client.Database(dbName).RunCommand(context.TODO(), cmd).Err()
	}
	return err
}

func NewMongoDB(ctx context.Context, URL string) (Interface, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))
	return &MongoDB{
		Client:       client,
		DBName:       DefaultDBName,
		MeteringConn: DefaultMeteringConn,
		MonitorConn:  DefaultMonitorConn,
		BillingConn:  DefaultBillingConn,
		PricesConn:   DefaultPricesConn,
	}, err
}
