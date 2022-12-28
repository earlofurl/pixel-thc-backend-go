// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"

	"github.com/gobuffalo/nulls"
	"github.com/google/uuid"
)

type Querier interface {
	// description: Add quantity to a package (can be negative to subtract)
	// arguments: package_id int, quantity float
	AddPackageQuantity(ctx context.Context, arg AddPackageQuantityParams) (Package, error)
	// description: Assign a lab test to a package via junction table
	AssignLabTestToPackage(ctx context.Context, arg AssignLabTestToPackageParams) error
	// description: Assign a source package child package relationship on junction table
	AssignSourcePackageChildPackage(ctx context.Context, arg AssignSourcePackageChildPackageParams) (SourcePackagesChildPackage, error)
	// description: Create a new facility
	CreateFacility(ctx context.Context, arg CreateFacilityParams) (Facility, error)
	// description: Create a new location within a facility
	CreateFacilityLocation(ctx context.Context, arg CreateFacilityLocationParams) (FacilityLocation, error)
	// description: Create a new item
	CreateItem(ctx context.Context, arg CreateItemParams) (Item, error)
	// description: Create a new item type
	CreateItemType(ctx context.Context, arg CreateItemTypeParams) (ItemType, error)
	// description: Create a lab test
	CreateLabTest(ctx context.Context, arg CreateLabTestParams) (LabTest, error)
	// description: Create a new order
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	// description: Create a package
	CreatePackage(ctx context.Context, arg CreatePackageParams) (Package, error)
	// description: Create a package adjustment entry
	CreatePackageAdjEntry(ctx context.Context, arg CreatePackageAdjEntryParams) (PackageAdjEntry, error)
	// description: Create a package adjustment
	CreatePackageAdjustment(ctx context.Context, arg CreatePackageAdjustmentParams) (PackageAdjustment, error)
	// description: Create a package tag
	CreatePackageTag(ctx context.Context, arg CreatePackageTagParams) (PackageTag, error)
	// description: Create a new product category
	CreateProductCategory(ctx context.Context, name string) (ProductCategory, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	// description: Create a strain
	CreateStrain(ctx context.Context, arg CreateStrainParams) (Strain, error)
	// description: Create a new UOM
	CreateUom(ctx context.Context, arg CreateUomParams) (Uom, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	// description: Delete an item by ID
	DeleteItem(ctx context.Context, id int64) error
	// description: Delete an item type by ID
	DeleteItemType(ctx context.Context, id int64) error
	// description: Delete a lab test by ID
	DeleteLabTest(ctx context.Context, id nulls.Int64) error
	// description: Delete a single order by id
	DeleteOrder(ctx context.Context, id int64) error
	// description: Delete a package
	DeletePackage(ctx context.Context, id int64) error
	// description: Delete a package tag by ID
	DeletePackageTag(ctx context.Context, id int64) error
	// description: Delete a product category
	DeleteProductCategory(ctx context.Context, id int64) error
	// description: Delete a strain by ID
	DeleteStrain(ctx context.Context, id int64) error
	// description: Delete a UOM by ID
	DeleteUom(ctx context.Context, id int64) error
	// description: Get a facility by ID
	GetFacility(ctx context.Context, id int64) (Facility, error)
	// description: Get a location within a facility by ID
	GetFacilityLocation(ctx context.Context, id int64) (FacilityLocation, error)
	// description: Get an item by ID
	GetItem(ctx context.Context, id int64) (Item, error)
	// description: Get an item type by id
	GetItemType(ctx context.Context, id int64) (ItemType, error)
	// description: Get a lab test by ID
	GetLabTest(ctx context.Context, id nulls.Int64) (LabTest, error)
	// description: Get a lab test connected to package by package id in lab_tests_packages junction table
	GetLabTestByPackageID(ctx context.Context, packageID int64) (LabTestsPackage, error)
	// description: Get a single order by id
	GetOrder(ctx context.Context, id int64) (Order, error)
	// description: Get a package
	GetPackage(ctx context.Context, id int64) (Package, error)
	// description: Get a package adjustment entry
	GetPackageAdjEntry(ctx context.Context, id int64) (PackageAdjEntry, error)
	// description: Get a package adjustment by id
	GetPackageAdjustment(ctx context.Context, id int64) (PackageAdjustment, error)
	// description: Get a package by tag id
	GetPackageByTagID(ctx context.Context, tagID nulls.Int64) (Package, error)
	// description: Get a package tag by ID
	GetPackageTag(ctx context.Context, id int64) (PackageTag, error)
	// description: Get a package tag by tag number
	GetPackageTagByTagNumber(ctx context.Context, tagNumber string) (PackageTag, error)
	// description: Get a product category by ID
	GetProductCategory(ctx context.Context, id int64) (ProductCategory, error)
	// description: Get a product category by name
	GetProductCategoryByName(ctx context.Context, name string) (ProductCategory, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	// description: Get a strain by ID
	GetStrain(ctx context.Context, id int64) (Strain, error)
	// description: Get a UOM by ID
	GetUom(ctx context.Context, id int64) (Uom, error)
	// description: Get a UOM by name
	GetUomByName(ctx context.Context, name string) (Uom, error)
	GetUser(ctx context.Context, email string) (User, error)
	// description: List all ACTIVE packages with related tag_number, uom, item, lab test, and source package
	ListActivePackages(ctx context.Context) ([]ListActivePackagesRow, error)
	// description: List all facilities
	ListFacilities(ctx context.Context) ([]Facility, error)
	// description: List all locations within facilities
	ListFacilityLocations(ctx context.Context) ([]Facility, error)
	// description: List all locations within a facility
	ListFacilityLocationsByFacility(ctx context.Context, facilityID int64) ([]FacilityLocation, error)
	// description: List all item types
	ListItemTypes(ctx context.Context) ([]ItemType, error)
	// description: List all items
	ListItems(ctx context.Context) ([]ListItemsRow, error)
	// description: List all lab tests
	ListLabTests(ctx context.Context) ([]LabTest, error)
	// description: List all orders
	ListOrders(ctx context.Context) ([]Order, error)
	// description: List package adjustment entries by package id
	ListPackageAdjEntries(ctx context.Context, arg ListPackageAdjEntriesParams) ([]PackageAdjEntry, error)
	// description: List package adjustments
	ListPackageAdjustments(ctx context.Context, arg ListPackageAdjustmentsParams) ([]PackageAdjustment, error)
	// description: List all package tags
	ListPackageTags(ctx context.Context) ([]PackageTag, error)
	// description: List all packages with related tag_number, uom, item, lab test, and source package
	ListPackages(ctx context.Context) ([]ListPackagesRow, error)
	// description: List all product categories
	ListProductCategories(ctx context.Context) ([]ProductCategory, error)
	// description: List all strains
	ListStrains(ctx context.Context) ([]Strain, error)
	// description: List all UOMs
	ListUoms(ctx context.Context) ([]Uom, error)
	// description: Update an item by ID
	UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error)
	// description: Update an item type by ID
	UpdateItemType(ctx context.Context, arg UpdateItemTypeParams) (ItemType, error)
	// description: Update a lab test
	UpdateLabTest(ctx context.Context, arg UpdateLabTestParams) (LabTest, error)
	// description: Update a single order by id
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error)
	// description: Update a package
	UpdatePackage(ctx context.Context, arg UpdatePackageParams) (Package, error)
	// description: Update a package tag
	UpdatePackageTag(ctx context.Context, arg UpdatePackageTagParams) (PackageTag, error)
	// description: Update a product category
	UpdateProductCategory(ctx context.Context, arg UpdateProductCategoryParams) (ProductCategory, error)
	// description: Update a strain
	UpdateStrain(ctx context.Context, arg UpdateStrainParams) (Strain, error)
	// description: Update a UOM by ID
	UpdateUom(ctx context.Context, arg UpdateUomParams) (Uom, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
