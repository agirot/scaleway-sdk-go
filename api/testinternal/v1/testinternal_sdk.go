// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package testinternal provides methods and message types of the testinternal v1 API.
package testinternal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ fmt.Stringer
	_ json.Unmarshaler
	_ url.URL
	_ net.IP
	_ http.Header
	_ bytes.Reader
	_ time.Time
	_ = strings.Join

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

// API: this API allows us to test.
// Test API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// NoAuthAPI: no Auth Service.
type NoAuthAPI struct {
	client *scw.Client
}

// NewNoAuthAPI returns a NoAuthAPI object from a Scaleway client.
func NewNoAuthAPI(client *scw.Client) *NoAuthAPI {
	return &NoAuthAPI{
		client: client,
	}
}

// RegionalAPI: regional API.
type RegionalAPI struct {
	client *scw.Client
}

// NewRegionalAPI returns a RegionalAPI object from a Scaleway client.
func NewRegionalAPI(client *scw.Client) *RegionalAPI {
	return &RegionalAPI{
		client: client,
	}
}

// StreamAPI: stream Service.
type StreamAPI struct {
	client *scw.Client
}

// NewStreamAPI returns a StreamAPI object from a Scaleway client.
func NewStreamAPI(client *scw.Client) *StreamAPI {
	return &StreamAPI{
		client: client,
	}
}

// ZoneAPI: zone Test API.
type ZoneAPI struct {
	client *scw.Client
}

// NewZoneAPI returns a ZoneAPI object from a Scaleway client.
func NewZoneAPI(client *scw.Client) *ZoneAPI {
	return &ZoneAPI{
		client: client,
	}
}

type ComplexTestEnum string

const (
	ComplexTestEnumComplexZero = ComplexTestEnum("ComplexZero")
	ComplexTestEnumComplexONE  = ComplexTestEnum("ComplexONE")
	ComplexTestEnumComplexTWO  = ComplexTestEnum("ComplexTWO")
)

func (enum ComplexTestEnum) String() string {
	if enum == "" {
		// return default value if empty
		return "ComplexZero"
	}
	return string(enum)
}

func (enum ComplexTestEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ComplexTestEnum) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ComplexTestEnum(ComplexTestEnum(tmp).String())
	return nil
}

type ListCharactersRequestOrderBy string

const (
	ListCharactersRequestOrderByAsc  = ListCharactersRequestOrderBy("asc")
	ListCharactersRequestOrderByDesc = ListCharactersRequestOrderBy("desc")
)

func (enum ListCharactersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "asc"
	}
	return string(enum)
}

func (enum ListCharactersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCharactersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCharactersRequestOrderBy(ListCharactersRequestOrderBy(tmp).String())
	return nil
}

type MapEnum string

const (
	MapEnumMAPENUMFOO = MapEnum("MAP_ENUM_FOO")
	MapEnumMAPENUMBAR = MapEnum("MAP_ENUM_BAR")
	MapEnumMAPENUMBAZ = MapEnum("MAP_ENUM_BAZ")
)

func (enum MapEnum) String() string {
	if enum == "" {
		// return default value if empty
		return "MAP_ENUM_FOO"
	}
	return string(enum)
}

func (enum MapEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MapEnum) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MapEnum(MapEnum(tmp).String())
	return nil
}

type NullValue string

const (
	NullValueNULLVALUE = NullValue("NULL_VALUE")
)

func (enum NullValue) String() string {
	if enum == "" {
		// return default value if empty
		return "NULL_VALUE"
	}
	return string(enum)
}

func (enum NullValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NullValue) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NullValue(NullValue(tmp).String())
	return nil
}

type PostAllOptionalMessageNestedEnum string

const (
	PostAllOptionalMessageNestedEnumNEG         = PostAllOptionalMessageNestedEnum("NEG")
	PostAllOptionalMessageNestedEnumUNSPECIFIED = PostAllOptionalMessageNestedEnum("UNSPECIFIED")
	PostAllOptionalMessageNestedEnumFOO         = PostAllOptionalMessageNestedEnum("FOO")
	PostAllOptionalMessageNestedEnumBAR         = PostAllOptionalMessageNestedEnum("BAR")
	PostAllOptionalMessageNestedEnumBAZ         = PostAllOptionalMessageNestedEnum("BAZ")
)

func (enum PostAllOptionalMessageNestedEnum) String() string {
	if enum == "" {
		// return default value if empty
		return "UNSPECIFIED"
	}
	return string(enum)
}

func (enum PostAllOptionalMessageNestedEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PostAllOptionalMessageNestedEnum) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PostAllOptionalMessageNestedEnum(PostAllOptionalMessageNestedEnum(tmp).String())
	return nil
}

type PostAllTypesMessageNestedEnum string

const (
	// The fifth enum value
	PostAllTypesMessageNestedEnumNEG = PostAllTypesMessageNestedEnum("NEG")
	// The first enum value
	PostAllTypesMessageNestedEnumZERO = PostAllTypesMessageNestedEnum("ZERO")
	// The second enum value
	PostAllTypesMessageNestedEnumFOO = PostAllTypesMessageNestedEnum("FOO")
	// The third enum value
	PostAllTypesMessageNestedEnumBAR = PostAllTypesMessageNestedEnum("BAR")
	// The fourth enum value
	PostAllTypesMessageNestedEnumBAZ = PostAllTypesMessageNestedEnum("BAZ")
)

func (enum PostAllTypesMessageNestedEnum) String() string {
	if enum == "" {
		// return default value if empty
		return "ZERO"
	}
	return string(enum)
}

func (enum PostAllTypesMessageNestedEnum) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PostAllTypesMessageNestedEnum) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PostAllTypesMessageNestedEnum(PostAllTypesMessageNestedEnum(tmp).String())
	return nil
}

type TransientStatus string

const (
	TransientStatusUnknownStatus = TransientStatus("unknown_status")
	TransientStatusPending       = TransientStatus("pending")
	TransientStatusError         = TransientStatus("error")
	TransientStatusDone          = TransientStatus("done")
)

func (enum TransientStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum TransientStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TransientStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TransientStatus(TransientStatus(tmp).String())
	return nil
}

// Type: this is an enum type.
type Type string

const (
	// Unknown type
	TypeUnknown = Type("unknown")
	// Start1 XS type
	TypeStart1Xs = Type("start1_xs")
	// Start1 S type
	TypeStart1S = Type("start1_s")
	// Start1 M type
	TypeStart1M = Type("start1_m")
)

func (enum Type) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum Type) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Type) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Type(Type(tmp).String())
	return nil
}

type ComplexValidateMsg struct {
	Const string `json:"const"`

	Nested *ComplexValidateMsg `json:"nested"`

	IntConst int32 `json:"int_const"`

	BoolConst bool `json:"bool_const"`

	FloatVal *float32 `json:"float_val"`

	DurVal *scw.Duration `json:"dur_val"`

	TsVal *time.Time `json:"ts_val"`

	Another *ComplexValidateMsg `json:"another"`

	FloatConst float32 `json:"float_const"`

	DoubleIn float64 `json:"double_in"`
	// EnumConst: default value: ComplexZero
	EnumConst ComplexTestEnum `json:"enum_const"`

	AnyVal interface{} `json:"any_val"`

	RepTsVal []*time.Time `json:"rep_ts_val"`

	MapVal map[int32]string `json:"map_val"`

	BytesVal []byte `json:"bytes_val"`

	// Precisely one of X, Y must be set.
	X *string `json:"x,omitempty"`

	// Precisely one of X, Y must be set.
	Y *int32 `json:"y,omitempty"`
}

// EchoMessage: echo message.
type EchoMessage struct {
	Str string `json:"str"`

	Strs []string `json:"strs"`
}

// GetEnumMessage: get enum message.
type GetEnumMessage struct {
	// Type: default value: unknown
	Type Type `json:"type"`
}

type GetZoneResponse struct {
	Zone scw.Zone `json:"zone"`
}

// ListCharactersResponse: list characters response.
type ListCharactersResponse struct {
	TotalCount uint32 `json:"total_count"`

	Characters []string `json:"characters"`
}

