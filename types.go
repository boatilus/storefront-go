package storefront

import "time"

// QueryRoot: The schema’s entry-point for queries. This acts as the public, top-level API from which all queries must start.
type QueryRoot struct {
	// Articles: List of the shop's articles.
	Articles Connection[Article] `json:"articles,omitempty"`
	// Blog: Fetch a specific `Blog` by one of its unique attributes.
	Blog Blog `json:"blog,omitempty"`
	// Deprecated: BlogByHandle: Find a blog by its handle.
	BlogByHandle Blog `json:"blogByHandle,omitempty"`
	// Blogs: List of the shop's blogs.
	Blogs Connection[Blog] `json:"blogs,omitempty"`
	// Cart: Find a cart by its ID.
	Cart Cart `json:"cart,omitempty"`
	// Collection: Fetch a specific `Collection` by one of its unique attributes.
	Collection Collection `json:"collection,omitempty"`
	// Deprecated: CollectionByHandle: Find a collection by its handle.
	CollectionByHandle Collection `json:"collectionByHandle,omitempty"`
	// Collections: List of the shop’s collections.
	Collections Connection[Collection] `json:"collections,omitempty"`
	// Customer: Find a customer by its access token.
	Customer Customer `json:"customer,omitempty"`
	// Localization: Returns the localized experiences configured for the shop.
	Localization Localization `json:"localization,omitempty"`
	/*
	   Locations: List of the shop's locations that support in-store pickup.

	   When sorting by distance, you must specify a location via the `near` argument.
	*/
	Locations Connection[Location] `json:"locations,omitempty"`
	// Node: Returns a specific node by ID.
	Node Node `json:"node,omitempty"`
	// Nodes: Returns the list of nodes with the given IDs.
	Nodes []Node `json:"nodes,omitempty"`
	// Page: Fetch a specific `Page` by one of its unique attributes.
	Page Page `json:"page,omitempty"`
	// Deprecated: PageByHandle: Find a page by its handle.
	PageByHandle Page `json:"pageByHandle,omitempty"`
	// Pages: List of the shop's pages.
	Pages Connection[Page] `json:"pages,omitempty"`
	// Product: Fetch a specific `Product` by one of its unique attributes.
	Product Product `json:"product,omitempty"`
	// Deprecated: ProductByHandle: Find a product by its handle.
	ProductByHandle Product `json:"productByHandle,omitempty"`
	/*
	   ProductRecommendations: Find recommended products related to a given `product_id`.
	   To learn more about how recommendations are generated, see
	   [*Showing product recommendations on product pages*](https://help.shopify.com/themes/development/recommended-products).
	*/
	ProductRecommendations []Product `json:"productRecommendations,omitempty"`
	/*
	   ProductTags: Tags added to products.
	   Additional access scope required: unauthenticated_read_product_tags.
	*/
	ProductTags Connection[string] `json:"productTags,omitempty"`
	// ProductTypes: List of product types for the shop's products that are published to your app.
	ProductTypes Connection[string] `json:"productTypes,omitempty"`
	// Products: List of the shop’s products.
	Products Connection[Product] `json:"products,omitempty"`
	// PublicApiVersions: The list of public Storefront API versions, including supported, release candidate and unstable versions.
	PublicApiVersions []ApiVersion `json:"publicApiVersions,omitempty"`
	// Shop: The shop associated with the storefront access token.
	Shop Shop `json:"shop,omitempty"`
}

// ArticleSortKeys: The set of valid sort keys for the Article query.
type ArticleSortKeys string

const (
	ArticleSortKeysTitle       ArticleSortKeys = "TITLE"
	ArticleSortKeysBlogTitle   ArticleSortKeys = "BLOG_TITLE"
	ArticleSortKeysAuthor      ArticleSortKeys = "AUTHOR"
	ArticleSortKeysUpdatedAt   ArticleSortKeys = "UPDATED_AT"
	ArticleSortKeysPublishedAt ArticleSortKeys = "PUBLISHED_AT"
	ArticleSortKeysId          ArticleSortKeys = "ID"
	ArticleSortKeysRelevance   ArticleSortKeys = "RELEVANCE"
)

// Article: An article in an online store blog.
type Article struct {
	// Deprecated: Author: The article's author.
	Author ArticleAuthor `json:"author,omitempty"`
	// AuthorV2: The article's author.
	AuthorV2 ArticleAuthor `json:"authorV2,omitempty"`
	// Blog: The blog that the article belongs to.
	Blog interface{} `json:"blog,omitempty"`
	// Comments: List of comments posted on the article.
	Comments Connection[Comment] `json:"comments,omitempty"`
	// Content: Stripped content of the article, single line with HTML tags removed.
	Content string `json:"content,omitempty"`
	// ContentHTML: The content of the article, complete with HTML formatting.
	ContentHTML string `json:"contentHtml,omitempty"`
	// Excerpt: Stripped excerpt of the article, single line with HTML tags removed.
	Excerpt string `json:"excerpt,omitempty"`
	// ExcerptHTML: The excerpt of the article, complete with HTML formatting.
	ExcerptHTML string `json:"excerptHtml,omitempty"`
	/*
	   Handle: A human-friendly unique string for the Article automatically generated from its title.
	*/
	Handle string `json:"handle,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Image: The image associated with the article.
	Image Image `json:"image,omitempty"`
	// Metafield: Returns a metafield found by namespace and key.
	Metafield Metafield `json:"metafield,omitempty"`
	// Deprecated: Metafields: A paginated list of metafields associated with the resource.
	Metafields Connection[Metafield] `json:"metafields,omitempty"`
	// OnlineStoreURL: The URL used for viewing the resource on the shop's Online Store. Returns `null` if the resource is currently not published to the Online Store sales channel.
	OnlineStoreURL string `json:"onlineStoreUrl,omitempty"`
	// PublishedAt: The date and time when the article was published.
	PublishedAt time.Time `json:"publishedAt,omitempty"`
	// SEO: The article’s SEO information.
	SEO SEO `json:"seo,omitempty"`
	// Tags: A categorization that a article can be tagged with.
	Tags []string `json:"tags,omitempty"`
	// Title: The article’s name.
	Title string `json:"title,omitempty"`
}

/*
Node: An object with an ID field to support global identification, in accordance with the
[Relay specification](https://relay.dev/graphql/objectidentification.htm#sec-Node-Interface).
This interface is used by the [node](https://shopify.dev/api/admin-graphql/unstable/queries/node)
and [nodes](https://shopify.dev/api/admin-graphql/unstable/queries/nodes) queries.
*/
type Node struct {
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
}

// HasMetafields: Represents information about the metafields associated to the specified resource.
type HasMetafields struct {
	// Metafield: Returns a metafield found by namespace and key.
	Metafield Metafield `json:"metafield,omitempty"`
	// Deprecated: Metafields: A paginated list of metafields associated with the resource.
	Metafields Connection[Metafield] `json:"metafields,omitempty"`
}

/*
Metafield: Metafields represent custom metadata attached to a resource. Metafields can be sorted into namespaces and are
comprised of keys, values, and value types.
*/
type Metafield struct {
	// CreatedAt: The date and time when the storefront metafield was created.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Description: The description of a metafield.
	Description string `json:"description,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Key: The key name for a metafield.
	Key string `json:"key,omitempty"`
	// Namespace: The namespace for a metafield.
	Namespace string `json:"namespace,omitempty"`
	// ParentResource: The parent object that the metafield belongs to.
	ParentResource string `json:"parentResource,omitempty"`
	// Reference: Returns a reference object if the metafield definition's type is a resource reference.
	Reference string `json:"reference,omitempty"`
	/*
	   Type: The type name of the metafield.
	   See the list of [supported types](https://shopify.dev/apps/metafields/definitions/types).
	*/
	Type string `json:"type,omitempty"`
	// UpdatedAt: The date and time when the storefront metafield was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// Value: The value of a metafield.
	Value string `json:"value,omitempty"`
}

