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
)

func main() {
	harvDB, err := gorm.Open("mysql", "root:@/harvester_development?charset=utf8&parseTime=True&loc=Local")
	defer harvDB.Close()
	if err != nil { fmt.Println(err) }
	webDB, err := gorm.Open("mysql", "root:@/eol_development?charset=utf8&parseTime=True&loc=Local")
	webDB.SingularTable(false) // I should not need this, but Node wasn't working. :S
	defer webDB.Close()
	if err != nil { fmt.Println(err) }
	resource := h.Resource{}
	harvDB.Where("id = ?", 1).First(&resource)
	fmt.Println(resource.ID, "->", resource.Name)
	var nodes []h.Node

	nodes, err = loadNodes(harvDB, resource)
	if err != nil { panic(err.Error()) }
	fmt.Println("Found and loaded", len(nodes), "Nodes... Last 10:")
	showNames(nodes[len(nodes)-10:])
	// ancestors, err := loadAncestors(harvDB, resource)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println("Loaded", len(ancestors), "Ancestors.")
	// TODO: load ancestors
	// TODO: load identifiers

	tStart := time.Now()
	webNodes := convNodes(&nodes)
	delta := time.Now().Sub(tStart)
	fmt.Println("Converted", len(webNodes), "Nodes. Took ", delta.Seconds())
	tStart = time.Now()

	err = saveAll(webDB, webNodes, resource)
	if err != nil { panic(err.Error()) }
	delta = time.Now().Sub(tStart)
	tStart = time.Now()
	fmt.Println("Wrote them to the DB,", delta.Seconds())

	err = saveAllAssocs(webDB, webNodes, resource)
	if err != nil { panic(err.Error()) }
	delta = time.Now().Sub(tStart)
	fmt.Println("Wrote all assocs to the DB,", delta.Seconds())
}

func loadNodes(db *gorm.DB, resource h.Resource) (nodeSlice []h.Node, err error) {
	tStart := time.Now()
	var count int64
	db.Model(&h.NodeAncestor{}).Where("resource_id = ?", resource.ID).Count(&count)
	nodesAll := make([]h.Node, 0, count)
	var nodeBuff []h.Node
	const perPage = 100 // TODO - CHAAAAAAAAAAAAAAAANGE THIIIIIIIIIIS SIIIIIIIIIIIIZZZZZZZZE!
	pages := PageCount(count, perPage)
	for currentPage := int64(1); currentPage <= pages; currentPage++ {
		offset := perPage * (currentPage - 1)
		fmt.Println("Reading page", currentPage, "of", pages, "with offset of", offset)
		// db.Limit(perPage).Offset(offset).Model(&resource).Preload("ScientificName").Preload("NodeAncestors").Related(&nodeBuff)
		db.Limit(perPage).Offset(offset).Model(&resource).Preload("ScientificName").Related(&nodeBuff)
		// TODO: err-checking here. ...There is none in the documentation. :S
		nodesAll = append(nodesAll, nodeBuff...)
		showNames(nodesAll[offset+1 : offset+2])
		if currentPage >= 3 { break }
	}
	fmt.Println("LoadNodes() t=", time.Now().Sub(tStart).Seconds())
	return nodesAll[:], nil
}

// func loadAncestors(db *gorm.DB, resource h.Resource) (ancestorSlice []h.NodeAncestor, err error) {
// 	tStart := time.Now()
// 	var count int64
// 	db.Model(&h.NodeAncestor{}).Where("resource_id = ?", resource.ID).Count(&count)
// 	ancestorsAll := make([]h.NodeAncestor, 0, count)
// 	var ancestorBuff []h.NodeAncestor
// 	const perPage = 1000
// 	pages := PageCount(count, perPage)
// 	fmt.Println("PageCount:", pages)
// 	for currentPage := int64(1); currentPage <= pages; currentPage++ {
// 		offset := perPage * (currentPage - 1)
// 		fmt.Println("Reading NodeAncestor page", currentPage, "of", pages, "with offset of", offset)
// 		db.Limit(perPage).Offset(offset).Model(&resource).Preload("ScientificName").Related(&ancestorBuff)
// 		// TODO: err-checking here. ...There is none in the documentation. :S
// 		ancestorsAll = append(ancestorsAll, ancestorBuff...)
// 	}
// 	fmt.Println("LoadAncestors() t=", time.Now().Sub(tStart).Seconds())
// 	return ancestorsAll[:], nil
// }

// TODO: figure out which interface to use for models, argh:
func saveAll(db *gorm.DB, models []w.Node, resource h.Resource) (err error) {
	transact := db.Begin()
	// Curious that there's no error returned from this method... it returns the DB, so you can chain...
  transact.Exec("SET innodb_flush_log_at_trx_commit = 0")
	transact.Debug().Where("resource_id = ?", resource.ID).Delete(w.Node{})
	fmt.Println("Saving", len(models))
	for _, model := range models {
		if err = transact.Debug().Create(&model).Error; err != nil {
			transact.Rollback()
			panic(err.Error())
		}
	}
	transact.Commit()
	return nil
}

