package repo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBSetting Repository
	inMem     = NewSettingInServiceMemory()
)

type Mgo struct {
	db *mongo.Collection
}

func NewMongoDatabase(uri, db, col string) *Mgo {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database(db).Collection(col)
	o := &Mgo{db: collection}
	defer o.loadSettingIntoMemory(ctx)
	return o
}
func (o *Mgo) loadSettingIntoMemory(ctx context.Context) {
	cur, e := o.db.Find(ctx, bson.M{})
	if e != nil && e != mongo.ErrNoDocuments {
		return
	}

	settings := []Setting{}
	e = cur.All(ctx, &settings)
	inMem.SetA(settings...)
}
func (o *Mgo) AddSetting(ctx context.Context, doc Setting) (Setting, error) {
	defer o.loadSettingIntoMemory(ctx)

	r, e := o.db.InsertOne(ctx, doc)
	if e != nil {
		return doc, e
	}
	doc.ID = r.InsertedID.(primitive.ObjectID)

	inMem.SetA(doc)
	return doc, nil
}
func (o *Mgo) GetSetting(ctx context.Context, id primitive.ObjectID) (Setting, error) {
	setting := Setting{}
	if s := inMem.Get(id.Hex()); s != nil {
		return *s, nil
	}

	r := o.db.FindOne(ctx, bson.M{"_id": id})
	if e := r.Decode(&setting); e != nil {
		return setting, e
	}

	inMem.SetA(setting)
	return setting, r.Err()
}
func (o *Mgo) GetSettingList(ctx context.Context) ([]Setting, error) {
	settings := []Setting{}
	if s := inMem.GetAll(); s != nil {
		return *s, nil
	}

	cur, e := o.db.Find(ctx, bson.M{})
	if e != nil && e != mongo.ErrNoDocuments {
		return settings, e
	}
	e = cur.All(ctx, &settings)

	inMem.SetA(settings...)
	return settings, e
}
func (o *Mgo) DeleteSetting(ctx context.Context, id primitive.ObjectID) error {
	defer o.loadSettingIntoMemory(ctx)

	r, e := o.db.DeleteOne(ctx, bson.M{"_id": id})
	if e != nil {
		return e
	}
	if r.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	inMem.Unset(id.Hex())
	return nil
}
func (o *Mgo) UpdateSetting(ctx context.Context, id primitive.ObjectID, data USetting) (Setting, error) {
	defer o.loadSettingIntoMemory(ctx)

	update := bson.M{"$set": bson.M{"config": data}}
	r, e := o.db.UpdateOne(ctx, bson.M{"_id": id}, update)
	if e != nil {
		return Setting{}, e
	}
	if r.MatchedCount == 0 {
		return Setting{}, mongo.ErrNoDocuments
	}

	inMem.Unset(id.Hex())
	return o.GetSetting(ctx, id)
}
func (o *Mgo) GetSettingByMethodAndURI(ctx context.Context, method string, uri string) (Setting, error) {
	setting := Setting{}

	if s := inMem.GetAll(); s != nil {
		for _, i := range *s {
			if i.Method == method && i.URI == uri {
				return i, nil
			}
		}
	}

	r := o.db.FindOne(ctx, bson.M{"config.uri": uri, "config.method": method})
	if e := r.Decode(&setting); e != nil {
		return setting, e
	}
	return setting, r.Err()
}
