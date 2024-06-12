package err

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ErrTestSuite struct {
	suite.Suite
}

func (suite *ErrTestSuite) TestNew() {
	e := New()
	assert.NotNil(suite.T(), e)
	assert.Nil(suite.T(), e.(*errs).err)
	assert.Equal(suite.T(), uint32(0), e.(*errs).code)
	assert.Empty(suite.T(), e.(*errs).msg)
}

func (suite *ErrTestSuite) TestErrSetAndAdd() {
	e := New().Err("initial error")
	assert.Equal(
		suite.T(), "initial error;", e.(*errs).err.Error(),
	)

	e.Err("second error")
	assert.Equal(
		suite.T(),
		"second error; initial error;",
		e.(*errs).err.Error(),
	)

	e.Err("third error")
	assert.Equal(
		suite.T(),
		"third error; second error; initial error;",
		e.(*errs).err.Error(),
	)
}

func (suite *ErrTestSuite) TestFailed() {
	e := New().Failed("operation")
	assert.Equal(
		suite.T(),
		"operation failed;",
		e.(*errs).err.Error(),
	)
}

func (suite *ErrTestSuite) TestInfoSetAndAdd() {
	e := New().Info("context", "detail")
	assert.Equal(
		suite.T(),
		"context: [detail];",
		e.(*errs).err.Error(),
	)

	e.Info("additional context", "more detail")
	assert.Equal(
		suite.T(),
		"context: [detail]; additional context: [more detail];",
		e.(*errs).err.Error(),
	)
}

func (suite *ErrTestSuite) TestCode() {
	e := New().Code(404)
	assert.Equal(suite.T(), uint32(404), e.(*errs).code)

	e.Code(500)
	assert.Equal(suite.T(), uint32(500), e.(*errs).code)
}

func (suite *ErrTestSuite) TestMsgSetAndAdd() {
	e := New().Msg("initial message")
	assert.Equal(suite.T(), "initial message", e.(*errs).msg)

	e.Msg("additional message")
	assert.Equal(
		suite.T(),
		"initial message; additional message",
		e.(*errs).msg,
	)
}

func (suite *ErrTestSuite) TestGetMsg() {
	e := New().Msg("test message")
	assert.Equal(
		suite.T(), "test message", e.GetMsg(),
	)
}

func (suite *ErrTestSuite) TestGetCode() {
	e := New().Code(200)
	assert.Equal(suite.T(), "200", e.GetCode())
}

func (suite *ErrTestSuite) TestGetCodeUint32() {
	e := New().Code(1234)
	assert.Equal(
		suite.T(), uint32(1234), e.GetCodeUint32(),
	)
}

func (suite *ErrTestSuite) TestGetCodeInt32() {
	e := New().Code(1234)
	assert.Equal(suite.T(), int32(1234), e.GetCodeInt32())
}

func (suite *ErrTestSuite) TestError() {
	e := New().Err("test error").Code(400).Msg("bad request")
	expected := `{"err": 'test error;',"code": 400,"msg": "bad request"}`
	assert.Equal(suite.T(), expected, e.Error())
}

func (suite *ErrTestSuite) TestString() {
	e := New().Err("test string").Code(500).Msg("server error")
	expected := `{"err": 'test string;',"code": 500,"msg": "server error"}`
	assert.Equal(suite.T(), expected, e.String())
}

func TestErrTestSuite(t *testing.T) {
	suite.Run(t, new(ErrTestSuite))
}