func saveAllAssocs(db *gorm.DB, models []w.Node, resource h.Resource) (err error) {
	transact := db.Begin()
	// Curious that there's no error returned from this method... it returns the DB, so you can chain...
  transact.Exec("SET innodb_flush_log_at_trx_commit = 0")
	transact.Debug().Where("resource_id = ?", resource.ID).Delete(w.NodeAncestor{})
	transact.Debug().Where("resource_id = ?", resource.ID).Delete(w.ScientificName{})
	fmt.Println("Saving", len(models))
	for i := range models {
		for _, assoc := range models[i].NodeAncestors {
			if err = transact.Debug().Create(&assoc).Error; err != nil {
				transact.Rollback()
				panic(err.Error())
			}
		}
		if err = transact.Debug().Create(&models[i].ScientificName).Error; err != nil {
			transact.Rollback()
			panic(err.Error())
		}
	}
	transact.Commit()
	return nil
}

func convNodes(harvNodes *[]h.Node) ([]w.Node) {
	webNodes := make([]w.Node, 0, len(*harvNodes))
	for _, harvNode := range *harvNodes {
		webNode := w.Node{
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
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			HasBreadcrumb: (harvNode.Landmark > 0),
			ParentResourcePk: harvNode.ParentResourcePk,
			Landmark: harvNode.Landmark, // enum: no_landmark minimal abbreviated extended full
			ScientificNames: []w.ScientificName{convSciName(harvNode.ScientificName, true)},
			NodeAncestors: []
		}
		webNodes = append(webNodes, webNode)
	}
	return webNodes
}

json.canonical_form name.canonical_italicized

func convSciName(harvName h.ScientificName, preferred bool) ([]w.ScientificName) {
	def attribution_html
		webName := w.ScientificName{
			// YOU WERE HERE ... converting all of these.... fun-fun.
			PageID               harvName.PageID,
			Italicized           harvName.Italicized,
			CanonicalForm        getCanonicalItal(harvName),
			TaxonomicStatusID    getTaxonomicStatus(harvName.TaxonomicStatusID),
			IsPreferred          preferred,
			CreatedAt            time.Now(),
			UpdatedAt            time.Now(),
			ResourceID           harvNode.ResourceID,
			NodeResourcePk       harvName.NodeResourcePk,
			// TODO: SourceReference      string,
			Genus                harvName.Genus,
			SpecificEpithet      harvName.SpecificEpithet,
			InfraspecificEpithet harvName.InfraspecificEpithet,
			InfragenericEpithet  harvName.InfragenericEpithet,
			Uninomial            harvName.Uninomial,
			Verbatim             harvName.Verbatim,
			Authorship           harvName.Authorship,
			Publication          harvName.Publication,
			Remarks              harvName.Remarks,
			ParseQuality         harvName.ParseQuality,
			Year                 harvName.Year,
			Hybrid               harvName.Hybrid,
			Surrogate            harvName.Surrogate,
			Virus                harvName.Virus,
			Attribution          harvName.AttributionHTML,
		}
		webNames = append(webNames, webName)
	return webNames
}

func getRankID(rank string) (id uint) {
	// TODO
	return uint(len(rank))
}

func getSourceUrl(id uint) (url string) {
	// TODO
	return fmt.Sprintf("%v", id)
}

func getTaxonomicStatus(id uint) (uint) {
	// TODO
	return id
}

func getTaxonomicStatus(name h.ScientificName) (string) {
	// TODO: implement
	// def canonical_italicized
	// 	italicize(canonical)
	// end
	//
	// def italicize(name)
	// 	name = verbatim if name.blank?
	// 	name.gsub!(/\s+/, ' ') # This is just aesthetic cleanup.
	// 	name = name.sub(genus, "<i>#{genus}</i>") if genus
	// 	name = name.sub(specific_epithet, "<i>#{specific_epithet}</i>") if specific_epithet
	// 	name = name.sub(infraspecific_epithet, "<i>#{infraspecific_epithet}</i>") if infraspecific_epithet
	// 	name = name.sub(infrageneric_epithet, "<i>#{infrageneric_epithet}</i>") if infrageneric_epithet
	// 	name.gsub('</i> <i>', ' ') # This is just aesthetic cleanup.
	// end

	return name.Canonical
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
	if result-floor > 0 { return int64(floor + 1) }
	return int64(floor)
}

func showNames(nodes []h.Node) {
	for i := range nodes[:] {
		ancestor := "(no ancestors)"
		if len(nodes[i].NodeAncestors) > 0 { ancestor = nodes[i].NodeAncestors[0].AncestorFk }
		fmt.Println("Node", nodes[i].ID, "-> '"+nodes[i].ScientificName.Verbatim+"' ("+nodes[i].Canonical+")", ancestor)
	}
	return
}
