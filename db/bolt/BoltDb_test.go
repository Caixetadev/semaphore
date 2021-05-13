package bolt

import (
	"fmt"
	"github.com/ansible-semaphore/semaphore/db"
	"testing"
)

type Test1 struct {
	FistName string `db:"first_name" json:"firstName"`
	LastName string `db:"last_name" json:"lastName"`
	Password string `db:"-" json:"password"`
	PasswordRepeat string `db:"-" json:"passwordRepeat"`
	PasswordHash string `db:"password" json:"-"`
}

func TestMarshalObject(t *testing.T) {
	test1 := Test1{
		FistName: "Denis",
		LastName: "Gukov",
		Password: "1234556",
		PasswordRepeat: "123456",
		PasswordHash: "9347502348723",
	}

	bytes, err := marshalObject(test1)

	if err != nil {
		t.Fatal(fmt.Errorf("function returns error: " + err.Error()))
	}

	str := string(bytes)
	if str != `{"first_name":"Denis","last_name":"Gukov","password":"9347502348723"}` {
		t.Fatal(fmt.Errorf("incorrect marshalling result"))
	}

	fmt.Println(str)
}

func TestUnmarshalObject(t *testing.T) {
	test1 := Test1{}
	data := `{
	"first_name": "Denis", 
	"last_name": "Gukov",
	"password": "9347502348723"
}`
	err := unmarshalObject([]byte(data), &test1)
	if err != nil {
		t.Fatal(fmt.Errorf("function returns error: " + err.Error()))
	}
	if test1.FistName != "Denis" ||
		test1.LastName != "Gukov" ||
		test1.Password != "" ||
		test1.PasswordRepeat != "" ||
		test1.PasswordHash != "9347502348723" {
		t.Fatal(fmt.Errorf("object unmarshalled incorrectly"))
	}
}

func TestSortObjects(t *testing.T) {
	objects := []db.Inventory{
		{
			ID: 1,
			Name: "x",
		},
		{
			ID: 2,
			Name: "a",
		},
		{
			ID: 3,
			Name: "d",
		},
		{
			ID: 4,
			Name: "b",
		},
		{
			ID: 5,
			Name: "r",
		},
	}

	err := sortObjects(&objects, "name", false)
	if err != nil {
		t.Fatal(err)
	}

	expected := objects[0].Name == "a" &&
		objects[1].Name == "b" &&
		objects[2].Name == "d" &&
		objects[3].Name == "r" &&
		objects[4].Name == "x"


	if !expected {
		t.Fatal(fmt.Errorf("objects not sorted"))
	}
}