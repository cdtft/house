package repository

import "testing"

func TestId_test(t *testing.T) {
	type fields struct {
		Id int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "test",
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := &Id{
				Id: tt.fields.Id,
			}
			id.test()
		})
	}
}
