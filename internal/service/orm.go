package service

import (
	"context"
	"fmt"
	"orm-demo/internal/model"
	"strconv"
)

func (s *Service) Retrieve(c context.Context, req *model.RetrieveReq) (res []map[string]interface{}, err error) {
	if list := req.List; list != nil {
		res = []map[string]interface{}{}
		for tableName, cons := range list {
			db := s.dao.DB().Table(tableName)
			for condition, value := range cons {
				switch condition {
				case "where":
					db.Where(value)
				case "order":
					db.Order(value)
				case "select":
					db.Select(value)
				}
			}
			var size, page int
			sizeStr, ok := cons["size"]
			if !ok {
				size = 10
			} else {
				size, err = strconv.Atoi(sizeStr)
				if err != nil {
					return
				}
			}
			pageStr, ok := cons["page"]
			if !ok {
				page = 1
			} else {
				page, err = strconv.Atoi(pageStr)
				if err != nil {
					return
				}
			}
			rows, _ := db.Limit(size).Offset(page*size - size).Rows()
			columns, _ := rows.Columns()
			columnLength := len(columns)
			for rows.Next() {
				cache := make([]interface{}, columnLength) //临时存储每行数据
				for index, _ := range cache { //为每一列初始化一个指针
					var a interface{}
					cache[index] = &a
				}
				//var a, b int
				//var c, d string
				//cache[0] = &a
				//cache[1] = &b
				//cache[2] = &c
				//cache[3] = &d
				_ = rows.Scan(cache...)
				item := make(map[string]interface{})
				for i, data := range cache {
					typ := fmt.Sprintf("%T, %+v", data, *data.(*interface{}))
					fmt.Println(typ)
					item[columns[i]] = fmt.Sprintf("%s", *data.(*interface{}))//取实际类型
				}
				res = append(res, item)
			}
			_ = rows.Close()
			return res, nil
		}
	}
	return
}
