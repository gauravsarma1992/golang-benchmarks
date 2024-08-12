package copybyvalueandreference

type (
	JustOneField struct {
		FieldOne uint32
	}

	ManyFields struct {
		FieldOne   uint32
		FieldTwo   uint32
		FieldThree uint32
		FieldFour  uint32
		FieldFive  uint32
		FieldSix   uint32
		FieldSeven uint32
		FieldEight uint32
	}
)

func copyByValueForSmallStruct(obj JustOneField) (result uint32) {
	result += 100 + obj.FieldOne
	return
}

func copyByRefForSmallStruct(obj *JustOneField) (result uint32) {
	result += 100 + obj.FieldOne
	return
}

func copyByValueForLargeStruct(obj ManyFields) (result uint32) {
	result += 100 + obj.FieldOne +
		obj.FieldTwo +
		obj.FieldThree +
		obj.FieldFour +
		obj.FieldFive +
		obj.FieldSix +
		obj.FieldSeven +
		obj.FieldEight
	return
}

func copyByRefForLargeStruct(obj *ManyFields) (result uint32) {
	result += 100 + obj.FieldOne +
		obj.FieldTwo +
		obj.FieldThree +
		obj.FieldFour +
		obj.FieldFive +
		obj.FieldSix +
		obj.FieldSeven +
		obj.FieldEight
	return
}
