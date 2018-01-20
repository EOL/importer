package main

import (
	"fmt"
	"math"
	"time"

	h "github.com/eol/importer/harvdb"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type webNode struct {
	ID                 int
	resource_id        int
	page_id            int
	rank_id            int
	parent_id          int
	scientific_name    string
	canonical_form     string
	resource_pk        string
	source_url         string
	is_hidden          bool
	in_unmapped_area   bool
	children_count     int
	created_at         time.Time
	updated_at         time.Time
	has_breadcrumb     bool
	parent_resource_pk string
	landmark           int
}

func (n webNode) TableName() string {
	return "nodes"
}

func main() {
	harvDB, err := gorm.Open("mysql", "root:@/harvester_development?charset=utf8&parseTime=True&loc=Local")
	defer harvDB.Close()
	if err != nil {
		fmt.Println(err)
	}
	webDB, err := gorm.Open("mysql", "root:@/eol_development?charset=utf8&parseTime=True&loc=Local")
	defer webDB.Close()
	if err != nil {
		fmt.Println(err)
	}
	resource := h.Resource{}
	harvDB.Where("id = ?", 16).First(&resource)
	fmt.Println(resource.ID, "->", resource.Name)
	var nodes []h.Node

	tStart := time.Now()
	var count int64
	harvDB.Model(&nodes).Where("resource_id = ?", resource.ID).Count(&count)
	nodes, err = loadNodes(harvDB, count, resource)
	delta := time.Now().Sub(tStart)
	fmt.Println("Found and loaded", len(nodes), "Nodes. First 10:")
	showNames(nodes[:9])
	fmt.Println(delta.Seconds())
}

func loadNodes(db *gorm.DB, count int64, resource h.Resource) (nodeSlice []h.Node, err error) {
	nodesAll := make([]h.Node, 0, count)
	var nodeBuff []h.Node
	const perPage = 1000
	pages := PageCount(count, perPage)
	fmt.Println("PageCount:", pages)
	for currentPage := int64(1); currentPage <= pages; currentPage++ {
		offset := perPage * (currentPage - 1)
		fmt.Println("Reading page", currentPage, "of", pages, "with offset of", offset)
		db.Limit(perPage).Offset(offset).Model(&resource).Preload("ScientificName").Related(&nodeBuff)
		// TODO: err-checking here. ...There is none in the documentation. :S
		nodesAll = append(nodesAll, nodeBuff...)
		showNames(nodesAll[offset+1 : offset+2])
	}
	return nodesAll[:], nil
}

func PageCount(numRows, perPage int64) int64 {
	result := float64(numRows) / float64(perPage)
	floor := math.Floor(result)
	if result-floor > 0 {
		return int64(floor + 1)
	}
	return int64(floor)
}

func showNames(nodes []h.Node) {
	for i := range nodes[:] {
		fmt.Println("Node", nodes[i].ID, "-> '"+nodes[i].ScientificName.Verbatim+"' ("+nodes[i].Canonical+")")
	}
	return
}
