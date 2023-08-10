package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
	"github.com/stevan-iskandar/learn-beego/models/book"
)

// DO NOT MODIFY
type CreateBookTable_20230810_114311 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateBookTable_20230810_114311{}
	m.Created = "20230810_114311"

	migration.Register("CreateBookTable_20230810_114311", m)
}

// Run the migrations
func (m *CreateBookTable_20230810_114311) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.CreateTable(book.TABLE, "InnoDB", "utf8mb4")
	m.PriCol("id").SetDataType("int").SetAuto(true).SetUnsigned(true)
	m.NewCol("name").SetDataType("varchar(255)").SetNullable(true)
	m.SQL(m.GetSQL())
}

// Reverse the migrations
func (m *CreateBookTable_20230810_114311) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
