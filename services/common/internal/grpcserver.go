package internal

import (
	"context"
	"encoding/json"
	"server/constants"
	"server/proto"
	"server/services/common/dao"
	"server/utils"
	"strconv"
)

type Server struct {
	dao *dao.Dao
}

func (s *Server) Init() {
	s.dao = dao.NewDAO(mysqlClient)
}

func (s *Server) CallService(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	data, ok := s.prefixCheck(ctx, req)
	if !ok {
		return utils.GrpcResponse(constants.ERR_JSON_UNMARSHAL, ""), nil
	}
	method := req.Method
	switch method {
	case "upload_pic":
		return s.UploadPic(ctx, data)
	}
	return utils.GrpcResponse(constants.ERR_FAILED, ""), nil
}

func (s *Server) prefixCheck(ctx context.Context, req *proto.Request) (map[string]interface{}, bool) {
	var data map[string]interface{}
	json.Unmarshal([]byte(req.Data), &data)
	token := utils.InterfaceToString(req.Header.Token)
	if token != "" {
		uidStr := utils.GetUidFromToken(token)
		uid, err := strconv.Atoi(uidStr)
		if err != nil {
			return nil, false
		}
		data["uid"] = uid
	} else if token == "" { //token 不能为空
		return nil, false
	}

	return data, true
}
