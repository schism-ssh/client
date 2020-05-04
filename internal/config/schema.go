package config

import (
	"fmt"
)

type SchismProfile struct {
	LambdaName   string `toml:"lambda_name"`
	LambdaRegion string `toml:"lambda_region"`
	S3Bucket     string `toml:"s3_bucket"`
	S3Region     string `toml:"s3_region"`
	KmsRegion    string `toml:"kms_region"`
}

func (prof SchismProfile) String() string {
	var profStr = ""
	profStr += fmt.Sprintf("\tLambda:\n\t\tName: %s\n\t\tRegion: %s\n", prof.LambdaName, prof.LambdaRegion)
	profStr += fmt.Sprintf("\tS3:\n\t\tBucket: %s\n\t\tRegion: %s\n", prof.S3Bucket, prof.S3Region)
	profStr += fmt.Sprintf("\tKms:\n\t\tRegion: %s\n", prof.KmsRegion)
	return profStr
}

type SchismConfig struct {
	Profiles map[string]SchismProfile `toml:"profiles"`
}

func (conf SchismConfig) String() string {
	var strConf = ""
	for name, prof := range conf.Profiles {
		profStr := fmt.Sprintf("Profile: %s\n%s", name, prof)
		strConf += fmt.Sprintf("%s\n----|====|----\n", profStr)
	}
	return strConf
}
