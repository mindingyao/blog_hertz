namespace go tag
namespace py tag
namespace java tag

enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}

enum State {
    Invalid = 0
    Valid = 1
}

struct Tag {
    1: i32 tag_id
    2: string name
    3: i8  state
}

struct TagListRequest {
	1: string  name (api.body="name", api.form="name",api.vd="(len($) > 0 && len($) < 100)")
	2: i8  state (api.body="state", api.form="state,default=1", api.vd="($ == 0||$ == 1)")
   3: i64 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   4: i64 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct TagListResponse{
   1: Code code
   2: string msg
   3: list<Tag> tags
   4: i64 totoal
}

struct CreateTagRequest {
	1: string  name (api.body="name", api.form="name",api.vd="(len($) > 0 && len($) < 100)")
	2: i8  state (api.body="state", api.default="1", api.form="state", api.vd="($ == 0||$ == 1)")
}

struct CreateTagResponse{
   1: Code code
   2: string msg
}

struct UpdateTagRequest {
    1: i32     tag_id   (api.path="tag_id",api.vd="$>0")
    2: string  name (api.body="name", api.form="name",api.vd="(len($) > 0 && len($) < 100)")
	 3: i8   state (api.body="state", api.form="state", api.vd="($ == 0||$ == 1)")
}

struct UpdateTagResponse{
   1: Code code
   2: string msg
}

struct DeleteTagRequest {
	1: i32    tag_id   (api.path="tag_id",api.vd="$>0")
}

struct DeleteTagResponse{
   1: Code code
   2: string msg
}

service TagService {
   UpdateTagResponse UpdateTag(1:UpdateTagRequest req)(api.post="/v1/tag/update/:tag_id")
   DeleteTagResponse DeleteTag(1:DeleteTagRequest req)(api.post="/v1/tag/delete/:tag_id")
   TagListResponse  GetTagList(1:TagListRequest req)(api.post="/v1/tag/query/")
   CreateTagResponse CreateTag(1:CreateTagRequest req)(api.post="/v1/tag/create/")
}