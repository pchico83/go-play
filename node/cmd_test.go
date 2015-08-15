package node_test

import (
	"github.com/pchico83/d2k8/client"
	"github.com/pchico83/d2k8/cluster"
	"github.com/pchico83/d2k8/node"
	"github.com/pchico83/d2k8/utils"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"os"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpTest(c *C) {
	utils.SetUpTestCommon()
}

func (s *MySuite) TearDownTest(c *C) {
	utils.TearDownTestCommon()
}

func (s *MySuite) TestNodeCreate(c *C) {
	err := node.CreateCmd("test", "127.0.0.1", "/", "/id_rsa", "default", []string{"testing", "staging"})
	c.Assert(err, IsNil)
	data, _ := ioutil.ReadFile(node.FilePath("test"))
	c.Assert(string(data), Equals, "name: test\nurl: tcp://127.0.0.1:2376\ncert: /\nkey: /id_rsa\ncluster: default\ntags:\n- testing\n- staging\n")
	err = node.CreateCmd("test", "127.0.0.1", "/", "/id_rsa", "default", []string{"testing", "staging"})
	c.Assert(err, NotNil)
	err = node.CreateCmd("test2", "127.0.0.1", "/", "/id_rsa", "cluster", []string{"testing", "staging"})
	c.Assert(err, NotNil)
	err = cluster.CreateCmd("cluster", true, true, "pwd", true)
	c.Assert(err, IsNil)
	err = node.CreateCmd("test2", "127.0.0.1", "/", "/id_rsa", "cluster", []string{"testing", "staging"})
	c.Assert(err, IsNil)
	data, _ = ioutil.ReadFile(node.FilePath("test2"))
	c.Assert(string(data), Equals, "name: test2\nurl: tcp://127.0.0.1:2376\ncert: /\nkey: /id_rsa\ncluster: cluster\ntags:\n- testing\n- staging\n")
}

func (s *MySuite) TestNodeEnv(c *C) {
	err := node.CreateCmd("test", "127.0.0.1", "/", "/id_rsa", "default", []string{"testing", "staging"})
	c.Assert(err, IsNil)
	err = node.EnvCmd("test", false)
	c.Assert(err, IsNil)
	err = node.EnvCmd("test", true)
	c.Assert(err, IsNil)
}

func (s *MySuite) TestNodeInspect(c *C) {
	err := node.CreateCmd("test", "127.0.0.1", "/", "/id_rsa", "default", []string{"testing", "staging"})
	c.Assert(err, IsNil)
	err = node.InspectCmd("test")
	c.Assert(err, IsNil)
	client.Factory = client.ErrorPingFactory{}
	err = node.InspectCmd("test")
	c.Assert(err, IsNil)
}

func (s *MySuite) TestNodeList(c *C) {
	err := node.CreateCmd("test1", "127.0.0.1", "/", "/id_rsa", "default", []string{"testing", "staging"})
	c.Assert(err, IsNil)
	err = node.CreateCmd("test2", "127.0.0.1", "/", "/id_rsa", "default", []string{"testing", "staging"})
	c.Assert(err, IsNil)
	err = node.ListCmd()
	c.Assert(err, IsNil)
	client.Factory = client.ErrorPingFactory{}
	err = node.ListCmd()
	c.Assert(err, IsNil)
}

func (s *MySuite) TestNodeUpdate(c *C) {
	err := node.CreateCmd("test", "127.0.0.1", "/", "/id_rsa", "default", []string{"testing", "staging"})
	c.Assert(err, IsNil)
	err = node.UpdateCmd("test", "192.168.99.100", "/tmp", "/tmp/id_rsa", []string{"prod"})
	c.Assert(err, IsNil)
	data, _ := ioutil.ReadFile(node.FilePath("test"))
	c.Assert(string(data), Equals, "name: test\nurl: tcp://192.168.99.100:2376\ncert: /tmp\nkey: /tmp/id_rsa\ncluster: default\ntags:\n- prod\n")
	err = node.UpdateCmd("test2", "192.168.99.100", "/tmp", "/tmp/id_rsa", []string{"prod"})
	c.Assert(err, NotNil)
}

func (s *MySuite) TestNodeRemove(c *C) {
	err := node.CreateCmd("test", "127.0.0.1", "/", "/id_rsa", "default", []string{"testing", "staging"})
	c.Assert(err, IsNil)
	err = node.RemoveCmd("test")
	c.Assert(err, IsNil)
	filePath := node.FilePath("test")
	_, err = os.Stat(filePath)
	c.Assert(err, NotNil)
	err = node.RemoveCmd("test")
	c.Assert(err, NotNil)
}
