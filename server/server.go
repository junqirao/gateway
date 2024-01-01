package server

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/component/registry"
	"github.com/junqirao/gateway/lib/response"
	"github.com/junqirao/gateway/model"
)

// ListServerInfo server info
func ListServerInfo(ctx context.Context) (infos []*model.ServerInfo, err error) {
	infos = make([]*model.ServerInfo, 0)

	cfgMap, err := registry.Instance().Get(ctx, registry.ServerConfigRegPath())
	if err != nil {
		return nil, err
	}

	for name, cfgStr := range cfgMap {
		cfg := new(model.ServerConfig)
		if err = json.Unmarshal([]byte(cfgStr), &cfg); err != nil {
			g.Log().Warningf(ctx, "unmarshal server config [%s](value=%s) failed: %v", name, cfgStr, err)
			continue
		}

		infos = append(infos, &model.ServerInfo{
			Name:         name,
			ServerConfig: *cfg,
		})
	}

	return
}

// GetServerInfo server info
func GetServerInfo(ctx context.Context, name string) (rsp interface{}, err error) {
	cfg := new(model.ServerConfig)
	if err = getRegistryOne(ctx, registry.ServerConfigRegKey(name), &cfg); err != nil {
		return
	}
	rsp = &model.ServerInfo{
		Name:         name,
		ServerConfig: *cfg,
	}
	return
}

func getRegistryOne(ctx context.Context, key string, ptr interface{}) (err error) {
	vs, err := registry.Instance().Get(ctx, key)
	if err != nil {
		return
	}

	has := false
	for n, s := range vs {
		if n != key {
			continue
		}
		has = true
		if ptr != nil {
			err = json.Unmarshal([]byte(s), &ptr)
		}
		break
	}
	if !has {
		err = response.ErrorResourceNotFound
	}
	return
}

// DeleteConfig ...
func DeleteConfig(ctx context.Context, name string) (rsp interface{}, err error) {
	key := registry.ServerConfigRegKey(name)
	if err = getRegistryOne(ctx, key, nil); err != nil {
		return
	}

	return name, registry.Instance().Delete(ctx, key)
}

// SetConfig ...
func SetConfig(ctx context.Context, name string, config *model.ServerConfig) (rsp interface{}, err error) {
	key := registry.ServerConfigRegKey(name)

	if config.Properties == nil {
		// update enabled only
		raw := new(model.ServerConfig)
		if err = getRegistryOne(ctx, key, &raw); err != nil {
			return
		}
		raw.Enabled = config.Enabled
		config = raw
	} else {
		config.FillDefault()
	}
	return name, registry.Instance().Set(ctx, key, config)
}
