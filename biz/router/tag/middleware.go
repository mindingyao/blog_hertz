// Code generated by hertz generator.

package Tag

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _v1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _tagMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{func(c context.Context, ctx *app.RequestContext) {
		fmt.Printf("create tag named %s\n", ctx.Request.Body())
	}}
}

func _createtagMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deleteMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletetagMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _gettaglistMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatetagMw() []app.HandlerFunc {
	// your code...
	return nil
}
