package visualizer

import (
	"fmt"

	"github.com/shellyln/go-open-soql-parser/soql/parser/types"
)

func Visualize(q *types.SoqlQuery) string {
	relations := "erDiagram\n\n"

	for viewId := range q.Meta.ViewGraph {
		leaf := q.Meta.ViewGraph[viewId]

		if leaf.ParentViewId == 0 {
			qLeaf := q.Meta.QueryGraph[leaf.QueryId]
			if qLeaf.ParentQueryId != 0 {
				parentQLeaf := q.Meta.QueryGraph[qLeaf.ParentQueryId]
				parentQuery := parentQLeaf.Query

				parentName := parentQuery.From[0].Name
				parentColName := ""
				fields := leaf.Object.PerObjectQuery.Fields

				namesFound := false
				collectNames := func(conditions []types.SoqlCondition) {
					for i := range conditions {
						if conditions[i].Opcode == types.SoqlConditionOpcode_FieldInfo &&
							conditions[i].Value.Type == types.SoqlFieldInfo_SubQuery {

							sq := conditions[i].Value.SubQuery
							if sq == leaf.Query {
								if i > 0 && conditions[i-1].Opcode == types.SoqlConditionOpcode_FieldInfo &&
									conditions[i-1].Value.Type == types.SoqlFieldInfo_Field {

									nm := conditions[i-1].Value.Name
									parentNameLen := len(nm)
									if parentNameLen > 0 {
										parentName = nm[:parentNameLen-1]
										parentColName = nm[parentNameLen-1] + " "
										namesFound = true
									}
								}
							}
						}
					}
				}

				collectNames(parentQuery.PostProcessWhere)
				if !namesFound {
					collectNames(parentQuery.Having)
				}

				subqColName := ""
				if len(fields) > 0 && len(fields[0].Name) > 0 {
					subqColName = fields[0].Name[len(fields[0].Name)-1]
				}

				relations += fmt.Sprintf("%v ||..o{ %v: \"%vin %v\"\n", parentName[len(parentName)-1], leaf.Name, parentColName, subqColName)
			}
			continue
		}
		parentLeaf := q.Meta.ViewGraph[leaf.ParentViewId]

		var format string
		var lineType string
		if leaf.NonResult {
			lineType = ".."
		} else {
			lineType = "--"
		}

		if leaf.Many {
			// one to zero or many (subquery root)
			format = "%v ||%vo{ %v: \"subquery\"\n"
		} else if leaf.InnerJoin {
			// one to one
			format = "%v ||%v|| %v: \"\"\n"
		} else {
			// one to zero or one
			format = "%v ||%vo| %v: \"\"\n"
		}
		relations += fmt.Sprintf(format, parentLeaf.Name, lineType, leaf.Name)
	}

	relations += "\n"

	fieldMap := make(map[string]struct{})

	for viewId := range q.Meta.ViewGraph {
		leaf := q.Meta.ViewGraph[viewId]
		relations += fmt.Sprintf("%v {\n", leaf.Name)

		fields := leaf.Object.PerObjectQuery.Fields
		for j := range fields {
			if _, ok := fieldMap[fields[j].Key]; ok {
				continue
			}
			name := fields[j].Name[len(fields[j].Name)-1]
			relations += fmt.Sprintf("    unknown %v\n", name)
			fieldMap[fields[j].Key] = struct{}{}
		}

		relations += "}\n\n"
	}

	return relations
}
