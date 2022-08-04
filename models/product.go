package models

type ProductStatus string
type Size string

var IdCounter uint32 = 6

const (
	Available    ProductStatus = "available"
	BackOrdered  ProductStatus = "backordered"
	Discontinued ProductStatus = "discontinued"
)

const (
	XS   Size = "XS"
	S    Size = "S"
	M    Size = "M"
	L    Size = "L"
	XL   Size = "XL"
	XXL  Size = "XXL"
	XXXL Size = "XXXL"
)

type Product struct {
	ProductID     string        `json:"productID,omitempty" gorm:"primary_key;not_null"`
	Name          string        `json:"name" gorm:"not_null"`
	Description   string        `json:"description" gorm:"not_null"`
	ProductStatus ProductStatus `json:"productStatus,omitempty"`
	Sizes         Size          `json:"size,omitempty"`
	Color         string        `json:"color,omitempty"`
	ProductMSRP   float64       `json:"productMSRP,omitempty"`
	ProductSKU    string        `json:"productSKU,omitempty"`
}

var ProductsMasterList = [...]Product{
	{
		ProductID:     "PRD001",
		Name:          "M-Shield Blue Logo T-Shirt",
		Description:   "4.2 oz., 52/48 Airlume combed and ring spun cotton/poly with retail fit.",
		ProductStatus: Available,
		Sizes:         M,
		Color:         "Heather Navy",
		ProductMSRP:   20.00,
		ProductSKU:    "MU00-TS-HN-E",
	},
	{
		ProductID:     "PRD002",
		Name:          "Equality T-Shirt",
		Description:   "In partnership with #HeForShe, MuleSoft is focusing on gender equality and the importance of allyship. We have designed a special shirt to raise funds in support of UN Women to aid in their efforts to positively move the needle on this important issue everyday. 100% of the proceeds from the shirt sales will be donated directly to UN Women in support of #HeForShe. 4.2 oz., 100% Airlume combed and ring spun cotton with retail fit",
		ProductStatus: Available,
		Sizes:         S,
		Color:         "Heather Navy",
		ProductMSRP:   30.00,
		ProductSKU:    "MU00-ETS-BLK-E",
	},
	{
		ProductID:     "PRD003",
		Name:          "Women's Basecamp Puffer Vest",
		Description:   "Feel cozy to the core in this quality D/W/R water-resistant thermal vest. Not only does it give you the convenience of impromptu insulation, but its original packable design allows for easy stowage into its own pocket. Features a chin saver and an internal full-length storm flap. 100% Nylon. COLD [-10C (14F) to +10C (50F)].",
		ProductStatus: Available,
		Sizes:         M,
		Color:         "Navy",
		ProductMSRP:   78.00,
		ProductSKU:    "SKU: MU00-0180-NVY",
	},
	{
		ProductID:     "PRD004",
		Name:          "Keyboard Notes Journal",
		Description:   "This spiral bound journal has it all! Open it up and place it under your keyboard and you are ready for the day. Includes 2 recycled 3\"x 3\" Post-it pads (2 colors), 60 white lined sheets 5\"x 3.75\" and 75 white sheets with check-list/dot-grid 4.75x3.75. Overall dimensions: 12.25\"L x 4\"W",
		ProductStatus: Available,
		Color:         "Heather Navy",
		ProductMSRP:   9.50,
		ProductSKU:    "MU00-0204-000-000",
	},
	{
		ProductID:     "PRD005",
		Name:          "Mulesoft Baby Socks",
		Description:   "Start them out right with 100% custom jacquard/woven socks to keep their little toes warm. Fits approximately 0-3 years.",
		ProductStatus: Available,
		ProductMSRP:   10.00,
		ProductSKU:    "MU00-0178-BLU-000",
	},
	{
		ProductID:     "PRD006",
		Name:          "M-Shield Cap",
		Description:   "This low profile, unstructured, laundered Chino Twill (100% Cotton), 6-panel cap will be your new favorite. Has adjustable fabric strap back to fit most sizes",
		ProductStatus: Available,
		ProductMSRP:   12.00,
		ProductSKU:    "MU00-0200-NVY-000",
		Color:         "Navy",
	},
}
