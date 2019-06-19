// Copyright 2016-2019 Alex Stocks
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hessian

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestEncList(t *testing.T) {
	var (
		list []interface{}
		err  error
		e    *Encoder
		d    *Decoder
		res  interface{}
	)

	e = NewEncoder()
	list = []interface{}{100, 10.001, "hello", []byte{0, 2, 4, 6, 8, 10}, true, nil, false}
	e.Encode(list)
	assert.NotEqual(t, 0, len(e.Buffer()))

	d = NewDecoder(e.Buffer())
	res, err = d.Decode()
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%v", list), fmt.Sprintf("%v", res))

	// typed list - int32
	e = NewEncoder()
	list_1 := []int32{1, 2, 3}
	e.Encode(list_1)
	assert.NotEqual(t, 0, len(e.Buffer()))

	d = NewDecoder(e.Buffer())
	res, err = d.Decode()
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(res, list_1))

}

func TestList(t *testing.T) {
	RegisterPOJOs(new(A0), new(A1))

	testDecodeFramework(t, "replyTypedFixedList_0", []string{})
	testDecodeFramework(t, "replyTypedFixedList_1", []string{"1"})
	testDecodeFramework(t, "replyTypedFixedList_7", []string{"1", "2", "3", "4", "5", "6", "7"})
	testDecodeFramework(t, "replyTypedFixedList_8", []string{"1", "2", "3", "4", "5", "6", "7", "8"})
	testDecodeFramework(t, "replyUntypedFixedList_0", []interface{}{})
	testDecodeFramework(t, "replyUntypedFixedList_1", []interface{}{"1"})
	testDecodeFramework(t, "replyUntypedFixedList_7", []interface{}{"1", "2", "3", "4", "5", "6", "7"})
	testDecodeFramework(t, "replyUntypedFixedList_8", []interface{}{"1", "2", "3", "4", "5", "6", "7", "8"})

	testDecodeFramework(t, "customReplyTypedFixedListHasNull", []interface{}{new(A0), new(A1), nil})
	testDecodeFramework(t, "customReplyTypedVariableListHasNull", []interface{}{new(A0), new(A1), nil})
	testDecodeFramework(t, "customReplyUntypedFixedListHasNull", []interface{}{new(A0), new(A1), nil})
	testDecodeFramework(t, "customReplyUntypedVariableListHasNull", []interface{}{new(A0), new(A1), nil})

	testDecodeFramework(t, "customReplyTypedFixedList_A0", []*A0{new(A0), new(A0), nil})
	testDecodeFramework(t, "customReplyTypedVariableList_A0", []*A0{new(A0), new(A0), nil})

	testDecodeFramework(t, "customReplyTypedFixedList_int", []int32{1, 2, 3})
	testDecodeFramework(t, "customReplyTypedVariableList_int", []int32{1, 2, 3})

	testDecodeFramework(t, "customReplyTypedFixedList_long", []int64{1, 2, 3})
	testDecodeFramework(t, "customReplyTypedVariableList_long", []int64{1, 2, 3})

	testDecodeFramework(t, "customReplyTypedFixedList_float", []float64{1, 2, 3})
	testDecodeFramework(t, "customReplyTypedVariableList_float", []float64{1, 2, 3})

	testDecodeFramework(t, "customReplyTypedFixedList_double", []float64{1, 2, 3})
	testDecodeFramework(t, "customReplyTypedVariableList_double", []float64{1, 2, 3})

	testDecodeFramework(t, "customReplyTypedFixedList_short", []int32{1, 2, 3})
	testDecodeFramework(t, "customReplyTypedVariableList_short", []int32{1, 2, 3})

	testDecodeFramework(t, "customReplyTypedFixedList_char", []string{"1", "2", "3"})
	testDecodeFramework(t, "customReplyTypedVariableList_char", []string{"1", "2", "3"})

	testDecodeFramework(t, "customReplyTypedFixedList_boolean", []bool{true, false, true})
	testDecodeFramework(t, "customReplyTypedVariableList_boolean", []bool{true, false, true})

	testDecodeFramework(t, "customReplyTypedFixedList_date", []time.Time{time.Unix(1560864, 0),
		time.Unix(1560864, 0), time.Unix(1560864, 0)})
	testDecodeFramework(t, "customReplyTypedVariableList_date", []time.Time{time.Unix(1560864, 0),
		time.Unix(1560864, 0), time.Unix(1560864, 0)})
}

