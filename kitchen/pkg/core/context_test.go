package core_test

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"testing"

	. "kitchen/pkg/core"
	"kitchen/pkg/test"
)

var conf1 string = `
ingestInterval: 200
shelfCapacity:
  hot: 5
  cold: 5
  frozen: 5
  overflow: 10
numOfCouriers: 12
minPickDuration: 5
maxPickDuration: 10
`

var conf2 string = `
ingestInterval: "asdf"
shelfCapacity:
  hot: 5
  cold: 5
  frozen: 5
  overflow: 10
numOfCouriers: 12
minPickDuration: 5
maxPickDuration: 10
`
var conf3 string = `
ingestInterval: 200
shelfCapacity:
  hot: 5
  cold: 5
  frozen: 5
  overflow: 10
numOfCouriers: 12
minPickDuration: 5
maxPickDuration: 3
`

func setup(conf string) error {
	os.MkdirAll("./temp", os.ModePerm)
	err := ioutil.WriteFile("./temp/conf.yaml", []byte(conf), 0777)
	if err != nil {
		return err
	}
	return nil
}

func teardown() error {
	return os.RemoveAll("./temp")
}

func TestContext(tt *testing.T) {
	t := test.NewTestWithoutContext(tt)

	t.Run("default config", func() {
		ctx, err := NewContext("")
		t.Expect(err).To(BeNil())
		t.Expect(ctx.NumOfCouriers).To(Equal(20))
	})
	t.Run("load config", func() {
		err := setup(conf1)
		t.Expect(err).NotTo(HaveOccurred())
		defer teardown()
		ctx, err := NewContext("./temp/conf.yaml")
		t.Expect(err).To(BeNil())
		t.Expect(ctx.NumOfCouriers).To(Equal(12))
	})
	t.Run("load invalid config", func() {
		err := setup(conf2)
		t.Expect(err).NotTo(HaveOccurred())
		defer teardown()
		_, err = NewContext("./temp/conf.yaml")
		t.Expect(err).NotTo(BeNil())
	})

	t.Run("load config with invalid value", func() {
		err := setup(conf3)
		t.Expect(err).NotTo(HaveOccurred())
		defer teardown()
		_, err = NewContext("./temp/conf.yaml")
		t.Expect(err).NotTo(BeNil())
	})

}
