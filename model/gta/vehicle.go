package gta

type Vehicle struct {
	ID              int     //
	Category        string  //
	CategoryName    string  //
	Brand           string  //
	Name            string  //
	URL             string  //
	Type            string  //
	Conditional     string  //
	Speed           float64 //
	Acceleration    float64 //
	Braking         float64 //
	Handling        float64 //
	TopSpeed        bool    //
	TopAcceleration bool    //
	TopBraking      bool    //
	TopHandling     bool    //
	ForSale         bool    //
	Website         string  //
	Cost            string  //
	Seats           int     //
	Personal        bool    //
	Premium         bool    //
	Moddable        bool    //
	SuperModdable   bool    //
	Sellable        bool    //
	SellPrice       string  //
	MainImageURL    string  //
	MainImage       []byte  //
	ActionImageURL  string  //
	ActionImage     []byte  //
}

func (Vehicle) TableName() string {
	return "gta5_vehicles"
}
