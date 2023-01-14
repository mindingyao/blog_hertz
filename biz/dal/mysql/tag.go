package mysql

import "blog_hertz/biz/model"

// func CountTag(t *model.Tag) (int64, error) {
// 	var count int64
// 	var db = DB
// 	if t.Name != "" {
// 		db = db.Where("name = ?", t.Name)
// 	}
// 	db = db.Where("state = ?", t.State)
// 	if err := db.Model(t).Count(&count).Error; err != nil {
// 		return 0, err
// 	}
// 	return count, nil
// }

func QueryTag(name string, state int8, page, pageSize int64) ([]*model.Tag, int64, error) {
	db := DB.Model(&model.Tag{})
	if name != "" {
		db = db.Where("name = ?", name)
	}
	db = db.Where("state = ?", state)

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var res []*model.Tag
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, total, nil
}

func CreateTag(tags []*model.Tag) error {
	return DB.Create(tags).Error
}

func DeleteTag(tagId int32) error {
	return DB.Where("id = ?", tagId).Delete(&model.Tag{}).Error
}

func UpdateTag(tag *model.Tag) error {
	return DB.Updates(tag).Error
}
