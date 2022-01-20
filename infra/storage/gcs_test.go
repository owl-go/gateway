package storage

import (
	"testing"
)

func TestWriteObject(t *testing.T) {
	projectID := ""
	gClient := NewGCSClient(projectID)
	if gClient == nil {
		t.Error("create google client")
	}

	defer gClient.Close()

	bucketName := "dev-qt"

	objName := "test.txt"

	content := "Hello,Google Cloud Storage!\n"
	uri, err := gClient.WriteObjectToBucket(bucketName, objName, []byte(content))
	if err != nil {
		t.Error(err)
	}
	t.Log(uri)
}
