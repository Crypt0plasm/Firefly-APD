package Firefly_Precision

import "math/big"

//
//	Scnd_F.Shortcuts.go					Secondary Function Shortcuts
//
//================================================
// 	Function List:
//
//	01 - NFS						new from string
//	02 - NFI						new from integer
//	03 - INT64						decimal to int64
//
//================================================
// Shortcuts functions of Main Functions
// Eases up function usage, discarding the Condition and Error outputs
//
//================================================
//
// Function 01 - NFS
//
// NFS creates a new decimal from a string s.
// Same as NewFromString, it only outputs the decimal
// It has no restriction on precision. Tested Accepted characters are "." and "-" for a negative sign
func NFS(s string) *Decimal {
	d := new(Decimal)
	//Further Tests can be added to detect wrong string construction.
	//It is always assumed the string number used with this function is of correct construction.
	d, _, _ = NewFromString(s)
	return d
}

//================================================
//
// Function 02 - NFI
//
// NFI creates a new decimal from a string integer
// THe integer must be within int64 range
// From 	-9,223,372,036,854,775,808 to
//		 9,223,372,036,854,775,807
func NFI(coeff int64) *Decimal {
	d := new(Decimal)
	d = New(coeff, 0)
	return d
}

//================================================
//
// Function 03 - NFBI
//
// NFBI creates a new decimal from a big.Int number
func NFBI(Number *big.Int) *Decimal {
	d := new(Decimal)
	d = NewWithBigInt(Number, 0)
	return d
}

//================================================
//
// Function 04 - INT64
//
// INT64 returns the int64 representation of x.
// No error is returned with this function
func INT64(Decimal2Convert *Decimal) int64 {
	var Result int64
	Result, _ = Decimal2Convert.Int64()
	return Result
}

//================================================
//================================================
//
// Function 05 - DTS
//
// DTS returns the string representation of x.
// No error is returned with this function
// Assumes "." as separator for decimals
func DTS(Input *Decimal) (Output string) {

	//Function makes a rune chain from a text string
	MakeRuneChain := func(Text string) []rune {
		Result := []rune(Text)
		return Result
	}

	//Function to insert an element (Value to insert) in a slice (in this example runes) at a given position (Index)
	InsertIntoRuneSlice := func(RuneSlice []rune, Index int, ValueToInsert rune) []rune {
		if len(RuneSlice) == Index { //nil or empty slice or after the last element
			return append(RuneSlice, ValueToInsert)
		}
		RuneSlice = append(RuneSlice[:Index+1], RuneSlice[Index:]...) //Index < len(RuneSlice)
		RuneSlice[Index] = ValueToInsert
		return RuneSlice
	}

	//Separator that is to be inserted into the string to be created.
	//The "." Character is used
	Separator := MakeRuneChain(".")[0]

	//Creating the Chain of runes representing the Decimal Number
	Coefficient := Input.Coeff
	DecimalAsLongText := Coefficient.Text(10)
	OriginalRuneSlice := MakeRuneChain(DecimalAsLongText)

	//Getting the Position where the Separator must be inserted
	//Assumes Exponent is lower than Zero, that is, there is precision in the decimal, ie: "8888.12345"
	//That Example would have a -5 Exponent
	Exponent := int(Input.Exponent)
	Position := len(OriginalRuneSlice) + Exponent

	if Exponent >= 0 {
		Output = DecimalAsLongText
	} else {
		Output = string(InsertIntoRuneSlice(OriginalRuneSlice, Position, Separator))
	}

	return
}
