package syncmap

import (
	"context"
	"errors"
	"sync"

	"github.com/rlawnsxo131/madre-server-v2/constants"
)

func GenerateHttpContext(parent context.Context) context.Context {
	ctx := context.WithValue(
		parent,
		constants.Key_HttpSyncMap,
		&sync.Map{},
	)
	return ctx
}

func GetFromHttpContext(ctx context.Context) (*sync.Map, error) {
	v := ctx.Value(constants.Key_HttpSyncMap)
	syncMap, ok := v.(*sync.Map)

	if ok {
		return syncMap, nil
	}

	return syncMap, errors.New("GetFromHttpcontext: syncMap is not exist")
}

func SetNewValueFromHttpContext(parent context.Context, key string, value interface{}) (context.Context, error) {
	v := parent.Value(constants.Key_HttpSyncMap)
	syncMap, ok := v.(*sync.Map)

	if ok {
		syncMap.Store(key, value)
		ctx := context.WithValue(
			parent,
			constants.Key_HttpSyncMap,
			syncMap,
		)
		return ctx, nil
	}

	return nil, errors.New("SetNewValueFromHttpContext: syncMap is not exist")
}
