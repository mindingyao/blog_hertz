package pack

import (
	"blog_hertz/biz/model"
	"blog_hertz/biz/model/tag"
)

func Tags(models []*model.Tag) []*tag.Tag {
	tags := make([]*tag.Tag, 0, len(models))
	for _, m := range models {
		if t := Tag(m); t != nil {
			tags = append(tags, t)
		}
	}
	return tags
}

func Tag(model *model.Tag) *tag.Tag {
	if model == nil {
		return nil
	}

	return &tag.Tag{
		TagID: int32(model.ID),
		Name:  model.Name,
		State: int8(model.State),
	}
}
