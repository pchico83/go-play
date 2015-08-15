package stack_test

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pchico83/d2k8/stack"
	"github.com/pchico83/d2k8/utils"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct {
	StackFile string
	Data      string
}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpTest(c *C) {
	utils.SetUpTestCommon()
	folder := os.Getenv("ELORA_STORAGE_FOLDER")
	s.StackFile = filepath.Join(folder, "elora.yml")
	s.Data = `web:
    image: ubuntu
`
	err := ioutil.WriteFile(s.StackFile, []byte(s.Data), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *MySuite) TearDownTest(c *C) {
	utils.TearDownTestCommon()
}

func (s *MySuite) TestStackCreate(c *C) {
	err := stack.CreateCmd("test", s.StackFile)
	c.Assert(err, IsNil)
	data, _ := ioutil.ReadFile(stack.DefinitionFilePath("test"))
	c.Assert(string(data), Equals, s.Data)
	err = stack.CreateCmd("test", s.StackFile)
	c.Assert(err, NotNil)
}

func (s *MySuite) TestStackExport(c *C) {
	err := stack.CreateCmd("test", s.StackFile)
	c.Assert(err, IsNil)
	err = stack.ExportCmd("test")
	c.Assert(err, IsNil)
}

func (s *MySuite) TestStackList(c *C) {
	err := stack.CreateCmd("test1", s.StackFile)
	c.Assert(err, IsNil)
	err = stack.CreateCmd("test2", s.StackFile)
	c.Assert(err, IsNil)
	err = stack.ListCmd()
	c.Assert(err, IsNil)
}

func (s *MySuite) TestStackUpdate(c *C) {
	err := stack.CreateCmd("test", s.StackFile)
	c.Assert(err, IsNil)
	data := `web:
    image: ubuntu
    privileged: true
`
	err = ioutil.WriteFile(s.StackFile, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = stack.UpdateCmd("test", s.StackFile)
	c.Assert(err, IsNil)
	content, err := ioutil.ReadFile(stack.DefinitionFilePath("test"))
	if err != nil {
		log.Fatal(err)
	}
	c.Assert(string(content), Equals, data)
	err = stack.UpdateCmd("test2", s.StackFile)
	c.Assert(err, NotNil)
}

func (s *MySuite) TestStackRemove(c *C) {
	err := stack.CreateCmd("test", s.StackFile)
	c.Assert(err, IsNil)
	err = stack.RemoveCmd("test")
	c.Assert(err, IsNil)
	filePath := stack.DefinitionFilePath("test")
	_, err = os.Stat(filePath)
	c.Assert(err, NotNil)
	err = stack.RemoveCmd("test")
	c.Assert(err, NotNil)
}
