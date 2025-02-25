// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package common

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Address) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Address[number], err)
}

func (x *Address) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Street, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.City, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.State, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Country, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Address) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Zip, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ProductItem) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductItem[number], err)
}

func (x *ProductItem) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.ProductId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *ProductItem) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Quantity, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Address) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *Address) fastWriteField1(buf []byte) (offset int) {
	if x.Street == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetStreet())
	return offset
}

func (x *Address) fastWriteField2(buf []byte) (offset int) {
	if x.City == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetCity())
	return offset
}

func (x *Address) fastWriteField3(buf []byte) (offset int) {
	if x.State == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetState())
	return offset
}

func (x *Address) fastWriteField4(buf []byte) (offset int) {
	if x.Country == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetCountry())
	return offset
}

func (x *Address) fastWriteField5(buf []byte) (offset int) {
	if x.Zip == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetZip())
	return offset
}

func (x *ProductItem) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ProductItem) fastWriteField1(buf []byte) (offset int) {
	if x.ProductId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetProductId())
	return offset
}

func (x *ProductItem) fastWriteField2(buf []byte) (offset int) {
	if x.Quantity == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 2, x.GetQuantity())
	return offset
}

func (x *Address) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *Address) sizeField1() (n int) {
	if x.Street == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetStreet())
	return n
}

func (x *Address) sizeField2() (n int) {
	if x.City == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetCity())
	return n
}

func (x *Address) sizeField3() (n int) {
	if x.State == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetState())
	return n
}

func (x *Address) sizeField4() (n int) {
	if x.Country == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetCountry())
	return n
}

func (x *Address) sizeField5() (n int) {
	if x.Zip == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetZip())
	return n
}

func (x *ProductItem) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ProductItem) sizeField1() (n int) {
	if x.ProductId == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetProductId())
	return n
}

func (x *ProductItem) sizeField2() (n int) {
	if x.Quantity == 0 {
		return n
	}
	n += fastpb.SizeUint32(2, x.GetQuantity())
	return n
}

var fieldIDToName_Address = map[int32]string{
	1: "Street",
	2: "City",
	3: "State",
	4: "Country",
	5: "Zip",
}

var fieldIDToName_ProductItem = map[int32]string{
	1: "ProductId",
	2: "Quantity",
}