// MetadataResponse: metadata response.
type MetadataResponse struct {
	Metadata map[string]string `json:"metadata"`
}

type PatchEnumMessage struct {
	// Type: default value: unknown
	Type Type `json:"type"`
}

type PostAllMapTypesMessage struct {
	MapInt32Int32 map[int32]int32 `json:"map_int32_int32"`

	MapInt64Int64 map[int64]int64 `json:"map_int64_int64"`

	MapUint32Uint32 map[uint32]uint32 `json:"map_uint32_uint32"`

	MapUint64Uint64 map[uint64]uint64 `json:"map_uint64_uint64"`

	MapSint32Sint32 map[int32]int32 `json:"map_sint32_sint32"`

	MapSint64Sint64 map[int64]int64 `json:"map_sint64_sint64"`

	MapFixed32Fixed32 map[uint32]uint32 `json:"map_fixed32_fixed32"`

	MapFixed64Fixed64 map[uint64]uint64 `json:"map_fixed64_fixed64"`

	MapSfixed32Sfixed32 map[int32]int32 `json:"map_sfixed32_sfixed32"`

	MapSfixed64Sfixed64 map[int64]int64 `json:"map_sfixed64_sfixed64"`

	MapInt32Float map[int32]float32 `json:"map_int32_float"`

	MapInt32Double map[int32]float64 `json:"map_int32_double"`

	MapStringString map[string]string `json:"map_string_string"`

	MapInt32Bytes map[int32][]byte `json:"map_int32_bytes"`

	MapInt32Enum map[int32]MapEnum `json:"map_int32_enum"`

	MapInt32AllTypes map[int32]*PostAllTypesMessage `json:"map_int32_all_types"`

	MapInt32IP map[int32]net.IP `json:"map_int32_ip"`

	MapInt32StdDuration map[int32]*time.Duration `json:"map_int32_std_duration"`

	MapInt32StdLongDuration map[int32]*time.Duration `json:"map_int32_std_long_duration"`

	MapInt32Size map[int32]scw.Size `json:"map_int32_size"`

	MapInt32Uint64Size map[int32]scw.Size `json:"map_int32_uint64_size"`

	MapInt32Uint64valueSize map[int32]*scw.Size `json:"map_int32_uint64value_size"`

	MapInt32StringIP map[int32]net.IP `json:"map_int32_string_ip"`

	MapInt32StringValueIP map[int32]*net.IP `json:"map_int32_string_value_ip"`

	MapInt32IPv4 map[int32]net.IP `json:"map_int32_ipv4"`

	MapInt32StringIPv4 map[int32]net.IP `json:"map_int32_string_ipv4"`

	MapInt32StringValueIPv4 map[int32]*net.IP `json:"map_int32_string_value_ipv4"`

	MapInt32IPv6 map[int32]net.IP `json:"map_int32_ipv6"`

	MapInt32StringIPv6 map[int32]net.IP `json:"map_int32_string_ipv6"`

	MapInt32StringValueIPv6 map[int32]*net.IP `json:"map_int32_string_value_ipv6"`

	MapInt32StringsValue map[int32]*[]string `json:"map_int32_strings_value"`

	MapInt32Duration map[int32]*scw.Duration `json:"map_int32_duration"`
}