func TestListEncode(t *testing.T) {
	RegisterPOJOs(new(A0))

	testJavaDecode(t, "argUntypedFixedList_0", []interface{}{})
	testJavaDecode(t, "argUntypedFixedList_1", []interface{}{"1"})
	testJavaDecode(t, "argUntypedFixedList_7", []interface{}{"1", "2", "3", "4", "5", "6", "7"})
	testJavaDecode(t, "argUntypedFixedList_8", []interface{}{"1", "2", "3", "4", "5", "6", "7", "8"})

	testJavaDecode(t, "customArgUntypedFixedListHasNull", []interface{}{new(A0), new(A1), nil})

	testJavaDecode(t, "customArgTypedFixedList", []*A0{new(A0)})

	testJavaDecode(t, "argTypedFixedList_0", []string{})
	testJavaDecode(t, "argTypedFixedList_7", []string{"1", "2", "3", "4", "5", "6", "7"})

	testJavaDecode(t, "customArgTypedFixedList_short_0", []int8{})
	testJavaDecode(t, "customArgTypedFixedList_short_7", []int8{1, 2, 3, 4, 5, 6, 7})
	testJavaDecode(t, "customArgTypedFixedList_short_0", []int16{})
	testJavaDecode(t, "customArgTypedFixedList_short_7", []int16{1, 2, 3, 4, 5, 6, 7})
	testJavaDecode(t, "customArgTypedFixedList_short_0", []uint16{})
	testJavaDecode(t, "customArgTypedFixedList_short_7", []uint16{1, 2, 3, 4, 5, 6, 7})

	testJavaDecode(t, "customArgTypedFixedList_int_0", []int32{})
	testJavaDecode(t, "customArgTypedFixedList_int_7", []uint32{1, 2, 3, 4, 5, 6, 7})

	testJavaDecode(t, "customArgTypedFixedList_long_0", []int{})
	testJavaDecode(t, "customArgTypedFixedList_long_7", []int{1, 2, 3, 4, 5, 6, 7})
	testJavaDecode(t, "customArgTypedFixedList_long_0", []uint{})
	testJavaDecode(t, "customArgTypedFixedList_long_7", []uint{1, 2, 3, 4, 5, 6, 7})
	testJavaDecode(t, "customArgTypedFixedList_long_0", []int64{})
	testJavaDecode(t, "customArgTypedFixedList_long_7", []int64{1, 2, 3, 4, 5, 6, 7})
	testJavaDecode(t, "customArgTypedFixedList_long_0", []uint64{})
	testJavaDecode(t, "customArgTypedFixedList_long_7", []uint64{1, 2, 3, 4, 5, 6, 7})

	testJavaDecode(t, "customArgTypedFixedList_float_0", []float32{})
	testJavaDecode(t, "customArgTypedFixedList_float_7", []float32{1, 2, 3, 4, 5, 6, 7})

	testJavaDecode(t, "customArgTypedFixedList_double_0", []float64{})
	testJavaDecode(t, "customArgTypedFixedList_double_7", []float64{1, 2, 3, 4, 5, 6, 7})

	testJavaDecode(t, "customArgTypedFixedList_boolean_0", []bool{})
	testJavaDecode(t, "customArgTypedFixedList_boolean_7", []bool{true, false, true, false, true, false, true})

	testJavaDecode(t, "customArgTypedFixedList_date_0", []time.Time{})
	testJavaDecode(t, "customArgTypedFixedList_date_3", []time.Time{time.Unix(1560864, 0),
		time.Unix(1560864, 0), time.Unix(1560864, 0)})

}
