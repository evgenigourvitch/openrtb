package rtb

// 3.2.15 Object: Segment
// 
// Segment objects are essentially key-value pairs that convey specific units of data about the user. The
// parent Data object is a collection of such values from a given data provider. The specific segment
// names and value options must be published by the exchange a priori to its bidders.
type Segment struct {

	// Attribute:
	//   id
	// Type:
	//   string
	// Description:
	//   ID of the data segment specific to the data provider.
	ID string `json:"id"`

	// Attribute:
	//   name
	// Type:
	//   string
	// Description:
	//   Name of the data segment specific to the data provider.
	Name string `json:"name"`

	// Attribute:
	//   value
	// Type:
	//   string
	// Description:
	//   String representation of the data segment value.
	Value string `json:"value"`

	// Attribute:
	//   ext
	// Type:
	//   object
	// Description:
	//   Placeholder for exchange-specific extensions to OpenRTB.
	Ext Ext `json:"ext"`
}