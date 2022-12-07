// Code generated by MockGen. DO NOT EDIT.
// Source: pixel-thc-backend-go/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	db "pixel-thc-backend-go/db/sqlc"
	reflect "reflect"

	nulls "github.com/gobuffalo/nulls"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAccountBalance mocks base method.
func (m *MockStore) AddAccountBalance(arg0 context.Context, arg1 db.AddAccountBalanceParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountBalance", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAccountBalance indicates an expected call of AddAccountBalance.
func (mr *MockStoreMockRecorder) AddAccountBalance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountBalance", reflect.TypeOf((*MockStore)(nil).AddAccountBalance), arg0, arg1)
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateEntry mocks base method.
func (m *MockStore) CreateEntry(arg0 context.Context, arg1 db.CreateEntryParams) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEntry indicates an expected call of CreateEntry.
func (mr *MockStoreMockRecorder) CreateEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEntry", reflect.TypeOf((*MockStore)(nil).CreateEntry), arg0, arg1)
}

// CreateItem mocks base method.
func (m *MockStore) CreateItem(arg0 context.Context, arg1 db.CreateItemParams) (db.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", arg0, arg1)
	ret0, _ := ret[0].(db.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockStoreMockRecorder) CreateItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockStore)(nil).CreateItem), arg0, arg1)
}