// Blog: An online store blog.
type Blog struct {
	// ArticleByHandle: Find an article by its handle.
	ArticleByHandle Article `json:"articleByHandle,omitempty"`
	// Articles: List of the blog's articles.
	Articles Connection[Article] `json:"articles,omitempty"`
	// Authors: The authors who have contributed to the blog.
	Authors []ArticleAuthor `json:"authors,omitempty"`
	/*
	   Handle: A human-friendly unique string for the Blog automatically generated from its title.
	*/
	Handle string `json:"handle,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Metafield: Returns a metafield found by namespace and key.
	Metafield Metafield `json:"metafield,omitempty"`
	// Deprecated: Metafields: A paginated list of metafields associated with the resource.
	Metafields Connection[Metafield] `json:"metafields,omitempty"`
	// OnlineStoreURL: The URL used for viewing the resource on the shop's Online Store. Returns `null` if the resource is currently not published to the Online Store sales channel.
	OnlineStoreURL string `json:"onlineStoreUrl,omitempty"`
	// SEO: The blog's SEO information.
	SEO SEO `json:"seo,omitempty"`
	// Title: The blogs’s title.
	Title string `json:"title,omitempty"`
}

// OnlineStorePublishable: Represents a resource that can be published to the Online Store sales channel.
type OnlineStorePublishable struct {
	// OnlineStoreURL: The URL used for viewing the resource on the shop's Online Store. Returns `null` if the resource is currently not published to the Online Store sales channel.
	OnlineStoreURL string `json:"onlineStoreUrl,omitempty"`
}

// ArticleAuthor: The author of an article.
type ArticleAuthor struct {
	// Bio: The author's bio.
	Bio string `json:"bio,omitempty"`
	// Email: The author’s email.
	Email string `json:"email,omitempty"`
	// FirstName: The author's first name.
	FirstName string `json:"firstName,omitempty"`
	// LastName: The author's last name.
	LastName string `json:"lastName,omitempty"`
	// Name: The author's full name.
	Name string `json:"name,omitempty"`
}

/*
PageInfo: Returns information about pagination in a connection, in accordance with the
[Relay specification](https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo).
*/
type PageInfo struct {
	// HasNextPage: Whether there are more pages to fetch following the current page.
	HasNextPage bool `json:"hasNextPage,omitempty"`
	// HasPreviousPage: Whether there are any pages prior to the current page.
	HasPreviousPage bool `json:"hasPreviousPage,omitempty"`
}

// SEO: SEO information.
type SEO struct {
	// Description: The meta description.
	Description string `json:"description,omitempty"`
	// Title: The SEO title.
	Title string `json:"title,omitempty"`
}

// Collection: A collection represents a grouping of products that a shop owner can create to organize them or make their shops easier to browse.
type Collection struct {
	// Description: Stripped description of the collection, single line with HTML tags removed.
	Description string `json:"description,omitempty"`
	// DescriptionHTML: The description of the collection, complete with HTML formatting.
	DescriptionHTML string `json:"descriptionHtml,omitempty"`
	/*
	   Handle: A human-friendly unique string for the collection automatically generated from its title.
	   Limit of 255 characters.
	*/
	Handle string `json:"handle,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Image: Image associated with the collection.
	Image Image `json:"image,omitempty"`
	// Metafield: Returns a metafield found by namespace and key.
	Metafield Metafield `json:"metafield,omitempty"`
	// Deprecated: Metafields: A paginated list of metafields associated with the resource.
	Metafields Connection[Metafield] `json:"metafields,omitempty"`
	// OnlineStoreURL: The URL used for viewing the resource on the shop's Online Store. Returns `null` if the resource is currently not published to the Online Store sales channel.
	OnlineStoreURL string `json:"onlineStoreUrl,omitempty"`
	// Products: List of products in the collection.
	Products Connection[Product] `json:"products,omitempty"`
	// Title: The collection’s name. Limit of 255 characters.
	Title string `json:"title,omitempty"`
	// UpdatedAt: The date and time when the collection was last modified.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// Image: Represents an image resource.
type Image struct {
	// AltText: A word or phrase to share the nature or contents of an image.
	AltText string `json:"altText,omitempty"`
	// Height: The original height of the image in pixels. Returns `null` if the image is not hosted by Shopify.
	Height int `json:"height,omitempty"`
	// Id: A unique identifier for the image.
	Id string `json:"id,omitempty"`
	/*
	   Deprecated: OriginalSrc: The location of the original image as a URL.

	   If there are any existing transformations in the original source URL, they will remain and not be stripped.
	*/
	OriginalSrc string `json:"originalSrc,omitempty"`
	// Deprecated: Src: The location of the image as a URL.
	Src string `json:"src,omitempty"`
	/*
	   Deprecated: TransformedSrc: The location of the transformed image as a URL.

	   All transformation arguments are considered "best-effort". If they can be applied to an image, they will be.
	   Otherwise any transformations which an image type does not support will be ignored.
	*/
	TransformedSrc string `json:"transformedSrc,omitempty"`
	/*
	   URL: The location of the image as a URL.

	   If no transform options are specified, then the original image will be preserved including any pre-applied transforms.

	   All transformation options are considered "best-effort". Any transformation that the original image type doesn't support will be ignored.

	   If you need multiple variations of the same image, then you can use [GraphQL aliases](https://graphql.org/learn/queries/#aliases).
	*/
	URL string `json:"url,omitempty"`
	// Width: The original width of the image in pixels. Returns `null` if the image is not hosted by Shopify.
	Width int `json:"width,omitempty"`
}

// CropRegion: The part of the image that should remain after cropping.
type CropRegion string

const (
	CropRegionCenter CropRegion = "CENTER"
	CropRegionTop    CropRegion = "TOP"
	CropRegionBottom CropRegion = "BOTTOM"
	CropRegionLeft   CropRegion = "LEFT"
	CropRegionRight  CropRegion = "RIGHT"
)

// ImageContentType: List of supported image content types.
type ImageContentType string

const (
	ImageContentTypePNG  ImageContentType = "PNG"
	ImageContentTypeJPG  ImageContentType = "JPG"
	ImageContentTypeWebP ImageContentType = "WEBP"
)

// ProductCollectionSortKeys: The set of valid sort keys for the ProductCollection query.
type ProductCollectionSortKeys string

const (
	ProductCollectionSortKeysTitle             ProductCollectionSortKeys = "TITLE"
	ProductCollectionSortKeysPrice             ProductCollectionSortKeys = "PRICE"
	ProductCollectionSortKeysBestSelling       ProductCollectionSortKeys = "BEST_SELLING"
	ProductCollectionSortKeysCreated           ProductCollectionSortKeys = "CREATED"
	ProductCollectionSortKeysId                ProductCollectionSortKeys = "ID"
	ProductCollectionSortKeysManual            ProductCollectionSortKeys = "MANUAL"
	ProductCollectionSortKeysCollectionDefault ProductCollectionSortKeys = "COLLECTION_DEFAULT"
	ProductCollectionSortKeysRelevance         ProductCollectionSortKeys = "RELEVANCE"
)

/*
Product: A product represents an individual item for sale in a Shopify store. Products are often physical, but they don't have to be.
For example, a digital download (such as a movie, music or ebook file) also qualifies as a product, as do services (such as equipment rental, work for hire, customization of another product or an extended warranty).
*/
type Product struct {
	// AvailableForSale: Indicates if at least one product variant is available for sale.
	AvailableForSale bool `json:"availableForSale,omitempty"`
	// Collections: List of collections a product belongs to.
	Collections Connection[Collection] `json:"collections,omitempty"`
	// CompareAtPriceRange: The compare at price of the product across all variants.
	CompareAtPriceRange ProductPriceRange `json:"compareAtPriceRange,omitempty"`
	// CreatedAt: The date and time when the product was created.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Description: Stripped description of the product, single line with HTML tags removed.
	Description string `json:"description,omitempty"`
	// DescriptionHTML: The description of the product, complete with HTML formatting.
	DescriptionHTML string `json:"descriptionHtml,omitempty"`
	/*
	   FeaturedImage: The featured image for the product.

	   This field is functionally equivalent to `images(first: 1)`.
	*/
	FeaturedImage Image `json:"featuredImage,omitempty"`
	/*
	   Handle: A human-friendly unique string for the Product automatically generated from its title.
	   They are used by the Liquid templating language to refer to objects.
	*/
	Handle string `json:"handle,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Images: List of images associated with the product.
	Images Connection[Image] `json:"images,omitempty"`
	// Media: The media associated with the product.
	Media Connection[Media] `json:"media,omitempty"`
	// Metafield: Returns a metafield found by namespace and key.
	Metafield Metafield `json:"metafield,omitempty"`
	// Deprecated: Metafields: A paginated list of metafields associated with the resource.
	Metafields Connection[Metafield] `json:"metafields,omitempty"`
	// OnlineStoreURL: The URL used for viewing the resource on the shop's Online Store. Returns `null` if the resource is currently not published to the Online Store sales channel.
	OnlineStoreURL string `json:"onlineStoreUrl,omitempty"`
	// Options: List of product options.
	Options []ProductOption `json:"options,omitempty"`
	// PriceRange: The price range.
	PriceRange ProductPriceRange `json:"priceRange,omitempty"`
	// ProductType: A categorization that a product can be tagged with, commonly used for filtering and searching.
	ProductType string `json:"productType,omitempty"`
	// PublishedAt: The date and time when the product was published to the channel.
	PublishedAt time.Time `json:"publishedAt,omitempty"`
	// RequiresSellingPlan: Whether the product can only be purchased with a selling plan.
	RequiresSellingPlan bool `json:"requiresSellingPlan,omitempty"`
	// SellingPlanGroups: A list of a product's available selling plan groups. A selling plan group represents a selling method. For example, 'Subscribe and save' is a selling method where customers pay for goods or services per delivery. A selling plan group contains individual selling plans.
	SellingPlanGroups Connection[SellingPlanGroup] `json:"sellingPlanGroups,omitempty"`
	// SEO: The product's SEO information.
	SEO SEO `json:"seo,omitempty"`
	/*
	   Tags: A comma separated list of tags that have been added to the product.
	   Additional access scope required for private apps: unauthenticated_read_product_tags.
	*/
	Tags []string `json:"tags,omitempty"`
	// Title: The product’s title.
	Title string `json:"title,omitempty"`
	// TotalInventory: The total quantity of inventory in stock for this Product.
	TotalInventory int `json:"totalInventory,omitempty"`
	/*
	   UpdatedAt: The date and time when the product was last modified.
	   A product's `updatedAt` value can change for different reasons. For example, if an order
	   is placed for a product that has inventory tracking set up, then the inventory adjustment
	   is counted as an update.
	*/
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	/*
	   VariantBySelectedOptions: Find a product’s variant based on its selected options.
	   This is useful for converting a user’s selection of product options into a single matching variant.
	   If there is not a variant for the selected options, `null` will be returned.
	*/
	VariantBySelectedOptions ProductVariant `json:"variantBySelectedOptions,omitempty"`
	// Variants: List of the product’s variants.
	Variants Connection[ProductVariant] `json:"variants,omitempty"`
	// Vendor: The product’s vendor name.
	Vendor string `json:"vendor,omitempty"`
}

// ProductPriceRange: The price range of the product.
type ProductPriceRange struct {
	// MaxVariantPrice: The highest variant's price.
	MaxVariantPrice MoneyV2 `json:"maxVariantPrice,omitempty"`
	// MinVariantPrice: The lowest variant's price.
	MinVariantPrice MoneyV2 `json:"minVariantPrice,omitempty"`
}

/*
MoneyV2: A monetary value with currency.
*/
type MoneyV2 struct {
	// Amount: Decimal money amount.
	Amount float64 `json:"amount,omitempty"`
	// CurrencyCode: Currency of the money.
	CurrencyCode string `json:"currencyCode,omitempty"`
}

// ProductImageSortKeys: The set of valid sort keys for the ProductImage query.
type ProductImageSortKeys string

const (
	ProductImageSortKeysCreatedAt ProductImageSortKeys = "CREATED_AT"
	ProductImageSortKeysPosition  ProductImageSortKeys = "POSITION"
	ProductImageSortKeysId        ProductImageSortKeys = "ID"
	ProductImageSortKeysRelevance ProductImageSortKeys = "RELEVANCE"
)

// ProductMediaSortKeys: The set of valid sort keys for the ProductMedia query.
type ProductMediaSortKeys string

const (
	ProductMediaSortKeysPosition  ProductMediaSortKeys = "POSITION"
	ProductMediaSortKeysId        ProductMediaSortKeys = "ID"
	ProductMediaSortKeysRelevance ProductMediaSortKeys = "RELEVANCE"
)

// Media: Represents a media interface.
type Media struct {
	// Alt: A word or phrase to share the nature or contents of a media.
	Alt string `json:"alt,omitempty"`
	// MediaContentType: The media content type.
	MediaContentType MediaContentType `json:"mediaContentType,omitempty"`
	// PreviewImage: The preview image for the media.
	PreviewImage Image `json:"previewImage,omitempty"`
}

// MediaContentType: The possible content types for a media object.
type MediaContentType string

const (
	MediaContentTypeExternalVideo MediaContentType = "EXTERNAL_VIDEO"
	MediaContentTypeImage         MediaContentType = "IMAGE"
	MediaContentTypeModel3D       MediaContentType = "MODEL_3D"
	MediaContentTypeVideo         MediaContentType = "VIDEO"
)

/*
ProductOption: Product property names like "Size", "Color", and "Material" that the customers can select.
Variants are selected based on permutations of these options.
255 characters limit each.
*/
type ProductOption struct {
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Name: The product option’s name.
	Name string `json:"name,omitempty"`
	// Values: The corresponding value to the product option name.
	Values []string `json:"values,omitempty"`
}

// SellingPlanGroup: Represents a selling method. For example, 'Subscribe and save' is a selling method where customers pay for goods or services per delivery. A selling plan group contains individual selling plans.
type SellingPlanGroup struct {
	// AppName: A display friendly name for the app that created the selling plan group.
	AppName string `json:"appName,omitempty"`
	// Name: The name of the selling plan group.
	Name string `json:"name,omitempty"`
	// Options: Represents the selling plan options available in the drop-down list in the storefront. For example, 'Delivery every week' or 'Delivery every 2 weeks' specifies the delivery frequency options for the product.
	Options []SellingPlanGroupOption `json:"options,omitempty"`
	// SellingPlans: A list of selling plans in a selling plan group. A selling plan is a representation of how products and variants can be sold and purchased. For example, an individual selling plan could be '6 weeks of prepaid granola, delivered weekly'.
	SellingPlans Connection[SellingPlan] `json:"sellingPlans,omitempty"`
}

// SellingPlanGroupOption: Represents an option on a selling plan group that's available in the drop-down list in the storefront.
type SellingPlanGroupOption struct {
	// Name: The name of the option. For example, 'Delivery every'.
	Name string `json:"name,omitempty"`
	// Values: The values for the options specified by the selling plans in the selling plan group. For example, '1 week', '2 weeks', '3 weeks'.
	Values []string `json:"values,omitempty"`
}

// SellingPlan: Represents how products and variants can be sold and purchased.
type SellingPlan struct {
	// Description: The description of the selling plan.
	Description string `json:"description,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Name: The name of the selling plan. For example, '6 weeks of prepaid granola, delivered weekly'.
	Name string `json:"name,omitempty"`
	// Options: Represents the selling plan options available in the drop-down list in the storefront. For example, 'Delivery every week' or 'Delivery every 2 weeks' specifies the delivery frequency options for the product.
	Options []SellingPlanOption `json:"options,omitempty"`
	// PriceAdjustments: Represents how a selling plan affects pricing when a variant is purchased with a selling plan.
	PriceAdjustments []SellingPlanPriceAdjustment `json:"priceAdjustments,omitempty"`
	// RecurringDeliveries: Whether purchasing the selling plan will result in multiple deliveries.
	RecurringDeliveries bool `json:"recurringDeliveries,omitempty"`
}

// SellingPlanOption: An option provided by a Selling Plan.
type SellingPlanOption struct {
	// Name: The name of the option (ie "Delivery every").
	Name string `json:"name,omitempty"`
	// Value: The value of the option (ie "Month").
	Value string `json:"value,omitempty"`
}

// SellingPlanPriceAdjustment: Represents by how much the price of a variant associated with a selling plan is adjusted. Each variant can have up to two price adjustments.
type SellingPlanPriceAdjustment struct {
	// AdjustmentValue: The type of price adjustment. An adjustment value can have one of three types: percentage, amount off, or a new price.
	AdjustmentValue string `json:"adjustmentValue,omitempty"`
	// OrderCount: The number of orders that the price adjustment applies to If the price adjustment always applies, then this field is `null`.
	OrderCount int `json:"orderCount,omitempty"`
}

// SellingPlanFixedAmountPriceAdjustment: A fixed amount that's deducted from the original variant price. For example, $10.00 off.
type SellingPlanFixedAmountPriceAdjustment struct {
	// AdjustmentAmount: The money value of the price adjustment.
	AdjustmentAmount MoneyV2 `json:"adjustmentAmount,omitempty"`
}

// SellingPlanFixedPriceAdjustment: A fixed price adjustment for a variant that's purchased with a selling plan.
type SellingPlanFixedPriceAdjustment struct {
	// Price: A new price of the variant when it's purchased with the selling plan.
	Price MoneyV2 `json:"price,omitempty"`
}

// SellingPlanPercentagePriceAdjustment: A percentage amount that's deducted from the original variant price. For example, 10% off.
type SellingPlanPercentagePriceAdjustment struct {
	// AdjustmentPercentage: The percentage value of the price adjustment.
	AdjustmentPercentage int `json:"adjustmentPercentage,omitempty"`
}

// ProductVariant: A product variant represents a different version of a product, such as differing sizes or differing colors.
type ProductVariant struct {
	// AvailableForSale: Indicates if the product variant is available for sale.
	AvailableForSale bool `json:"availableForSale,omitempty"`
	// Barcode: The barcode (for example, ISBN, UPC, or GTIN) associated with the variant.
	Barcode string `json:"barcode,omitempty"`
	// Deprecated: CompareAtPrice: The compare at price of the variant. This can be used to mark a variant as on sale, when `compareAtPrice` is higher than `price`.
	CompareAtPrice string `json:"compareAtPrice,omitempty"`
	// CompareAtPriceV2: The compare at price of the variant. This can be used to mark a variant as on sale, when `compareAtPriceV2` is higher than `priceV2`.
	CompareAtPriceV2 MoneyV2 `json:"compareAtPriceV2,omitempty"`
	// CurrentlyNotInStock: Whether a product is out of stock but still available for purchase (used for backorders).
	CurrentlyNotInStock bool `json:"currentlyNotInStock,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	/*
	   Image: Image associated with the product variant. This field falls back to the product image if no image is available.
	*/
	Image Image `json:"image,omitempty"`
	// Metafield: Returns a metafield found by namespace and key.
	Metafield Metafield `json:"metafield,omitempty"`
	// Deprecated: Metafields: A paginated list of metafields associated with the resource.
	Metafields Connection[Metafield] `json:"metafields,omitempty"`
	// Deprecated: Price: The product variant’s price.
	Price string `json:"price,omitempty"`
	// PriceV2: The product variant’s price.
	PriceV2 MoneyV2 `json:"priceV2,omitempty"`
	// Product: The product object that the product variant belongs to.
	Product interface{} `json:"product,omitempty"`
	// QuantityAvailable: The total sellable quantity of the variant for online sales channels.
	QuantityAvailable int `json:"quantityAvailable,omitempty"`
	// RequiresShipping: Whether a customer needs to provide a shipping address when placing an order for the product variant.
	RequiresShipping bool `json:"requiresShipping,omitempty"`
	// SelectedOptions: List of product options applied to the variant.
	SelectedOptions []SelectedOption `json:"selectedOptions,omitempty"`
	// SellingPlanAllocations: Represents an association between a variant and a selling plan. Selling plan allocations describe which selling plans are available for each variant, and what their impact is on pricing.
	SellingPlanAllocations Connection[SellingPlanAllocation] `json:"sellingPlanAllocations,omitempty"`
	// Sku: The SKU (stock keeping unit) associated with the variant.
	Sku string `json:"sku,omitempty"`
	// StoreAvailability: The in-store pickup availability of this variant by location.
	StoreAvailability Connection[StoreAvailability] `json:"storeAvailability,omitempty"`
	// Title: The product variant’s title.
	Title string `json:"title,omitempty"`
	// UnitPrice: The unit price value for the variant based on the variant's measurement.
	UnitPrice MoneyV2 `json:"unitPrice,omitempty"`
	// UnitPriceMeasurement: The unit price measurement for the variant.
	UnitPriceMeasurement UnitPriceMeasurement `json:"unitPriceMeasurement,omitempty"`
	// Weight: The weight of the product variant in the unit system specified with `weight_unit`.
	Weight float64 `json:"weight,omitempty"`
	// WeightUnit: Unit of measurement for weight.
	WeightUnit string `json:"weightUnit,omitempty"`
}

/*
SelectedOption: Properties used by customers to select a product variant.
Products can have multiple options, like different sizes or colors.
*/
type SelectedOption struct {
	// Name: The product option’s name.
	Name string `json:"name,omitempty"`
	// Value: The product option’s value.
	Value string `json:"value,omitempty"`
}

// SellingPlanAllocation: Represents an association between a variant and a selling plan. Selling plan allocations describe the options offered for each variant, and the price of the variant when purchased with a selling plan.
type SellingPlanAllocation struct {
	// PriceAdjustments: A list of price adjustments, with a maximum of two. When there are two, the first price adjustment goes into effect at the time of purchase, while the second one starts after a certain number of orders.
	PriceAdjustments []SellingPlanAllocationPriceAdjustment `json:"priceAdjustments,omitempty"`
	// SellingPlan: A representation of how products and variants can be sold and purchased. For example, an individual selling plan could be '6 weeks of prepaid granola, delivered weekly'.
	SellingPlan SellingPlan `json:"sellingPlan,omitempty"`
}

// SellingPlanAllocationPriceAdjustment: The resulting prices for variants when they're purchased with a specific selling plan.
type SellingPlanAllocationPriceAdjustment struct {
	// CompareAtPrice: The price of the variant when it's purchased without a selling plan for the same number of deliveries. For example, if a customer purchases 6 deliveries of $10.00 granola separately, then the price is 6 x $10.00 = $60.00.
	CompareAtPrice MoneyV2 `json:"compareAtPrice,omitempty"`
	// PerDeliveryPrice: The effective price for a single delivery. For example, for a prepaid subscription plan that includes 6 deliveries at the price of $48.00, the per delivery price is $8.00.
	PerDeliveryPrice MoneyV2 `json:"perDeliveryPrice,omitempty"`
	// Price: The price of the variant when it's purchased with a selling plan For example, for a prepaid subscription plan that includes 6 deliveries of $10.00 granola, where the customer gets 20% off, the price is 6 x $10.00 x 0.80 = $48.00.
	Price MoneyV2 `json:"price,omitempty"`
	// UnitPrice: The resulting price per unit for the variant associated with the selling plan. If the variant isn't sold by quantity or measurement, then this field returns `null`.
	UnitPrice MoneyV2 `json:"unitPrice,omitempty"`
}

/*
StoreAvailability: The availability of a product variant at a particular location.
Local pick-up must be enabled in the  store's shipping settings, otherwise this will return an empty result.
*/
type StoreAvailability struct {
	// Available: Whether or not this product variant is in-stock at this location.
	Available bool `json:"available,omitempty"`
	// Location: The location where this product variant is stocked at.
	Location Location `json:"location,omitempty"`
	// PickUpTime: Returns the estimated amount of time it takes for pickup to be ready (Example: Usually ready in 24 hours).
	PickUpTime string `json:"pickUpTime,omitempty"`
}

// Location: Represents a location where product inventory is held.
type Location struct {
	// Address: The address of the location.
	Address LocationAddress `json:"address,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Name: The name of the location.
	Name string `json:"name,omitempty"`
}

/*
LocationAddress: Represents the address of a location.
*/
type LocationAddress struct {
	// Address1: The first line of the address for the location.
	Address1 string `json:"address1,omitempty"`
	// Address2: The second line of the address for the location.
	Address2 string `json:"address2,omitempty"`
	// City: The city of the location.
	City string `json:"city,omitempty"`
	// Country: The country of the location.
	Country string `json:"country,omitempty"`
	// CountryCode: The country code of the location.
	CountryCode string `json:"countryCode,omitempty"`
	// Formatted: A formatted version of the address for the location.
	Formatted []string `json:"formatted,omitempty"`
	// Latitude: The latitude coordinates of the location.
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude: The longitude coordinates of the location.
	Longitude float64 `json:"longitude,omitempty"`
	// Phone: The phone number of the location.
	Phone string `json:"phone,omitempty"`
	// Province: The province of the location.
	Province string `json:"province,omitempty"`
	/*
	   ProvinceCode: The code for the province, state, or district of the address of the location.
	*/
	ProvinceCode string `json:"provinceCode,omitempty"`
	// ZIP: The ZIP code of the location.
	ZIP string `json:"zip,omitempty"`
}

/*
UnitPriceMeasurement: The measurement used to calculate a unit price for a product variant (e.g. $9.99 / 100ml).
*/
type UnitPriceMeasurement struct {
	// MeasuredType: The type of unit of measurement for the unit price measurement.
	MeasuredType string `json:"measuredType,omitempty"`
	// QuantityUnit: The quantity unit for the unit price measurement.
	QuantityUnit string `json:"quantityUnit,omitempty"`
	// QuantityValue: The quantity value for the unit price measurement.
	QuantityValue float64 `json:"quantityValue,omitempty"`
	// ReferenceUnit: The reference unit for the unit price measurement.
	ReferenceUnit string `json:"referenceUnit,omitempty"`
	// ReferenceValue: The reference value for the unit price measurement.
	ReferenceValue int `json:"referenceValue,omitempty"`
}

// ProductVariantSortKeys: The set of valid sort keys for the ProductVariant query.
type ProductVariantSortKeys string

const (
	ProductVariantSortKeysTitle     ProductVariantSortKeys = "TITLE"
	ProductVariantSortKeysSku       ProductVariantSortKeys = "SKU"
	ProductVariantSortKeysPosition  ProductVariantSortKeys = "POSITION"
	ProductVariantSortKeysId        ProductVariantSortKeys = "ID"
	ProductVariantSortKeysRelevance ProductVariantSortKeys = "RELEVANCE"
)

// Filter: A filter that is supported on the parent field.
type Filter struct {
	// Id: A unique identifier.
	Id string `json:"id,omitempty"`
	// Label: A human-friendly string for this filter.
	Label string `json:"label,omitempty"`
	// Type: An enumeration that denotes the type of data this filter represents.
	Type FilterType `json:"type,omitempty"`
	// Values: The list of values for this filter.
	Values []FilterValue `json:"values,omitempty"`
}

// FilterType: Denotes the type of data this filter group represents.
type FilterType string

const (
	FilterTypeList       FilterType = "LIST"
	FilterTypePriceRange FilterType = "PRICE_RANGE"
)

// FilterValue: A selectable value within a filter.
type FilterValue struct {
	// Count: The number of results that match this filter value.
	Count int `json:"count,omitempty"`
	// Id: A unique identifier.
	Id string `json:"id,omitempty"`
	/*
	   Input: An input object that can be used to filter by this value on the parent field.

	   The value is provided as a helper for building dynamic filtering UI. For example, if you have a list of selected `FilterValue` objects, you can combine their respective `input` values to use in a subsequent query.
	*/
	Input map[string]interface{} `json:"input,omitempty"`
	// Label: A human-friendly string for this filter value.
	Label string `json:"label,omitempty"`
}

// Customer: A customer represents a customer account with the shop. Customer accounts store contact information for the customer, saving logged-in customers the trouble of having to provide it at every checkout.
type Customer struct {
	// AcceptsMarketing: Indicates whether the customer has consented to be sent marketing material via email.
	AcceptsMarketing bool `json:"acceptsMarketing,omitempty"`
	// Addresses: A list of addresses for the customer.
	Addresses Connection[MailingAddress] `json:"addresses,omitempty"`
	// CreatedAt: The date and time when the customer was created.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// DefaultAddress: The customer’s default address.
	DefaultAddress MailingAddress `json:"defaultAddress,omitempty"`
	// DisplayName: The customer’s name, email or phone number.
	DisplayName string `json:"displayName,omitempty"`
	// Email: The customer’s email address.
	Email string `json:"email,omitempty"`
	// FirstName: The customer’s first name.
	FirstName string `json:"firstName,omitempty"`
	// Id: A unique identifier for the customer.
	Id string `json:"id,omitempty"`
	// LastIncompleteCheckout: The customer's most recently updated, incomplete checkout.
	LastIncompleteCheckout Checkout `json:"lastIncompleteCheckout,omitempty"`
	// LastName: The customer’s last name.
	LastName string `json:"lastName,omitempty"`
	// Metafield: Returns a metafield found by namespace and key.
	Metafield Metafield `json:"metafield,omitempty"`
	// Deprecated: Metafields: A paginated list of metafields associated with the resource.
	Metafields Connection[Metafield] `json:"metafields,omitempty"`
	// Orders: The orders associated with the customer.
	Orders Connection[Order] `json:"orders,omitempty"`
	// Phone: The customer’s phone number.
	Phone string `json:"phone,omitempty"`
	/*
	   Tags: A comma separated list of tags that have been added to the customer.
	   Additional access scope required: unauthenticated_read_customer_tags.
	*/
	Tags []string `json:"tags,omitempty"`
	// UpdatedAt: The date and time when the customer information was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// MailingAddress: Represents a mailing address for customers and shipping.
type MailingAddress struct {
	// Address1: The first line of the address. Typically the street address or PO Box number.
	Address1 string `json:"address1,omitempty"`
	/*
	   Address2: The second line of the address. Typically the number of the apartment, suite, or unit.
	*/
	Address2 string `json:"address2,omitempty"`
	/*
	   City: The name of the city, district, village, or town.
	*/
	City string `json:"city,omitempty"`
	/*
	   Company: The name of the customer's company or organization.
	*/
	Company string `json:"company,omitempty"`
	/*
	   Country: The name of the country.
	*/
	Country string `json:"country,omitempty"`
	/*
	   Deprecated: CountryCode: The two-letter code for the country of the address.

	   For example, US.
	*/
	CountryCode string `json:"countryCode,omitempty"`
	/*
	   CountryCodeV2: The two-letter code for the country of the address.

	   For example, US.
	*/
	CountryCodeV2 string `json:"countryCodeV2,omitempty"`
	// FirstName: The first name of the customer.
	FirstName string `json:"firstName,omitempty"`
	// Formatted: A formatted version of the address, customized by the provided arguments.
	Formatted []string `json:"formatted,omitempty"`
	// FormattedArea: A comma-separated list of the values for city, province, and country.
	FormattedArea string `json:"formattedArea,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// LastName: The last name of the customer.
	LastName string `json:"lastName,omitempty"`
	// Latitude: The latitude coordinate of the customer address.
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude: The longitude coordinate of the customer address.
	Longitude float64 `json:"longitude,omitempty"`
	/*
	   Name: The full name of the customer, based on firstName and lastName.
	*/
	Name string `json:"name,omitempty"`
	/*
	   Phone: A unique phone number for the customer.

	   Formatted using E.164 standard. For example, _+16135551111_.
	*/
	Phone string `json:"phone,omitempty"`
	// Province: The region of the address, such as the province, state, or district.
	Province string `json:"province,omitempty"`
	/*
	   ProvinceCode: The two-letter code for the region.

	   For example, ON.
	*/
	ProvinceCode string `json:"provinceCode,omitempty"`
	// ZIP: The zip or postal code of the address.
	ZIP string `json:"zip,omitempty"`
}

// Checkout: A container for all the information required to checkout items and pay.
type Checkout struct {
	// AppliedGiftCards: The gift cards used on the checkout.
	AppliedGiftCards []AppliedGiftCard `json:"appliedGiftCards,omitempty"`
	/*
	   AvailableShippingRates: The available shipping rates for this Checkout.
	   Should only be used when checkout `requiresShipping` is `true` and
	   the shipping address is valid.
	*/
	AvailableShippingRates AvailableShippingRates `json:"availableShippingRates,omitempty"`
	// BuyerIdentity: The identity of the customer associated with the checkout.
	BuyerIdentity CheckoutBuyerIdentity `json:"buyerIdentity,omitempty"`
	// CompletedAt: The date and time when the checkout was completed.
	CompletedAt time.Time `json:"completedAt,omitempty"`
	// CreatedAt: The date and time when the checkout was created.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// CurrencyCode: The currency code for the Checkout.
	CurrencyCode string `json:"currencyCode,omitempty"`
	// CustomAttributes: A list of extra information that is added to the checkout.
	CustomAttributes []Attribute `json:"customAttributes,omitempty"`
	// DiscountApplications: Discounts that have been applied on the checkout.
	DiscountApplications Connection[DiscountApplication] `json:"discountApplications,omitempty"`
	// Email: The email attached to this checkout.
	Email string `json:"email,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// LineItems: A list of line item objects, each one containing information about an item in the checkout.
	LineItems Connection[CheckoutLineItem] `json:"lineItems,omitempty"`
	// LineItemsSubtotalPrice: The sum of all the prices of all the items in the checkout. Duties, taxes, shipping and discounts excluded.
	LineItemsSubtotalPrice MoneyV2 `json:"lineItemsSubtotalPrice,omitempty"`
	// Note: The note associated with the checkout.
	Note string `json:"note,omitempty"`
	// Order: The resulting order from a paid checkout.
	Order Order `json:"order,omitempty"`
	// OrderStatusURL: The Order Status Page for this Checkout, null when checkout is not completed.
	OrderStatusURL string `json:"orderStatusUrl,omitempty"`
	// Deprecated: PaymentDue: The amount left to be paid. This is equal to the cost of the line items, taxes and shipping minus discounts and gift cards.
	PaymentDue string `json:"paymentDue,omitempty"`
	// PaymentDueV2: The amount left to be paid. This is equal to the cost of the line items, duties, taxes and shipping minus discounts and gift cards.
	PaymentDueV2 MoneyV2 `json:"paymentDueV2,omitempty"`
	/*
	   Ready: Whether or not the Checkout is ready and can be completed. Checkouts may
	   have asynchronous operations that can take time to finish. If you want
	   to complete a checkout or ensure all the fields are populated and up to
	   date, polling is required until the value is true.
	*/
	Ready bool `json:"ready,omitempty"`
	// RequiresShipping: States whether or not the fulfillment requires shipping.
	RequiresShipping bool `json:"requiresShipping,omitempty"`
	// ShippingAddress: The shipping address to where the line items will be shipped.
	ShippingAddress MailingAddress `json:"shippingAddress,omitempty"`
	/*
	   ShippingDiscountAllocations: The discounts that have been allocated onto the shipping line by discount applications.
	*/
	ShippingDiscountAllocations []DiscountAllocation `json:"shippingDiscountAllocations,omitempty"`
	// ShippingLine: Once a shipping rate is selected by the customer it is transitioned to a `shipping_line` object.
	ShippingLine ShippingRate `json:"shippingLine,omitempty"`
	// Deprecated: SubtotalPrice: Price of the checkout before shipping and taxes.
	SubtotalPrice string `json:"subtotalPrice,omitempty"`
	// SubtotalPriceV2: Price of the checkout before duties, shipping and taxes.
	SubtotalPriceV2 MoneyV2 `json:"subtotalPriceV2,omitempty"`
	// TaxExempt: Specifies if the Checkout is tax exempt.
	TaxExempt bool `json:"taxExempt,omitempty"`
	// TaxesIncluded: Specifies if taxes are included in the line item and shipping line prices.
	TaxesIncluded bool `json:"taxesIncluded,omitempty"`
	// TotalDuties: The sum of all the duties applied to the line items in the checkout.
	TotalDuties MoneyV2 `json:"totalDuties,omitempty"`
	// Deprecated: TotalPrice: The sum of all the prices of all the items in the checkout, taxes and discounts included.
	TotalPrice string `json:"totalPrice,omitempty"`
	// TotalPriceV2: The sum of all the prices of all the items in the checkout, duties, taxes and discounts included.
	TotalPriceV2 MoneyV2 `json:"totalPriceV2,omitempty"`
	// Deprecated: TotalTax: The sum of all the taxes applied to the line items and shipping lines in the checkout.
	TotalTax string `json:"totalTax,omitempty"`
	// TotalTaxV2: The sum of all the taxes applied to the line items and shipping lines in the checkout.
	TotalTaxV2 MoneyV2 `json:"totalTaxV2,omitempty"`
	// UpdatedAt: The date and time when the checkout was last updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// WebURL: The url pointing to the checkout accessible from the web.
	WebURL string `json:"webUrl,omitempty"`
}

// AppliedGiftCard: Details about the gift card used on the checkout.
type AppliedGiftCard struct {
	// Deprecated: AmountUsed: The amount that was taken from the gift card by applying it.
	AmountUsed string `json:"amountUsed,omitempty"`
	// AmountUsedV2: The amount that was taken from the gift card by applying it.
	AmountUsedV2 MoneyV2 `json:"amountUsedV2,omitempty"`
	// Deprecated: Balance: The amount left on the gift card.
	Balance string `json:"balance,omitempty"`
	// BalanceV2: The amount left on the gift card.
	BalanceV2 MoneyV2 `json:"balanceV2,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// LastCharacters: The last characters of the gift card.
	LastCharacters string `json:"lastCharacters,omitempty"`
	// PresentmentAmountUsed: The amount that was applied to the checkout in its currency.
	PresentmentAmountUsed MoneyV2 `json:"presentmentAmountUsed,omitempty"`
}

// AvailableShippingRates: A collection of available shipping rates for a checkout.
type AvailableShippingRates struct {
	/*
	   Ready: Whether or not the shipping rates are ready.
	   The `shippingRates` field is `null` when this value is `false`.
	   This field should be polled until its value becomes `true`.
	*/
	Ready bool `json:"ready,omitempty"`
	// ShippingRates: The fetched shipping rates. `null` until the `ready` field is `true`.
	ShippingRates []ShippingRate `json:"shippingRates,omitempty"`
}

// ShippingRate: A shipping rate to be applied to a checkout.
type ShippingRate struct {
	// Handle: Human-readable unique identifier for this shipping rate.
	Handle string `json:"handle,omitempty"`
	// Deprecated: Price: Price of this shipping rate.
	Price string `json:"price,omitempty"`
	// PriceV2: Price of this shipping rate.
	PriceV2 MoneyV2 `json:"priceV2,omitempty"`
	// Title: Title of this shipping rate.
	Title string `json:"title,omitempty"`
}

// CheckoutBuyerIdentity: The identity of the customer associated with the checkout.
type CheckoutBuyerIdentity struct {
	// CountryCode: The country code for the checkout. For example, `CA`.
	CountryCode string `json:"countryCode,omitempty"`
}

// Attribute: Represents a generic custom attribute.
type Attribute struct {
	// Key: Key or name of the attribute.
	Key string `json:"key,omitempty"`
	// Value: Value of the attribute.
	Value string `json:"value,omitempty"`
}

/*
DiscountApplication: Discount applications capture the intentions of a discount source at
the time of application.
*/
type DiscountApplication struct {
	// AllocationMethod: The method by which the discount's value is allocated to its entitled items.
	AllocationMethod DiscountApplicationAllocationMethod `json:"allocationMethod,omitempty"`
	// TargetSelection: Which lines of targetType that the discount is allocated over.
	TargetSelection DiscountApplicationTargetSelection `json:"targetSelection,omitempty"`
	// TargetType: The type of line that the discount is applicable towards.
	TargetType DiscountApplicationTargetType `json:"targetType,omitempty"`
	// Value: The value of the discount application.
	Value string `json:"value,omitempty"`
}

// DiscountApplicationAllocationMethod: The method by which the discount's value is allocated onto its entitled lines.
type DiscountApplicationAllocationMethod string

const (
	DiscountApplicationAllocationMethodAcross DiscountApplicationAllocationMethod = "ACROSS"
	DiscountApplicationAllocationMethodEach   DiscountApplicationAllocationMethod = "EACH"
	DiscountApplicationAllocationMethodOne    DiscountApplicationAllocationMethod = "ONE"
)

/*
DiscountApplicationTargetSelection: Which lines on the order that the discount is allocated over, of the type
defined by the Discount Application's target_type.
*/
type DiscountApplicationTargetSelection string

const (
	DiscountApplicationTargetSelectionAll      DiscountApplicationTargetSelection = "ALL"
	DiscountApplicationTargetSelectionEntitled DiscountApplicationTargetSelection = "ENTITLED"
	DiscountApplicationTargetSelectionExplicit DiscountApplicationTargetSelection = "EXPLICIT"
)

/*
DiscountApplicationTargetType: The type of line (i.e. line item or shipping line) on an order that the discount is applicable towards.
*/
type DiscountApplicationTargetType string

const (
	DiscountApplicationTargetTypeLineItem     DiscountApplicationTargetType = "LINE_ITEM"
	DiscountApplicationTargetTypeShippingLine DiscountApplicationTargetType = "SHIPPING_LINE"
)

// PricingPercentageValue: The value of the percentage pricing object.
type PricingPercentageValue struct {
	// Percentage: The percentage value of the object.
	Percentage float64 `json:"percentage,omitempty"`
}

// CheckoutLineItem: A single line item in the checkout, grouped by variant and attributes.
type CheckoutLineItem struct {
	// CustomAttributes: Extra information in the form of an array of Key-Value pairs about the line item.
	CustomAttributes []Attribute `json:"customAttributes,omitempty"`
	// DiscountAllocations: The discounts that have been allocated onto the checkout line item by discount applications.
	DiscountAllocations []DiscountAllocation `json:"discountAllocations,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Quantity: The quantity of the line item.
	Quantity int `json:"quantity,omitempty"`
	// Title: Title of the line item. Defaults to the product's title.
	Title string `json:"title,omitempty"`
	// UnitPrice: Unit price of the line item.
	UnitPrice MoneyV2 `json:"unitPrice,omitempty"`
	// Variant: Product variant of the line item.
	Variant ProductVariant `json:"variant,omitempty"`
}

/*
DiscountAllocation: An amount discounting the line that has been allocated by a discount.
*/
type DiscountAllocation struct {
	// AllocatedAmount: Amount of discount allocated.
	AllocatedAmount MoneyV2 `json:"allocatedAmount,omitempty"`
	// DiscountApplication: The discount this allocated amount originated from.
	DiscountApplication DiscountApplication `json:"discountApplication,omitempty"`
}

// Order: An order is a customer’s completed request to purchase one or more products from a shop. An order is created when a customer completes the checkout process, during which time they provides an email address, billing address and payment information.
type Order struct {
	// CancelReason: The reason for the order's cancellation. Returns `null` if the order wasn't canceled.
	CancelReason OrderCancelReason `json:"cancelReason,omitempty"`
	// CanceledAt: The date and time when the order was canceled. Returns null if the order wasn't canceled.
	CanceledAt time.Time `json:"canceledAt,omitempty"`
	// CurrencyCode: The code of the currency used for the payment.
	CurrencyCode string `json:"currencyCode,omitempty"`
	// CurrentSubtotalPrice: The subtotal of line items and their discounts, excluding line items that have been removed. Does not contain order-level discounts, duties, shipping costs, or shipping discounts. Taxes are not included unless the order is a taxes-included order.
	CurrentSubtotalPrice MoneyV2 `json:"currentSubtotalPrice,omitempty"`
	// CurrentTotalDuties: The total cost of duties for the order, including refunds.
	CurrentTotalDuties MoneyV2 `json:"currentTotalDuties,omitempty"`
	// CurrentTotalPrice: The total amount of the order, including duties, taxes and discounts, minus amounts for line items that have been removed.
	CurrentTotalPrice MoneyV2 `json:"currentTotalPrice,omitempty"`
	// CurrentTotalTax: The total of all taxes applied to the order, excluding taxes for returned line items.
	CurrentTotalTax MoneyV2 `json:"currentTotalTax,omitempty"`
	// CustomerLocale: The locale code in which this specific order happened.
	CustomerLocale string `json:"customerLocale,omitempty"`
	// CustomerURL: The unique URL that the customer can use to access the order.
	CustomerURL string `json:"customerUrl,omitempty"`
	// DiscountApplications: Discounts that have been applied on the order.
	DiscountApplications Connection[DiscountApplication] `json:"discountApplications,omitempty"`
	// Edited: Whether the order has had any edits applied or not.
	Edited bool `json:"edited,omitempty"`
	// Email: The customer's email address.
	Email string `json:"email,omitempty"`
	// FinancialStatus: The financial status of the order.
	FinancialStatus OrderFinancialStatus `json:"financialStatus,omitempty"`
	// FulfillmentStatus: The fulfillment status for the order.
	FulfillmentStatus OrderFulfillmentStatus `json:"fulfillmentStatus,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// LineItems: List of the order’s line items.
	LineItems Connection[OrderLineItem] `json:"lineItems,omitempty"`
	// Metafield: Returns a metafield found by namespace and key.
	Metafield Metafield `json:"metafield,omitempty"`
	// Deprecated: Metafields: A paginated list of metafields associated with the resource.
	Metafields Connection[Metafield] `json:"metafields,omitempty"`
	/*
	   Name: Unique identifier for the order that appears on the order.
	   For example, _#1000_ or _Store1001.
	*/
	Name string `json:"name,omitempty"`
	// OrderNumber: A unique numeric identifier for the order for use by shop owner and customer.
	OrderNumber int `json:"orderNumber,omitempty"`
	// OriginalTotalDuties: The total cost of duties charged at checkout.
	OriginalTotalDuties MoneyV2 `json:"originalTotalDuties,omitempty"`
	// OriginalTotalPrice: The total price of the order before any applied edits.
	OriginalTotalPrice MoneyV2 `json:"originalTotalPrice,omitempty"`
	// Phone: The customer's phone number for receiving SMS notifications.
	Phone string `json:"phone,omitempty"`
	/*
	   ProcessedAt: The date and time when the order was imported.
	   This value can be set to dates in the past when importing from other systems.
	   If no value is provided, it will be auto-generated based on current date and time.
	*/
	ProcessedAt time.Time `json:"processedAt,omitempty"`
	// ShippingAddress: The address to where the order will be shipped.
	ShippingAddress MailingAddress `json:"shippingAddress,omitempty"`
	/*
	   ShippingDiscountAllocations: The discounts that have been allocated onto the shipping line by discount applications.
	*/
	ShippingDiscountAllocations []DiscountAllocation `json:"shippingDiscountAllocations,omitempty"`
	// StatusURL: The unique URL for the order's status page.
	StatusURL string `json:"statusUrl,omitempty"`
	// Deprecated: SubtotalPrice: Price of the order before shipping and taxes.
	SubtotalPrice string `json:"subtotalPrice,omitempty"`
	// SubtotalPriceV2: Price of the order before duties, shipping and taxes.
	SubtotalPriceV2 MoneyV2 `json:"subtotalPriceV2,omitempty"`
	// SuccessfulFulfillments: List of the order’s successful fulfillments.
	SuccessfulFulfillments []Fulfillment `json:"successfulFulfillments,omitempty"`
	// Deprecated: TotalPrice: The sum of all the prices of all the items in the order, taxes and discounts included (must be positive).
	TotalPrice string `json:"totalPrice,omitempty"`
	// TotalPriceV2: The sum of all the prices of all the items in the order, duties, taxes and discounts included (must be positive).
	TotalPriceV2 MoneyV2 `json:"totalPriceV2,omitempty"`
	// Deprecated: TotalRefunded: The total amount that has been refunded.
	TotalRefunded string `json:"totalRefunded,omitempty"`
	// TotalRefundedV2: The total amount that has been refunded.
	TotalRefundedV2 MoneyV2 `json:"totalRefundedV2,omitempty"`
	// Deprecated: TotalShippingPrice: The total cost of shipping.
	TotalShippingPrice string `json:"totalShippingPrice,omitempty"`
	// TotalShippingPriceV2: The total cost of shipping.
	TotalShippingPriceV2 MoneyV2 `json:"totalShippingPriceV2,omitempty"`
	// Deprecated: TotalTax: The total cost of taxes.
	TotalTax string `json:"totalTax,omitempty"`
	// TotalTaxV2: The total cost of taxes.
	TotalTaxV2 MoneyV2 `json:"totalTaxV2,omitempty"`
}

// OrderCancelReason: Represents the reason for the order's cancellation.
type OrderCancelReason string

const (
	OrderCancelReasonCustomer  OrderCancelReason = "CUSTOMER"
	OrderCancelReasonFraud     OrderCancelReason = "FRAUD"
	OrderCancelReasonInventory OrderCancelReason = "INVENTORY"
	OrderCancelReasonDeclined  OrderCancelReason = "DECLINED"
	OrderCancelReasonOther     OrderCancelReason = "OTHER"
)

// OrderFinancialStatus: Represents the order's current financial status.
type OrderFinancialStatus string

const (
	OrderFinancialStatusPending           OrderFinancialStatus = "PENDING"
	OrderFinancialStatusAuthorized        OrderFinancialStatus = "AUTHORIZED"
	OrderFinancialStatusPartiallyPaid     OrderFinancialStatus = "PARTIALLY_PAID"
	OrderFinancialStatusPartiallyRefunded OrderFinancialStatus = "PARTIALLY_REFUNDED"
	OrderFinancialStatusVoided            OrderFinancialStatus = "VOIDED"
	OrderFinancialStatusPaid              OrderFinancialStatus = "PAID"
	OrderFinancialStatusRefunded          OrderFinancialStatus = "REFUNDED"
)

// OrderFulfillmentStatus: Represents the order's aggregated fulfillment status for display purposes.
type OrderFulfillmentStatus string

const (
	OrderFulfillmentStatusUnfulfilled        OrderFulfillmentStatus = "UNFULFILLED"
	OrderFulfillmentStatusPartiallyFulfilled OrderFulfillmentStatus = "PARTIALLY_FULFILLED"
	OrderFulfillmentStatusFulfilled          OrderFulfillmentStatus = "FULFILLED"
	OrderFulfillmentStatusRestocked          OrderFulfillmentStatus = "RESTOCKED"
	OrderFulfillmentStatusPendingFulfillment OrderFulfillmentStatus = "PENDING_FULFILLMENT"
	OrderFulfillmentStatusOpen               OrderFulfillmentStatus = "OPEN"
	OrderFulfillmentStatusInProgress         OrderFulfillmentStatus = "IN_PROGRESS"
	OrderFulfillmentStatusOnHold             OrderFulfillmentStatus = "ON_HOLD"
	OrderFulfillmentStatusScheduled          OrderFulfillmentStatus = "SCHEDULED"
)

// OrderLineItem: Represents a single line in an order. There is one line item for each distinct product variant.
type OrderLineItem struct {
	// CurrentQuantity: The number of entries associated to the line item minus the items that have been removed.
	CurrentQuantity int `json:"currentQuantity,omitempty"`
	// CustomAttributes: List of custom attributes associated to the line item.
	CustomAttributes []Attribute `json:"customAttributes,omitempty"`
	// DiscountAllocations: The discounts that have been allocated onto the order line item by discount applications.
	DiscountAllocations []DiscountAllocation `json:"discountAllocations,omitempty"`
	// DiscountedTotalPrice: The total price of the line item, including discounts, and displayed in the presentment currency.
	DiscountedTotalPrice MoneyV2 `json:"discountedTotalPrice,omitempty"`
	// OriginalTotalPrice: The total price of the line item, not including any discounts. The total price is calculated using the original unit price multiplied by the quantity, and it is displayed in the presentment currency.
	OriginalTotalPrice MoneyV2 `json:"originalTotalPrice,omitempty"`
	// Quantity: The number of products variants associated to the line item.
	Quantity int `json:"quantity,omitempty"`
	// Title: The title of the product combined with title of the variant.
	Title string `json:"title,omitempty"`
	// Variant: The product variant object associated to the line item.
	Variant ProductVariant `json:"variant,omitempty"`
}

// Fulfillment: Represents a single fulfillment in an order.
type Fulfillment struct {
	// FulfillmentLineItems: List of the fulfillment's line items.
	FulfillmentLineItems Connection[FulfillmentLineItem] `json:"fulfillmentLineItems,omitempty"`
	// TrackingCompany: The name of the tracking company.
	TrackingCompany string `json:"trackingCompany,omitempty"`
	/*
	   TrackingInfo: Tracking information associated with the fulfillment,
	   such as the tracking number and tracking URL.
	*/
	TrackingInfo []FulfillmentTrackingInfo `json:"trackingInfo,omitempty"`
}

// FulfillmentLineItem: Represents a single line item in a fulfillment. There is at most one fulfillment line item for each order line item.
type FulfillmentLineItem struct {
	// LineItem: The associated order's line item.
	LineItem OrderLineItem `json:"lineItem,omitempty"`
	// Quantity: The amount fulfilled in this fulfillment.
	Quantity int `json:"quantity,omitempty"`
}

// FulfillmentTrackingInfo: Tracking information associated with the fulfillment.
type FulfillmentTrackingInfo struct {
	// Number: The tracking number of the fulfillment.
	Number string `json:"number,omitempty"`
	// URL: The URL to track the fulfillment.
	URL string `json:"url,omitempty"`
}

// OrderSortKeys: The set of valid sort keys for the Order query.
type OrderSortKeys string

const (
	OrderSortKeysProcessedAt OrderSortKeys = "PROCESSED_AT"
	OrderSortKeysTotalPrice  OrderSortKeys = "TOTAL_PRICE"
	OrderSortKeysId          OrderSortKeys = "ID"
	OrderSortKeysRelevance   OrderSortKeys = "RELEVANCE"
)

// Page: Shopify merchants can create pages to hold static HTML content. Each Page object represents a custom page on the online store.
type Page struct {
	// Body: The description of the page, complete with HTML formatting.
	Body string `json:"body,omitempty"`
	// BodySummary: Summary of the page body.
	BodySummary string `json:"bodySummary,omitempty"`
	// CreatedAt: The timestamp of the page creation.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Handle: A human-friendly unique string for the page automatically generated from its title.
	Handle string `json:"handle,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Metafield: Returns a metafield found by namespace and key.
	Metafield Metafield `json:"metafield,omitempty"`
	// Deprecated: Metafields: A paginated list of metafields associated with the resource.
	Metafields Connection[Metafield] `json:"metafields,omitempty"`
	// OnlineStoreURL: The URL used for viewing the resource on the shop's Online Store. Returns `null` if the resource is currently not published to the Online Store sales channel.
	OnlineStoreURL string `json:"onlineStoreUrl,omitempty"`
	// SEO: The page's SEO information.
	SEO SEO `json:"seo,omitempty"`
	// Title: The title of the page.
	Title string `json:"title,omitempty"`
	// UpdatedAt: The timestamp of the latest page update.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// Shop: Shop represents a collection of the general settings and information about the shop.
type Shop struct {
	// Description: A description of the shop.
	Description string `json:"description,omitempty"`
	// Metafield: Returns a metafield found by namespace and key.
	Metafield Metafield `json:"metafield,omitempty"`
	// Deprecated: Metafields: A paginated list of metafields associated with the resource.
	Metafields Connection[Metafield] `json:"metafields,omitempty"`
	// MoneyFormat: A string representing the way currency is formatted when the currency isn’t specified.
	MoneyFormat string `json:"moneyFormat,omitempty"`
	// Name: The shop’s name.
	Name string `json:"name,omitempty"`
	// PaymentSettings: Settings related to payments.
	PaymentSettings PaymentSettings `json:"paymentSettings,omitempty"`
	// PrimaryDomain: The shop’s primary domain.
	PrimaryDomain Domain `json:"primaryDomain,omitempty"`
	// PrivacyPolicy: The shop’s privacy policy.
	PrivacyPolicy ShopPolicy `json:"privacyPolicy,omitempty"`
	// RefundPolicy: The shop’s refund policy.
	RefundPolicy ShopPolicy `json:"refundPolicy,omitempty"`
	// ShippingPolicy: The shop’s shipping policy.
	ShippingPolicy ShopPolicy `json:"shippingPolicy,omitempty"`
	// ShipsToCountries: Countries that the shop ships to.
	ShipsToCountries []string `json:"shipsToCountries,omitempty"`
	// SubscriptionPolicy: The shop’s subscription policy.
	SubscriptionPolicy ShopPolicyWithDefault `json:"subscriptionPolicy,omitempty"`
	// TermsOfService: The shop’s terms of service.
	TermsOfService ShopPolicy `json:"termsOfService,omitempty"`
}

// PaymentSettings: Settings related to payments.
type PaymentSettings struct {
	// AcceptedCardBrands: List of the card brands which the shop accepts.
	AcceptedCardBrands []CardBrand `json:"acceptedCardBrands,omitempty"`
	// CardVaultURL: The url pointing to the endpoint to vault credit cards.
	CardVaultURL string `json:"cardVaultUrl,omitempty"`
	// CountryCode: The country where the shop is located.
	CountryCode string `json:"countryCode,omitempty"`
	// CurrencyCode: The three-letter code for the shop's primary currency.
	CurrencyCode string `json:"currencyCode,omitempty"`
	// EnabledPresentmentCurrencies: A list of enabled currencies (ISO 4217 format) that the shop accepts. Merchants can enable currencies from their Shopify Payments settings in the Shopify admin.
	EnabledPresentmentCurrencies []string `json:"enabledPresentmentCurrencies,omitempty"`
	// ShopifyPaymentsAccountId: The shop’s Shopify Payments account id.
	ShopifyPaymentsAccountId string `json:"shopifyPaymentsAccountId,omitempty"`
	// SupportedDigitalWallets: List of the digital wallets which the shop supports.
	SupportedDigitalWallets []DigitalWallet `json:"supportedDigitalWallets,omitempty"`
}

// CardBrand: Card brand, such as Visa or Mastercard, which can be used for payments.
type CardBrand string

const (
	CardBrandVisa            CardBrand = "VISA"
	CardBrandMastercard      CardBrand = "MASTERCARD"
	CardBrandDiscover        CardBrand = "DISCOVER"
	CardBrandAmericanExpress CardBrand = "AMERICAN_EXPRESS"
	CardBrandDinersClub      CardBrand = "DINERS_CLUB"
	CardBrandJCB             CardBrand = "JCB"
)

// DigitalWallet: Digital wallet, such as Apple Pay, which can be used for accelerated checkouts.
type DigitalWallet string

const (
	DigitalWalletApplePay   DigitalWallet = "APPLE_PAY"
	DigitalWalletAndroidPay DigitalWallet = "ANDROID_PAY"
	DigitalWalletGooglePay  DigitalWallet = "GOOGLE_PAY"
	DigitalWalletShopifyPay DigitalWallet = "SHOPIFY_PAY"
)

// Domain: Represents a web address.
type Domain struct {
	// Host: The host name of the domain (eg: `example.com`).
	Host string `json:"host,omitempty"`
	// SSLEnabled: Whether SSL is enabled or not.
	SSLEnabled bool `json:"sslEnabled,omitempty"`
	// URL: The URL of the domain (eg: `https://example.com`).
	URL string `json:"url,omitempty"`
}

// ShopPolicy: Policy that a merchant has configured for their store, such as their refund or privacy policy.
type ShopPolicy struct {
	// Body: Policy text, maximum size of 64kb.
	Body string `json:"body,omitempty"`
	// Handle: Policy’s handle.
	Handle string `json:"handle,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Title: Policy’s title.
	Title string `json:"title,omitempty"`
	// URL: Public URL to the policy.
	URL string `json:"url,omitempty"`
}

/*
ShopPolicyWithDefault: A policy for the store that comes with a default value, such as a subscription policy.
If the merchant hasn't configured a policy for their store, then the policy will return the default value.
Otherwise, the policy will return the merchant-configured value.
*/
type ShopPolicyWithDefault struct {
	// Body: The text of the policy. Maximum size: 64KB.
	Body string `json:"body,omitempty"`
	// Handle: The handle of the policy.
	Handle string `json:"handle,omitempty"`
	// Id: The unique identifier of the policy. A default policy doesn't have an ID.
	Id string `json:"id,omitempty"`
	// Title: The title of the policy.
	Title string `json:"title,omitempty"`
	// URL: Public URL to the policy.
	URL string `json:"url,omitempty"`
}

// MediaImage: Represents a Shopify hosted image.
type MediaImage struct {
	// Alt: A word or phrase to share the nature or contents of a media.
	Alt string `json:"alt,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Image: The image for the media.
	Image Image `json:"image,omitempty"`
	// MediaContentType: The media content type.
	MediaContentType MediaContentType `json:"mediaContentType,omitempty"`
	// PreviewImage: The preview image for the media.
	PreviewImage Image `json:"previewImage,omitempty"`
}

// Comment: A comment on an article.
type Comment struct {
	// Author: The comment’s author.
	Author CommentAuthor `json:"author,omitempty"`
	// Content: Stripped content of the comment, single line with HTML tags removed.
	Content string `json:"content,omitempty"`
	// ContentHTML: The content of the comment, complete with HTML formatting.
	ContentHTML string `json:"contentHtml,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
}

// CommentAuthor: The author of a comment.
type CommentAuthor struct {
	// Email: The author's email.
	Email string `json:"email,omitempty"`
	// Name: The author’s name.
	Name string `json:"name,omitempty"`
}

// BlogSortKeys: The set of valid sort keys for the Blog query.
type BlogSortKeys string

const (
	BlogSortKeysHandle    BlogSortKeys = "HANDLE"
	BlogSortKeysTitle     BlogSortKeys = "TITLE"
	BlogSortKeysId        BlogSortKeys = "ID"
	BlogSortKeysRelevance BlogSortKeys = "RELEVANCE"
)

// Cart: A cart represents the merchandise that a buyer intends to purchase, and the estimated cost associated with the cart. To learn how to interact with a cart during a customer's session, refer to [Manage a cart with the Storefront API](https://shopify.dev/custom-storefronts/cart).
type Cart struct {
	// Attributes: The attributes associated with the cart. Attributes are represented as key-value pairs.
	Attributes []Attribute `json:"attributes,omitempty"`
	// BuyerIdentity: Information about the buyer that is interacting with the cart.
	BuyerIdentity CartBuyerIdentity `json:"buyerIdentity,omitempty"`
	// CheckoutURL: The URL of the checkout for the cart.
	CheckoutURL string `json:"checkoutUrl,omitempty"`
	// CreatedAt: The date and time when the cart was created.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// DiscountCodes: The discount codes that have been applied to the cart.
	DiscountCodes []CartDiscountCode `json:"discountCodes,omitempty"`
	// EstimatedCost: The estimated costs that the buyer will pay at checkout. The estimated costs are subject to change and changes will be reflected at checkout. The `estimatedCost` field uses the `buyerIdentity` field to determine [international pricing](https://shopify.dev/custom-storefronts/products/international-pricing#create-a-cart).
	EstimatedCost CartEstimatedCost `json:"estimatedCost,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Lines: A list of lines containing information about the items the customer intends to purchase.
	Lines Connection[CartLine] `json:"lines,omitempty"`
	// Note: A note that is associated with the cart. For example, the note can be a personalized message to the buyer.
	Note string `json:"note,omitempty"`
	// UpdatedAt: The date and time when the cart was updated.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// CartBuyerIdentity: Represents information about the buyer that is interacting with the cart.
type CartBuyerIdentity struct {
	// CountryCode: The country where the buyer is located.
	CountryCode string `json:"countryCode,omitempty"`
	// Customer: The customer account associated with the cart.
	Customer Customer `json:"customer,omitempty"`
	// Email: The email address of the buyer that is interacting with the cart.
	Email string `json:"email,omitempty"`
	// Phone: The phone number of the buyer that is interacting with the cart.
	Phone string `json:"phone,omitempty"`
}

// CartDiscountCode: The discount codes applied to the cart.
type CartDiscountCode struct {
	// Applicable: Whether the discount code is applicable to the cart's current contents.
	Applicable bool `json:"applicable,omitempty"`
	// Code: The code for the discount.
	Code string `json:"code,omitempty"`
}

/*
CartEstimatedCost: The estimated costs that the buyer will pay at checkout.
It uses [`CartBuyerIdentity`](https://shopify.dev/api/storefront/reference/cart/cartbuyeridentity) to determine
[international pricing](https://shopify.dev/custom-storefronts/products/international-pricing#create-a-cart).
*/
type CartEstimatedCost struct {
	// SubtotalAmount: The estimated amount, before taxes and discounts, for the customer to pay at checkout.
	SubtotalAmount MoneyV2 `json:"subtotalAmount,omitempty"`
	// TotalAmount: The estimated total amount for the customer to pay at checkout.
	TotalAmount MoneyV2 `json:"totalAmount,omitempty"`
	// TotalDutyAmount: The estimated duty amount for the customer to pay at checkout.
	TotalDutyAmount MoneyV2 `json:"totalDutyAmount,omitempty"`
	// TotalTaxAmount: The estimated tax amount for the customer to pay at checkout.
	TotalTaxAmount MoneyV2 `json:"totalTaxAmount,omitempty"`
}

// CartLine: Represents information about the merchandise in the cart.
type CartLine struct {
	// Attributes: The attributes associated with the cart line. Attributes are represented as key-value pairs.
	Attributes []Attribute `json:"attributes,omitempty"`
	// DiscountAllocations: The discounts that have been applied to the cart line.
	DiscountAllocations []CartDiscountAllocation `json:"discountAllocations,omitempty"`
	// EstimatedCost: The estimated cost of the merchandise that the buyer will pay for at checkout. The estimated costs are subject to change and changes will be reflected at checkout.
	EstimatedCost CartLineEstimatedCost `json:"estimatedCost,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// Merchandise: The merchandise that the buyer intends to purchase.
	Merchandise string `json:"merchandise,omitempty"`
	// Quantity: The quantity of the merchandise that the customer intends to purchase.
	Quantity int `json:"quantity,omitempty"`
	// SellingPlanAllocation: The selling plan associated with the cart line and the effect that each selling plan has on variants when they're purchased.
	SellingPlanAllocation SellingPlanAllocation `json:"sellingPlanAllocation,omitempty"`
}

// CartDiscountAllocation: The discounts that have been applied to the cart line.
type CartDiscountAllocation struct {
	// DiscountedAmount: The discounted amount that has been applied to the cart line.
	DiscountedAmount MoneyV2 `json:"discountedAmount,omitempty"`
}

// CartLineEstimatedCost: The estimated cost of the merchandise line that the buyer will pay at checkout.
type CartLineEstimatedCost struct {
	// SubtotalAmount: The estimated cost of the merchandise line before discounts.
	SubtotalAmount MoneyV2 `json:"subtotalAmount,omitempty"`
	// TotalAmount: The estimated total cost of the merchandise line.
	TotalAmount MoneyV2 `json:"totalAmount,omitempty"`
}

// CollectionSortKeys: The set of valid sort keys for the Collection query.
type CollectionSortKeys string

const (
	CollectionSortKeysTitle     CollectionSortKeys = "TITLE"
	CollectionSortKeysUpdatedAt CollectionSortKeys = "UPDATED_AT"
	CollectionSortKeysId        CollectionSortKeys = "ID"
	CollectionSortKeysRelevance CollectionSortKeys = "RELEVANCE"
)

// Localization: Information about the localized experiences configured for the shop.
type Localization struct {
	// AvailableCountries: List of countries with enabled localized experiences.
	AvailableCountries []Country `json:"availableCountries,omitempty"`
	// Country: The country of the active localized experience. Use the `@inContext` directive to change this value.
	Country Country `json:"country,omitempty"`
}

// Country: A country.
type Country struct {
	// Currency: The currency of the country.
	Currency Currency `json:"currency,omitempty"`
	// IsoCode: The ISO code of the country.
	IsoCode string `json:"isoCode,omitempty"`
	// Name: The name of the country.
	Name string `json:"name,omitempty"`
	// UnitSystem: The unit system used in the country.
	UnitSystem UnitSystem `json:"unitSystem,omitempty"`
}

// Currency: A currency.
type Currency struct {
	// IsoCode: The ISO code of the currency.
	IsoCode string `json:"isoCode,omitempty"`
	// Name: The name of the currency.
	Name string `json:"name,omitempty"`
	// Symbol: The symbol of the currency.
	Symbol string `json:"symbol,omitempty"`
}

// UnitSystem: Systems of weights and measures.
type UnitSystem string

const (
	UnitSystemImperialSystem UnitSystem = "IMPERIAL_SYSTEM"
	UnitSystemMetricSystem   UnitSystem = "METRIC_SYSTEM"
)

// LocationSortKeys: The set of valid sort keys for the Location query.
type LocationSortKeys string

const (
	LocationSortKeysId       LocationSortKeys = "ID"
	LocationSortKeysName     LocationSortKeys = "NAME"
	LocationSortKeysCity     LocationSortKeys = "CITY"
	LocationSortKeysDistance LocationSortKeys = "DISTANCE"
)

// PageSortKeys: The set of valid sort keys for the Page query.
type PageSortKeys string

const (
	PageSortKeysTitle     PageSortKeys = "TITLE"
	PageSortKeysUpdatedAt PageSortKeys = "UPDATED_AT"
	PageSortKeysId        PageSortKeys = "ID"
	PageSortKeysRelevance PageSortKeys = "RELEVANCE"
)

// ProductSortKeys: The set of valid sort keys for the Product query.
type ProductSortKeys string

const (
	ProductSortKeysTitle       ProductSortKeys = "TITLE"
	ProductSortKeysProductType ProductSortKeys = "PRODUCT_TYPE"
	ProductSortKeysVendor      ProductSortKeys = "VENDOR"
	ProductSortKeysUpdatedAt   ProductSortKeys = "UPDATED_AT"
	ProductSortKeysCreatedAt   ProductSortKeys = "CREATED_AT"
	ProductSortKeysBestSelling ProductSortKeys = "BEST_SELLING"
	ProductSortKeysPrice       ProductSortKeys = "PRICE"
	ProductSortKeysId          ProductSortKeys = "ID"
	ProductSortKeysRelevance   ProductSortKeys = "RELEVANCE"
)

/*
ApiVersion: A version of the API, as defined by [Shopify API versioning](https://shopify.dev/api/usage/versioning).
Versions are commonly referred to by their handle (for example, `2021-10`).
*/
type ApiVersion struct {
	// DisplayName: The human-readable name of the version.
	DisplayName string `json:"displayName,omitempty"`
	// Handle: The unique identifier of an ApiVersion. All supported API versions have a date-based (YYYY-MM) or `unstable` handle.
	Handle string `json:"handle,omitempty"`
	// Supported: Whether the version is actively supported by Shopify. Supported API versions are guaranteed to be stable. Unsupported API versions include unstable, release candidate, and end-of-life versions that are marked as unsupported. For more information, refer to [Versioning](https://shopify.dev/api/usage/versioning).
	Supported bool `json:"supported,omitempty"`
}

// Mutation: The schema’s entry-point for mutations. This acts as the public, top-level API from which all mutation queries must start.
type Mutation struct {
	// CartAttributesUpdate: Updates the attributes on a cart.
	CartAttributesUpdate CartAttributesUpdatePayload `json:"cartAttributesUpdate,omitempty"`
	/*
	   CartBuyerIdentityUpdate: Updates customer information associated with a cart.
	   Buyer identity is used to determine
	   [international pricing](https://shopify.dev/custom-storefronts/products/international-pricing#create-a-checkout)
	   and should match the customer's shipping address.
	*/
	CartBuyerIdentityUpdate CartBuyerIdentityUpdatePayload `json:"cartBuyerIdentityUpdate,omitempty"`
	// CartCreate: Creates a new cart.
	CartCreate CartCreatePayload `json:"cartCreate,omitempty"`
	// CartDiscountCodesUpdate: Updates the discount codes applied to the cart.
	CartDiscountCodesUpdate CartDiscountCodesUpdatePayload `json:"cartDiscountCodesUpdate,omitempty"`
	// CartLinesAdd: Adds a merchandise line to the cart.
	CartLinesAdd CartLinesAddPayload `json:"cartLinesAdd,omitempty"`
	// CartLinesRemove: Removes one or more merchandise lines from the cart.
	CartLinesRemove CartLinesRemovePayload `json:"cartLinesRemove,omitempty"`
	// CartLinesUpdate: Updates one or more merchandise lines on a cart.
	CartLinesUpdate CartLinesUpdatePayload `json:"cartLinesUpdate,omitempty"`
	// CartNoteUpdate: Updates the note on the cart.
	CartNoteUpdate CartNoteUpdatePayload `json:"cartNoteUpdate,omitempty"`
	// Deprecated: CheckoutAttributesUpdate: Updates the attributes of a checkout if `allowPartialAddresses` is `true`.
	CheckoutAttributesUpdate CheckoutAttributesUpdatePayload `json:"checkoutAttributesUpdate,omitempty"`
	// CheckoutAttributesUpdateV2: Updates the attributes of a checkout if `allowPartialAddresses` is `true`.
	CheckoutAttributesUpdateV2 CheckoutAttributesUpdateV2Payload `json:"checkoutAttributesUpdateV2,omitempty"`
	// CheckoutCompleteFree: Completes a checkout without providing payment information. You can use this mutation for free items or items whose purchase price is covered by a gift card.
	CheckoutCompleteFree CheckoutCompleteFreePayload `json:"checkoutCompleteFree,omitempty"`
	// Deprecated: CheckoutCompleteWithCreditCard: Completes a checkout using a credit card token from Shopify's Vault.
	CheckoutCompleteWithCreditCard CheckoutCompleteWithCreditCardPayload `json:"checkoutCompleteWithCreditCard,omitempty"`
	// CheckoutCompleteWithCreditCardV2: Completes a checkout using a credit card token from Shopify's card vault. Before you can complete checkouts using CheckoutCompleteWithCreditCardV2, you need to  [_request payment processing_](https://shopify.dev/apps/channels/getting-started#request-payment-processing).
	CheckoutCompleteWithCreditCardV2 CheckoutCompleteWithCreditCardV2Payload `json:"checkoutCompleteWithCreditCardV2,omitempty"`
	// Deprecated: CheckoutCompleteWithTokenizedPayment: Completes a checkout with a tokenized payment.
	CheckoutCompleteWithTokenizedPayment CheckoutCompleteWithTokenizedPaymentPayload `json:"checkoutCompleteWithTokenizedPayment,omitempty"`
	// Deprecated: CheckoutCompleteWithTokenizedPaymentV2: Completes a checkout with a tokenized payment.
	CheckoutCompleteWithTokenizedPaymentV2 CheckoutCompleteWithTokenizedPaymentV2Payload `json:"checkoutCompleteWithTokenizedPaymentV2,omitempty"`
	// CheckoutCompleteWithTokenizedPaymentV3: Completes a checkout with a tokenized payment.
	CheckoutCompleteWithTokenizedPaymentV3 CheckoutCompleteWithTokenizedPaymentV3Payload `json:"checkoutCompleteWithTokenizedPaymentV3,omitempty"`
	// CheckoutCreate: Creates a new checkout.
	CheckoutCreate CheckoutCreatePayload `json:"checkoutCreate,omitempty"`
	// Deprecated: CheckoutCustomerAssociate: Associates a customer to the checkout.
	CheckoutCustomerAssociate CheckoutCustomerAssociatePayload `json:"checkoutCustomerAssociate,omitempty"`
	// CheckoutCustomerAssociateV2: Associates a customer to the checkout.
	CheckoutCustomerAssociateV2 CheckoutCustomerAssociateV2Payload `json:"checkoutCustomerAssociateV2,omitempty"`
	// Deprecated: CheckoutCustomerDisassociate: Disassociates the current checkout customer from the checkout.
	CheckoutCustomerDisassociate CheckoutCustomerDisassociatePayload `json:"checkoutCustomerDisassociate,omitempty"`
	// CheckoutCustomerDisassociateV2: Disassociates the current checkout customer from the checkout.
	CheckoutCustomerDisassociateV2 CheckoutCustomerDisassociateV2Payload `json:"checkoutCustomerDisassociateV2,omitempty"`
	// Deprecated: CheckoutDiscountCodeApply: Applies a discount to an existing checkout using a discount code.
	CheckoutDiscountCodeApply CheckoutDiscountCodeApplyPayload `json:"checkoutDiscountCodeApply,omitempty"`
	// CheckoutDiscountCodeApplyV2: Applies a discount to an existing checkout using a discount code.
	CheckoutDiscountCodeApplyV2 CheckoutDiscountCodeApplyV2Payload `json:"checkoutDiscountCodeApplyV2,omitempty"`
	// CheckoutDiscountCodeRemove: Removes the applied discount from an existing checkout.
	CheckoutDiscountCodeRemove CheckoutDiscountCodeRemovePayload `json:"checkoutDiscountCodeRemove,omitempty"`
	// Deprecated: CheckoutEmailUpdate: Updates the email on an existing checkout.
	CheckoutEmailUpdate CheckoutEmailUpdatePayload `json:"checkoutEmailUpdate,omitempty"`
	// CheckoutEmailUpdateV2: Updates the email on an existing checkout.
	CheckoutEmailUpdateV2 CheckoutEmailUpdateV2Payload `json:"checkoutEmailUpdateV2,omitempty"`
	// Deprecated: CheckoutGiftCardApply: Applies a gift card to an existing checkout using a gift card code. This will replace all currently applied gift cards.
	CheckoutGiftCardApply CheckoutGiftCardApplyPayload `json:"checkoutGiftCardApply,omitempty"`
	// Deprecated: CheckoutGiftCardRemove: Removes an applied gift card from the checkout.
	CheckoutGiftCardRemove CheckoutGiftCardRemovePayload `json:"checkoutGiftCardRemove,omitempty"`
	// CheckoutGiftCardRemoveV2: Removes an applied gift card from the checkout.
	CheckoutGiftCardRemoveV2 CheckoutGiftCardRemoveV2Payload `json:"checkoutGiftCardRemoveV2,omitempty"`
	// CheckoutGiftCardsAppend: Appends gift cards to an existing checkout.
	CheckoutGiftCardsAppend CheckoutGiftCardsAppendPayload `json:"checkoutGiftCardsAppend,omitempty"`
	// CheckoutLineItemsAdd: Adds a list of line items to a checkout.
	CheckoutLineItemsAdd CheckoutLineItemsAddPayload `json:"checkoutLineItemsAdd,omitempty"`
	// CheckoutLineItemsRemove: Removes line items from an existing checkout.
	CheckoutLineItemsRemove CheckoutLineItemsRemovePayload `json:"checkoutLineItemsRemove,omitempty"`
	// CheckoutLineItemsReplace: Sets a list of line items to a checkout.
	CheckoutLineItemsReplace CheckoutLineItemsReplacePayload `json:"checkoutLineItemsReplace,omitempty"`
	// CheckoutLineItemsUpdate: Updates line items on a checkout.
	CheckoutLineItemsUpdate CheckoutLineItemsUpdatePayload `json:"checkoutLineItemsUpdate,omitempty"`
	// Deprecated: CheckoutShippingAddressUpdate: Updates the shipping address of an existing checkout.
	CheckoutShippingAddressUpdate CheckoutShippingAddressUpdatePayload `json:"checkoutShippingAddressUpdate,omitempty"`
	// CheckoutShippingAddressUpdateV2: Updates the shipping address of an existing checkout.
	CheckoutShippingAddressUpdateV2 CheckoutShippingAddressUpdateV2Payload `json:"checkoutShippingAddressUpdateV2,omitempty"`
	// CheckoutShippingLineUpdate: Updates the shipping lines on an existing checkout.
	CheckoutShippingLineUpdate CheckoutShippingLineUpdatePayload `json:"checkoutShippingLineUpdate,omitempty"`
	/*
	   CustomerAccessTokenCreate: Creates a customer access token.
	   The customer access token is required to modify the customer object in any way.
	*/
	CustomerAccessTokenCreate CustomerAccessTokenCreatePayload `json:"customerAccessTokenCreate,omitempty"`
	/*
	   CustomerAccessTokenCreateWithMultipass: Creates a customer access token using a multipass token instead of email and password.
	   A customer record is created if customer does not exist. If a customer record already
	   exists but the record is disabled, then it's enabled.
	*/
	CustomerAccessTokenCreateWithMultipass CustomerAccessTokenCreateWithMultipassPayload `json:"customerAccessTokenCreateWithMultipass,omitempty"`
	// CustomerAccessTokenDelete: Permanently destroys a customer access token.
	CustomerAccessTokenDelete CustomerAccessTokenDeletePayload `json:"customerAccessTokenDelete,omitempty"`
	/*
	   CustomerAccessTokenRenew: Renews a customer access token.

	   Access token renewal must happen *before* a token expires.
	   If a token has already expired, a new one should be created instead via `customerAccessTokenCreate`.
	*/
	CustomerAccessTokenRenew CustomerAccessTokenRenewPayload `json:"customerAccessTokenRenew,omitempty"`
	// CustomerActivate: Activates a customer.
	CustomerActivate CustomerActivatePayload `json:"customerActivate,omitempty"`
	// CustomerActivateByURL: Activates a customer with the activation url received from `customerCreate`.
	CustomerActivateByURL CustomerActivateByURLPayload `json:"customerActivateByUrl,omitempty"`
	// CustomerAddressCreate: Creates a new address for a customer.
	CustomerAddressCreate CustomerAddressCreatePayload `json:"customerAddressCreate,omitempty"`
	// CustomerAddressDelete: Permanently deletes the address of an existing customer.
	CustomerAddressDelete CustomerAddressDeletePayload `json:"customerAddressDelete,omitempty"`
	// CustomerAddressUpdate: Updates the address of an existing customer.
	CustomerAddressUpdate CustomerAddressUpdatePayload `json:"customerAddressUpdate,omitempty"`
	// CustomerCreate: Creates a new customer.
	CustomerCreate CustomerCreatePayload `json:"customerCreate,omitempty"`
	// CustomerDefaultAddressUpdate: Updates the default address of an existing customer.
	CustomerDefaultAddressUpdate CustomerDefaultAddressUpdatePayload `json:"customerDefaultAddressUpdate,omitempty"`
	// CustomerRecover: Sends a reset password email to the customer, as the first step in the reset password process.
	CustomerRecover CustomerRecoverPayload `json:"customerRecover,omitempty"`
	// CustomerReset: Resets a customer’s password with a token received from `CustomerRecover`.
	CustomerReset CustomerResetPayload `json:"customerReset,omitempty"`
	// CustomerResetByURL: Resets a customer’s password with the reset password url received from `CustomerRecover`.
	CustomerResetByURL CustomerResetByURLPayload `json:"customerResetByUrl,omitempty"`
	// CustomerUpdate: Updates an existing customer.
	CustomerUpdate CustomerUpdatePayload `json:"customerUpdate,omitempty"`
}

// CartAttributesUpdatePayload: Return type for `cartAttributesUpdate` mutation.
type CartAttributesUpdatePayload struct {
	// Cart: The updated cart.
	Cart Cart `json:"cart,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []CartUserError `json:"userErrors,omitempty"`
}

// CartUserError: Represents an error that happens during execution of a cart mutation.
type CartUserError struct {
	// Code: The error code.
	Code CartErrorCode `json:"code,omitempty"`
	// Field: The path to the input field that caused the error.
	Field []string `json:"field,omitempty"`
	// Message: The error message.
	Message string `json:"message,omitempty"`
}

// DisplayableError: Represents an error in the input of a mutation.
type DisplayableError struct {
	// Field: The path to the input field that caused the error.
	Field []string `json:"field,omitempty"`
	// Message: The error message.
	Message string `json:"message,omitempty"`
}

// CartErrorCode: Possible error codes that can be returned by `CartUserError`.
type CartErrorCode string

const (
	CartErrorCodeInvalid                CartErrorCode = "INVALID"
	CartErrorCodeLessThan               CartErrorCode = "LESS_THAN"
	CartErrorCodeInvalidMerchandiseLine CartErrorCode = "INVALID_MERCHANDISE_LINE"
	CartErrorCodeMissingDiscountCode    CartErrorCode = "MISSING_DISCOUNT_CODE"
	CartErrorCodeMissingNote            CartErrorCode = "MISSING_NOTE"
)

// CartBuyerIdentityUpdatePayload: Return type for `cartBuyerIdentityUpdate` mutation.
type CartBuyerIdentityUpdatePayload struct {
	// Cart: The updated cart.
	Cart Cart `json:"cart,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []CartUserError `json:"userErrors,omitempty"`
}

// CartCreatePayload: Return type for `cartCreate` mutation.
type CartCreatePayload struct {
	// Cart: The new cart.
	Cart Cart `json:"cart,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []CartUserError `json:"userErrors,omitempty"`
}

// CartDiscountCodesUpdatePayload: Return type for `cartDiscountCodesUpdate` mutation.
type CartDiscountCodesUpdatePayload struct {
	// Cart: The updated cart.
	Cart Cart `json:"cart,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []CartUserError `json:"userErrors,omitempty"`
}

// CartLinesAddPayload: Return type for `cartLinesAdd` mutation.
type CartLinesAddPayload struct {
	// Cart: The updated cart.
	Cart Cart `json:"cart,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []CartUserError `json:"userErrors,omitempty"`
}

// CartLinesRemovePayload: Return type for `cartLinesRemove` mutation.
type CartLinesRemovePayload struct {
	// Cart: The updated cart.
	Cart Cart `json:"cart,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []CartUserError `json:"userErrors,omitempty"`
}

// CartLinesUpdatePayload: Return type for `cartLinesUpdate` mutation.
type CartLinesUpdatePayload struct {
	// Cart: The updated cart.
	Cart Cart `json:"cart,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []CartUserError `json:"userErrors,omitempty"`
}

// CartNoteUpdatePayload: Return type for `cartNoteUpdate` mutation.
type CartNoteUpdatePayload struct {
	// Cart: The updated cart.
	Cart Cart `json:"cart,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []CartUserError `json:"userErrors,omitempty"`
}

// CheckoutAttributesUpdatePayload: Return type for `checkoutAttributesUpdate` mutation.
type CheckoutAttributesUpdatePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutUserError: Represents an error that happens during execution of a checkout mutation.
type CheckoutUserError struct {
	// Code: The error code.
	Code CheckoutErrorCode `json:"code,omitempty"`
	// Field: The path to the input field that caused the error.
	Field []string `json:"field,omitempty"`
	// Message: The error message.
	Message string `json:"message,omitempty"`
}

// CheckoutErrorCode: Possible error codes that can be returned by `CheckoutUserError`.
type CheckoutErrorCode string

const (
	CheckoutErrorCodeBlank                                            CheckoutErrorCode = "BLANK"
	CheckoutErrorCodeInvalid                                          CheckoutErrorCode = "INVALID"
	CheckoutErrorCodeTooLong                                          CheckoutErrorCode = "TOO_LONG"
	CheckoutErrorCodePresent                                          CheckoutErrorCode = "PRESENT"
	CheckoutErrorCodeLessThan                                         CheckoutErrorCode = "LESS_THAN"
	CheckoutErrorCodeGreaterThanOrEqualTo                             CheckoutErrorCode = "GREATER_THAN_OR_EQUAL_TO"
	CheckoutErrorCodeLessThanOrEqualTo                                CheckoutErrorCode = "LESS_THAN_OR_EQUAL_TO"
	CheckoutErrorCodeAlreadyCompleted                                 CheckoutErrorCode = "ALREADY_COMPLETED"
	CheckoutErrorCodeLocked                                           CheckoutErrorCode = "LOCKED"
	CheckoutErrorCodeNotSupported                                     CheckoutErrorCode = "NOT_SUPPORTED"
	CheckoutErrorCodeBadDomain                                        CheckoutErrorCode = "BAD_DOMAIN"
	CheckoutErrorCodeInvalidForCountry                                CheckoutErrorCode = "INVALID_FOR_COUNTRY"
	CheckoutErrorCodeInvalidForCountryAndProvince                     CheckoutErrorCode = "INVALID_FOR_COUNTRY_AND_PROVINCE"
	CheckoutErrorCodeInvalidStateInCountry                            CheckoutErrorCode = "INVALID_STATE_IN_COUNTRY"
	CheckoutErrorCodeInvalidProvinceInCountry                         CheckoutErrorCode = "INVALID_PROVINCE_IN_COUNTRY"
	CheckoutErrorCodeInvalidRegionInCountry                           CheckoutErrorCode = "INVALID_REGION_IN_COUNTRY"
	CheckoutErrorCodeShippingRateExpired                              CheckoutErrorCode = "SHIPPING_RATE_EXPIRED"
	CheckoutErrorCodeGiftCardUnusable                                 CheckoutErrorCode = "GIFT_CARD_UNUSABLE"
	CheckoutErrorCodeGiftCardDisabled                                 CheckoutErrorCode = "GIFT_CARD_DISABLED"
	CheckoutErrorCodeGiftCardCodeInvalid                              CheckoutErrorCode = "GIFT_CARD_CODE_INVALID"
	CheckoutErrorCodeGiftCardAlreadyApplied                           CheckoutErrorCode = "GIFT_CARD_ALREADY_APPLIED"
	CheckoutErrorCodeGiftCardCurrencyMismatch                         CheckoutErrorCode = "GIFT_CARD_CURRENCY_MISMATCH"
	CheckoutErrorCodeGiftCardExpired                                  CheckoutErrorCode = "GIFT_CARD_EXPIRED"
	CheckoutErrorCodeGiftCardDepleted                                 CheckoutErrorCode = "GIFT_CARD_DEPLETED"
	CheckoutErrorCodeGiftCardNotFound                                 CheckoutErrorCode = "GIFT_CARD_NOT_FOUND"
	CheckoutErrorCodeCartDoesNotMeetDiscountRequirementsNotice        CheckoutErrorCode = "CART_DOES_NOT_MEET_DISCOUNT_REQUIREMENTS_NOTICE"
	CheckoutErrorCodeDiscountExpired                                  CheckoutErrorCode = "DISCOUNT_EXPIRED"
	CheckoutErrorCodeDiscountDisabled                                 CheckoutErrorCode = "DISCOUNT_DISABLED"
	CheckoutErrorCodeDiscountLimitReached                             CheckoutErrorCode = "DISCOUNT_LIMIT_REACHED"
	CheckoutErrorCodeDiscountNotFound                                 CheckoutErrorCode = "DISCOUNT_NOT_FOUND"
	CheckoutErrorCodeCustomerAlreadyUsedOncePerCustomerDiscountNotice CheckoutErrorCode = "CUSTOMER_ALREADY_USED_ONCE_PER_CUSTOMER_DISCOUNT_NOTICE"
	CheckoutErrorCodeEmpty                                            CheckoutErrorCode = "EMPTY"
	CheckoutErrorCodeNotEnoughInStock                                 CheckoutErrorCode = "NOT_ENOUGH_IN_STOCK"
	CheckoutErrorCodeMissingPaymentInput                              CheckoutErrorCode = "MISSING_PAYMENT_INPUT"
	CheckoutErrorCodeTotalPriceMismatch                               CheckoutErrorCode = "TOTAL_PRICE_MISMATCH"
	CheckoutErrorCodeLineItemNotFound                                 CheckoutErrorCode = "LINE_ITEM_NOT_FOUND"
	CheckoutErrorCodeUnableToApply                                    CheckoutErrorCode = "UNABLE_TO_APPLY"
	CheckoutErrorCodeDiscountAlreadyApplied                           CheckoutErrorCode = "DISCOUNT_ALREADY_APPLIED"
	CheckoutErrorCodeThrottledDuringCheckout                          CheckoutErrorCode = "THROTTLED_DURING_CHECKOUT"
	CheckoutErrorCodeExpiredQueueToken                                CheckoutErrorCode = "EXPIRED_QUEUE_TOKEN"
	CheckoutErrorCodeInvalidQueueToken                                CheckoutErrorCode = "INVALID_QUEUE_TOKEN"
	CheckoutErrorCodeInvalidCountryAndCurrency                        CheckoutErrorCode = "INVALID_COUNTRY_AND_CURRENCY"
)

// UserError: Represents an error in the input of a mutation.
type UserError struct {
	// Field: The path to the input field that caused the error.
	Field []string `json:"field,omitempty"`
	// Message: The error message.
	Message string `json:"message,omitempty"`
}

// CheckoutAttributesUpdateV2Payload: Return type for `checkoutAttributesUpdateV2` mutation.
type CheckoutAttributesUpdateV2Payload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutCompleteFreePayload: Return type for `checkoutCompleteFree` mutation.
type CheckoutCompleteFreePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutCompleteWithCreditCardPayload: Return type for `checkoutCompleteWithCreditCard` mutation.
type CheckoutCompleteWithCreditCardPayload struct {
	// Checkout: The checkout on which the payment was applied.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Payment: A representation of the attempted payment.
	Payment Payment `json:"payment,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// Payment: A payment applied to a checkout.
type Payment struct {
	// Deprecated: Amount: The amount of the payment.
	Amount string `json:"amount,omitempty"`
	// AmountV2: The amount of the payment.
	AmountV2 MoneyV2 `json:"amountV2,omitempty"`
	// BillingAddress: The billing address for the payment.
	BillingAddress MailingAddress `json:"billingAddress,omitempty"`
	// Checkout: The checkout to which the payment belongs.
	Checkout Checkout `json:"checkout,omitempty"`
	// CreditCard: The credit card used for the payment in the case of direct payments.
	CreditCard CreditCard `json:"creditCard,omitempty"`
	// ErrorMessage: A message describing a processing error during asynchronous processing.
	ErrorMessage string `json:"errorMessage,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	/*
	   IdempotencyKey: A client-side generated token to identify a payment and perform idempotent operations.
	   For more information, refer to
	   [Idempotent requests](https://shopify.dev/api/usage/idempotent-requests).
	*/
	IdempotencyKey string `json:"idempotencyKey,omitempty"`
	// NextActionURL: The URL where the customer needs to be redirected so they can complete the 3D Secure payment flow.
	NextActionURL string `json:"nextActionUrl,omitempty"`
	// Ready: Whether or not the payment is still processing asynchronously.
	Ready bool `json:"ready,omitempty"`
	// Test: A flag to indicate if the payment is to be done in test mode for gateways that support it.
	Test bool `json:"test,omitempty"`
	// Transaction: The actual transaction recorded by Shopify after having processed the payment with the gateway.
	Transaction Transaction `json:"transaction,omitempty"`
}

// CreditCard: Credit card information used for a payment.
type CreditCard struct {
	// Brand: The brand of the credit card.
	Brand string `json:"brand,omitempty"`
	// ExpiryMonth: The expiry month of the credit card.
	ExpiryMonth int `json:"expiryMonth,omitempty"`
	// ExpiryYear: The expiry year of the credit card.
	ExpiryYear int `json:"expiryYear,omitempty"`
	// FirstDigits: The credit card's BIN number.
	FirstDigits string `json:"firstDigits,omitempty"`
	// FirstName: The first name of the card holder.
	FirstName string `json:"firstName,omitempty"`
	// LastDigits: The last 4 digits of the credit card.
	LastDigits string `json:"lastDigits,omitempty"`
	// LastName: The last name of the card holder.
	LastName string `json:"lastName,omitempty"`
	// MaskedNumber: The masked credit card number with only the last 4 digits displayed.
	MaskedNumber string `json:"maskedNumber,omitempty"`
}

// Transaction: An object representing exchange of money for a product or service.
type Transaction struct {
	// Deprecated: Amount: The amount of money that the transaction was for.
	Amount string `json:"amount,omitempty"`
	// AmountV2: The amount of money that the transaction was for.
	AmountV2 MoneyV2 `json:"amountV2,omitempty"`
	// Kind: The kind of the transaction.
	Kind TransactionKind `json:"kind,omitempty"`
	// Deprecated: Status: The status of the transaction.
	Status TransactionStatus `json:"status,omitempty"`
	// StatusV2: The status of the transaction.
	StatusV2 TransactionStatus `json:"statusV2,omitempty"`
	// Test: Whether the transaction was done in test mode or not.
	Test bool `json:"test,omitempty"`
}

// TransactionKind: The different kinds of order transactions.
type TransactionKind string

const (
	TransactionKindSale             TransactionKind = "SALE"
	TransactionKindCapture          TransactionKind = "CAPTURE"
	TransactionKindAuthorization    TransactionKind = "AUTHORIZATION"
	TransactionKindEmvAuthorization TransactionKind = "EMV_AUTHORIZATION"
	TransactionKindChange           TransactionKind = "CHANGE"
)

// TransactionStatus: Transaction statuses describe the status of a transaction.
type TransactionStatus string

const (
	TransactionStatusPending TransactionStatus = "PENDING"
	TransactionStatusSuccess TransactionStatus = "SUCCESS"
	TransactionStatusFailure TransactionStatus = "FAILURE"
	TransactionStatusError   TransactionStatus = "ERROR"
)

// CheckoutCompleteWithCreditCardV2Payload: Return type for `checkoutCompleteWithCreditCardV2` mutation.
type CheckoutCompleteWithCreditCardV2Payload struct {
	// Checkout: The checkout on which the payment was applied.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Payment: A representation of the attempted payment.
	Payment Payment `json:"payment,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutCompleteWithTokenizedPaymentPayload: Return type for `checkoutCompleteWithTokenizedPayment` mutation.
type CheckoutCompleteWithTokenizedPaymentPayload struct {
	// Checkout: The checkout on which the payment was applied.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Payment: A representation of the attempted payment.
	Payment Payment `json:"payment,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutCompleteWithTokenizedPaymentV2Payload: Return type for `checkoutCompleteWithTokenizedPaymentV2` mutation.
type CheckoutCompleteWithTokenizedPaymentV2Payload struct {
	// Checkout: The checkout on which the payment was applied.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Payment: A representation of the attempted payment.
	Payment Payment `json:"payment,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// PaymentTokenType: The valid values for the types of payment token.
type PaymentTokenType string

const (
	PaymentTokenTypeApplePay         PaymentTokenType = "APPLE_PAY"
	PaymentTokenTypeVault            PaymentTokenType = "VAULT"
	PaymentTokenTypeShopifyPay       PaymentTokenType = "SHOPIFY_PAY"
	PaymentTokenTypeGooglePay        PaymentTokenType = "GOOGLE_PAY"
	PaymentTokenTypeStripeVaultToken PaymentTokenType = "STRIPE_VAULT_TOKEN"
)

// CheckoutCompleteWithTokenizedPaymentV3Payload: Return type for `checkoutCompleteWithTokenizedPaymentV3` mutation.
type CheckoutCompleteWithTokenizedPaymentV3Payload struct {
	// Checkout: The checkout on which the payment was applied.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Payment: A representation of the attempted payment.
	Payment Payment `json:"payment,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutCreatePayload: Return type for `checkoutCreate` mutation.
type CheckoutCreatePayload struct {
	// Checkout: The new checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// QueueToken: The checkout queue token. Available only to selected stores.
	QueueToken string `json:"queueToken,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutCustomerAssociatePayload: Return type for `checkoutCustomerAssociate` mutation.
type CheckoutCustomerAssociatePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// Customer: The associated customer object.
	Customer Customer `json:"customer,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutCustomerAssociateV2Payload: Return type for `checkoutCustomerAssociateV2` mutation.
type CheckoutCustomerAssociateV2Payload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Customer: The associated customer object.
	Customer Customer `json:"customer,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutCustomerDisassociatePayload: Return type for `checkoutCustomerDisassociate` mutation.
type CheckoutCustomerDisassociatePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutCustomerDisassociateV2Payload: Return type for `checkoutCustomerDisassociateV2` mutation.
type CheckoutCustomerDisassociateV2Payload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutDiscountCodeApplyPayload: Return type for `checkoutDiscountCodeApply` mutation.
type CheckoutDiscountCodeApplyPayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutDiscountCodeApplyV2Payload: Return type for `checkoutDiscountCodeApplyV2` mutation.
type CheckoutDiscountCodeApplyV2Payload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutDiscountCodeRemovePayload: Return type for `checkoutDiscountCodeRemove` mutation.
type CheckoutDiscountCodeRemovePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutEmailUpdatePayload: Return type for `checkoutEmailUpdate` mutation.
type CheckoutEmailUpdatePayload struct {
	// Checkout: The checkout object with the updated email.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutEmailUpdateV2Payload: Return type for `checkoutEmailUpdateV2` mutation.
type CheckoutEmailUpdateV2Payload struct {
	// Checkout: The checkout object with the updated email.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutGiftCardApplyPayload: Return type for `checkoutGiftCardApply` mutation.
type CheckoutGiftCardApplyPayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutGiftCardRemovePayload: Return type for `checkoutGiftCardRemove` mutation.
type CheckoutGiftCardRemovePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutGiftCardRemoveV2Payload: Return type for `checkoutGiftCardRemoveV2` mutation.
type CheckoutGiftCardRemoveV2Payload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutGiftCardsAppendPayload: Return type for `checkoutGiftCardsAppend` mutation.
type CheckoutGiftCardsAppendPayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutLineItemsAddPayload: Return type for `checkoutLineItemsAdd` mutation.
type CheckoutLineItemsAddPayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutLineItemsRemovePayload: Return type for `checkoutLineItemsRemove` mutation.
type CheckoutLineItemsRemovePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutLineItemsReplacePayload: Return type for `checkoutLineItemsReplace` mutation.
type CheckoutLineItemsReplacePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []CheckoutUserError `json:"userErrors,omitempty"`
}

// CheckoutLineItemsUpdatePayload: Return type for `checkoutLineItemsUpdate` mutation.
type CheckoutLineItemsUpdatePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutShippingAddressUpdatePayload: Return type for `checkoutShippingAddressUpdate` mutation.
type CheckoutShippingAddressUpdatePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutShippingAddressUpdateV2Payload: Return type for `checkoutShippingAddressUpdateV2` mutation.
type CheckoutShippingAddressUpdateV2Payload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CheckoutShippingLineUpdatePayload: Return type for `checkoutShippingLineUpdate` mutation.
type CheckoutShippingLineUpdatePayload struct {
	// Checkout: The updated checkout object.
	Checkout Checkout `json:"checkout,omitempty"`
	// CheckoutUserErrors: The list of errors that occurred from executing the mutation.
	CheckoutUserErrors []CheckoutUserError `json:"checkoutUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerAccessTokenCreatePayload: Return type for `customerAccessTokenCreate` mutation.
type CustomerAccessTokenCreatePayload struct {
	// CustomerAccessToken: The newly created customer access token object.
	CustomerAccessToken CustomerAccessToken `json:"customerAccessToken,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerAccessToken: A CustomerAccessToken represents the unique token required to make modifications to the customer object.
type CustomerAccessToken struct {
	// AccessToken: The customer’s access token.
	AccessToken string `json:"accessToken,omitempty"`
	// ExpiresAt: The date and time when the customer access token expires.
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
}

// CustomerUserError: Represents an error that happens during execution of a customer mutation.
type CustomerUserError struct {
	// Code: The error code.
	Code CustomerErrorCode `json:"code,omitempty"`
	// Field: The path to the input field that caused the error.
	Field []string `json:"field,omitempty"`
	// Message: The error message.
	Message string `json:"message,omitempty"`
}

// CustomerErrorCode: Possible error codes that can be returned by `CustomerUserError`.
type CustomerErrorCode string

const (
	CustomerErrorCodeBlank                              CustomerErrorCode = "BLANK"
	CustomerErrorCodeInvalid                            CustomerErrorCode = "INVALID"
	CustomerErrorCodeTaken                              CustomerErrorCode = "TAKEN"
	CustomerErrorCodeTooLong                            CustomerErrorCode = "TOO_LONG"
	CustomerErrorCodeTooShort                           CustomerErrorCode = "TOO_SHORT"
	CustomerErrorCodeUnidentifiedCustomer               CustomerErrorCode = "UNIDENTIFIED_CUSTOMER"
	CustomerErrorCodeCustomerDisabled                   CustomerErrorCode = "CUSTOMER_DISABLED"
	CustomerErrorCodePasswordStartsOrEndsWithWhitespace CustomerErrorCode = "PASSWORD_STARTS_OR_ENDS_WITH_WHITESPACE"
	CustomerErrorCodeContainsHTMLTags                   CustomerErrorCode = "CONTAINS_HTML_TAGS"
	CustomerErrorCodeContainsURL                        CustomerErrorCode = "CONTAINS_URL"
	CustomerErrorCodeTokenInvalid                       CustomerErrorCode = "TOKEN_INVALID"
	CustomerErrorCodeAlreadyEnabled                     CustomerErrorCode = "ALREADY_ENABLED"
	CustomerErrorCodeNotFound                           CustomerErrorCode = "NOT_FOUND"
	CustomerErrorCodeBadDomain                          CustomerErrorCode = "BAD_DOMAIN"
	CustomerErrorCodeInvalidMultipassRequest            CustomerErrorCode = "INVALID_MULTIPASS_REQUEST"
)

// CustomerAccessTokenCreateWithMultipassPayload: Return type for `customerAccessTokenCreateWithMultipass` mutation.
type CustomerAccessTokenCreateWithMultipassPayload struct {
	// CustomerAccessToken: An access token object associated with the customer.
	CustomerAccessToken CustomerAccessToken `json:"customerAccessToken,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
}

// CustomerAccessTokenDeletePayload: Return type for `customerAccessTokenDelete` mutation.
type CustomerAccessTokenDeletePayload struct {
	// DeletedAccessToken: The destroyed access token.
	DeletedAccessToken string `json:"deletedAccessToken,omitempty"`
	// DeletedCustomerAccessTokenId: ID of the destroyed customer access token.
	DeletedCustomerAccessTokenId string `json:"deletedCustomerAccessTokenId,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerAccessTokenRenewPayload: Return type for `customerAccessTokenRenew` mutation.
type CustomerAccessTokenRenewPayload struct {
	// CustomerAccessToken: The renewed customer access token object.
	CustomerAccessToken CustomerAccessToken `json:"customerAccessToken,omitempty"`
	// UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerActivatePayload: Return type for `customerActivate` mutation.
type CustomerActivatePayload struct {
	// Customer: The customer object.
	Customer Customer `json:"customer,omitempty"`
	// CustomerAccessToken: A newly created customer access token object for the customer.
	CustomerAccessToken CustomerAccessToken `json:"customerAccessToken,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerActivateByURLPayload: Return type for `customerActivateByUrl` mutation.
type CustomerActivateByURLPayload struct {
	// Customer: The customer that was activated.
	Customer Customer `json:"customer,omitempty"`
	// CustomerAccessToken: A new customer access token for the customer.
	CustomerAccessToken CustomerAccessToken `json:"customerAccessToken,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
}

// CustomerAddressCreatePayload: Return type for `customerAddressCreate` mutation.
type CustomerAddressCreatePayload struct {
	// CustomerAddress: The new customer address object.
	CustomerAddress MailingAddress `json:"customerAddress,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerAddressDeletePayload: Return type for `customerAddressDelete` mutation.
type CustomerAddressDeletePayload struct {
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// DeletedCustomerAddressId: ID of the deleted customer address.
	DeletedCustomerAddressId string `json:"deletedCustomerAddressId,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerAddressUpdatePayload: Return type for `customerAddressUpdate` mutation.
type CustomerAddressUpdatePayload struct {
	// CustomerAddress: The customer’s updated mailing address.
	CustomerAddress MailingAddress `json:"customerAddress,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerCreatePayload: Return type for `customerCreate` mutation.
type CustomerCreatePayload struct {
	// Customer: The created customer object.
	Customer Customer `json:"customer,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerDefaultAddressUpdatePayload: Return type for `customerDefaultAddressUpdate` mutation.
type CustomerDefaultAddressUpdatePayload struct {
	// Customer: The updated customer object.
	Customer Customer `json:"customer,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerRecoverPayload: Return type for `customerRecover` mutation.
type CustomerRecoverPayload struct {
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerResetPayload: Return type for `customerReset` mutation.
type CustomerResetPayload struct {
	// Customer: The customer object which was reset.
	Customer Customer `json:"customer,omitempty"`
	// CustomerAccessToken: A newly created customer access token object for the customer.
	CustomerAccessToken CustomerAccessToken `json:"customerAccessToken,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerResetByURLPayload: Return type for `customerResetByUrl` mutation.
type CustomerResetByURLPayload struct {
	// Customer: The customer object which was reset.
	Customer Customer `json:"customer,omitempty"`
	// CustomerAccessToken: A newly created customer access token object for the customer.
	CustomerAccessToken CustomerAccessToken `json:"customerAccessToken,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

// CustomerUpdatePayload: Return type for `customerUpdate` mutation.
type CustomerUpdatePayload struct {
	// Customer: The updated customer object.
	Customer Customer `json:"customer,omitempty"`
	/*
	   CustomerAccessToken: The newly created customer access token. If the customer's password is updated, all previous access tokens
	   (including the one used to perform this mutation) become invalid, and a new token is generated.
	*/
	CustomerAccessToken CustomerAccessToken `json:"customerAccessToken,omitempty"`
	// CustomerUserErrors: The list of errors that occurred from executing the mutation.
	CustomerUserErrors []CustomerUserError `json:"customerUserErrors,omitempty"`
	// Deprecated: UserErrors: The list of errors that occurred from executing the mutation.
	UserErrors []UserError `json:"userErrors,omitempty"`
}

/*
AutomaticDiscountApplication: Automatic discount applications capture the intentions of a discount that was automatically applied.
*/
type AutomaticDiscountApplication struct {
	// AllocationMethod: The method by which the discount's value is allocated to its entitled items.
	AllocationMethod DiscountApplicationAllocationMethod `json:"allocationMethod,omitempty"`
	// TargetSelection: Which lines of targetType that the discount is allocated over.
	TargetSelection DiscountApplicationTargetSelection `json:"targetSelection,omitempty"`
	// TargetType: The type of line that the discount is applicable towards.
	TargetType DiscountApplicationTargetType `json:"targetType,omitempty"`
	// Title: The title of the application.
	Title string `json:"title,omitempty"`
	// Value: The value of the discount application.
	Value string `json:"value,omitempty"`
}

// CartAutomaticDiscountAllocation: The discounts automatically applied to the cart line based on prerequisites that have been met.
type CartAutomaticDiscountAllocation struct {
	// DiscountedAmount: The discounted amount that has been applied to the cart line.
	DiscountedAmount MoneyV2 `json:"discountedAmount,omitempty"`
	// Title: The title of the allocated discount.
	Title string `json:"title,omitempty"`
}

// CartCodeDiscountAllocation: The discount that has been applied to the cart line using a discount code.
type CartCodeDiscountAllocation struct {
	// Code: The code used to apply the discount.
	Code string `json:"code,omitempty"`
	// DiscountedAmount: The discounted amount that has been applied to the cart line.
	DiscountedAmount MoneyV2 `json:"discountedAmount,omitempty"`
}

/*
DiscountCodeApplication: Discount code applications capture the intentions of a discount code at
the time that it is applied.
*/
type DiscountCodeApplication struct {
	// AllocationMethod: The method by which the discount's value is allocated to its entitled items.
	AllocationMethod DiscountApplicationAllocationMethod `json:"allocationMethod,omitempty"`
	// Applicable: Specifies whether the discount code was applied successfully.
	Applicable bool `json:"applicable,omitempty"`
	// Code: The string identifying the discount code that was used at the time of application.
	Code string `json:"code,omitempty"`
	// TargetSelection: Which lines of targetType that the discount is allocated over.
	TargetSelection DiscountApplicationTargetSelection `json:"targetSelection,omitempty"`
	// TargetType: The type of line that the discount is applicable towards.
	TargetType DiscountApplicationTargetType `json:"targetType,omitempty"`
	// Value: The value of the discount application.
	Value string `json:"value,omitempty"`
}

// ExternalVideo: Represents a video hosted outside of Shopify.
type ExternalVideo struct {
	// Alt: A word or phrase to share the nature or contents of a media.
	Alt string `json:"alt,omitempty"`
	// Deprecated: EmbeddedURL: The URL.
	EmbeddedURL string `json:"embeddedUrl,omitempty"`
	// Host: The host of the external video.
	Host MediaHost `json:"host,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// MediaContentType: The media content type.
	MediaContentType MediaContentType `json:"mediaContentType,omitempty"`
	// PreviewImage: The preview image for the media.
	PreviewImage Image `json:"previewImage,omitempty"`
}

// MediaHost: Host for a Media Resource.
type MediaHost string

const (
	MediaHostYouTube MediaHost = "YOUTUBE"
	MediaHostVimeo   MediaHost = "VIMEO"
)

/*
ManualDiscountApplication: Manual discount applications capture the intentions of a discount that was manually created.
*/
type ManualDiscountApplication struct {
	// AllocationMethod: The method by which the discount's value is allocated to its entitled items.
	AllocationMethod DiscountApplicationAllocationMethod `json:"allocationMethod,omitempty"`
	// Description: The description of the application.
	Description string `json:"description,omitempty"`
	// TargetSelection: Which lines of targetType that the discount is allocated over.
	TargetSelection DiscountApplicationTargetSelection `json:"targetSelection,omitempty"`
	// TargetType: The type of line that the discount is applicable towards.
	TargetType DiscountApplicationTargetType `json:"targetType,omitempty"`
	// Title: The title of the application.
	Title string `json:"title,omitempty"`
	// Value: The value of the discount application.
	Value string `json:"value,omitempty"`
}

// Model3D: Represents a Shopify hosted 3D model.
type Model3D struct {
	// Alt: A word or phrase to share the nature or contents of a media.
	Alt string `json:"alt,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// MediaContentType: The media content type.
	MediaContentType MediaContentType `json:"mediaContentType,omitempty"`
	// PreviewImage: The preview image for the media.
	PreviewImage Image `json:"previewImage,omitempty"`
	// Sources: The sources for a 3d model.
	Sources []Model3DSource `json:"sources,omitempty"`
}

// Model3DSource: Represents a source for a Shopify hosted 3d model.
type Model3DSource struct {
	// Filesize: The filesize of the 3d model.
	Filesize int `json:"filesize,omitempty"`
	// Format: The format of the 3d model.
	Format string `json:"format,omitempty"`
	// MimeType: The MIME type of the 3d model.
	MimeType string `json:"mimeType,omitempty"`
	// URL: The URL of the 3d model.
	URL string `json:"url,omitempty"`
}

/*
ScriptDiscountApplication: Script discount applications capture the intentions of a discount that
was created by a Shopify Script.
*/
type ScriptDiscountApplication struct {
	// AllocationMethod: The method by which the discount's value is allocated to its entitled items.
	AllocationMethod DiscountApplicationAllocationMethod `json:"allocationMethod,omitempty"`
	// TargetSelection: Which lines of targetType that the discount is allocated over.
	TargetSelection DiscountApplicationTargetSelection `json:"targetSelection,omitempty"`
	// TargetType: The type of line that the discount is applicable towards.
	TargetType DiscountApplicationTargetType `json:"targetType,omitempty"`
	// Title: The title of the application as defined by the Script.
	Title string `json:"title,omitempty"`
	// Value: The value of the discount application.
	Value string `json:"value,omitempty"`
}

// Video: Represents a Shopify hosted video.
type Video struct {
	// Alt: A word or phrase to share the nature or contents of a media.
	Alt string `json:"alt,omitempty"`
	// Id: A globally-unique identifier.
	Id string `json:"id,omitempty"`
	// MediaContentType: The media content type.
	MediaContentType MediaContentType `json:"mediaContentType,omitempty"`
	// PreviewImage: The preview image for the media.
	PreviewImage Image `json:"previewImage,omitempty"`
	// Sources: The sources for a video.
	Sources []VideoSource `json:"sources,omitempty"`
}

// VideoSource: Represents a source for a Shopify hosted video.
type VideoSource struct {
	// Format: The format of the video source.
	Format string `json:"format,omitempty"`
	// Height: The height of the video.
	Height int `json:"height,omitempty"`
	// MimeType: The video MIME type.
	MimeType string `json:"mimeType,omitempty"`
	// URL: The URL of the video.
	URL string `json:"url,omitempty"`
	// Width: The width of the video.
	Width int `json:"width,omitempty"`
}
