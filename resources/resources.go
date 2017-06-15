package resources

import "fmt"

//go:generate counterfeiter -o ../fakes/fake_test_interface.go . MyTestInterface

type MyTestInterface interface {
	MyTestMethod(string) (string, error)
}

type interfImpl struct {
}

func NewInterfaceImpl() MyTestInterface {
	return &interfImpl{}
}
func (i *interfImpl) MyTestMethod(string) (string, error) {
	return "", nil
}

type MyStruct struct {
}

func NewMyStruct() *MyStruct {
	return &MyStruct{}
}

func (m *MyStruct) MyFunc(myS string, interfaceImpl MyTestInterface) error {
	mystring, err := interfaceImpl.MyTestMethod(myS)
	if err != nil {
		return err
	}

	fmt.Printf(mystring)
	return nil
}
