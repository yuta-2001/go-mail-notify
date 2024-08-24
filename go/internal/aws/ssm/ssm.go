package ssm

import (
    "fmt"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ssm"
)

type SsmInstanceStruct struct {
    Client *ssm.SSM
}

var SsmInstance *SsmInstanceStruct

func GetSsmInstance(env string) (*SsmInstanceStruct, error) {
    if SsmInstance == nil {
        sess, err := session.NewSession()
        if env == "local" {
            localKmsEndpoint := os.Getenv("local_kms_endpoint")
            if localKmsEndpoint == "" {
                return nil, fmt.Errorf("local_kms_endpoint is empty")
            }

            sess, err = session.NewSession(&aws.Config{
                Endpoint: aws.String("http://host.docker.internal:4566"),
            })
        }
        if err != nil {
            return nil, err
        }
        svc := ssm.New(sess)
        SsmInstance = &SsmInstanceStruct{
            Client: svc,
        }
    }

    return SsmInstance, nil
}


func (SsmInstance *SsmInstanceStruct)GetParamValue(param string, decrypt bool) (string, error) {
    input := &ssm.GetParameterInput{
        Name:           aws.String(param),
        WithDecryption: aws.Bool(decrypt),
    }

    result, err := SsmInstance.Client.GetParameter(input)
    if err != nil {
        return "", err
    }

    return *result.Parameter.Value, nil
}
