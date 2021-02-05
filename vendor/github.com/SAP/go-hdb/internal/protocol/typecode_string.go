// Code generated by "stringer -type=typeCode"; DO NOT EDIT.

package protocol

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[tcNullL-0]
	_ = x[tcTinyint-1]
	_ = x[tcSmallint-2]
	_ = x[tcInteger-3]
	_ = x[tcBigint-4]
	_ = x[tcDecimal-5]
	_ = x[tcReal-6]
	_ = x[tcDouble-7]
	_ = x[tcChar-8]
	_ = x[tcVarchar-9]
	_ = x[tcNchar-10]
	_ = x[tcNvarchar-11]
	_ = x[tcBinary-12]
	_ = x[tcVarbinary-13]
	_ = x[tcDate-14]
	_ = x[tcTime-15]
	_ = x[tcTimestamp-16]
	_ = x[tcTimetz-17]
	_ = x[tcTimeltz-18]
	_ = x[tcTimestampTz-19]
	_ = x[tcTimestampLtz-20]
	_ = x[tcIntervalYm-21]
	_ = x[tcIntervalDs-22]
	_ = x[tcRowid-23]
	_ = x[tcUrowid-24]
	_ = x[tcClob-25]
	_ = x[tcNclob-26]
	_ = x[tcBlob-27]
	_ = x[tcBoolean-28]
	_ = x[tcString-29]
	_ = x[tcNstring-30]
	_ = x[tcLocator-31]
	_ = x[tcNlocator-32]
	_ = x[tcBstring-33]
	_ = x[tcDecimalDigitArray-34]
	_ = x[tcVarchar2-35]
	_ = x[tcTable-45]
	_ = x[tcSmalldecimal-47]
	_ = x[tcAbapstream-48]
	_ = x[tcAbapstruct-49]
	_ = x[tcAarray-50]
	_ = x[tcText-51]
	_ = x[tcShorttext-52]
	_ = x[tcBintext-53]
	_ = x[tcAlphanum-55]
	_ = x[tcLongdate-61]
	_ = x[tcSeconddate-62]
	_ = x[tcDaydate-63]
	_ = x[tcSecondtime-64]
	_ = x[tcClocator-70]
	_ = x[tcBlobDiskReserved-71]
	_ = x[tcClobDiskReserved-72]
	_ = x[tcNclobDiskReserved-73]
	_ = x[tcStGeometry-74]
	_ = x[tcStPoint-75]
	_ = x[tcFixed16-76]
	_ = x[tcAbapItab-77]
	_ = x[tcRecordRowStore-78]
	_ = x[tcRecordColumnStore-79]
	_ = x[tcFixed8-81]
	_ = x[tcFixed12-82]
	_ = x[tcCiphertext-90]
	_ = x[tcTableRef-126]
	_ = x[tcTableRows-127]
}

const (
	_typeCode_name_0 = "tcNullLtcTinyinttcSmallinttcIntegertcBiginttcDecimaltcRealtcDoubletcChartcVarchartcNchartcNvarchartcBinarytcVarbinarytcDatetcTimetcTimestamptcTimetztcTimeltztcTimestampTztcTimestampLtztcIntervalYmtcIntervalDstcRowidtcUrowidtcClobtcNclobtcBlobtcBooleantcStringtcNstringtcLocatortcNlocatortcBstringtcDecimalDigitArraytcVarchar2"
	_typeCode_name_1 = "tcTable"
	_typeCode_name_2 = "tcSmalldecimaltcAbapstreamtcAbapstructtcAarraytcTexttcShorttexttcBintext"
	_typeCode_name_3 = "tcAlphanum"
	_typeCode_name_4 = "tcLongdatetcSeconddatetcDaydatetcSecondtime"
	_typeCode_name_5 = "tcClocatortcBlobDiskReservedtcClobDiskReservedtcNclobDiskReservedtcStGeometrytcStPointtcFixed16tcAbapItabtcRecordRowStoretcRecordColumnStore"
	_typeCode_name_6 = "tcFixed8tcFixed12"
	_typeCode_name_7 = "tcCiphertext"
	_typeCode_name_8 = "tcTableReftcTableRows"
)

var (
	_typeCode_index_0 = [...]uint16{0, 7, 16, 26, 35, 43, 52, 58, 66, 72, 81, 88, 98, 106, 117, 123, 129, 140, 148, 157, 170, 184, 196, 208, 215, 223, 229, 236, 242, 251, 259, 268, 277, 287, 296, 315, 325}
	_typeCode_index_2 = [...]uint8{0, 14, 26, 38, 46, 52, 63, 72}
	_typeCode_index_4 = [...]uint8{0, 10, 22, 31, 43}
	_typeCode_index_5 = [...]uint8{0, 10, 28, 46, 65, 77, 86, 95, 105, 121, 140}
	_typeCode_index_6 = [...]uint8{0, 8, 17}
	_typeCode_index_8 = [...]uint8{0, 10, 21}
)

func (i typeCode) String() string {
	switch {
	case i <= 35:
		return _typeCode_name_0[_typeCode_index_0[i]:_typeCode_index_0[i+1]]
	case i == 45:
		return _typeCode_name_1
	case 47 <= i && i <= 53:
		i -= 47
		return _typeCode_name_2[_typeCode_index_2[i]:_typeCode_index_2[i+1]]
	case i == 55:
		return _typeCode_name_3
	case 61 <= i && i <= 64:
		i -= 61
		return _typeCode_name_4[_typeCode_index_4[i]:_typeCode_index_4[i+1]]
	case 70 <= i && i <= 79:
		i -= 70
		return _typeCode_name_5[_typeCode_index_5[i]:_typeCode_index_5[i+1]]
	case 81 <= i && i <= 82:
		i -= 81
		return _typeCode_name_6[_typeCode_index_6[i]:_typeCode_index_6[i+1]]
	case i == 90:
		return _typeCode_name_7
	case 126 <= i && i <= 127:
		i -= 126
		return _typeCode_name_8[_typeCode_index_8[i]:_typeCode_index_8[i+1]]
	default:
		return "typeCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