func (m *PostAllMapTypesMessage) UnmarshalJSON(b []byte) error {
	type tmpType PostAllMapTypesMessage
	tmp := struct {
		tmpType

		TmpMapInt32StdDuration     marshaler.Durationint32Map     `json:"map_int32_std_duration"`
		TmpMapInt32StdLongDuration marshaler.LongDurationint32Map `json:"map_int32_std_long_duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostAllMapTypesMessage(tmp.tmpType)

	m.MapInt32StdDuration = tmp.TmpMapInt32StdDuration.Standard()
	m.MapInt32StdLongDuration = tmp.TmpMapInt32StdLongDuration.Standard()
	return nil
}

func (m PostAllMapTypesMessage) MarshalJSON() ([]byte, error) {
	type tmpType PostAllMapTypesMessage
	tmp := struct {
		tmpType

		TmpMapInt32StdDuration     marshaler.Durationint32Map     `json:"map_int32_std_duration"`
		TmpMapInt32StdLongDuration marshaler.LongDurationint32Map `json:"map_int32_std_long_duration"`
	}{
		tmpType: tmpType(m),

		TmpMapInt32StdDuration:     marshaler.NewDurationint32Map(m.MapInt32StdDuration),
		TmpMapInt32StdLongDuration: marshaler.NewLongDurationint32Map(m.MapInt32StdLongDuration),
	}
	return json.Marshal(tmp)
}

type PostAllOptionalMessage struct {

	// Precisely one of OptionalInt32 must be set.
	OptionalInt32 *int32 `json:"optional_int32,omitempty"`

	// Precisely one of OptionalInt64 must be set.
	OptionalInt64 *int64 `json:"optional_int64,omitempty"`

	// Precisely one of OptionalUint32 must be set.
	OptionalUint32 *uint32 `json:"optional_uint32,omitempty"`

	// Precisely one of OptionalUint64 must be set.
	OptionalUint64 *uint64 `json:"optional_uint64,omitempty"`

	// Precisely one of OptionalSint32 must be set.
	OptionalSint32 *int32 `json:"optional_sint32,omitempty"`

	// Precisely one of OptionalSint64 must be set.
	OptionalSint64 *int64 `json:"optional_sint64,omitempty"`

	// Precisely one of OptionalFixed32 must be set.
	OptionalFixed32 *uint32 `json:"optional_fixed32,omitempty"`

	// Precisely one of OptionalFixed64 must be set.
	OptionalFixed64 *uint64 `json:"optional_fixed64,omitempty"`

	// Precisely one of OptionalSfixed32 must be set.
	OptionalSfixed32 *int32 `json:"optional_sfixed32,omitempty"`

	// Precisely one of OptionalSfixed64 must be set.
	OptionalSfixed64 *int64 `json:"optional_sfixed64,omitempty"`

	// Precisely one of OptionalFloat must be set.
	OptionalFloat *float32 `json:"optional_float,omitempty"`

	// Precisely one of OptionalDouble must be set.
	OptionalDouble *float64 `json:"optional_double,omitempty"`

	// Precisely one of OptionalBool must be set.
	OptionalBool *bool `json:"optional_bool,omitempty"`

	// Precisely one of OptionalString must be set.
	OptionalString *string `json:"optional_string,omitempty"`

	// Precisely one of OptionalBytes must be set.
	OptionalBytes *[]byte `json:"optional_bytes,omitempty"`

	// Precisely one of OptionalCord must be set.
	OptionalCord *string `json:"optional_cord,omitempty"`

	// Precisely one of OptionalNestedMessageWithOptional must be set.
	OptionalNestedMessageWithOptional *PostAllOptionalMessageNestedMessageWithOptional `json:"optional_nested_message_with_optional,omitempty"`

	// Precisely one of LazyNestedMessageWithOptional must be set.
	LazyNestedMessageWithOptional *PostAllOptionalMessageNestedMessageWithOptional `json:"lazy_nested_message_with_optional,omitempty"`
	// OptionalNestedEnum: default value: UNSPECIFIED
	// Precisely one of OptionalNestedEnum must be set.
	OptionalNestedEnum *PostAllOptionalMessageNestedEnum `json:"optional_nested_enum,omitempty"`

	// Precisely one of OptionalNestedMessage must be set.
	OptionalNestedMessage *PostAllOptionalMessageNestedMessage `json:"optional_nested_message,omitempty"`

	SingularInt32 int32 `json:"singular_int32"`

	SingularInt64 int64 `json:"singular_int64"`

	NestedMessage *PostAllOptionalMessageNestedMessage `json:"nested_message"`

	// Precisely one of SingularDoubleValue must be set.
	SingularDoubleValue *float64 `json:"singular_double_value,omitempty"`

	// Precisely one of SingularFloatValue must be set.
	SingularFloatValue *float32 `json:"singular_float_value,omitempty"`

	// Precisely one of SingularInt64Value must be set.
	SingularInt64Value *int64 `json:"singular_int64_value,omitempty"`

	// Precisely one of SingularUint64Value must be set.
	SingularUint64Value *uint64 `json:"singular_uint64_value,omitempty"`

	// Precisely one of SingularInt32Value must be set.
	SingularInt32Value *int32 `json:"singular_int32_value,omitempty"`

	// Precisely one of SingularUint32Value must be set.
	SingularUint32Value *uint32 `json:"singular_uint32_value,omitempty"`

	// Precisely one of SingularBoolValue must be set.
	SingularBoolValue *bool `json:"singular_bool_value,omitempty"`

	// Precisely one of SingularStringValue must be set.
	SingularStringValue *string `json:"singular_string_value,omitempty"`

	// Precisely one of SingularBytesValue must be set.
	SingularBytesValue *[]byte `json:"singular_bytes_value,omitempty"`

	// Precisely one of SingularTimestamp must be set.
	SingularTimestamp *time.Time `json:"singular_timestamp,omitempty"`

	// Precisely one of SingularAny must be set.
	SingularAny *interface{} `json:"singular_any,omitempty"`

	// Precisely one of SingularStruct must be set.
	SingularStruct *scw.JSONObject `json:"singular_struct,omitempty"`

	// Precisely one of SingularMoney must be set.
	SingularMoney *scw.Money `json:"singular_money,omitempty"`

	// Precisely one of SingularStringsValue must be set.
	SingularStringsValue *[]string `json:"singular_strings_value,omitempty"`

	// Precisely one of SingularDuration must be set.
	SingularDuration *scw.Duration `json:"singular_duration,omitempty"`

	// Precisely one of MapStringString must be set.
	MapStringString *map[string]string `json:"map_string_string,omitempty"`

	// Precisely one of Timeseries must be set.
	Timeseries *scw.TimeSeries `json:"timeseries,omitempty"`
}

type PostAllOptionalMessageNestedMessage struct {
	S string `json:"s"`
}

type PostAllOptionalMessageNestedMessageWithOptional struct {

	// Precisely one of Bb must be set.
	Bb *int32 `json:"bb,omitempty"`
}

// PostAllTypesMessage: post all types message.
type PostAllTypesMessage struct {
	SingularInt32 int32 `json:"singular_int32"`

	SingularInt64 int64 `json:"singular_int64"`

	SingularUint32 uint32 `json:"singular_uint32"`

	SingularUint64 uint64 `json:"singular_uint64"`

	SingularSint32 int32 `json:"singular_sint32"`

	SingularSint64 int64 `json:"singular_sint64"`

	SingularFixed32 uint32 `json:"singular_fixed32"`

	SingularFixed64 uint64 `json:"singular_fixed64"`

	SingularSfixed32 int32 `json:"singular_sfixed32"`

	SingularSfixed64 int64 `json:"singular_sfixed64"`

	SingularFloat float32 `json:"singular_float"`

	SingularDouble float64 `json:"singular_double"`

	SingularBool bool `json:"singular_bool"`

	SingularString string `json:"singular_string"`

	SingularBytes []byte `json:"singular_bytes"`

	SingularNestedMessage *PostAllTypesMessageNestedMessage `json:"singular_nested_message"`
	// SingularNestedEnum: default value: ZERO
	SingularNestedEnum PostAllTypesMessageNestedEnum `json:"singular_nested_enum"`

	RepeatedInt32 []int32 `json:"repeated_int32"`

	RepeatedInt64 []int64 `json:"repeated_int64"`

	RepeatedUint32 []uint32 `json:"repeated_uint32"`

	RepeatedUint64 []uint64 `json:"repeated_uint64"`

	RepeatedSint32 []int32 `json:"repeated_sint32"`

	RepeatedSint64 []int64 `json:"repeated_sint64"`

	RepeatedFixed32 []uint32 `json:"repeated_fixed32"`

	RepeatedFixed64 []uint64 `json:"repeated_fixed64"`

	RepeatedSfixed32 []int32 `json:"repeated_sfixed32"`

	RepeatedSfixed64 []int64 `json:"repeated_sfixed64"`

	RepeatedFloat []float32 `json:"repeated_float"`

	RepeatedDouble []float64 `json:"repeated_double"`

	RepeatedBool []bool `json:"repeated_bool"`

	RepeatedString []string `json:"repeated_string"`

	RepeatedBytes [][]byte `json:"repeated_bytes"`

	RepeatedNestedMessage []*PostAllTypesMessageNestedMessage `json:"repeated_nested_message"`

	RepeatedNestedEnum []PostAllTypesMessageNestedEnum `json:"repeated_nested_enum"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofUint32 *uint32 `json:"oneof_uint32,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofNestedMessage *PostAllTypesMessageNestedMessage `json:"oneof_nested_message,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofString *string `json:"oneof_string,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofBytes *[]byte `json:"oneof_bytes,omitempty"`

	SingularDoubleValue *float64 `json:"singular_double_value"`

	SingularFloatValue *float32 `json:"singular_float_value"`

	SingularInt64Value *int64 `json:"singular_int64_value"`

	SingularUint64Value *uint64 `json:"singular_uint64_value"`

	SingularInt32Value *int32 `json:"singular_int32_value"`

	SingularUint32Value *uint32 `json:"singular_uint32_value"`

	SingularBoolValue *bool `json:"singular_bool_value"`

	SingularStringValue *string `json:"singular_string_value"`

	SingularBytesValue *[]byte `json:"singular_bytes_value"`

	SingularTimestamp *time.Time `json:"singular_timestamp"`

	SingularAny interface{} `json:"singular_any"`

	SingularStruct *scw.JSONObject `json:"singular_struct"`

	SingularMoney *scw.Money `json:"singular_money"`

	SingularStringsValue *[]string `json:"singular_strings_value"`

	SingularDuration *scw.Duration `json:"singular_duration"`

	SingularIP net.IP `json:"singular_ip"`

	SingularStringIP net.IP `json:"singular_string_ip"`

	SingularStringValueIP *net.IP `json:"singular_string_value_ip"`

	SingularIPv4 net.IP `json:"singular_ipv4"`

	SingularStringIPv4 net.IP `json:"singular_string_ipv4"`

	SingularStringValueIPv4 *net.IP `json:"singular_string_value_ipv4"`

	SingularIPv6 net.IP `json:"singular_ipv6"`

	SingularStringIPv6 net.IP `json:"singular_string_ipv6"`

	SingularStringValueIPv6 *net.IP `json:"singular_string_value_ipv6"`

	SingularStdDuration *time.Duration `json:"singular_std_duration"`

	SingularStdLongDuration *time.Duration `json:"singular_std_long_duration"`

	SingularSize scw.Size `json:"singular_size"`

	SingularUint64Size scw.Size `json:"singular_uint64_size"`

	SingularUint64valueSize *scw.Size `json:"singular_uint64value_size"`

	SingularStringIpnet scw.IPNet `json:"singular_string_ipnet"`

	SingularStringValueIpnet *scw.IPNet `json:"singular_string_value_ipnet"`

	RepeatedDoubleValue []*float64 `json:"repeated_double_value"`

	RepeatedFloatValue []*float32 `json:"repeated_float_value"`

	RepeatedInt64Value []*int64 `json:"repeated_int64_value"`

	RepeatedUint64Value []*uint64 `json:"repeated_uint64_value"`

	RepeatedInt32Value []*int32 `json:"repeated_int32_value"`

	RepeatedUint32Value []*uint32 `json:"repeated_uint32_value"`

	RepeatedBoolValue []*bool `json:"repeated_bool_value"`

	RepeatedStringValue []*string `json:"repeated_string_value"`

	RepeatedBytesValue []*[]byte `json:"repeated_bytes_value"`

	RepeatedTimestamp []*time.Time `json:"repeated_timestamp"`

	RepeatedAny []interface{} `json:"repeated_any"`

	RepeatedStruct []*scw.JSONObject `json:"repeated_struct"`

	RepeatedMoney []*scw.Money `json:"repeated_money"`

	RepeatedStringsValue []*[]string `json:"repeated_strings_value"`

	RepeatedDuration []*scw.Duration `json:"repeated_duration"`

	RepeatedIP []net.IP `json:"repeated_ip"`

	RepeatedStringIP []net.IP `json:"repeated_string_ip"`

	RepeatedStringValueIP []*net.IP `json:"repeated_string_value_ip"`

	RepeatedIPv4 []net.IP `json:"repeated_ipv4"`

	RepeatedStringIPv4 []net.IP `json:"repeated_string_ipv4"`

	RepeatedStringValueIPv4 []*net.IP `json:"repeated_string_value_ipv4"`

	RepeatedIPv6 []net.IP `json:"repeated_ipv6"`

	RepeatedStringIPv6 []net.IP `json:"repeated_string_ipv6"`

	RepeatedStringValueIPv6 []*net.IP `json:"repeated_string_value_ipv6"`

	RepeatedStdDuration []*time.Duration `json:"repeated_std_duration"`

	RepeatedStdLongDuration []*time.Duration `json:"repeated_std_long_duration"`

	RepeatedSize []scw.Size `json:"repeated_size"`

	RepeatedUint64Size []scw.Size `json:"repeated_uint64_size"`

	RepeatedUint64valueSize []*scw.Size `json:"repeated_uint64value_size"`

	RepeatedStringIpnet []scw.IPNet `json:"repeated_string_ipnet"`

	RepeatedStringValueIpnet []*scw.IPNet `json:"repeated_string_value_ipnet"`
}

func (m *PostAllTypesMessage) UnmarshalJSON(b []byte) error {
	type tmpType PostAllTypesMessage
	tmp := struct {
		tmpType

		TmpSingularStdDuration     *marshaler.Duration         `json:"singular_std_duration"`
		TmpSingularStdLongDuration *marshaler.LongDuration     `json:"singular_std_long_duration"`
		TmpRepeatedStdDuration     marshaler.DurationSlice     `json:"repeated_std_duration"`
		TmpRepeatedStdLongDuration marshaler.LongDurationSlice `json:"repeated_std_long_duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostAllTypesMessage(tmp.tmpType)

	m.SingularStdDuration = tmp.TmpSingularStdDuration.Standard()
	m.SingularStdLongDuration = tmp.TmpSingularStdLongDuration.Standard()
	m.RepeatedStdDuration = tmp.TmpRepeatedStdDuration.Standard()
	m.RepeatedStdLongDuration = tmp.TmpRepeatedStdLongDuration.Standard()
	return nil
}

func (m PostAllTypesMessage) MarshalJSON() ([]byte, error) {
	type tmpType PostAllTypesMessage
	tmp := struct {
		tmpType

		TmpSingularStdDuration     *marshaler.Duration         `json:"singular_std_duration"`
		TmpSingularStdLongDuration *marshaler.LongDuration     `json:"singular_std_long_duration"`
		TmpRepeatedStdDuration     marshaler.DurationSlice     `json:"repeated_std_duration"`
		TmpRepeatedStdLongDuration marshaler.LongDurationSlice `json:"repeated_std_long_duration"`
	}{
		tmpType: tmpType(m),

		TmpSingularStdDuration:     marshaler.NewDuration(m.SingularStdDuration),
		TmpSingularStdLongDuration: marshaler.NewLongDuration(m.SingularStdLongDuration),
		TmpRepeatedStdDuration:     marshaler.NewDurationSlice(m.RepeatedStdDuration),
		TmpRepeatedStdLongDuration: marshaler.NewLongDurationSlice(m.RepeatedStdLongDuration),
	}
	return json.Marshal(tmp)
}

type PostAllTypesMessageNestedMessage struct {
	Bb int32 `json:"bb"`
}

type PostBodyAndPathAndQueryMessage struct {
	Path string `json:"path"`

	Query string `json:"query"`

	Body *PostBodyAndPathSimpleMessage `json:"body"`
}

type PostBodyAndPathComplexMessage struct {
	Path string `json:"path"`

	Body *PostBodyAndPathSimpleMessage `json:"body"`
}

type PostBodyAndPathSimple2Message struct {
	Path string `json:"path"`

	Body string `json:"body"`
}

// PostBodyAndPathSimpleMessage: post body and path simple message.
type PostBodyAndPathSimpleMessage struct {
	Path string `json:"path"`

	Body string `json:"body"`
}

type PostComplexValidateMessage struct {
	Val *ComplexValidateMsg `json:"val"`
}

// PostDeprecatedOrganizationMessage: post deprecated organization message.
type PostDeprecatedOrganizationMessage struct {
	// Deprecated
	Organization *string `json:"organization,omitempty"`
}

type PostDeprecatedProjectMessage struct {
	// Deprecated
	Project *string `json:"project,omitempty"`
}

// PostEchoTimeSeriesMessage: post echo time series message.
type PostEchoTimeSeriesMessage struct {
	Metrics *scw.TimeSeries `json:"metrics"`
}

// PostEnumMessage: post enum message.
type PostEnumMessage struct {
	// Type: default value: unknown
	Type Type `json:"type"`
	// Type2: default value: unknown
	Type2 *Type `json:"type2"`
	// Type3: default value: unknown
	Type3 *Type `json:"type3"`
}

// PostIPMessage: post ip message.
type PostIPMessage struct {
	IPV4 net.IP `json:"ip_v4"`

	IPV6 net.IP `json:"ip_v6"`

	IP net.IP `json:"ip"`
}

type PostOneOfMessage struct {

	// Precisely one of Test, Test2 must be set.
	Test *string `json:"test,omitempty"`

	// Precisely one of Test, Test2 must be set.
	Test2 *int32 `json:"test2,omitempty"`

	// Precisely one of Test3, TestNested must be set.
	TestNested *EchoMessage `json:"test_nested,omitempty"`

	// Precisely one of Test3, TestNested must be set.
	Test3 *int32 `json:"test3,omitempty"`
}

// PostOrganizationIDMessage: post organization id message.
type PostOrganizationIDMessage struct {
	OrganizationID string `json:"organization_id"`
}

type PostProjectIDMessage struct {
	ProjectID string `json:"project_id"`
}

// PostScalarTypesMessage: post scalar types message.
type PostScalarTypesMessage struct {
	DoubleField float64 `json:"double_field"`

	FloatField float32 `json:"float_field"`

	Int32Field int32 `json:"int32_field"`

	Int64Field int64 `json:"int64_field"`

	Uint32Field uint32 `json:"uint32_field"`

	Uint64Field uint64 `json:"uint64_field"`

	BoolField bool `json:"bool_field"`

	StringField string `json:"string_field"`
}

// PostTagsMessage: post tags message.
type PostTagsMessage struct {
	Tags *[]string `json:"tags"`
}

// Transient: transient.
type Transient struct {
	// Status: default value: unknown_status
	Status TransientStatus `json:"status"`
}

// getRegionResponse: get region response.
type getRegionResponse struct {
	Region scw.Region `json:"region"`
}

// Service API

type GetServiceInfoRequest struct {
}

func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1",
		Headers: http.Header{},
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteNothingRequest struct {
}

// DeleteNothing: this method deletes nothing.
func (s *API) DeleteNothing(req *DeleteNothingRequest, opts ...scw.RequestOption) error {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/test-internal/v1",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetEchoRequest struct {
	// Str: a string.
	Str string `json:"-"`
	// Strs: a slice of strings.
	Strs []string `json:"-"`
}

// GetEcho: echo the request message.
// ### This is a multiline test.
func (s *API) GetEcho(req *GetEchoRequest, opts ...scw.RequestOption) (*EchoMessage, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "str", req.Str)
	parameter.AddToQuery(query, "strs", req.Strs)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/echo",
		Query:   query,
		Headers: http.Header{},
	}

	var resp EchoMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostEchoRequest struct {
	Str *string `json:"str"`

	Strs []string `json:"strs"`
}

// postEcho: echo the request message.
func (s *API) postEcho(req *PostEchoRequest, opts ...scw.RequestOption) (*EchoMessage, error) {
	var err error

	if req.Str == nil {
		req.Str = scw.StringPtr(namegenerator.GetRandomName("name"))
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/echo",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EchoMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetEnumRequest struct {
	// Type: default value: unknown
	Type Type `json:"-"`
}

// GetEnum: get enum.
func (s *API) GetEnum(req *GetEnumRequest, opts ...scw.RequestOption) (*GetEnumMessage, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "type", req.Type)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/enum",
		Query:   query,
		Headers: http.Header{},
	}

	var resp GetEnumMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostEnumRequest struct {
	// Type: default value: unknown
	Type Type `json:"type"`
	// Type2: default value: unknown
	Type2 *Type `json:"type2"`
	// Type3: default value: unknown
	Type3 *Type `json:"type3"`
}

// PostEnum: post enum.
func (s *API) PostEnum(req *PostEnumRequest, opts ...scw.RequestOption) (*PostEnumMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/enum",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostEnumMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PatchEnumRequest struct {
	// Type: default value: unknown
	Type Type `json:"type"`
}

// PatchEnum: patch enum.
func (s *API) PatchEnum(req *PatchEnumRequest, opts ...scw.RequestOption) (*PatchEnumMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/test-internal/v1/enum",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PatchEnumMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostTagsRequest struct {
	Tags *[]string `json:"tags"`
}

// PostTags: post tags.
func (s *API) PostTags(req *PostTagsRequest, opts ...scw.RequestOption) (*PostTagsMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/test-internal/v1/tags",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostTagsMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostFileRequest struct {
	Body *scw.File
}

// PostFile: post file.
func (s *API) PostFile(req *PostFileRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/file",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req.Body)
	if err != nil {
		return nil, err
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostIPRequest struct {
	IPV4 net.IP `json:"ip_v4"`

	IPV6 net.IP `json:"ip_v6"`

	IP net.IP `json:"ip"`
}

// PostIP: post IP.
func (s *API) PostIP(req *PostIPRequest, opts ...scw.RequestOption) (*PostIPMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/ip",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostIPMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostScalarTypesRequest struct {
	DoubleField float64 `json:"double_field"`

	FloatField float32 `json:"float_field"`

	Int32Field int32 `json:"int32_field"`

	Int64Field int64 `json:"int64_field"`

	Uint32Field uint32 `json:"uint32_field"`

	Uint64Field uint64 `json:"uint64_field"`

	BoolField bool `json:"bool_field"`

	StringField string `json:"string_field"`
}

// PostScalarTypes: post scalar types.
func (s *API) PostScalarTypes(req *PostScalarTypesRequest, opts ...scw.RequestOption) (*PostScalarTypesMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/scalar-types",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostScalarTypesMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetMetadataRequest struct {
}

func (s *API) GetMetadata(req *GetMetadataRequest, opts ...scw.RequestOption) (*MetadataResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/metadata",
		Headers: http.Header{},
	}

	var resp MetadataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostWaitRequest struct {
	// Duration: waiting duration.
	Duration *time.Duration `json:"duration"`
}

func (m *PostWaitRequest) UnmarshalJSON(b []byte) error {
	type tmpType PostWaitRequest
	tmp := struct {
		tmpType

		TmpDuration *marshaler.Duration `json:"duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostWaitRequest(tmp.tmpType)

	m.Duration = tmp.TmpDuration.Standard()
	return nil
}

func (m PostWaitRequest) MarshalJSON() ([]byte, error) {
	type tmpType PostWaitRequest
	tmp := struct {
		tmpType

		TmpDuration *marshaler.Duration `json:"duration"`
	}{
		tmpType: tmpType(m),

		TmpDuration: marshaler.NewDuration(m.Duration),
	}
	return json.Marshal(tmp)
}

// PostWait: wait until a given time in second.
func (s *API) PostWait(req *PostWaitRequest, opts ...scw.RequestOption) error {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/wait",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetRegionRequest struct {
}

// GetRegion: get a region.
func (s *API) GetRegion(req *GetRegionRequest, opts ...scw.RequestOption) (*getRegionResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/region",
		Headers: http.Header{},
	}

	var resp getRegionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostOrganizationIDRequest struct {
	OrganizationID string `json:"organization_id"`
}

// PostOrganizationID: post an organization ID.
func (s *API) PostOrganizationID(req *PostOrganizationIDRequest, opts ...scw.RequestOption) (*PostOrganizationIDMessage, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/organization-id",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostOrganizationIDMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostDeprecatedOrganizationRequest struct {
	// Deprecated
	Organization *string `json:"organization,omitempty"`
}

// Deprecated: PostDeprecatedOrganization: post a deprecated organization ID.
func (s *API) PostDeprecatedOrganization(req *PostDeprecatedOrganizationRequest, opts ...scw.RequestOption) (*PostDeprecatedOrganizationMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/deprecated-organization",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostDeprecatedOrganizationMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostProjectIDRequest struct {
	ProjectID string `json:"project_id"`
}

func (s *API) PostProjectID(req *PostProjectIDRequest, opts ...scw.RequestOption) (*PostProjectIDMessage, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/project-id",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostProjectIDMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostDeprecatedProjectRequest struct {
	// Deprecated
	Project *string `json:"project,omitempty"`
}

// Deprecated
func (s *API) PostDeprecatedProject(req *PostDeprecatedProjectRequest, opts ...scw.RequestOption) (*PostDeprecatedProjectMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/deprecated-project",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostDeprecatedProjectMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostOneOfRequest struct {

	// Precisely one of Test, Test2 must be set.
	Test *string `json:"test,omitempty"`

	// Precisely one of Test, Test2 must be set.
	Test2 *int32 `json:"test2,omitempty"`

	// Precisely one of Test3, TestNested must be set.
	TestNested *EchoMessage `json:"test_nested,omitempty"`

	// Precisely one of Test3, TestNested must be set.
	Test3 *int32 `json:"test3,omitempty"`
}

func (s *API) PostOneOf(req *PostOneOfRequest, opts ...scw.RequestOption) (*PostOneOfMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/oneof",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostOneOfMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostBodyAndPathSimpleRequest struct {
	Path string `json:"-"`

	Body string `json:"body"`
}

func (s *API) PostBodyAndPathSimple(req *PostBodyAndPathSimpleRequest, opts ...scw.RequestOption) (*PostBodyAndPathSimpleMessage, error) {
	var err error

	if fmt.Sprint(req.Path) == "" {
		return nil, errors.New("field Path cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/paths/" + fmt.Sprint(req.Path) + "/simple",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req.Body)
	if err != nil {
		return nil, err
	}

	var resp PostBodyAndPathSimpleMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostBodyAndPathSimple2Request struct {
	Path string `json:"-"`

	Body string `json:"body"`
}

func (s *API) PostBodyAndPathSimple2(req *PostBodyAndPathSimple2Request, opts ...scw.RequestOption) (*PostBodyAndPathSimple2Message, error) {
	var err error

	if fmt.Sprint(req.Path) == "" {
		return nil, errors.New("field Path cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/paths/" + fmt.Sprint(req.Path) + "/simple2",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostBodyAndPathSimple2Message

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostBodyAndPathComplexRequest struct {
	Path string `json:"-"`

	Body *PostBodyAndPathSimpleMessage `json:"body"`
}

func (s *API) PostBodyAndPathComplex(req *PostBodyAndPathComplexRequest, opts ...scw.RequestOption) (*PostBodyAndPathComplexMessage, error) {
	var err error

	if fmt.Sprint(req.Path) == "" {
		return nil, errors.New("field Path cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/paths/" + fmt.Sprint(req.Path) + "/complex",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req.Body)
	if err != nil {
		return nil, err
	}

	var resp PostBodyAndPathComplexMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostBodyAndPathAndQueryRequest struct {
	Path string `json:"-"`

	Query string `json:"-"`

	Body *PostBodyAndPathSimpleMessage `json:"body"`
}

func (s *API) PostBodyAndPathAndQuery(req *PostBodyAndPathAndQueryRequest, opts ...scw.RequestOption) (*PostBodyAndPathAndQueryMessage, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "query", req.Query)

	if fmt.Sprint(req.Path) == "" {
		return nil, errors.New("field Path cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/paths/" + fmt.Sprint(req.Path) + "/bodyquery",
		Query:   query,
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req.Body)
	if err != nil {
		return nil, err
	}

	var resp PostBodyAndPathAndQueryMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostAllMapTypesRequest struct {
	MapInt32Int32 map[int32]int32 `json:"map_int32_int32"`

	MapInt64Int64 map[int64]int64 `json:"map_int64_int64"`

	MapUint32Uint32 map[uint32]uint32 `json:"map_uint32_uint32"`

	MapUint64Uint64 map[uint64]uint64 `json:"map_uint64_uint64"`

	MapSint32Sint32 map[int32]int32 `json:"map_sint32_sint32"`

	MapSint64Sint64 map[int64]int64 `json:"map_sint64_sint64"`

	MapFixed32Fixed32 map[uint32]uint32 `json:"map_fixed32_fixed32"`

	MapFixed64Fixed64 map[uint64]uint64 `json:"map_fixed64_fixed64"`

	MapSfixed32Sfixed32 map[int32]int32 `json:"map_sfixed32_sfixed32"`

	MapSfixed64Sfixed64 map[int64]int64 `json:"map_sfixed64_sfixed64"`

	MapInt32Float map[int32]float32 `json:"map_int32_float"`

	MapInt32Double map[int32]float64 `json:"map_int32_double"`

	MapStringString map[string]string `json:"map_string_string"`

	MapInt32Bytes map[int32][]byte `json:"map_int32_bytes"`

	MapInt32Enum map[int32]MapEnum `json:"map_int32_enum"`

	MapInt32AllTypes map[int32]*PostAllTypesMessage `json:"map_int32_all_types"`

	MapInt32IP map[int32]net.IP `json:"map_int32_ip"`

	MapInt32StdDuration map[int32]*time.Duration `json:"map_int32_std_duration"`

	MapInt32StdLongDuration map[int32]*time.Duration `json:"map_int32_std_long_duration"`

	MapInt32Size map[int32]scw.Size `json:"map_int32_size"`

	MapInt32Uint64Size map[int32]scw.Size `json:"map_int32_uint64_size"`

	MapInt32Uint64valueSize map[int32]*scw.Size `json:"map_int32_uint64value_size"`

	MapInt32StringIP map[int32]net.IP `json:"map_int32_string_ip"`

	MapInt32StringValueIP map[int32]*net.IP `json:"map_int32_string_value_ip"`

	MapInt32IPv4 map[int32]net.IP `json:"map_int32_ipv4"`

	MapInt32StringIPv4 map[int32]net.IP `json:"map_int32_string_ipv4"`

	MapInt32StringValueIPv4 map[int32]*net.IP `json:"map_int32_string_value_ipv4"`

	MapInt32IPv6 map[int32]net.IP `json:"map_int32_ipv6"`

	MapInt32StringIPv6 map[int32]net.IP `json:"map_int32_string_ipv6"`

	MapInt32StringValueIPv6 map[int32]*net.IP `json:"map_int32_string_value_ipv6"`

	MapInt32StringsValue map[int32]*[]string `json:"map_int32_strings_value"`

	MapInt32Duration map[int32]*scw.Duration `json:"map_int32_duration"`
}

func (m *PostAllMapTypesRequest) UnmarshalJSON(b []byte) error {
	type tmpType PostAllMapTypesRequest
	tmp := struct {
		tmpType

		TmpMapInt32StdDuration     marshaler.Durationint32Map     `json:"map_int32_std_duration"`
		TmpMapInt32StdLongDuration marshaler.LongDurationint32Map `json:"map_int32_std_long_duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostAllMapTypesRequest(tmp.tmpType)

	m.MapInt32StdDuration = tmp.TmpMapInt32StdDuration.Standard()
	m.MapInt32StdLongDuration = tmp.TmpMapInt32StdLongDuration.Standard()
	return nil
}

func (m PostAllMapTypesRequest) MarshalJSON() ([]byte, error) {
	type tmpType PostAllMapTypesRequest
	tmp := struct {
		tmpType

		TmpMapInt32StdDuration     marshaler.Durationint32Map     `json:"map_int32_std_duration"`
		TmpMapInt32StdLongDuration marshaler.LongDurationint32Map `json:"map_int32_std_long_duration"`
	}{
		tmpType: tmpType(m),

		TmpMapInt32StdDuration:     marshaler.NewDurationint32Map(m.MapInt32StdDuration),
		TmpMapInt32StdLongDuration: marshaler.NewLongDurationint32Map(m.MapInt32StdLongDuration),
	}
	return json.Marshal(tmp)
}

func (s *API) PostAllMapTypes(req *PostAllMapTypesRequest, opts ...scw.RequestOption) (*PostAllMapTypesMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/map",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostAllMapTypesMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostAllTypesRequest struct {
	SingularInt32 int32 `json:"singular_int32"`

	SingularInt64 int64 `json:"singular_int64"`

	SingularUint32 uint32 `json:"singular_uint32"`

	SingularUint64 uint64 `json:"singular_uint64"`

	SingularSint32 int32 `json:"singular_sint32"`

	SingularSint64 int64 `json:"singular_sint64"`

	SingularFixed32 uint32 `json:"singular_fixed32"`

	SingularFixed64 uint64 `json:"singular_fixed64"`

	SingularSfixed32 int32 `json:"singular_sfixed32"`

	SingularSfixed64 int64 `json:"singular_sfixed64"`

	SingularFloat float32 `json:"singular_float"`

	SingularDouble float64 `json:"singular_double"`

	SingularBool bool `json:"singular_bool"`

	SingularString string `json:"singular_string"`

	SingularBytes []byte `json:"singular_bytes"`

	SingularNestedMessage *PostAllTypesMessageNestedMessage `json:"singular_nested_message"`
	// SingularNestedEnum: default value: ZERO
	SingularNestedEnum PostAllTypesMessageNestedEnum `json:"singular_nested_enum"`

	RepeatedInt32 []int32 `json:"repeated_int32"`

	RepeatedInt64 []int64 `json:"repeated_int64"`

	RepeatedUint32 []uint32 `json:"repeated_uint32"`

	RepeatedUint64 []uint64 `json:"repeated_uint64"`

	RepeatedSint32 []int32 `json:"repeated_sint32"`

	RepeatedSint64 []int64 `json:"repeated_sint64"`

	RepeatedFixed32 []uint32 `json:"repeated_fixed32"`

	RepeatedFixed64 []uint64 `json:"repeated_fixed64"`

	RepeatedSfixed32 []int32 `json:"repeated_sfixed32"`

	RepeatedSfixed64 []int64 `json:"repeated_sfixed64"`

	RepeatedFloat []float32 `json:"repeated_float"`

	RepeatedDouble []float64 `json:"repeated_double"`

	RepeatedBool []bool `json:"repeated_bool"`

	RepeatedString []string `json:"repeated_string"`

	RepeatedBytes [][]byte `json:"repeated_bytes"`

	RepeatedNestedMessage []*PostAllTypesMessageNestedMessage `json:"repeated_nested_message"`

	RepeatedNestedEnum []PostAllTypesMessageNestedEnum `json:"repeated_nested_enum"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofUint32 *uint32 `json:"oneof_uint32,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofNestedMessage *PostAllTypesMessageNestedMessage `json:"oneof_nested_message,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofString *string `json:"oneof_string,omitempty"`

	// Precisely one of OneofBytes, OneofNestedMessage, OneofString, OneofUint32 must be set.
	OneofBytes *[]byte `json:"oneof_bytes,omitempty"`

	SingularDoubleValue *float64 `json:"singular_double_value"`

	SingularFloatValue *float32 `json:"singular_float_value"`

	SingularInt64Value *int64 `json:"singular_int64_value"`

	SingularUint64Value *uint64 `json:"singular_uint64_value"`

	SingularInt32Value *int32 `json:"singular_int32_value"`

	SingularUint32Value *uint32 `json:"singular_uint32_value"`

	SingularBoolValue *bool `json:"singular_bool_value"`

	SingularStringValue *string `json:"singular_string_value"`

	SingularBytesValue *[]byte `json:"singular_bytes_value"`

	SingularTimestamp *time.Time `json:"singular_timestamp"`

	SingularAny interface{} `json:"singular_any"`

	SingularStruct *scw.JSONObject `json:"singular_struct"`

	SingularMoney *scw.Money `json:"singular_money"`

	SingularStringsValue *[]string `json:"singular_strings_value"`

	SingularDuration *scw.Duration `json:"singular_duration"`

	SingularIP net.IP `json:"singular_ip"`

	SingularStringIP net.IP `json:"singular_string_ip"`

	SingularStringValueIP *net.IP `json:"singular_string_value_ip"`

	SingularIPv4 net.IP `json:"singular_ipv4"`

	SingularStringIPv4 net.IP `json:"singular_string_ipv4"`

	SingularStringValueIPv4 *net.IP `json:"singular_string_value_ipv4"`

	SingularIPv6 net.IP `json:"singular_ipv6"`

	SingularStringIPv6 net.IP `json:"singular_string_ipv6"`

	SingularStringValueIPv6 *net.IP `json:"singular_string_value_ipv6"`

	SingularStdDuration *time.Duration `json:"singular_std_duration"`

	SingularStdLongDuration *time.Duration `json:"singular_std_long_duration"`

	SingularSize scw.Size `json:"singular_size"`

	SingularUint64Size scw.Size `json:"singular_uint64_size"`

	SingularUint64valueSize *scw.Size `json:"singular_uint64value_size"`

	SingularStringIpnet scw.IPNet `json:"singular_string_ipnet"`

	SingularStringValueIpnet *scw.IPNet `json:"singular_string_value_ipnet"`

	RepeatedDoubleValue []*float64 `json:"repeated_double_value"`

	RepeatedFloatValue []*float32 `json:"repeated_float_value"`

	RepeatedInt64Value []*int64 `json:"repeated_int64_value"`

	RepeatedUint64Value []*uint64 `json:"repeated_uint64_value"`

	RepeatedInt32Value []*int32 `json:"repeated_int32_value"`

	RepeatedUint32Value []*uint32 `json:"repeated_uint32_value"`

	RepeatedBoolValue []*bool `json:"repeated_bool_value"`

	RepeatedStringValue []*string `json:"repeated_string_value"`

	RepeatedBytesValue []*[]byte `json:"repeated_bytes_value"`

	RepeatedTimestamp []*time.Time `json:"repeated_timestamp"`

	RepeatedAny []interface{} `json:"repeated_any"`

	RepeatedStruct []*scw.JSONObject `json:"repeated_struct"`

	RepeatedMoney []*scw.Money `json:"repeated_money"`

	RepeatedStringsValue []*[]string `json:"repeated_strings_value"`

	RepeatedDuration []*scw.Duration `json:"repeated_duration"`

	RepeatedIP []net.IP `json:"repeated_ip"`

	RepeatedStringIP []net.IP `json:"repeated_string_ip"`

	RepeatedStringValueIP []*net.IP `json:"repeated_string_value_ip"`

	RepeatedIPv4 []net.IP `json:"repeated_ipv4"`

	RepeatedStringIPv4 []net.IP `json:"repeated_string_ipv4"`

	RepeatedStringValueIPv4 []*net.IP `json:"repeated_string_value_ipv4"`

	RepeatedIPv6 []net.IP `json:"repeated_ipv6"`

	RepeatedStringIPv6 []net.IP `json:"repeated_string_ipv6"`

	RepeatedStringValueIPv6 []*net.IP `json:"repeated_string_value_ipv6"`

	RepeatedStdDuration []*time.Duration `json:"repeated_std_duration"`

	RepeatedStdLongDuration []*time.Duration `json:"repeated_std_long_duration"`

	RepeatedSize []scw.Size `json:"repeated_size"`

	RepeatedUint64Size []scw.Size `json:"repeated_uint64_size"`

	RepeatedUint64valueSize []*scw.Size `json:"repeated_uint64value_size"`

	RepeatedStringIpnet []scw.IPNet `json:"repeated_string_ipnet"`

	RepeatedStringValueIpnet []*scw.IPNet `json:"repeated_string_value_ipnet"`
}

func (m *PostAllTypesRequest) UnmarshalJSON(b []byte) error {
	type tmpType PostAllTypesRequest
	tmp := struct {
		tmpType

		TmpSingularStdDuration     *marshaler.Duration         `json:"singular_std_duration"`
		TmpSingularStdLongDuration *marshaler.LongDuration     `json:"singular_std_long_duration"`
		TmpRepeatedStdDuration     marshaler.DurationSlice     `json:"repeated_std_duration"`
		TmpRepeatedStdLongDuration marshaler.LongDurationSlice `json:"repeated_std_long_duration"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = PostAllTypesRequest(tmp.tmpType)

	m.SingularStdDuration = tmp.TmpSingularStdDuration.Standard()
	m.SingularStdLongDuration = tmp.TmpSingularStdLongDuration.Standard()
	m.RepeatedStdDuration = tmp.TmpRepeatedStdDuration.Standard()
	m.RepeatedStdLongDuration = tmp.TmpRepeatedStdLongDuration.Standard()
	return nil
}

func (m PostAllTypesRequest) MarshalJSON() ([]byte, error) {
	type tmpType PostAllTypesRequest
	tmp := struct {
		tmpType

		TmpSingularStdDuration     *marshaler.Duration         `json:"singular_std_duration"`
		TmpSingularStdLongDuration *marshaler.LongDuration     `json:"singular_std_long_duration"`
		TmpRepeatedStdDuration     marshaler.DurationSlice     `json:"repeated_std_duration"`
		TmpRepeatedStdLongDuration marshaler.LongDurationSlice `json:"repeated_std_long_duration"`
	}{
		tmpType: tmpType(m),

		TmpSingularStdDuration:     marshaler.NewDuration(m.SingularStdDuration),
		TmpSingularStdLongDuration: marshaler.NewLongDuration(m.SingularStdLongDuration),
		TmpRepeatedStdDuration:     marshaler.NewDurationSlice(m.RepeatedStdDuration),
		TmpRepeatedStdLongDuration: marshaler.NewLongDurationSlice(m.RepeatedStdLongDuration),
	}
	return json.Marshal(tmp)
}

// PostAllTypes: post all types.
func (s *API) PostAllTypes(req *PostAllTypesRequest, opts ...scw.RequestOption) (*PostAllTypesMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/alltypes",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostAllTypesMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostAllOptionalRequest struct {

	// Precisely one of OptionalInt32 must be set.
	OptionalInt32 *int32 `json:"optional_int32,omitempty"`

	// Precisely one of OptionalInt64 must be set.
	OptionalInt64 *int64 `json:"optional_int64,omitempty"`

	// Precisely one of OptionalUint32 must be set.
	OptionalUint32 *uint32 `json:"optional_uint32,omitempty"`

	// Precisely one of OptionalUint64 must be set.
	OptionalUint64 *uint64 `json:"optional_uint64,omitempty"`

	// Precisely one of OptionalSint32 must be set.
	OptionalSint32 *int32 `json:"optional_sint32,omitempty"`

	// Precisely one of OptionalSint64 must be set.
	OptionalSint64 *int64 `json:"optional_sint64,omitempty"`

	// Precisely one of OptionalFixed32 must be set.
	OptionalFixed32 *uint32 `json:"optional_fixed32,omitempty"`

	// Precisely one of OptionalFixed64 must be set.
	OptionalFixed64 *uint64 `json:"optional_fixed64,omitempty"`

	// Precisely one of OptionalSfixed32 must be set.
	OptionalSfixed32 *int32 `json:"optional_sfixed32,omitempty"`

	// Precisely one of OptionalSfixed64 must be set.
	OptionalSfixed64 *int64 `json:"optional_sfixed64,omitempty"`

	// Precisely one of OptionalFloat must be set.
	OptionalFloat *float32 `json:"optional_float,omitempty"`

	// Precisely one of OptionalDouble must be set.
	OptionalDouble *float64 `json:"optional_double,omitempty"`

	// Precisely one of OptionalBool must be set.
	OptionalBool *bool `json:"optional_bool,omitempty"`

	// Precisely one of OptionalString must be set.
	OptionalString *string `json:"optional_string,omitempty"`

	// Precisely one of OptionalBytes must be set.
	OptionalBytes *[]byte `json:"optional_bytes,omitempty"`

	// Precisely one of OptionalCord must be set.
	OptionalCord *string `json:"optional_cord,omitempty"`

	// Precisely one of OptionalNestedMessageWithOptional must be set.
	OptionalNestedMessageWithOptional *PostAllOptionalMessageNestedMessageWithOptional `json:"optional_nested_message_with_optional,omitempty"`

	// Precisely one of LazyNestedMessageWithOptional must be set.
	LazyNestedMessageWithOptional *PostAllOptionalMessageNestedMessageWithOptional `json:"lazy_nested_message_with_optional,omitempty"`
	// OptionalNestedEnum: default value: UNSPECIFIED
	// Precisely one of OptionalNestedEnum must be set.
	OptionalNestedEnum *PostAllOptionalMessageNestedEnum `json:"optional_nested_enum,omitempty"`

	// Precisely one of OptionalNestedMessage must be set.
	OptionalNestedMessage *PostAllOptionalMessageNestedMessage `json:"optional_nested_message,omitempty"`

	SingularInt32 int32 `json:"singular_int32"`

	SingularInt64 int64 `json:"singular_int64"`

	NestedMessage *PostAllOptionalMessageNestedMessage `json:"nested_message"`

	// Precisely one of SingularDoubleValue must be set.
	SingularDoubleValue *float64 `json:"singular_double_value,omitempty"`

	// Precisely one of SingularFloatValue must be set.
	SingularFloatValue *float32 `json:"singular_float_value,omitempty"`

	// Precisely one of SingularInt64Value must be set.
	SingularInt64Value *int64 `json:"singular_int64_value,omitempty"`

	// Precisely one of SingularUint64Value must be set.
	SingularUint64Value *uint64 `json:"singular_uint64_value,omitempty"`

	// Precisely one of SingularInt32Value must be set.
	SingularInt32Value *int32 `json:"singular_int32_value,omitempty"`

	// Precisely one of SingularUint32Value must be set.
	SingularUint32Value *uint32 `json:"singular_uint32_value,omitempty"`

	// Precisely one of SingularBoolValue must be set.
	SingularBoolValue *bool `json:"singular_bool_value,omitempty"`

	// Precisely one of SingularStringValue must be set.
	SingularStringValue *string `json:"singular_string_value,omitempty"`

	// Precisely one of SingularBytesValue must be set.
	SingularBytesValue *[]byte `json:"singular_bytes_value,omitempty"`

	// Precisely one of SingularTimestamp must be set.
	SingularTimestamp *time.Time `json:"singular_timestamp,omitempty"`

	// Precisely one of SingularAny must be set.
	SingularAny *interface{} `json:"singular_any,omitempty"`

	// Precisely one of SingularStruct must be set.
	SingularStruct *scw.JSONObject `json:"singular_struct,omitempty"`

	// Precisely one of SingularMoney must be set.
	SingularMoney *scw.Money `json:"singular_money,omitempty"`

	// Precisely one of SingularStringsValue must be set.
	SingularStringsValue *[]string `json:"singular_strings_value,omitempty"`

	// Precisely one of SingularDuration must be set.
	SingularDuration *scw.Duration `json:"singular_duration,omitempty"`

	// Precisely one of MapStringString must be set.
	MapStringString *map[string]string `json:"map_string_string,omitempty"`

	// Precisely one of Timeseries must be set.
	Timeseries *scw.TimeSeries `json:"timeseries,omitempty"`
}

func (s *API) PostAllOptional(req *PostAllOptionalRequest, opts ...scw.RequestOption) (*PostAllOptionalMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/alloptional",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostAllOptionalMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostComplexValidateRequest struct {
	Val *ComplexValidateMsg `json:"val"`
}

func (s *API) PostComplexValidate(req *PostComplexValidateRequest, opts ...scw.RequestOption) (*PostComplexValidateMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/complex-validate",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostComplexValidateMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListCharactersRequest struct {
	// Page: a positive integer to choose the page to display.
	Page *int32 `json:"-"`
	// PageSize: a positive integer lower or equal to 100 to select the number of items to display.
	// Default value: 10
	PageSize *uint32 `json:"-"`
	// OrderBy: order the listing.
	// Default value: asc
	OrderBy ListCharactersRequestOrderBy `json:"-"`
	// Name: filter characters by name, this is a "contains" matching.
	// Default value: i-am-fake
	Name *string `json:"-"`
	// Tags: dummy tags to check the comma_separated_list argument.
	Tags []string `json:"-"`
}

// ListCharacters: list The Lord of the Rings characters.
func (s *API) ListCharacters(req *ListCharactersRequest, opts ...scw.RequestOption) (*ListCharactersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)
	if len(req.Tags) != 0 {
		parameter.AddToQuery(query, "tags", strings.Join(req.Tags, ","))
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/characters",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListCharactersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostEchoTimeSeriesRequest struct {
	Metrics *scw.TimeSeries `json:"metrics"`
}

// PostEchoTimeSeries: echo metrics.
func (s *API) PostEchoTimeSeries(req *PostEchoTimeSeriesRequest, opts ...scw.RequestOption) (*PostEchoTimeSeriesMessage, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/echo-timeseries",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PostEchoTimeSeriesMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PostTransientRequest struct {
}

func (s *API) PostTransient(req *PostTransientRequest, opts ...scw.RequestOption) (*Transient, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/transients",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Transient

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetTransientRequest struct {
	TransientID string `json:"-"`
}

func (s *API) GetTransient(req *GetTransientRequest, opts ...scw.RequestOption) (*Transient, error) {
	var err error

	if fmt.Sprint(req.TransientID) == "" {
		return nil, errors.New("field TransientID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/transients/" + fmt.Sprint(req.TransientID) + "",
		Headers: http.Header{},
	}

	var resp Transient

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Service NoAuthAPI

// Service RegionalAPI

// Regions list localities the api is available in
func (s *RegionalAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms}
}

type RegionalAPIGetServiceInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

func (s *RegionalAPI) GetServiceInfo(req *RegionalAPIGetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "",
		Headers: http.Header{},
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RegionalAPIGetRegionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

func (s *RegionalAPI) GetRegion(req *RegionalAPIGetRegionRequest, opts ...scw.RequestOption) (*getRegionResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "/region",
		Headers: http.Header{},
	}

	var resp getRegionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RegionalAPIGetMetadataRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

func (s *RegionalAPI) GetMetadata(req *RegionalAPIGetMetadataRequest, opts ...scw.RequestOption) (*MetadataResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "/metadata",
		Headers: http.Header{},
	}

	var resp MetadataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RegionalAPIPostEchoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Str *string `json:"str"`

	Strs []string `json:"strs"`
}

func (s *RegionalAPI) PostEcho(req *RegionalAPIPostEchoRequest, opts ...scw.RequestOption) (*EchoMessage, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Str == nil {
		req.Str = scw.StringPtr(namegenerator.GetRandomName("name"))
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/regions/" + fmt.Sprint(req.Region) + "/echo",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EchoMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Service StreamAPI

// Service ZoneAPI

// Zones list localities the api is available in
func (s *ZoneAPI) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1}
}

type ZoneAPIGetZoneRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// GetZone: get a zone.
func (s *ZoneAPI) GetZone(req *ZoneAPIGetZoneRequest, opts ...scw.RequestOption) (*GetZoneResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/zone",
		Headers: http.Header{},
	}

	var resp GetZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ZoneAPIGetMetadataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// GetMetadata: get metadata.
func (s *ZoneAPI) GetMetadata(req *ZoneAPIGetMetadataRequest, opts ...scw.RequestOption) (*MetadataResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/test-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/metadata",
		Headers: http.Header{},
	}

	var resp MetadataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ZoneAPIPostEchoRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Str *string `json:"str"`

	Strs []string `json:"strs"`
}

func (s *ZoneAPI) PostEcho(req *ZoneAPIPostEchoRequest, opts ...scw.RequestOption) (*EchoMessage, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Str == nil {
		req.Str = scw.StringPtr(namegenerator.GetRandomName("name"))
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/test-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/echo",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EchoMessage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCharactersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCharactersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCharactersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Characters = append(r.Characters, results.Characters...)
	r.TotalCount += uint32(len(results.Characters))
	return uint32(len(results.Characters)), nil
}
