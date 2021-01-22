package main

import pageable "github.com/BillSJC/gorm-pageable"

func getResultSet(page int, rowsPerPage int) (*pageable.Response, error) {
	resultSet := make([]*UserInfo, 0, 30)
	handler := DB.Model(&UserInfo{}).
		Where(&UserInfo{Age: 1})
	resp, err := pageable.PageQuery(page, rowsPerPage, handler, &resultSet)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
