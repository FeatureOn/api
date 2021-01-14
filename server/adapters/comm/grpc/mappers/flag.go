package mappers

import (
	pb "github.com/FeatureOn/api/flagpb"
	"github.com/FeatureOn/api/server/domain"
)

// MapEnvironmentFlag2EnvironmentFlags maps domain EnvironmentFlag to grpc EnvirionmentFlag
func MapEnvironmentFlag2EnvironmentFlags(envFlag domain.EnvironmentFlag) *pb.EnvironmentFlags {
	envFlagDTO := pb.EnvironmentFlags{
		EnvironmentID: envFlag.EnvironmentID,
	}
	for _, flag := range envFlag.Flags {
		envFlagDTO.Flags = append(envFlagDTO.Flags, mapFlag2FlagResponse(flag))
	}
	return &envFlagDTO
}

func mapFlag2FlagResponse(flag domain.Flag) *pb.Flag {
	return &pb.Flag{
		FeatureKey: flag.FeatureKey,
		Value:      flag.Value,
	}
}
