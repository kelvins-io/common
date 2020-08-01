package convert

import (
	"fmt"
	"testing"
)

type UsrUserGetTableResponse struct {
	UserId      int64     `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserName    string    `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	NickName    string    `protobuf:"bytes,3,opt,name=nick_name,json=nickName" json:"nick_name,omitempty"`
	CreateTime  int64     `protobuf:"varint,4,opt,name=create_time,json=createTime" json:"create_time,omitempty"`
	ModifyTime  int64     `protobuf:"varint,5,opt,name=modify_time,json=modifyTime" json:"modify_time,omitempty"`
	DeletedTime int64     `protobuf:"varint,6,opt,name=deleted_time,json=deletedTime" json:"deleted_time,omitempty"`
	IsDel       int32     `protobuf:"varint,7,opt,name=is_del,json=isDel" json:"is_del,omitempty"`
	LoginList   []UsrUser `protobuf:"bytes,8,rep,name=login_list,json=loginList" json:"login_list,omitempty"`
}

type UsrUser struct {
	ColumnNameNickName
	UserId int64 `gorm:"primary_key;AUTO_INCREMENT" json:"user_id" db:"user_id" db:"user_id"`
}

type ColumnNameNickName struct {
	UserName string `json:"user_name" db:"user_name"`
	NickName string `json:"nick_name" db:"nick_name"`
}

func TestParseObject(t *testing.T) {
	fmt.Println(ParseObject(UsrUser{}, 1))
	fmt.Println(ParseObject(UsrUserGetTableResponse{}, 2))
}
