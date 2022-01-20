package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"google.golang.org/api/googleapi"
)

type GCSClient struct {
	ctx       context.Context
	projectID string
	*storage.Client
}

func NewGCSClient(projectID string) *GCSClient {
	if projectID != "" {
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			return nil
		}
		return &GCSClient{
			ctx:       ctx,
			projectID: projectID,
			Client:    client,
		}
	}
	return nil
}

func (g *GCSClient) GetContext() context.Context {
	return g.ctx
}

func (g *GCSClient) WriteObjectToBucket(bucketName, objectName string, buf []byte) (string, error) {
	bufLen := len(buf)
	bucket := g.Bucket(bucketName)
	obj := bucket.Object(objectName)
	w := obj.If(storage.Conditions{DoesNotExist: true}).NewWriter(g.ctx)
	wLen, err := w.Write(buf)
	if err != nil {
		return "", err
	}
	log.Printf("write incomplete should be :%d,actually: %d", bufLen, wLen)
	if bufLen != wLen {
		return "", errors.New(fmt.Sprintf("write incomplete should be :%d,actually: %d", bufLen, wLen))
	}

	err = w.Close()
	if err != nil {
		switch e := err.(type) {
		case *googleapi.Error:
			if e.Code == http.StatusPreconditionFailed {
				break
			}
			errMsg := fmt.Sprintf("ecode:%d,msg:%s", e.Code, e.Message)
			return "", errors.New(errMsg)
		default:
			return "", err
		}
	}

	url := "https://storage.cloud.google.com/" + bucketName + "/" + objectName

	return url, nil
}
