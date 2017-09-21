package mongodb

import (
	//standard library
	"errors"
	"time"
	//t1cg library
	"github.com/t1cg/util/apperror"
	"github.com/t1cg/util/logger"
	"github.com/t1cg/util/runstat"
	//third party library
	mgo "gopkg.in/mgo.v2"
)

// Errors constants.
var Errors = struct {
	DatabaseAggregatePipeFailed       error
	DatabaseFindFailed                error
	DatabaseFindAndApplyFailed        error
	DatabaseIncorrectConnectionString error
	DatabaseInsertFailed              error
	DatabaseCollectionNotFound        error
	DatabaseFoundZeroRecords          error
	DatabaseInvalidQuery              error
	DatabaseSessionNil                error
	DatabaseUpdateFailed              error
}{
	DatabaseAggregatePipeFailed:       errors.New("Database aggregate pipe operation failed"),
	DatabaseFindFailed:                errors.New("Database find operation failed"),
	DatabaseFindAndApplyFailed:        errors.New("Database find & apply operation failed"),
	DatabaseInsertFailed:              errors.New("Database insert operation failed"),
	DatabaseIncorrectConnectionString: errors.New("Database connection string incorrect"),
	DatabaseCollectionNotFound:        errors.New("Database collection not found"),
	DatabaseFoundZeroRecords:          errors.New("Database find returned 0 records"),
	DatabaseInvalidQuery:              errors.New("Database invalid query"),
	DatabaseSessionNil:                errors.New("Database session returned nil"),
	DatabaseUpdateFailed:              errors.New("Database update operation failed"),
}

// ConnectionInfo struct type defines the struct used to hold the database connection string.
type ConnectionInfo struct {
	Host       string
	DBName     string
	User       string
	Pw         string
	To         int
	Collection string
}

// GetSession function returns the MongdDB session.
func GetSession(cn ConnectionInfo) (*mgo.Session, *apperror.AppInfo) {

	//get the caller and callee (if any) function names
	fname := logger.GetFuncName()

	//performance analysis - begin
	trace := &runstat.RunInfo{Name: fname, StartTime: time.Now()}
	defer trace.MeasureRuntime()

	//check parameter
	if len(cn.Host) == 0 || len(cn.DBName) == 0 || len(cn.User) == 0 || len(cn.Pw) == 0 {
		a := &apperror.AppInfo{Msg: Errors.DatabaseIncorrectConnectionString}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		return nil, a
	}

	//set timeout value
	var timeout int

	//TODO: revisit the hard coded value
	if cn.To == 0 {
		timeout = 3
	} else {
		timeout = cn.To
	}

	//setup mongodb
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{cn.Host},
		Database: cn.DBName,
		Username: cn.User,
		Password: cn.Pw,
		Timeout:  time.Duration(timeout) * time.Second,
	}

	msession, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		a := &apperror.AppInfo{Msg: err}
		a.LogError(a.Error())
		trace.SetEndTime(time.Now())
		return nil, a
	}

	defer msession.Close()

	//performance analysis - end
	trace.SetEndTime(time.Now())

	//Clone works just like Copy, but also reuses the same socket as the original session,
	//in case it had already reserved one due to its consistency guarantees.
	return msession.Clone(), nil
}