// CreateItemType mocks base method.
func (m *MockStore) CreateItemType(arg0 context.Context, arg1 db.CreateItemTypeParams) (db.ItemType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItemType", arg0, arg1)
	ret0, _ := ret[0].(db.ItemType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItemType indicates an expected call of CreateItemType.
func (mr *MockStoreMockRecorder) CreateItemType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItemType", reflect.TypeOf((*MockStore)(nil).CreateItemType), arg0, arg1)
}

// CreateLabTest mocks base method.
func (m *MockStore) CreateLabTest(arg0 context.Context, arg1 db.CreateLabTestParams) (db.LabTest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLabTest", arg0, arg1)
	ret0, _ := ret[0].(db.LabTest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLabTest indicates an expected call of CreateLabTest.
func (mr *MockStoreMockRecorder) CreateLabTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLabTest", reflect.TypeOf((*MockStore)(nil).CreateLabTest), arg0, arg1)
}

// CreatePackage mocks base method.
func (m *MockStore) CreatePackage(arg0 context.Context, arg1 db.CreatePackageParams) (db.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePackage", arg0, arg1)
	ret0, _ := ret[0].(db.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePackage indicates an expected call of CreatePackage.
func (mr *MockStoreMockRecorder) CreatePackage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePackage", reflect.TypeOf((*MockStore)(nil).CreatePackage), arg0, arg1)
}

// CreatePackageTag mocks base method.
func (m *MockStore) CreatePackageTag(arg0 context.Context, arg1 db.CreatePackageTagParams) (db.PackageTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePackageTag", arg0, arg1)
	ret0, _ := ret[0].(db.PackageTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePackageTag indicates an expected call of CreatePackageTag.
func (mr *MockStoreMockRecorder) CreatePackageTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePackageTag", reflect.TypeOf((*MockStore)(nil).CreatePackageTag), arg0, arg1)
}

// CreateProductCategory mocks base method.
func (m *MockStore) CreateProductCategory(arg0 context.Context, arg1 string) (db.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductCategory", arg0, arg1)
	ret0, _ := ret[0].(db.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProductCategory indicates an expected call of CreateProductCategory.
func (mr *MockStoreMockRecorder) CreateProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductCategory", reflect.TypeOf((*MockStore)(nil).CreateProductCategory), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateStrain mocks base method.
func (m *MockStore) CreateStrain(arg0 context.Context, arg1 db.CreateStrainParams) (db.Strain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStrain", arg0, arg1)
	ret0, _ := ret[0].(db.Strain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStrain indicates an expected call of CreateStrain.
func (mr *MockStoreMockRecorder) CreateStrain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStrain", reflect.TypeOf((*MockStore)(nil).CreateStrain), arg0, arg1)
}

// CreateTransfer mocks base method.
func (m *MockStore) CreateTransfer(arg0 context.Context, arg1 db.CreateTransferParams) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransfer indicates an expected call of CreateTransfer.
func (mr *MockStoreMockRecorder) CreateTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransfer", reflect.TypeOf((*MockStore)(nil).CreateTransfer), arg0, arg1)
}

// CreateUom mocks base method.
func (m *MockStore) CreateUom(arg0 context.Context, arg1 db.CreateUomParams) (db.Uom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUom", arg0, arg1)
	ret0, _ := ret[0].(db.Uom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUom indicates an expected call of CreateUom.
func (mr *MockStoreMockRecorder) CreateUom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUom", reflect.TypeOf((*MockStore)(nil).CreateUom), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateUserTx mocks base method.
func (m *MockStore) CreateUserTx(arg0 context.Context, arg1 db.CreateUserTxParams) (db.CreateUserTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserTx", arg0, arg1)
	ret0, _ := ret[0].(db.CreateUserTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserTx indicates an expected call of CreateUserTx.
func (mr *MockStoreMockRecorder) CreateUserTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserTx", reflect.TypeOf((*MockStore)(nil).CreateUserTx), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// DeleteItem mocks base method.
func (m *MockStore) DeleteItem(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockStoreMockRecorder) DeleteItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockStore)(nil).DeleteItem), arg0, arg1)
}

// DeleteItemType mocks base method.
func (m *MockStore) DeleteItemType(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItemType", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItemType indicates an expected call of DeleteItemType.
func (mr *MockStoreMockRecorder) DeleteItemType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItemType", reflect.TypeOf((*MockStore)(nil).DeleteItemType), arg0, arg1)
}

// DeleteLabTest mocks base method.
func (m *MockStore) DeleteLabTest(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLabTest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLabTest indicates an expected call of DeleteLabTest.
func (mr *MockStoreMockRecorder) DeleteLabTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLabTest", reflect.TypeOf((*MockStore)(nil).DeleteLabTest), arg0, arg1)
}

// DeletePackage mocks base method.
func (m *MockStore) DeletePackage(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePackage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePackage indicates an expected call of DeletePackage.
func (mr *MockStoreMockRecorder) DeletePackage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePackage", reflect.TypeOf((*MockStore)(nil).DeletePackage), arg0, arg1)
}

// DeletePackageTag mocks base method.
func (m *MockStore) DeletePackageTag(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePackageTag", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePackageTag indicates an expected call of DeletePackageTag.
func (mr *MockStoreMockRecorder) DeletePackageTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePackageTag", reflect.TypeOf((*MockStore)(nil).DeletePackageTag), arg0, arg1)
}

// DeleteProductCategory mocks base method.
func (m *MockStore) DeleteProductCategory(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductCategory", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductCategory indicates an expected call of DeleteProductCategory.
func (mr *MockStoreMockRecorder) DeleteProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductCategory", reflect.TypeOf((*MockStore)(nil).DeleteProductCategory), arg0, arg1)
}

// DeleteStrain mocks base method.
func (m *MockStore) DeleteStrain(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStrain", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStrain indicates an expected call of DeleteStrain.
func (mr *MockStoreMockRecorder) DeleteStrain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStrain", reflect.TypeOf((*MockStore)(nil).DeleteStrain), arg0, arg1)
}

// DeleteUom mocks base method.
func (m *MockStore) DeleteUom(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUom indicates an expected call of DeleteUom.
func (mr *MockStoreMockRecorder) DeleteUom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUom", reflect.TypeOf((*MockStore)(nil).DeleteUom), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), arg0, arg1)
}

// GetAccountForUpdate mocks base method.
func (m *MockStore) GetAccountForUpdate(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountForUpdate indicates an expected call of GetAccountForUpdate.
func (mr *MockStoreMockRecorder) GetAccountForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountForUpdate", reflect.TypeOf((*MockStore)(nil).GetAccountForUpdate), arg0, arg1)
}

// GetEntry mocks base method.
func (m *MockStore) GetEntry(arg0 context.Context, arg1 int64) (db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEntry", arg0, arg1)
	ret0, _ := ret[0].(db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEntry indicates an expected call of GetEntry.
func (mr *MockStoreMockRecorder) GetEntry(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEntry", reflect.TypeOf((*MockStore)(nil).GetEntry), arg0, arg1)
}

// GetItem mocks base method.
func (m *MockStore) GetItem(arg0 context.Context, arg1 int64) (db.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", arg0, arg1)
	ret0, _ := ret[0].(db.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockStoreMockRecorder) GetItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockStore)(nil).GetItem), arg0, arg1)
}

// GetItemType mocks base method.
func (m *MockStore) GetItemType(arg0 context.Context, arg1 int64) (db.ItemType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemType", arg0, arg1)
	ret0, _ := ret[0].(db.ItemType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItemType indicates an expected call of GetItemType.
func (mr *MockStoreMockRecorder) GetItemType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemType", reflect.TypeOf((*MockStore)(nil).GetItemType), arg0, arg1)
}

// GetLabTest mocks base method.
func (m *MockStore) GetLabTest(arg0 context.Context, arg1 int64) (db.LabTest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLabTest", arg0, arg1)
	ret0, _ := ret[0].(db.LabTest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLabTest indicates an expected call of GetLabTest.
func (mr *MockStoreMockRecorder) GetLabTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabTest", reflect.TypeOf((*MockStore)(nil).GetLabTest), arg0, arg1)
}

// GetPackage mocks base method.
func (m *MockStore) GetPackage(arg0 context.Context, arg1 int64) (db.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPackage", arg0, arg1)
	ret0, _ := ret[0].(db.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPackage indicates an expected call of GetPackage.
func (mr *MockStoreMockRecorder) GetPackage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPackage", reflect.TypeOf((*MockStore)(nil).GetPackage), arg0, arg1)
}

// GetPackageByTagID mocks base method.
func (m *MockStore) GetPackageByTagID(arg0 context.Context, arg1 nulls.Int64) (db.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPackageByTagID", arg0, arg1)
	ret0, _ := ret[0].(db.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPackageByTagID indicates an expected call of GetPackageByTagID.
func (mr *MockStoreMockRecorder) GetPackageByTagID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPackageByTagID", reflect.TypeOf((*MockStore)(nil).GetPackageByTagID), arg0, arg1)
}

// GetPackageTag mocks base method.
func (m *MockStore) GetPackageTag(arg0 context.Context, arg1 int64) (db.PackageTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPackageTag", arg0, arg1)
	ret0, _ := ret[0].(db.PackageTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPackageTag indicates an expected call of GetPackageTag.
func (mr *MockStoreMockRecorder) GetPackageTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPackageTag", reflect.TypeOf((*MockStore)(nil).GetPackageTag), arg0, arg1)
}

// GetPackageTagByTagNumber mocks base method.
func (m *MockStore) GetPackageTagByTagNumber(arg0 context.Context, arg1 string) (db.PackageTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPackageTagByTagNumber", arg0, arg1)
	ret0, _ := ret[0].(db.PackageTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPackageTagByTagNumber indicates an expected call of GetPackageTagByTagNumber.
func (mr *MockStoreMockRecorder) GetPackageTagByTagNumber(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPackageTagByTagNumber", reflect.TypeOf((*MockStore)(nil).GetPackageTagByTagNumber), arg0, arg1)
}

// GetProductCategory mocks base method.
func (m *MockStore) GetProductCategory(arg0 context.Context, arg1 int64) (db.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductCategory", arg0, arg1)
	ret0, _ := ret[0].(db.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductCategory indicates an expected call of GetProductCategory.
func (mr *MockStoreMockRecorder) GetProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductCategory", reflect.TypeOf((*MockStore)(nil).GetProductCategory), arg0, arg1)
}

// GetProductCategoryByName mocks base method.
func (m *MockStore) GetProductCategoryByName(arg0 context.Context, arg1 string) (db.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductCategoryByName", arg0, arg1)
	ret0, _ := ret[0].(db.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductCategoryByName indicates an expected call of GetProductCategoryByName.
func (mr *MockStoreMockRecorder) GetProductCategoryByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductCategoryByName", reflect.TypeOf((*MockStore)(nil).GetProductCategoryByName), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetStrain mocks base method.
func (m *MockStore) GetStrain(arg0 context.Context, arg1 int64) (db.Strain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStrain", arg0, arg1)
	ret0, _ := ret[0].(db.Strain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStrain indicates an expected call of GetStrain.
func (mr *MockStoreMockRecorder) GetStrain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStrain", reflect.TypeOf((*MockStore)(nil).GetStrain), arg0, arg1)
}

// GetTransfer mocks base method.
func (m *MockStore) GetTransfer(arg0 context.Context, arg1 int64) (db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransfer", arg0, arg1)
	ret0, _ := ret[0].(db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransfer indicates an expected call of GetTransfer.
func (mr *MockStoreMockRecorder) GetTransfer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransfer", reflect.TypeOf((*MockStore)(nil).GetTransfer), arg0, arg1)
}

// GetUom mocks base method.
func (m *MockStore) GetUom(arg0 context.Context, arg1 int64) (db.Uom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUom", arg0, arg1)
	ret0, _ := ret[0].(db.Uom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUom indicates an expected call of GetUom.
func (mr *MockStoreMockRecorder) GetUom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUom", reflect.TypeOf((*MockStore)(nil).GetUom), arg0, arg1)
}

// GetUomByName mocks base method.
func (m *MockStore) GetUomByName(arg0 context.Context, arg1 string) (db.Uom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUomByName", arg0, arg1)
	ret0, _ := ret[0].(db.Uom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUomByName indicates an expected call of GetUomByName.
func (mr *MockStoreMockRecorder) GetUomByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUomByName", reflect.TypeOf((*MockStore)(nil).GetUomByName), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), arg0, arg1)
}

// ListEntries mocks base method.
func (m *MockStore) ListEntries(arg0 context.Context, arg1 db.ListEntriesParams) ([]db.Entry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntries", arg0, arg1)
	ret0, _ := ret[0].([]db.Entry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntries indicates an expected call of ListEntries.
func (mr *MockStoreMockRecorder) ListEntries(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntries", reflect.TypeOf((*MockStore)(nil).ListEntries), arg0, arg1)
}

// ListItemTypes mocks base method.
func (m *MockStore) ListItemTypes(arg0 context.Context) ([]db.ItemType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListItemTypes", arg0)
	ret0, _ := ret[0].([]db.ItemType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListItemTypes indicates an expected call of ListItemTypes.
func (mr *MockStoreMockRecorder) ListItemTypes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItemTypes", reflect.TypeOf((*MockStore)(nil).ListItemTypes), arg0)
}

// ListItems mocks base method.
func (m *MockStore) ListItems(arg0 context.Context) ([]db.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListItems", arg0)
	ret0, _ := ret[0].([]db.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListItems indicates an expected call of ListItems.
func (mr *MockStoreMockRecorder) ListItems(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListItems", reflect.TypeOf((*MockStore)(nil).ListItems), arg0)
}

// ListLabTests mocks base method.
func (m *MockStore) ListLabTests(arg0 context.Context) ([]db.LabTest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLabTests", arg0)
	ret0, _ := ret[0].([]db.LabTest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLabTests indicates an expected call of ListLabTests.
func (mr *MockStoreMockRecorder) ListLabTests(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLabTests", reflect.TypeOf((*MockStore)(nil).ListLabTests), arg0)
}

// ListPackageTags mocks base method.
func (m *MockStore) ListPackageTags(arg0 context.Context) ([]db.PackageTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPackageTags", arg0)
	ret0, _ := ret[0].([]db.PackageTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPackageTags indicates an expected call of ListPackageTags.
func (mr *MockStoreMockRecorder) ListPackageTags(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPackageTags", reflect.TypeOf((*MockStore)(nil).ListPackageTags), arg0)
}

// ListPackages mocks base method.
func (m *MockStore) ListPackages(arg0 context.Context) ([]db.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPackages", arg0)
	ret0, _ := ret[0].([]db.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPackages indicates an expected call of ListPackages.
func (mr *MockStoreMockRecorder) ListPackages(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPackages", reflect.TypeOf((*MockStore)(nil).ListPackages), arg0)
}

// ListProductCategories mocks base method.
func (m *MockStore) ListProductCategories(arg0 context.Context) ([]db.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductCategories", arg0)
	ret0, _ := ret[0].([]db.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProductCategories indicates an expected call of ListProductCategories.
func (mr *MockStoreMockRecorder) ListProductCategories(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductCategories", reflect.TypeOf((*MockStore)(nil).ListProductCategories), arg0)
}

// ListStrains mocks base method.
func (m *MockStore) ListStrains(arg0 context.Context) ([]db.Strain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStrains", arg0)
	ret0, _ := ret[0].([]db.Strain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStrains indicates an expected call of ListStrains.
func (mr *MockStoreMockRecorder) ListStrains(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStrains", reflect.TypeOf((*MockStore)(nil).ListStrains), arg0)
}

// ListTransfers mocks base method.
func (m *MockStore) ListTransfers(arg0 context.Context, arg1 db.ListTransfersParams) ([]db.Transfer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTransfers", arg0, arg1)
	ret0, _ := ret[0].([]db.Transfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTransfers indicates an expected call of ListTransfers.
func (mr *MockStoreMockRecorder) ListTransfers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTransfers", reflect.TypeOf((*MockStore)(nil).ListTransfers), arg0, arg1)
}

// ListUoms mocks base method.
func (m *MockStore) ListUoms(arg0 context.Context) ([]db.Uom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUoms", arg0)
	ret0, _ := ret[0].([]db.Uom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUoms indicates an expected call of ListUoms.
func (mr *MockStoreMockRecorder) ListUoms(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUoms", reflect.TypeOf((*MockStore)(nil).ListUoms), arg0)
}

// TransferTx mocks base method.
func (m *MockStore) TransferTx(arg0 context.Context, arg1 db.TransferTxParams) (db.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(db.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockStoreMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockStore)(nil).TransferTx), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(arg0 context.Context, arg1 db.UpdateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), arg0, arg1)
}

// UpdateItem mocks base method.
func (m *MockStore) UpdateItem(arg0 context.Context, arg1 db.UpdateItemParams) (db.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", arg0, arg1)
	ret0, _ := ret[0].(db.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockStoreMockRecorder) UpdateItem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockStore)(nil).UpdateItem), arg0, arg1)
}

// UpdateItemType mocks base method.
func (m *MockStore) UpdateItemType(arg0 context.Context, arg1 db.UpdateItemTypeParams) (db.ItemType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItemType", arg0, arg1)
	ret0, _ := ret[0].(db.ItemType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItemType indicates an expected call of UpdateItemType.
func (mr *MockStoreMockRecorder) UpdateItemType(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItemType", reflect.TypeOf((*MockStore)(nil).UpdateItemType), arg0, arg1)
}

// UpdateLabTest mocks base method.
func (m *MockStore) UpdateLabTest(arg0 context.Context, arg1 db.UpdateLabTestParams) (db.LabTest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLabTest", arg0, arg1)
	ret0, _ := ret[0].(db.LabTest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLabTest indicates an expected call of UpdateLabTest.
func (mr *MockStoreMockRecorder) UpdateLabTest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLabTest", reflect.TypeOf((*MockStore)(nil).UpdateLabTest), arg0, arg1)
}

// UpdatePackage mocks base method.
func (m *MockStore) UpdatePackage(arg0 context.Context, arg1 db.UpdatePackageParams) (db.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackage", arg0, arg1)
	ret0, _ := ret[0].(db.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePackage indicates an expected call of UpdatePackage.
func (mr *MockStoreMockRecorder) UpdatePackage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackage", reflect.TypeOf((*MockStore)(nil).UpdatePackage), arg0, arg1)
}

// UpdatePackageTag mocks base method.
func (m *MockStore) UpdatePackageTag(arg0 context.Context, arg1 db.UpdatePackageTagParams) (db.PackageTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackageTag", arg0, arg1)
	ret0, _ := ret[0].(db.PackageTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePackageTag indicates an expected call of UpdatePackageTag.
func (mr *MockStoreMockRecorder) UpdatePackageTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackageTag", reflect.TypeOf((*MockStore)(nil).UpdatePackageTag), arg0, arg1)
}

// UpdateProductCategory mocks base method.
func (m *MockStore) UpdateProductCategory(arg0 context.Context, arg1 db.UpdateProductCategoryParams) (db.ProductCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductCategory", arg0, arg1)
	ret0, _ := ret[0].(db.ProductCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductCategory indicates an expected call of UpdateProductCategory.
func (mr *MockStoreMockRecorder) UpdateProductCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductCategory", reflect.TypeOf((*MockStore)(nil).UpdateProductCategory), arg0, arg1)
}

// UpdateStrain mocks base method.
func (m *MockStore) UpdateStrain(arg0 context.Context, arg1 db.UpdateStrainParams) (db.Strain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStrain", arg0, arg1)
	ret0, _ := ret[0].(db.Strain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStrain indicates an expected call of UpdateStrain.
func (mr *MockStoreMockRecorder) UpdateStrain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStrain", reflect.TypeOf((*MockStore)(nil).UpdateStrain), arg0, arg1)
}

// UpdateUom mocks base method.
func (m *MockStore) UpdateUom(arg0 context.Context, arg1 db.UpdateUomParams) (db.Uom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUom", arg0, arg1)
	ret0, _ := ret[0].(db.Uom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUom indicates an expected call of UpdateUom.
func (mr *MockStoreMockRecorder) UpdateUom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUom", reflect.TypeOf((*MockStore)(nil).UpdateUom), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
