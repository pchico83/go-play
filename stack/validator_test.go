package stack_test

import (
	log "github.com/Sirupsen/logrus"
	"github.com/pchico83/d2k8/stack"
	. "gopkg.in/check.v1"
	"io/ioutil"
)

func (s *MySuite) TestFullStackValidation(c *C) {
	data := `web:
    image: ubuntu
    command: ls
    links:
        - db:database
        - database
    ports:
        - 80:80
        - 443:4443
    volumes:
        - /tmp1
        - /tmp2:/tmp2
        - /tmp3:/tmp3:ro
        - /tmp4:/tmp4:rw
    volumes_from:
        - database
    environment:
        - PATH
        - TEST=test
    entrypoint: ls
    mem_limit: 1000m
    privileged: true
    restart: always
    strategy: every_node
    scale: 2
    tags:
        - node-1
        - staging
database:
    image: mysql
`
	err := ioutil.WriteFile(s.StackFile, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = stack.CreateCmd("test", s.StackFile)
	c.Assert(err, IsNil)
	stackObj, err := stack.Read("test")
	if err != nil {
		log.Fatal(err)
	}
	s1 := stackObj.Definition["web"]
	s2 := stackObj.Definition["database"]

	c.Assert(s1.Image, Equals, "ubuntu")
	c.Assert(s1.Command, Equals, "ls")
	c.Assert(s1.Links, DeepEquals, []string{"db:database", "database"})
	c.Assert(s1.Ports, DeepEquals, []string{"80:80", "443:4443"})
	c.Assert(s1.Volumes, DeepEquals, []string{"/tmp1", "/tmp2:/tmp2", "/tmp3:/tmp3:ro", "/tmp4:/tmp4:rw"})
	c.Assert(s1.Volumes_from, DeepEquals, []string{"database"})
	c.Assert(s1.Environment, DeepEquals, []string{"PATH", "TEST=test"})
	c.Assert(s1.Entrypoint, Equals, "ls")
	c.Assert(s1.Mem_limit, Equals, "1000m")
	c.Assert(s1.Privileged, Equals, true)
	c.Assert(s1.Restart, Equals, "always")
	c.Assert(s1.Strategy, Equals, "every_node")
	c.Assert(s1.Scale, Equals, 2)
	c.Assert(s1.Tags, DeepEquals, []string{"node-1", "staging"})
	c.Assert(s2.Image, Equals, "mysql")
}

func (s *MySuite) TestMinimumStackValidation(c *C) {
	data := `web:
    image: ubuntu
`
	err := ioutil.WriteFile(s.StackFile, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = stack.CreateCmd("test", s.StackFile)
	c.Assert(err, IsNil)
	stackObj, err := stack.Read("test")
	if err != nil {
		log.Fatal(err)
	}

	s1 := stackObj.Definition["web"]
	c.Assert(s1.Image, Equals, "ubuntu")
	c.Assert(s1.Command, Equals, "")
	c.Assert(s1.Links, DeepEquals, []string(nil))
	c.Assert(s1.Ports, DeepEquals, []string(nil))
	c.Assert(s1.Volumes, DeepEquals, []string(nil))
	c.Assert(s1.Volumes_from, DeepEquals, []string(nil))
	c.Assert(s1.Environment, DeepEquals, []string(nil))
	c.Assert(s1.Entrypoint, Equals, "")
	c.Assert(s1.Mem_limit, Equals, "")
	c.Assert(s1.Privileged, Equals, false)
	c.Assert(s1.Restart, Equals, "no")
	c.Assert(s1.Strategy, Equals, "balance")
	c.Assert(s1.Scale, Equals, 1)
	c.Assert(s1.Tags, DeepEquals, []string(nil))
}

func (s *MySuite) TestWrongStackName(c *C) {
	data := `web:
    image: ubuntu
`
	err := ioutil.WriteFile(s.StackFile, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = stack.CreateCmd("test-stack", s.StackFile)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Stack name not valid: 'test-stack' does not match pattern '^[0-9a-zA-Z]{1,25}$'\n")
}

func (s *MySuite) TestStackWithTabs(c *C) {
	data := `web:
	image: ubuntu
`
	err := ioutil.WriteFile(s.StackFile, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = stack.CreateCmd("test", s.StackFile)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "'\t' characters are not allowed in YAML files")
}

func (s *MySuite) TestWrongServiceName(c *C) {
	data := `web-test:
    image: ubuntu
`
	err := ioutil.WriteFile(s.StackFile, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = stack.CreateCmd("Test01", s.StackFile)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Service name not valid: 'web-test' does not match pattern '^[0-9a-zA-Z]{1,25}$'\n")
}

func (s *MySuite) TestWrongLinks(c *C) {
	data := `web:
    image: ubuntu
    links:
        - data-base
        - d-b:database
`
	err := ioutil.WriteFile(s.StackFile, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = stack.CreateCmd("test", s.StackFile)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Service web not valid. see errors:\n"+
		"links.0: 'data-base' does not match pattern '^[0-9a-zA-Z]{1,20}(:[0-9a-zA-Z]{1,20})?$'\n"+
		"links.1: 'd-b:database' does not match pattern '^[0-9a-zA-Z]{1,20}(:[0-9a-zA-Z]{1,20})?$'\n")
}

func (s *MySuite) TestWrongPorts(c *C) {
	data := `web:
    image: ubuntu
    ports:
        - 0:10
        - 888-888
        - 12a34
`
	err := ioutil.WriteFile(s.StackFile, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = stack.CreateCmd("test", s.StackFile)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Service web not valid. see errors:\n"+
		"ports.0: '0:10' does not match pattern '^[1-9][0-9]{0,4}(:[1-9][0-9]{0,4})?$'\n"+
		"ports.1: '888-888' does not match pattern '^[1-9][0-9]{0,4}(:[1-9][0-9]{0,4})?$'\n"+
		"ports.2: '12a34' does not match pattern '^[1-9][0-9]{0,4}(:[1-9][0-9]{0,4})?$'\n")
}

func (s *MySuite) TestLinksDoNotExist(c *C) {
	data := `web:
    image: ubuntu
    links:
        - no
`
	err := ioutil.WriteFile(s.StackFile, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = stack.CreateCmd("test", s.StackFile)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Linked service 'no' in service 'web' does not exist\n")

	data = `web:
    image: ubuntu
    links:
        - alias:no
`
	err = ioutil.WriteFile(s.StackFile, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = stack.CreateCmd("test", s.StackFile)
	c.Assert(err, NotNil)
	c.Assert(err.Error(), Equals, "Linked service 'no' in service 'web' does not exist\n")
}
