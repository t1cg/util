package awssns

import (
	"strings"
	"time"

	//t1cg library
	"github.com/t1cg/util/apperror"
	"github.com/t1cg/util/logger"
	"github.com/t1cg/util/runstat"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// SendSnsSms function ...
func SendSnsSms(phoneNumbers []string, message, awsRegion, awsProfile string, development bool) *apperror.AppInfo {

	//get the caller and callee (if any) function names
	fname := logger.GetFuncName()

	//performance analysis - begin
	trace := &runstat.RunInfo{Name: fname, StartTime: time.Now()}
	defer trace.MeasureRuntime()

	//ready to set attributes
	attrs := map[string]*sns.MessageAttributeValue{}

	//set it to Transactional vs Promotional
	attrs["AWS.SNS.SMS.SMSType"] = &sns.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String("Transactional"),
		//StringValue: aws.String("Promotional"),
	}

	var sess *session.Session
	var service *sns.SNS
	var err error

	//set the aws profile to use for development
	//NOTE: without the region setting, it fails
	if development {
		sess, err = session.NewSession(&aws.Config{
			Region:      aws.String(awsRegion),
			Credentials: credentials.NewSharedCredentials("", awsProfile),
		})

		//check if credentials found
		_, err = sess.Config.Credentials.Get()
		if err != nil {
			a := &apperror.AppInfo{Msg: err}
			a.LogError(a.Error("sess.Config.Credentials.Get()"))
			trace.SetEndTime(time.Now())
			return a
		}

		service = sns.New(sess)

	} else {

		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(awsRegion)},
		)

		service = sns.New(sess)

		_, err = service.Config.Credentials.Get()

		if err != nil {
			a := &apperror.AppInfo{Msg: err}
			a.LogError(a.Error("service.Config.Credentials.Get()"))
			trace.SetEndTime(time.Now())
			return a
		}
	}

	for _, phoneNumber := range phoneNumbers {

		phoneNumber = "+1" + strings.Join(strings.Split(phoneNumber, "-"), "")

		params := &sns.PublishInput{
			Message:           aws.String(message),
			PhoneNumber:       aws.String(phoneNumber),
			MessageAttributes: attrs,
		}
		resp, err := service.Publish(params)

		if err != nil {
			a := &apperror.AppInfo{Msg: err}
			a.LogError(a.Error("sns.PublishInput()"))
			trace.SetEndTime(time.Now())
			return a
		}
		logger.L.Info.Printf("SMS sent successfully[%v] to phone[%v]", resp, phoneNumber)
	}

	//performance analysis - end
	trace.SetEndTime(time.Now())

	return nil
}
