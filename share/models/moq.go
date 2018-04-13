package models

// moq -out mocks/mocks_test.go models IMyData
type IMyData interface {
	GetData(i int) (string, error)
}

type (

	MyData struct {
		I int
		S string
		A []int
		M map[string]string
	}
)
