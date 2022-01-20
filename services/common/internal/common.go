package internal

import (
	"context"
	"fmt"
	"math/rand"
	"server/constants"
	"server/infra/storage"
	"server/log"
	"server/proto"
	"server/services/common/conf"
	"server/utils"
	"strings"
	"time"
)

func (s *Server) UploadPic(ctx context.Context, data map[string]interface{}) (*proto.Response, error) {
	name := utils.InterfaceToString(data["pic_name"])
	pic := utils.InterfaceToString(data["pic"])
	category_name := utils.InterfaceToString(data["category_name"])
	uid := utils.InterfaceToInt(data["uid"])

	projectID := conf.Gcs.ProjectID
	bucketName := conf.Gcs.AvatarBucket

	gClient := storage.NewGCSClient(projectID)
	if gClient == nil {
		log.Errorf("create google client")
	}
	defer gClient.Close()
	ojbTypeIndex := strings.LastIndex(name, ".")
	ojbType := name[ojbTypeIndex:]
	directory := fmt.Sprintf("%s/%s/", category_name, time.Unix(time.Now().Unix(), 0).Format("20060102"))

	objName := fmt.Sprintf("%d-%d-%d", time.Now().UnixNano(), uid, rand.Intn(100))
	objName += ojbType

	content, err := utils.Base64Decode(pic)
	if err != nil {
		return utils.GrpcResponse(constants.ERR_UPLOAD_FAILED, ""), nil
	}
	uri, err := gClient.WriteObjectToBucket(bucketName, directory+objName, []byte(content))
	if err != nil {
		return utils.GrpcResponse(constants.ERR_UPLOAD_FAILED, ""), nil
	}
	result := make(map[string]interface{})
	result["url"] = uri

	return utils.GrpcResponse(constants.OK, utils.Marshal(result)), nil

}
