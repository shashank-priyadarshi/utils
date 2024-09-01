package sql

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.ssnk.in/utils/types"
	"strings"
	"testing"
)

func Test_Create(t *testing.T) {

	//logger := ports.NewMockLogger(t)
	db, mock, err := sqlmock.New()
	assert.Equal(t, nil, err)

	handle := Handle(db)

	tests := []types.Test{{
		Name: "happy path",
		TestCase: func(t *testing.T) {
			query := strings.TrimSpace(`INSERT INTO person (name, age) VALUES (test, 3000)`)

			mock.ExpectExec(query).WillReturnResult(nil)

			_, err = handle.Create(context.Background(), query)
			assert.Equal(t, nil, err)
		},
	}}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}

func Test_Query(t *testing.T) {

	//logger := ports.NewMockLogger(t)
	db, mock, err := sqlmock.New()
	assert.Equal(t, nil, err)

	handle := Handle(db)

	tests := []types.Test{{
		Name: "happy path",
		TestCase: func(t *testing.T) {
			mock.ExpectExec("")

			_, err = handle.Create(context.Background())
			assert.Equal(t, nil, err)
		},
	}}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}

func Test_Update(t *testing.T) {

	//logger := ports.NewMockLogger(t)
	db, mock, err := sqlmock.New()
	assert.Equal(t, nil, err)

	handle := Handle(db)

	tests := []types.Test{{
		Name: "happy path",
		TestCase: func(t *testing.T) {
			mock.ExpectExec("")

			_, err = handle.Create(context.Background())
			assert.Equal(t, nil, err)
		},
	}}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}

func Test_Delete(t *testing.T) {

	//logger := ports.NewMockLogger(t)
	db, mock, err := sqlmock.New()
	assert.Equal(t, nil, err)

	handle := Handle(db)

	tests := []types.Test{{
		Name: "happy path",
		TestCase: func(t *testing.T) {
			mock.ExpectExec("")

			_, err = handle.Create(context.Background())
			assert.Equal(t, nil, err)
		},
	}}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}
