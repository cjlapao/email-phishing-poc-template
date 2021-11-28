package executioncontext

import (
	"os"
	"strings"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/log"
	"github.com/cjlapao/go-template/entities"
)

var globalContext *Context
var logger = log.Get()

// Context entity
type Context struct {
	MongoConnectionString string              `json:"mongoConnectionString"`
	Mongo                 entities.LogicState `json:"mongo"`
	Database              string              `json:"database"`
	ShowHelp              bool                `json:"help"`
	ApiPrefix             string              `json:"apiPrefix"`
	ApiPort               string              `json:"apiPort"`
}

func Get() *Context {
	if globalContext != nil {
		return globalContext
	}

	logger.Debug("Creating Execution Context")
	globalContext = &Context{
		ShowHelp: helper.GetFlagSwitch("help", false),
	}

	globalContext.Getenv()

	return globalContext
}

// Getenv gets the environment variables for the entities
func (e *Context) Getenv() {

	e.Database = os.Getenv("MOCKER_DATABASENAME")
	e.MongoConnectionString = os.Getenv("MOCKER_MONGO_CONNECTION_STRING")
	e.ApiPrefix = os.Getenv("MOCKER_API_PREFIX")
	e.ApiPort = os.Getenv("MOCKER_API_PORT")

	if e.MongoConnectionString == "" {
		e.Mongo = entities.Disabled
	} else {
		e.Mongo = entities.Enabled
	}

	if e.Database == "" {
		e.Database = "database"
	}

	if e.ApiPort == "" {
		e.ApiPort = "80"
	}

	if strings.HasSuffix(e.ApiPrefix, "/") {
		e.ApiPrefix = strings.TrimSuffix(e.ApiPrefix, "/")
	}
}
