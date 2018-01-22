package main

import (
	"fmt"
	"math" // For Floor
	// "strings"  // For Join
	"time" // For Now (heh heh)

	h "github.com/eol/importer/harvdb"
	w "github.com/eol/importer/webdb"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/btfak/sqlext"
)

func main() {
	harvDB, err := gorm.Open("mysql", "root:@/harvester_development?charset=utf8&parseTime=True&loc=Local")
	defer harvDB.Close()
	if err != nil {
		fmt.Println(err)
	}
	webDB, err := gorm.Open("mysql", "root:@/eol_development?charset=utf8&parseTime=True&loc=Local")
	webDB.SingularTable(false) // I should not need this, but Node wasn't working. :S
	defer webDB.Close()
	if err != nil {
		fmt.Println(err)
	}
	resource := h.Resource{}
	harvDB.Where("id = ?", 1).First(&resource)
	fmt.Println(resource.ID, "->", resource.Name)
	var nodes []h.Node

	tStart := time.Now()
	var count int64
	harvDB.Model(&nodes).Where("resource_id = ?", resource.ID).Count(&count)
	nodes, err = loadNodes(harvDB, count, resource)
	delta := time.Now().Sub(tStart)
	fmt.Println("Found and loaded", len(nodes), "Nodes in", delta.Seconds(), "... Last 10:")
	showNames(nodes[len(nodes)-10:])
	tStart = time.Now()
	webNodes := setNodes(&nodes)
	delta = time.Now().Sub(tStart)
	fmt.Println("Converted", len(webNodes), "Nodes. Took ", delta.Seconds())
	tStart = time.Now()
	err = saveNodes(webDB, webNodes)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	delta = time.Now().Sub(tStart)
	fmt.Println("Wrote them to the DB,", delta.Seconds())
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

func saveNodes(db *gorm.DB, nodes []w.Nodes) (err error) {
	const perPage = 1000
	pages := PageCount(int64(len(nodes)), perPage)
	for currentPage := int64(1); currentPage <= pages; currentPage++ {
		offset := perPage * (currentPage - 1)
		resultOfInsert, err := sqlext.BatchInsert(db.DB(), nodes[offset:offset+perPage])
		fmt.Println("Saving... ", resultOfInsert)
		if err != nil {
			return err
		}
	}
	return nil
}

func setNodes(harvNodes *[]h.Node) ([]w.Nodes) {
	webNodes := make([]w.Nodes, 0, len(*harvNodes))
	for _, harvNode := range *harvNodes {
		webNode := w.Nodes{
			ResourceID: harvNode.ResourceID,
			PageID: harvNode.PageID,
			RankID: getRankID(harvNode.Rank),
			ParentID: harvNode.ParentID,
			ScientificName: harvNode.ScientificName.Verbatim,
			CanonicalForm: harvNode.ScientificName.Canonical,
			ResourcePk: harvNode.ResourcePk,
			SourceUrl: getSourceUrl(harvNode.ID),
			IsHidden: false,
			InUnmappedArea: harvNode.InUnmappedArea,
			ChildrenCount: 0, // TODO...
			CreatedAt: harvNode.CreatedAt,
			UpdatedAt: harvNode.UpdatedAt,
			HasBreadcrumb: (harvNode.Landmark > 0),
			ParentResourcePk: harvNode.ParentResourcePk,
			Landmark: harvNode.Landmark,
		}
		webNodes = append(webNodes, webNode)
	}
	return webNodes
}

func getRankID(rank string) (id uint) {
	// TODO
	return uint(len(rank))
}

func getSourceUrl(id uint) (url string) {
	// TODO
	return fmt.Sprintf("%v", id)
}

// func bulkInsertNodes(db *gorm.DB, nodes []w.Node) {
	// sets := "SET autocommit=0; SET unique_checks=0; SET foreign_key_checks=0;"
	// insert := "INSERT INTO nodes(resource_id, page_id, rank_id, parent_id, scientific_name, canonical_form, "+
	// 			  	"resource_pk, source_url, is_hidden, in_unmapped_area, children_count, created_at, updated_at, "+
	// 			  	"has_breadcrumb, parent_resource_pk, landmark)"
	// values := "VALUES"
	// valList := make([]string, 0, len(nodes))
	// for node := range nodes {
	// 	nodeVals := []string
	// 	nodeVals = append(nodeVals, node.ResourceID, node.PageID, node.RankID, rank.ParentID, rank.ScientificName)
	// 	nodeVals = append(nodeVals, node.CanonicalForm, node.ResourcePk, node.SourceUrl, rank.IsHidden, rank.InUnmappedArea)
	// 	nodeVals = append(nodeVals, node.ChildrenCount, node.CreatedAt, node.UpdatedAt, rank.HasBreadcrumb, rank.ParentResourcePk)
	// 	val_list = append(val_list, "("+strings.Join(nodeVals, ",")+")")
	// }
// }

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
