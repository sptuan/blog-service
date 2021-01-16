package global

import (
	"github.com/sptuan/blog-service/pkg/logger"
	"github.com/sptuan/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)

var (
	Logger *logger.Logger
)
