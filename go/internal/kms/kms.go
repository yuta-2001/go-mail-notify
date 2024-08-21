package kms

import (
    "encoding/base64"
    "fmt"
    "strings"
    "os"

    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/kms"

    "github.com/aws/aws-sdk-go/aws"
)

type KmsInstanceStruct struct {
    Client *kms.KMS
}

var kmsInstance *KmsInstanceStruct

func GetKmsInstance(env string) (*KmsInstanceStruct, error) {
    if kmsInstance == nil {
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
        svc := kms.New(sess)
        kmsInstance = &KmsInstanceStruct{
            Client: svc,
        }
    }

    return kmsInstance, nil
}

func (kmsInstance *KmsInstanceStruct)DecryptVariable(variable string) (string, error) {
    v, err := base64.StdEncoding.DecodeString(variable)
    if err != nil {
        return "", fmt.Errorf("failed to decode base64 string: %w", err)
    }

    input := &kms.DecryptInput{
        CiphertextBlob: v,
    }

    result, err := kmsInstance.Client.Decrypt(input)
    if err != nil {
        return "", fmt.Errorf("failed to decrypt variable: %w", err)
    }

    return strings.Replace(string(result.Plaintext), "\n", "", -1), nil
}
