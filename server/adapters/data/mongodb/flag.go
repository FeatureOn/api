package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/FeatureOn/api/adapters/data/mongodb/dao"
	"github.com/FeatureOn/api/adapters/data/mongodb/mappers"
	"github.com/FeatureOn/api/domain"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FlagRepository represent a structure that will communicate to MongoDB to accomplish Flag related transactions
type FlagRepository struct {
	dbClient *mongo.Client
	dbName   string
}

func newFlagRepository(client *mongo.Client, databaseName string) FlagRepository {
	return FlagRepository{
		dbClient: client,
		dbName:   databaseName,
	}
}

// GetFlags gets values of all active flags for a given environment
func (fr FlagRepository) GetFlags(environmentID string) (domain.EnvironmentFlag, error) {
	collection := fr.dbClient.Database(fr.dbName).Collection(viper.GetString("FlagsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var envFlagDAO dao.EnvironmentFlagDAO
	objID, err := primitive.ObjectIDFromHex(environmentID)
	if err != nil {
		log.Error().Err(err).Msgf("Error parsing ProductID: %s", environmentID)
		return domain.EnvironmentFlag{}, err
	}
	err = collection.FindOne(ctx, bson.M{"environmentID": objID}).Decode(&envFlagDAO)
	if err != nil {
		log.Error().Err(err).Msgf("Error getting Products")
		return domain.EnvironmentFlag{}, err
	}
	return mappers.MapEnvironmentFlagDAO2EnvironmentFlag(envFlagDAO), nil
}

// UpdateFlag sets new value to a spesific flag
func (fr FlagRepository) UpdateFlag(environmentID string, featureKey string, value bool) error {
	collection := fr.dbClient.Database(fr.dbName).Collection(viper.GetString("FlagsCollection"))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	envID, err := primitive.ObjectIDFromHex(environmentID)
	if err != nil {
		log.Error().Err(err).Msgf("Cannot parse environmentID: %s into ObjectID", environmentID)
		return errors.New("EnvironmentID format is not as expected")
	}

	idDoc := bson.M{"environmentID": envID, "flags.featureKey": featureKey}
	upDoc := bson.D{{Key: "$set", Value: bson.M{"flags.$.value": value}}}
	var updateOpts options.UpdateOptions
	updateOpts.SetUpsert(false)
	_, err = collection.UpdateOne(ctx, idDoc, upDoc, &updateOpts)
	if err != nil {
		log.Error().Err(err).Msgf("Error updating the flag with environmentID: %s and featureKey: %s", environmentID, featureKey)
		return errors.New("Error updating the flag")
	}
	return nil

}
